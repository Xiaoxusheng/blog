// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"
	"time"

	user "blog/internal/handler/user"
	"blog/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: user.UserLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/rigister",
				Handler: user.UserRigisterHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
		rest.WithTimeout(3000*time.Millisecond),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.ParseToken},
			[]rest.Route{
				{
					Method:  http.MethodPut,
					Path:    "/user/updateRole",
					Handler: user.UpdateRoleHandler(serverCtx),
				},
			}...,
		),
		rest.WithSignature(serverCtx.Config.Signature),
		rest.WithPrefix("/v1"),
		rest.WithTimeout(3000*time.Millisecond),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/user/updatePassword",
				Handler: UpdatePasswordHandler(serverCtx),
			},
		},
	)
}
