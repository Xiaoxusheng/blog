package comment

import (
	"blog/internal/config"
	"blog/internal/models"
	"blog/internal/svc"
	"blog/internal/types"
	"blog/internal/utils"
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApplyCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApplyCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApplyCommentLogic {
	return &ApplyCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApplyCommentLogic) ApplyComment(req *types.ApplyCommentReq) (resp *types.ApplyCommentResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.ApplyCommentResp)
	//判断评论是否存在
	err = models.GetByCommentId(req.Id)
	if err != nil {
		resp.Code = config.COMMENT
		resp.Msg = err.Error()
		return
	}
	err = models.InsertAppleComment(models.Comment{
		Identity:   utils.GetUuid(),
		ParentId:   req.Id,
		IsRead:     false,
		Type:       0,
		ThumbsUp:   0,
		FromId:     req.FromId,
		FromName:   req.FromName,
		FromAvatar: req.FromAvatar,
		ToId:       req.Id,
		ToName:     req.ToName,
		ToAvatar:   req.ToAvatar,
		Content:    req.Content,
		Ip:         req.Id,
		Status:     0,
	})
	if err != nil {
		resp.Code = config.COMMENT
		resp.Msg = err.Error()
		return
	}
	resp.Code = http.StatusOK
	resp.Msg = "回复成功！"
	return resp, nil
}
