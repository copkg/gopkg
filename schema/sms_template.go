package schema

type SendSmsRequest struct {
	TID     int64       `json:"tid"`
	Mobiles []string    `json:"mobiles"`
	Cond    interface{} `json:"cond"`
}

type SendSmsResponse struct {
	*Error
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
	*Error
}

type SmsTemplateRemoveRequest struct {
	TID int64 `json:"tid" binding:"required" error:"tid不能为空"`
}

type SmsTemplateRemoveResponse struct {
	*Error
}

type SmsTemplateRequest struct {
	Title   string `json:"title" binding:"required" error:"title不能为空"`     // 模板标题
	Content string `json:"content" binding:"required" error:"content不能为空"` // 模板内容
}

type SmsTemplateResponse struct {
	TID int64 `json:"tid,omitempty"`
	*Error
}

type SmsTemplateUpdateRequest struct {
	Tid     int64  `json:"tid" binding:"required" error:"tid不能为空"`
	Title   string `json:"title" binding:"required" error:"title不能为空"`     // 模板标题
	Content string `json:"content" binding:"required" error:"content不能为空"` // 模板内容
}
