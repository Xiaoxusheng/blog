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

type GetUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserListLogic) GetUserList(req *types.GetUserListReq, identity string) (resp *types.GetUserListResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.GetUserListResp)
	Userlist := make([]*types.Userlist, 0)
	//查看是否为管理
	ok, err := models.GetByIdentity(identity)
	if !ok {
		resp.Msg = err.Error()
		resp.Code = config.USER
		return
	}
	list, err := models.GetUserList(identity, req.Offset, req.Limit, Userlist)
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.USER
		return nil, err
	}
	resp.Msg = "获取成功！"
	resp.Code = http.StatusOK
	resp.Data = map[string][]*types.Userlist{
		"userlist": list,
	}

	return
}
