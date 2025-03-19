package graph

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteGraphNodeRelationShipLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteGraphNodeRelationShipLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteGraphNodeRelationShipLogic {
	return &DeleteGraphNodeRelationShipLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteGraphNodeRelationShipLogic) DeleteGraphNodeRelationShip(req *types.DeleteGraphNodeRelationShipReq) (resp *types.DeleteGraphNodeRelationShipResp, err error) {
	elementID := req.ElementID
	filename := req.Filename

	err = service.DeleteGraphNodeRelationShip(elementID, filename)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrDeleteGraphNode, err)
	}

	return &types.DeleteGraphNodeRelationShipResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: "",
	}, nil
}
