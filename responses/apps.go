/*
Copyright 2019 Konstantin Vasilev (burmuley@gmail.com)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

/////////////////////////////////////////////////////////////////////////////////////////////////////////////
// WARNING! AUTOMATICALLY GENERATED CONTENT! DON'T CHANGE IT MANUALLY!                                     //
// Source schema can be found at https://github.com/VKCOM/vk-api-schema/blob/master/responses.json         //
// Code generator location: https://gitlab.com/Burmuley/go-vkapi-gen                                       //
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

package responses

import (
	"gitlab.com/Burmuley/go-vkapi/objects"
)

/////////////////////////////////////////////////////////////
// `apps` group of responses
/////////////////////////////////////////////////////////////

// AppsGetcatalog type represents `apps_getCatalog_response` API response object
type AppsGetcatalog struct {
	Count int                `json:"count"` // Total number
	Items []*objects.AppsApp `json:"items"`
}

// AppsGetfriendslist type represents `apps_getFriendsList_response` API response object
type AppsGetfriendslist struct {
	Count int                      `json:"count"` // Total number
	Items []*objects.UsersUserFull `json:"items"`
}

// AppsGetleaderboardExtended type represents `apps_getLeaderboard_extended_response` API response object
type AppsGetleaderboardExtended struct {
	Count    int                        `json:"count"` // Total number
	Items    []*objects.AppsLeaderboard `json:"items"`
	Profiles []*objects.UsersUserMin    `json:"profiles"`
}

// AppsGetleaderboard type represents `apps_getLeaderboard_response` API response object
type AppsGetleaderboard struct {
	Count int                        `json:"count"` // Total number
	Items []*objects.AppsLeaderboard `json:"items"`
}

// AppsGetsc type represents `apps_getScopes_response` API response object
type AppsGetsc struct {
	Count int                  `json:"count"` // Total number
	Items []*objects.AppsScope `json:"items"`
}

// AppsGetscor type represents `apps_getScore_response` API response object
type AppsGetscor int // Score number

// AppsGet type represents `apps_get_response` API response object
type AppsGet struct {
	Count int                `json:"count"` // Total number
	Items []*objects.AppsApp `json:"items"`
}

// AppsSendrequest type represents `apps_sendRequest_response` API response object
type AppsSendrequest int // Request ID
