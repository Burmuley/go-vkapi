package wall

import (
	"fmt"
	"gitlab.com/Burmuley/go-vkapi/objects/audio"
	"gitlab.com/Burmuley/go-vkapi/objects/base"
	"gitlab.com/Burmuley/go-vkapi/objects/docs"
	"gitlab.com/Burmuley/go-vkapi/objects/events"
	"gitlab.com/Burmuley/go-vkapi/objects/market"
	"gitlab.com/Burmuley/go-vkapi/objects/pages"
	"gitlab.com/Burmuley/go-vkapi/objects/photos"
	"gitlab.com/Burmuley/go-vkapi/objects/polls"
	"gitlab.com/Burmuley/go-vkapi/objects/videos"
)

/////////////////////////////////////////////////////////////
// Wall related API objects	                               //
/////////////////////////////////////////////////////////////

// AppPost represents `wall_app_post` API object
type AppPost struct {
	ID       int    `json:"id"`        // Application ID
	Name     string `json:"name"`      // Application name
	Photo130 string `json:"photo_130"` // URL of the preview image with 130 px in width
	Photo604 string `json:"photo_604"` // URL of the preview image with 604 px in width
}

// AttachedNote represents `wall_attached_note` API object
type AttachedNote struct {
	Comments     int    `json:"comments"`      // Comments number
	Date         int    `json:"date"`          // Date when the note has been created in Unixtime
	ID           int    `json:"id"`            // Note ID
	OwnerID      int    `json:"owner_id"`      // Note owner's ID
	ReadComments int    `json:"read_comments"` // Read comments number
	Title        string `json:"title"`         // Note title
	ViewURL      string `json:"view_url"`      // URL of the page with note preview
}

// CommentAttachment represents `wall_comment_attachment` API object
type CommentAttachment struct {
	Audio       audio.AudioFull       `json:"audio"`
	Doc         docs.Doc              `json:"doc"`
	Link        base.Link             `json:"link"`
	Market      market.Item           `json:"market"`
	MarketAlbum market.Album          `json:"market_market_album"`
	Note        AttachedNote          `json:"note"`
	Page        pages.WikipageFull    `json:"page"`
	Photo       photos.Photo          `json:"photo"`
	Sticker     base.Sticker          `json:"sticker"`
	Type        CommentAttachmentType `json:"type"`
	Video       videos.Video          `json:"video"`
}

// CommentAttachmentType represents `wall_comment_attachment_type` API object
type CommentAttachmentType string

func (c *CommentAttachmentType) MarshalJSON() ([]byte, error) {
	switch *c {
	case "photo", "video", "audio", "doc",
		"link", "note", "page", "market_market_album",
		"market", "sticker":
		return []byte(*c), nil
	}
	return []byte{}, fmt.Errorf("value is not in allowed range")
}

func (c *CommentAttachmentType) GetName() string {
	return string(*c)
}

// CommentThread represents `wall_comment_thread` API object
type CommentThread struct {
	CanPost       base.BoolInt `json:"can_post"`        // Information whether current user can comment the post
	Count         int          `json:"count"`           // Comments number
	GroupsCanPost base.BoolInt `json:"groups_can_post"` // Information whether groups can comment the post
}

// Graffiti represents `wall_graffiti` API object
type Graffiti struct {
	ID       int    `json:"id"`        // Graffiti ID
	OwnerID  int    `json:"owner_id"`  // Graffiti owner's ID
	Photo200 string `json:"photo_200"` // URL of the preview image with 200 px in width
	Photo586 string `json:"photo_586"` // URL of the preview image with 586 px in width
}

// PostSourceType represents `wall_post_source_type` API object
type PostSourceType string

func (p *PostSourceType) MarshalJSON() ([]byte, error) {
	switch *p {
	case "vk", "widget", "api",
		"rss", "sms":
		return []byte(*p), nil
	}

	return []byte{}, fmt.Errorf("value is not in allowed range")
}

func (p *PostSourceType) GetName() string {
	return string(*p)
}

// PostSource represents `wall_post_source` API object
type PostSource struct {
	Data     string         `json:"data"`     // Additional data
	Platform string         `json:"platform"` // Platform name
	Type     PostSourceType `json:"type"`
	Url      string         `json:"url"` // URL to an external site used to publish the post
}

// PostType represents `wall_post_type` API object
type PostType string

func (p *PostType) MarshalJSON() ([]byte, error) {
	switch *p {
	case "post", "copy", "reply",
		"postpone", "suggest":
		return []byte(*p), nil
	}

	return []byte{}, fmt.Errorf("value is not in allowed range")
}

func (p *PostType) GetName() string {
	return string(*p)
}

// PostedPhoto represents `wall_posted_photo` API object
type PostedPhoto struct {
	ID       int    `json:"id"`        // Photo ID
	OwnerID  int    `json:"owner_id"`  // Photo owner's ID
	Photo130 string `json:"photo_130"` // URL of the preview image with 130 px in width
	Photo604 string `json:"photo_604"` // URL of the preview image with 604 px in width
}

// Views represents `wall_views` API object
type Views struct {
	Count int `json:"count"`
}

