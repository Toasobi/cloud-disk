package helper

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/tencentyun/cos-go-sdk-v5"
)

func CosUpload(r *http.Request) (string, error) {
	// 存储桶名称，由bucketname-appid 组成，appid必须填入，可以在COS控制台查看存储桶名称。 https://console.cloud.tencent.com/cos5/bucket
	// 替换为用户的 region，存储桶region可以在COS控制台“存储桶概览”查看 https://console.cloud.tencent.com/ ，关于地域的详情见 https://cloud.tencent.com/document/product/436/6224 。
	u, _ := url.Parse("https://toasobi-1313075089.cos.ap-guangzhou.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 通过环境变量获取密钥
			// 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretID: os.Getenv("AKIDd7OTLyObbWO5e1B8wc8SbZKsHVYO5bw4"),
			// 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretKey: os.Getenv("CNPQgv1xzykKhsqb9XlmaJrBIF5jEZgZ"),
		},
	})

	file, fileHeader, err := r.FormFile("file")
	key := "cloud-disk/" + GenerateUUID() + path.Ext(fileHeader.Filename)

	// f, err := os.ReadFile("./img/testingJPG.jpg")

	_, err = client.Object.Put(
		context.Background(), key, file, nil,
	)
	if err != nil {
		panic(err)
	}
	return "https://toasobi-1313075089.cos.ap-guangzhou.myqcloud.com" + "/" + key, nil
}

func CosInitUploader(ext string) (string, string, error) {
	u, _ := url.Parse("https://toasobi-1313075089.cos.ap-guangzhou.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 通过环境变量获取密钥
			// 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretID: os.Getenv("AKIDd7OTLyObbWO5e1B8wc8SbZKsHVYO5bw4"),
			// 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretKey: os.Getenv("CNPQgv1xzykKhsqb9XlmaJrBIF5jEZgZ"),
		},
	})
	key := "cloud-disk/" + GenerateUUID() + ext
	// 可选opt,如果不是必要操作，建议上传文件时不要给单个文件设置权限，避免达到限制。若不设置默认继承桶的权限。
	v, _, err := client.Object.InitiateMultipartUpload(context.Background(), key, nil)
	if err != nil {
		return "", "", err
	}
	// UploadID := v.UploadID //16649622071a149cec5e5babf4f45c25e591b9ba1e78181d2469838440be7a5811d305fa0c
	return key, v.UploadID, nil

}

//分片上传

func CosPartUpload(r *http.Request) (string, error) {
	u, _ := url.Parse("https://toasobi-1313075089.cos.ap-guangzhou.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 通过环境变量获取密钥
			// 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretID: os.Getenv("AKIDd7OTLyObbWO5e1B8wc8SbZKsHVYO5bw4"),
			// 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretKey: os.Getenv("CNPQgv1xzykKhsqb9XlmaJrBIF5jEZgZ"),
		},
	})
	key := r.PostForm.Get("key")
	UploadID := r.PostForm.Get("upload_id")
	partNumber, _ := strconv.Atoi(r.PostForm.Get("part_number"))
	f, _, err := r.FormFile("file")
	if err != nil {
		fmt.Println("--------")
		fmt.Println("--------")
		fmt.Println("--------")
		fmt.Println("--------")
		fmt.Println("--------")
		fmt.Println("--------")
		fmt.Println("--------")
		fmt.Println("--------")
		return "", err
	}

	buf := bytes.NewBuffer(nil)
	io.Copy(buf, f)

	// opt可选
	resp, err := client.Object.UploadPart(
		context.Background(), key, UploadID, partNumber, bytes.NewReader(buf.Bytes()), nil,
	)
	if err != nil {
		return "", err
	}
	PartETag := resp.Header.Get("ETag")
	fmt.Println(PartETag) //md5
	return strings.Trim(resp.Header.Get("ETag"), "\""), nil

}

//分片上传完成
func PartUploadComplete(key, uploadId string, cs []cos.Object) error {
	u, _ := url.Parse("https://toasobi-1313075089.cos.ap-guangzhou.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 通过环境变量获取密钥
			// 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretID: os.Getenv("AKIDd7OTLyObbWO5e1B8wc8SbZKsHVYO5bw4"),
			// 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretKey: os.Getenv("CNPQgv1xzykKhsqb9XlmaJrBIF5jEZgZ"),
		},
	})

	opt := &cos.CompleteMultipartUploadOptions{}
	opt.Parts = append(opt.Parts, cs...,
	)
	_, _, err := client.Object.CompleteMultipartUpload(
		context.Background(), key, uploadId, opt,
	)
	if err != nil {
		return err
	}
	return nil
}
