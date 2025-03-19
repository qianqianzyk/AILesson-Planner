package disk

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecoverDiskFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecoverDiskFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecoverDiskFileLogic {
	return &RecoverDiskFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecoverDiskFileLogic) RecoverDiskFile(req *types.RecoverDiskFileReq) (resp *types.RecoverDiskFileResp, err error) {
	ids := req.Ids

	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

	err = service.RecoverFile(ids, int(userID))
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrRecoverFile, err)
	}

	return &types.RecoverDiskFileResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: "",
	}, nil
}
