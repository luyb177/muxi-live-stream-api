package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"muxi-live-stream-api/internal/logic"
	"muxi-live-stream-api/internal/svc"
)

func GetRoomIdsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookieHeader := r.Header.Get("Cookie")

		l := logic.NewGetRoomIdsLogic(r.Context(), svcCtx)
		resp, err := l.GetRoomIds(cookieHeader)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
