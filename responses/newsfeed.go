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
// `newsfeed` group of responses
/////////////////////////////////////////////////////////////

// NewsfeedGetbannedExtended type represents `newsfeed_getBanned_extended_response` API response object
type NewsfeedGetbannedExtended struct {
	Groups   []objects.UsersUserFull   `json:"groups"`
	Profiles []objects.GroupsGroupFull `json:"profiles"`
}

// NewsfeedGetbanned type represents `newsfeed_getBanned_response` API response object
type NewsfeedGetbanned struct {
	Groups  []int `json:"groups"`
	Members []int `json:"members"`
}

// NewsfeedGetcomments type represents `newsfeed_getComments_response` API response object
type NewsfeedGetcomments struct {
	Groups   []objects.GroupsGroupFull      `json:"groups"`
	Items    []objects.NewsfeedNewsfeedItem `json:"items"`
	NextFrom string                         `json:"next_from"` // New from value
	Profiles []objects.UsersUserFull        `json:"profiles"`
}

// NewsfeedGetlistsExtended type represents `newsfeed_getLists_extended_response` API response object
type NewsfeedGetlistsExtended struct {
	Count int                        `json:"count"` // Total number
	Items []objects.NewsfeedListFull `json:"items"`
}

// NewsfeedGetlists type represents `newsfeed_getLists_response` API response object
type NewsfeedGetlists struct {
	Count int                    `json:"count"` // Total number
	Items []objects.NewsfeedList `json:"items"`
}

// NewsfeedGetmentions type represents `newsfeed_getMentions_response` API response object
type NewsfeedGetmentions struct {
	Count int                        `json:"count"` // Total number
	Items []objects.WallWallpostToId `json:"items"`
}

// NewsfeedGetrecommended type represents `newsfeed_getRecommended_response` API response object
type NewsfeedGetrecommended struct {
	Groups    []objects.GroupsGroupFull      `json:"groups"`
	Items     []objects.NewsfeedNewsfeedItem `json:"items"`
	NewOffset string                         `json:"new_offset"` // New offset value
	NextFrom  string                         `json:"next_from"`  // Next from value
	Profiles  []objects.UsersUserFull        `json:"profiles"`
}

