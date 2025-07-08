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

type GeneratePPTDesignLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGeneratePPTDesignLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GeneratePPTDesignLogic {
	return &GeneratePPTDesignLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GeneratePPTDesignLogic) GeneratePPTDesign(req *types.GeneratePPTDesignReq) (resp *types.GeneratePPTDesignResp, err error) {
	sessionID := req.SessionID
	chatAnswer := "https://disk.qianqianzyk.top/aihelper/share/2025/pdf/a7c50e61-1142-11f0-9c2d-00163e0daf65.pdf,https://disk.qianqianzyk.top/aihelper/share/2025/pptx/中国近代史重点复习.pptx"

	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

	messages := []model.ConversationMessage{
		{
			SessionID:   sessionID,
			UserID:      int(userID),
			Role:        "user",
			Message:     "您的 中国近代史重点复习 PPT已生成完毕，请及时下载保存！",
			MessageType: 1,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			SessionID:   sessionID,
			UserID:      int(userID),
			Role:        "ai",
			Message:     chatAnswer,
			MessageType: 9,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}
	_, err = service.SaveMessage(sessionID, messages)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrSyncToMySQL, err)
	}

	time.Sleep(5 * time.Second)

	return &types.GeneratePPTDesignResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: types.ChatAnswer{Message: chatAnswer},
	}, nil
}
