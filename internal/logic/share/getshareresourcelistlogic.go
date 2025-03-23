package share

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"

	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShareResourceListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetShareResourceListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShareResourceListLogic {
	return &GetShareResourceListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetShareResourceListLogic) GetShareResourceList(req *types.GetShareResourceListReq) (resp *types.GetShareResourceListResp, err error) {
	resourceType := req.ResourceType

	posts, err := service.GetShareResourceList(resourceType)
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
			CoverImg:  p.CoverImg,
			Title:     p.Title,
			Content:   p.Content,
			CreatedAt: p.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: p.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &types.GetShareResourceListResp{
		Base: types.Base{
			Code: 200,
			Msg:  "ok",
		},
		Data: experiencePosts,
	}, nil
}
