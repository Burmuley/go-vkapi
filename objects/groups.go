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

import (
	"encoding/json"
)

/////////////////////////////////////////////////////////////////////////////////////////////////////////////
// `groups` group of objects
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

// GroupsAddress type represents `groups_address` API object
type GroupsAddress struct {
	AdditionalAddress string                      `json:"additional_address"` // Additional address to the place (6 floor, left door)
	Address           string                      `json:"address"`            // String address to the place (Nevsky, 28)
	CityId            int                         `json:"city_id"`            // City id of address
	CountryId         int                         `json:"country_id"`         // Country id of address
	Distance          int                         `json:"distance"`           // Distance from the point
	Id                int                         `json:"id"`                 // Address id
	Latitude          json.Number                 `json:"latitude"`           // Address latitude
	Longitude         json.Number                 `json:"longitude"`          // Address longitude
	MetroStationId    int                         `json:"metro_station_id"`   // Metro id of address
	Phone             string                      `json:"phone"`              // Address phone
	TimeOffset        int                         `json:"time_offset"`        // Time offset int minutes from utc time
	Timetable         GroupsAddressTimetable      `json:"timetable"`          // Week timetable for the address
	Title             string                      `json:"title"`              // Title of the place (Zinger, etc)
	WorkInfoStatus    GroupsAddressWorkInfoStatus `json:"work_info_status"`   // Status of information about timetable
}

// GroupsAddressTimetable type represents `groups_address_timetable` API object
type GroupsAddressTimetable struct {
	Fri GroupsAddressTimetableDay `json:"fri"` // Timetable for friday
	Mon GroupsAddressTimetableDay `json:"mon"` // Timetable for monday
	Sat GroupsAddressTimetableDay `json:"sat"` // Timetable for saturday
	Sun GroupsAddressTimetableDay `json:"sun"` // Timetable for sunday
	Thu GroupsAddressTimetableDay `json:"thu"` // Timetable for thursday
	Tue GroupsAddressTimetableDay `json:"tue"` // Timetable for tuesday
	Wed GroupsAddressTimetableDay `json:"wed"` // Timetable for wednesday
}

// GroupsAddressTimetableDay type represents `groups_address_timetable_day` API object
type GroupsAddressTimetableDay struct {
	BreakCloseTime int `json:"break_close_time"` // Close time of the break in minutes
	BreakOpenTime  int `json:"break_open_time"`  // Start time of the break in minutes
	CloseTime      int `json:"close_time"`       // Close time in minutes
	OpenTime       int `json:"open_time"`        // Open time in minutes
}

// GroupsAddressWorkInfoStatus type represents `groups_address_work_info_status` API object
type GroupsAddressWorkInfoStatus string // Status of information about timetable

// GroupsAddressesInfo type represents `groups_addresses_info` API object
type GroupsAddressesInfo struct {
	IsEnabled     bool `json:"is_enabled"`      // Information whether addresses is enabled
	MainAddressId int  `json:"main_address_id"` // Main address id for group
}

// GroupsBanInfo type represents `groups_ban_info` API object
type GroupsBanInfo struct {
	AdminId        int                 `json:"admin_id"`        // Administrator ID
	Comment        string              `json:"comment"`         // Comment for a ban
	CommentVisible bool                `json:"comment_visible"` // Show comment for user
	Date           int                 `json:"date"`            // Date when user has been added to blacklist in Unixtime
	EndDate        int                 `json:"end_date"`        // Date when user will be removed from blacklist in Unixtime
	IsClosed       bool                `json:"is_closed"`
	Reason         GroupsBanInfoReason `json:"reason"`
}

// GroupsBanInfoReason type represents `groups_ban_info_reason` API object
type GroupsBanInfoReason int // Ban reason

// GroupsBannedItem type represents `groups_banned_item` API object
type GroupsBannedItem struct {
	GroupsOwnerXtrBanInfo GroupsOwnerXtrBanInfo `json:"groups_owner_xtr_ban_info"`
}

// GroupsCallbackServer type represents `groups_callback_server` API object
type GroupsCallbackServer struct {
	CreatorId int    `json:"creator_id"`
	Id        int    `json:"id"`
	SecretKey string `json:"secret_key"`
	Status    string `json:"status"`
	Title     string `json:"title"`
	Url       string `json:"url"`
}

// GroupsCallbackSettings type represents `groups_callback_settings` API object
type GroupsCallbackSettings struct {
	ApiVersion string               `json:"api_version"` // API version used for the events
	Events     GroupsLongPollEvents `json:"events"`
}

