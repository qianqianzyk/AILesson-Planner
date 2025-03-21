package share

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/model"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"
	"go.uber.org/zap"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadShareFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadShareFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadShareFileLogic {
	return &UploadShareFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadShareFileLogic) UploadShareFile(w http.ResponseWriter, r *http.Request, req *types.Empty) (resp *types.UploadShareFileResp, err error) {
	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

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

	objectKey := service.GenerateObjectKey("share", service.GetFileExt(header.Filename))
	objectUrl, err := service.PutObject(objectKey, file, fileSize, header.Header.Get("Content-Type"))
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrFileUpload, err)
	}

	fileType, fType := service.JudgeFileType(header.Filename)

	err = service.CreateAttachment(&model.Attachment{
		UserID:    int(userID),
		FileName:  header.Filename,
		FileUrl:   objectUrl,
		Size:      int(fileSize),
		FileType:  fileType,
		FType:     fType,
		CreatedAt: time.Now(),
	})
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrServer, err)
	}

	resp = &types.UploadShareFileResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: types.FileUrl{
			Url: objectUrl,
		},
	}
	return resp, nil
}
