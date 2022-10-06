package handler

import (
	"net/http"

	"cloud-disk/user-api/internal/logic"
	"cloud-disk/user-api/internal/svc"
	"cloud-disk/user-api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func userFolderCreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserFolderCreateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserFolderCreateLogic(r.Context(), svcCtx)
		resp, err := l.UserFolderCreate(&req, r.Header.Get("UserIdentity"))
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
