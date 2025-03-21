package share

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchShareResourceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchShareResourceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchShareResourceLogic {
	return &SearchShareResourceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchShareResourceLogic) SearchShareResource(req *types.SearchShareResourceReq) (resp *types.SearchShareResourceResp, err error) {
	resourceType := req.ResourceType
	keyword := req.Keyword

	posts, err := service.SearchShareResource(resourceType, keyword)
	if err != nil {
		return nil, utils.AbortWithException(utils.ErrGetExperiencePost, err)
	}

	var experiencePosts []types.ShareResource
	for _, p := range posts {
		user, err := service.GetUserByUserID(uint(p.UserID))
		if err != nil {
			return nil, utils.AbortWithException(utils.ErrServer, err)
		}

		experiencePosts = append(experiencePosts, types.ShareResource{
			ID:        int(p.ID),
			Username:  user.Username,
			Avatar:    user.Avatar,
			Title:     p.Title,
			Content:   p.Content,
			CreatedAt: p.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: p.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &types.SearchShareResourceResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: experiencePosts,
	}, nil
}
