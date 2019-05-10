package newsfeed

import (
	"fmt"
	"gitlab.com/Burmuley/go-vkapi/objects/audio"
	"gitlab.com/Burmuley/go-vkapi/objects/base"
	"gitlab.com/Burmuley/go-vkapi/objects/groups"
	"gitlab.com/Burmuley/go-vkapi/objects/photos"
	"gitlab.com/Burmuley/go-vkapi/objects/videos"
	"gitlab.com/Burmuley/go-vkapi/objects/wall"
)

// CommentsFilters represents `newsfeed_comments_filters` API object
type CommentsFilters string

func (c *CommentsFilters) MarshalJSON() ([]byte, error) {
	switch *c {
	case "post",
		"photo",
		"video",
		"topic",
		"note":
		return []byte(*c), nil
	}

	return []byte{}, fmt.Errorf("value is not in valid range")
}

func (c *CommentsFilters) GetName() string {
	return string(*c)
}

// EventActivity represents `newsfeed_event_activity` API object
type EventActivity struct {
	Address      string                       `json:"address"`       // address of event
	ButtonText   string                       `json:"button_text"`   // text of attach
	Friends      []int                        `json:"friends"`       // array of friends ids
	MemberStatus groups.GroupFullMemberStatus `json:"member_status"` // Current user's member status
	Text         string                       `json:"text"`          // text of attach
	Time         int                          `json:"time"`          // event start time
}

// Filters represents `newsfeed_filters` API object
type Filters string

func (f *Filters) MarshalJSON() ([]byte, error) {
	switch *f {
	case "post",
		"photo",
		"photo_tag",
		"wall_photo",
		"friend",
		"note",
		"audio",
		"video":
		return []byte(*f), nil
	}

	return []byte{}, fmt.Errorf("value is not in valid range")
}

func (f *Filters) GetName() string {
	return string(*f)
}

// IgnoreItemType represents `newsfeed_ignore_item_type` API object
type IgnoreItemType string

func (i *IgnoreItemType) MarshalJSON() ([]byte, error) {
	switch *i {
	case "wall",
		"tag",
		"profilephoto",
		"video",
		"photo",
		"audio":
		return []byte(*i), nil
	}

	return []byte{}, fmt.Errorf("value is not in valid range")
}

func (i *IgnoreItemType) GetName() string {
	switch *i {
	case "wall":
		return "post on the wall"
	case "tag":
		return "tag on a photo"
	case "profilephoto":
		return "profile photo"
	case "video":
		return "video"
	case "photo":
		return "photo"
	case "audio":
		return "audio"
	default:
		return "unknown value"
	}
}

// ItemAudio represents `newsfeed_item_audio` API object
type ItemAudio struct {
	Audio  ItemAudioAudio `json:"audio"`
	PostID int            `json:"post_id"`
}

// ItemAudioAudio represents `newsfeed_item_audio_audio` API object
type ItemAudioAudio struct {
	Count int               `json:"count"` // Audios number
	Items []audio.AudioFull `json:"items"`
}

// ItemDigest represents `newsfeed_item_digest` API object
type ItemDigest struct {
	ButtonText  string          `json:"button_text"`
	FeedID      int             `json:"feed_id"` // id of feed in digest
	Items       []wall.Wallpost `json:"items"`
	MainPostIDs []string        `json:"main_post_ids"`
	Template    string          `json:"template"` // type of digest (can be 'list' or 'grid')
	Title       string          `json:"title"`
	TrackCode   int             `json:"track_code"`
	Type        string          `json:"type"` // type of digest (can be 'digest`)
}

// ItemFriend represents `newsfeed_item_friend` API object
type ItemFriend struct {
	Friends ItemFriendFriends `json:"friends"`
}

// ItemFriendFriend `newsfeed_item_friend_friends` API object
type ItemFriendFriends struct {
	Count int           `json:"count"` // Number of friends has been added
	Items []base.UserID `json:"items"`
}

// ItemNote represents `newsfeed_item_note` API object
type ItemNote struct {
	Notes ItemNoteNotes `json:"notes"`
}

// ItemNoteNotes represents `newsfeed_item_note_notes` API object
type ItemNoteNotes struct {
	Count int  `json:"count"` // Notes number
	Items Note `json:"items"`
}

// ItemPhoto represents `newsfeed_item_photo` API object
type ItemPhoto struct {
	PostID int             `json:"post_id"` // Post ID
	Photos ItemPhotoPhotos `json:"photos"`
}

// ItemPhotoPhotos represents `newsfeed_item_photo_photos` API object
type ItemPhotoPhotos struct {
	Count int     `json:"count"` // Photos number
	Items []Photo `json:"items"`
}

