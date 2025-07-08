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

type StoreResourceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStoreResourceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StoreResourceLogic {
	return &StoreResourceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StoreResourceLogic) StoreResource(req *types.StoreResourceReq) (resp *types.StoreResourceResp, err error) {
	ids := req.Ids
	parentID := req.ParentID
	path := req.Path

	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

	if len(ids) > 0 {
		for _, id := range ids {
			file, err := service.GetFileByID(id)
			if err != nil {
				return nil, utils.AbortWithException(utils.ErrGetFile, err)
			}

			if file.IsDir {
				err = service.StoreDirectoryContent(file, int(userID), parentID, path)
				if err != nil {
					return nil, utils.AbortWithException(utils.ErrCreateFile, err)
				}
			} else {
				newFile := model.File{
					UserID:    int(userID),
					Name:      file.Name,
					Path:      path,
					Size:      file.Size,
					FileType:  file.FileType,
					FileUrl:   file.FileUrl,
					IsDir:     file.IsDir,
					IsCollect: false,
					ParentID:  parentID,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}
				err = service.CreateFile(&newFile)
				if err != nil {
					return nil, utils.AbortWithException(utils.ErrCreateFile, err)
				}
			}
		}
	}

	return &types.StoreResourceResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: "",
	}, nil
}
