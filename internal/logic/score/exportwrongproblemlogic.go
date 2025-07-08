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

type ExportWrongProblemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExportWrongProblemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExportWrongProblemLogic {
	return &ExportWrongProblemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExportWrongProblemLogic) ExportWrongProblem(req *types.Empty) (resp *types.ExportWrongProblemResp, err error) {
	err = service.SendStudentWrongProblem(l.svcCtx.Config.Email.Name, "1831711261@qq.com", l.svcCtx.Config.Email.Key)
	if err != nil {
		if errors.Is(err, utils.ErrTimeLimited) {
			return nil, utils.AbortWithException(utils.ErrSendCodeLimited, err)
		}
		return nil, utils.AbortWithException(utils.ErrServer, err)
	}

	return &types.ExportWrongProblemResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: types.FileUrl{
			Url: "https://disk.qianqianzyk.top/aihelper/scores/2025/docx/浅浅_中国近代史_第1章_错题.docx",
		},
	}, nil
}
