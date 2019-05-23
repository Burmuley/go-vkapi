package objects

// CommentsFilters represents `newsfeed_comments_filters` API object
type NewsfeedCommentsFilters string

func (c *NewsfeedCommentsFilters) MarshalJSON() ([]byte, error) {
	return GetStringFromRange(string(*c), "post", "photo", "video", "topic", "note")
}

func (c *NewsfeedCommentsFilters) String() string {
	return string(*c)
}

// EventActivity represents `newsfeed_event_activity` API object
type NewsfeedEventActivity struct {
	Address      string                      `json:"address"`       // address of event
	ButtonText   string                      `json:"button_text"`   // text of attach
	Friends      []int                       `json:"friends"`       // array of friends ids
	MemberStatus GroupsGroupFullMemberStatus `json:"member_status"` // Current user's member status
	Text         string                      `json:"text"`          // text of attach
	Time         int                         `json:"time"`          // event start time
}

// Filters represents `newsfeed_filters` API object
type NewsfeedFilters string

func (f *NewsfeedFilters) MarshalJSON() ([]byte, error) {
	return GetStringFromRange(string(*f), "post", "photo", "photo_tag", "wall_photo", "friend",
		"note", "audio", "video")
}

func (f *NewsfeedFilters) String() string {
	return string(*f)
}

// IgnoreItemType represents `newsfeed_ignore_item_type` API object
type NewsfeedIgnoreItemType string

func (i *NewsfeedIgnoreItemType) MarshalJSON() ([]byte, error) {
	return GetStringFromRange(string(*i), "wall", "tag", "profilephoto", "video", "photo", "audio")
}

