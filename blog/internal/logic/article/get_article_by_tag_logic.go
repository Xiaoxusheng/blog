package article

import (
	"blog/internal/config"
	"blog/internal/models"
	"context"
	"errors"
	"net/http"

	"blog/internal/svc"
	"blog/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleByTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetArticleByTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleByTagLogic {
	return &GetArticleByTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticleByTagLogic) GetArticleByTag(req *types.GetArticleByTagReq) (resp *types.GetArticleByTagResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.GetArticleByTagResp)
	list := make([]*types.ArticleList, 0)

	//判断tag是否存在
	err = models.GetTagName(req.Tag)
	if err == nil {
		resp.Msg = errors.New("tag不存在！").Error()
		resp.Code = config.ARTICLE
		return
	}
	//获取文章列表
	err = models.GetByTagArticleList(list, req.Tag)
	if err != nil {
		resp.Msg = errors.New("tag不存在！").Error()
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
