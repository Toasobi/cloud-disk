package handler

import (
	"net/http"

	"cloud-disk/user-api/internal/logic"
	"cloud-disk/user-api/internal/svc"
	"cloud-disk/user-api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func userRefreshTokenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserRefreshTokenReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserRefreshTokenLogic(r.Context(), svcCtx)
		resp, err := l.UserRefreshToken(&req, r.Header.Get("Authorization"))
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
