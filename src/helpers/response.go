package helpers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  string      `json:"status"`
	IsError bool        `json:"isError"`
	Data    interface{} `json:"data,omitempty"`
	Message interface{} `json:"message,omitempty"`
}

func (res *Response) Send(w http.ResponseWriter) {
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		w.Write([]byte("Error When Encode Response"))
	}
}

func NewResponse(data interface{}, code int, isError bool, message string) *Response {

	if isError {
		return &Response{
			Status:  getStatus(code),
			IsError: isError,
			Message: message,
			Data:    data,
		}

	}
	return &Response{
		Status:  getStatus(code),
		IsError: isError,
		Message: message,
		Data:    data,
	}
}

func getStatus(status int) string {
	var description string
	switch status {
	case 200:
		description = "OK"
	case 201:
		description = "Created"
	case 400:
		description = "Bad Request"
	case 401:
		description = "Unauthorized"
	case 404:
		description = "Not Found"
	case 500:
		description = "Internal Server Error"
	case 501:
		description = "Bad Gateway"
	case 304:
		description = "Not Modified"
	default:
		description = ""
	}

	return description
}
