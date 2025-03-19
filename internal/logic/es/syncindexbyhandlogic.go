package es

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SyncIndexByHandLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSyncIndexByHandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncIndexByHandLogic {
	return &SyncIndexByHandLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SyncIndexByHandLogic) SyncIndexByHand(req *types.SyncIndexByHandReq) (resp *types.SyncIndexByHandResp, err error) {
	indexType := req.IndexType

	if indexType == 1 {
		err = service.SyncIndex("conversation_messages")
		if err != nil {
			return nil, utils.AbortWithException(utils.ErrSyncIndex, err)
		}
	}

	return &types.SyncIndexByHandResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: "",
	}, nil
}
