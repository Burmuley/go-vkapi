package go_vkapi

import "encoding/json"

type RequestParams struct {
	Key   string `json:key`
	Value string `json:value`
}

type ApiResponse struct {
	Error         *ApiError       `json:error`
	RequestParams []RequestParams `json:request_params`
	Response      json.RawMessage `json:response`
}
