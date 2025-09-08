package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"muxi-live-stream-api/internal/logic"
	"muxi-live-stream-api/internal/svc"
	"muxi-live-stream-api/internal/types"
)

// GetAllSeatInformationHandler 获取所有座位信息
// @Summary 获取所有座位信息
// @Description 获取系统中所有座位的信息
// @Tags 图书馆信息
// @Accept json
// @Produce json
// @Param Cookie header string false "用户认证Cookie"
// @Param req body types.GetAllSeatInformationRequest true "请求参数"
// @Success 200 {object} types.Response "返回所有座位信息"
// @Failure 401 {string} string "用户未授权"
// @Failure 500 {string} string "服务器内部错误"
// @Router /library/seatinfo [post]
func GetAllSeatInformationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie := r.Header.Get("Cookie")
		var req types.GetAllSeatInformationRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetAllSeatInformationLogic(r.Context(), svcCtx)
		resp, err := l.GetAllSeatInformation(&req, cookie)
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, resp)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
