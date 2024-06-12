package asynq

import (
	"github.com/copkg/gopkg/json"
	"github.com/hibiken/asynq"
	"time"
)

const (
	TypeSmsSend = "sms:send"
)

type SmsSendPayload struct {
	Mobile  string `json:"mobile"`
	Content string `json:"content"`
	SendID  int64  `json:"send_id"`
}

func NewEmailDeliveryTask(mobile, content string, sendid int64) (*asynq.Task, error) {
	payload, err := json.Marshal(SmsSendPayload{Mobile: mobile, Content: content, SendID: sendid})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeSmsSend, payload, asynq.MaxRetry(5), asynq.Timeout(20*time.Minute)), nil
}
