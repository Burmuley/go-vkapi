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
	Activity               string                    `json:"activity"`          // User's status
	Bdate                  string                    `json:"bdate"`             // User's date of birth
	Blacklisted            BaseBoolInt               `json:"blacklisted"`       // Information whether current user is in the requested user's blacklist.
	BlacklistedByMe        BaseBoolInt               `json:"blacklisted_by_me"` // Information whether the requested user is in current user's blacklist
	CanAccessClosed        bool                      `json:"can_access_closed"`
	CanPost                BaseBoolInt               `json:"can_post"`                  // Information whether current user can post on the user's wall
	CanSeeAllPosts         BaseBoolInt               `json:"can_see_all_posts"`         // Information whether current user can see other users' audio on the wall
	CanSeeAudio            BaseBoolInt               `json:"can_see_audio"`             // Information whether current user can see the user's audio
	CanSendFriendRequest   BaseBoolInt               `json:"can_send_friend_request"`   // Information whether current user can send a friend request
	CanSubscribePodcasts   bool                      `json:"can_subscribe_podcasts"`    // Owner in whitelist or not
	CanSubscribePosts      bool                      `json:"can_subscribe_posts"`       // Can subscribe to wall
	CanWritePrivateMessage BaseBoolInt               `json:"can_write_private_message"` // Information whether current user can write private message
	Career                 []UsersCareer             `json:"career"`
	City                   BaseObject                `json:"city"`
	CommonCount            int                       `json:"common_count"` // Number of common friends with current user
	Country                BaseCountry               `json:"country"`
	CropPhoto              UsersCropPhoto            `json:"crop_photo"`
	Deactivated            string                    `json:"deactivated"`      // Returns if a profile is deleted or blocked
	Domain                 string                    `json:"domain"`           // Domain name of the user's page
	EducationForm          string                    `json:"education_form"`   // Education form
	EducationStatus        string                    `json:"education_status"` // User's education status
	Exports                UsersExports              `json:"exports"`
	Faculty                int                       `json:"faculty"`         // Faculty ID
	FacultyName            string                    `json:"faculty_name"`    // Faculty name
	FirstName              string                    `json:"first_name"`      // User first name
	FollowersCount         int                       `json:"followers_count"` // Number of user's followers
	FriendStatus           FriendsFriendStatusStatus `json:"friend_status"`
	Graduation             int                       `json:"graduation"` // Graduation year
	HasMobile              BaseBoolInt               `json:"has_mobile"` // Information whether the user specified his phone number
	HasPhoto               BaseBoolInt               `json:"has_photo"`  // Information whether the user has main photo
	Hidden                 int                       `json:"hidden"`     // Returns if a profile is hidden.
	HomePhone              string                    `json:"home_phone"` // User's mobile phone number
	HomeTown               string                    `json:"home_town"`  // User hometown
	Id                     int                       `json:"id"`         // User ID
	IsClosed               bool                      `json:"is_closed"`
	IsFavorite             BaseBoolInt               `json:"is_favorite"`            // Information whether the requested user is in faves of current user
	IsFriend               BaseBoolInt               `json:"is_friend"`              // Information whether the user is a friend of current user
	IsHiddenFromFeed       BaseBoolInt               `json:"is_hidden_from_feed"`    // Information whether the requested user is hidden from current user's newsfeed
	IsSubscribedPodcasts   bool                      `json:"is_subscribed_podcasts"` // Information whether current user is subscribed to podcasts
	LastName               string                    `json:"last_name"`              // User last name
	LastSeen               UsersLastSeen             `json:"last_seen"`
	Lists                  []int                     `json:"lists"`
	MaidenName             string                    `json:"maiden_name"` // User maiden name
	Military               []UsersMilitary           `json:"military"`
	MobilePhone            string                    `json:"mobile_phone"` // Information whether current user can see
	Mutual                 FriendsRequestsMutual     `json:"mutual"`
	Nickname               string                    `json:"nickname"` // User nickname
	Occupation             UsersOccupation           `json:"occupation"`
	Online                 BaseBoolInt               `json:"online"`        // Information whether the user is online
	OnlineApp              int                       `json:"online_app"`    // Application ID
	OnlineMobile           BaseBoolInt               `json:"online_mobile"` // Information whether the user is online in mobile site or application
	Personal               UsersPersonal             `json:"personal"`
	Photo100               string                    `json:"photo_100"`      // URL of square photo of the user with 100 pixels in width
	Photo200               string                    `json:"photo_200"`      // URL of square photo of the user with 200 pixels in width
	Photo200Orig           string                    `json:"photo_200_orig"` // URL of user's photo with 200 pixels in width
	Photo400Orig           string                    `json:"photo_400_orig"` // URL of user's photo with 400 pixels in width
	Photo50                string                    `json:"photo_50"`       // URL of square photo of the user with 50 pixels in width
	PhotoId                string                    `json:"photo_id"`       // ID of the user's main photo
	PhotoMax               string                    `json:"photo_max"`      // URL of square photo of the user with maximum width
	PhotoMaxOrig           string                    `json:"photo_max_orig"` // URL of user's photo of maximum size
	Relation               UsersUserRelation         `json:"relation"`       // User relationship status
	RelationPartner        UsersUserMin              `json:"relation_partner"`
	Relatives              []UsersRelative           `json:"relatives"`
	Schools                []UsersSchool             `json:"schools"`
	ScreenName             string                    `json:"screen_name"` // Domain name of the user's page
	Sex                    BaseSex                   `json:"sex"`         // User sex
	Site                   string                    `json:"site"`        // User's website
	Status                 string                    `json:"status"`      // User's status
	StatusAudio            AudioAudio                `json:"status_audio"`
	Timezone               int                       `json:"timezone"` // User's timezone
	Trending               BaseBoolInt               `json:"trending"` // Information whether the user has a "fire" pictogram.
	Universities           []UsersUniversity         `json:"universities"`
	University             int                       `json:"university"`      // University ID
	UniversityName         string                    `json:"university_name"` // University name
	Verified               BaseBoolInt               `json:"verified"`        // Information whether the user is verified
	WallComments           BaseBoolInt               `json:"wall_comments"`   // Information whether current user can comment wall posts
}

