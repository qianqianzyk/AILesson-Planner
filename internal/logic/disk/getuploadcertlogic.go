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

type GetUploadCertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUploadCertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUploadCertLogic {
	return &GetUploadCertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUploadCertLogic) GetUploadCert(req *types.GetUploadCertReq) (resp *types.GetUploadCertResp, err error) {
	md5 := req.MD5
	fileName := req.FileName
	contentType := req.ContentType

	uploadID, err := service.InitMultipartUpload(md5, fileName, contentType)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrGenFileCert, err)
	}

	attachment := model.Attachment{
		FileName:  fileName,
		MD5:       md5,
		UploadID:  uploadID,
		CreatedAt: time.Now(),
	}
	err = service.CreateAttachment(&attachment)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrServer, err)
	}

	return &types.GetUploadCertResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: types.UploadCert{
			UploadID: uploadID,
		},
	}, nil
}
