package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"muxi-live-stream-api/internal/logic"
	"muxi-live-stream-api/internal/svc"
	"muxi-live-stream-api/internal/types"
)

func InLibraryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie := r.Header.Get("Cookie")
		var req types.InLibraryRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewInLibraryLogic(r.Context(), svcCtx)
		resp, err := l.InLibrary(&req, cookie)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
