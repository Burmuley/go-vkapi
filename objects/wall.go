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
// `wall` group of objects
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

// WallAppPost type represents `wall_app_post` API object
type WallAppPost struct {
	Id       int    `json:"id"`        // Application ID
	Name     string `json:"name"`      // Application name
	Photo130 string `json:"photo_130"` // URL of the preview image with 130 px in width
	Photo604 string `json:"photo_604"` // URL of the preview image with 604 px in width
}

// WallAttachedNote type represents `wall_attached_note` API object
type WallAttachedNote struct {
	Comments     int    `json:"comments"`      // Comments number
	Date         int    `json:"date"`          // Date when the note has been created in Unixtime
	Id           int    `json:"id"`            // Note ID
	OwnerId      int    `json:"owner_id"`      // Note owner's ID
	ReadComments int    `json:"read_comments"` // Read comments number
	Title        string `json:"title"`         // Note title
	ViewUrl      string `json:"view_url"`      // URL of the page with note preview
}

// WallCommentAttachment type represents `wall_comment_attachment` API object
type WallCommentAttachment struct {
	Audio             AudioAudio                `json:"audio"`
	Doc               DocsDoc                   `json:"doc"`
	Link              BaseLink                  `json:"link"`
	Market            MarketMarketItem          `json:"market"`
	MarketMarketAlbum MarketMarketAlbum         `json:"market_market_album"`
	Note              WallAttachedNote          `json:"note"`
	Page              PagesWikipageFull         `json:"page"`
	Photo             PhotosPhoto               `json:"photo"`
	Sticker           BaseSticker               `json:"sticker"`
	Type              WallCommentAttachmentType `json:"type"`
	Video             VideoVideo                `json:"video"`
}

// WallCommentAttachmentType type represents `wall_comment_attachment_type` API object
type WallCommentAttachmentType string // Attachment type

// WallGeo type represents `wall_geo` API object
type WallGeo struct {
	Coordinates string    `json:"coordinates"` // Coordinates as string. <latitude> <longtitude>
	Place       BasePlace `json:"place"`
	Showmap     int       `json:"showmap"` // Information whether a map is showed
	Type        string    `json:"type"`    // Place type
}

// WallGraffiti type represents `wall_graffiti` API object
type WallGraffiti struct {
	Id       int    `json:"id"`        // Graffiti ID
	OwnerId  int    `json:"owner_id"`  // Graffiti owner's ID
	Photo200 string `json:"photo_200"` // URL of the preview image with 200 px in width
	Photo586 string `json:"photo_586"` // URL of the preview image with 586 px in width
}

// WallPostSource type represents `wall_post_source` API object
type WallPostSource struct {
	Data     string             `json:"data"`     // Additional data
	Platform string             `json:"platform"` // Platform name
	Type     WallPostSourceType `json:"type"`
	Url      string             `json:"url"` // URL to an external site used to publish the post
}

// WallPostSourceType type represents `wall_post_source_type` API object
type WallPostSourceType string // Type of post source

// WallPostType type represents `wall_post_type` API object
type WallPostType string // Post type

// WallPostedPhoto type represents `wall_posted_photo` API object
type WallPostedPhoto struct {
	Id       int    `json:"id"`        // Photo ID
	OwnerId  int    `json:"owner_id"`  // Photo owner's ID
	Photo130 string `json:"photo_130"` // URL of the preview image with 130 px in width
	Photo604 string `json:"photo_604"` // URL of the preview image with 604 px in width
}

// WallViews type represents `wall_views` API object
type WallViews struct {
	Count int `json:"count"` // Count
}

// WallWallComment type represents `wall_wall_comment` API object
type WallWallComment struct {
	Attachments    []WallCommentAttachment `json:"attachments"`
	Date           int                     `json:"date"` // Date when the comment has been added in Unixtime
	Deleted        bool                    `json:"deleted"`
	FromId         int                     `json:"from_id"` // Author ID
	Id             int                     `json:"id"`      // Comment ID
	Likes          BaseLikesInfo           `json:"likes"`
	OwnerId        int                     `json:"owner_id"`
	ParentsStack   []int                   `json:"parents_stack"`
	PostId         int                     `json:"post_id"`
	RealOffset     int                     `json:"real_offset"`      // Real position of the comment
	ReplyToComment int                     `json:"reply_to_comment"` // Replied comment ID
	ReplyToUser    int                     `json:"reply_to_user"`    // Replied user ID
	Text           string                  `json:"text"`             // Comment text
	Thread         CommentThread           `json:"thread"`
}

