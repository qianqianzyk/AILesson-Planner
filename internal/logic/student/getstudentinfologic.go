package student

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStudentInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStudentInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStudentInfoLogic {
	return &GetStudentInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStudentInfoLogic) GetStudentInfo(req *types.GetStudentInfoReq) (resp *types.GetStudentInfoResp, err error) {
	courseID := req.CourseID
	class := req.Class

	students, err := service.GetStudentsByCourse(courseID, class)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrServer, err)
	}

	var studentsInfo []types.StudentInfo
	for _, s := range students {
		studentsInfo = append(studentsInfo, types.StudentInfo{
			StudentID: s.StudentID,
			Name:      s.Name,
			College:   s.College,
			Class:     s.Class,
			Major:     s.Major,
		})
	}

	return &types.GetStudentInfoResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: studentsInfo,
	}, nil
}
