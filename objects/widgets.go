package objects

/////////////////////////////////////////////////////////////
// Widgets related API objects	                           //
/////////////////////////////////////////////////////////////

type WidgetsCommentMediaType string

func (c *WidgetsCommentMediaType) MarshalJSON() ([]byte, error) {
	return GetStringFromRange(string(*c), "audio", "photo", "video")
}

func (c *WidgetsCommentMediaType) GetName() string {
	return string(*c)
}

// CommentMedia represents `widgets_comment_media` API object
type WidgetsCommentMedia struct {
	ItemID   int                     `json:"item_id"`   // Media item ID
	OwnerID  int                     `json:"owner_id"`  // Media owner's ID
	ThumbSrc string                  `json:"thumb_src"` // URL of the preview image (type=photo only)
	Type     WidgetsCommentMediaType `json:"type"`
}

// CommentRepliesItem represents `widgets_comment_replies_item` API object
type WidgetsCommentRepliesItem struct {
	CommentID int                `json:"cid"`  // Comment ID
	Date      int                `json:"date"` // Date when the comment has been added in Unixtime
	Likes     WidgetsWidgetLikes `json:"likes"`
	Text      string             `json:"text"` // Comment text
	UserID    int                `json:"uid"`  // User ID
	User      UsersUserFull      `json:"user"`
}

// CommentReplies represents `widgets_comment_replies` API object
type WidgetsCommentReplies struct {
	CanPost BaseBoolInt                 `json:"can_post"` // Information whether current user can comment the post
	Count   int                         `json:"count"`    // Comments number
	Replies []WidgetsCommentRepliesItem `json:"replies"`
}

// WidgetComment represents `widgets_widget_comment` API object
type WidgetsWidgetComment struct {
	Attachments []WallCommentAttachment `json:"attachments"`
	CanDelete   BaseBoolInt             `json:"can_delete"` // Information whether current user can delete the comment
	Comments    WidgetsCommentReplies   `json:"comments"`
	Date        int                     `json:"date"`    // Date when the comment has been added in Unixtime
	AuthorID    int                     `json:"from_id"` // Comment author ID
	ID          int                     `json:"id"`      // Comment ID
	Likes       BaseLikesInfo           `json:"likes"`
	Media       WidgetsCommentMedia     `json:"media"`
	PostSource  WallPostSource          `json:"post_source"`
	PostType    int                     `json:"post_type"` // Post type
	Reposts     BaseRepostsInfo         `json:"reposts"`
	Text        string                  `json:"text"`  // Comment text
	ToID        int                     `json:"to_id"` // Wall owner
	User        UsersUserFull           `json:"user"`
}

// WidgetLikes represents `widgets_widget_likes` API object
type WidgetsWidgetLikes struct {
	Count int `json:"count"` // Likes number
}

// WidgetPage represents `widgets_widget_page` API object
type WidgetsWidgetPage struct {
	Comments    BaseObjectCount `json:"comments"`
	Date        int             `json:"date"`        // Date when widgets on the page has been initialized firstly in Unixtime
	Description string          `json:"description"` // Page description
	ID          int             `json:"id"`          // Page ID
	Likes       BaseObjectCount `json:"likes"`
	PageID      int             `json:"page_id"` // page_id parameter value
	Photo       string          `json:"photo"`   // URL of the preview image
	Title       string          `json:"title"`   // Page title
	Url         string          `json:"url"`     // Page absolute URL
}
