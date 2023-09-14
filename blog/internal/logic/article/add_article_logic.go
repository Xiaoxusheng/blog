package article

import (
	"blog/internal/config"
	"blog/internal/models"
	"blog/internal/svc"
	"blog/internal/types"
	"blog/internal/utils"
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddArticleLogic {
	return &AddArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddArticleLogic) AddArticle(req *types.AddArticleReq) (resp *types.AddArticleResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.AddArticleResp)
	err = models.InsertArticle(models.Article{
		Identity:           utils.GetUuid(),
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
	resp.Msg = "新增文章成功！"
	resp.Code = http.StatusOK
	return resp, nil
}
