package user

import (
	"context"
	"net/http"
	"strconv"

	"blog/internal/config"
	"blog/internal/models"
	"blog/internal/svc"
	"blog/internal/types"
	"blog/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.UserLoginReq) (resp *types.UserLoginResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.UserLoginResp)
	//判断是否存在
	u, ok := models.GetByUsername(req.Username)
	if !ok {
		resp.Msg = "用户不存在！"
		resp.Code = config.USER
		return
	}

	user := models.GetByUsePwd(req.Username, utils.GetMd5(req.Password), u.Role)
	if user.Username == "" {
		resp.Msg = "用户不存在在或者用户密码错误！"
		resp.Code = config.USER
		return
	}
	resp.Msg = "登陆成功！"
	resp.Code = http.StatusOK
	resp.Data = map[string]string{
		"token":    utils.CreateToken(user.Identity),
		"username": user.Username,
		"role":     strconv.Itoa(user.Role),
		"identity": user.Identity,
	}
	return resp, nil
}
