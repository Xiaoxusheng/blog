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

type GetArticleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetArticleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleListLogic {
	return &GetArticleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticleListLogic) GetArticleList(req *types.GetArticleListReq) (resp *types.GetArticleListResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.GetArticleListResp)
	list := make([]*types.ArticleList, 0)
	err = models.GetArticleList(list, req.Limit, req.Offset)
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
