package chat

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateChatTopicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateChatTopicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateChatTopicLogic {
	return &UpdateChatTopicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateChatTopicLogic) UpdateChatTopic(req *types.UpdateChatTopicReq) (resp *types.UpdateChatTopicResp, err error) {
	id := req.ID
	newTopic := req.NewTopic

	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

	cov, err := service.GetTopicByID(uint(id))
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrGetTopic, err)
	}
	if cov.UserID != int(userID) {
		return nil, utils.AbortWithException(utils.ErrAuthUser, err)
	}

	cov.Title = newTopic
	err = service.UpdateTopic(*cov)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUpdateTopic, err)
	}

	return &types.UpdateChatTopicResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: "",
	}, nil
}
