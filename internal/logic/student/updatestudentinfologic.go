package student

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStudentInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateStudentInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStudentInfoLogic {
	return &UpdateStudentInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateStudentInfoLogic) UpdateStudentInfo(req *types.UpdateStudentInfoReq) (resp *types.UpdateStudentInfoResp, err error) {
	studentID := req.StudentID
	name := req.Name
	college := req.College
	class := req.Class
	major := req.Major

	student, err := service.GetStudentByStudentID(studentID)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrServer, err)
	}

	student.Name = name
	student.College = college
	student.Class = class
	student.Major = major
	err = service.UpdateStudent(student)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrServer, err)
	}

	return &types.UpdateStudentInfoResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: ""}, nil
}
