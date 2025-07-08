package disk

import (
	"context"
	"fmt"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type GetShareLinkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetShareLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShareLinkLogic {
	return &GetShareLinkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetShareLinkLogic) GetShareLink(req *types.GetShareLinkReq) (resp *types.GetShareLinkResp, err error) {
	linkStr := req.Link
	code := req.Code

	link, err := service.GetLink(linkStr)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrGetLink, err)
	}

	if link.ExpiresAt.Before(time.Now()) {
		return nil, utils.AbortWithException(utils.ErrExpireLink, fmt.Errorf("share link has expired"))
	}

	if link.ShareCode != code {
		return nil, utils.AbortWithException(utils.ErrLinkCode, fmt.Errorf("share code is wrong"))
	}

	ids, err := service.GetFileIDs(link.FileIDs)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrServer, err)
	}

	files, err := service.GetFilesByIDs(ids)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrServer, err)
	}

	var diskDirectory []types.DiskDirectory
	for _, f := range files {
		diskDirectory = append(diskDirectory, types.DiskDirectory{
			ID:        f.ID,
			Name:      f.Name,
			Path:      f.Path,
			Size:      service.FormatFileSize(int64(f.Size)),
			FileType:  f.FileType,
			FileUrl:   f.FileUrl,
			IsDir:     f.IsDir,
			UpdatedAt: f.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &types.GetShareLinkResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: diskDirectory,
	}, nil
}
