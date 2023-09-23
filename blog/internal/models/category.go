package models

import (
	"blog/internal/types"
	"errors"
	"time"
)

type Category struct {
	Identity     string    `json:"identity"`
	Status       int       `json:"status"`
	CategoryName string    `json:"categoryName"`
	CreatedAt    time.Time ` xorm:"created " json:"createdAt" `
	UpdatedAt    time.Time ` xorm:"updated" json:"updatedAt" `
}

func GetCategory(name string) error {
	category := new(Category)
	ok, err := Engine.Where("category_name=?", name).Get(category)
	if err != nil || !ok {
		return errors.New("该分类不存在！")
	}
	return nil
}

func InsertCategory(c *Category) error {
	_, err := Engine.Insert(c)
	if err != nil {
		return errors.New("新增分类错误！")
	}
	return nil
}

func GetCategoryById(id string) error {
	category := new(Category)
	ok, err := Engine.Where("identity=?", id).Get(category)
	if err != nil || !ok {
		return errors.New("该分类不存在！")
	}
	return nil
}

func UpdateCategory(id, name string) error {
	_, err := Engine.Where("identity=?", id).Update(&Category{
		CategoryName: name,
	})
	if err != nil {
		return errors.New("修改分类失败！")
	}
	return nil
}

func DeleteCategory(id string) error {
	_, err := Engine.Where("identity=?", id).Update(&Category{
		Status: 1,
	})
	if err != nil {
		return errors.New("删除分类失败！")
	}
	return nil
}

func GetCategoryList(list []*types.CategoryList) error {
	err := Engine.Table("blog_category").Find(&list)
	if err != nil {
		return errors.New("获取分类列表失败！")
	}
	return nil
}

func GetCategoryCount(list map[string]string) error {
	category := make([]Category, 0)
	err := Engine.Table("blog_category").Find(&category)
	if err != nil {
		return errors.New("获取失败！")
	}
	for i := 0; i < len(category); i++ {
		list[category[i].Identity] = category[i].CategoryName
	}
	return nil
}
