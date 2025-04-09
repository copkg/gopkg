package upload

import (
	"context"
	"github.com/qiniu/go-sdk/v7/storagev2/credentials"
	"github.com/qiniu/go-sdk/v7/storagev2/uptoken"
	"time"
)

var QiniuCredentials *credentials.Credentials

type QiniuConf struct {
	AccessKey string
	SecretKey string
}

func New(c *QiniuConf) *credentials.Credentials {
	return credentials.NewCredentials(c.AccessKey, c.SecretKey)
}

func PutPolicy(bucket, keyToOverwrite, returnBody string, ttl time.Duration) (uptoken.PutPolicy, error) {
	putPolicy, err := uptoken.NewPutPolicy(bucket, time.Now().Add(ttl))
	if keyToOverwrite != "" {
		putPolicy, err = uptoken.NewPutPolicyWithKey(bucket, keyToOverwrite, time.Now().Add(ttl))
	}
	if returnBody != "" {
		putPolicy.SetReturnBody(returnBody)
	}
	return putPolicy, err
}
func UploadToken(putPolicy uptoken.PutPolicy, mac *credentials.Credentials) (string, error) {
	return uptoken.NewSigner(putPolicy, mac).GetUpToken(context.Background())
}