// WallComment represents `wall_wall_comment` API object
type WallComment struct {
	Attachments    CommentAttachment `json:"attachments"`
	Date           int               `json:"date"`    // Date when the comment has been added in Unixtime
	FromID         int               `json:"from_id"` // Author ID
	ID             int               `json:"id"`      // Comment ID
	Likes          base.LikesInfo    `json:"likes"`
	RealOffset     int               `json:"real_offset"`      // Real position of the comment
	ReplyToComment int               `json:"reply_to_comment"` // Replied comment ID
	ReplyToUser    int               `json:"reply_to_user"`    // Replied user ID
	Text           string            `json:"text"`             // Comment text
	Thread         CommentThread     `json:"thread"`
}

// Wallpost represents `wall_wallpost` API object
type Wallpost struct {
	AccessKey   string               `json:"access_key"` // Access key to private object
	Attachments []WallpostAttachment `json:"attachments"`
	Date        int                  `json:"date"`    // Date of publishing in Unixtime
	Edited      int                  `json:"edited"`  // Date of editing in Unixtime
	FromID      int                  `json:"from_id"` // Post author ID
	Geo         base.Geo             `json:"geo"`
	ID          int                  `json:"id"`          // Post ID
	IsArchived  bool                 `json:"is_archived"` // Is post archived, only for post owners
	IsFavorite  bool                 `json:"is_favorite"` // Information whether the post in favorites list
	Likes       base.LikesInfo       `json:"likes"`       // Count of likes
	OwnerID     int                  `json:"owner_id"`    // Wall owner's ID
	PostSource  PostSource           `json:"post_source"`
	PostType    PostType             `json:"post_type"`
	Reposts     base.RepostsInfo     `json:"reposts"`   // Count of views
	SignerID    int                  `json:"signer_id"` // Post signer ID
	Text        string               `json:"text"`      // Post text
	Views       Views                `json:"views"`     // Count of views
}

// WallpostAttachment represents `wall_wallpost_attachment` API object
type WallpostAttachment struct {
	AccessKey   string                 `json:"access_key"` // Access key for the audio
	Album       photos.Album           `json:"album"`
	App         AppPost                `json:"app"`
	Audio       audio.AudioFull        `json:"audio"`
	Doc         docs.Doc               `json:"doc"`
	Event       events.EventAttach     `json:"event"`
	Graffiti    Graffiti               `json:"graffiti"`
	Link        base.Link              `json:"link"`
	Market      market.Item            `json:"market"`
	MarketAlbum market.Album           `json:"market_album"`
	Note        AttachedNote           `json:"note"`
	Page        pages.WikipageFull     `json:"page"`
	Photo       photos.Photo           `json:"photo"`
	PhotosList  []string               `json:"photos_list"` // String ID of photo
	Poll        polls.Poll             `json:"poll"`
	PostedPhoto PostedPhoto            `json:"posted_photo"`
	Type        WallpostAttachmentType `json:"type"`
	Video       videos.Video           `json:"video"`
}

// WallpostAttachmentType represents `wall_wallpost_attachment_type` API object
type WallpostAttachmentType string

func (w *WallpostAttachmentType) MarshalJSON() ([]byte, error) {
	switch *w {
	case "photo", "posted_photo", "audio", "video", "doc",
		"link", "graffiti", "note", "app", "poll", "page",
		"album", "photos_list", "market_market_album", "market",
		"event":
		return []byte(*w), nil
	}

	return []byte{}, fmt.Errorf("value is not in allowed range")
}

func (w *WallpostAttachmentType) GetName() string {
	return string(*w)
}

// WallpostFull represents `wall_wallpost_full` API object
type WallpostFull struct {
	Wallpost
	CopyHistory Wallpost          `json:"copy_history"`
	CanEdit     base.BoolInt      `json:"can_edit"`   // Information whether current user can edit the post
	CreatorID   int               `json:"created_by"` // Post creator ID (if post still can be edited)
	CanDelete   base.BoolInt      `json:"can_delete"` // Information whether current user can delete the post
	CanPin      base.BoolInt      `json:"can_pin"`    // Information whether current user can pin the post
	IsPinned    int               `json:"is_pinned"`  // Information whether the post is pinned
	Comments    base.CommentsInfo `json:"comments"`
	IsAds       base.BoolInt      `json:"marked_as_ads"` // Information whether the post is marked as ads
}

// WallpostToId represents `wall_wallpost_to_id` API object
type WallpostToId struct {
	Attachments []WallpostAttachment `json:"attachments"`
	Comments    base.CommentsInfo    `json:"comments"`
	CopyOwnerID int                  `json:"copy_owner_id"` // ID of the source post owner
	CopyPostID  int                  `json:"copy_post_id"`  // ID of the source post
	Date        int                  `json:"date"`          // Date of publishing in Unixtime
	FromID      int                  `json:"from_id"`       // Post author ID
	Geo         base.Geo             `json:"geo"`
	ID          int                  `json:"id"` // Post ID
	Likes       base.LikesInfo       `json:"likes"`
	PostID      int                  `json:"post_id"` // wall post ID (if comment)
	PostSource  PostSource           `json:"post_source"`
	PostType    PostType             `json:"post_type"`
	Reposts     base.RepostsInfo     `json:"reposts"`
	SignerID    int                  `json:"signer_id"` // Post signer ID
	Text        string               `json:"text"`      // Post text
	ToID        int                  `json:"to_id"`     // Wall owner's ID
}
