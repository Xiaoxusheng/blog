package article

import (
	"blog/internal/config"
	"blog/internal/models"
	"context"
	"net/http"

	"blog/internal/svc"
	"blog/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHotArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetHotArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHotArticleLogic {
	return &GetHotArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetHotArticleLogic) GetHotArticle(req *types.GetHotArticleReq) (resp *types.GetHotArticleResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.GetHotArticleResp)
	//获取热门文章
	user, err := models.GetHotArticle()
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.ARTICLE
		return
	}
	resp.Msg = "获取热搜文章成功！"
	resp.Code = http.StatusOK
	resp.Data = map[string]types.ArticleList{
		"article_list": {
			Identity:           user.Identity,
			ArticleTitle:       user.ArticleTitle,
			AuthorId:           user.AuthorId,
			CategoryId:         user.CategoryId,
			ArticleContent:     user.ArticleContent,
			ArticleCover:       user.ArticleCover,
			IsTop:              user.IsTop,
			Status:             user.Status,
			Type:               user.Type,
			OriginUrl:          user.OriginUrl,
			ViewTimes:          user.ViewTimes,
			ArticleDescription: user.ArticleDescription,
			ThumbsUpTimes:      user.ThumbsUpTimes,
			ReadingDuration:    user.ReadingDuration,
		},
	}
	return resp, nil
}
