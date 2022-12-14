// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"cloud-disk/user-api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: userLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/register",
				Handler: userRegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/send-code",
				Handler: userSendCodeHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/user/upload",
					Handler: userUploadFileHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/repository/save",
					Handler: userRepositorySaveHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/file/list",
					Handler: userFileListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/file/update",
					Handler: userFileUpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/folder/create",
					Handler: userFolderCreateHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/user/file/delete",
					Handler: userFileDeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/refresh/token",
					Handler: userRefreshTokenHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/upload/prepare",
					Handler: userUploadPrepareHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/file/upload/chunk",
					Handler: userFileUploadChunkHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/file/upload/chunk/complete",
					Handler: userFileUploadChunkCompleteHandler(serverCtx),
				},
			}...,
		),
	)
}
