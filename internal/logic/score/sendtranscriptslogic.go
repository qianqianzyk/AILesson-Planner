package score

import (
	"context"
	"errors"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendTranscriptsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendTranscriptsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendTranscriptsLogic {
	return &SendTranscriptsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendTranscriptsLogic) SendTranscripts(req *types.SendtTranscriptsReq) (resp *types.SendtTranscriptsResp, err error) {
	studentID := req.StudentID

	student, err := service.GetStudentByStudentID(studentID)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrServer, err)
	}

	err = service.SendStudentTranscripts(l.svcCtx.Config.Email.Name, student.Email, l.svcCtx.Config.Email.Key)
	if err != nil {
		if errors.Is(err, utils.ErrTimeLimited) {
			return nil, utils.AbortWithException(utils.ErrSendCodeLimited, err)
		}
		return nil, utils.AbortWithException(utils.ErrServer, err)
	}

	return &types.SendtTranscriptsResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: "",
	}, nil
}
