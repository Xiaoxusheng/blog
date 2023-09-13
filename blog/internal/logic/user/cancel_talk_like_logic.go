package user

import (
	"blog/internal/config"
	"blog/internal/models"
	"context"
	"net/http"

	"blog/internal/svc"
	"blog/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelTalkLikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCancelTalkLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelTalkLikeLogic {
	return &CancelTalkLikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancelTalkLikeLogic) CancelTalkLike(req *types.CancelTalkLikeReq) (resp *types.CancelTalkLikeResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.CancelTalkLikeResp)
	err = models.GetTalkById(req.Id)
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.TALK
		return
	}
	err = models.CancelTalkLike(req.Id)
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.TALK
		return
	}
	resp.Msg = "取消点赞成功！"
	resp.Code = http.StatusOK
	return
}
