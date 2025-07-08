package disk

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRecycleDiskFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRecycleDiskFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRecycleDiskFileLogic {
	return &GetRecycleDiskFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRecycleDiskFileLogic) GetRecycleDiskFile(req *types.Empty) (resp *types.GetRecycleDiskFileResp, err error) {
	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

	files, err := service.GetRecycleFileList(int(userID))
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrGetFile, err)
	}

	var ryFiles []types.RecycleDiskFile
	for _, f := range files {
		ryFiles = append(ryFiles, types.RecycleDiskFile{
			ID:        int(f.ID),
			Name:      f.Name,
			Size:      service.FormatFileSize(int64(f.Size)),
			DeletedAt: f.DeletedAt.Time.Format("2006-01-02 15:04:05"),
		})
	}

	return &types.GetRecycleDiskFileResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: ryFiles,
	}, nil
}
