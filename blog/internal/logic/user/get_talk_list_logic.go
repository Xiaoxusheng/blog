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

type GetTalkListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTalkListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTalkListLogic {
	return &GetTalkListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTalkListLogic) GetTalkList(req *types.GetTalkListReq) (resp *types.GetTalkListResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.GetTalkListResp)
	list := make([]*types.TalkList, 0)
	m, err := models.GetTalkList(list, req.Limit, req.Offset)
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.TALK
		return
	}
	resp.Msg = "获取成功!"
	resp.Code = http.StatusOK
	resp.Data = map[string][]*types.TalkList{
		"talk_list": m,
	}
	return resp, nil
}