// GroupsContactsItem type represents `groups_contacts_item` API object
type GroupsContactsItem struct {
	Desc   string `json:"desc"`    // Contact description
	Email  string `json:"email"`   // Contact email
	Phone  string `json:"phone"`   // Contact phone
	UserId int    `json:"user_id"` // User ID
}

// GroupsCountersGroup type represents `groups_counters_group` API object
type GroupsCountersGroup struct {
	Addresses int `json:"addresses"` // Addresses number
	Albums    int `json:"albums"`    // Photo albums number
	Audios    int `json:"audios"`    // Audios number
	Docs      int `json:"docs"`      // Docs number
	Market    int `json:"market"`    // Market items number
	Photos    int `json:"photos"`    // Photos number
	Topics    int `json:"topics"`    // Topics number
	Videos    int `json:"videos"`    // Videos number
}

// GroupsCover type represents `groups_cover` API object
type GroupsCover struct {
	Enabled BaseBoolInt `json:"enabled"` // Information whether cover is enabled
	Images  []BaseImage `json:"images"`
}

// GroupsFields type represents `groups_fields` API object
type GroupsFields string

// GroupsFilter type represents `groups_filter` API object
type GroupsFilter string

// GroupsGroup type represents `groups_group` API object
type GroupsGroup struct {
	AdminLevel   GroupsGroupAdminLevel `json:"admin_level"`
	Deactivated  string                `json:"deactivated"`   // Information whether community is banned
	FinishDate   int                   `json:"finish_date"`   // Finish date in Unixtime format
	Id           int                   `json:"id"`            // Community ID
	IsAdmin      BaseBoolInt           `json:"is_admin"`      // Information whether current user is administrator
	IsAdvertiser BaseBoolInt           `json:"is_advertiser"` // Information whether current user is advertiser
	IsClosed     GroupsGroupIsClosed   `json:"is_closed"`
	IsMember     BaseBoolInt           `json:"is_member"`   // Information whether current user is member
	Name         string                `json:"name"`        // Community name
	Photo100     string                `json:"photo_100"`   // URL of square photo of the community with 100 pixels in width
	Photo200     string                `json:"photo_200"`   // URL of square photo of the community with 200 pixels in width
	Photo50      string                `json:"photo_50"`    // URL of square photo of the community with 50 pixels in width
	ScreenName   string                `json:"screen_name"` // Domain of the community page
	StartDate    int                   `json:"start_date"`  // Start date in Unixtime format
	Type         GroupsGroupType       `json:"type"`
}

// GroupsGroupAccess type represents `groups_group_access` API object
type GroupsGroupAccess int

// GroupsGroupAdminLevel type represents `groups_group_admin_level` API object
type GroupsGroupAdminLevel int // Level of current user's credentials as manager

// GroupsGroupAgeLimits type represents `groups_group_age_limits` API object
type GroupsGroupAgeLimits int

// GroupsGroupAudio type represents `groups_group_audio` API object
type GroupsGroupAudio int

// GroupsGroupBanInfo type represents `groups_group_ban_info` API object
type GroupsGroupBanInfo struct {
	Comment string `json:"comment"`  // Ban comment
	EndDate int    `json:"end_date"` // End date of ban in Unixtime
}

// GroupsGroupCategory type represents `groups_group_category` API object
type GroupsGroupCategory struct {
	Id            int                  `json:"id"`   // Category ID
	Name          string               `json:"name"` // Category name
	Subcategories []BaseObjectWithName `json:"subcategories"`
}

// GroupsGroupCategoryFull type represents `groups_group_category_full` API object
type GroupsGroupCategoryFull struct {
	Id            int                   `json:"id"`         // Category ID
	Name          string                `json:"name"`       // Category name
	PageCount     int                   `json:"page_count"` // Pages number
	PagePreviews  []GroupsGroup         `json:"page_previews"`
	Subcategories []GroupsGroupCategory `json:"subcategories"`
}

// GroupsGroupCategoryType type represents `groups_group_category_type` API object
type GroupsGroupCategoryType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// GroupsGroupDocs type represents `groups_group_docs` API object
type GroupsGroupDocs int

