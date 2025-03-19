package user

import (
	"context"
	"errors"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"
	"go.uber.org/zap"
	"image"
	"mime/multipart"
	"net/http"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadFileLogic {
	return &UploadFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadFileLogic) UploadFile(w http.ResponseWriter, r *http.Request, req *types.UploadFileReq) (resp *types.UploadFileResp, err error) {
	fileType := req.Type

	file, header, err := r.FormFile("file")
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrParam, err)
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			zap.L().Warn("文件关闭错误", zap.Error(err))
		}
	}(file)

	fileSize := header.Size
	//buf, err := io.ReadAll(file)
	//if err != nil {
	//	return nil, utils.AbortWithException(utils.ErrServer, err)
	//}

	// 上传头像 & 课程图片
	if fileType == 1 {
		if fileSize > service.ImageLimit {
			return nil, utils.AbortWithException(utils.ErrAvatarLimited, err)
		}

		reader, size, err := service.ConvertToWebP(file)
		if errors.Is(err, image.ErrFormat) {
			return nil, utils.AbortWithException(utils.ErrFileNotImage, err)
		}
		if err != nil {
			return nil, utils.AbortWithException(utils.ErrServer, err)
		}

		objectKey := service.GenerateObjectKey("image", ".webp")
		objectUrl, err := service.PutObject(objectKey, reader, size, "image/webp")
		if err != nil {
			return nil, utils.AbortWithException(utils.ErrAvatarUpload, err)
		}

		return &types.UploadFileResp{
			Base: types.Base{
				Code: 200,
				Msg:  "ok",
			},
			Data: types.FileUrl{
				Url: objectUrl,
			},
		}, nil
	}

	return
}
