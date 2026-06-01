package schema

import (
	"github.com/copkg/gopkg/errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type SendSmsRequest struct {
	Mobile  string      `json:"mobile"`
	Content string      `json:"content"`
	Cond    interface{} `json:"cond"`
	UID     int
}

func (a SendSmsRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Mobile, validation.Required.Error("不能为空")),
	)
}

type SendSmsResponse struct {
	*errors.Error
	TaskID int64 `json:"task_id"`
}

type SmsTemplate struct {
	TID       int64  `json:"tid"`        // 模板ID
	Title     string `json:"title"`      // 模板标题
	Content   string `json:"content"`    // 模板内容
	CreatedAt int64  `json:"created_at"` // 创建时间
	UpdatedAt int64  `json:"updated_at"` // 更新时间
}

type SmsTemplateListRequest struct {
}

type SmsTemplateListResponse struct {
	Templates []*SmsTemplate `json:"templates"`
	*errors.Error
}

type SmsTemplateRemoveRequest struct {
	TID int64 `json:"tid" binding:"required"`
}

func (a SmsTemplateRemoveRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.TID, validation.Required.Error("模板id不能为空")),
	)
}

type SmsTemplateRemoveResponse struct {
	*errors.Error
}

type SmsTemplateRequest struct {
	Title   string `json:"title"`   // 模板标题
	Content string `json:"content"` // 模板内容
}

func (a SmsTemplateRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Title, validation.Required.Error("标题不能为空")),
		validation.Field(&a.Content, validation.Required.Error("内容不能为空")),
	)
}

type SmsTemplateResponse struct {
	TID int64 `json:"tid,omitempty"`
	*errors.Error
}

type SmsTemplateUpdateRequest struct {
	Tid     int64  `json:"tid" `
	Title   string `json:"title" `   // 模板标题
	Content string `json:"content" ` // 模板内容
}

func (a SmsTemplateUpdateRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Tid, validation.Required.Error("模板id不能为空")),
		validation.Field(&a.Title, validation.Required.Error("标题不能为空")),
		validation.Field(&a.Content, validation.Required.Error("内容不能为空")),
	)
}
