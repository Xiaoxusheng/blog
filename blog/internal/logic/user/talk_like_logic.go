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

type TalkLikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTalkLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TalkLikeLogic {
	return &TalkLikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TalkLikeLogic) TalkLike(req *types.TalkLikeReq) (resp *types.TalkLikeResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.TalkLikeResp)
	//查询id是否存在
	err = models.GetTalkById(req.Id)
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.TALK
		return
	}
	//点赞
	err = models.UpdateTalkLike(req.Id)
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.TALK
		return
	}
	resp.Msg = "点赞成功！"
	resp.Code = http.StatusOK
	return resp, nil
}
