package chat

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetChatListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetChatListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChatListLogic {
	return &GetChatListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetChatListLogic) GetChatList(req *types.Empty) (resp *types.GetChatListResp, err error) {
	userID, err := service.GetUserID(l.ctx)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrUserID, err)
	}

	cov, err := service.GetTopicList(userID)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrGetTopicList, err)
	}

	var chatList []types.ChatList
	for _, c := range cov {
		chatList = append(chatList, types.ChatList{
			ID:        c.ID,
			Title:     c.Title,
			CreatedAt: c.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: c.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &types.GetChatListResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: chatList,
	}, nil
}