// GroupsGroupFull type represents `groups_group_full` API object
type GroupsGroupFull struct {
	Activity             string                      `json:"activity"`  // Type of group, start date of event or category of public page
	Addresses            GroupsAddressesInfo         `json:"addresses"` // Info about addresses in groups
	AdminLevel           GroupsGroupAdminLevel       `json:"admin_level"`
	AgeLimits            GroupsGroupFullAgeLimits    `json:"age_limits"`             // Information whether age limit
	BanInfo              GroupsGroupBanInfo          `json:"ban_info"`               // User ban info
	CanCreateTopic       BaseBoolInt                 `json:"can_create_topic"`       // Information whether current user can create topic
	CanMessage           BaseBoolInt                 `json:"can_message"`            // Information whether current user can send a message to community
	CanPost              BaseBoolInt                 `json:"can_post"`               // Information whether current user can post on community's wall
	CanSeeAllPosts       BaseBoolInt                 `json:"can_see_all_posts"`      // Information whether current user can see all posts on community's wall
	CanSendNotify        BaseBoolInt                 `json:"can_send_notify"`        // Information whether community can send notifications by phone number to current user
	CanSubscribePodcasts bool                        `json:"can_subscribe_podcasts"` // Owner in whitelist or not
	CanSubscribePosts    bool                        `json:"can_subscribe_posts"`    // Can subscribe to wall
	CanUploadVideo       BaseBoolInt                 `json:"can_upload_video"`       // Information whether current user can upload video
	City                 BaseObject                  `json:"city"`
	Contacts             []GroupsContactsItem        `json:"contacts"`
	Counters             GroupsCountersGroup         `json:"counters"`
	Country              BaseCountry                 `json:"country"`
	Cover                GroupsCover                 `json:"cover"`
	Deactivated          string                      `json:"deactivated"`   // Information whether community is banned
	Description          string                      `json:"description"`   // Community description
	FinishDate           int                         `json:"finish_date"`   // Finish date in Unixtime format
	FixedPost            int                         `json:"fixed_post"`    // Fixed post ID
	HasPhoto             BaseBoolInt                 `json:"has_photo"`     // Information whether community has photo
	Id                   int                         `json:"id"`            // Community ID
	IsAdmin              BaseBoolInt                 `json:"is_admin"`      // Information whether current user is administrator
	IsAdvertiser         BaseBoolInt                 `json:"is_advertiser"` // Information whether current user is advertiser
	IsClosed             GroupsGroupIsClosed         `json:"is_closed"`
	IsFavorite           BaseBoolInt                 `json:"is_favorite"`            // Information whether community is in faves
	IsMember             BaseBoolInt                 `json:"is_member"`              // Information whether current user is member
	IsMessagesBlocked    BaseBoolInt                 `json:"is_messages_blocked"`    // Information whether community can send a message to current user
	IsSubscribed         BaseBoolInt                 `json:"is_subscribed"`          // Information whether current user is subscribed
	IsSubscribedPodcasts bool                        `json:"is_subscribed_podcasts"` // Information whether current user is subscribed to podcasts
	Links                []GroupsLinksItem           `json:"links"`
	MainAlbumId          int                         `json:"main_album_id"` // Community's main photo album ID
	MainSection          GroupsGroupFullMainSection  `json:"main_section"`
	Market               GroupsMarketInfo            `json:"market"`
	MemberStatus         GroupsGroupFullMemberStatus `json:"member_status"` // Current user's member status
	MembersCount         int                         `json:"members_count"` // Community members number
	Name                 string                      `json:"name"`          // Community name
	OnlineStatus         GroupsOnlineStatus          `json:"online_status"` // Status of replies in community messages
	Photo100             string                      `json:"photo_100"`     // URL of square photo of the community with 100 pixels in width
	Photo200             string                      `json:"photo_200"`     // URL of square photo of the community with 200 pixels in width
	Photo50              string                      `json:"photo_50"`      // URL of square photo of the community with 50 pixels in width
	ScreenName           string                      `json:"screen_name"`   // Domain of the community page
	Site                 string                      `json:"site"`          // Community's website
	StartDate            int                         `json:"start_date"`    // Start date in Unixtime format
	Status               string                      `json:"status"`        // Community status
	Trending             BaseBoolInt                 `json:"trending"`      // Information whether the community has a "fire" pictogram.
	Type                 GroupsGroupType             `json:"type"`
	Verified             BaseBoolInt                 `json:"verified"`  // Information whether community is verified
	WikiPage             string                      `json:"wiki_page"` // Community's main wiki page title
}

// GroupsGroupFullAgeLimits type represents `groups_group_full_age_limits` API object
type GroupsGroupFullAgeLimits int

// GroupsGroupFullMainSection type represents `groups_group_full_main_section` API object
type GroupsGroupFullMainSection int // Main section of community

