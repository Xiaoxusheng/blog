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

type UpdateCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCategoryLogic {
	return &UpdateCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCategoryLogic) UpdateCategory(req *types.UpdatecategoryReq) (resp *types.UpdatecategoryResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.UpdatecategoryResp)
	//判断是否存在
	err = models.GetCategoryById(req.Id)
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.CATEGORY
		return
	}
	err = models.UpdateCategory(req.Id, req.CategoryName)
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.CATEGORY
		return
	}
	resp.Msg = "修改分类成功！"
	resp.Code = http.StatusOK
	return resp, nil
}
