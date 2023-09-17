package models

import (
	"blog/internal/types"
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
	if article.Status > 3 || article.Status < 0 || article.IsTop > 2 || article.IsTop < 0 || article.Type < 0 || article.Type > 2 {
		return errors.New("非法操作！")
	}
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
func GetByArticle(id string) *Article {
	article := new(Article)
	ok, err := Engine.Where("identity=?", id).Get(article)
	if err != nil || !ok {
		return nil
	}
	return article
}

func GetArticleList(list []*types.ArticleList, limit, offset int) error {
	err := Engine.Table("blog_article").Where("status=?", 1).Desc("thumbs_up_times").OrderBy("created_at", "desc").Limit(limit, (offset-1)*limit).Find(list)
	if err != nil {
		return errors.New("获取文章列表失败！")
	}
	return nil
}

func GetCreateArticleList(list []*types.ArticleList, limit, offset int) error {
	err := Engine.Table("blog_article").Where("status=?", 1).Desc("created_at").Limit(limit, (offset-1)*limit).Find(list)
	if err != nil {
		return errors.New("获取文章列表失败！")
	}
	return nil

}

func GetByTagArticleList(list []*types.ArticleList, tag string) error {
	err := Engine.Table("blog_article").Where("status=? and tag=?", 1, tag).Desc("created_at").Find(list)
	if err != nil {
		return errors.New("获取文章列表失败！")
	}
	return nil
}