// GroupsGroupFullMemberStatus type represents `groups_group_full_member_status` API object
type GroupsGroupFullMemberStatus int

// GroupsGroupIsClosed type represents `groups_group_is_closed` API object
type GroupsGroupIsClosed int // Information whether community is closed

// GroupsGroupLink type represents `groups_group_link` API object
type GroupsGroupLink struct {
	Desc            string      `json:"desc"`             // Link description
	EditTitle       BaseBoolInt `json:"edit_title"`       // Information whether the title can be edited
	Id              int         `json:"id"`               // Link ID
	ImageProcessing BaseBoolInt `json:"image_processing"` // Information whether the image on processing
	Name            string      `json:"name"`             // Link label
	Url             string      `json:"url"`              // Link URL
}

// GroupsGroupMarketCurrency type represents `groups_group_market_currency` API object
type GroupsGroupMarketCurrency int

// GroupsGroupPhotos type represents `groups_group_photos` API object
type GroupsGroupPhotos int

// GroupsGroupPublicCategoryList type represents `groups_group_public_category_list` API object
type GroupsGroupPublicCategoryList struct {
	Id           int                       `json:"id"`
	Name         string                    `json:"name"`
	SubtypesList []GroupsGroupCategoryType `json:"subtypes_list"`
}

// GroupsGroupRole type represents `groups_group_role` API object
type GroupsGroupRole string

// GroupsGroupSettings type represents `groups_group_settings` API object
type GroupsGroupSettings struct {
	Access             int                             `json:"access"`            // Community access settings
	Address            string                          `json:"address"`           // Community's page domain
	Audio              int                             `json:"audio"`             // Audio settings
	Description        string                          `json:"description"`       // Community description
	Docs               int                             `json:"docs"`              // Docs settings
	ObsceneFilter      BaseBoolInt                     `json:"obscene_filter"`    // Information whether the obscene filter is enabled
	ObsceneStopwords   BaseBoolInt                     `json:"obscene_stopwords"` // Information whether the stopwords filter is enabled
	ObsceneWords       string                          `json:"obscene_words"`     // The list of stop words
	Photos             int                             `json:"photos"`            // Photos settings
	PublicCategory     int                             `json:"public_category"`   // Information about the group category
	PublicCategoryList []GroupsGroupPublicCategoryList `json:"public_category_list"`
	PublicSubcategory  int                             `json:"public_subcategory"` // Information about the group subcategory
	Rss                string                          `json:"rss"`                // URL of the RSS feed
	Subject            int                             `json:"subject"`            // Community subject ID
	SubjectList        []GroupsSubjectItem             `json:"subject_list"`
	Title              string                          `json:"title"`   // Community title
	Topics             int                             `json:"topics"`  // Topics settings
	Video              int                             `json:"video"`   // Video settings
	Wall               int                             `json:"wall"`    // Wall settings
	Website            string                          `json:"website"` // Community website
	Wiki               int                             `json:"wiki"`    // Wiki settings
}

// GroupsGroupSubject type represents `groups_group_subject` API object
type GroupsGroupSubject string

// GroupsGroupTopics type represents `groups_group_topics` API object
type GroupsGroupTopics int

// GroupsGroupType type represents `groups_group_type` API object
type GroupsGroupType string // Community type

// GroupsGroupVideo type represents `groups_group_video` API object
type GroupsGroupVideo int

// GroupsGroupWall type represents `groups_group_wall` API object
type GroupsGroupWall int

// GroupsGroupWiki type represents `groups_group_wiki` API object
type GroupsGroupWiki int

// GroupsGroupXtrInvitedBy type represents `groups_group_xtr_invited_by` API object
type GroupsGroupXtrInvitedBy struct {
	AdminLevel   GroupsGroupXtrInvitedByAdminLevel `json:"admin_level"`
	Id           string                            `json:"id"`            // Community ID
	InvitedBy    int                               `json:"invited_by"`    // Inviter ID
	IsAdmin      BaseBoolInt                       `json:"is_admin"`      // Information whether current user is manager
	IsAdvertiser BaseBoolInt                       `json:"is_advertiser"` // Information whether current user is advertiser
	IsClosed     BaseBoolInt                       `json:"is_closed"`     // Information whether community is closed
	IsMember     BaseBoolInt                       `json:"is_member"`     // Information whether current user is member
	Name         string                            `json:"name"`          // Community name
	Photo100     string                            `json:"photo_100"`     // URL of square photo of the community with 100 pixels in width
	Photo200     string                            `json:"photo_200"`     // URL of square photo of the community with 200 pixels in width
	Photo50      string                            `json:"photo_50"`      // URL of square photo of the community with 50 pixels in width
	ScreenName   string                            `json:"screen_name"`   // Domain of the community page
	Type         GroupsGroupXtrInvitedByType       `json:"type"`
}

