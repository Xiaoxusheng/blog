package user

import (
	"net/http"

	"blog/internal/logic/user"
	"blog/internal/svc"
	"blog/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AdminRigisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AdminRigisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewAdminRigisterLogic(r.Context(), svcCtx)
		resp, err := l.AdminRigister(&req, r.RemoteAddr)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
