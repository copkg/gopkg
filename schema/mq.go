package schema

type MqMsgBody struct {
	TypeTask string   `json:"type_task"`
	Data     *SmsBody `json:"data"`
}
type SmsBody struct {
	Mobile  string `json:"mobile"`
	Content string `json:"content"`
	SendID  int64  `json:"send_id"`
}
