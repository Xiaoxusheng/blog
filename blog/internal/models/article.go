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
	ReadingDuration    float64   `json:"readingDuration"`
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

func GetId(id string) error {
	article := new(Article)
	ok, err := Engine.Where("category_id=?", id).Get(article)
	if err != nil || !ok {
		return errors.New("id不存在！")
	}
	return nil
}

func GetByIdArticleList(id string, list []*types.ArticleList) error {
	err := Engine.Table("blog_article").Where("status=? and category_id=?", 1, id).Desc("created_at").Find(list)
	if err != nil {
		return errors.New("获取文章列表失败！")
	}
	return nil
}

func GetRecommendArticleById(id string, list []*types.ArticleList) error {
	err := Engine.Table("blog_article").Where("status=?  ", 1).Desc("created_at").Find(list)
	if err != nil {
		return errors.New("获取文章列表失败！")
	}
	for i := 0; i < len(list); i++ {
		if list[i].Identity == id && i >= 1 && i < len(list)-1 {
			list[0] = list[i-1]
			list[1] = list[i+1]
		} else if list[i].Identity == id && i < 1 {
			list[0] = nil
			list[1] = list[i+1]
		} else if list[i].Identity == id && i >= 1 && i == len(list)-1 {
			list[0] = list[i-1]
			list[1] = nil
		}
	}

	return nil
}

func GetByContent(content string) (*Article, error) {
	article := new(Article)
	ok, err := Engine.Where("content like %=? %", content).Get(article)
	if err != nil || !ok {
		return nil, errors.New("内容不存在！")
	}
	return article, nil
}

func GetHotArticle() (*Article, error) {
	article := new(Article)
	ok, err := Engine.Table("blog_article").Select("MAX(thumbs_up_times)").Get(article)
	if err != nil || !ok {
		return nil, errors.New("热搜不存在！")
	}
	return article, nil
}

func UpdateLike(id string) error {
	article := new(Article)
	article.ThumbsUpTimes += 1

	_, err := Engine.Update(article)
	if err != nil {
		return errors.New("点赞失败！")
	}
	return nil
}

func UpdateUnLike(id string) error {
	article := new(Article)
	article.ThumbsUpTimes -= 1
	_, err := Engine.Where("identity=?", id).Update(article)
	if err != nil {
		return errors.New("点赞失败！")
	}
	return nil
}

func UpdateTime(id string, time float64) error {
	article := new(Article)
	article.ReadingDuration += time
	_, err := Engine.Where("identity=?", id).Update(article)
	if err != nil {
		return errors.New("增加时长失败！")
	}
	return nil
}
