package chat

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/model"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"
	"time"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateImageDesignLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGenerateImageDesignLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateImageDesignLogic {
	return &GenerateImageDesignLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenerateImageDesignLogic) GenerateImageDesign(req *types.GenerateImageDesignReq) (resp *types.GenerateImageDesignResp, err error) {
	sessionID := req.SessionID

	chatAnswer := "https://disk.qianqianzyk.top/aihelper/share/2025/png/271e4e37-30c5-11f0-9a72-04bf1b6f946c.png"

	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

	messages := []model.ConversationMessage{
		{
			SessionID:   sessionID,
			UserID:      int(userID),
			Role:        "user",
			Message:     "假设你是一名中国近代史的教师，现在你要授课，需要一张鸦片战争海战的图片，请生成！",
			MessageType: 1,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			SessionID:   sessionID,
			UserID:      int(userID),
			Role:        "ai",
			Message:     chatAnswer,
			MessageType: 10,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}
	_, err = service.SaveMessage(sessionID, messages)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrSyncToMySQL, err)
	}

	time.Sleep(5 * time.Second)

	return &types.GenerateImageDesignResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: types.ChatAnswer{Message: chatAnswer},
	}, nil
}
