package article

import (
	"blog/internal/config"
	"blog/internal/models"
	"context"
	"net/http"

	"blog/internal/svc"
	"blog/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRecommendArticleByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRecommendArticleByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRecommendArticleByIdLogic {
	return &GetRecommendArticleByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRecommendArticleByIdLogic) GetRecommendArticleById(req *types.GetRecommendArticleByIdReq) (resp *types.GetRecommendArticleByIdResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.GetRecommendArticleByIdResp)
	list := make([]*types.ArticleList, 0)
	err = models.GetByArticleId(req.Id)
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.ARTICLE
		return
	}
	err = models.GetRecommendArticleById(req.Id, list)
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.ARTICLE
		return
	}
	resp.Msg = "获取文章列表成功！"
	resp.Code = http.StatusOK
	resp.Data = map[string][]*types.ArticleList{
		"article_list": list,
	}
	return resp, nil
}
