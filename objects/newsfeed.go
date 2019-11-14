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
// `newsfeed` group of objects
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

// NewsfeedCommentsFilters type represents `newsfeed_comments_filters` API object
type NewsfeedCommentsFilters string

// NewsfeedEventActivity type represents `newsfeed_event_activity` API object
type NewsfeedEventActivity struct {
	Address      string                      `json:"address"`       // address of event
	ButtonText   string                      `json:"button_text"`   // text of attach
	Friends      []int                       `json:"friends"`       // array of friends ids
	MemberStatus GroupsGroupFullMemberStatus `json:"member_status"` // Current user's member status
	Text         string                      `json:"text"`          // text of attach
	Time         int                         `json:"time"`          // event start time
}

// NewsfeedFilters type represents `newsfeed_filters` API object
type NewsfeedFilters string

// NewsfeedIgnoreItemType type represents `newsfeed_ignore_item_type` API object
type NewsfeedIgnoreItemType string

// NewsfeedItemAudio type represents `newsfeed_item_audio` API object
type NewsfeedItemAudio struct {
	Audio    NewsfeedItemAudioAudio   `json:"audio"`
	Date     int                      `json:"date"`      // Date when item has been added in Unixtime
	PostId   int                      `json:"post_id"`   // Post ID
	SourceId int                      `json:"source_id"` // Item source ID
	Type     NewsfeedNewsfeedItemType `json:"type"`
}

// NewsfeedItemAudioAudio type represents `newsfeed_item_audio_audio` API object
type NewsfeedItemAudioAudio struct {
	Count int          `json:"count"` // Audios number
	Items []AudioAudio `json:"items"`
}

// NewsfeedItemBase type represents `newsfeed_item_base` API object
type NewsfeedItemBase struct {
	Date     int                      `json:"date"`      // Date when item has been added in Unixtime
	SourceId int                      `json:"source_id"` // Item source ID
	Type     NewsfeedNewsfeedItemType `json:"type"`
}

// NewsfeedItemDigest type represents `newsfeed_item_digest` API object
type NewsfeedItemDigest struct {
	ButtonText  string                   `json:"button_text"`
	Date        int                      `json:"date"`    // Date when item has been added in Unixtime
	FeedId      string                   `json:"feed_id"` // id of feed in digest
	Items       []WallWallpost           `json:"items"`
	MainPostIds []string                 `json:"main_post_ids"`
	SourceId    int                      `json:"source_id"` // Item source ID
	Template    string                   `json:"template"`  // type of digest
	Title       string                   `json:"title"`
	TrackCode   string                   `json:"track_code"`
	Type        NewsfeedNewsfeedItemType `json:"type"`
}

// NewsfeedItemFriend type represents `newsfeed_item_friend` API object
type NewsfeedItemFriend struct {
	Date     int                       `json:"date"` // Date when item has been added in Unixtime
	Friends  NewsfeedItemFriendFriends `json:"friends"`
	SourceId int                       `json:"source_id"` // Item source ID
	Type     NewsfeedNewsfeedItemType  `json:"type"`
}

// NewsfeedItemFriendFriends type represents `newsfeed_item_friend_friends` API object
type NewsfeedItemFriendFriends struct {
	Count int          `json:"count"` // Number of friends has been added
	Items []BaseUserId `json:"items"`
}

// NewsfeedItemNote type represents `newsfeed_item_note` API object
type NewsfeedItemNote struct {
	Date     int                      `json:"date"` // Date when item has been added in Unixtime
	Notes    NewsfeedItemNoteNotes    `json:"notes"`
	SourceId int                      `json:"source_id"` // Item source ID
	Type     NewsfeedNewsfeedItemType `json:"type"`
}

// NewsfeedItemNoteNotes type represents `newsfeed_item_note_notes` API object
type NewsfeedItemNoteNotes struct {
	Count int                    `json:"count"` // Notes number
	Items []NewsfeedNewsfeedNote `json:"items"`
}

