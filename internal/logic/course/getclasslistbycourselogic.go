package course

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetClassListByCourseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetClassListByCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetClassListByCourseLogic {
	return &GetClassListByCourseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetClassListByCourseLogic) GetClassListByCourse(req *types.GetClassListByCourseReq) (resp *types.GetClassListByCourseResp, err error) {
	courseID := req.CourseID

	classList, err := service.GetClassListByCourseID(courseID)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrGetClassList, err)
	}

	return &types.GetClassListByCourseResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: types.ClassList{
			ClassList: classList,
		},
	}, nil
}
