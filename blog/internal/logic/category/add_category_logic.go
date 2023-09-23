package category

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

type AddCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCategoryLogic {
	return &AddCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCategoryLogic) AddCategory(req *types.CategoryReq) (resp *types.CategoryResp, err error) {
	// todo: add your logic here and delete this line
	// todo: add your logic here and delete this line
	resp = new(types.CategoryResp)
	//判断是否存在
	err = models.GetCategory(req.CategoryName)
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.CATEGORY
		return
	}
	err = models.InsertCategory(&models.Category{
		Identity:     utils.GetUuid(),
		CategoryName: req.CategoryName,
		Status:       0,
	})
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.CATEGORY
		return
	}
	resp.Msg = "新增分类成功！"
	resp.Code = http.StatusOK
	return resp, nil
}
