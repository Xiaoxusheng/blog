package user

import (
	"blog/internal/config"
	"blog/internal/models"
	"blog/internal/svc"
	"blog/internal/types"
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateIpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateIpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateIpLogic {
	return &UpdateIpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateIpLogic) UpdateIp(req *types.UpdateIpReq, identity string) (resp *types.UpdateIpResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.UpdateIpResp)

	if req.Ip == "" {
		resp.Msg = "ip不能为空！"
		resp.Code = config.USER
		return
	}
	//判断权限

	//修改ip
	err = models.UpdateIp(identity, req.Ip)
	if err != nil {
		resp.Msg = "ip不能为空！"
		resp.Code = config.USER
		return
	}
	resp.Msg = "修改成功！！"
	resp.Code = http.StatusOK
	return
}
