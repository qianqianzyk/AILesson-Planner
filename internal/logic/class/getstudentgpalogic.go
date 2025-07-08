package class

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStudentGPALogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStudentGPALogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStudentGPALogic {
	return &GetStudentGPALogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStudentGPALogic) GetStudentGPA(req *types.GetStudentTermGPAReq) (resp *types.GetStudentTermGPAResp, err error) {
	studentID := req.StudentID

	termGPAs, err := service.GetStudentGPAAndRank(studentID)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrGetStudentGPA, err)
	}

	var gpas []types.TermGPA
	for _, g := range termGPAs {
		gpas = append(gpas, types.TermGPA{
			AcademicYear: g.AcademicYear,
			AcademicTerm: g.AcademicTerm,
			AvgGPA:       g.AvgGPA,
			Rank:         g.Rank,
			Percentile:   g.Percentile,
		})
	}

	return &types.GetStudentTermGPAResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: gpas,
	}, nil
}
