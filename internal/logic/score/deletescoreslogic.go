package score

import (
	"context"
	"fmt"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteScoresLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteScoresLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteScoresLogic {
	return &DeleteScoresLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteScoresLogic) DeleteScores(req *types.DeleteStudentScoreReq) (resp *types.DeleteStudentScoreResp, err error) {
	studentIDs := req.StudentIDs
	courseID := req.CourseID

	fmt.Println(studentIDs)

	err = service.DeleteStudentScores(studentIDs, courseID)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrDeleteScores, err)
	}

	return &types.DeleteStudentScoreResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: "",
	}, nil
}