// FriendsUserXtrPhone type represents `friends_user_xtr_phone` API object
type FriendsUserXtrPhone struct {
	Activity               string                    `json:"activity"`          // User's status
	Bdate                  string                    `json:"bdate"`             // User's date of birth
	Blacklisted            BaseBoolInt               `json:"blacklisted"`       // Information whether current user is in the requested user's blacklist.
	BlacklistedByMe        BaseBoolInt               `json:"blacklisted_by_me"` // Information whether the requested user is in current user's blacklist
	CanAccessClosed        bool                      `json:"can_access_closed"`
	CanPost                BaseBoolInt               `json:"can_post"`                  // Information whether current user can post on the user's wall
	CanSeeAllPosts         BaseBoolInt               `json:"can_see_all_posts"`         // Information whether current user can see other users' audio on the wall
	CanSeeAudio            BaseBoolInt               `json:"can_see_audio"`             // Information whether current user can see the user's audio
	CanSendFriendRequest   BaseBoolInt               `json:"can_send_friend_request"`   // Information whether current user can send a friend request
	CanSubscribePodcasts   bool                      `json:"can_subscribe_podcasts"`    // Owner in whitelist or not
	CanSubscribePosts      bool                      `json:"can_subscribe_posts"`       // Can subscribe to wall
	CanWritePrivateMessage BaseBoolInt               `json:"can_write_private_message"` // Information whether current user can write private message
	Career                 []UsersCareer             `json:"career"`
	City                   BaseObject                `json:"city"`
	CommonCount            int                       `json:"common_count"` // Number of common friends with current user
	Country                BaseCountry               `json:"country"`
	CropPhoto              UsersCropPhoto            `json:"crop_photo"`
	Deactivated            string                    `json:"deactivated"`      // Returns if a profile is deleted or blocked
	Domain                 string                    `json:"domain"`           // Domain name of the user's page
	EducationForm          string                    `json:"education_form"`   // Education form
	EducationStatus        string                    `json:"education_status"` // User's education status
	Exports                UsersExports              `json:"exports"`
	Faculty                int                       `json:"faculty"`         // Faculty ID
	FacultyName            string                    `json:"faculty_name"`    // Faculty name
	FirstName              string                    `json:"first_name"`      // User first name
	FollowersCount         int                       `json:"followers_count"` // Number of user's followers
	FriendStatus           FriendsFriendStatusStatus `json:"friend_status"`
	Graduation             int                       `json:"graduation"` // Graduation year
	HasMobile              BaseBoolInt               `json:"has_mobile"` // Information whether the user specified his phone number
	HasPhoto               BaseBoolInt               `json:"has_photo"`  // Information whether the user has main photo
	Hidden                 int                       `json:"hidden"`     // Returns if a profile is hidden.
	HomePhone              string                    `json:"home_phone"` // User's mobile phone number
	HomeTown               string                    `json:"home_town"`  // User hometown
	Id                     int                       `json:"id"`         // User ID
	IsClosed               bool                      `json:"is_closed"`
	IsFavorite             BaseBoolInt               `json:"is_favorite"`            // Information whether the requested user is in faves of current user
	IsFriend               BaseBoolInt               `json:"is_friend"`              // Information whether the user is a friend of current user
	IsHiddenFromFeed       BaseBoolInt               `json:"is_hidden_from_feed"`    // Information whether the requested user is hidden from current user's newsfeed
	IsSubscribedPodcasts   bool                      `json:"is_subscribed_podcasts"` // Information whether current user is subscribed to podcasts
	LastName               string                    `json:"last_name"`              // User last name
	LastSeen               UsersLastSeen             `json:"last_seen"`
	MaidenName             string                    `json:"maiden_name"` // User maiden name
	Military               []UsersMilitary           `json:"military"`
	MobilePhone            string                    `json:"mobile_phone"` // Information whether current user can see
	Mutual                 FriendsRequestsMutual     `json:"mutual"`
	Nickname               string                    `json:"nickname"` // User nickname
	Occupation             UsersOccupation           `json:"occupation"`
	Online                 BaseBoolInt               `json:"online"`        // Information whether the user is online
	OnlineApp              int                       `json:"online_app"`    // Application ID
	OnlineMobile           BaseBoolInt               `json:"online_mobile"` // Information whether the user is online in mobile site or application
	Personal               UsersPersonal             `json:"personal"`
	Phone                  string                    `json:"phone"`          // User phone
	Photo100               string                    `json:"photo_100"`      // URL of square photo of the user with 100 pixels in width
	Photo200               string                    `json:"photo_200"`      // URL of square photo of the user with 200 pixels in width
	Photo200Orig           string                    `json:"photo_200_orig"` // URL of user's photo with 200 pixels in width
	Photo400Orig           string                    `json:"photo_400_orig"` // URL of user's photo with 400 pixels in width
	Photo50                string                    `json:"photo_50"`       // URL of square photo of the user with 50 pixels in width
	PhotoId                string                    `json:"photo_id"`       // ID of the user's main photo
	PhotoMax               string                    `json:"photo_max"`      // URL of square photo of the user with maximum width
	PhotoMaxOrig           string                    `json:"photo_max_orig"` // URL of user's photo of maximum size
	Relation               UsersUserRelation         `json:"relation"`       // User relationship status
	RelationPartner        UsersUserMin              `json:"relation_partner"`
	Relatives              []UsersRelative           `json:"relatives"`
	Schools                []UsersSchool             `json:"schools"`
	ScreenName             string                    `json:"screen_name"` // Domain name of the user's page
	Sex                    BaseSex                   `json:"sex"`         // User sex
	Site                   string                    `json:"site"`        // User's website
	Status                 string                    `json:"status"`      // User's status
	StatusAudio            AudioAudio                `json:"status_audio"`
	Timezone               int                       `json:"timezone"` // User's timezone
	Trending               BaseBoolInt               `json:"trending"` // Information whether the user has a "fire" pictogram.
	Universities           []UsersUniversity         `json:"universities"`
	University             int                       `json:"university"`      // University ID
	UniversityName         string                    `json:"university_name"` // University name
	Verified               BaseBoolInt               `json:"verified"`        // Information whether the user is verified
	WallComments           BaseBoolInt               `json:"wall_comments"`   // Information whether current user can comment wall posts
}
