package user

import (
	"context"
	"log"
	"net/http"

	"blog/internal/config"
	"blog/internal/models"
	"blog/internal/svc"
	"blog/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRoleLogic) UpdateRole(req *types.UpdateRoleReq, identity string) (resp *types.UpdateRoleResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.UpdateRoleResp)
	log.Println("identity", identity)
	resp = new(types.UpdateRoleResp)
	//修改
	user, ok := models.GetById(req.Id)
	if !ok {
		resp.Msg = "id不存在"
		resp.Code = config.USER
		return
	}
	if user.Role != 9999 {
		resp.Msg = "您没有权限修改"
		resp.Code = config.USER
		return
	}
	err = models.UpdateRole(identity, req.Role)
	if err != nil {
		return nil, err
	}
	resp.Msg = "更新成功"
	resp.Code = http.StatusOK
	return resp, nil
}
