package models

type RestResponse struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Result  interface{} `json:"result,omitempty"` // map[string]interface
	Error   *RestError  `json:"error,omitempty"`
}

type RestError struct {
	ID         string                 `json:"id,omitempty"`
	Title      string                 `json:"title,omitempty"`
	Message    string                 `json:"message,omitempty"`
	StatusCode int                    `json:"status_code,omitempty"`
	EventCode  string                 `json:"event_code,omitempty"`
	Meta       map[string]interface{} `json:"meta,omitempty"`
	Error      error                  `json:"error,omitempty"`
	Alert      bool                   `json:"alert,omitempty"`
}
