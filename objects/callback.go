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
// `callback` group of objects
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

// CallbackBoardPostDelete type represents `callback_board_post_delete` API object
type CallbackBoardPostDelete struct {
	Id           int `json:"id"`
	TopicId      int `json:"topic_id"`
	TopicOwnerId int `json:"topic_owner_id"`
}

// CallbackConfirmationMessage type represents `callback_confirmation_message` API object
type CallbackConfirmationMessage struct {
	GroupId int                 `json:"group_id"`
	Secret  string              `json:"secret"`
	Type    CallbackMessageType `json:"type"`
}

// CallbackGroupChangePhoto type represents `callback_group_change_photo` API object
type CallbackGroupChangePhoto struct {
	Photo  PhotosPhoto `json:"photo"`
	UserId int         `json:"user_id"`
}

// CallbackGroupChangeSettings type represents `callback_group_change_settings` API object
type CallbackGroupChangeSettings struct {
	Self   BaseBoolInt `json:"self"`
	UserId int         `json:"user_id"`
}

// CallbackGroupJoin type represents `callback_group_join` API object
type CallbackGroupJoin struct {
	JoinType CallbackGroupJoinType `json:"join_type"`
	UserId   int                   `json:"user_id"`
}

// CallbackGroupJoinType type represents `callback_group_join_type` API object
type CallbackGroupJoinType string

// CallbackGroupLeave type represents `callback_group_leave` API object
type CallbackGroupLeave struct {
	Self   BaseBoolInt `json:"self"`
	UserId int         `json:"user_id"`
}

// CallbackGroupMarket type represents `callback_group_market` API object
type CallbackGroupMarket int

// CallbackGroupOfficerRole type represents `callback_group_officer_role` API object
type CallbackGroupOfficerRole int

// CallbackGroupOfficersEdit type represents `callback_group_officers_edit` API object
type CallbackGroupOfficersEdit struct {
	AdminId  int                      `json:"admin_id"`
	LevelNew CallbackGroupOfficerRole `json:"level_new"`
	LevelOld CallbackGroupOfficerRole `json:"level_old"`
	UserId   int                      `json:"user_id"`
}

// CallbackGroupSettingsChanges type represents `callback_group_settings_changes` API object
type CallbackGroupSettingsChanges struct {
	Access              GroupsGroupIsClosed      `json:"access"`
	AgeLimits           GroupsGroupFullAgeLimits `json:"age_limits"`
	Description         string                   `json:"description"`
	EnableAudio         GroupsGroupAudio         `json:"enable_audio"`
	EnableMarket        CallbackGroupMarket      `json:"enable_market"`
	EnablePhoto         GroupsGroupPhotos        `json:"enable_photo"`
	EnableStatusDefault GroupsGroupWall          `json:"enable_status_default"`
	EnableVideo         GroupsGroupVideo         `json:"enable_video"`
	PublicCategory      int                      `json:"public_category"`
	PublicSubcategory   int                      `json:"public_subcategory"`
	ScreenName          string                   `json:"screen_name"`
	Title               string                   `json:"title"`
	Website             string                   `json:"website"`
}

// CallbackMarketComment type represents `callback_market_comment` API object
type CallbackMarketComment struct {
	Date          int    `json:"date"`
	FromId        int    `json:"from_id"`
	Id            int    `json:"id"`
	MarketOwnerOd int    `json:"market_owner_od"`
	PhotoId       int    `json:"photo_id"`
	Text          string `json:"text"`
}

// CallbackMarketCommentDelete type represents `callback_market_comment_delete` API object
type CallbackMarketCommentDelete struct {
	Id      int `json:"id"`
	ItemId  int `json:"item_id"`
	OwnerId int `json:"owner_id"`
	UserId  int `json:"user_id"`
}

// CallbackMessageAllow type represents `callback_message_allow` API object
type CallbackMessageAllow struct {
	Key    string `json:"key"`
	UserId int    `json:"user_id"`
}

// CallbackMessageBase type represents `callback_message_base` API object
type CallbackMessageBase struct {
	GroupId int                 `json:"group_id"`
	Object  object              `json:"object"`
	Type    CallbackMessageType `json:"type"`
}

// CallbackMessageDeny type represents `callback_message_deny` API object
type CallbackMessageDeny struct {
	UserId int `json:"user_id"`
}

// CallbackMessageType type represents `callback_message_type` API object
type CallbackMessageType string

// CallbackPhotoComment type represents `callback_photo_comment` API object
type CallbackPhotoComment struct {
	Date         int    `json:"date"`
	FromId       int    `json:"from_id"`
	Id           int    `json:"id"`
	PhotoOwnerOd int    `json:"photo_owner_od"`
	Text         string `json:"text"`
}

// CallbackPhotoCommentDelete type represents `callback_photo_comment_delete` API object
type CallbackPhotoCommentDelete struct {
	Id      int `json:"id"`
	OwnerId int `json:"owner_id"`
	PhotoId int `json:"photo_id"`
	UserId  int `json:"user_id"`
}

// CallbackPollVoteNew type represents `callback_poll_vote_new` API object
type CallbackPollVoteNew struct {
	OptionId int `json:"option_id"`
	OwnerId  int `json:"owner_id"`
	PollId   int `json:"poll_id"`
	UserId   int `json:"user_id"`
}

// CallbackUserBlock type represents `callback_user_block` API object
type CallbackUserBlock struct {
	AdminId     int    `json:"admin_id"`
	Comment     string `json:"comment"`
	Reason      int    `json:"reason"`
	UnblockDate int    `json:"unblock_date"`
	UserId      int    `json:"user_id"`
}

// CallbackUserUnblock type represents `callback_user_unblock` API object
type CallbackUserUnblock struct {
	AdminId   int `json:"admin_id"`
	ByEndDate int `json:"by_end_date"`
	UserId    int `json:"user_id"`
}

// CallbackVideoComment type represents `callback_video_comment` API object
type CallbackVideoComment struct {
	Date         int    `json:"date"`
	FromId       int    `json:"from_id"`
	Id           int    `json:"id"`
	Text         string `json:"text"`
	VideoOwnerOd int    `json:"video_owner_od"`
}

// CallbackVideoCommentDelete type represents `callback_video_comment_delete` API object
type CallbackVideoCommentDelete struct {
	Id      int `json:"id"`
	OwnerId int `json:"owner_id"`
	UserId  int `json:"user_id"`
	VideoId int `json:"video_id"`
}

// CallbackWallCommentDelete type represents `callback_wall_comment_delete` API object
type CallbackWallCommentDelete struct {
	Id      int `json:"id"`
	OwnerId int `json:"owner_id"`
	PostId  int `json:"post_id"`
	UserId  int `json:"user_id"`
}
