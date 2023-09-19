package comment

import (
	`blog/internal/config`
	`blog/internal/models`
	"blog/internal/svc"
	"blog/internal/types"
	`blog/internal/utils`
	"context"
	`encoding/json`
	`net/http`

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCommentLogic {
	return &AddCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCommentLogic) AddComment(req *types.AddCommentReq) (resp *types.AddCommentResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.AddCommentResp)
	//判断id是否存在
	err = models.GetByCommentId(req.Id)
	if err != nil {
		resp.Code = config.COMMENT
		resp.Msg = err.Error()
	}
	//是否有违禁词

	//进入队列
	c := &models.Comment{
		Identity:   utils.GetUuid(),
		Type:       req.Type,
		IsRead:     false,
		ThumbsUp:   0,
		FromId:     req.FromId,
		FromName:   req.FromName,
		FromAvatar: req.FromAvatar,
		Content:    req.Content,
		Ip:         req.Id,
	}

	m, err := json.Marshal(c)
	if err != nil {
		resp.Code = config.COMMENT
		resp.Msg = err.Error()
		return
	}

	err = utils.Pub(m)
	if err != nil {
		resp.Code = config.COMMENT
		resp.Msg = err.Error()
		return
	}
	//进入数据库
	err = models.InsertComment(c)
	if err != nil {
		resp.Code = config.COMMENT
		resp.Msg = err.Error()
		return
	}
	//响应
	resp.Code = http.StatusOK
	resp.Msg = "评论成功！"
	return resp, nil
}
