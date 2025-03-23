package chat

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"
	"time"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLessonPlanLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLessonPlanLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLessonPlanLogic {
	return &UpdateLessonPlanLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLessonPlanLogic) UpdateLessonPlan(req *types.UpdateLessonPlanReq) (resp *types.UpdateLessonPlanResp, err error) {
	id := req.ID
	subject := req.Subject
	textBookName := req.TextBookName
	topicName := req.TopicName
	topicHours := req.TopicHours
	templateFile := req.TemplateFile
	resourceFile := req.ResourceFile
	textBookImg := req.TextBookImg
	description := req.Description
	tPlanContent := req.TPlanContent
	tPlanUrl := req.TPlanUrl

	tPlan, err := service.GetTPlanByID(id)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrServer, err)
	}

	previousUrls := service.FormatUrls(tPlan.TemplateFile, tPlan.ResourceFile, tPlan.TextBookImg, tPlan.TPlanUrl)
	nowUrls := service.FormatUrls(templateFile, resourceFile, textBookImg, tPlanUrl)
	deletedUrls := service.ExtractDelUrls(previousUrls, nowUrls)
	service.DeleteFiles(deletedUrls)

	tPlan.Subject = subject
	tPlan.TextBookName = textBookName
	tPlan.TopicHours = topicHours
	tPlan.TopicName = topicName
	tPlan.TemplateFile = templateFile
	tPlan.ResourceFile = resourceFile
	tPlan.TextBookImg = textBookImg
	tPlan.Description = description
	tPlan.TPlanContent = tPlanContent
	tPlan.TPlanUrl = tPlanUrl
	tPlan.UpdatedAt = time.Now()
	err = service.UpdateTPlan(tPlan)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUpdateTPlan, err)
	}

	return &types.UpdateLessonPlanResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: "",
	}, nil
}
