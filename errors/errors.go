package errors

import (
	"fmt"
)

type VKErrors interface {
	GetCode() int
	GetDescription() string
}

var (
	ApiErrorUnknown               = ApiError{Code: 1, Message: "Unknown error occurred"}
	ApiErrorDisabled              = ApiError{Code: 2, Message: "Application is disabled. Enable your application or use test mode"}
	ApiErrorMethod                = ApiError{Code: 3, Message: "Unknown method passed"}
	ApiErrorSignature             = ApiError{Code: 4, Message: "Incorrect signature"}
	ApiErrorAuth                  = ApiError{Code: 5, Message: "User authorization failed"}
	ApiErrorTooMany               = ApiError{Code: 6, Message: "Too many requests per second"}
	ApiErrorPermission            = ApiError{Code: 7, Message: "Permission to perform this action is denied"}
	ApiErrorRequest               = ApiError{Code: 8, Message: "Invalid request"}
	ApiErrorFlood                 = ApiError{Code: 9, Message: "Flood control"}
	ApiErrorServer                = ApiError{Code: 10, Message: "Internal server error"}
	ApiErrorEnabledInTest         = ApiError{Code: 11, Message: "In test mode application should be disabled or user should be authorized"}
	ApiErrorCaptcha               = ApiError{Code: 14, Message: "Captcha needed"}
	ApiErrorAccess                = ApiError{Code: 15, Message: "Access denied"}
	ApiErrorAuthHttps             = ApiError{Code: 16, Message: "HTTP authorization failed"}
	ApiErrorAuthValidation        = ApiError{Code: 17, Message: "Validation required"}
	ApiErrorUserDeleted           = ApiError{Code: 18, Message: "User was deleted or banned"}
	ApiErrorMethodPermission      = ApiError{Code: 20, Message: "Permission to perform this action is denied for non-standalone applications"}
	ApiErrorMethodAds             = ApiError{Code: 21, Message: "Permission to perform this action is allowed only for standalone and OpenAPI applications"}
	ApiErrorMethodDisabled        = ApiError{Code: 23, Message: "This method was disabled"}
	ApiErrorNeedConfirmation      = ApiError{Code: 24, Message: "Confirmation required"}
	ApiErrorNeedTokenConfirmation = ApiError{Code: 25, Message: "Token confirmation required"}
	ApiErrorGroupAuth             = ApiError{Code: 27, Message: "Group authorization failed"}
	ApiErrorAppAuth               = ApiError{Code: 28, Message: "Application authorization failed"}
	ApiErrorRateLimit             = ApiError{Code: 29, Message: "Rate limit reached"}
	ApiErrorPrivateProfile        = ApiError{Code: 30, Message: "This profile is private"}
	ApiErrorParam                 = ApiError{Code: 100, Message: "One of the parameters specified was missing or invalid"}
	ApiErrorParamApiId            = ApiError{Code: 101, Message: "Invalid application API ID"}
	ApiErrorParamUserId           = ApiError{Code: 113, Message: "Invalid user id"}
	ApiErrorParamTimestamp        = ApiError{Code: 150, Message: "Invalid timestamp"}
	ApiErrorAccessAlbum           = ApiError{Code: 200, Message: "Access denied"}
	ApiErrorAccessAudio           = ApiError{Code: 201, Message: "Access denied"}
	ApiErrorAccessGroup           = ApiError{Code: 203, Message: "Access to group denied"}
	ApiErrorAlbumFull             = ApiError{Code: 300, Message: "This album is full"}
	ApiErrorVotesPermission       = ApiError{Code: 500, Message: "Permission denied. You must enable votes processing in application settings"}
	ApiErrorAdsPermission         = ApiError{Code: 600, Message: "Permission denied. You have no access to operations specified with given object(s)"}
	ApiErrorAdsSpecific           = ApiError{Code: 603, Message: "Some ads error occurred"}
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
