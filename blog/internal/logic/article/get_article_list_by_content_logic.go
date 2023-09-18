package article

import (
	"blog/internal/models"
	"context"
	"net/http"

	"blog/internal/svc"
	"blog/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleListByContentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetArticleListByContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleListByContentLogic {
	return &GetArticleListByContentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticleListByContentLogic) GetArticleListByContent(req *types.GetArticleListByContentReq) (resp *types.GetArticleListByContentResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.GetArticleListByContentResp)
	user, err := models.GetByContent(req.Content)
	resp.Msg = "获取文章列表成功！"
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
