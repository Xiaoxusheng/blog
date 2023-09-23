package models

import (
	"errors"
	"time"
)

type Comment struct {
	Identity   string    `json:"identity"`
	ParentId   string    `json:"parentId"`
	IsRead     bool      `json:"isRead"`
	Type       int       `json:"type"`
	ThumbsUp   int       `json:"thumbs_Up"`
	FromId     string    `json:"fromId"`
	FromName   string    `json:"fromName"`
	FromAvatar string    `json:"fromAvatar"`
	ToId       string    `json:"toId"`
	ToName     string    `json:"toName"`
	ToAvatar   string    `json:"toAvatar"`
	Content    string    `json:"content"`
	CreatedAt  time.Time ` xorm:"created " json:"createdAt" `
	UpdatedAt  time.Time ` xorm:"updated" json:"updatedAt" `
	Ip         string    `json:"ip"`
	Status     int       `json:"status"`
}

func GetByCommentId(id string) error {
	comment := new(Comment)
	ok, err := Engine.Where("identity=?", id).Get(comment)
	if err != nil || !ok {
		return errors.New("内容不存在！")
	}
	return nil
}

func InsertComment(c *Comment) error {
	_, err := Engine.Insert(c)
	if err != nil {
		return errors.New("发表评论失败！")
	}
	return nil
}

func DeleteComment(id string) error {
	comment := new(Comment)
	ok, err := Engine.Where("identity=?", id).Get(comment)
	if err != nil {
		return errors.New("删除失败！")
	}
	if ok && comment.ParentId != "" {
		comment.Status = 1
		Engine.Where("identity=?", id).Update(comment)
	} else {
		_, err := Engine.Where("identity=?", id).Update(&Comment{
			Status: 1,
		})
		if err != nil {
			return errors.New("删除失败！")
		}
		_, err = Engine.Where("identity=?", comment.ParentId).Update(&Comment{
			Status: 1,
		})
		if err != nil {
			return errors.New("删除失败！")
		}
	}
	return nil
}

// 回复评论
func InsertAppleComment(c Comment) error {
	_, err := Engine.Insert(c)
	if err != nil {
		return errors.New("回复失败！")
	}
	return nil
}
