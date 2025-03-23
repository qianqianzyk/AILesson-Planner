package chat

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/model"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"
	"time"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateLessonPlanLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGenerateLessonPlanLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateLessonPlanLogic {
	return &GenerateLessonPlanLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenerateLessonPlanLogic) GenerateLessonPlan(req *types.GenerateLessonPlanReq) (resp *types.GenerateLessonPlanResp, err error) {
	subject := req.Subject
	textBookName := req.TextBookName
	topicName := req.TopicName
	topicHours := req.TopicHours
	templateFile := req.TemplateFile
	resourceFile := req.ResourceFile
	textBookImg := req.TextBookImg
	description := req.Description

	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

	answer, err := service.GenerateLessonPlan(l.svcCtx.Config.AI.TPlanEndpoint, textBookName, subject, topicHours, topicName, templateFile, resourceFile, textBookImg, description)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrGenAnswer, err)
	}

	tPlan := model.TPlan{
		UserID:       int(userID),
		Subject:      subject,
		TextBookName: textBookName,
		TopicHours:   topicHours,
		TopicName:    topicName,
		TemplateFile: templateFile,
		ResourceFile: resourceFile,
		TextBookImg:  textBookImg,
		Description:  description,
		TPlanContent: answer,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	err = service.CreateTPlan(&tPlan)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrGenTPlan, err)
	}

	return &types.GenerateLessonPlanResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: types.ChatAnswer{
			Message: answer,
		},
	}, nil
}
