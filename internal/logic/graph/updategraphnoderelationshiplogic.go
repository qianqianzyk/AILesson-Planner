package graph

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGraphNodeRelationShipLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateGraphNodeRelationShipLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGraphNodeRelationShipLogic {
	return &UpdateGraphNodeRelationShipLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateGraphNodeRelationShipLogic) UpdateGraphNodeRelationShip(req *types.UpdateGraphNodeRelationShipReq) (resp *types.UpdateGraphNodeRelationShipResp, err error) {
	elementID := req.ElementID
	relationshipType := req.RelationshipType
	filename := req.Filename

	err = service.UpdateGraphNodeRelationShip(elementID, relationshipType, filename)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUpdateGraphNodeRelationship, err)
	}

	return &types.UpdateGraphNodeRelationShipResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: "",
	}, nil
}