// ItemPhotoTag represents `newsfeed_item_photo_tag` API object
type ItemPhotoTag struct {
	PostID    int                   `json:"post_id"` // Post ID
	PhotoTags ItemPhotoTagPhotoTags `json:"photo_tags"`
}

// ItemPhotoTagPhotoTags represents `newsfeed_item_photo_tag_photo_tags` API object
type ItemPhotoTagPhotoTags struct {
	Count int     `json:"count"` // Tags number
	Items []Photo `json:"items"`
}

// ItemTopic represents `newsfeed_item_topic` API object
type ItemTopic struct {
	Comments base.CommentsInfo `json:"comments"`
	Likes    base.LikesInfo    `json:"likes"`
	PostID   int               `json:"post_id"` // Topic post ID
	Text     string            `json:"text"`    // Post text
}

// ItemVideo represents `newsfeed_item_video` API object
type ItemVideo struct {
	Video ItemVideoVideo `json:"video"`
}

// ItemVideoVideo represents `newsfeed_item_video_video` API object
type ItemVideoVideo struct {
	Count int          `json:"count"` // Tags number
	Items videos.Video `json:"items"`
}

// ItemWallpost represents `newsfeed_item_wallpost` API object
type ItemWallpost struct {
	Activity    EventActivity           `json:"activity"`
	Attachments wall.WallpostAttachment `json:"attachments"`
	Comments    base.CommentsInfo       `json:"comments"`
	CopyHistory []wall.Wallpost         `json:"copy_history"`
	Geo         base.Geo                `json:"geo"`
	Likes       base.LikesInfo          `json:"likes"`
	PostID      int                     `json:"post_id"`
	PostSource  wall.PostSource         `json:"post_source"`
	PostType    ItemWallpostType        `json:"post_type"`
	Reposts     base.RepostsInfo        `json:"reposts"`
	Text        string                  `json:"text"` // Post text
}

// ItemWallpostType represents `newsfeed_item_wallpost_type` API object
type ItemWallpostType string

func (i *ItemWallpostType) MarshalJSON() ([]byte, error) {
	switch *i {
	case "post",
		"copy",
		"reply":
		return []byte(*i), nil
	}

	return []byte{}, fmt.Errorf("value is not in valid range")
}

func (i *ItemWallpostType) GetName() string {
	return string(*i)
}

// List represents `newsfeed_list` API object
type List struct {
	ID    int    `json:"id"`    // List ID
	Title string `json:"title"` // List title
}

// ListFull represents `newsfeed_list_full` API object
type ListFull struct {
	List
	NoReposts base.BoolInt `json:"no_reposts"` // Information whether reposts hiding is enabled
	SourceIDS []int        `json:"source_ids"` // Users and communities IDs
}

// Item represents `newsfeed_newsfeed_item` API object
// TODO: Finish implementation!
// WARNING: NOT IMPLEMENTED YET
type Item struct {
	ItemWallpost `json:",omitempty"`
	ItemPhoto    `json:",omitempty"`
	ItemPhotoTag `json:",omitempty"`
	ItemFriend   `json:",omitempty"`
	ItemNote     `json:",omitempty"`
	ItemAudio    `json:",omitempty"`
	ItemVideo    `json:",omitempty"`
	ItemTopic    `json:",omitempty"`
	ItemDigest   `json:",omitempty"`
	Type         ItemType `json:"type"`
	SourceID     int      `json:"source_id"` // Item source ID
	Date         int      `json:"date"`      // Date when item has been added in Unixtime
}

// ItemType represents `newsfeed_newsfeed_item_type` API object
type ItemType string

func (i *ItemType) MarshalJSON() ([]byte, error) {
	switch *i {
	case "post",
		"photo",
		"photo_tag",
		"wall_photo",
		"friend",
		"note",
		"audio",
		"video",
		"topic":
		return []byte(*i), nil
	}

	return []byte{}, fmt.Errorf("value is not in valid range")
}

func (i *ItemType) GetName() string {
	return string(*i)
}

// Note represents `newsfeed_newsfeed_note` API object
type Note struct {
	Comments int    `json:"comments"` // Comments Number
	ID       int    `json:"id"`       // Note ID
	OwnerID  int    `json:"owner_id"` // Owner ID
	Title    string `json:"title"`    // Note title
}

// Photo represents `newsfeed_newsfeed_photo` API object
type Photo struct {
	photos.Photo
	Likes      base.Likes       `json:"likes"`
	Comments   base.ObjectCount `json:"comments"`
	CanComment base.BoolInt     `json:"can_comment"` // Information whether current user can comment the photo
	CanRepost  base.BoolInt     `json:"can_repost"`  // Information whether current user can repost the photo
}
