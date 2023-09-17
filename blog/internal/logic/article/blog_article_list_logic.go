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

type BlogArticleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBlogArticleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BlogArticleListLogic {
	return &BlogArticleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BlogArticleListLogic) BlogArticleList(req *types.BlogArticleListReq) (resp *types.BlogArticleListResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.BlogArticleListResp)
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
