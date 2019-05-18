package errors

import "fmt"

const (
	VK_API_ERROR_UNKNOWN                 int = 1
	VK_API_ERROR_DISABLED                int = 2
	VK_API_ERROR_METHOD                  int = 3
	VK_API_ERROR_SIGNATURE               int = 4
	VK_API_ERROR_AUTH                    int = 5
	VK_API_ERROR_TOO_MANY                int = 6
	VK_API_ERROR_PERMISSION              int = 7
	VK_API_ERROR_REQUEST                 int = 8
	VK_API_ERROR_FLOOD                   int = 9
	VK_API_ERROR_SERVER                  int = 10
	VK_API_ERROR_ENABLED_IN_TEST         int = 11
	VK_API_ERROR_CAPTCHA                 int = 14
	VK_API_ERROR_ACCESS                  int = 15
	VK_API_ERROR_AUTH_HTTPS              int = 16
	VK_API_ERROR_AUTH_VALIDATION         int = 17
	VK_API_ERROR_USER_DELETED            int = 18
	VK_API_ERROR_METHOD_PERMISSION       int = 20
	VK_API_ERROR_METHOD_ADS              int = 21
	VK_API_ERROR_METHOD_DISABLED         int = 23
	VK_API_ERROR_NEED_CONFIRMATION       int = 24
	VK_API_ERROR_NEED_TOKEN_CONFIRMATION int = 25
	VK_API_ERROR_GROUP_AUTH              int = 27
	VK_API_ERROR_APP_AUTH                int = 28
	VK_API_ERROR_RATE_LIMIT              int = 29
	VK_API_ERROR_PRIVATE_PROFILE         int = 30
	VK_API_ERROR_PARAM                   int = 100
	VK_API_ERROR_PARAM_VK_API_ID         int = 101
	VK_API_ERROR_PARAM_USER_ID           int = 113
	VK_API_ERROR_PARAM_TIMESTAMP         int = 150
	VK_API_ERROR_ACCESS_ALBUM            int = 200
	VK_API_ERROR_ACCESS_AUDIO            int = 201
	VK_API_ERROR_ACCESS_GROUP            int = 203
	VK_API_ERROR_ALBUM_FULL              int = 300
	VK_API_ERROR_VOTES_PERMISSION        int = 500
	VK_API_ERROR_ADS_PERMISSION          int = 600
	VK_API_ERROR_ADS_SPECIFIC            int = 603
)

type ApiError struct {
	Code    int    `json:"error_code"`
	Message string `json:"error_msg"`
}

func (e *ApiError) GetCode() int {
	return e.Code
}

func (e *ApiError) GetDescription() string {
	return e.Message
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("API ERROR! Code: %d, Message: %s", e.Code, e.Message)
}
