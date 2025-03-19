package disk

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"
	"time"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutDiskDirectoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutDiskDirectoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutDiskDirectoryLogic {
	return &PutDiskDirectoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutDiskDirectoryLogic) PutDiskDirectory(req *types.PutDiskDirectoryReq) (resp *types.PutDiskDirectoryResq, err error) {
	id := req.ID
	name := req.Name

	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

	file, err := service.GetFile(id, int(userID))
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrGetFile, err)
	}

	ok, err := service.IsFileNameExisted(int(userID), name, file.ParentID)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrServer, err)
	}
	if ok {
		return nil, utils.AbortWithException(utils.ErrFileExisted, err)
	}

	file.Name = name
	file.UpdatedAt = time.Now()
	err = service.UpdateFile(file)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUpdateFile, err)
	}

	return &types.PutDiskDirectoryResq{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: "",
	}, nil
}
