package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"muxi-live-stream-api/internal/logic"
	"muxi-live-stream-api/internal/svc"
	"muxi-live-stream-api/internal/types"
)

// LoginHandler 用户登录接口
// @Summary 用户登录
// @Description 用户通过学号和密码进行登录，获取访问令牌
// @Tags 用户
// @Accept json
// @Produce json
// @Param req body types.LoginRequest true "登录请求参数"
// @Success 200 {object} types.Response "登录成功返回信息"
// @Failure 400 {object} types.Response "登录失败返回信息"
// @Router /library/login [post]
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
