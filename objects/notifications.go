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
// `notifications` group of objects
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

// NotificationsSendMessageItem type represents `notifications_send_message_item` API object
type NotificationsSendMessageItem struct {
	Error  NotificationsSendMessageError `json:"error"`
	Status bool                          `json:"status"`  // Notification status
	UserId int                           `json:"user_id"` // User ID
}

// NotificationsReply type represents `notifications_reply` API object
type NotificationsReply struct {
	Date int `json:"date"` // Date when the reply has been created in Unixtime
	Id   int `json:"id"`   // Reply ID
	Text int `json:"text"` // Reply text
}

// NotificationsSendMessageError type represents `notifications_send_message_error` API object
type NotificationsSendMessageError struct {
	Code        int    `json:"code"`        // Error code
	Description string `json:"description"` // Error description
}

// NotificationsNotificationParent type represents `notifications_notification_parent` API object
type NotificationsNotificationParent struct {
	BoardTopic                        BoardTopic                        `json:"BoardTopic"`
	NotificationsNotificationsComment NotificationsNotificationsComment `json:"NotificationsNotificationsComment"`
	PhotosPhoto                       PhotosPhoto                       `json:"PhotosPhoto"`
	VideoVideo                        VideoVideo                        `json:"VideoVideo"`
	WallWallpostToId                  WallWallpostToId                  `json:"WallWallpostToId"`
}

// NotificationsFeedback type represents `notifications_feedback` API object
type NotificationsFeedback struct {
	Attachments []WallWallpostAttachment `json:"attachments"`
	FromId      int                      `json:"from_id"` // Reply author's ID
	Geo         BaseGeo                  `json:"geo"`
	Id          int                      `json:"id"` // Item ID
	Likes       BaseLikesInfo            `json:"likes"`
	Text        string                   `json:"text"`  // Reply text
	ToId        int                      `json:"to_id"` // Wall owner's ID
}

// NotificationsNotification type represents `notifications_notification` API object
type NotificationsNotification struct {
	Date     int                             `json:"date"` // Date when the event has been occurred
	Feedback NotificationsFeedback           `json:"feedback"`
	Parent   NotificationsNotificationParent `json:"parent"`
	Reply    NotificationsReply              `json:"reply"`
	Type     string                          `json:"type"` // Notification type
}

// NotificationsNotificationsComment type represents `notifications_notifications_comment` API object
type NotificationsNotificationsComment struct {
	Date    int          `json:"date"`     // Date when the comment has been added in Unixtime
	Id      int          `json:"id"`       // Comment ID
	OwnerId int          `json:"owner_id"` // Author ID
	Photo   PhotosPhoto  `json:"photo"`
	Post    WallWallpost `json:"post"`
	Text    string       `json:"text"` // Comment text
	Topic   BoardTopic   `json:"topic"`
	Video   VideoVideo   `json:"video"`
}
