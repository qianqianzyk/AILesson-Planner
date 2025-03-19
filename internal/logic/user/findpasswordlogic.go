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

type FindPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindPasswordLogic {
	return &FindPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindPasswordLogic) FindPassword(req *types.FindPasswordReq) (resp *types.FindPasswordResp, err error) {
	email := req.Email
	newPassword := req.NewPassword
	code := req.Code

	user, err := service.GetUserByEmail(email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, utils.AbortWithException(utils.ErrServer, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, utils.AbortWithException(utils.ErrUserNotExist, err)
	}

	verifyCode, err := service.GetVerificationCode(user.ID)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrVerifyEmail, err)
	}
	if verifyCode != code {
		return nil, utils.AbortWithException(utils.ErrVerifyCode, err)
	}

	passwordPattern := "^[!-~]{8,24}$"
	if !utils.MatchRegexp(passwordPattern, newPassword) {
		return nil, utils.AbortWithException(utils.ErrUsernameOrPassword, err)
	}

	user.Password = newPassword
	err = service.UpdateUser(*user)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUpdateUser, err)
	}

	return &types.FindPasswordResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: "",
	}, nil
}
