package widgets

import (
	"gitlab.com/Burmuley/go-vkapi/objects"
	"gitlab.com/Burmuley/go-vkapi/objects/base"
	"gitlab.com/Burmuley/go-vkapi/objects/users"
	"gitlab.com/Burmuley/go-vkapi/objects/wall"
)

/////////////////////////////////////////////////////////////
// Widgets related API objects	                           //
/////////////////////////////////////////////////////////////

type CommentMediaType string

func (c *CommentMediaType) MarshalJSON() ([]byte, error) {
	return objects.GetStringFromRange(string(*c), "audio", "photo", "video")
}

func (c *CommentMediaType) GetName() string {
	return string(*c)
}

// CommentMedia represents `widgets_comment_media` API object
type CommentMedia struct {
	ItemID   int              `json:"item_id"`   // Media item ID
	OwnerID  int              `json:"owner_id"`  // Media owner's ID
	ThumbSrc string           `json:"thumb_src"` // URL of the preview image (type=photo only)
	Type     CommentMediaType `json:"type"`
}

// CommentRepliesItem represents `widgets_comment_replies_item` API object
type CommentRepliesItem struct {
	CommentID int            `json:"cid"`  // Comment ID
	Date      int            `json:"date"` // Date when the comment has been added in Unixtime
	Likes     WidgetLikes    `json:"likes"`
	Text      string         `json:"text"` // Comment text
	UserID    int            `json:"uid"`  // User ID
	User      users.UserFull `json:"user"`
}

// CommentReplies represents `widgets_comment_replies` API object
type CommentReplies struct {
	CanPost base.BoolInt         `json:"can_post"` // Information whether current user can comment the post
	Count   int                  `json:"count"`    // Comments number
	Replies []CommentRepliesItem `json:"replies"`
}

// WidgetComment represents `widgets_widget_comment` API object
type WidgetComment struct {
	Attachments []wall.CommentAttachment `json:"attachments"`
	CanDelete   base.BoolInt             `json:"can_delete"` // Information whether current user can delete the comment
	Comments    CommentReplies           `json:"comments"`
	Date        int                      `json:"date"`    // Date when the comment has been added in Unixtime
	AuthorID    int                      `json:"from_id"` // Comment author ID
	ID          int                      `json:"id"`      // Comment ID
	Likes       base.LikesInfo           `json:"likes"`
	Media       CommentMedia             `json:"media"`
	PostSource  wall.PostSource          `json:"post_source"`
	PostType    int                      `json:"post_type"` // Post type
	Reposts     base.RepostsInfo         `json:"reposts"`
	Text        string                   `json:"text"`  // Comment text
	ToID        int                      `json:"to_id"` // Wall owner
	User        users.UserFull           `json:"user"`
}

// WidgetLikes represents `widgets_widget_likes` API object
type WidgetLikes struct {
	Count int `json:"count"` // Likes number
}

// WidgetPage represents `widgets_widget_page` API object
type WidgetPage struct {
	Comments    base.ObjectCount `json:"comments"`
	Date        int              `json:"date"`        // Date when widgets on the page has been initialized firstly in Unixtime
	Description string           `json:"description"` // Page description
	ID          int              `json:"id"`          // Page ID
	Likes       base.ObjectCount `json:"likes"`
	PageID      int              `json:"page_id"` // page_id parameter value
	Photo       string           `json:"photo"`   // URL of the preview image
	Title       string           `json:"title"`   // Page title
	Url         string           `json:"url"`     // Page absolute URL
}
