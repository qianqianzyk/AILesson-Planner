package class

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetClassListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetClassListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetClassListLogic {
	return &GetClassListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetClassListLogic) GetClassList(req *types.Empty) (resp *types.GetClassListResp, err error) {
	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

	classList, err := service.GetClassListByUserID(int(userID))
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
