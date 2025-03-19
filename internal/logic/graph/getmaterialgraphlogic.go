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

type GetMaterialGraphLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMaterialGraphLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMaterialGraphLogic {
	return &GetMaterialGraphLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMaterialGraphLogic) GetMaterialGraph(req *types.GetMaterialGraphReq) (resp *types.GetMaterialGraphResp, err error) {
	filename := req.Filename
	authorizationID := req.AuthorizationID

	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}
	if authorizationID == "" {
		authorizationID = strconv.FormatInt(userID, 10)
	}

	graph, err := service.FetchGraphDataByFileName(filename, authorizationID)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrGetGraph, err)
	}

	return &types.GetMaterialGraphResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: graph,
	}, nil
}
