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

type ShareLinkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareLinkLogic {
	return &ShareLinkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareLinkLogic) ShareLink(req *types.ShareLinkGenReq) (resp *types.ShareLinkGenResp, err error) {
	ids := req.Ids
	code := req.Code
	validity := req.Validity

	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

	_, ok, err := service.CheckFilesExistenceByIDs(ids, int(userID))
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrServer, err)
	}
	if !ok {
		return nil, utils.AbortWithException(utils.ErrFile, err)
	}

	formattedIds, err := service.SetFileIDs(ids)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrServer, err)
	}

	if code == "" {
		code, err = service.GenerateCode()
		if err != nil {
			return nil, utils.AbortWithException(utils.ErrGenCode, err)
		}
	} else {
		ok := service.CheckCodeValidity(code)
		if !ok {
			return nil, utils.AbortWithException(utils.ErrFormattedCode, err)
		}
	}

	formattedValidity := service.SetLinkExpiresAt(validity)
	link, err := service.GenerateShareLink()
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrGenLink, err)
	}

	shareLink := model.ShareLink{
		UserID:    int(userID),
		ShareCode: code,
		FileIDs:   formattedIds,
		Link:      link,
		ExpiresAt: formattedValidity,
		CreatedAt: time.Now(),
	}
	err = service.CreateLink(&shareLink)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrServer, err)
	}

	return &types.ShareLinkGenResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: types.Link{
			Code: code,
			Url:  link,
		},
	}, nil
}
