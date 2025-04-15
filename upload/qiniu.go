package upload

import (
	"github.com/qiniu/go-sdk/v7/storagev2/credentials"
)

var QiniuCredentials *credentials.Credentials

type QiniuConf struct {
	AccessKey string
	SecretKey string
}

func New(c *QiniuConf) *credentials.Credentials {
	return credentials.NewCredentials(c.AccessKey, c.SecretKey)
}
