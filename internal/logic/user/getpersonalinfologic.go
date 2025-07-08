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

type GetPersonalInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPersonalInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPersonalInfoLogic {
	return &GetPersonalInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPersonalInfoLogic) GetPersonalInfo(req *types.Empty) (resp *types.GetPersonalInfoResp, err error) {
	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

	user, err := service.GetUserByUserID(uint(userID))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, utils.AbortWithException(utils.ErrServer, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, utils.AbortWithException(utils.ErrUserNotExist, err)
	}

	return &types.GetPersonalInfoResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: types.PersonalInfo{
			UserID:    user.ID,
			Username:  user.Username,
			Email:     user.Email,
			Avatar:    user.Avatar,
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		},
	}, nil
}
