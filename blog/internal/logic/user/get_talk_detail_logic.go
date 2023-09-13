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

type GetTalkDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTalkDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTalkDetailLogic {
	return &GetTalkDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTalkDetailLogic) GetTalkDetail(req *types.GetTalkDetailReq) (resp *types.GetTalkDetailResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.GetTalkDetailResp)
	err = models.GetTalkById(req.Id)
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.TALK
		return
	}
	Talk := new(types.TalkList)
	err = models.GetTalkDetail(req.Id, Talk)
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.TALK
		return
	}
	resp.Msg = "获取成功！"
	resp.Code = http.StatusOK
	resp.Data = map[string]*types.TalkList{
		"talk_list": Talk,
	}
	return resp, nil
}
