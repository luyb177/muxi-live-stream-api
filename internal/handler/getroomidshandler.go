package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"muxi-live-stream-api/internal/logic"
	"muxi-live-stream-api/internal/svc"
)

// GetRoomIdsHandler 获取所有房间ID
// @Summary 获取所有房间ID
// @Description 获取系统中所有活跃的房间ID列表
// @Tags 图书馆信息
// @Accept json
// @Produce json
// @Param Cookie header string false "用户认证Cookie"
// @Success 200 {object} types.Response "成功返回房间ID列表"
// @Failure 400 {object} types.Response "请求参数错误"
// @Failure 401 {string} string "未授权访问"
// @Router /library/roomids [get]
func GetRoomIdsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookieHeader := r.Header.Get("Cookie")

		l := logic.NewGetRoomIdsLogic(r.Context(), svcCtx)
		resp, err := l.GetRoomIds(cookieHeader)
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, resp)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
