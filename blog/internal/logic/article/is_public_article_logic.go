package article

import (
	"blog/internal/config"
	"blog/internal/models"
	"context"
	"errors"
	"net/http"
	"unicode"

	"blog/internal/svc"
	"blog/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsPublicArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIsPublicArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsPublicArticleLogic {
	return &IsPublicArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IsPublicArticleLogic) IsPublicArticle(req *types.IsPublicArticleReq) (resp *types.IsPublicArticleResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.IsPublicArticleResp)
	article := models.GetByArticle(req.Id)
	if article == nil {
		resp.Msg = errors.New("文章不存在！").Error()
		resp.Code = config.ARTICLE
		return
	}
	if req.Status < 0 || req.Status > 2 || !unicode.IsNumber(rune(req.Status)) {
		resp.Msg = errors.New("非法数据！").Error()
		resp.Code = config.ARTICLE
		return
	}
	if req.Status == article.Status {
		resp.Msg = errors.New("无效修改！").Error()
		resp.Code = config.ARTICLE
		return
	}
	err = models.UpdateArticle(models.Article{
		Identity: req.Id,
		Status:   req.Status,
	})
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.ARTICLE
		return
	}
	resp.Msg = "修改文章私密性成功！"
	resp.Code = http.StatusOK
	return resp, nil
}
