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

type UnlikeArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUnlikeArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnlikeArticleLogic {
	return &UnlikeArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UnlikeArticleLogic) UnlikeArticle(req *types.UnlikeArticleReq) (resp *types.UnlikeArticleResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.UnlikeArticleResp)
	err = models.GetByArticleId(req.Id)
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.ARTICLE
		return
	}
	err = models.UpdateUnLike(req.Id)
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.ARTICLE
		return
	}
	resp.Msg = "点赞成功！"
	resp.Code = http.StatusOK
	return resp, nil
}
