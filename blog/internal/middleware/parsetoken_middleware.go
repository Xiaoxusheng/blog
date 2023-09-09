package middleware

import (
	"blog/internal/utils"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

type ParseTokenMiddleware struct {
}

func NewParseTokenMiddleware() *ParseTokenMiddleware {
	return &ParseTokenMiddleware{}
}

func (m *ParseTokenMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		// Token from another example.  This token is expired
		tokens := r.Header.Get("Authorization")
		if tokens == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("token不能为空"))
			return
		}
		//fmt.Println(strings.Contains(tokens, "Bearer"))
		if !strings.Contains(tokens, "Bearer") {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("token格式不对"))
			return
		}
		tokenString := strings.Split(tokens, "Bearer ")
		fmt.Println(tokenString)
		user := utils.MyCustomClaims{}
		token, err := jwt.ParseWithClaims(tokenString[len(tokenString)-1], &user, func(token *jwt.Token) (interface{}, error) {
			return utils.MySigningKey, nil
		})
		if token.Valid {
			fmt.Println("验证通过")
		} else if errors.Is(err, jwt.ErrTokenMalformed) {
			w.WriteHeader(http.StatusBadRequest)

			fmt.Println("That's not even a token")
			w.Write([]byte("这不是一个token!"))
			return
		} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
			// Invalid signature
			fmt.Println("Invalid signature")
			w.WriteHeader(http.StatusBadRequest)

			w.Write([]byte("无效的签名!"))

			return
		} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
			// Token is either expired or not active yet
			fmt.Println("Timing is everything")
			w.WriteHeader(http.StatusBadRequest)

			w.Write([]byte("token失效或过期!"))

			return
		} else {
			fmt.Println("Couldn't handle this token:", err)
			w.WriteHeader(http.StatusBadRequest)

			w.Write([]byte("未知错误!"))

			return
		}
		r.Header.Set("identity", user.Identity)
		//fmt.Println("username", user.Identification)
		//ctx := context.Background()
		//exit, _ := db.Rdb.Exists(ctx, user.Identity).Result()
		//fmt.Println("exit", exit)
		//if exit != 1 {
		//	c.JSON(http.StatusBadRequest, gin.H{
		//		"code": 1,
		//		"msg":  "已经退出登录!",
		//	})
		//	return
		//}
		//
		////fmt.Println(t, "\n", rt)
		//if t != tokenString[len(tokenString)-1] && rt != tokenString[len(tokenString)-1] {
		//	c.JSON(http.StatusOK, gin.H{
		//		"code": 1,
		//		"msg":  "账号在其他地方登录，只允许一台设备登录！",
		//	})
		//
		//	return
		//}
		// Passthrough to next handler if need
		next(w, r)
	}
}