// NewsfeedItemPhoto type represents `newsfeed_item_photo` API object
type NewsfeedItemPhoto struct {
	Date     int                      `json:"date"` // Date when item has been added in Unixtime
	Photos   NewsfeedItemPhotoPhotos  `json:"photos"`
	PostId   int                      `json:"post_id"`   // Post ID
	SourceId int                      `json:"source_id"` // Item source ID
	Type     NewsfeedNewsfeedItemType `json:"type"`
}

// NewsfeedItemPhotoPhotos type represents `newsfeed_item_photo_photos` API object
type NewsfeedItemPhotoPhotos struct {
	Count int                     `json:"count"` // Photos number
	Items []NewsfeedNewsfeedPhoto `json:"items"`
}

// NewsfeedItemPhotoTag type represents `newsfeed_item_photo_tag` API object
type NewsfeedItemPhotoTag struct {
	Date      int                           `json:"date"` // Date when item has been added in Unixtime
	PhotoTags NewsfeedItemPhotoTagPhotoTags `json:"photo_tags"`
	PostId    int                           `json:"post_id"`   // Post ID
	SourceId  int                           `json:"source_id"` // Item source ID
	Type      NewsfeedNewsfeedItemType      `json:"type"`
}

// NewsfeedItemPhotoTagPhotoTags type represents `newsfeed_item_photo_tag_photo_tags` API object
type NewsfeedItemPhotoTagPhotoTags struct {
	Count int                     `json:"count"` // Tags number
	Items []NewsfeedNewsfeedPhoto `json:"items"`
}

// NewsfeedItemStoriesBlock type represents `newsfeed_item_stories_block` API object
type NewsfeedItemStoriesBlock struct {
	BlockType string                   `json:"block_type"`
	Date      int                      `json:"date"`      // Date when item has been added in Unixtime
	SourceId  int                      `json:"source_id"` // Item source ID
	Stories   []StoriesStory           `json:"stories"`
	Title     string                   `json:"title"`
	TrackCode string                   `json:"track_code"`
	Type      NewsfeedNewsfeedItemType `json:"type"`
}

// NewsfeedItemTopic type represents `newsfeed_item_topic` API object
type NewsfeedItemTopic struct {
	Comments BaseCommentsInfo         `json:"comments"`
	Date     int                      `json:"date"` // Date when item has been added in Unixtime
	Likes    BaseLikesInfo            `json:"likes"`
	PostId   int                      `json:"post_id"`   // Topic post ID
	SourceId int                      `json:"source_id"` // Item source ID
	Text     string                   `json:"text"`      // Post text
	Type     NewsfeedNewsfeedItemType `json:"type"`
}

// NewsfeedItemVideo type represents `newsfeed_item_video` API object
type NewsfeedItemVideo struct {
	Date     int                      `json:"date"`      // Date when item has been added in Unixtime
	SourceId int                      `json:"source_id"` // Item source ID
	Type     NewsfeedNewsfeedItemType `json:"type"`
	Video    NewsfeedItemVideoVideo   `json:"video"`
}

// NewsfeedItemVideoVideo type represents `newsfeed_item_video_video` API object
type NewsfeedItemVideoVideo struct {
	Count int          `json:"count"` // Tags number
	Items []VideoVideo `json:"items"`
}

// NewsfeedItemWallpost type represents `newsfeed_item_wallpost` API object
type NewsfeedItemWallpost struct {
	Activity    NewsfeedEventActivity    `json:"activity"`
	Attachments []WallWallpostAttachment `json:"attachments"`
	Comments    BaseCommentsInfo         `json:"comments"`
	CopyHistory []WallWallpost           `json:"copy_history"`
	Date        int                      `json:"date"` // Date when item has been added in Unixtime
	Geo         BaseGeo                  `json:"geo"`
	Likes       BaseLikesInfo            `json:"likes"`
	PostId      int                      `json:"post_id"` // Post ID
	PostSource  WallPostSource           `json:"post_source"`
	PostType    NewsfeedItemWallpostType `json:"post_type"`
	Reposts     BaseRepostsInfo          `json:"reposts"`
	SourceId    int                      `json:"source_id"` // Item source ID
	Text        string                   `json:"text"`      // Post text
	Type        NewsfeedNewsfeedItemType `json:"type"`
}

