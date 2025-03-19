package chat

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteChatTopicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteChatTopicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteChatTopicLogic {
	return &DeleteChatTopicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteChatTopicLogic) DeleteChatTopic(req *types.DeleteChatTopicReq) (resp *types.DeleteChatTopicResp, err error) {
	id := req.ID

	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

	err = service.DelTopicByID(uint(id), userID)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrDelTopic, err)
	}

	return &types.DeleteChatTopicResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: "",
	}, nil
}
