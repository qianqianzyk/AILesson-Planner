package disk

import (
	"context"
	"fmt"
	"github.com/qianqianzyk/AILesson-Planner/internal/model"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"
	"go.uber.org/zap"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadDiskFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadDiskFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadDiskFileLogic {
	return &UploadDiskFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadDiskFileLogic) UploadDiskFile(w http.ResponseWriter, r *http.Request, req *types.UploadDiskFileReq) (resp *types.UploadDiskFileResp, err error) {
	parentID := req.ParentID
	path := req.Path

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

	objectKey := service.GenerateObjectKey("disk", service.GetFileExt(header.Filename))
	objectUrl, err := service.PutObject(objectKey, file, fileSize, header.Header.Get("Content-Type"))
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrFileUpload, err)
	}

	fileType, fType := service.JudgeFileType(header.Filename)

	newFile := model.File{
		UserID:    int(userID),
		Name:      header.Filename,
		Path:      path,
		Size:      int(fileSize),
		FileType:  fileType,
		FType:     fType,
		FileUrl:   objectUrl,
		IsDir:     false,
		ParentID:  parentID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsCollect: false,
	}
	err = service.CreateFile(&newFile)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrCreateFile, err)
	}

	err = service.CreateAttachment(&model.Attachment{
		UserID:    int(userID),
		FileUrl:   objectUrl,
		CreatedAt: time.Now(),
	})
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrServer, err)
	}

	resp = &types.UploadDiskFileResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: types.FileInfo{
			ID:        fmt.Sprintf("%d", newFile.ID),
			Name:      newFile.Name,
			Size:      service.FormatFileSize(int64(newFile.Size)),
			FileType:  newFile.FileType,
			UpdatedAt: newFile.UpdatedAt.Format("2006-01-02 15:04:05"),
			FileUrl:   newFile.FileUrl,
			IsCollect: newFile.IsCollect,
		},
	}
	return resp, nil
}
