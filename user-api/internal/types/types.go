// Code generated by goctl. DO NOT EDIT.
package types

type UserLoginReq struct {
	UserName string `form:"username"`
	Password string `form:"password"`
}

type UserLoginResp struct {
	Token        string `form:"token"`
	RefreshToken string `form:"refresh_Token"`
}

type UserRegisterReq struct {
	Name     string `form:"name"`
	Password string `form:"password"`
	Email    string `form:"email"`
	Code     string `form:"code"`
}

type UserRegisterResp struct {
	Status string `form:"status"`
}

type UserSendCodeReq struct {
	Email string `form:"email"`
}

type UserSendCodeResp struct {
	Status string `form:"status"`
}

type FileUploadReq struct {
	Hash string `form:"hash,optional"`
	Name string `form:"name,optional"`
	Ext  string `form:"ext,optional"`
	Size int64  `form:"size,optional"`
	Path string `form:"path,optional"`
}

type FileUploadResp struct {
	Identity string `form:"identity"`
	Ext      string `form:"ext"`
	Name     string `form:"name"`
}

type UserRepositorySaveReq struct {
	ParentId           int    `form:"parentId"`
	RepositoryIdentity string `form:"repositoryIdentity"`
	Ext                string `form:"ext"`
	Name               string `form:"name"`
}

type UserRepositorySaveResp struct {
	Status string `form:"status"`
}

type UserFileListReq struct {
	Id   int64 `form:"id,optional"`
	Page int64 `form:"page,optional"`
	Size int64 `form:"size,optional"`
}

type UserFileListResp struct {
	List  []*UserFile `form:"list"`
	Count int64       `form:"count"`
}

type UserFile struct {
	Id                 int64  `form:"id"`
	Identity           string `form:"identity"`
	RepositoryIdentity string `form:"repository_identity"`
	Ext                string `form:"ext"` // 文件或文件夹类型
	Name               string `form:"name"`
	Path               string `form:"path"`
	Size               string `form:"size"`
}

type UserFileUpdateReq struct {
	Identity string `form:"identity"`
	Name     string `form:"name"`
}

type UserFileUpdateResp struct {
	Name string `form:"name"`
}

type UserFolderCreateReq struct {
	Name     string `form:"name"`
	ParentId int    `form:"parent_id"`
}

type UserFolderCreateResp struct {
	Status string `form:"status"`
}

type UserFileDeleteReq struct {
	Identity string `form:"identity"`
}

type UserFileDeleteResp struct {
	Status string `form:"status"`
}

type UserRefreshTokenReq struct {
}

type UserRefreshTokenResp struct {
	Token        string `form:"token"`
	RefreshToken string `form:"refresh_Token"`
}

type UserUploadPrepareReq struct {
	Md5  string `form:"md5"`
	Name string `form:"name"`
	Ext  string `form:"ext"`
}

type UserUploadPrepareResp struct {
	Identity string `form:"identity"`
	UploadId string `form:"upload_id"`
	Key      string `form:"key"`
}

type UserFileUploadChunkReq struct {
	UploadId   string `form:"upload_id"`
	Key        string `form:"key"`
	PartNumber string `form:"part_number"`
}

type UserFileUploadChunkResp struct {
	Etag string `form:"etag"`
}

type UserFileUploadChunkCompleteReq struct {
	Key        string      `json:"key"`
	UploadId   string      `json:"upload_id"`
	CosObjects []CosObject `json:"cos_objects"`
}

type UserFileUploadChunkCompleteResp struct {
	Status string `json:"status"`
}

type CosObject struct {
	PartNumber int64  `json:"part_number"`
	Etag       string `json:"etag"`
}
