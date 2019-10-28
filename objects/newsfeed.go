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
	NewsfeedItemBase
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
	NewsfeedItemBase
}

// NewsfeedItemFriend type represents `newsfeed_item_friend` API object
type NewsfeedItemFriend struct {
	NewsfeedItemBase
}

// NewsfeedItemFriendFriends type represents `newsfeed_item_friend_friends` API object
type NewsfeedItemFriendFriends struct {
	Count int          `json:"count"` // Number of friends has been added
	Items []BaseUserId `json:"items"`
}

// NewsfeedItemNote type represents `newsfeed_item_note` API object
type NewsfeedItemNote struct {
	NewsfeedItemBase
}

// NewsfeedItemNoteNotes type represents `newsfeed_item_note_notes` API object
type NewsfeedItemNoteNotes struct {
	Count int                    `json:"count"` // Notes number
	Items []NewsfeedNewsfeedNote `json:"items"`
}

// NewsfeedItemPhoto type represents `newsfeed_item_photo` API object
type NewsfeedItemPhoto struct {
	NewsfeedItemBase
}

// NewsfeedItemPhotoPhotos type represents `newsfeed_item_photo_photos` API object
type NewsfeedItemPhotoPhotos struct {
	Count int                     `json:"count"` // Photos number
	Items []NewsfeedNewsfeedPhoto `json:"items"`
}

// NewsfeedItemPhotoTag type represents `newsfeed_item_photo_tag` API object
type NewsfeedItemPhotoTag struct {
	NewsfeedItemBase
}

// NewsfeedItemPhotoTagPhotoTags type represents `newsfeed_item_photo_tag_photo_tags` API object
type NewsfeedItemPhotoTagPhotoTags struct {
	Count int                     `json:"count"` // Tags number
	Items []NewsfeedNewsfeedPhoto `json:"items"`
}

// NewsfeedItemStoriesBlock type represents `newsfeed_item_stories_block` API object
type NewsfeedItemStoriesBlock struct {
	NewsfeedItemBase
}

// NewsfeedItemTopic type represents `newsfeed_item_topic` API object
type NewsfeedItemTopic struct {
	NewsfeedItemBase
}

// NewsfeedItemVideo type represents `newsfeed_item_video` API object
type NewsfeedItemVideo struct {
	NewsfeedItemBase
}

// NewsfeedItemVideoVideo type represents `newsfeed_item_video_video` API object
type NewsfeedItemVideoVideo struct {
	Count int          `json:"count"` // Tags number
	Items []VideoVideo `json:"items"`
}

// NewsfeedItemWallpost type represents `newsfeed_item_wallpost` API object
type NewsfeedItemWallpost struct {
	NewsfeedItemBase
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
	NewsfeedList
}

// NewsfeedNewsfeedItem type represents `newsfeed_newsfeed_item` API object
type NewsfeedNewsfeedItem struct {
	NewsfeedItemWallpost
	NewsfeedItemPhoto
	NewsfeedItemPhotoTag
	NewsfeedItemFriend
	NewsfeedItemNote
	NewsfeedItemAudio
	NewsfeedItemVideo
	NewsfeedItemTopic
	NewsfeedItemDigest
	NewsfeedItemStoriesBlock
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
	PhotosPhoto
}