// GroupsGroupXtrInvitedByAdminLevel type represents `groups_group_xtr_invited_by_admin_level` API object
type GroupsGroupXtrInvitedByAdminLevel int // Level of current user's credentials as manager

// GroupsGroupXtrInvitedByType type represents `groups_group_xtr_invited_by_type` API object
type GroupsGroupXtrInvitedByType string // Community type

// GroupsGroupsArray type represents `groups_groups_array` API object
type GroupsGroupsArray struct {
	Count int   `json:"count"` // Communities number
	Items []int `json:"items"`
}

// GroupsLinksItem type represents `groups_links_item` API object
type GroupsLinksItem struct {
	Desc      string      `json:"desc"`       // Link description
	EditTitle BaseBoolInt `json:"edit_title"` // Information whether the link title can be edited
	Id        int         `json:"id"`         // Link ID
	Name      string      `json:"name"`       // Link title
	Photo100  string      `json:"photo_100"`  // URL of square image of the link with 100 pixels in width
	Photo50   string      `json:"photo_50"`   // URL of square image of the link with 50 pixels in width
	Url       string      `json:"url"`        // Link URL
}

// GroupsLongPollEvents type represents `groups_long_poll_events` API object
type GroupsLongPollEvents struct {
	AudioNew             BaseBoolInt `json:"audio_new"`
	BoardPostDelete      BaseBoolInt `json:"board_post_delete"`
	BoardPostEdit        BaseBoolInt `json:"board_post_edit"`
	BoardPostNew         BaseBoolInt `json:"board_post_new"`
	BoardPostRestore     BaseBoolInt `json:"board_post_restore"`
	GroupChangePhoto     BaseBoolInt `json:"group_change_photo"`
	GroupChangeSettings  BaseBoolInt `json:"group_change_settings"`
	GroupJoin            BaseBoolInt `json:"group_join"`
	GroupLeave           BaseBoolInt `json:"group_leave"`
	GroupOfficersEdit    BaseBoolInt `json:"group_officers_edit"`
	LeadFormsNew         BaseBoolInt `json:"lead_forms_new"`
	MarketCommentDelete  BaseBoolInt `json:"market_comment_delete"`
	MarketCommentEdit    BaseBoolInt `json:"market_comment_edit"`
	MarketCommentNew     BaseBoolInt `json:"market_comment_new"`
	MarketCommentRestore BaseBoolInt `json:"market_comment_restore"`
	MessageAllow         BaseBoolInt `json:"message_allow"`
	MessageDeny          BaseBoolInt `json:"message_deny"`
	MessageNew           BaseBoolInt `json:"message_new"`
	MessageRead          BaseBoolInt `json:"message_read"`
	MessageReply         BaseBoolInt `json:"message_reply"`
	MessageTypingState   BaseBoolInt `json:"message_typing_state"`
	MessagesEdit         BaseBoolInt `json:"messages_edit"`
	PhotoCommentDelete   BaseBoolInt `json:"photo_comment_delete"`
	PhotoCommentEdit     BaseBoolInt `json:"photo_comment_edit"`
	PhotoCommentNew      BaseBoolInt `json:"photo_comment_new"`
	PhotoCommentRestore  BaseBoolInt `json:"photo_comment_restore"`
	PhotoNew             BaseBoolInt `json:"photo_new"`
	PollVoteNew          BaseBoolInt `json:"poll_vote_new"`
	UserBlock            BaseBoolInt `json:"user_block"`
	UserUnblock          BaseBoolInt `json:"user_unblock"`
	VideoCommentDelete   BaseBoolInt `json:"video_comment_delete"`
	VideoCommentEdit     BaseBoolInt `json:"video_comment_edit"`
	VideoCommentNew      BaseBoolInt `json:"video_comment_new"`
	VideoCommentRestore  BaseBoolInt `json:"video_comment_restore"`
	VideoNew             BaseBoolInt `json:"video_new"`
	WallPostNew          BaseBoolInt `json:"wall_post_new"`
	WallReplyDelete      BaseBoolInt `json:"wall_reply_delete"`
	WallReplyEdit        BaseBoolInt `json:"wall_reply_edit"`
	WallReplyNew         BaseBoolInt `json:"wall_reply_new"`
	WallReplyRestore     BaseBoolInt `json:"wall_reply_restore"`
	WallRepost           BaseBoolInt `json:"wall_repost"`
}

