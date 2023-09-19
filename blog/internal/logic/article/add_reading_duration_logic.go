package article

import (
	`blog/internal/config`
	`blog/internal/models`
	"context"
	`net/http`

	"blog/internal/svc"
	"blog/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddReadingDurationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddReadingDurationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddReadingDurationLogic {
	return &AddReadingDurationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddReadingDurationLogic) AddReadingDuration(req *types.AddReadingDurationReq) (resp *types.AddReadingDurationResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.AddReadingDurationResp)
	err = models.GetByArticleId(req.Id)
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.ARTICLE
		return
	}
	err = models.UpdateTime(req.Id, req.Duration)
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = config.ARTICLE
		return
	}
	resp.Msg = "增加文章时长成功！"
	resp.Code = http.StatusOK
	return resp, nil
}
