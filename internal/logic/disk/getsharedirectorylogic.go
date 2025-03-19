package disk

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShareDirectoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetShareDirectoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShareDirectoryLogic {
	return &GetShareDirectoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetShareDirectoryLogic) GetShareDirectory(req *types.GetShareDirectoryReq) (resp *types.GetShareDirectoryResp, err error) {
	parentID := req.ParentID
	pageNum := req.PageNum
	pageSize := req.PageSize

	files, totalSize, err := service.GetFileListByParentID(parentID, pageNum, pageSize)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrGetFile, err)
	}

	var diskDirectory []types.DiskDirectory
	for _, f := range files {
		diskDirectory = append(diskDirectory, types.DiskDirectory{
			ID:        f.ID,
			Name:      f.Name,
			Size:      service.FormatFileSize(int64(f.Size)),
			FileType:  f.FileType,
			FileUrl:   f.FileUrl,
			IsDir:     f.IsDir,
			UpdatedAt: f.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &types.GetShareDirectoryResp{
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
