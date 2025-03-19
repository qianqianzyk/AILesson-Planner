package disk

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/model"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CollectFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCollectFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CollectFileLogic {
	return &CollectFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CollectFileLogic) CollectFile(req *types.CollectFileReq) (resp *types.CollectFileResp, err error) {
	ids := req.Ids

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

	var fs []*model.File
	var ds []*model.File
	for _, f := range files {
		if f.IsDir {
			ds = append(ds, f)
		} else {
			fs = append(fs, f)
		}
		f.IsCollect = !f.IsCollect
	}
	err = service.UpdateFilesCollect(fs, ds)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrCollectFile, err)
	}

	return &types.CollectFileResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: "",
	}, nil
}
