package disk

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckFileMD5Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckFileMD5Logic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckFileMD5Logic {
	return &CheckFileMD5Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckFileMD5Logic) CheckFileMD5(req *types.CheckFileMD5Req) (resp *types.CheckFileMD5Resp, err error) {
	md5 := req.MD5
	fileName := req.FileName

	status, uploadID, uploadedParts, err := service.CheckFileMD5(md5, fileName)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrCheckFileMD5, err)
	}

	var minioObjectPart []types.MinioObjectPart
	for _, m := range uploadedParts {
		minioObjectPart = append(minioObjectPart, types.MinioObjectPart{
			PartNumber: m.PartNumber,
			ETag:       m.ETag,
		})
	}

	return &types.CheckFileMD5Resp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: types.FileMD5{
			Status:     status,
			UploadID:   uploadID,
			ObjectPart: minioObjectPart,
		},
	}, nil
}
