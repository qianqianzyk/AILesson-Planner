package disk

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/model"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"
	"time"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDiskDirectoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateDiskDirectoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDiskDirectoryLogic {
	return &CreateDiskDirectoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateDiskDirectoryLogic) CreateDiskDirectory(req *types.CreateDiskDirectoryReq) (resp *types.CreateDiskDirectoryResp, err error) {
	name := req.Name
	path := req.Path
	parentID := req.ParentID

	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

	ok, err := service.IsFileNameExisted(int(userID), name, parentID)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrServer, err)
	}
	if ok {
		return nil, utils.AbortWithException(utils.ErrFileExisted, err)
	}

	file := model.File{
		UserID:    int(userID),
		Name:      name,
		Path:      path,
		Size:      0,
		FileType:  "文件夹",
		IsDir:     true,
		IsCollect: false,
		ParentID:  parentID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = service.CreateFile(&file)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrCreateFile, err)
	}

	return &types.CreateDiskDirectoryResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: types.DirectoryInfo{
			ID:        file.ID,
			Name:      file.Name,
			FileType:  file.FileType,
			UpdatedAt: file.UpdatedAt.Format("2006-01-02 15:04:05"),
			IsCollect: file.IsCollect,
		},
	}, nil
}
