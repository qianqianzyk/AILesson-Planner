package disk

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDiskDirectoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDiskDirectoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDiskDirectoryLogic {
	return &GetDiskDirectoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDiskDirectoryLogic) GetDiskDirectory(req *types.GetDiskDirectoryReq) (resp *types.GetDiskDirectoryResp, err error) {
	parentID := req.ParentID
	pageNum := req.PageNum
	pageSize := req.PageSize

	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

	files, totalSize, err := service.GetFileList(int(userID), parentID, pageNum, pageSize)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrGetFile, err)
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

	return &types.GetDiskDirectoryResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: types.DiskDirectoryList{
			TotalNum: *totalSize,
			FileList: diskDirectory,
		},
	}, nil
}
