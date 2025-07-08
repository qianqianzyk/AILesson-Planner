package email

import (
	"context"
	"errors"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendCodeLogic {
	return &SendCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendCodeLogic) SendCode(req *types.SendCodeReq) (resp *types.SendCodeResp, err error) {
	email := req.Email

	user, err := service.GetUserByEmail(email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, utils.AbortWithException(utils.ErrServer, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, utils.AbortWithException(utils.ErrUserNotExist, err)
	}

	err = service.SendMsgToEmail(l.svcCtx.Config.Email.Name, email, l.svcCtx.Config.Email.Key, user.ID)
	if err != nil {
		if errors.Is(err, utils.ErrTimeLimited) {
			return nil, utils.AbortWithException(utils.ErrSendCodeLimited, err)
		}
		return nil, utils.AbortWithException(utils.ErrServer, err)
	}

	return &types.SendCodeResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: "",
	}, nil
}