// WallWallpost type represents `wall_wallpost` API object
type WallWallpost struct {
	AccessKey   string                   `json:"access_key"` // Access key to private object
	Attachments []WallWallpostAttachment `json:"attachments"`
	Date        int                      `json:"date"`    // Date of publishing in Unixtime
	Edited      int                      `json:"edited"`  // Date of editing in Unixtime
	FromId      int                      `json:"from_id"` // Post author ID
	Geo         WallGeo                  `json:"geo"`
	Id          int                      `json:"id"`          // Post ID
	IsArchived  bool                     `json:"is_archived"` // Is post archived, only for post owners
	IsFavorite  bool                     `json:"is_favorite"` // Information whether the post in favorites list
	Likes       BaseLikesInfo            `json:"likes"`       // Count of likes
	OwnerId     int                      `json:"owner_id"`    // Wall owner's ID
	PostSource  WallPostSource           `json:"post_source"`
	PostType    WallPostType             `json:"post_type"`
	Reposts     BaseRepostsInfo          `json:"reposts"`   // Count of views
	SignerId    int                      `json:"signer_id"` // Post signer ID
	Text        string                   `json:"text"`      // Post text
	Views       WallViews                `json:"views"`     // Count of views
}

// WallWallpostAttachment type represents `wall_wallpost_attachment` API object
type WallWallpostAttachment struct {
	AccessKey   string                     `json:"access_key"` // Access key for the audio
	Album       PhotosPhotoAlbum           `json:"album"`
	App         WallAppPost                `json:"app"`
	Audio       AudioAudio                 `json:"audio"`
	Doc         DocsDoc                    `json:"doc"`
	Event       EventsEventAttach          `json:"event"`
	Graffiti    WallGraffiti               `json:"graffiti"`
	Link        BaseLink                   `json:"link"`
	Market      MarketMarketItem           `json:"market"`
	MarketAlbum MarketMarketAlbum          `json:"market_album"`
	Note        WallAttachedNote           `json:"note"`
	Page        PagesWikipageFull          `json:"page"`
	Photo       PhotosPhoto                `json:"photo"`
	PhotosList  []string                   `json:"photos_list"`
	Poll        PollsPoll                  `json:"poll"`
	PostedPhoto WallPostedPhoto            `json:"posted_photo"`
	Type        WallWallpostAttachmentType `json:"type"`
	Video       VideoVideo                 `json:"video"`
}

// WallWallpostAttachmentType type represents `wall_wallpost_attachment_type` API object
type WallWallpostAttachmentType string // Attachment type

// WallWallpostFull type represents `wall_wallpost_full` API object
type WallWallpostFull struct {
	WallWallpost WallWallpost     `json:"WallWallpost"`
	CanDelete    BaseBoolInt      `json:"can_delete"` // Information whether current user can delete the post
	CanEdit      BaseBoolInt      `json:"can_edit"`   // Information whether current user can edit the post
	CanPin       BaseBoolInt      `json:"can_pin"`    // Information whether current user can pin the post
	Comments     BaseCommentsInfo `json:"comments"`
	CopyHistory  []WallWallpost   `json:"copy_history"`
	CreatedBy    int              `json:"created_by"`    // Post creator ID (if post still can be edited)
	IsPinned     int              `json:"is_pinned"`     // Information whether the post is pinned
	MarkedAsAds  BaseBoolInt      `json:"marked_as_ads"` // Information whether the post is marked as ads
}

// WallWallpostToId type represents `wall_wallpost_to_id` API object
type WallWallpostToId struct {
	Attachments []WallWallpostAttachment `json:"attachments"`
	Comments    BaseCommentsInfo         `json:"comments"`
	CopyOwnerId int                      `json:"copy_owner_id"` // ID of the source post owner
	CopyPostId  int                      `json:"copy_post_id"`  // ID of the source post
	Date        int                      `json:"date"`          // Date of publishing in Unixtime
	FromId      int                      `json:"from_id"`       // Post author ID
	Geo         WallGeo                  `json:"geo"`
	Id          int                      `json:"id"` // Post ID
	Likes       BaseLikesInfo            `json:"likes"`
	PostId      int                      `json:"post_id"` // wall post ID (if comment)
	PostSource  WallPostSource           `json:"post_source"`
	PostType    WallPostType             `json:"post_type"`
	Reposts     BaseRepostsInfo          `json:"reposts"`
	SignerId    int                      `json:"signer_id"` // Post signer ID
	Text        string                   `json:"text"`      // Post text
	ToId        int                      `json:"to_id"`     // Wall owner's ID
}
