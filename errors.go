package go_vkapi

import "fmt"

type ApiError struct {
	Code    int    `json:"error_code"`
	Message string `json:"error_msg"`
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("API Error! Code: %d, Message: %s", e.Code, e.Message)
}
