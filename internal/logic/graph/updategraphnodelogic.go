package graph

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"
	"strconv"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGraphNodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateGraphNodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGraphNodeLogic {
	return &UpdateGraphNodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateGraphNodeLogic) UpdateGraphNode(req *types.UpdateGraphNodeReq) (resp *types.UpdateGraphNodeResp, err error) {
	updateInformation := req.UpdateInformation
	nodeType := req.NodeType
	elementID := req.ElementID
	authorizationID := req.AuthorizationID

	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}
	if authorizationID == "" {
		authorizationID = strconv.FormatInt(userID, 10)
	}

	if nodeType == "Chunk" {
		text, exists := updateInformation["text"]
		if exists {
			textStr, ok := text.(string)
			if ok {
				length := len(textStr)
				updateInformation["length"] = length
			}
		}
	}

	err = service.UpdateNodeByElementID(nodeType, elementID, authorizationID, updateInformation)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUpdateGraphNode, err)
	}

	return &types.UpdateGraphNodeResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: "",
	}, nil
}
