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

type GetTagListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTagListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTagListLogic {
	return &GetTagListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTagListLogic) GetTagList(req *types.GetTagListReq) (resp *types.GetTagListResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.GetTagListResp)
	list := make([]*types.TagList, 0)
	err = models.GetTagList(list, req.Offset, req.Limit)
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.TALK
		return
	}
	resp.Msg = "获取tag标签列表成功!"
	resp.Code = http.StatusOK
	resp.Data = map[string][]*types.TagList{
		"tag_list": list,
	}
	return resp, nil
}
