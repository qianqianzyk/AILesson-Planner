package chat

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchMessagesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchMessagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchMessagesLogic {
	return &SearchMessagesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchMessagesLogic) SearchMessages(req *types.SearchMessagesReq) (resp *types.SearchMessagesResp, err error) {
	key := req.Key

	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

	responseConversationSession, err := service.SearchMessages(key, int(userID))
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrSearch, err)
	}

	var chat []types.Chat
	for _, cov := range responseConversationSession {
		chatItem := types.Chat{
			Title:     cov.Title,
			CreatedAt: cov.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: cov.UpdatedAt.Format("2006-01-02 15:04:05"),
			Messages:  []types.Message{},
		}

		for _, msg := range cov.ConversationMessage {
			messageItem := types.Message{
				Role:      msg.Role,
				Message:   msg.Message,
				CreatedAt: msg.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt: msg.UpdatedAt.Format("2006-01-02 15:04:05"),
			}
			chatItem.Messages = append(chatItem.Messages, messageItem)
		}

		chat = append(chat, chatItem)
	}

	return &types.SearchMessagesResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: chat,
	}, nil
}
