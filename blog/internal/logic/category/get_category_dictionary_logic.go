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

type GetCategoryDictionaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCategoryDictionaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryDictionaryLogic {
	return &GetCategoryDictionaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCategoryDictionaryLogic) GetCategoryDictionary(req *types.GetCategoryDictionaryReq) (resp *types.GetCategoryDictionaryResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.GetCategoryDictionaryResp)
	list := make(map[string]string)
	err = models.GetCategoryCount(list)

	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.CATEGORY
		return
	}
	resp.Msg = "获取成功！"
	resp.Code = http.StatusOK
	return resp, nil
}
