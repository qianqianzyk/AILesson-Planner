package chat

import (
	"context"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetExerciseListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetExerciseListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetExerciseListLogic {
	return &GetExerciseListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetExerciseListLogic) GetExerciseList(req *types.Empty) (resp *types.GetExerciseListResp, err error) {
	mockData := map[string]interface{}{
		"resource_list": []map[string]interface{}{
			{
				"resource_name":      "《中国近现代史纲要》习题集",
				"resource_url":       "https://disk.qianqianzyk.top/aihelper/share/2025/pdf/d1606fdc-11c4-11f0-bfdf-00163e0daf65.pdf",
				"resource_size":      "191KB",
				"resource_cover_url": "https://disk.qianqianzyk.top/aihelper/share/2025/png/ffb6b3f6-11c4-11f0-bfdf-00163e0daf65.png",
			},
			{
				"resource_name":      "高等数学（上）2019-2020学年第一学期联考试卷",
				"resource_url":       "",
				"resource_size":      "",
				"resource_cover_url": "https://disk.qianqianzyk.top/aihelper/share/2025/png/82c0f26a-11c5-11f0-bfdf-00163e0daf65.png",
			},
			{
				"resource_name":      "高等数学（下）2019-2020学年第二学期期末试卷",
				"resource_url":       "",
				"resource_size":      "",
				"resource_cover_url": "https://disk.qianqianzyk.top/aihelper/share/2025/png/d2156db9-11c5-11f0-bfdf-00163e0daf65.png",
			},
		},
	}

	return &types.GetExerciseListResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: mockData,
	}, nil
}
