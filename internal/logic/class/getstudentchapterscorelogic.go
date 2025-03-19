package class

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStudentChapterScoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStudentChapterScoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStudentChapterScoreLogic {
	return &GetStudentChapterScoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStudentChapterScoreLogic) GetStudentChapterScore(req *types.GetStudentChapterScoreReq) (resp *types.GetStudentChapterScoreResp, err error) {
	studentID := req.StudentID
	courseID := req.CourseID

	chapterS, err := service.GetChapterScoresWithAvg(studentID, courseID)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrGetChapterScores, err)
	}

	var chapterScores []types.ChapterScoreWithAvg
	for _, c := range chapterS {
		chapterScores = append(chapterScores, types.ChapterScoreWithAvg{
			Chapter:      c.Chapter,
			StudentScore: c.StudentScore,
			AvgScore:     c.AvgScore,
		})
	}

	return &types.GetStudentChapterScoreResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: chapterScores,
	}, nil
}
