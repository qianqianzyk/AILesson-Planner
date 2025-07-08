package share

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteShareFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteShareFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteShareFileLogic {
	return &DeleteShareFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteShareFileLogic) DeleteShareFile(req *types.DeleteShareFileReq) (resp *types.DeleteShareFileResp, err error) {
	fileUrl := req.FileUrl

	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

	err = service.DeleteAttachmentByUrl(int(userID), fileUrl)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrServer, err)
	}
	service.DeleteObjectByUrlAsync(fileUrl)

	return &types.DeleteShareFileResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: "",
	}, nil
}
