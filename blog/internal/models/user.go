package models

import (
	"log"
	"time"
)

type blogUser struct {
	Identity  string    `json:"identity"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Role      int       `json:"role"`
	Nick_name string    `json:"nick_Name"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time ` xorm:"created " json:"createdAt" `
	UpdatedAt time.Time ` xorm:"updated" json:"updatedAt" `
	Qq        int       `json:"qq"`
	Ip        string    `json:"ip"`
}

func GetByUsePwd(username, password string, role int) *blogUser {
	user := new(blogUser)
	_, err := Engine.Where("username=? and password=? and role=?", username, password, role).Get(user)
	if err != nil {
		log.Println(err)
	}
	return user
}

func GetUser(username string) bool {
	user := new(blogUser)
	ok, err := Engine.Where("username=? ", username).Get(user)
	if err != nil {
		log.Println(err)
	}
	return ok
}

func GetByUsername(username string) bool {
	user := new(blogUser)
	ok, err := Engine.Where("username=?", username).Get(user)
	if err != nil {
		return false
	}
	return ok

}

func InsertUser(identity, username, password, nickName, ip string, qq int) error {
	_, err := Engine.Insert(blogUser{
		Identity:  identity,
		Username:  username,
		Password:  password,
		Role:      2,
		Nick_name: nickName,
		Qq:        qq,
		Ip:        ip,
	})
	if err != nil {
		return err
	}
	return nil
}

func GetById(identity string) (*blogUser, bool) {
	user := new(blogUser)
	ok, err := Engine.Where("identity=?", identity).Get(user)
	if err != nil {
		log.Println(err)
	}
	return user, ok
}

func UpdateRole(id string, role int) error {
	user := new(blogUser)
	user.Role = role
	_, err := Engine.Where("identity=?", id).Update(user)
	if err != nil {
		return err
	}
	return nil
}
