package disk

import (
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"
	"net/http"

	"github.com/qianqianzyk/AILesson-Planner/internal/logic/disk"
	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetShareDirectoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetShareDirectoryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := disk.NewGetShareDirectoryLogic(r.Context(), svcCtx)
		resp, err := l.GetShareDirectory(&req)
		if err != nil {
			utils.HandleError(w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
