package userl

import (
	"context"

	"blog/internal/svc"
	"blog/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *UserRigisterLogic) UserRigister(req *types.UserRigisterReq) (resp *types.UserRigisterResp, err error) {
	// todo: add your logic here and delete this line

	return
}
