package models

//HTTPErrorMessage |
type HTTPErrorMessage struct {
	Message string `json:"message"`
}

// NewHTTPErrorMessage |
func NewHTTPErrorMessage(m string) *HTTPErrorMessage {
	return &HTTPErrorMessage{Message: m}
}