// NewsfeedItemWallpostType type represents `newsfeed_item_wallpost_type` API object
type NewsfeedItemWallpostType string // Post type

// NewsfeedList type represents `newsfeed_list` API object
type NewsfeedList struct {
	Id    int    `json:"id"`    // List ID
	Title string `json:"title"` // List title
}

// NewsfeedListFull type represents `newsfeed_list_full` API object
type NewsfeedListFull struct {
	Id        int         `json:"id"`         // List ID
	NoReposts BaseBoolInt `json:"no_reposts"` // Information whether reposts hiding is enabled
	SourceIds []int       `json:"source_ids"`
	Title     string      `json:"title"` // List title
}

// NewsfeedNewsfeedItem type represents `newsfeed_newsfeed_item` API object
type NewsfeedNewsfeedItem struct {
	NewsfeedItemAudio        NewsfeedItemAudio        `json:"newsfeed_item_audio"`
	NewsfeedItemDigest       NewsfeedItemDigest       `json:"newsfeed_item_digest"`
	NewsfeedItemFriend       NewsfeedItemFriend       `json:"newsfeed_item_friend"`
	NewsfeedItemNote         NewsfeedItemNote         `json:"newsfeed_item_note"`
	NewsfeedItemPhoto        NewsfeedItemPhoto        `json:"newsfeed_item_photo"`
	NewsfeedItemPhotoTag     NewsfeedItemPhotoTag     `json:"newsfeed_item_photo_tag"`
	NewsfeedItemStoriesBlock NewsfeedItemStoriesBlock `json:"newsfeed_item_stories_block"`
	NewsfeedItemTopic        NewsfeedItemTopic        `json:"newsfeed_item_topic"`
	NewsfeedItemVideo        NewsfeedItemVideo        `json:"newsfeed_item_video"`
	NewsfeedItemWallpost     NewsfeedItemWallpost     `json:"newsfeed_item_wallpost"`
}

// NewsfeedNewsfeedItemType type represents `newsfeed_newsfeed_item_type` API object
type NewsfeedNewsfeedItemType string // Item type

// NewsfeedNewsfeedNote type represents `newsfeed_newsfeed_note` API object
type NewsfeedNewsfeedNote struct {
	Comments int    `json:"comments"` // Comments Number
	Id       int    `json:"id"`       // Note ID
	OwnerId  int    `json:"owner_id"` // integer
	Title    string `json:"title"`    // Note title
}

// NewsfeedNewsfeedPhoto type represents `newsfeed_newsfeed_photo` API object
type NewsfeedNewsfeedPhoto struct {
	AccessKey  string             `json:"access_key"`  // Access key for the photo
	AlbumId    int                `json:"album_id"`    // Album ID
	CanComment BaseBoolInt        `json:"can_comment"` // Information whether current user can comment the photo
	CanRepost  BaseBoolInt        `json:"can_repost"`  // Information whether current user can repost the photo
	Comments   BaseObjectCount    `json:"comments"`
	Date       int                `json:"date"`   // Date when uploaded
	Height     int                `json:"height"` // Original photo height
	Id         int                `json:"id"`     // Photo ID
	Images     []PhotosImage      `json:"images"`
	Lat        json.Number        `json:"lat"` // Latitude
	Likes      BaseLikes          `json:"likes"`
	Long       json.Number        `json:"long"`     // Longitude
	OwnerId    int                `json:"owner_id"` // Photo owner's ID
	PostId     int                `json:"post_id"`  // Post ID
	Sizes      []PhotosPhotoSizes `json:"sizes"`
	Text       string             `json:"text"`    // Photo caption
	UserId     int                `json:"user_id"` // ID of the user who have uploaded the photo
	Width      int                `json:"width"`   // Original photo width
}
