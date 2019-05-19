package errors

import (
	"fmt"
)

type VKErrors interface {
	GetCode() int
	GetDescription() string
}

var (
	API_ERROR_UNKNOWN                 = ApiError{Code: 1, Message: "Unknown error occurred"}
	API_ERROR_DISABLED                = ApiError{Code: 2, Message: "Application is disabled. Enable your application or use test mode"}
	API_ERROR_METHOD                  = ApiError{Code: 3, Message: "Unknown method passed"}
	API_ERROR_SIGNATURE               = ApiError{Code: 4, Message: "Incorrect signature"}
	API_ERROR_AUTH                    = ApiError{Code: 5, Message: "User authorization failed"}
	API_ERROR_TOO_MANY                = ApiError{Code: 6, Message: "Too many requests per second"}
	API_ERROR_PERMISSION              = ApiError{Code: 7, Message: "Permission to perform this action is denied"}
	API_ERROR_REQUEST                 = ApiError{Code: 8, Message: "Invalid request"}
	API_ERROR_FLOOD                   = ApiError{Code: 9, Message: "Flood control"}
	API_ERROR_SERVER                  = ApiError{Code: 10, Message: "Internal server error"}
	API_ERROR_ENABLED_IN_TEST         = ApiError{Code: 11, Message: "In test mode application should be disabled or user should be authorized"}
	API_ERROR_CAPTCHA                 = ApiError{Code: 14, Message: "Captcha needed"}
	API_ERROR_ACCESS                  = ApiError{Code: 15, Message: "Access denied"}
	API_ERROR_AUTH_HTTPS              = ApiError{Code: 16, Message: "HTTP authorization failed"}
	API_ERROR_AUTH_VALIDATION         = ApiError{Code: 17, Message: "Validation required"}
	API_ERROR_USER_DELETED            = ApiError{Code: 18, Message: "User was deleted or banned"}
	API_ERROR_METHOD_PERMISSION       = ApiError{Code: 20, Message: "Permission to perform this action is denied for non-standalone applications"}
	API_ERROR_METHOD_ADS              = ApiError{Code: 21, Message: "Permission to perform this action is allowed only for standalone and OpenAPI applications"}
	API_ERROR_METHOD_DISABLED         = ApiError{Code: 23, Message: "This method was disabled"}
	API_ERROR_NEED_CONFIRMATION       = ApiError{Code: 24, Message: "Confirmation required"}
	API_ERROR_NEED_TOKEN_CONFIRMATION = ApiError{Code: 25, Message: "Token confirmation required"}
	API_ERROR_GROUP_AUTH              = ApiError{Code: 27, Message: "Group authorization failed"}
	API_ERROR_APP_AUTH                = ApiError{Code: 28, Message: "Application authorization failed"}
	API_ERROR_RATE_LIMIT              = ApiError{Code: 29, Message: "Rate limit reached"}
	API_ERROR_PRIVATE_PROFILE         = ApiError{Code: 30, Message: "This profile is private"}
	API_ERROR_PARAM                   = ApiError{Code: 100, Message: "One of the parameters specified was missing or invalid"}
	API_ERROR_PARAM_API_ID            = ApiError{Code: 101, Message: "Invalid application API ID"}
	API_ERROR_PARAM_USER_ID           = ApiError{Code: 113, Message: "Invalid user id"}
	API_ERROR_PARAM_TIMESTAMP         = ApiError{Code: 150, Message: "Invalid timestamp"}
	API_ERROR_ACCESS_ALBUM            = ApiError{Code: 200, Message: "Access denied"}
	API_ERROR_ACCESS_AUDIO            = ApiError{Code: 201, Message: "Access denied"}
	API_ERROR_ACCESS_GROUP            = ApiError{Code: 203, Message: "Access to group denied"}
	API_ERROR_ALBUM_FULL              = ApiError{Code: 300, Message: "This album is full"}
	API_ERROR_VOTES_PERMISSION        = ApiError{Code: 500, Message: "Permission denied. You must enable votes processing in application settings"}
	API_ERROR_ADS_PERMISSION          = ApiError{Code: 600, Message: "Permission denied. You have no access to operations specified with given object(s)"}
	API_ERROR_ADS_SPECIFIC            = ApiError{Code: 603, Message: "Some ads error occurred"}
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
