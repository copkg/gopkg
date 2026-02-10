package sms

import (
	"crypto/md5"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"strings"
	"time"
)

var SmsClient *Smser
var client = resty.New()
var smsUrl = "http://115.239.134.217/smsservice/httpservices/capService"

type SIConfig struct {
	Siid      string
	SecretKey string
	User      string
}
type SIRes struct {
	TransactionID string `json:"transactionID"`
	RetCode       string `json:"retCode"`
	RetMsg        string `json:"retMsg"`
}
type Smser struct {
	Siid      string
	SecretKey string
	User      string
}

func NewSmser(c *SIConfig) *Smser {
	if c == nil {
		panic("sms config cannot be nil")
	}
	return &Smser{
		Siid:      c.Siid,
		SecretKey: c.SecretKey,
		User:      c.User,
	}
}

// 格式yyyymmddhhiissSSS
func (s *Smser) timeStamp() string {
	return strings.Join(strings.Split(time.Now().Format("20060102150405.999"), "."), "")
}
func (s *Smser) authenticator(str string) string {
	hash := md5.Sum([]byte(str))
	return base64.StdEncoding.EncodeToString(hash[:])
}
func (s *Smser) Send(content, mobile string) (*SIRes, error) {
	var data = make(map[string]string)
	timeStamp := s.timeStamp()
	data["siid"] = s.Siid
	data["user"] = s.User
	data["streamingNo"] = s.Siid + timeStamp
	data["timeStamp"] = timeStamp
	data["authenticator"] = s.authenticator(timeStamp + timeStamp + data["streamingNo"] + s.SecretKey)

	data["content"] = content
	data["mobile"] = mobile
	data["transactionID"] = timeStamp
	var result *SIRes
	_, err := client.R().
		SetBody(data).
		SetResult(&result).
		Post(smsUrl)
	if err != nil {
		return nil, err
	}
	if result.RetCode != "0000" {
		return result, errors.New(fmt.Sprintf("call smsservice err,retCode:%s,retMsg:%s", result.RetCode, result.RetMsg))
	}
	return result, nil
}
