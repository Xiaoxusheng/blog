package user

import (
	"blog/internal/config"
	"blog/internal/models"
	"blog/internal/utils"
	"context"
	"net/http"

	"blog/internal/svc"
	"blog/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminRigisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminRigisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminRigisterLogic {
	return &AdminRigisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminRigisterLogic) AdminRigister(req *types.AdminRigisterReq, ip string) (resp *types.AdminRigisterResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.AdminRigisterResp)
	//判断用户是否存在
	_, ok := models.GetByUsername(req.Username)
	if ok {
		resp.Msg = "用户已存在！"
		resp.Code = config.USER
		return
	}
	m := utils.Verify(req.Password)
	if !m {
		resp.Msg = "用户密码不符合要求！"
		resp.Code = config.USER
		return
	}
	err = models.InsertAdminUser(utils.GetUuid(), req.Username, utils.GetMd5(req.Password), req.NickName, ip, req.QQ)
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.USER
		return
	}
	resp.Msg = "注册成功！"
	resp.Code = http.StatusOK

	return resp, nil
}
