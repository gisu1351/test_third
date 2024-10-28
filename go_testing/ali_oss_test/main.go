package main

import (
	"fmt"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type defaultCredentials struct {
	config *oss.Config
}

func (defCre *defaultCredentials) GetAccessKeyID() string {
	return defCre.config.AccessKeyID
}

func (defCre *defaultCredentials) GetAccessKeySecret() string {
	return defCre.config.AccessKeySecret
}

func (defCre *defaultCredentials) GetSecurityToken() string {
	return defCre.config.SecurityToken
}

type defaultCredentialsProvider struct {
	config *oss.Config
}

func (defBuild *defaultCredentialsProvider) GetCredentials() oss.Credentials {
	return &defaultCredentials{config: defBuild.config}
}
func NewDefaultCredentialsProvider(accessID, accessKey, token string) (defaultCredentialsProvider, error) {
	var provider defaultCredentialsProvider
	if accessID == "" {
		return provider, fmt.Errorf("access key id is empty!")
	}
	if accessKey == "" {
		return provider, fmt.Errorf("access key secret is empty!")
	}
	config := &oss.Config{
		AccessKeyID:     accessID,
		AccessKeySecret: accessKey,
		SecurityToken:   token,
	}
	return defaultCredentialsProvider{
		config,
	}, nil
}

func HandleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}

func main() {
	accessKeyID := ""
	accessKeySecret := ""
	provider, err := NewDefaultCredentialsProvider(accessKeyID, accessKeySecret, "")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	client, err := oss.New("http://oss-cn-hangzhou.aliyuncs.com", "", "", oss.SetCredentialsProvider(&provider))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	fmt.Printf("client:%#v\n", client)
	marker := ""
	// for {
	// 	lsRes, err := client.ListBuckets(oss.Marker(marker))
	// 	if err != nil {
	// 		fmt.Println("Error:", err)
	// 		os.Exit(-1)
	// 	}

	// 	// 默认情况下一次返回100条记录。
	// 	for _, bucket := range lsRes.Buckets {
	// 		fmt.Println("Bucket: ", bucket.Name)
	// 	}

	// 	if lsRes.IsTruncated {
	// 		marker = lsRes.NextMarker
	// 	} else {
	// 		break
	// 	}
	// }

	bucketName := "haiwell-firmware-update-cache-test"
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		HandleError(err)
	}

	marker = ""
	for {
		lsRes, err := bucket.ListObjects(oss.Marker(marker))
		if err != nil {
			HandleError(err)
		}
		// 打印列举结果。默认情况下，一次返回100条记录。
		for _, object := range lsRes.Objects {
			fmt.Println("Object Name: ", object.Key)
		}
		if lsRes.IsTruncated {
			marker = lsRes.NextMarker
		} else {
			break
		}
	}
}
