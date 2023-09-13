package models

import (
	"blog/internal/types"
	"errors"
	"time"
)

type Tag struct {
	Identity  string    `json:"identity"`
	TagName   string    `json:"tag_Name"`
	Status    int       `json:"status"`
	CreatedAt time.Time ` xorm:"created " json:"createdAt" `
	UpdatedAt time.Time ` xorm:"updated" json:"updatedAt" `
}

func GetTagName(tagName string) error {
	tag := new(Tag)
	ok, _ := Engine.Where("tag_name=?", tagName).Get(tag)
	if ok {
		return errors.New("tag标签已经存在！")
	}
	return nil
}

func GetTagById(id string) error {
	tag := new(Tag)
	ok, err := Engine.Where("identity=?", id).Get(tag)
	if err != nil || !ok {
		return errors.New("tag标签不存在！")
	}
	return nil
}

func InsertTag(id, tagName string) error {
	_, err := Engine.Insert(&Tag{Identity: id, TagName: tagName, Status: 0})
	if err != nil {
		return errors.New("新增标签失败！")
	}
	return nil
}

func UpdateTag(id, TagName string) error {
	_, err := Engine.Where("identity=?", id).Update(&Tag{
		TagName: TagName,
	})
	if err != nil {
		return errors.New("修改tag标签失败")
	}
	return nil
}

func DeleteTag(id string) error {
	_, err := Engine.Where("identity=?", id).Update(&Tag{
		Status: 1,
	})
	if err != nil {
		return errors.New("删除tag标签失败")
	}
	return nil
}

func GetTagList(list []*types.TagList, offset, limit int) error {
	err := Engine.Table("blog_tag").Limit(limit, (offset-1)*limit).Find(&list)
	if err != nil {
		return errors.New("获取tag列表标签失败")
	}
	return nil
}
