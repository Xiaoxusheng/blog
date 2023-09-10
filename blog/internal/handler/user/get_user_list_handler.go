package user

import (
	"net/http"

	"blog/internal/logic/user"
	"blog/internal/svc"
	"blog/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetUserListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewGetUserListLogic(r.Context(), svcCtx)
		resp, err := l.GetUserList(&req, r.Header.Get("identity"))
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
