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

type GetCategoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryListLogic {
	return &GetCategoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCategoryListLogic) GetCategoryList(req *types.GetCategoryListReq) (resp *types.GetCategoryListResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.GetCategoryListResp)
	list := make([]*types.CategoryList, 0)
	err = models.GetCategoryList(list)
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.CATEGORY
		return
	}
	resp.Msg = "删除分类成功！"
	resp.Code = http.StatusOK
	resp.Data = map[string][]*types.CategoryList{
		"category_list": list,
	}
	return resp, nil
}
