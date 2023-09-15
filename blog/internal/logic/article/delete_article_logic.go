package article

import (
	"blog/internal/config"
	"blog/internal/models"
	"blog/internal/svc"
	"blog/internal/types"
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteArticleLogic {
	return &DeleteArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteArticleLogic) DeleteArticle(req *types.DeleteArticleReq) (resp *types.DeleteArticleResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.DeleteArticleResp)
	//判断id是否存在
	err = models.GetByArticleId(req.Id)
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.ARTICLE
		return
	}
	err = models.UpdateArticle(models.Article{
		Identity: req.Id,
		Status:   3,
	})
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.ARTICLE
		return
	}
	resp.Msg = "删除文章成功！"
	resp.Code = http.StatusOK
	return resp, nil
}
