package share

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

type CreateShareResourceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateShareResourceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateShareResourceLogic {
	return &CreateShareResourceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateShareResourceLogic) CreateShareResource(req *types.CreateShareResourceReq) (resp *types.CreateShareResourceResp, err error) {
	resourceType := req.ResourceType
	title := req.Title
	content := req.Content

	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

	post := model.ShareResource{
		UserID:       int(userID),
		ResourceType: resourceType,
		Title:        title,
		Content:      content,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	err = service.CreateShareResource(&post)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrCreateExperiencePost, err)
	}

	return &types.CreateShareResourceResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: "",
	}, nil
}
