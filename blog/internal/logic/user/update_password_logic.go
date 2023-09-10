package user

import (
	"blog/internal/config"
	"blog/internal/models"
	"context"
	"net/http"

	"blog/internal/svc"
	"blog/internal/types"
	"blog/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePasswordLogic {
	return &UpdatePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePasswordLogic) UpdatePassword(req *types.UpdatePasswordReq, identity string) (resp *types.UpdatePasswordResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.UpdatePasswordResp)
	//获取密码
	ok := utils.Verify(req.Password)
	if !ok {
		resp.Msg = "用户密码不符合要求！"
		resp.Code = config.USER
		return
	}
	//修改密码
	err = models.UpdatePwd(identity, utils.GetMd5(req.Password))
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.USER
		return
	}
	resp.Msg = "修改成功！"
	resp.Code = http.StatusOK
	return resp, nil
}
