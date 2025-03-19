package user

import (
	"context"
	"errors"
	"github.com/qianqianzyk/AILesson-Planner/internal/model"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"
	"gorm.io/gorm"
	"time"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	username := req.Username
	password := req.Password
	email := req.Email

	// 检查用户名和密码是否合法
	// 用户名1-20位
	usernamePattern := "^[^\\s]{1,20}$"
	// 密码8-24位
	passwordPattern := "^[!-~]{8,24}$"
	// 邮箱限定格式
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	if !utils.MatchRegexp(usernamePattern, username) || !utils.MatchRegexp(passwordPattern, password) ||
		!utils.MatchRegexp(emailPattern, email) {
		return nil, utils.AbortWithException(utils.ErrUsernameOrPassword, err)
	}

	// 检查用户是否已经存在
	_, err = service.GetUserByUsername(username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, utils.AbortWithException(utils.ErrServer, err)
	}
	if err == nil {
		return nil, utils.AbortWithException(utils.ErrUserExist, err)
	}

	// 检查邮箱是否已经存在
	_, err = service.GetUserByEmail(email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, utils.AbortWithException(utils.ErrServer, err)
	}
	if err == nil {
		return nil, utils.AbortWithException(utils.ErrUserExist, err)
	}

	// 创建用户
	err = service.CreateUser(model.User{
		Username:  username,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
	})
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrCreateUser, err)
	}

	// 响应返回
	return &types.RegisterResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: "",
	}, nil
}
