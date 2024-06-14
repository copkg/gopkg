package schema

type SmsLog struct {
	SendID        int64  `json:"send_id"`        // 主键ID
	TaskID        int64  `json:"task_id"`        // 任务ID
	SendCond      string `json:"send_cond"`      // 发送条件
	SendMobile    string `json:"send_mobile"`    // 手机号
	TemplateTitle string `json:"template_title"` // 标题
	SendContent   string `json:"send_content"`   // 内容
	SendReason    string `json:"send_reason"`    // 失败原因
	SendCount     int    `json:"send_count"`     // 发送手机号数量
	SendStatus    string `json:"send_status"`    // 发送状态
	TransactionID string `json:"transaction_id"` // 平台交易号
	SendAt        int64  `json:"send_at"`        // 发送时间
}

type SmsLogListResponse struct {
	Logs  []*SmsLog `json:"logs"`
	Total int       `json:"total"`
	Page  int       `json:"page,omitempty"`
	Size  int       `json:"size,omitempty"`
	*Error
}

type SmsLogRequest struct {
	Page        int64  `json:"page"`
	Size        int64  `json:"size"` // 每页大小
	SendTimeMin string `json:"send_time_min,optional"`
	SendTimeMax string `json:"send_time_max,optional"`
	Status      string `json:"status,optional"`
	Mobile      string `json:"mobile,optional"`
	TaskID      int64  `json:"task_id,optional"`
}
