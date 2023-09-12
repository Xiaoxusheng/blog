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

type PublishTalkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishTalkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishTalkLogic {
	return &PublishTalkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishTalkLogic) PublishTalk(req *types.PublishTalkReq, identity string) (resp *types.PublishTalkResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.PublishTalkResp)
	uid := utils.GetUuid()
	err = models.InsertTalk(&models.Talk{
		Identity:  uid,
		UserId:    identity,
		Content:   req.Content,
		Status:    1,
		IsTop:     1,
		LikeTimes: 0,
	})
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.TALK
		return
	}
	err = models.InsertTalkPhoto(&models.TalkPhoto{
		Identity: utils.GetUuid(),
		Url:      req.Imglist,
		TalkId:   uid,
	})
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.TALK
		return
	}
	resp.Msg = "发表成功！"
	resp.Code = http.StatusOK
	resp.Data = map[string]string{
		"talk_id": uid,
		"url":     req.Imglist,
	}
	return resp, nil
}
