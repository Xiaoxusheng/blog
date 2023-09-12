package models

import (
	"blog/internal/types"
	"errors"
	"time"
	"unicode"
)

type Talk struct {
	Identity  string    `json:"identity" `
	UserId    string    `json:"userId,omitempty"`
	Content   string    `json:"content,omitempty"`
	Status    int       `json:"status,omitempty"`
	IsTop     int       `json:"isTop,omitempty"`
	LikeTimes int       `json:"likeTimes" json:"likeTimes,omitempty"`
	CreatedAt time.Time ` xorm:"created " json:"createdAt" json:"createdAt"`
	UpdatedAt time.Time ` xorm:"updated" json:"updatedAt" json:"updatedAt"`
}

// 插入说说
func InsertTalk(talk *Talk) error {
	_, err := Engine.Insert(talk)
	if err != nil {
		return errors.New("发表说说失败！")
	}
	return nil

}

func UpdateTalk(id string, content string, status, istop int) error {
	talk := new(Talk)
	if unicode.IsNumber(rune(status)) && status > 0 && status < 3 {
		talk.Status = status
	}
	if unicode.IsNumber(rune(istop)) && istop > 0 && istop < 3 {
		talk.IsTop = istop
	}
	talk.Content = content
	_, err := Engine.Where("identity=?", id).Update(talk)
	if err != nil {
		return errors.New("修改说说失败！")
	}
	return nil
}

func DeleteTalk(id string) error {
	talk := new(Talk)
	talk.Status = 3
	_, err := Engine.Where("identity=?", id).Update(talk)
	if err != nil {
		return errors.New("删除说说失败！")
	}
	return nil
}

func RecoverTalk(id string) error {
	talk := new(Talk)
	talk.Status = 1
	_, err := Engine.Where("identity=?", id).Update(talk)
	if err != nil {
		return errors.New("恢复说说失败！")
	}
	return nil
}

func GeTalkById(id string) error {
	talk := new(Talk)
	ok, err := Engine.Where("identity=?", id).Get(talk)
	if err != nil || !ok {
		return errors.New("说说不存在！")
	}
	return nil
}

func GetTalkList(list []*types.TalkList, limit, offset int) ([]*types.TalkList, error) {
	err := Engine.Where("blog_talk").Limit(limit, offset).OrderBy("created_at", "desc").Find(list)
	if err != nil {
		return nil, errors.New("查询说说列表失败!")
	}
	return list, nil
}
