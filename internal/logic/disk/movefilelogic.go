package disk

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MoveFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMoveFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MoveFileLogic {
	return &MoveFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MoveFileLogic) MoveFile(req *types.MoveFileReq) (resp *types.MoveFileResp, err error) {
	ids := req.Ids
	parentID := req.ParentID

	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

	files, ok, err := service.CheckFilesExistenceByIDs(ids, int(userID))
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrServer, err)
	}
	if !ok {
		return nil, utils.AbortWithException(utils.ErrFile, err)
	}

	for _, f := range files {
		f.ParentID = parentID
	}
	err = service.UpdateFilesMove(files)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrMoveFile, err)
	}

	return &types.MoveFileResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: "",
	}, nil
}
