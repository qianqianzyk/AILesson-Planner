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

type GetMaterialGraphListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMaterialGraphListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMaterialGraphListLogic {
	return &GetMaterialGraphListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMaterialGraphListLogic) GetMaterialGraphList(req *types.GetMaterialGraphListReq) (resp *types.GetMaterialGraphListResp, err error) {
	authorizationID := req.AuthorizationID
	graphType := req.GraphType

	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}
	if authorizationID == "" {
		authorizationID = strconv.FormatInt(userID, 10)
	}

	graphList, err := service.GetDocumentList(authorizationID, graphType)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrGetGraph, err)
	}

	return &types.GetMaterialGraphListResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: graphList,
	}, nil
}
