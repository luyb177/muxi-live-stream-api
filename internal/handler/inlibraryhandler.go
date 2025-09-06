package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"muxi-live-stream-api/internal/logic"
	"muxi-live-stream-api/internal/svc"
	"muxi-live-stream-api/internal/types"
)

// InLibraryHandler 查询用户是否在图书馆
// @Summary 查询用户是否在图书馆
// @Description 借助用户Cookie信息查询用户当前是否在图书馆内
// @Tags 用户
// @Accept json
// @Produce json
// @Param Cookie header string true "用户认证Cookie"
// @Param req body types.InLibraryRequest true "请求参数"
// @Success 200 {object} types.Response "查询成功返回用户在馆状态"
// @Failure 400 {object} types.Response "请求参数错误"
// @Failure 401 {string} string "未授权访问"
// @Router /library/inlibrary [post]
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
			httpx.OkJsonCtx(r.Context(), w, resp)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
