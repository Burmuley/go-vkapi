package stories

import (
	"fmt"
	"gitlab.com/Burmuley/go-vkapi/objects/base"
	"gitlab.com/Burmuley/go-vkapi/objects/photos"
	"gitlab.com/Burmuley/go-vkapi/objects/videos"
)

/////////////////////////////////////////////////////////////
// Stories related API objects	                           //
/////////////////////////////////////////////////////////////

// Replies represents `stories_replies` API object
type Replies struct {
	Count int `json:"count"` // Replies number.
	New   int `json:"new"`   // New replies number.
}

// Story represents `stories_story` API object
type Story struct {
	AccessKey            string       `json:"access_key"`  // Access key for private object.
	CanComment           base.BoolInt `json:"can_comment"` // Information whether current user can comment the story (0 - no, 1 - yes).
	CanReply             base.BoolInt `json:"can_reply"`   // Information whether current user can reply to the story (0 - no, 1 - yes).
	CanSee               base.BoolInt `json:"can_see"`     // Information whether current user can see the story (0 - no, 1 - yes).
	CanShare             base.BoolInt `json:"can_share"`   // Information whether current user can share the story (0 - no, 1 - yes).
	Date                 int          `json:"date"`        // Date when story has been added in Unixtime.
	ID                   int          `json:"id"`          // Story ID.
	Deleted              bool         `json:"is_deleted"`  // Information whether the story is deleted (false - no, true - yes).
	Expired              bool         `json:"is_expired"`  // Information whether the story is expired (false - no, true - yes).
	Link                 Link         `json:"link"`
	OwnerID              int          `json:"owner_id"` // Story owner's ID.
	ParentStory          Story        `json:"parent_story"`
	ParentStoryAccessKey string       `json:"parent_story_access_key"` // Access key for private object.
	ParentStoryID        int          `json:"parent_story_id"`         // Parent story ID.
	ParentStoryOwnerID   int          `json:"parent_story_owner_id"`   // Parent story owner's ID.
	Photo                photos.Photo `json:"photo"`
	Replies              []Replies    `json:"replies"` // Replies to current story.
	Seen                 base.BoolInt `json:"seen"`    // Information whether current user has seen the story or not (0 - no, 1 - yes).
	Type                 Type         `json:"type"`
	Video                Video        `json:"video"`
	Views                int          `json:"views"` // Views number.
}

// Link represents `stories_story_link` API object
type Link struct {
	Text string `json:"text"` // Link text
	Url  string `json:"url"`  // Link URL
}

// Stats represents `stories_story_stats` API object
type Stats struct {
	Answer      StatsStat `json:"answer"`
	Bans        StatsStat `json:"bans"`
	OpenLink    StatsStat `json:"open_link"`
	Replies     StatsStat `json:"replies"`
	Shares      StatsStat `json:"shares"`
	Subscribers StatsStat `json:"subscribers"`
	Views       StatsStat `json:"views"`
}

// StatsStat represents `stories_story_stats_stat` API object
type StatsStat struct {
	Count int        `json:"count"` // Stat value
	State StatsState `json:"state"`
}

// StatsState represents `stories_story_stats_state` API object
type StatsState string

func (s *StatsState) MarshalJSON() ([]byte, error) {
	switch *s {
	case "on", "off", "hidden":
		return []byte(*s), nil
	}

	return []byte{}, fmt.Errorf("value is not in valid range")
}

func (s *StatsState) GetName() string {
	return string(*s)
}

// Type represents `stories_story_type` API object
type Type string

func (t *Type) MarshalJSON() ([]byte, error) {
	switch *t {
	case "photo", "video":
		return []byte(*t), nil
	}

	return []byte{}, fmt.Errorf("value is not in valid range")
}

func (t *Type) GetName() string {
	return string(*t)
}

// Video represents `stories_story_video` API object
type Video struct {
	videos.Video
	Private base.BoolInt `json:"private"` // Information whether story is private (0 - no, 1 - yes).
}

// UploadLinkText represents `stories_upload_link_text` API object
type UploadLinkText string

func (u *UploadLinkText) MarshalJSON() ([]byte, error) {
	switch *u {
	case "to_store",
		"vote",
		"more",
		"book",
		"order",
		"enroll",
		"fill",
		"signup",
		"buy",
		"ticket",
		"write",
		"open",
		"learn_more",
		"view",
		"go_to",
		"contact",
		"watch",
		"play",
		"install",
		"read":
		return []byte(*u), nil
	}

	return []byte{}, fmt.Errorf("value is not in valid range")
}

func (u *UploadLinkText) GetName() string {
	return string(*u)
}
