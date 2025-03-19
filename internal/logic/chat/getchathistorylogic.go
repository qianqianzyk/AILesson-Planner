package chat

import (
	"context"
	"errors"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"
	"gorm.io/gorm"
	"math"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetChatHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetChatHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChatHistoryLogic {
	return &GetChatHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetChatHistoryLogic) GetChatHistory(req *types.GetChatHistoryReq) (resp *types.GetChatHistoryResp, err error) {
	id := req.ID
	pageNum := req.PageNum
	pageSize := req.PageSize

	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

	session, err := service.GetTopicByID(id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, utils.AbortWithException(utils.ErrServer, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, utils.AbortWithException(utils.ErrGetTopic, err)
	}
	if session.UserID != int(userID) {
		return nil, utils.AbortWithException(utils.ErrAuthUser, err)
	}

	messages, totalSize, err := service.GetMessageList(int(id), pageNum, pageSize)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrGetMessageList, err)
	}

	var messagesList []types.History
	for _, msg := range messages {
		messagesList = append(messagesList, types.History{
			ID:        msg.ID,
			Role:      msg.Role,
			Message:   msg.Message,
			CreatedAt: msg.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: msg.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &types.GetChatHistoryResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: types.ChatHistory{
			SessionID:    int(id),
			TotalPageNum: math.Ceil(float64(*totalSize) / float64(pageSize)),
			ChatHistory:  messagesList,
		},
	}, nil
}
