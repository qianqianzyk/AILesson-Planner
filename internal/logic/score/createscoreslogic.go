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

type CreateScoresLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateScoresLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateScoresLogic {
	return &CreateScoresLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateScoresLogic) CreateScores(req *types.CreateStudentScoreReq) (resp *types.CreateStudentScoreResp, err error) {
	courseID := req.CourseID
	createScores := req.CreateScores

	var studentScores []model.Score
	for _, u := range createScores {
		regularScore, _ := strconv.ParseFloat(u.RegularScore, 64)
		finalScore, _ := strconv.ParseFloat(u.FinalScore, 64)
		totalScore, _ := strconv.ParseFloat(u.TotalScore, 64)
		studentScores = append(studentScores, model.Score{
			StudentID:    u.StudentID,
			CourseID:     courseID,
			RegularScore: regularScore,
			FinalScore:   finalScore,
			TotalScore:   totalScore,
		})
	}

	err = service.UpsertStudentScore(studentScores, courseID)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUpsertScores, err)
	}

	return &types.CreateStudentScoreResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: "",
	}, nil
}