// NewsfeedGetsuggestedsources type represents `newsfeed_getSuggestedSources_response` API response object
type NewsfeedGetsuggestedsources struct {
	Count int `json:"count"` // Total number
	Items []struct {
		Activity             string                              `json:"activity"`  // Type of group, start date of event or category of public page
		Addresses            objects.GroupsAddressesInfo         `json:"addresses"` // Info about addresses in groups
		AdminLevel           objects.GroupsGroupAdminLevel       `json:"admin_level"`
		AgeLimits            objects.GroupsGroupFullAgeLimits    `json:"age_limits"` // Information whether age limit
		BanInfo              objects.GroupsGroupBanInfo          `json:"ban_info"`   // User ban info
		CanAccessClosed      bool                                `json:"can_access_closed"`
		CanCreateTopic       objects.BaseBoolInt                 `json:"can_create_topic"`       // Information whether current user can create topic
		CanMessage           objects.BaseBoolInt                 `json:"can_message"`            // Information whether current user can send a message to community
		CanPost              objects.BaseBoolInt                 `json:"can_post"`               // Information whether current user can post on community's wall
		CanSeeAllPosts       objects.BaseBoolInt                 `json:"can_see_all_posts"`      // Information whether current user can see all posts on community's wall
		CanSendNotify        objects.BaseBoolInt                 `json:"can_send_notify"`        // Information whether community can send notifications by phone number to current user
		CanSubscribePodcasts bool                                `json:"can_subscribe_podcasts"` // Owner in whitelist or not
		CanSubscribePosts    bool                                `json:"can_subscribe_posts"`    // Can subscribe to wall
		CanUploadVideo       objects.BaseBoolInt                 `json:"can_upload_video"`       // Information whether current user can upload video
		City                 objects.BaseObject                  `json:"city"`
		Contacts             []objects.GroupsContactsItem        `json:"contacts"`
		Counters             objects.GroupsCountersGroup         `json:"counters"`
		Country              objects.BaseCountry                 `json:"country"`
		Cover                objects.GroupsCover                 `json:"cover"`
		Deactivated          string                              `json:"deactivated"` // Returns if a profile is deleted or blocked
		Description          string                              `json:"description"` // Community description
		FinishDate           int                                 `json:"finish_date"` // Finish date in Unixtime format
		FirstName            string                              `json:"first_name"`  // User first name
		FixedPost            int                                 `json:"fixed_post"`  // Fixed post ID
		FriendStatus         objects.FriendsFriendStatusStatus   `json:"friend_status"`
		HasPhoto             objects.BaseBoolInt                 `json:"has_photo"`     // Information whether community has photo
		Hidden               int                                 `json:"hidden"`        // Returns if a profile is hidden.
		Id                   int                                 `json:"id"`            // User ID
		IsAdmin              objects.BaseBoolInt                 `json:"is_admin"`      // Information whether current user is administrator
		IsAdvertiser         objects.BaseBoolInt                 `json:"is_advertiser"` // Information whether current user is advertiser
		IsClosed             bool                                `json:"is_closed"`
		IsFavorite           objects.BaseBoolInt                 `json:"is_favorite"`            // Information whether community is in faves
		IsMember             objects.BaseBoolInt                 `json:"is_member"`              // Information whether current user is member
		IsMessagesBlocked    objects.BaseBoolInt                 `json:"is_messages_blocked"`    // Information whether community can send a message to current user
		IsSubscribed         objects.BaseBoolInt                 `json:"is_subscribed"`          // Information whether current user is subscribed
		IsSubscribedPodcasts bool                                `json:"is_subscribed_podcasts"` // Information whether current user is subscribed to podcasts
		LastName             string                              `json:"last_name"`              // User last name
		Links                []objects.GroupsLinksItem           `json:"links"`
		MainAlbumId          int                                 `json:"main_album_id"` // Community's main photo album ID
		MainSection          objects.GroupsGroupFullMainSection  `json:"main_section"`
		Market               objects.GroupsMarketInfo            `json:"market"`
		MemberStatus         objects.GroupsGroupFullMemberStatus `json:"member_status"` // Current user's member status
		MembersCount         int                                 `json:"members_count"` // Community members number
		Mutual               objects.FriendsRequestsMutual       `json:"mutual"`
		Name                 string                              `json:"name"`          // Community name
		Online               objects.BaseBoolInt                 `json:"online"`        // Information whether the user is online
		OnlineApp            int                                 `json:"online_app"`    // Application ID
		OnlineMobile         objects.BaseBoolInt                 `json:"online_mobile"` // Information whether the user is online in mobile site or application
		OnlineStatus         objects.GroupsOnlineStatus          `json:"online_status"` // Status of replies in community messages
		Photo100             string                              `json:"photo_100"`     // URL of square photo of the user with 100 pixels in width
		Photo200             string                              `json:"photo_200"`     // URL of square photo of the community with 200 pixels in width
		Photo50              string                              `json:"photo_50"`      // URL of square photo of the user with 50 pixels in width
		ScreenName           string                              `json:"screen_name"`   // Domain name of the user's page
		Sex                  objects.BaseSex                     `json:"sex"`           // User sex
		Site                 string                              `json:"site"`          // Community's website
		StartDate            int                                 `json:"start_date"`    // Start date in Unixtime format
		Status               string                              `json:"status"`        // Community status
		Trending             objects.BaseBoolInt                 `json:"trending"`      // Information whether the user has a "fire" pictogram.
		Type                 objects.UsersUserType               `json:"type"`
		Verified             objects.BaseBoolInt                 `json:"verified"`  // Information whether the user is verified
		WikiPage             string                              `json:"wiki_page"` // Community's main wiki page title
	} `json:"items"`
}

// NewsfeedGet type represents `newsfeed_get_response` API response object
type NewsfeedGet struct {
	Groups   []objects.GroupsGroupFull      `json:"groups"`
	Items    []objects.NewsfeedNewsfeedItem `json:"items"`
	NextFrom string                         `json:"next_from"` // New from value
	Profiles []objects.UsersUserFull        `json:"profiles"`
}

// NewsfeedSavelist type represents `newsfeed_saveList_response` API response object
type NewsfeedSavelist int // List ID

// NewsfeedSearchExtended type represents `newsfeed_search_extended_response` API response object
type NewsfeedSearchExtended struct {
	Count      int                        `json:"count"` // Filtered number
	Groups     []objects.GroupsGroupFull  `json:"groups"`
	Items      []objects.WallWallpostFull `json:"items"`
	NextFrom   string                     `json:"next_from"`
	Profiles   []objects.UsersUserFull    `json:"profiles"`
	TotalCount int                        `json:"total_count"` // Total number
}

// NewsfeedSearch type represents `newsfeed_search_response` API response object
type NewsfeedSearch struct {
	Items            []objects.WallWallpostFull `json:"items"`
	SuggestedQueries []string                   `json:"suggested_queries"`
}
