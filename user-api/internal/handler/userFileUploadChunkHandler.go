package handler

import (
	"errors"
	"fmt"
	"net/http"

	"cloud-disk/user-api/helper"
	"cloud-disk/user-api/internal/logic"
	"cloud-disk/user-api/internal/svc"
	"cloud-disk/user-api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func userFileUploadChunkHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserFileUploadChunkReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		if r.PostForm.Get("key") == "" {
			httpx.Error(w, errors.New("key is empty"))
		}
		if r.PostForm.Get("upload_id") == "" {
			httpx.Error(w, errors.New("upload_id is empty"))
		}
		if r.PostForm.Get("part_number") == "" {
			httpx.Error(w, errors.New("part_number is empty"))
		}

		md5, err := helper.CosPartUpload(r)
		if err != nil {
			httpx.Error(w, errors.New("cos is shit!"))
			fmt.Println(err)
		}

		l := logic.NewUserFileUploadChunkLogic(r.Context(), svcCtx)
		resp, err := l.UserFileUploadChunk(&req)
		resp = new(types.UserFileUploadChunkResp)
		resp.Etag = md5
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
