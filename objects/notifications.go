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

import (
	"encoding/json"
)

/////////////////////////////////////////////////////////////////////////////////////////////////////////////
// `notifications` group of objects
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

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

// NotificationsNotificationParent type represents `notifications_notification_parent` API object
type NotificationsNotificationParent struct {
	AccessKey   string                   `json:"access_key"`  // Video access key
	AddingDate  int                      `json:"adding_date"` // Date when the video has been added in Unixtime
	AlbumId     int                      `json:"album_id"`    // Album ID
	Attachments []WallWallpostAttachment `json:"attachments"`
	CanAdd      BaseBoolInt              `json:"can_add"`       // Information whether current user can add the video
	CanComment  BaseBoolInt              `json:"can_comment"`   // Information whether current user can comment the video
	CanEdit     BaseBoolInt              `json:"can_edit"`      // Information whether current user can edit the video
	CanLike     BaseBoolInt              `json:"can_like"`      // Information whether current user can like the video
	CanRepost   BaseBoolInt              `json:"can_repost"`    // Information whether current user can repost this video
	Comments    int                      `json:"comments"`      // Number of comments
	CopyOwnerId int                      `json:"copy_owner_id"` // ID of the source post owner
	CopyPostId  int                      `json:"copy_post_id"`  // ID of the source post
	Created     int                      `json:"created"`       // Date when the topic has been created in Unixtime
	CreatedBy   int                      `json:"created_by"`    // Creator ID
	Date        int                      `json:"date"`          // Date when the comment has been added in Unixtime
	Description string                   `json:"description"`   // Video description
	Duration    int                      `json:"duration"`      // Video duration in seconds
	Files       VideoVideoFiles          `json:"files"`
	FirstFrame  []VideoVideoImage        `json:"first_frame"`
	FromId      int                      `json:"from_id"` // Post author ID
	Geo         WallGeo                  `json:"geo"`
	Height      int                      `json:"height"` // Video height
	Id          int                      `json:"id"`     // Comment ID
	Image       []VideoVideoImage        `json:"image"`
	Images      []PhotosImage            `json:"images"`
	IsClosed    BaseBoolInt              `json:"is_closed"` // Information whether the topic is closed
	IsFavorite  bool                     `json:"is_favorite"`
	IsFixed     BaseBoolInt              `json:"is_fixed"` // Information whether the topic is fixed
	Lat         json.Number              `json:"lat"`      // Latitude
	Likes       BaseLikesInfo            `json:"likes"`
	Live        BasePropertyExists       `json:"live"`     // Returns if the video is a live stream
	Long        json.Number              `json:"long"`     // Longitude
	OwnerId     int                      `json:"owner_id"` // Author ID
	Photo       PhotosPhoto              `json:"photo"`
	Player      string                   `json:"player"` // URL of the page with a player that can be used to play the video in the browser.
	Post        WallWallpost             `json:"post"`
	PostId      int                      `json:"post_id"` // Post ID
	PostSource  WallPostSource           `json:"post_source"`
	PostType    WallPostType             `json:"post_type"`
	Processing  BasePropertyExists       `json:"processing"` // Returns if the video is processing
	Reposts     BaseRepostsInfo          `json:"reposts"`
	SignerId    int                      `json:"signer_id"` // Post signer ID
	Sizes       []PhotosPhotoSizes       `json:"sizes"`
	Text        string                   `json:"text"`  // Comment text
	Title       string                   `json:"title"` // Video title
	ToId        int                      `json:"to_id"` // Wall owner's ID
	Topic       BoardTopic               `json:"topic"`
	Type        string                   `json:"type"`
	Updated     int                      `json:"updated"`    // Date when the topic has been updated in Unixtime
	UpdatedBy   int                      `json:"updated_by"` // ID of user who updated the topic
	UserId      int                      `json:"user_id"`    // ID of the user who have uploaded the photo
	Video       VideoVideo               `json:"video"`
	Views       int                      `json:"views"` // Number of views
	Width       int                      `json:"width"` // Video width
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

// NotificationsSendMessageItem type represents `notifications_send_message_item` API object
type NotificationsSendMessageItem struct {
	Error  NotificationsSendMessageError `json:"error"`
	Status bool                          `json:"status"`  // Notification status
	UserId int                           `json:"user_id"` // User ID
}
