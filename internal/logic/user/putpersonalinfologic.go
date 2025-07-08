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

type PutPersonalInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutPersonalInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutPersonalInfoLogic {
	return &PutPersonalInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutPersonalInfoLogic) PutPersonalInfo(req *types.PutPersonalInfoReq) (resp *types.PutPersonalInfoResp, err error) {
	username := req.Username
	avatar := req.Avatar

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

	user.Username = username
	if user.Avatar != avatar {
		service.DeleteObjectByUrlAsync(user.Avatar)
		user.Avatar = avatar
	}
	err = service.UpdateUser(*user)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUpdateUser, err)
	}

	return &types.PutPersonalInfoResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: "",
	}, nil
}
