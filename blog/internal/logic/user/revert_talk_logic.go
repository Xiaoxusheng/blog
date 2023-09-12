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

type RevertTalkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRevertTalkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RevertTalkLogic {
	return &RevertTalkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RevertTalkLogic) RevertTalk(req *types.RevertTalkReq, identity string) (resp *types.RevertTalkResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.RevertTalkResp)
	//判断id是否存在
	err = models.GeTalkById(req.Id)
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.TALK
		return
	}
	err = models.RecoverTalk(identity)
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.TALK
		return
	}
	resp.Msg = "恢复成功！"
	resp.Code = http.StatusOK
	return resp, nil
}
