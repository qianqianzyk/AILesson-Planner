package score

import (
	"context"
	"errors"
	"github.com/qianqianzyk/AILesson-Planner/internal/model"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"
	"gorm.io/gorm"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStudentScoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStudentScoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStudentScoreLogic {
	return &GetStudentScoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStudentScoreLogic) GetStudentScore(req *types.GetStudentScoreReq) (resp *types.GetStudentScoreResp, err error) {
	courseID := req.CourseID
	class := req.Class

	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

	course, err := service.GetCourseByID(courseID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, utils.AbortWithException(utils.ErrServer, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, utils.AbortWithException(utils.ErrGetCourse, err)
	}
	if course.UserID != int(userID) {
		return nil, utils.AbortWithException(utils.ErrAuthUser, err)
	}

	var courses []model.StudentScore
	if class == "" {
		courses, err = service.GetStudentScore(courseID)
		if err != nil {
			return nil, utils.AbortWithException(utils.ErrGetStudentScore, err)
		}
	} else {
		courses, err = service.GetStudentScoreByClass(courseID, class)
		if err != nil {
			return nil, utils.AbortWithException(utils.ErrGetStudentScore, err)
		}
	}

	var studentScores []types.StudentScore
	for _, c := range courses {
		studentScores = append(studentScores, types.StudentScore{
			StudentID:    c.StudentID,
			Name:         c.Name,
			Class:        c.Class,
			Major:        c.Major,
			College:      c.College,
			RegularScore: c.RegularScore,
			FinalScore:   c.FinalScore,
			TotalScore:   c.TotalScore,
			CreditEarned: c.CreditEarned,
			GradePoint:   c.GradePoint,
		})
	}

	return &types.GetStudentScoreResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: studentScores,
	}, nil
}
