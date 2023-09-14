package models

import (
	"errors"
	"time"
)

type Article struct {
	Identity           string    `json:"identity"`
	ArticleTitle       string    `json:"articleTitle"`
	AuthorId           int       `json:"authorId"`
	CategoryId         int       `json:"categoryId"`
	ArticleContent     string    `json:"articleContent"`
	ArticleCover       string    `json:"articleCover"`
	IsTop              int       `json:"isTop"`
	Status             int       `json:"status"`
	Type               int       `json:"type"`
	OriginUrl          string    `json:"originUrl"`
	ViewTimes          int       `json:"viewTimes"`
	ArticleDescription string    `json:"articleDescription"`
	ThumbsUpTimes      int       `json:"thumbsUpTimes"`
	ReadingDuration    int       `json:"readingDuration"`
	CreatedAt          time.Time ` xorm:"created " json:"createdAt" `
	UpdatedAt          time.Time ` xorm:"updated" json:"updatedAt" `
}

func InsertArticle(article Article) error {
	_, err := Engine.Insert(&Article{
		Identity:           article.Identity,
		ArticleTitle:       article.ArticleTitle,
		AuthorId:           article.AuthorId,
		CategoryId:         article.CategoryId,
		ArticleContent:     article.ArticleContent,
		ArticleCover:       article.ArticleCover,
		IsTop:              article.IsTop,
		Status:             article.Status,
		Type:               article.Type,
		OriginUrl:          article.OriginUrl,
		ViewTimes:          article.ViewTimes,
		ArticleDescription: article.ArticleDescription,
		ThumbsUpTimes:      article.ThumbsUpTimes,
		ReadingDuration:    article.ReadingDuration,
	})
	if err != nil {
		return errors.New("新增文章失败！")
	}
	return nil
}

// 修改文章内容
func UpdateArticle(article Article) error {
	_, err := Engine.Where("identity=?", article.Identity).Update(Article{
		ArticleTitle:       article.ArticleTitle,
		AuthorId:           article.AuthorId,
		CategoryId:         article.CategoryId,
		ArticleContent:     article.ArticleContent,
		ArticleCover:       article.ArticleCover,
		IsTop:              article.IsTop,
		Status:             article.Status,
		Type:               article.Type,
		OriginUrl:          article.OriginUrl,
		ViewTimes:          article.ViewTimes,
		ArticleDescription: article.ArticleDescription,
		ThumbsUpTimes:      article.ThumbsUpTimes,
		ReadingDuration:    article.ReadingDuration,
	})
	if err != nil {
		return errors.New("修改文章失败！")
	}
	return nil
}

// 查询是否存在
func GetByArticleId(id string) error {
	article := new(Article)
	ok, err := Engine.Where("identity=?", id).Get(article)
	if err != nil || !ok {
		return errors.New("文章不存在！")
	}
	return nil
}
