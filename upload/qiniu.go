package upload

import (
	"context"
	"github.com/qiniu/go-sdk/v7/storagev2/credentials"
	"github.com/qiniu/go-sdk/v7/storagev2/uploader"
	"github.com/qiniu/go-sdk/v7/storagev2/uptoken"
	"time"
)

var QiniuCredentials *credentials.Credentials
var Uploader *QiniuUploader

type QiniuUploader struct {
	Mac     *credentials.Credentials
	Policy  uptoken.PutPolicy
	Options *uploader.UploadManagerOptions
}

func NewQiniuUploader(accessKey, secretKey string) *QiniuUploader {
	mac := NewMac(accessKey, secretKey)
	return &QiniuUploader{
		Mac: mac,
	}
}
func (c *QiniuUploader) Upload(localFile, key string, customVars map[string]string, ret interface{}) error {
	uploadManager := uploader.NewUploadManager(c.Options)
	err := uploadManager.UploadFile(context.Background(), localFile, &uploader.ObjectOptions{
		UpToken:    uptoken.NewSigner(c.Policy, c.Mac),
		ObjectName: &key,
		CustomVars: customVars,
	}, &ret)
	return err
}
func (c *QiniuUploader) UploadToken() (string, error) {
	return uptoken.NewSigner(c.Policy, c.Mac).GetUpToken(context.Background())
}
func (c *QiniuUploader) SetUploadManagerOptions(options *uploader.UploadManagerOptions) *QiniuUploader {
	c.Options = options
	return c
}
func (c *QiniuUploader) SetPutPolicy(bucket string, ttl time.Time) *QiniuUploader {
	policy, err := uptoken.NewPutPolicy(bucket, ttl)
	if err != nil {
		panic(err)
	}
	c.Policy = policy
	return c
}
func NewMac(accessKey, secretKey string) *credentials.Credentials {
	return credentials.NewCredentials(accessKey, secretKey)
}

func NewPutPolicy(bucket string, ttl time.Time) (uptoken.PutPolicy, error) {
	return uptoken.NewPutPolicy(bucket, ttl)
}
