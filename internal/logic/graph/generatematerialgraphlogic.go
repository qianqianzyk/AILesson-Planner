package graph

import (
	"context"
	"encoding/json"
	"github.com/qianqianzyk/AILesson-Planner/internal/model"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"
	"strconv"
	"time"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateMaterialGraphLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGenerateMaterialGraphLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateMaterialGraphLogic {
	return &GenerateMaterialGraphLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenerateMaterialGraphLogic) GenerateMaterialGraph(req *types.GenerateMaterialGraphReq) (resp *types.GenerateMaterialGraphResp, err error) {
	sessionID := req.SessionID

	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

	authorizationID := strconv.FormatInt(userID, 10)

	graph, err := service.FetchGraphDataByFileName("中国近代史重点.pdf", authorizationID)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrGetGraph, err)
	}

	jsonBytes, err := json.Marshal(graph)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrServer, err)
	}

	messages := []model.ConversationMessage{
		{
			SessionID:   sessionID,
			UserID:      int(userID),
			Role:        "user",
			Message:     "请按照要求为我生成一份知识图谱",
			MessageType: 1,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			SessionID:   sessionID,
			UserID:      int(userID),
			Role:        "ai",
			Message:     string(jsonBytes),
			MessageType: 11,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}
	_, err = service.SaveMessage(sessionID, messages)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrSyncToMySQL, err)
	}

	time.Sleep(30 * time.Second)

	return &types.GenerateMaterialGraphResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: graph,
	}, nil
}
