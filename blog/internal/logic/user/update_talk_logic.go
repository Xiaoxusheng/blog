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

type UpdateTalkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateTalkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTalkLogic {
	return &UpdateTalkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateTalkLogic) UpdateTalk(req *types.UpdateTalkReq, identity string) (resp *types.UpdateTalkResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.UpdateTalkResp)
	//判断说说id是否存在
	err = models.GeTalkById(req.Id)
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.TALK
		return
	}
	err = models.UpdateTalk(identity, req.Content, req.Status, req.Istop)
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.TALK
		return
	}
	err = models.UpdateTalkPhoto(req.Id, req.Imglist)
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.TALK
		return
	}
	resp.Msg = "修改成功！"
	resp.Code = http.StatusOK
	return resp, nil
}
