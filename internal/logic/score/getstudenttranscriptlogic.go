package score

import (
	"context"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStudentTranscriptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStudentTranscriptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStudentTranscriptLogic {
	return &GetStudentTranscriptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStudentTranscriptLogic) GetStudentTranscript(req *types.GetStudentTranscriptReq) (resp *types.GetStudentTranscriptResp, err error) {
	return &types.GetStudentTranscriptResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: types.FileUrl{
			Url: "https://disk.qianqianzyk.top/aihelper/transcripts/2025/xlsx/302023311111_2023-2024_1_期末成绩单.xlsx",
		},
	}, nil
}
