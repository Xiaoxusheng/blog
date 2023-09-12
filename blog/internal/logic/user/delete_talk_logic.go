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

type DeleteTalkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteTalkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTalkLogic {
	return &DeleteTalkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteTalkLogic) DeleteTalk(req *types.DeleteTalkReq, identity string) (resp *types.DeleteTalkResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.DeleteTalkResp)
	err = models.GeTalkById(req.Id)
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.TALK
		return
	}
	err = models.DeleteTalk(identity)
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.TALK
		return
	}
	resp.Msg = "修改成功！"
	resp.Code = http.StatusOK
	return resp, nil
}
