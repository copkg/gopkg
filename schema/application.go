package schema

type ApplicationRequest struct {
	AppID   string `json:"app_id"`
	AppName string `json:"app_name"`
	AID     int    `json:"aid"`
}

type Application struct {
	ID          int    `json:"id"`
	AppID       string `json:"app_id"`
	AppSecret   string `json:"app_secret"`
	AgentID     int    `json:"agent_id,omitempty"`
	AppType     string `json:"app_type,omitempty"`
	AppName     string `json:"app_name,omitempty"`
	RedirectURL string `json:"redirect_url,omitempty"`
	Director    string `json:"director,omitempty"`
	Status      int    `json:"status,omitempty"`
}
type ApplicationResponse struct {
	Applicatin *Application `json:"applicatin"`
	*Comm
}

type ApplicationListRequest struct {
}

type ApplicationListResponse struct {
	Applications []*Application `json:"applications"`
	*Comm
}
