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

type BlogTimeArticleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBlogTimeArticleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BlogTimeArticleListLogic {
	return &BlogTimeArticleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BlogTimeArticleListLogic) BlogTimeArticleList(req *types.BlogTimeArticleListReq) (resp *types.BlogTimeArticleListResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.BlogTimeArticleListResp)
	list := make([]*types.ArticleList, 0)
	err = models.GetCreateArticleList(list, req.Limit, req.Offset)
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
