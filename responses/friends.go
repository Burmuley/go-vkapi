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
// `friends` group of responses
/////////////////////////////////////////////////////////////

// FriendsAddlist type represents `friends_addList_response` API response object
type FriendsAddlist struct {
	ListId int `json:"list_id"` // List ID
}

// FriendsAdd type represents `friends_add_response` API response object
type FriendsAdd int // Friend request status

// FriendsArefriends type represents `friends_areFriends_response` API response object
type FriendsArefriends objects.FriendsFriendStatus

// FriendsDelete type represents `friends_delete_response` API response object
type FriendsDelete struct {
	FriendDeleted     int                    `json:"friend_deleted"`      // Returns 1 if friend has been deleted
	InRequestDeleted  int                    `json:"in_request_deleted"`  // Returns 1 if incoming request has been declined
	OutRequestDeleted int                    `json:"out_request_deleted"` // Returns 1 if out request has been canceled
	Success           objects.BaseOkResponse `json:"success"`
	SuggestionDeleted int                    `json:"suggestion_deleted"` // Returns 1 if suggestion has been declined
}

// FriendsGetappusers type represents `friends_getAppUsers_response` API response object
type FriendsGetappusers int

// FriendsGetbyphones type represents `friends_getByPhones_response` API response object
type FriendsGetbyphones objects.FriendsUserXtrPhone

// FriendsGetlists type represents `friends_getLists_response` API response object
type FriendsGetlists struct {
	Count int                          `json:"count"` // Total communities number
	Items []objects.FriendsFriendsList `json:"items"`
}

// FriendsGetmutual type represents `friends_getMutual_response` API response object
type FriendsGetmutual int

// FriendsGetmutualTargetUids type represents `friends_getMutual_target_uids_response` API response object
type FriendsGetmutualTargetUids objects.FriendsMutualFriend

// FriendsGetonlineOnlineMobile type represents `friends_getOnline_online_mobile_response` API response object
type FriendsGetonlineOnlineMobile struct {
	Online       []int `json:"online"`
	OnlineMobile []int `json:"online_mobile"`
}

// FriendsGetonline type represents `friends_getOnline_response` API response object
type FriendsGetonline int

// FriendsGetrecent type represents `friends_getRecent_response` API response object
type FriendsGetrecent int

// FriendsGetrequestsExtended type represents `friends_getRequests_extended_response` API response object
type FriendsGetrequestsExtended struct {
	Count int                                 `json:"count"` // Total requests number
	Items []objects.FriendsRequestsXtrMessage `json:"items"`
}

// FriendsGetrequestsNeedMutual type represents `friends_getRequests_need_mutual_response` API response object
type FriendsGetrequestsNeedMutual struct {
	Count int                       `json:"count"` // Total requests number
	Items []objects.FriendsRequests `json:"items"`
}

// FriendsGetrequests type represents `friends_getRequests_response` API response object
type FriendsGetrequests struct {
	Count       int   `json:"count"`        // Total requests number
	CountUnread int   `json:"count_unread"` // Total unread requests number
	Items       []int `json:"items"`
}

// FriendsGetsuggestions type represents `friends_getSuggestions_response` API response object
type FriendsGetsuggestions struct {
	Count int                     `json:"count"` // Total results number
	Items []objects.UsersUserFull `json:"items"`
}

// FriendsGetFields type represents `friends_get_fields_response` API response object
type FriendsGetFields struct {
	Count int                           `json:"count"` // Total friends number
	Items []objects.FriendsUserXtrLists `json:"items"`
}

// FriendsGet type represents `friends_get_response` API response object
type FriendsGet struct {
	Count int   `json:"count"` // Total friends number
	Items []int `json:"items"`
}

// FriendsSearch type represents `friends_search_response` API response object
type FriendsSearch struct {
	Count int                     `json:"count"` // Total number
	Items []objects.UsersUserFull `json:"items"`
}
