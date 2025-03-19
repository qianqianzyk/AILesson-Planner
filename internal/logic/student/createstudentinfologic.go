package student

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/model"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateStudentInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateStudentInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateStudentInfoLogic {
	return &CreateStudentInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateStudentInfoLogic) CreateStudentInfo(req *types.CreateStudentInfoReq) (resp *types.CreateStudentInfoResp, err error) {
	courseID := req.CourseID
	studentID := req.StudentID
	name := req.Name
	college := req.College
	class := req.Class
	major := req.Major

	_, err = service.GetStudentByStudentID(studentID)
	if err == nil {
		return nil, utils.AbortWithException(utils.ErrStudentExist, err)
	}

	student := model.Student{
		StudentID: studentID,
		Name:      name,
		College:   college,
		Class:     class,
		Major:     major,
	}
	err = service.CreateStudent(&student)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrServer, err)
	}

	score := model.Score{
		StudentID:    studentID,
		CourseID:     courseID,
		RegularScore: 0,
		FinalScore:   0,
		TotalScore:   0,
		CreditEarned: false,
		GradePoint:   0,
	}
	err = service.CreateScore(&score)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrServer, err)
	}

	return &types.CreateStudentInfoResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: "",
	}, nil
}
