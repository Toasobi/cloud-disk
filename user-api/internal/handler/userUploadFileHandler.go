package handler

import (
	"context"
	"crypto/md5"
	"database/sql"
	"fmt"
	"net/http"
	"path"

	"cloud-disk/user-api/helper"
	"cloud-disk/user-api/internal/logic"
	"cloud-disk/user-api/internal/svc"
	"cloud-disk/user-api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func userUploadFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			fmt.Println(err)
			return
		}

		//生成hash
		d := make([]byte, fileHeader.Size)
		file.Read(d)
		hash := fmt.Sprintf("%x", md5.Sum(d))

		//查找文件是否已经存在
		var ctx context.Context

		has, err := svcCtx.RepositoryPool.FindOneByHash(ctx, hash)
		if err != nil {
			if err != sql.ErrNoRows {
				fmt.Println("数据库返回错误")
				fmt.Println(err)
				return
			}
		}
		if has != nil {
			httpx.OkJson(w, &types.FileUploadResp{
				Identity: has.Identity.String,
				Ext:      has.Ext.String,
				Name:     has.Name.String,
			})
			return
		}

		//若是没有该文件则去logic添加至数据库
		//往cos中存储文件
		cosPath, err := helper.CosUpload(r)
		if err != nil {
			fmt.Println("文件上传错误")
			return
		}

		//往logic里面传递
		req.Name = fileHeader.Filename
		req.Ext = path.Ext(fileHeader.Filename)
		req.Size = fileHeader.Size
		req.Hash = hash
		req.Path = cosPath

		l := logic.NewUserUploadFileLogic(r.Context(), svcCtx)
		resp, err := l.UserUploadFile(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
