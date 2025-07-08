package course

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCourseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCourseLogic {
	return &DeleteCourseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCourseLogic) DeleteCourse(req *types.DeleteCourseReq) (resp *types.DeleteCourseResp, err error) {
	courseID := req.CourseID

	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

	err = service.DeleteCourseByID(courseID, int(userID))
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrDeleteCourse, err)
	}

	return &types.DeleteCourseResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: "",
	}, nil
}
