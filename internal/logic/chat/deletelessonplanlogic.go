package chat

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLessonPlanLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLessonPlanLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLessonPlanLogic {
	return &DeleteLessonPlanLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLessonPlanLogic) DeleteLessonPlan(req *types.DeleteLessonPlanReq) (resp *types.DeleteLessonPlanResp, err error) {
	id := req.ID

	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

	err = service.DeleteTPlanByID(id, int(userID))
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrDeleteTPlan, err)
	}

	return &types.DeleteLessonPlanResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: "",
	}, nil
}
