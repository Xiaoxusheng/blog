package category

import (
	"blog/internal/config"
	"blog/internal/models"
	"context"
	"net/http"

	"blog/internal/svc"
	"blog/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCategoryLogic {
	return &DeleteCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCategoryLogic) DeleteCategory(req *types.DeleteCategoryReq) (resp *types.DeleteCategoryResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.DeleteCategoryResp)
	//判断是否存在
	err = models.GetCategoryById(req.Id)
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.CATEGORY
		return
	}
	err = models.DeleteCategory(req.Id)
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.CATEGORY
		return
	}
	resp.Msg = "删除分类成功！"
	resp.Code = http.StatusOK
	return resp, nil
}
