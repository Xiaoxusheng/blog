package user

import (
	"blog/internal/config"
	"blog/internal/models"
	"blog/internal/svc"
	"blog/internal/types"
	"blog/internal/utils"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type UserRigisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRigisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRigisterLogic {
	return &UserRigisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRigisterLogic) UserRigister(req *types.UserRigisterReq, ip string) (resp *types.UserRigisterResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.UserRigisterResp)
	ok := models.GetByUsername(req.Username)
	if ok {
		resp.Msg = "用户已存在！"
		resp.Code = config.USER
		return
	}
	err = models.InsertUser(utils.GetUuid(), req.Username, utils.GetMd5(req.Password), "", ip, 0)
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.USER
		return
	}
	resp.Msg = "注册成功！"
	resp.Code = http.StatusOK

	return resp, nil
}
