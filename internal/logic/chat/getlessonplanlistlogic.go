package chat

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLessonPlanListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLessonPlanListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLessonPlanListLogic {
	return &GetLessonPlanListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLessonPlanListLogic) GetLessonPlanList(req *types.Empty) (resp *types.GetLessonPlanListResp, err error) {
	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

	tPlans, err := service.GetTPlanList(int(userID))
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrGetTPlan, err)
	}

	var responseTPlans []types.LessonPlan
	for _, t := range tPlans {
		resourceFile := service.ExtractUrlsFromString(t.ResourceFile)
		textBookImg := service.ExtractUrlsFromString(t.TextBookImg)

		responseTPlans = append(responseTPlans, types.LessonPlan{
			ID:           int(t.ID),
			Subject:      t.Subject,
			TextBookName: t.TextBookName,
			TopicHours:   t.TopicHours,
			TopicName:    t.TopicName,
			TemplateFile: t.TemplateFile,
			ResourceFile: resourceFile,
			TextBookImg:  textBookImg,
			Description:  t.Description,
			TPlanContent: t.TPlanContent,
			TPlanUrl:     t.TPlanUrl,
		})
	}

	return &types.GetLessonPlanListResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: responseTPlans,
	}, nil
}
