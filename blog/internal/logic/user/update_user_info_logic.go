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

type UpdateUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(req *types.UpdateUserInfoReq, identity string) (resp *types.UpdateUserInfoResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.UpdateUserInfoResp)
	err = models.UpdateUserById(identity, req)
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.USER
		return
	}
	resp.Msg = "更新成功！"
	resp.Code = http.StatusOK
	return
}
