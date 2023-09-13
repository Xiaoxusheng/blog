package tag

import (
	"blog/internal/config"
	"blog/internal/models"
	"context"
	"net/http"

	"blog/internal/svc"
	"blog/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTagLogic {
	return &UpdateTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateTagLogic) UpdateTag(req *types.UpdateTagReq) (resp *types.UpdateTagResp, err error) {
	// todo: add your logic here and delete this line
	//判断标签是否存在
	resp = new(types.UpdateTagResp)
	err = models.GetTagById(req.Id)
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.TALK
		return
	}
	err = models.UpdateTag(req.Id, req.TagName)
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.TALK
		return
	}
	resp.Msg = "修改tag标签成功!"
	resp.Code = http.StatusOK
	return resp, nil
}