func (i *NewsfeedIgnoreItemType) String() string {
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
type NewsfeedItemAudio struct {
	Audio  NewsfeedItemAudioAudio `json:"audio"`
	PostID int                    `json:"post_id"`
}

// ItemAudioAudio represents `newsfeed_item_audio_audio` API object
type NewsfeedItemAudioAudio struct {
	Count int              `json:"count"` // Audios number
	Items []AudioAudioFull `json:"items"`
}

// ItemDigest represents `newsfeed_item_digest` API object
type NewsfeedItemDigest struct {
	ButtonText  string         `json:"button_text"`
	FeedID      int            `json:"feed_id"` // id of feed in digest
	Items       []WallWallpost `json:"items"`
	MainPostIDs []string       `json:"main_post_ids"`
	Template    string         `json:"template"` // type Newsfeedof digest (can be 'list' or 'grid')
	Title       string         `json:"title"`
	TrackCode   int            `json:"track_code"`
	Type        string         `json:"type"` // type Newsfeedof digest (can be 'digest`)
}

// ItemFriend represents `newsfeed_item_friend` API object
type NewsfeedItemFriend struct {
	Friends NewsfeedItemFriendFriends `json:"friends"`
}

// ItemFriendFriend `newsfeed_item_friend_friends` API object
type NewsfeedItemFriendFriends struct {
	Count int          `json:"count"` // Number of friends has been added
	Items []BaseUserID `json:"items"`
}

// ItemNote represents `newsfeed_item_note` API object
type NewsfeedItemNote struct {
	Notes NewsfeedItemNoteNotes `json:"notes"`
}

// ItemNoteNotes represents `newsfeed_item_note_notes` API object
type NewsfeedItemNoteNotes struct {
	Count int          `json:"count"` // Notes number
	Items NewsfeedNote `json:"items"`
}

// ItemPhoto represents `newsfeed_item_photo` API object
type NewsfeedItemPhoto struct {
	PostID int                     `json:"post_id"` // Post ID
	Photos NewsfeedItemPhotoPhotos `json:"photos"`
}

// ItemPhotoPhotos represents `newsfeed_item_photo_photos` API object
type NewsfeedItemPhotoPhotos struct {
	Count int             `json:"count"` // Photos number
	Items []NewsfeedPhoto `json:"items"`
}

// ItemPhotoTag represents `newsfeed_item_photo_tag` API object
type NewsfeedItemPhotoTag struct {
	PostID    int                           `json:"post_id"` // Post ID
	PhotoTags NewsfeedItemPhotoTagPhotoTags `json:"photo_tags"`
}

// ItemPhotoTagPhotoTags represents `newsfeed_item_photo_tag_photo_tags` API object
type NewsfeedItemPhotoTagPhotoTags struct {
	Count int             `json:"count"` // Tags number
	Items []NewsfeedPhoto `json:"items"`
}

// ItemTopic represents `newsfeed_item_topic` API object
type NewsfeedItemTopic struct {
	Comments BaseCommentsInfo `json:"comments"`
	Likes    BaseLikesInfo    `json:"likes"`
	PostID   int              `json:"post_id"` // Topic post ID
	Text     string           `json:"text"`    // Post text
}

// ItemVideo represents `newsfeed_item_video` API object
type NewsfeedItemVideo struct {
	Video NewsfeedItemVideoVideo `json:"video"`
}

// ItemVideoVideo represents `newsfeed_item_video_video` API object
type NewsfeedItemVideoVideo struct {
	Count int         `json:"count"` // Tags number
	Items VideosVideo `json:"items"`
}

// ItemWallpost represents `newsfeed_item_wallpost` API object
type NewsfeedItemWallpost struct {
	Activity    NewsfeedEventActivity    `json:"activity"`
	Attachments WallWallpostAttachment   `json:"attachments"`
	Comments    BaseCommentsInfo         `json:"comments"`
	CopyHistory []WallWallpost           `json:"copy_history"`
	Geo         BaseGeo                  `json:"geo"`
	Likes       BaseLikesInfo            `json:"likes"`
	PostID      int                      `json:"post_id"`
	PostSource  WallPostSource           `json:"post_source"`
	PostType    NewsfeedItemWallpostType `json:"post_type"`
	Reposts     BaseRepostsInfo          `json:"reposts"`
	Text        string                   `json:"text"` // Post text
}

// ItemWallpostType represents `newsfeed_item_wallpost_type` API object
type NewsfeedItemWallpostType string

func (i *NewsfeedItemWallpostType) MarshalJSON() ([]byte, error) {
	return GetStringFromRange(string(*i), "post", "copy", "reply")
}

func (i *NewsfeedItemWallpostType) String() string {
	return string(*i)
}

// List represents `newsfeed_list` API object
type NewsfeedList struct {
	ID    int    `json:"id"`    // List ID
	Title string `json:"title"` // List title
}

// ListFull represents `newsfeed_list_full` API object
type NewsfeedListFull struct {
	*NewsfeedList
	NoReposts BaseBoolInt `json:"no_reposts"` // Information whether reposts hiding is enabled
	SourceIDS []int       `json:"source_ids"` // Users and communities IDs
}

// Item represents `newsfeed_newsfeed_item` API object
// TODO: Finish implementation!
// WARNING: NOT IMPLEMENTED YET
type NewsfeedItem struct {
	NewsfeedItemWallpost `json:",omitempty"`
	NewsfeedItemPhoto    `json:",omitempty"`
	NewsfeedItemPhotoTag `json:",omitempty"`
	NewsfeedItemFriend   `json:",omitempty"`
	NewsfeedItemNote     `json:",omitempty"`
	NewsfeedItemAudio    `json:",omitempty"`
	NewsfeedItemVideo    `json:",omitempty"`
	NewsfeedItemTopic    `json:",omitempty"`
	NewsfeedItemDigest   `json:",omitempty"`
	Type                 NewsfeedItemType `json:"type"`
	SourceID             int              `json:"source_id"` // Item source ID
	Date                 int              `json:"date"`      // Date when item has been added in Unixtime
}

// ItemType represents `newsfeed_newsfeed_item_type` API object
type NewsfeedItemType string

func (i *NewsfeedItemType) MarshalJSON() ([]byte, error) {
	return GetStringFromRange(string(*i), "post",
		"photo", "photo_tag", "wall_photo", "friend", "note", "audio",
		"video", "topic")
}

func (i *NewsfeedItemType) String() string {
	return string(*i)
}

// Note represents `newsfeed_newsfeed_note` API object
type NewsfeedNote struct {
	Comments int    `json:"comments"` // Comments Number
	ID       int    `json:"id"`       // Note ID
	OwnerID  int    `json:"owner_id"` // Owner ID
	Title    string `json:"title"`    // Note title
}

// Photo represents `newsfeed_newsfeed_photo` API object
type NewsfeedPhoto struct {
	*PhotosPhoto
	Likes      BaseLikes       `json:"likes"`
	Comments   BaseObjectCount `json:"comments"`
	CanComment BaseBoolInt     `json:"can_comment"` // Information whether current user can comment the photo
	CanRepost  BaseBoolInt     `json:"can_repost"`  // Information whether current user can repost the photo
}
