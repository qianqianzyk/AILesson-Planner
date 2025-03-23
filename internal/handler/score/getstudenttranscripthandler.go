package score

import (
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"
	"net/http"

	"github.com/qianqianzyk/AILesson-Planner/internal/logic/score"
	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetStudentTranscriptHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetStudentTranscriptReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := score.NewGetStudentTranscriptLogic(r.Context(), svcCtx)
		resp, err := l.GetStudentTranscript(&req)
		if err != nil {
			utils.HandleError(w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
