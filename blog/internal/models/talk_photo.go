package models

import (
	"errors"
	"time"
)

type TalkPhoto struct {
	Identity  string    `json:"identity"`
	Url       string    `json:"url,omitempty"`
	TalkId    string    `json:"talk_Id"`
	CreatedAt time.Time ` xorm:"created " json:"createdAt" json:"createdAt"`
	UpdatedAt time.Time ` xorm:"updated" json:"updatedAt" json:"updatedAt"`
}

func InsertTalkPhoto(talkPhoto *TalkPhoto) error {
	_, err := Engine.Insert(talkPhoto)
	if err != nil {
		return errors.New("插入图片失败")
	}
	return nil
}

func UpdateTalkPhoto(id, Imglist string) error {
	talkPhoto := new(TalkPhoto)
	talkPhoto.Url = Imglist
	_, err := Engine.Where("identity=?", id).Update(talkPhoto)
	if err != nil {
		return errors.New("修改图片失败")
	}
	return nil

}
