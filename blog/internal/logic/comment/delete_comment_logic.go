package comment

import (
	"blog/internal/config"
	"blog/internal/models"
	"context"
	"net/http"

	"blog/internal/svc"
	"blog/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentLogic {
	return &DeleteCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCommentLogic) DeleteComment(req *types.DeleteCommentReq) (resp *types.DeleteCommentResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.DeleteCommentResp)
	err = models.GetByCommentId(req.Id)
	if err != nil {
		resp.Code = config.COMMENT
		resp.Msg = err.Error()
		return
	}
	err = models.DeleteComment(req.Id)
	if err != nil {
		resp.Code = config.COMMENT
		resp.Msg = err.Error()
		return
	}
	resp.Code = http.StatusOK
	resp.Msg = "删除成功！"
	return resp, nil
}
