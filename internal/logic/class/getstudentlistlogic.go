package class

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStudentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStudentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStudentListLogic {
	return &GetStudentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStudentListLogic) GetStudentList(req *types.GetStudentListReq) (resp *types.GetStudentListResp, err error) {
	class := req.Class

	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

	studentList, err := service.GetStudentsByUserAndClass(int(userID), class)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrGetStudentList, err)
	}

	var students []types.Student
	for _, s := range studentList {
		students = append(students, types.Student{
			StudentID: s.StudentID,
			Class:     s.Class,
			Name:      s.Name,
			Major:     s.Major,
		})
	}

	return &types.GetStudentListResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: students,
	}, nil
}
