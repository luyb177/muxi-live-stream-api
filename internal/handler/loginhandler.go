package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"muxi-live-stream-api/internal/logic"
	"muxi-live-stream-api/internal/svc"
	"muxi-live-stream-api/internal/types"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, resp)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
