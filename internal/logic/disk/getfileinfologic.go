package disk

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFileInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFileInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFileInfoLogic {
	return &GetFileInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFileInfoLogic) GetFileInfo(req *types.GetDiskFileInfoReq) (resp *types.GetDiskFileInfoResp, err error) {
	id := req.ID

	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

	file, err := service.GetFile(id, int(userID))
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrGetFile, err)
	}

	if file.IsDir {
		stats, err := service.GetDirectoryStats(int64(id))
		if err != nil {
			return nil, utils.AbortWithException(utils.ErrFileStats, err)
		}

		return &types.GetDiskFileInfoResp{
			Base: types.Base{
				Code: 200,
				Msg:  "ok",
			},
			Data: types.FileInfo{
				Name:           file.Name,
				FileType:       file.FileType,
				Size:           service.FormatFileSize(stats.TotalSize),
				FileCount:      int(stats.FileCount),
				DirectoryCount: int(stats.DirectoryCount),
				UpdatedAt:      file.UpdatedAt.Format("2006-01-02 15:04:05"),
				IsCollect:      file.IsCollect,
			},
		}, nil

	}
	if !file.IsDir {
		return &types.GetDiskFileInfoResp{
			Base: types.Base{
				Code: 200,
				Msg:  "ok",
			},
			Data: types.FileInfo{
				Name:      file.Name,
				FileType:  file.FileType,
				Path:      file.Path,
				Size:      service.FormatFileSize(int64(file.Size)),
				UpdatedAt: file.UpdatedAt.Format("2006-01-02 15:04:05"),
				FileUrl:   file.FileUrl,
			},
		}, nil
	}

	return
}
