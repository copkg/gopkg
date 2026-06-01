package schema

import (
	"github.com/copkg/gopkg/errors"
	"time"
)

type UploadTokenRequest struct {
	Key        string        `json:"key"`
	Bucket     string        `json:"bucket"`
	ReturnBody string        `json:"returnbody"`
	Ttl        time.Duration `json:"ttl"`
}
type UploadTokenResponse struct {
	Token       string `json:"token"`
	DownloadUrl string `json:"downloadurl"`
	*errors.Error
}
