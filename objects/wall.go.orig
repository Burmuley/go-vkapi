package objects

/////////////////////////////////////////////////////////////
// Wall related API objects	                               //
/////////////////////////////////////////////////////////////

// AppPost represents `wall_app_post` API object
type WallAppPost struct {
	ID       int    `json:"id"`        // Application ID
	Name     string `json:"name"`      // Application name
	Photo130 string `json:"photo_130"` // URL of the preview image with 130 px in width
	Photo604 string `json:"photo_604"` // URL of the preview image with 604 px in width
}

// AttachedNote represents `wall_attached_note` API object
type WallAttachedNote struct {
	Comments     int    `json:"comments"`      // Comments number
	Date         int    `json:"date"`          // Date when the note has been created in Unixtime
	ID           int    `json:"id"`            // Note ID
	OwnerID      int    `json:"owner_id"`      // Note owner's ID
	ReadComments int    `json:"read_comments"` // Read comments number
	Title        string `json:"title"`         // Note title
	ViewURL      string `json:"view_url"`      // URL of the page with note preview
}

// CommentAttachment represents `wall_comment_attachment` API object
type WallCommentAttachment struct {
	Audio       AudioAudioFull            `json:"audio"`
	Doc         DocsDoc                   `json:"doc"`
	Link        BaseLink                  `json:"link"`
	Market      MarketItem                `json:"market"`
	MarketAlbum MarketAlbum               `json:"market_market_album"`
	Note        WallAttachedNote          `json:"note"`
	Page        PagesWikipageFull         `json:"page"`
	Photo       PhotosPhoto               `json:"photo"`
	Sticker     BaseSticker               `json:"sticker"`
	Type        WallCommentAttachmentType `json:"type"`
	Video       VideosVideo               `json:"video"`
}

// CommentAttachmentType represents `wall_comment_attachment_type` API object
type WallCommentAttachmentType string

func (c *WallCommentAttachmentType) MarshalJSON() ([]byte, error) {
	return GetStringFromRange(string(*c), "photo", "video", "audio", "doc", "link", "note", "page",
		"market_market_album", "market", "sticker")
}

func (c *WallCommentAttachmentType) GetName() string {
	return string(*c)
}

// CommentThread represents `wall_comment_thread` API object
type WallCommentThread struct {
	CanPost       BaseBoolInt `json:"can_post"`        // Information whether current user can comment the post
	Count         int         `json:"count"`           // Comments number
	GroupsCanPost BaseBoolInt `json:"groups_can_post"` // Information whether groups can comment the post
}

// Graffiti represents `wall_graffiti` API object
type WallGraffiti struct {
	ID       int    `json:"id"`        // Graffiti ID
	OwnerID  int    `json:"owner_id"`  // Graffiti owner's ID
	Photo200 string `json:"photo_200"` // URL of the preview image with 200 px in width
	Photo586 string `json:"photo_586"` // URL of the preview image with 586 px in width
}

// PostSourceType represents `wall_post_source_type` API object
type WallPostSourceType string

func (p *WallPostSourceType) MarshalJSON() ([]byte, error) {
	return GetStringFromRange(string(*p), "vk", "widget", "api", "rss", "sms")
}

func (p *WallPostSourceType) GetName() string {
	return string(*p)
}

// PostSource represents `wall_post_source` API object
type WallPostSource struct {
	Data     string             `json:"data"`     // Additional data
	Platform string             `json:"platform"` // Platform name
	Type     WallPostSourceType `json:"type"`
	Url      string             `json:"url"` // URL to an external site used to publish the post
}

// PostType represents `wall_post_type` API object
type WallPostType string

func (p *WallPostType) MarshalJSON() ([]byte, error) {
	return GetStringFromRange(string(*p), "post", "copy", "reply", "postpone", "suggest")
}

func (p *WallPostType) GetName() string {
	return string(*p)
}

// PostedPhoto represents `wall_posted_photo` API object
type WallPostedPhoto struct {
	ID       int    `json:"id"`        // Photo ID
	OwnerID  int    `json:"owner_id"`  // Photo owner's ID
	Photo130 string `json:"photo_130"` // URL of the preview image with 130 px in width
	Photo604 string `json:"photo_604"` // URL of the preview image with 604 px in width
}

// Views represents `wall_views` API object
type WallViews struct {
	Count int `json:"count"`
}

// WallComment represents `wall_wall_comment` API object
type WallWallComment struct {
	Attachments    WallCommentAttachment `json:"attachments"`
	Date           int                   `json:"date"`    // Date when the comment has been added in Unixtime
	FromID         int                   `json:"from_id"` // Author ID
	ID             int                   `json:"id"`      // Comment ID
	Likes          BaseLikesInfo         `json:"likes"`
	RealOffset     int                   `json:"real_offset"`      // Real position of the comment
	ReplyToComment int                   `json:"reply_to_comment"` // Replied comment ID
	ReplyToUser    int                   `json:"reply_to_user"`    // Replied user ID
	Text           string                `json:"text"`             // Comment text
	Thread         WallCommentThread     `json:"thread"`
}

