package upload

import (
	"context"
	"github.com/qiniu/go-sdk/v7/storagev2/credentials"
	"github.com/qiniu/go-sdk/v7/storagev2/uploader"
	"github.com/qiniu/go-sdk/v7/storagev2/uptoken"
	"time"
)

var QiniuCredentials *credentials.Credentials

type QiniuUploader struct {
	Mac        *credentials.Credentials
	Policy     uptoken.PutPolicy
	LocalFile  string
	Options    *uploader.UploadManagerOptions
	Key        string
	CustomVars map[string]string
	Ret        interface{}
}

func (c *QiniuUploader) Upload() error {
	uploadManager := uploader.NewUploadManager(c.Options)
	err := uploadManager.UploadFile(context.Background(), c.LocalFile, &uploader.ObjectOptions{
		UpToken:    uptoken.NewSigner(c.Policy, c.Mac),
		ObjectName: &c.Key,
		CustomVars: c.CustomVars,
	}, &c.Ret)
	return err
}
func NewMac(accessKey, secretKey string) *credentials.Credentials {
	return credentials.NewCredentials(accessKey, secretKey)
}

func NewPutPolicy(bucket string, ttl time.Time) (uptoken.PutPolicy, error) {
	return uptoken.NewPutPolicy(bucket, ttl)
}
