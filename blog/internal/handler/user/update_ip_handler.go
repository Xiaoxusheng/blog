package user

import (
	"net/http"

	"blog/internal/logic/user"
	"blog/internal/svc"
	"blog/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateIpHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateIpReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUpdateIpLogic(r.Context(), svcCtx)
		resp, err := l.UpdateIp(&req, r.Header.Get("identity"))
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
