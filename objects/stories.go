package objects

/////////////////////////////////////////////////////////////
// Stories related API objects	                           //
/////////////////////////////////////////////////////////////

// Replies represents `stories_replies` API object
type StoriesReplies struct {
	Count int `json:"count"` // Replies number.
	New   int `json:"new"`   // New replies number.
}

// Story represents `stories_story` API object
type StoriesStory struct {
	AccessKey            string           `json:"access_key"`  // Access key for private object.
	CanComment           BaseBoolInt      `json:"can_comment"` // Information whether current user can comment the story (0 - no, 1 - yes).
	CanReply             BaseBoolInt      `json:"can_reply"`   // Information whether current user can reply to the story (0 - no, 1 - yes).
	CanSee               BaseBoolInt      `json:"can_see"`     // Information whether current user can see the story (0 - no, 1 - yes).
	CanShare             BaseBoolInt      `json:"can_share"`   // Information whether current user can share the story (0 - no, 1 - yes).
	Date                 int              `json:"date"`        // Date when story has been added in Unixtime.
	ID                   int              `json:"id"`          // Story ID.
	Deleted              bool             `json:"is_deleted"`  // Information whether the story is deleted (false - no, true - yes).
	Expired              bool             `json:"is_expired"`  // Information whether the story is expired (false - no, true - yes).
	Link                 StoriesLink      `json:"link"`
	OwnerID              int              `json:"owner_id"` // Story owner's ID.
	ParentStory          *StoriesStory    `json:"parent_story"`
	ParentStoryAccessKey string           `json:"parent_story_access_key"` // Access key for private object.
	ParentStoryID        int              `json:"parent_story_id"`         // Parent story ID.
	ParentStoryOwnerID   int              `json:"parent_story_owner_id"`   // Parent story owner's ID.
	Photo                PhotosPhoto      `json:"photo"`
	Replies              []StoriesReplies `json:"replies"` // Replies to current story.
	Seen                 BaseBoolInt      `json:"seen"`    // Information whether current user has seen the story or not (0 - no, 1 - yes).
	Type                 StoriesType      `json:"type"`
	Video                StoriesVideo     `json:"video"`
	Views                int              `json:"views"` // Views number.
}

// Link represents `stories_story_link` API object
type StoriesLink struct {
	Text string `json:"text"` // Link text
	Url  string `json:"url"`  // Link URL
}

// Stats represents `stories_story_stats` API object
type StoriesStats struct {
	Answer      StoriesStatsStat `json:"answer"`
	Bans        StoriesStatsStat `json:"bans"`
	OpenLink    StoriesStatsStat `json:"open_link"`
	Replies     StoriesStatsStat `json:"replies"`
	Shares      StoriesStatsStat `json:"shares"`
	Subscribers StoriesStatsStat `json:"subscribers"`
	Views       StoriesStatsStat `json:"views"`
}

// StatsStat represents `stories_story_stats_stat` API object
type StoriesStatsStat struct {
	Count int               `json:"count"` // Stat value
	State StoriesStatsState `json:"state"`
}

// StatsState represents `stories_story_stats_state` API object
type StoriesStatsState string

func (s *StoriesStatsState) MarshalJSON() ([]byte, error) {
	return GetStringFromRange(string(*s), "on", "off", "hidden")
}

func (s *StoriesStatsState) String() string {
	return string(*s)
}

// Type represents `stories_story_type` API object
type StoriesType string

func (t *StoriesType) MarshalJSON() ([]byte, error) {
	return GetStringFromRange(string(*t), "photo", "video")
}

func (t *StoriesType) String() string {
	return string(*t)
}

// Video represents `stories_story_video` API object
type StoriesVideo struct {
	*VideosVideo
	Private BaseBoolInt `json:"private"` // Information whether story is private (0 - no, 1 - yes).
}

// UploadLinkText represents `stories_upload_link_text` API object
type StoriesUploadLinkText string

func (u *StoriesUploadLinkText) MarshalJSON() ([]byte, error) {
	return GetStringFromRange(string(*u), "to_store", "vote", "more", "book", "order", "enroll",
		"fill", "signup", "buy", "ticket", "write", "open", "learn_more", "view", "go_to", "contact", "watch",
		"play", "install", "read")
}

func (u *StoriesUploadLinkText) String() string {
	return string(*u)
}
