package helper

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

type AliOss struct {
	ossCli             *oss.Client
	ossEndpoint        string
	ossAccessKeyID     string
	ossAccessKeySecret string
	ossBucket          string
	ossUrl             string
}

func (t AliOss) Create(ossEndpoint string, ossAccessKeyID string, ossAccessKeySecret string, ossBucket string, ossUrl string) *AliOss {
	tt := AliOss{
		ossEndpoint:        ossEndpoint,
		ossAccessKeyID:     ossAccessKeyID,
		ossAccessKeySecret: ossAccessKeySecret,
		ossBucket:          ossBucket,
		ossUrl:             ossUrl,
	}
	tt.ossCli, _ = oss.New(ossEndpoint, ossAccessKeyID, ossAccessKeySecret)
	return &tt
}

// Bucket: 获取bucket存储空间
func (t *AliOss) bucket() *oss.Bucket {
	bucket, err := t.ossCli.Bucket(t.ossBucket)
	if err != nil {
		return nil
	}
	return bucket
}

func (t *AliOss) Upload(filename string, ossPath string) string {
	file, _ := os.Open(filename)
	url := t.ossBucket + "/" + ossPath
	err := t.bucket().PutObject(url, file)
	if err != nil {
		return ""
	}
	return t.ossUrl + "/" + url
}
