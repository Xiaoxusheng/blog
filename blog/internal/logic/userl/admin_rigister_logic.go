package userl

import (
	"context"

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

func (l *AdminRigisterLogic) AdminRigister(req *types.AdminRigisterReq) (resp *types.AdminRigisterResp, err error) {
	// todo: add your logic here and delete this line

	return
}
