package disk

import (
	"context"
	"fmt"
	"github.com/qianqianzyk/AILesson-Planner/internal/model"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"
	"time"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MergeFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMergeFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MergeFileLogic {
	return &MergeFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MergeFileLogic) MergeFile(req *types.MergeFileReq) (resp *types.MergeFileResp, err error) {
	fileName := req.FileName
	md5 := req.MD5
	uploadID := req.UploadID
	contentType := req.ContentType
	parentID := req.ParentID
	path := req.Path

	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

	fileUrl, size, err := service.CompleteMultipartUpload(md5, fileName, uploadID, contentType)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrMergeFile, err)
	}

	attachment, err := service.GetAttachmentByUploadID(md5, fileName, uploadID)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrServer, err)
	}

	attachment.UserID = int(userID)
	attachment.FileUrl = fileUrl
	err = service.UpdateAttachment(attachment)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrServer, err)
	}

	fileType, fType := service.JudgeFileType(fileName)

	newFile := model.File{
		UserID:    int(userID),
		Name:      fileName,
		Path:      path,
		Size:      int(size),
		FileType:  fileType,
		FType:     fType,
		FileUrl:   fileUrl,
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

	return &types.MergeFileResp{
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
	}, nil
}
