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
// Code generator location: https://github.com/Burmuley/go-vkapi-gen                                       //
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

package responses

import (
	"github.com/Burmuley/go-vkapi/objects"
)

/////////////////////////////////////////////////////////////
// `users` group of responses
/////////////////////////////////////////////////////////////

// UsersGetfollowersFields type represents `users_getFollowers_fields_response` API response object
type UsersGetfollowersFields struct {
	Count int                     `json:"count"` // Total number of available results
	Items []objects.UsersUserFull `json:"items"`
}

// UsersGetfollowers type represents `users_getFollowers_response` API response object
type UsersGetfollowers struct {
	Count int   `json:"count"` // Total friends number
	Items []int `json:"items"`
}

// UsersGetsubscriptionsExtended type represents `users_getSubscriptions_extended_response` API response object
type UsersGetsubscriptionsExtended struct {
	Count int                              `json:"count"` // Total number of available results
	Items []objects.UsersSubscriptionsItem `json:"items"`
}

// UsersGetsubscriptions type represents `users_getSubscriptions_response` API response object
type UsersGetsubscriptions struct {
	Groups objects.GroupsGroupsArray `json:"groups"`
	Users  objects.UsersUsersArray   `json:"users"`
}

// UsersGet type represents `users_get_response` API response object
type UsersGet objects.UsersUserXtrCounters

// UsersIsappuser type represents `users_isAppUser_response` API response object
type UsersIsappuser objects.BaseBoolInt // Information whether the user have installed an app

// UsersSearch type represents `users_search_response` API response object
type UsersSearch struct {
	Count int                     `json:"count"` // Total number of available results
	Items []objects.UsersUserFull `json:"items"`
}
