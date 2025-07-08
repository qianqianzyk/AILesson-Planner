package chat

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLessonPlanLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLessonPlanLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLessonPlanLogic {
	return &GetLessonPlanLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLessonPlanLogic) GetLessonPlan(req *types.GetLessonPlanReq) (resp *types.GetLessonPlanResp, err error) {
	messageID := req.MessageID

	tPlan, err := service.GetTPlanByMessageID(messageID)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrServer, err)
	}

	resourceFile := service.ExtractUrlsFromString(tPlan.ResourceFile)
	textBookImg := service.ExtractUrlsFromString(tPlan.TextBookImg)
	responseTPlan := types.LessonPlan{
		MessageID:    messageID,
		Subject:      tPlan.Subject,
		TextBookName: tPlan.TextBookName,
		TopicHours:   tPlan.TopicHours,
		TopicName:    tPlan.TopicName,
		TemplateFile: tPlan.TemplateFile,
		ResourceFile: resourceFile,
		TextBookImg:  textBookImg,
		Description:  tPlan.Description,
		TPlanUrl:     tPlan.TPlanUrl,
	}

	return &types.GetLessonPlanResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: responseTPlan,
	}, nil
}
