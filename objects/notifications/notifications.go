package notifications

import (
	"gitlab.com/Burmuley/go-vkapi/objects/base"
	"gitlab.com/Burmuley/go-vkapi/objects/board"
	"gitlab.com/Burmuley/go-vkapi/objects/photos"
	"gitlab.com/Burmuley/go-vkapi/objects/videos"
	"gitlab.com/Burmuley/go-vkapi/objects/wall"
)

/////////////////////////////////////////////////////////////
// Notifications related API objects	                   //
/////////////////////////////////////////////////////////////

// Feedback represents `notifications_feedback` API objects
type Feedback struct {
	Attachments []wall.WallpostAttachment `json:"attachments"`
	FromID      int                       `json:"from_id"` // Reply author's ID
	ID          int                       `json:"id"`      // Item ID
	Geo         base.Geo                  `json:"geo"`
	Likes       base.LikesInfo            `json:"likes"`
	Text        string                    `json:"text"`  // Reply text
	ToID        int                       `json:"to_id"` // Wall owner's ID
}

// Notification represents `notifications_notification` API object
type Notification struct {
	Date     int                `json:"date"` // Date when the event has been occured
	Feedback Feedback           `json:"feedback"`
	Parent   NotificationParent `json:"parent"`
	Reply    Reply              `json:"reply"`
	Type     string             `json:"type"` // Notification type
}

// NotificationParent represents `notifications_notification_parent` API object
type NotificationParent struct {
	wall.WallpostToId
	photos.Photo
	board.Topic
	videos.Video
	Comment
}

// Comment represents `notifications_notifications_comment` API object
type Comment struct {
	Date    int           `json:"date"`     // Date when the comment has been added in Unixtime
	ID      int           `json:"id"`       // Comment ID
	OwnerID int           `json:"owner_id"` // Author ID
	Photo   photos.Photo  `json:"photo"`
	Post    wall.Wallpost `json:"post"`
	Text    string        `json:"text"` // Comment text
	Topic   board.Topic   `json:"topic"`
	Video   videos.Video  `json:"video"`
}

// Reply represents `notifications_notification_reply` API object
type Reply struct {
	Date int `json:"date"` // Date when the reply has been created in Unixtime
	ID   int `json:"id"`   // Reply ID
	Text int `json:"text"` // Reply text
}
