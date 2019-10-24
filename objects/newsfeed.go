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
// `newsfeed` group of objects
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

// NewsfeedCommentsFilters type represents `newsfeed_comments_filters` API object
type NewsfeedCommentsFilters string

// NewsfeedEventActivity type represents `newsfeed_event_activity` API object
type NewsfeedEventActivity struct {
	Address      string                       `json:"address"`       // address of event
	ButtonText   string                       `json:"button_text"`   // text of attach
	Friends      []int                        `json:"friends"`       // array of friends ids
	MemberStatus *GroupsGroupFullMemberStatus `json:"member_status"` // Current user's member status
	Text         string                       `json:"text"`          // text of attach
	Time         int                          `json:"time"`          // event start time
}

// NewsfeedFilters type represents `newsfeed_filters` API object
type NewsfeedFilters string

// NewsfeedIgnoreItemType type represents `newsfeed_ignore_item_type` API object
type NewsfeedIgnoreItemType string

// NewsfeedItemAudio type represents `newsfeed_item_audio` API object
type NewsfeedItemAudio struct {
	Newsfeeditembase *NewsfeedItemBase       `json:"Newsfeeditembase"`
	Audio            *NewsfeedItemAudioAudio `json:"audio"`
	PostId           int                     `json:"post_id"` // Post ID
}

// NewsfeedItemAudioAudio type represents `newsfeed_item_audio_audio` API object
type NewsfeedItemAudioAudio struct {
	Count int           `json:"count"` // Audios number
	Items []*AudioAudio `json:"items"`
}

// NewsfeedItemBase type represents `newsfeed_item_base` API object
type NewsfeedItemBase struct {
	Date     int                       `json:"date"`      // Date when item has been added in Unixtime
	SourceId int                       `json:"source_id"` // Item source ID
	Type     *NewsfeedNewsfeedItemType `json:"type"`
}

// NewsfeedItemDigest type represents `newsfeed_item_digest` API object
type NewsfeedItemDigest struct {
	Newsfeeditembase *NewsfeedItemBase `json:"Newsfeeditembase"`
	ButtonText       string            `json:"button_text"`
	FeedId           string            `json:"feed_id"` // id of feed in digest
	Items            []*WallWallpost   `json:"items"`
	MainPostIds      []string          `json:"main_post_ids"`
	Template         string            `json:"template"` // type of digest
	Title            string            `json:"title"`
	TrackCode        string            `json:"track_code"`
}

// NewsfeedItemFriend type represents `newsfeed_item_friend` API object
type NewsfeedItemFriend struct {
	Newsfeeditembase *NewsfeedItemBase          `json:"Newsfeeditembase"`
	Friends          *NewsfeedItemFriendFriends `json:"friends"`
}

// NewsfeedItemFriendFriends type represents `newsfeed_item_friend_friends` API object
type NewsfeedItemFriendFriends struct {
	Count int           `json:"count"` // Number of friends has been added
	Items []*BaseUserId `json:"items"`
}

// NewsfeedItemNote type represents `newsfeed_item_note` API object
type NewsfeedItemNote struct {
	Newsfeeditembase *NewsfeedItemBase      `json:"Newsfeeditembase"`
	Notes            *NewsfeedItemNoteNotes `json:"notes"`
}

// NewsfeedItemNoteNotes type represents `newsfeed_item_note_notes` API object
type NewsfeedItemNoteNotes struct {
	Count int                     `json:"count"` // Notes number
	Items []*NewsfeedNewsfeedNote `json:"items"`
}

// NewsfeedItemPhoto type represents `newsfeed_item_photo` API object
type NewsfeedItemPhoto struct {
	Newsfeeditembase *NewsfeedItemBase        `json:"Newsfeeditembase"`
	Photos           *NewsfeedItemPhotoPhotos `json:"photos"`
	PostId           int                      `json:"post_id"` // Post ID
}

// NewsfeedItemPhotoPhotos type represents `newsfeed_item_photo_photos` API object
type NewsfeedItemPhotoPhotos struct {
	Count int                      `json:"count"` // Photos number
	Items []*NewsfeedNewsfeedPhoto `json:"items"`
}

// NewsfeedItemPhotoTag type represents `newsfeed_item_photo_tag` API object
type NewsfeedItemPhotoTag struct {
	Newsfeeditembase *NewsfeedItemBase              `json:"Newsfeeditembase"`
	PhotoTags        *NewsfeedItemPhotoTagPhotoTags `json:"photo_tags"`
	PostId           int                            `json:"post_id"` // Post ID
}

