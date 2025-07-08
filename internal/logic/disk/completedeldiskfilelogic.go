package disk

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CompleteDelDiskFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCompleteDelDiskFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CompleteDelDiskFileLogic {
	return &CompleteDelDiskFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CompleteDelDiskFileLogic) CompleteDelDiskFile(req *types.CompleteDelDiskFileReq) (resp *types.CompleteDelDiskFileResp, err error) {
	ids := req.Ids

	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

	err = service.CompleteDelFile(ids, int(userID))
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrDelFile, err)
	}

	return &types.CompleteDelDiskFileResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: "",
	}, nil
}
