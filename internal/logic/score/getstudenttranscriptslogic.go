package score

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStudentTranscriptsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStudentTranscriptsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStudentTranscriptsLogic {
	return &GetStudentTranscriptsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStudentTranscriptsLogic) GetStudentTranscripts(req *types.GetStudentTranscriptsReq) (resp *types.GetStudentTranscriptsResp, err error) {
	academicYear := req.AcademicYear
	academicTerm := req.AcademicTerm
	name := req.Name
	studentID := req.StudentID

	transcripts, err := service.GetStudentTranscripts(academicYear, name, studentID, academicTerm)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrServer, err)
	}

	var responseTranscripts []types.StudentTranscripts
	for _, t := range transcripts {
		responseTranscripts = append(responseTranscripts, types.StudentTranscripts{
			StudentID: t.StudentID,
			Name:      t.Name,
			Class:     t.Class,
			Ranking:   t.Ranking,
			AvgGPA:    t.AvgGPA,
		})
	}

	return &types.GetStudentTranscriptsResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: responseTranscripts,
	}, nil
}
