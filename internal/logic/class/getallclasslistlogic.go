package class

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllClassListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllClassListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllClassListLogic {
	return &GetAllClassListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllClassListLogic) GetAllClassList(req *types.Empty) (resp *types.GetClassListResp, err error) {
	classList, err := service.GetClassList()
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrGetClassList, err)
	}

	return &types.GetClassListResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: types.ClassList{
			ClassList: classList,
		},
	}, nil
}
