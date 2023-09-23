package comment

import (
	"net/http"

	"blog/internal/logic/comment"
	"blog/internal/svc"
	"blog/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ApplyCommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ApplyCommentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := comment.NewApplyCommentLogic(r.Context(), svcCtx)
		resp, err := l.ApplyComment(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
