package chat

import (
	"context"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLessonResourceListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLessonResourceListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLessonResourceListLogic {
	return &GetLessonResourceListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLessonResourceListLogic) GetLessonResourceList(req *types.Empty) (resp *types.GetLessonResourceListResp, err error) {
	mockData := map[string]interface{}{
		"resource_list": []map[string]interface{}{
			{
				"resource_name":      "中国近现代史纲要：2023 年版",
				"resource_url":       "https://disk.qianqianzyk.top/aihelper/disk/2025/pdf/448d2d36-0d39-11f0-989a-00163e0daf65.pdf",
				"resource_size":      "50.43MB",
				"resource_cover_url": "https://disk.qianqianzyk.top/aihelper/share/2025/png/5bfcb365-1089-11f0-9c2d-00163e0daf65.png",
			},
			{
				"resource_name":      "大学物理：第二版",
				"resource_url":       "",
				"resource_size":      "",
				"resource_cover_url": "https://disk.qianqianzyk.top/aihelper/share/2025/jpg/ee649b39-1089-11f0-9c2d-00163e0daf65.jpg",
			},
			{
				"resource_name":      "学术英语",
				"resource_url":       "",
				"resource_size":      "",
				"resource_cover_url": "https://disk.qianqianzyk.top/aihelper/share/2025/png/0232840d-30c3-11f0-9a72-04bf1b6f946c.png",
			},
			{
				"resource_name":      "C++程序设计：第三版",
				"resource_url":       "",
				"resource_size":      "",
				"resource_cover_url": "https://disk.qianqianzyk.top/aihelper/share/2025/png/754d5c6e-108a-11f0-9c2d-00163e0daf65.png",
			},
			{
				"resource_name":      "高等数学（上册）：第七版",
				"resource_url":       "",
				"resource_size":      "",
				"resource_cover_url": "https://disk.qianqianzyk.top/aihelper/share/2025/png/add13ebe-108a-11f0-9c2d-00163e0daf65.png",
			},
			{
				"resource_name":      "数据结构教程（C++语言描述）：第二版",
				"resource_url":       "",
				"resource_size":      "",
				"resource_cover_url": "https://disk.qianqianzyk.top/aihelper/share/2025/png/ffeff0d8-108a-11f0-9c2d-00163e0daf65.png",
			},
			{
				"resource_name":      "高等数学（下册）：第七版",
				"resource_url":       "",
				"resource_size":      "",
				"resource_cover_url": "https://disk.qianqianzyk.top/aihelper/share/2025/png/4758479c-108b-11f0-9c2d-00163e0daf65.png",
			},
			{
				"resource_name":      "高等数学辅导及习题精解（下册）：同济 第七版",
				"resource_url":       "",
				"resource_size":      "",
				"resource_cover_url": "https://disk.qianqianzyk.top/aihelper/share/2025/png/95f0b89d-108b-11f0-9c2d-00163e0daf65.png",
			},
		},
	}

	return &types.GetLessonResourceListResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: mockData,
	}, nil
}