// GroupsLongPollServer type represents `groups_long_poll_server` API object
type GroupsLongPollServer struct {
	Key    string `json:"key"`    // Long Poll key
	Server string `json:"server"` // Long Poll server address
	Ts     string `json:"ts"`     // Number of the last event
}

// GroupsLongPollSettings type represents `groups_long_poll_settings` API object
type GroupsLongPollSettings struct {
	ApiVersion string               `json:"api_version"` // API version used for the events
	Events     GroupsLongPollEvents `json:"events"`
	IsEnabled  bool                 `json:"is_enabled"` // Shows whether Long Poll is enabled
}

// GroupsMarketInfo type represents `groups_market_info` API object
type GroupsMarketInfo struct {
	ContactId    int            `json:"contact_id"` // Contact person ID
	Currency     MarketCurrency `json:"currency"`
	CurrencyText string         `json:"currency_text"` // Currency name
	Enabled      BaseBoolInt    `json:"enabled"`       // Information whether the market is enabled
	MainAlbumId  int            `json:"main_album_id"` // Main market album ID
	PriceMax     int            `json:"price_max"`     // Maximum price
	PriceMin     int            `json:"price_min"`     // Minimum price
}

// GroupsMemberRole type represents `groups_member_role` API object
type GroupsMemberRole struct {
	Id          int                          `json:"id"` // User ID
	Permissions []GroupsMemberRolePermission `json:"permissions"`
	Role        GroupsMemberRoleStatus       `json:"role"`
}

// GroupsMemberRolePermission type represents `groups_member_role_permission` API object
type GroupsMemberRolePermission string

// GroupsMemberRoleStatus type represents `groups_member_role_status` API object
type GroupsMemberRoleStatus string // User's credentials as community admin

// GroupsMemberStatus type represents `groups_member_status` API object
type GroupsMemberStatus struct {
	Member BaseBoolInt `json:"member"`  // Information whether user is a member of the group
	UserId int         `json:"user_id"` // User ID
}

// GroupsMemberStatusFull type represents `groups_member_status_full` API object
type GroupsMemberStatusFull struct {
	CanInvite  BaseBoolInt `json:"can_invite"` // Information whether user can be invited
	CanRecall  BaseBoolInt `json:"can_recall"` // Information whether user's invite to the group can be recalled
	Invitation BaseBoolInt `json:"invitation"` // Information whether user has been invited to the group
	Member     BaseBoolInt `json:"member"`     // Information whether user is a member of the group
	Request    BaseBoolInt `json:"request"`    // Information whether user has send request to the group
	UserId     int         `json:"user_id"`    // User ID
}

// GroupsOnlineStatus type represents `groups_online_status` API object
type GroupsOnlineStatus struct {
	Minutes int                    `json:"minutes"` // Estimated time of answer (for status = answer_mark)
	Status  GroupsOnlineStatusType `json:"status"`
}

// GroupsOnlineStatusType type represents `groups_online_status_type` API object
type GroupsOnlineStatusType string // Type of online status of group

// GroupsOwnerXtrBanInfo type represents `groups_owner_xtr_ban_info` API object
type GroupsOwnerXtrBanInfo struct {
	BanInfo GroupsBanInfo             `json:"ban_info"`
	Group   GroupsGroup               `json:"group"`   // Information about group if type = group
	Profile UsersUser                 `json:"profile"` // Information about group if type = profile
	Type    GroupsOwnerXtrBanInfoType `json:"type"`
}

// GroupsOwnerXtrBanInfoType type represents `groups_owner_xtr_ban_info_type` API object
type GroupsOwnerXtrBanInfoType string // Owner type

// GroupsRoleOptions type represents `groups_role_options` API object
type GroupsRoleOptions string // User's credentials as community admin

// GroupsSubjectItem type represents `groups_subject_item` API object
type GroupsSubjectItem struct {
	Id   int    `json:"id"`   // Subject ID
	Name string `json:"name"` // Subject title
}

// GroupsTokenPermissionSetting type represents `groups_token_permission_setting` API object
type GroupsTokenPermissionSetting struct {
	Name    string `json:"name"`
	Setting int    `json:"setting"`
}

// GroupsUserXtrRole type represents `groups_user_xtr_role` API object
type GroupsUserXtrRole struct {
	Role GroupsRoleOptions `json:"role"`
}
