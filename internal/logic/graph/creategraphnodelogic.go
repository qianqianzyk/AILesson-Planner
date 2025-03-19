package graph

import (
	"context"
	"github.com/google/uuid"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"
	"strconv"
	"time"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateGraphNodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateGraphNodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGraphNodeLogic {
	return &CreateGraphNodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateGraphNodeLogic) CreateGraphNode(req *types.CreateGraphNodeReq) (resp *types.CreateGraphNodeResp, err error) {
	nodeType := req.NodeType
	createInformation := req.CreateInformation
	authorizationID := req.AuthorizationID

	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}
	if authorizationID == "" {
		authorizationID = strconv.FormatInt(userID, 10)
	}

	if nodeType == "Document" {
		createInformation["authorization_id"] = authorizationID
		createInformation["fileSource"] = "manual"
		createInformation["fileType"] = "text"
		createInformation["fileSize"] = 0
		createInformation["is_cancelled"] = false
		createInformation["model"] = "手动创建"
		createInformation["status"] = "Completed"
		createInformation["total_chunks"] = 0
		createInformation["processingTime"] = 0.0
		createInformation["processed_chunk"] = 0
		createInformation["relationshipCount"] = 0
		createInformation["nodeCount"] = 0
		createInformation["communityNodeCount"] = 0
		createInformation["communityRelCount"] = 0
		createInformation["entityNodeCount"] = 0
		createInformation["entityEntityRelCount"] = 0
		createInformation["errorMessage"] = ""
		createInformation["createdAt"] = time.Now().Format("2006-01-02 15:04:05")
		createInformation["updatedAt"] = time.Now().Format("2006-01-02 15:04:05")
	}
	if nodeType == "Chunk" {
		text, exists := createInformation["text"]
		if exists {
			createInformation["length"] = len(text.(string))
		} else {
			createInformation["length"] = 0
		}
		createInformation["id"] = uuid.New().String()
		createInformation["position"] = 1
	}

	elementID, err := service.CreateNode(nodeType, createInformation)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrCreateGraphNode, err)
	}

	return &types.CreateGraphNodeResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: types.NodeElementID{
			ElementID: elementID,
		},
	}, nil
}
