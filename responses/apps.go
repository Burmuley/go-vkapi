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

import "gitlab.com/Burmuley/go-vkapi/objects"

/////////////////////////////////////////////////////////////
// `apps` group of responses
/////////////////////////////////////////////////////////////

// AppsGetScore type represents `apps_getScore_response` API response object
type AppsGetScore int // Score number

// AppsSendRequest type represents `apps_sendRequest_response` API response object
type AppsSendRequest int // Request ID

// AppsGetFriendsList type represents `apps_getFriendsList_response` API response object
type AppsGetFriendsList struct {
	Count int                     `json:"count"` // Total number
	Items []objects.UsersUserFull `json:"items"`
}

// AppsGetLeaderboard type represents `apps_getLeaderboard_response` API response object
type AppsGetLeaderboard struct {
	Count int                       `json:"count"` // Total number
	Items []objects.AppsLeaderboard `json:"items"`
}

// AppsGetLeaderboardExtended type represents `apps_getLeaderboard_extended_response` API response object
type AppsGetLeaderboardExtended struct {
	Count    int                       `json:"count"` // Total number
	Items    []objects.AppsLeaderboard `json:"items"`
	Profiles []objects.UsersUserMin    `json:"profiles"`
}

// AppsGet type represents `apps_get_response` API response object
type AppsGet struct {
	Count int               `json:"count"` // Total number
	Items []objects.AppsApp `json:"items"`
}

// AppsGetScopes type represents `apps_getScopes_response` API response object
type AppsGetScopes struct {
	Count int                 `json:"count"` // Total number
	Items []objects.AppsScope `json:"items"`
}

// AppsGetCatalog type represents `apps_getCatalog_response` API response object
type AppsGetCatalog struct {
	Count int               `json:"count"` // Total number
	Items []objects.AppsApp `json:"items"`
}
