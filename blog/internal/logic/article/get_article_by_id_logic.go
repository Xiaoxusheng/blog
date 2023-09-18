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

type GetArticleByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetArticleByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleByIdLogic {
	return &GetArticleByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticleByIdLogic) GetArticleById(req *types.GetArticleByIdReq) (resp *types.GetArticleByIdResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.GetArticleByIdResp)
	list := make([]*types.ArticleList, 0)
	//判断id是否存在
	err = models.GetId(req.Id)
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.ARTICLE
		return
	}
	err = models.GetByIdArticleList(req.Id, list)
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
