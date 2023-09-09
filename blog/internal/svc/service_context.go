package svc

import (
	"blog/internal/config"
	"blog/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config     config.Config
	ParseToken rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		ParseToken: middleware.NewParseTokenMiddleware().Handle,
	}
}
