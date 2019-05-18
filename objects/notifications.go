package objects

/////////////////////////////////////////////////////////////
// Notifications related API objects	                   //
/////////////////////////////////////////////////////////////

// Feedback represents `notifications_feedback` API objects
type NotificationsFeedback struct {
	Attachments []WallWallpostAttachment `json:"attachments"`
	FromID      int                      `json:"from_id"` // Reply author's ID
	ID          int                      `json:"id"`      // Item ID
	Geo         BaseGeo                  `json:"geo"`
	Likes       BaseLikesInfo            `json:"likes"`
	Text        string                   `json:"text"`  // Reply text
	ToID        int                      `json:"to_id"` // Wall owner's ID
}

// Notification represents `notifications_notification` API object
type NotificationsNotification struct {
	Date     int                             `json:"date"` // Date when the event has been occured
	Feedback NotificationsFeedback           `json:"feedback"`
	Parent   NotificationsNotificationParent `json:"parent"`
	Reply    NotificationsReply              `json:"reply"`
	Type     string                          `json:"type"` // Notification type
}

// NotificationParent represents `notifications_notification_parent` API object
type NotificationsNotificationParent struct {
	*WallWallpostToId
	*PhotosPhoto
	*BoardTopic
	*VideosVideo
	*NotificationsComment
}

// Comment represents `notifications_notifications_comment` API object
type NotificationsComment struct {
	Date    int          `json:"date"`     // Date when the comment has been added in Unixtime
	ID      int          `json:"id"`       // Comment ID
	OwnerID int          `json:"owner_id"` // Author ID
	Photo   PhotosPhoto  `json:"photo"`
	Post    WallWallpost `json:"post"`
	Text    string       `json:"text"` // Comment text
	Topic   BoardTopic   `json:"topic"`
	Video   VideosVideo  `json:"video"`
}

// Reply represents `notifications_notification_reply` API object
type NotificationsReply struct {
	Date int `json:"date"` // Date when the reply has been created in Unixtime
	ID   int `json:"id"`   // Reply ID
	Text int `json:"text"` // Reply text
}
