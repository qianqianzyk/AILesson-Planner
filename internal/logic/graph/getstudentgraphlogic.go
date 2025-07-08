package graph

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStudentGraphLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStudentGraphLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStudentGraphLogic {
	return &GetStudentGraphLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStudentGraphLogic) GetStudentGraph(req *types.Empty) (resp *types.GetStudentGraphResp, err error) {
	graph, err := service.FetchGraphDataByFileName("浅浅的中国近代史知识点掌握情况", "20250001")
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrGetGraph, err)
	}

	return &types.GetStudentGraphResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: graph,
	}, nil
}
