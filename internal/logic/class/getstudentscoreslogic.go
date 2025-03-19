package class

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStudentScoresLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStudentScoresLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStudentScoresLogic {
	return &GetStudentScoresLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStudentScoresLogic) GetStudentScores(req *types.GetStudentScoresReq) (resp *types.GetStudentScoresResp, err error) {
	studentID := req.StudentID

	studentScores, err := service.GetStudentCourses(studentID)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrGetStudentScore, err)
	}

	var stuScores []types.StudentScores
	for _, s := range studentScores {
		stuScores = append(stuScores, types.StudentScores{
			StudentID:    s.StudentID,
			Name:         s.Name,
			Class:        s.Class,
			Major:        s.Major,
			College:      s.College,
			CourseID:     s.CourseID,
			CourseName:   s.CourseName,
			RegularScore: s.RegularScore,
			FinalScore:   s.FinalScore,
			TotalScore:   s.TotalScore,
			Credit:       s.Credit,
			CreditEarned: s.CreditEarned,
			GradePoint:   s.GradePoint,
			AcademicYear: s.AcademicYear,
			AcademicTerm: s.AcademicTerm,
		})
	}

	return &types.GetStudentScoresResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: stuScores,
	}, nil
}
