package tag

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

type AddTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTagLogic {
	return &AddTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddTagLogic) AddTag(req *types.AddTagReq) (resp *types.AddTagResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.AddTagResp)
	err = models.GetTagName(req.TagName)
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.TALK
		return
	}
	err = models.InsertTag(utils.GetUuid(), req.TagName)
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.TALK
		return
	}
	resp.Msg = "新增tag标签成功!"
	resp.Code = http.StatusOK
	return resp, nil
}
