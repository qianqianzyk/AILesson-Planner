package disk

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFilesByTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFilesByTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFilesByTypeLogic {
	return &GetFilesByTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFilesByTypeLogic) GetFilesByType(req *types.GetFilesByTypeReq) (resp *types.GetFilesByTypeResp, err error) {
	fileType := req.FileType
	pageNum := req.PageNum
	pageSize := req.PageSize

	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

	files, totalSize, err := service.GetFileListByType(int(userID), fileType, pageNum, pageSize)
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

	return &types.GetFilesByTypeResp{
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
