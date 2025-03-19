package chat

import (
	"context"
	"errors"
	"github.com/qianqianzyk/AILesson-Planner/internal/model"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"
	"gorm.io/gorm"
	"time"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetChatAnswerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetChatAnswerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChatAnswerLogic {
	return &GetChatAnswerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetChatAnswerLogic) GetChatAnswer(req *types.GetChatAnswerReq) (resp *types.GetChatAnswerResp, err error) {
	sessionID := req.SessionID
	message := req.Message

	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

	session, err := service.GetTopicByID(uint(sessionID))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, utils.AbortWithException(utils.ErrServer, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, utils.AbortWithException(utils.ErrGetTopic, err)
	}
	if session.UserID != int(userID) {
		return nil, utils.AbortWithException(utils.ErrAuthUser, err)
	}

	answer, err := service.GetAnswerTextByTongyi(
		l.svcCtx.Config.Tongyi.Endpoint,
		l.svcCtx.Config.Tongyi.APIKey,
		message,
		sessionID,
	)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrGenAnswer, err)
	}

	messages := []model.ConversationMessage{
		{
			SessionID: sessionID,
			UserID:    session.UserID,
			Role:      "user",
			Message:   message,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			SessionID: sessionID,
			UserID:    session.UserID,
			Role:      "ai",
			Message:   answer,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	err = service.SaveMessageToMySQL(messages)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrSyncToMySQL, err)
	}

	err = service.SyncMessageIndexToES(messages)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrSyncIndex, err)
	}

	return &types.GetChatAnswerResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: types.ChatAnswer{
			Message: answer,
		},
	}, nil
}
