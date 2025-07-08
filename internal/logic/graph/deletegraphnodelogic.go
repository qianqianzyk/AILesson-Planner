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

type DeleteGraphNodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteGraphNodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteGraphNodeLogic {
	return &DeleteGraphNodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteGraphNodeLogic) DeleteGraphNode(req *types.DeleteGraphNodeReq) (resp *types.DeleteGraphNodeResp, err error) {
	elementID := req.ElementID
	nodeType := req.NodeType
	filename := req.Filename
	authorizationID := req.AuthorizationID

	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}
	if authorizationID == "" {
		authorizationID = strconv.FormatInt(userID, 10)
	}

	err = service.DeleteNodeByElementID(elementID, nodeType, filename, authorizationID)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrDeleteGraphNode, err)
	}

	return &types.DeleteGraphNodeResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: "",
	}, nil
}
