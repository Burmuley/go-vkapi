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
// `widgets` group of objects
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

// WidgetsCommentMedia type represents `widgets_comment_media` API object
type WidgetsCommentMedia struct {
	ItemId   int                      `json:"item_id"`   // Media item ID
	OwnerId  int                      `json:"owner_id"`  // Media owner's ID
	ThumbSrc string                   `json:"thumb_src"` // URL of the preview image (type=photo only)
	Type     *WidgetsCommentMediaType `json:"type"`
}

// WidgetsCommentMediaType type represents `widgets_comment_media_type` API object
type WidgetsCommentMediaType string // Media type

// WidgetsCommentReplies type represents `widgets_comment_replies` API object
type WidgetsCommentReplies struct {
	CanPost *BaseBoolInt                 `json:"can_post"` // Information whether current user can comment the post
	Count   int                          `json:"count"`    // Comments number
	Replies []*WidgetsCommentRepliesItem `json:"replies"`
}

// WidgetsCommentRepliesItem type represents `widgets_comment_replies_item` API object
type WidgetsCommentRepliesItem struct {
	Cid   int                 `json:"cid"`  // Comment ID
	Date  int                 `json:"date"` // Date when the comment has been added in Unixtime
	Likes *WidgetsWidgetLikes `json:"likes"`
	Text  string              `json:"text"` // Comment text
	Uid   int                 `json:"uid"`  // User ID
	User  *UsersUserFull      `json:"user"`
}

// WidgetsWidgetComment type represents `widgets_widget_comment` API object
type WidgetsWidgetComment struct {
	Attachments []*WallCommentAttachment `json:"attachments"`
	CanDelete   *BaseBoolInt             `json:"can_delete"` // Information whether current user can delete the comment
	Comments    *WidgetsCommentReplies   `json:"comments"`
	Date        int                      `json:"date"`    // Date when the comment has been added in Unixtime
	FromId      int                      `json:"from_id"` // Comment author ID
	Id          int                      `json:"id"`      // Comment ID
	Likes       *BaseLikesInfo           `json:"likes"`
	Media       *WidgetsCommentMedia     `json:"media"`
	PostSource  *WallPostSource          `json:"post_source"`
	PostType    int                      `json:"post_type"` // Post type
	Reposts     *BaseRepostsInfo         `json:"reposts"`
	Text        string                   `json:"text"`  // Comment text
	ToId        int                      `json:"to_id"` // Wall owner
	User        *UsersUserFull           `json:"user"`
}

// WidgetsWidgetLikes type represents `widgets_widget_likes` API object
type WidgetsWidgetLikes struct {
	Count int `json:"count"` // Likes number
}

// WidgetsWidgetPage type represents `widgets_widget_page` API object
type WidgetsWidgetPage struct {
	Comments    *BaseObjectCount `json:"comments"`
	Date        int              `json:"date"`        // Date when widgets on the page has been initialized firstly in Unixtime
	Description string           `json:"description"` // Page description
	Id          int              `json:"id"`          // Page ID
	Likes       *BaseObjectCount `json:"likes"`
	PageId      string           `json:"page_id"` // page_id parameter value
	Photo       string           `json:"photo"`   // URL of the preview image
	Title       string           `json:"title"`   // Page title
	Url         string           `json:"url"`     // Page absolute URL
}
