package share

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

type StoreShareResourceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStoreShareResourceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StoreShareResourceLogic {
	return &StoreShareResourceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StoreShareResourceLogic) StoreShareResource(req *types.StoreShareResourceReq) (resp *types.StoreShareResourceResp, err error) {
	fileUrl := req.FileUrl
	subjectName := req.SubjectName

	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

	file, err := service.GetAttachmentByUrl(fileUrl)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrGetFile, err)
	}

	var parentID int
	var path string
	if file.FType == model.DocumentType {
		parentID, path, err = service.EnsureCourseResourceFolder(int(userID), subjectName, model.DocumentType)
		if err != nil {
			if err != nil {
				return nil, utils.AbortWithException(utils.ErrServer, err)
			}
		}
	} else if file.FType == model.VideoType {
		parentID, path, err = service.EnsureCourseResourceFolder(int(userID), subjectName, model.VideoType)
		if err != nil {
			if err != nil {
				return nil, utils.AbortWithException(utils.ErrServer, err)
			}
		}
	}

	newFile := model.File{
		UserID:    int(userID),
		Name:      file.FileName,
		Path:      path,
		Size:      file.Size,
		FileType:  file.FileType,
		FileUrl:   file.FileUrl,
		IsDir:     false,
		IsCollect: false,
		ParentID:  parentID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = service.CreateFile(&newFile)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrCreateFile, err)
	}

	return &types.StoreShareResourceResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: "",
	}, nil
}
