package student

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelStudentInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelStudentInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelStudentInfoLogic {
	return &DelStudentInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelStudentInfoLogic) DelStudentInfo(req *types.DelStudentInfoReq) (resp *types.DelStudentInfoResp, err error) {
	studentIDs := req.StudentIDs

	err = service.DeleteStudents(studentIDs)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrDeleteStudents, err)
	}

	return &types.DelStudentInfoResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: "",
	}, nil
}
