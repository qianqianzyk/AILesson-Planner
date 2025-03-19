package disk

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchFileLogic {
	return &SearchFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchFileLogic) SearchFile(req *types.SearchFilesReq) (resp *types.SearchFilesResp, err error) {
	key := req.Key

	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

	ids, err := service.SearchFiles(key, int(userID))
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrSearchFile, err)
	}

	files, ok, err := service.CheckFilesExistenceByIDs(ids, int(userID))
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrServer, err)
	}
	if !ok {
		return nil, utils.AbortWithException(utils.ErrFile, err)
	}

	var diskDirectory []types.DiskDirectory
	for _, f := range files {
		diskDirectory = append(diskDirectory, types.DiskDirectory{
			ID:        f.ID,
			Name:      f.Name,
			Path:      f.Path,
			Size:      service.FormatFileSize(int64(f.Size)),
			FileType:  f.FileType,
			FileUrl:   f.FileUrl,
			IsDir:     f.IsDir,
			UpdatedAt: f.UpdatedAt.Format("2006-01-02 15:04:05"),
			IsCollect: f.IsCollect,
		})
	}

	return &types.SearchFilesResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: diskDirectory,
	}, nil
}
