package article

import (
	"blog/internal/config"
	"blog/internal/models"
	"blog/internal/svc"
	"blog/internal/types"
	"context"
	"errors"
	"net/http"
	"unicode"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateArticleTopLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateArticleTopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateArticleTopLogic {
	return &UpdateArticleTopLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateArticleTopLogic) UpdateArticleTop(req *types.UpdateArticleTopReq) (resp *types.UpdateArticleTopResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.UpdateArticleTopResp)
	err = models.GetByArticleId(req.Id)
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.ARTICLE
		return
	}
	if !unicode.IsNumber(rune(req.IsTop)) || req.IsTop > 2 || req.IsTop < 0 {
		resp.Msg = errors.New("非法数据！").Error()
		resp.Code = config.ARTICLE
		return
	}
	err = models.UpdateArticle(models.Article{
		Identity: req.Id,
		IsTop:    req.IsTop,
	})
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.ARTICLE
		return
	}
	resp.Msg = "更改文章置顶成功！"
	resp.Code = http.StatusOK
	return resp, nil
}
