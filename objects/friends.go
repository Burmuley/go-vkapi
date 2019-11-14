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

package objects

/////////////////////////////////////////////////////////////////////////////////////////////////////////////
// `friends` group of objects
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

// FriendsFriendStatus type represents `friends_friend_status` API object
type FriendsFriendStatus struct {
	FriendStatus   FriendsFriendStatusStatus `json:"friend_status"`
	ReadState      BaseBoolInt               `json:"read_state"`      // Information whether request is unviewed
	RequestMessage string                    `json:"request_message"` // Message sent with request
	Sign           string                    `json:"sign"`            // MD5 hash for the result validation
	UserId         int                       `json:"user_id"`         // User ID
}

// FriendsFriendStatusStatus type represents `friends_friend_status_status` API object
type FriendsFriendStatusStatus int // Friend status with the user

// FriendsFriendsList type represents `friends_friends_list` API object
type FriendsFriendsList struct {
	Id   int    `json:"id"`   // List ID
	Name string `json:"name"` // List title
}

// FriendsMutualFriend type represents `friends_mutual_friend` API object
type FriendsMutualFriend struct {
	CommonCount   int   `json:"common_count"` // Total mutual friends number
	CommonFriends []int `json:"common_friends"`
	Id            int   `json:"id"` // User ID
}

// FriendsRequests type represents `friends_requests` API object
type FriendsRequests struct {
	From   string                `json:"from"` // ID of the user by whom friend has been suggested
	Mutual FriendsRequestsMutual `json:"mutual"`
	UserId int                   `json:"user_id"` // User ID
}

// FriendsRequestsMutual type represents `friends_requests_mutual` API object
type FriendsRequestsMutual struct {
	Count int   `json:"count"` // Total mutual friends number
	Users []int `json:"users"`
}

// FriendsRequestsXtrMessage type represents `friends_requests_xtr_message` API object
type FriendsRequestsXtrMessage struct {
	From    string                `json:"from"`    // ID of the user by whom friend has been suggested
	Message string                `json:"message"` // Message sent with a request
	Mutual  FriendsRequestsMutual `json:"mutual"`
	UserId  int                   `json:"user_id"` // User ID
}

// FriendsUserXtrLists type represents `friends_user_xtr_lists` API object
type FriendsUserXtrLists struct {
	Lists []int `json:"lists"`
}

// FriendsUserXtrPhone type represents `friends_user_xtr_phone` API object
type FriendsUserXtrPhone struct {
	Phone string `json:"phone"` // User phone
}
