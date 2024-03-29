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

package objects

/////////////////////////////////////////////////////////////////////////////////////////////////////////////
// `stories` group of objects
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

// StoriesPromoBlock type represents `stories_promo_block` API object
type StoriesPromoBlock struct {
	Name        string `json:"name"`         // Promo story title
	NotAnimated bool   `json:"not_animated"` // Hide animation for promo story
	Photo100    string `json:"photo_100"`    // RL of square photo of the story with 100 pixels in width
	Photo50     string `json:"photo_50"`     // RL of square photo of the story with 50 pixels in width
}

// StoriesReplies type represents `stories_replies` API object
type StoriesReplies struct {
	Count int `json:"count"` // Replies number.
	New   int `json:"new"`   // New replies number.
}

// StoriesStory type represents `stories_story` API object
type StoriesStory struct {
	AccessKey            string            `json:"access_key"`        // Access key for private object.
	CanAsk               BaseBoolInt       `json:"can_ask"`           // Information whether story has question sticker and current user can send question to the author
	CanAskAnonymous      BaseBoolInt       `json:"can_ask_anonymous"` // Information whether story has question sticker and current user can send anonymous question to the author
	CanComment           BaseBoolInt       `json:"can_comment"`       // Information whether current user can comment the story (0 - no, 1 - yes).
	CanReply             BaseBoolInt       `json:"can_reply"`         // Information whether current user can reply to the story (0 - no, 1 - yes).
	CanSee               BaseBoolInt       `json:"can_see"`           // Information whether current user can see the story (0 - no, 1 - yes).
	CanShare             BaseBoolInt       `json:"can_share"`         // Information whether current user can share the story (0 - no, 1 - yes).
	Date                 int               `json:"date"`              // Date when story has been added in Unixtime.
	ExpiresAt            int               `json:"expires_at"`        // Story expiration time. Unixtime.
	Id                   int               `json:"id"`                // Story ID.
	IsDeleted            bool              `json:"is_deleted"`        // Information whether the story is deleted (false - no, true - yes).
	IsExpired            bool              `json:"is_expired"`        // Information whether the story is expired (false - no, true - yes).
	IsRestricted         bool              `json:"is_restricted"`     // Does author have stories privacy restrictions
	Link                 StoriesStoryLink  `json:"link"`
	NeedMute             bool              `json:"need_mute"` // Does video need to be muted
	NoSound              bool              `json:"no_sound"`  // Is video without sound
	OwnerId              int               `json:"owner_id"`  // Story owner's ID.
	ParentStory          *StoriesStory     `json:"parent_story"`
	ParentStoryAccessKey string            `json:"parent_story_access_key"` // Access key for private object.
	ParentStoryId        int               `json:"parent_story_id"`         // Parent story ID.
	ParentStoryOwnerId   int               `json:"parent_story_owner_id"`   // Parent story owner's ID.
	Photo                PhotosPhoto       `json:"photo"`
	Replies              StoriesReplies    `json:"replies"` // Replies to current story.
	Seen                 BaseBoolInt       `json:"seen"`    // Information whether current user has seen the story or not (0 - no, 1 - yes).
	Type                 StoriesStoryType  `json:"type"`
	Video                StoriesStoryVideo `json:"video"`
	Views                int               `json:"views"` // Views number.
}

// StoriesStoryLink type represents `stories_story_link` API object
type StoriesStoryLink struct {
	Text string `json:"text"` // Link text
	Url  string `json:"url"`  // Link URL
}

// StoriesStoryStats type represents `stories_story_stats` API object
type StoriesStoryStats struct {
	Answer      StoriesStoryStatsStat `json:"answer"`
	Bans        StoriesStoryStatsStat `json:"bans"`
	OpenLink    StoriesStoryStatsStat `json:"open_link"`
	Replies     StoriesStoryStatsStat `json:"replies"`
	Shares      StoriesStoryStatsStat `json:"shares"`
	Subscribers StoriesStoryStatsStat `json:"subscribers"`
	Views       StoriesStoryStatsStat `json:"views"`
}

// StoriesStoryStatsStat type represents `stories_story_stats_stat` API object
type StoriesStoryStatsStat struct {
	Count int                    `json:"count"` // Stat value
	State StoriesStoryStatsState `json:"state"`
}

// StoriesStoryStatsState type represents `stories_story_stats_state` API object
type StoriesStoryStatsState string // Statistic state

// StoriesStoryType type represents `stories_story_type` API object
type StoriesStoryType string // Story type.

// StoriesStoryVideo type represents `stories_story_video` API object
type StoriesStoryVideo struct {
	AccessKey   string             `json:"access_key"`  // Video access key
	AddingDate  int                `json:"adding_date"` // Date when the video has been added in Unixtime
	CanAdd      BaseBoolInt        `json:"can_add"`     // Information whether current user can add the video
	CanComment  BaseBoolInt        `json:"can_comment"` // Information whether current user can comment the video
	CanEdit     BaseBoolInt        `json:"can_edit"`    // Information whether current user can edit the video
	CanLike     BaseBoolInt        `json:"can_like"`    // Information whether current user can like the video
	CanRepost   BaseBoolInt        `json:"can_repost"`  // Information whether current user can repost this video
	Comments    int                `json:"comments"`    // Number of comments
	Date        int                `json:"date"`        // Date when video has been uploaded in Unixtime
	Description string             `json:"description"` // Video description
	Duration    int                `json:"duration"`    // Video duration in seconds
	Files       VideoVideoFiles    `json:"files"`
	FirstFrame  []VideoVideoImage  `json:"first_frame"`
	Height      int                `json:"height"` // Video height
	Id          int                `json:"id"`     // Video ID
	Image       []VideoVideoImage  `json:"image"`
	IsFavorite  bool               `json:"is_favorite"`
	IsPrivate   BaseBoolInt        `json:"is_private"` // Information whether story is private (0 - no, 1 - yes).
	Live        BasePropertyExists `json:"live"`       // Returns if the video is a live stream
	OwnerId     int                `json:"owner_id"`   // Video owner ID
	Player      string             `json:"player"`     // URL of the page with a player that can be used to play the video in the browser.
	Processing  BasePropertyExists `json:"processing"` // Returns if the video is processing
	Title       string             `json:"title"`      // Video title
	Type        string             `json:"type"`
	Views       int                `json:"views"` // Number of views
	Width       int                `json:"width"` // Video width
}

// StoriesUploadLinkText type represents `stories_upload_link_text` API object
type StoriesUploadLinkText string