// Wallpost represents `wall_wallpost` API object
type WallWallpost struct {
	AccessKey   string                   `json:"access_key"` // Access key to private object
	Attachments []WallWallpostAttachment `json:"attachments"`
	Date        int                      `json:"date"`    // Date of publishing in Unixtime
	Edited      int                      `json:"edited"`  // Date of editing in Unixtime
	FromID      int                      `json:"from_id"` // Post author ID
	Geo         BaseGeo                  `json:"geo"`
	ID          int                      `json:"id"`          // Post ID
	IsArchived  bool                     `json:"is_archived"` // Is post archived, only for post owners
	IsFavorite  bool                     `json:"is_favorite"` // Information whether the post in favorites list
	Likes       BaseLikesInfo            `json:"likes"`       // Count of likes
	OwnerID     int                      `json:"owner_id"`    // Wall owner's ID
	PostSource  WallPostSource           `json:"post_source"`
	PostType    WallPostType             `json:"post_type"`
	Reposts     BaseRepostsInfo          `json:"reposts"`   // Count of views
	SignerID    int                      `json:"signer_id"` // Post signer ID
	Text        string                   `json:"text"`      // Post text
	Views       WallViews                `json:"views"`     // Count of views
}

// WallpostAttachment represents `wall_wallpost_attachment` API object
type WallWallpostAttachment struct {
	AccessKey   string                     `json:"access_key"` // Access key for the audio
	Album       PhotosAlbum                `json:"album"`
	App         WallAppPost                `json:"app"`
	Audio       AudioAudioFull             `json:"audio"`
	Doc         DocsDoc                    `json:"doc"`
	Event       EventsEventAttach          `json:"event"`
	Graffiti    WallGraffiti               `json:"graffiti"`
	Link        BaseLink                   `json:"link"`
	Market      MarketItem                 `json:"market"`
	MarketAlbum MarketAlbum                `json:"market_album"`
	Note        WallAttachedNote           `json:"note"`
	Page        PagesWikipageFull          `json:"page"`
	Photo       PhotosPhoto                `json:"photo"`
	PhotosList  []string                   `json:"photos_list"` // String ID of photo
	Poll        PollsPoll                  `json:"poll"`
	PostedPhoto WallPostedPhoto            `json:"posted_photo"`
	Type        WallWallpostAttachmentType `json:"type"`
	Video       VideosVideo                `json:"video"`
}

// WallpostAttachmentType represents `wall_wallpost_attachment_type` API object
type WallWallpostAttachmentType string

func (w *WallWallpostAttachmentType) MarshalJSON() ([]byte, error) {
	return GetStringFromRange(string(*w), "photo", "posted_photo", "audio", "video", "doc", "link",
		"graffiti", "note", "app", "poll", "page", "album", "photos_list", "market_market_album", "market", "event")
}

func (w *WallWallpostAttachmentType) GetName() string {
	return string(*w)
}

// WallpostFull represents `wall_wallpost_full` API object
type WallWallpostFull struct {
	*WallWallpost
	CopyHistory WallWallpost     `json:"copy_history"`
	CanEdit     BaseBoolInt      `json:"can_edit"`   // Information whether current user can edit the post
	CreatorID   int              `json:"created_by"` // Post creator ID (if post still can be edited)
	CanDelete   BaseBoolInt      `json:"can_delete"` // Information whether current user can delete the post
	CanPin      BaseBoolInt      `json:"can_pin"`    // Information whether current user can pin the post
	IsPinned    int              `json:"is_pinned"`  // Information whether the post is pinned
	Comments    BaseCommentsInfo `json:"comments"`
	IsAds       BaseBoolInt      `json:"marked_as_ads"` // Information whether the post is marked as ads
}

// WallpostToId represents `wall_wallpost_to_id` API object
type WallWallpostToId struct {
	Attachments []WallWallpostAttachment `json:"attachments"`
	Comments    BaseCommentsInfo         `json:"comments"`
	CopyOwnerID int                      `json:"copy_owner_id"` // ID of the source post owner
	CopyPostID  int                      `json:"copy_post_id"`  // ID of the source post
	Date        int                      `json:"date"`          // Date of publishing in Unixtime
	FromID      int                      `json:"from_id"`       // Post author ID
	Geo         BaseGeo                  `json:"geo"`
	ID          int                      `json:"id"` // Post ID
	Likes       BaseLikesInfo            `json:"likes"`
	PostID      int                      `json:"post_id"` // wall post ID (if comment)
	PostSource  WallPostSource           `json:"post_source"`
	PostType    WallPostType             `json:"post_type"`
	Reposts     BaseRepostsInfo          `json:"reposts"`
	SignerID    int                      `json:"signer_id"` // Post signer ID
	Text        string                   `json:"text"`      // Post text
	ToID        int                      `json:"to_id"`     // Wall owner's ID
}
