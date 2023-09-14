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

type UpdateArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateArticleLogic {
	return &UpdateArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateArticleLogic) UpdateArticle(req *types.UpdateArticleReq) (resp *types.UpdateArticleResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.UpdateArticleResp)

	//查询是否存在
	err = models.GetByArticleId(req.Article.Identity)
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.ARTICLE
		return
	}
	err = models.UpdateArticle(models.Article{
		ArticleTitle:       req.Article.ArticleTitle,
		AuthorId:           req.Article.AuthorId,
		CategoryId:         req.Article.CategoryId,
		ArticleContent:     req.Article.ArticleContent,
		ArticleCover:       req.Article.ArticleCover,
		IsTop:              req.Article.IsTop,
		Status:             req.Article.Status,
		Type:               req.Article.Type,
		OriginUrl:          req.Article.OriginUrl,
		ViewTimes:          req.Article.ViewTimes,
		ArticleDescription: req.Article.ArticleDescription,
		ThumbsUpTimes:      req.Article.ThumbsUpTimes,
		ReadingDuration:    req.Article.ReadingDuration,
	})
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.ARTICLE
		return
	}
	resp.Msg = "修改文章成功！"
	resp.Code = http.StatusOK
	return resp, nil
}
