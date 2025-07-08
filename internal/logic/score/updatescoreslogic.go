package score

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/model"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"
	"strconv"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateScoresLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateScoresLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateScoresLogic {
	return &UpdateScoresLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateScoresLogic) UpdateScores(req *types.UpdateStudentScoreReq) (resp *types.UpdateStudentScoreResp, err error) {
	courseID := req.CourseID
	updateScores := req.UpdateScores

	var studentScores []model.StudentScore
	for _, u := range updateScores {
		regularScore, _ := strconv.ParseFloat(u.RegularScore, 64)
		finalScore, _ := strconv.ParseFloat(u.FinalScore, 64)
		totalScore, _ := strconv.ParseFloat(u.TotalScore, 64)
		studentScores = append(studentScores, model.StudentScore{
			StudentID:    u.StudentID,
			Name:         u.Name,
			Class:        u.Class,
			Major:        u.Major,
			College:      u.College,
			RegularScore: regularScore,
			FinalScore:   finalScore,
			TotalScore:   totalScore,
		})
	}

	err = service.UpdateStudentScores(studentScores, courseID)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUpdateScores, err)
	}

	return &types.UpdateStudentScoreResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: "",
	}, nil
}
