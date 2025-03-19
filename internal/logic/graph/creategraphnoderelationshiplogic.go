package graph

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateGraphNodeRelationShipLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateGraphNodeRelationShipLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGraphNodeRelationShipLogic {
	return &CreateGraphNodeRelationShipLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateGraphNodeRelationShipLogic) CreateGraphNodeRelationShip(req *types.CreateGraphNodeRelationShipReq) (resp *types.CreateGraphNodeRelationShipResp, err error) {
	startNodeElementID := req.StartNodeElementID
	endNodeElementID := req.EndNodeElementID
	relationshipType := req.RelationshipType
	filename := req.Filename

	elementID, err := service.CreateGraphNodeRelationShip(startNodeElementID, endNodeElementID, relationshipType, filename)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrCreateGraphNodeRelationship, err)
	}

	return &types.CreateGraphNodeRelationShipResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: types.NodeElementID{
			ElementID: elementID,
		},
	}, nil
}
