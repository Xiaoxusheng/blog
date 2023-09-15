package article

import (
	"blog/internal/config"
	"blog/internal/models"
	"blog/internal/svc"
	"blog/internal/types"
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecoveArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecoveArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecoveArticleLogic {
	return &RecoveArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecoveArticleLogic) RecoveArticle(req *types.RecoveArticleReq) (resp *types.RecoveArticleResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.RecoveArticleResp)
	err = models.GetByArticleId(req.Id)
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.ARTICLE
		return
	}
	err = models.UpdateArticle(models.Article{
		Identity: req.Id,
		Status:   1,
	})
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.ARTICLE
		return
	}
	resp.Msg = "恢复文章成功！"
	resp.Code = http.StatusOK
	return resp, nil
}
