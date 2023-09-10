package utils

import (
	"crypto/md5"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	uuid "github.com/satori/go.uuid"
	"time"
)

var MySigningKey = []byte("blog by go-zero in 2023-09-09 14:14:02")

type MyCustomClaims struct {
	Identity string `json:"identity"`
	jwt.RegisteredClaims
}

// 生成token
func CreateToken(identity string) string {

	// Create claims with multiple fields populated
	claims := MyCustomClaims{
		identity,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Subject:   "somebody",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(MySigningKey)
	fmt.Printf("%v %v", ss, err)
	return ss
}

// 生成MD5
func GetMd5(password string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(password)))
}

// uuid生成
func GetUuid() string {
	return uuid.NewV4().String()
}

// 验证
func Verify(s string) bool {
	if len(s) < 1 || len(s) > 6 {
		return false
	}
	return true
}
