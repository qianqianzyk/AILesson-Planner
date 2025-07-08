package disk

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDirectoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDirectoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDirectoryLogic {
	return &GetDirectoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDirectoryLogic) GetDirectory(req *types.GetDirectoryReq) (resp *types.GetDirectoryResp, err error) {
	parentID := req.ParentID

	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

	files, err := service.GetDirectoryStructure(int(userID), parentID)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrServer, err)
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

	return &types.GetDirectoryResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: diskDirectory,
	}, nil
}
