package share

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteShareResourceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteShareResourceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteShareResourceLogic {
	return &DeleteShareResourceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteShareResourceLogic) DeleteShareResource(req *types.DeleteShareResourceReq) (resp *types.DeleteShareResourceResp, err error) {
	resourceID := req.ResourceID

	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

	err = service.DeleteShareResource(resourceID, int(userID))
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrDeleteExperiencePost, err)
	}

	return &types.DeleteShareResourceResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: "",
	}, nil
}
