package disk

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDiskFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteDiskFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDiskFileLogic {
	return &DeleteDiskFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteDiskFileLogic) DeleteDiskFile(req *types.DeleteDiskFileReq) (resp *types.DeleteDiskFileResp, err error) {
	ids := req.Ids

	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

	_, ok, err := service.CheckFilesExistenceByIDs(ids, int(userID))
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrServer, err)
	}
	if !ok {
		return nil, utils.AbortWithException(utils.ErrFile, err)
	}

	err = service.DeleteFile(ids, int(userID))
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrDelFile, err)
	}

	return &types.DeleteDiskFileResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: "",
	}, nil
}
