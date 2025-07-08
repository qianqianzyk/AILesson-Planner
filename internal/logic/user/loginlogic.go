package user

import (
	"context"
	"errors"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"
	"gorm.io/gorm"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	email := req.Email
	password := req.Password

	user, err := service.GetUserByEmail(email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, utils.AbortWithException(utils.ErrServer, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, utils.AbortWithException(utils.ErrUserNotExist, err)
	}

	if user.Password != password {
		return nil, utils.AbortWithException(utils.ErrLogin, err)
	}

	tokenStr, err := utils.GenerateToken(user.ID, l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.Config.Auth.AccessExpire)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrGenToken, err)
	}

	return &types.LoginResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: types.LoginToken{
			UserID: user.ID,
			Token:  tokenStr,
		},
	}, nil
}
