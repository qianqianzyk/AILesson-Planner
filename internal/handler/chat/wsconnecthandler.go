package chat

import (
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"
	"net/http"

	"github.com/qianqianzyk/AILesson-Planner/internal/logic/chat"
	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func WsConnectHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := chat.NewWsConnectLogic(r.Context(), svcCtx)
		err := l.WsConnect(w, r)
		if err != nil {
			utils.HandleError(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
