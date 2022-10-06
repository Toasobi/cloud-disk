package test

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/tencentyun/cos-go-sdk-v5"
)

// func TestFileUpload(t *testing.T) {

// 	// 存储桶名称，由bucketname-appid 组成，appid必须填入，可以在COS控制台查看存储桶名称。 https://console.cloud.tencent.com/cos5/bucket
// 	// 替换为用户的 region，存储桶region可以在COS控制台“存储桶概览”查看 https://console.cloud.tencent.com/ ，关于地域的详情见 https://cloud.tencent.com/document/product/436/6224 。
// 	u, _ := url.Parse("https://toasobi-1313075089.cos.ap-guangzhou.myqcloud.com")
// 	b := &cos.BaseURL{BucketURL: u}
// 	client := cos.NewClient(b, &http.Client{
// 		Transport: &cos.AuthorizationTransport{
// 			// 通过环境变量获取密钥
// 			// 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
// 			SecretID: os.Getenv("AKIDd7OTLyObbWO5e1B8wc8SbZKsHVYO5bw4"),
// 			// 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
// 			SecretKey: os.Getenv("CNPQgv1xzykKhsqb9XlmaJrBIF5jEZgZ"),
// 		},
// 	})

// 	key := "cloud-disk/exampleobject.jpg"

// 	_, _, err := client.Object.Upload(
// 		context.Background(), key, "./img/testingJPG.jpg", nil,
// 	)
// 	if err != nil {
// 		panic(err)
// 	}
// }

func TestPutUpload(t *testing.T) {
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

	key := "cloud-disk/exampleobject.jpg"

	f, err := os.ReadFile("./img/testingJPG.jpg")

	_, err = client.Object.Put(
		context.Background(), key, bytes.NewReader(f), nil,
	)
	if err != nil {
		panic(err)
	}
}

//分片上传初始化
func TestInitUploader(t *testing.T) {
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
	key := "cloud-disk/test.mp4"
	// 可选opt,如果不是必要操作，建议上传文件时不要给单个文件设置权限，避免达到限制。若不设置默认继承桶的权限。
	v, _, err := client.Object.InitiateMultipartUpload(context.Background(), key, nil)
	if err != nil {
		t.Fatal(err)
	}
	UploadID := v.UploadID //16649622071a149cec5e5babf4f45c25e591b9ba1e78181d2469838440be7a5811d305fa0c
	fmt.Println(UploadID)

}

//分片上传

func TestUpload(t *testing.T) {
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
	key := "cloud-disk/test.mp4"
	UploadID := "16649637087495f4a31ab3679f3f22b218ca7a064a32628efd22aaee0b03d3f6826e7d9f60"
	f := strings.NewReader("0.chunk")
	//

	// opt可选
	resp, err := client.Object.UploadPart(
		context.Background(), key, UploadID, 1, f, nil,
	)
	if err != nil {
		t.Fatal(err)
	}
	PartETag := resp.Header.Get("ETag")
	fmt.Println(PartETag) //md5

}

//分片上传完成
func TestPartUploadComplete(t *testing.T) {
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
	key := "cloud-disk/test.mp4"
	UploadID := "16649637087495f4a31ab3679f3f22b218ca7a064a32628efd22aaee0b03d3f6826e7d9f60"
	PartETag := "1"
	opt := &cos.CompleteMultipartUploadOptions{}
	opt.Parts = append(opt.Parts, cos.Object{
		PartNumber: 1, ETag: PartETag},
	)
	_, _, err := client.Object.CompleteMultipartUpload(
		context.Background(), key, UploadID, opt,
	)
	if err != nil {
		t.Fatal(err)
	}

}
