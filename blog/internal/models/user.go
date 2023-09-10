package models

import (
	"blog/internal/types"
	"errors"
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

//func GetUser(username string) bool {
//	user := new(blogUser)
//	ok, err := Engine.Where("username=? ", username).Get(user)
//	if err != nil {
//		log.Println(err)
//	}
//	return ok
//}

func GetByUsername(username string) (*blogUser, bool) {
	user := new(blogUser)
	ok, err := Engine.Where("username=?", username).Get(user)
	if err != nil {
		return nil, false
	}
	return user, ok
}

// 新增用户
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

func InsertAdminUser(identity, username, password, nickName, ip string, qq int) error {
	_, err := Engine.Insert(blogUser{
		Identity:  identity,
		Username:  username,
		Password:  password,
		Role:      9999,
		Nick_name: nickName,
		Qq:        qq,
		Ip:        ip,
	})
	if err != nil {
		return err
	}
	return nil
}

// 根据identity获取信息
func GetById(id string) (*blogUser, bool) {
	user := new(blogUser)
	ok, err := Engine.Where("identity=?", id).Get(user)
	if err != nil {
		log.Println(err)
	}
	return user, ok
}

// 管理员修改权限
func UpdateRole(id string, role int) error {
	user := new(blogUser)
	user.Role = role
	_, err := Engine.Where("identity=?", id).Update(user)
	if err != nil {
		return err
	}
	return nil
}

// 修改密码
func UpdatePwd(id, pwd string) error {
	user := new(blogUser)
	user.Password = pwd
	_, err := Engine.Where("identity=?", id).Update(user)
	if err != nil {
		return errors.New("修改失败!")
	}
	return nil
}

// 判断是否有权限
func GetByIdentity(id string) (bool, error) {
	user := new(blogUser)
	ok, err := Engine.Where("identity=?", id).Get(user)
	if err != nil || !ok {
		return false, errors.New("用户不存在！")
	}
	if user.Role == 9999 {
		return ok, nil
	}
	return false, errors.New("权限不足！")

}

// 获取用户列表
func GetUserList(id string, page, limit int, Userlist []*types.Userlist) ([]*types.Userlist, error) {

	err := Engine.Table("blog_user").Where("identity!=?", id).Limit(limit, (page-1)*limit).Find(&Userlist)
	if err != nil {
		return Userlist, err
	}
	return Userlist, nil

}

// 修改ip
func UpdateIp(id, ip string) error {
	user := new(blogUser)
	user.Ip = ip
	_, err := Engine.Where("identity=?", id).Update(user)
	if err != nil {
		return errors.New("ip修改失败！")
	}
	return nil
}

// 更新用户信息
func UpdateUserById(id string, req *types.UpdateUserInfoReq) error {
	use := new(blogUser)
	if req.QQ != 0 {
		use.Qq = req.QQ
	}
	if req.NickName != "" {
		use.Nick_name = req.NickName
	}
	if req.Avatar != "" {
		use.Avatar = req.Avatar
	}
	_, err := Engine.Where("identity=?", id).Update(use)
	if err != nil {
		return errors.New("更新失败！")
	}
	return nil
}