// NewsfeedItemPhotoTagPhotoTags type represents `newsfeed_item_photo_tag_photo_tags` API object
type NewsfeedItemPhotoTagPhotoTags struct {
	Count int                      `json:"count"` // Tags number
	Items []*NewsfeedNewsfeedPhoto `json:"items"`
}

// NewsfeedItemStoriesBlock type represents `newsfeed_item_stories_block` API object
type NewsfeedItemStoriesBlock struct {
	Newsfeeditembase *NewsfeedItemBase `json:"Newsfeeditembase"`
	BlockType        string            `json:"block_type"`
	Stories          []*StoriesStory   `json:"stories"`
	Title            string            `json:"title"`
	TrackCode        string            `json:"track_code"`
}

// NewsfeedItemTopic type represents `newsfeed_item_topic` API object
type NewsfeedItemTopic struct {
	Newsfeeditembase *NewsfeedItemBase `json:"Newsfeeditembase"`
	Comments         *BaseCommentsInfo `json:"comments"`
	Likes            *BaseLikesInfo    `json:"likes"`
	PostId           int               `json:"post_id"` // Topic post ID
	Text             string            `json:"text"`    // Post text
}

// NewsfeedItemVideo type represents `newsfeed_item_video` API object
type NewsfeedItemVideo struct {
	Newsfeeditembase *NewsfeedItemBase       `json:"Newsfeeditembase"`
	Video            *NewsfeedItemVideoVideo `json:"video"`
}

// NewsfeedItemVideoVideo type represents `newsfeed_item_video_video` API object
type NewsfeedItemVideoVideo struct {
	Count int           `json:"count"` // Tags number
	Items []*VideoVideo `json:"items"`
}

// NewsfeedItemWallpost type represents `newsfeed_item_wallpost` API object
type NewsfeedItemWallpost struct {
	Newsfeeditembase *NewsfeedItemBase         `json:"Newsfeeditembase"`
	Activity         *NewsfeedEventActivity    `json:"activity"`
	Attachments      []*WallWallpostAttachment `json:"attachments"`
	Comments         *BaseCommentsInfo         `json:"comments"`
	CopyHistory      []*WallWallpost           `json:"copy_history"`
	Geo              *BaseGeo                  `json:"geo"`
	Likes            *BaseLikesInfo            `json:"likes"`
	PostId           int                       `json:"post_id"` // Post ID
	PostSource       *WallPostSource           `json:"post_source"`
	PostType         *NewsfeedItemWallpostType `json:"post_type"`
	Reposts          *BaseRepostsInfo          `json:"reposts"`
	Text             string                    `json:"text"` // Post text
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
	Newsfeedlist *NewsfeedList `json:"Newsfeedlist"`
	NoReposts    *BaseBoolInt  `json:"no_reposts"` // Information whether reposts hiding is enabled
	SourceIds    []int         `json:"source_ids"`
}

// NewsfeedNewsfeedItem type represents `newsfeed_newsfeed_item` API object
type NewsfeedNewsfeedItem struct {
	Newsfeeditemaudio        *NewsfeedItemAudio        `json:"Newsfeeditemaudio"`
	Newsfeeditemdigest       *NewsfeedItemDigest       `json:"Newsfeeditemdigest"`
	Newsfeeditemfriend       *NewsfeedItemFriend       `json:"Newsfeeditemfriend"`
	Newsfeeditemnote         *NewsfeedItemNote         `json:"Newsfeeditemnote"`
	Newsfeeditemphoto        *NewsfeedItemPhoto        `json:"Newsfeeditemphoto"`
	Newsfeeditemphototag     *NewsfeedItemPhotoTag     `json:"Newsfeeditemphototag"`
	Newsfeeditemstoriesblock *NewsfeedItemStoriesBlock `json:"Newsfeeditemstoriesblock"`
	Newsfeeditemtopic        *NewsfeedItemTopic        `json:"Newsfeeditemtopic"`
	Newsfeeditemvideo        *NewsfeedItemVideo        `json:"Newsfeeditemvideo"`
	Newsfeeditemwallpost     *NewsfeedItemWallpost     `json:"Newsfeeditemwallpost"`
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
	Photosphoto *PhotosPhoto     `json:"Photosphoto"`
	CanComment  *BaseBoolInt     `json:"can_comment"` // Information whether current user can comment the photo
	CanRepost   *BaseBoolInt     `json:"can_repost"`  // Information whether current user can repost the photo
	Comments    *BaseObjectCount `json:"comments"`
	Likes       *BaseLikes       `json:"likes"`
}
