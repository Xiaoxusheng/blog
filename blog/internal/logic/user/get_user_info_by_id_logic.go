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

type GetUserInfoByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoByIdLogic {
	return &GetUserInfoByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoByIdLogic) GetUserInfoById(req *types.GetUserInfoByIdReq, identity string) (resp *types.GetUserInfoByIdResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.GetUserInfoByIdResp)
	user, ok := models.GetById(identity)
	if !ok {
		resp.Msg = "用户不存在！"
		resp.Code = config.USER
		return
	}
	resp.Msg = "获取用户信息成功！"
	resp.Code = http.StatusOK
	resp.Data = map[string]types.Userlist{
		"data": {
			Identity:  user.Identity,
			Username:  user.Username,
			Role:      user.Role,
			Nick_name: user.Nick_name,
			Avatar:    user.Avatar,
			Qq:        user.Qq,
			Ip:        user.Ip,
		},
	}
	return resp, nil
}
