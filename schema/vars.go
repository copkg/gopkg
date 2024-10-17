package schema

type Comm struct {
	Code int    `json:"code"`
	Msg  string `json:"msg,omitempty"`
}

func (e Comm) GetMessage() string {
	return e.Msg
}
