package commonresp

import "encoding/json"

type ResponseMessage struct {
	Status       string      `json:"status"`
	ResponseCode int      `json:"response_code"`
	Description  string      `json:"description"`
	Data         interface{} `json:"data"`
}

type ErrorMessage struct {
	HttpCode int `json:"http_code"`
	ErrorDescription
}

type ErrorDescription struct {
	Status       string `json:"status"`
	Message      string `json:"message"`
}

func (e ErrorMessage) ToJson() string {
	b, err := json.Marshal(e)
	if err != nil {
		return ""
	}
	return string(b)
}

func NewResponseMessage(httpCode int, description string, data interface{}) ResponseMessage {
	return ResponseMessage{
		"Success", httpCode, description, data,
	}
}
func NewErrorMessage(httpCode int, errorMessage string) ErrorMessage {
	return ErrorMessage{
		httpCode, ErrorDescription{
			Status:       "Error",
			Message: errorMessage,
		},
	}
}
