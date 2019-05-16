package photos

import (
	"fmt"
	"gitlab.com/Burmuley/go-vkapi/objects/base"
	"gitlab.com/Burmuley/go-vkapi/objects/wall"
)

/////////////////////////////////////////////////////////////
// Photos related API objects	                           //
/////////////////////////////////////////////////////////////

// CommentXtrPid represents `photos_comment_xtr_pid` API object
type CommentXtrPid struct {
	Attachments    []wall.CommentAttachment `json:"attachments"`
	Date           int                      `json:"date"`    // Date when the comment has been added in Unixtime
	FromID         int                      `json:"from_id"` // Author ID
	ID             int                      `json:"id"`      // Comment ID
	Likes          base.LikesInfo           `json:"likes"`
	PhotoID        int                      `json:"pid"`              // Photo ID
	ReplyToComment int                      `json:"reply_to_comment"` // Replied comment ID
	ReplyToUser    int                      `json:"reply_to_user"`    // Replied user ID
	Text           string                   `json:"text"`             // Comment text
}

// Image represents `photos_image` API object
type Image struct {
	Height int       `json:"height"` // Height of the photo in px.
	Url    string    `json:"url"`    // Photo URL
	Type   ImageType `json:"type"`
	Width  int       `json:"width"` // Width of the photo in px.
}

// ImageType represents `photos_image_type` API object
type ImageType string

func (i *ImageType) MarshalJSON() ([]byte, error) {
	switch *i {
	case "s", "m", "x", "o", "p", "q", "r", "y", "z", "w":
		return []byte(*i), nil
	}

	return []byte{}, fmt.Errorf("value is not in valid range")
}

func (i *ImageType) GetName() string {
	return string(*i)
}

// MarketAlbumUploadResponse represents `photos_market_album_upload_response` API object
type MarketAlbumUploadResponse struct {
	GroupID int    `json:"gid"`    // Community ID
	Hash    string `json:"hash"`   // Uploading hash
	Photo   string `json:"photo"`  // Uploaded photo data
	Server  int    `json:"server"` // Upload server number
}

// MarketUploadResponse represents `photos_market_upload_response` API object
type MarketUploadResponse struct {
	CropData string `json:"crop_data"` // Crop data
	CropHash string `json:"crop_hash"` // Crop hash
	GroupID  int    `json:"group_id"`  // Community ID
	Hash     string `json:"hash"`      // Uploading hash
	Photo    string `json:"photo"`     // Uploaded photo data
	Server   int    `json:"server"`    // Upload server number
}

// MessageUploadResponse represents `photos_message_upload_response` API object
type MessageUploadResponse struct {
	Hash   string `json:"hash"`   // Uploading hash
	Photo  string `json:"photo"`  // Uploaded photo data
	Server int    `json:"server"` // Upload server number
}

// OwnerUploadResponse represents `photos_owner_upload_response` API object
type OwnerUploadResponse struct {
	Hash   string `json:"hash"`   // Uploading hash
	Photo  string `json:"photo"`  // Uploaded photo data
	Server int    `json:"server"` // Upload server number
}

// Photo represents `photos_photo` APi object
type Photo struct {
	AccessKey string  `json:"access_key"` // Access key for the photo
	AlbumID   int     `json:"album_id"`   // Album ID
	Date      int     `json:"date"`       // Date when uploaded
	Height    int     `json:"height"`     // Original photo height
	ID        int     `json:"id"`         // Photo ID
	Images    []Image `json:"images"`
	Latitude  float64 `json:"lat"`      // Latitude
	Longitude float64 `json:"long"`     // Longitude
	OwnerID   int     `json:"owner_id"` // Photo owner's ID
	PostID    int     `json:"post_id"`  // Post ID
	Text      string  `json:"text"`     // Photo caption
	UserID    int     `json:"user_id"`  // ID of the user who have uploaded the photo
	Width     int     `json:"width"`    // Original photo width
}

// Album represents `photos_photo_album` API object
type Album struct {
	Created     int    `json:"created"`     // Date when the album has been created in Unixtime
	Description string `json:"description"` // Photo album description
	ID          int    `json:"id"`          // Photo album ID
	OwnerID     int    `json:"owner_id"`    // Album owner's ID
	Size        int    `json:"size"`        // Photos number
	Thumb       Photo  `json:"thumb"`
	Title       string `json:"title"`   // Photo album title
	Updated     int    `json:"updated"` // Date when the album has been updated last time in Unixtime
}

// AlbumFull represents `photos_photo_album_full` API object
type AlbumFull struct {
	CanUpload        base.BoolInt `json:"can_upload"`        // Information whether current user can upload photo to the album
	CommentsDisabled base.BoolInt `json:"comments_disabled"` // Information whether album comments are disabled
	Created          int          `json:"created"`           // Date when the album has been created in Unixtime
	Description      string       `json:"description"`       // Photo album description
	ID               int          `json:"id"`                // Photo album ID
	OwnerID          int          `json:"owner_id"`          // Album owner's ID
	PrivacyComment   []string     `json:"privacy_comment"`   // Privacy comment
	PrivacyView      []string     `json:"privacy_view"`      // Privacy view
	Size             int          `json:"size"`              // Photos number
	Sizes            []PhotoSizes `json:"sizes"`
	ThumbID          int          `json:"thumb_id"`              // Thumb photo ID
	ThumbIsLast      base.BoolInt `json:"thumb_is_last"`         // Information whether the album thumb is last photo
	ThumbSrc         string       `json:"thumb_src"`             // URL of the thumb image
	Title            string       `json:"title"`                 // Photo album title
	Updated          int          `json:"updated"`               // Date when the album has been updated last time in Unixtime
	AdminsUpload     base.BoolInt `json:"upload_by_admins_only"` // Information whether only community administrators can upload photos
}

// PhotoFull represents `photos_photo_full` APi object
type PhotoFull struct {
	AccessKey  string           `json:"access_key"`  // Access key for the photo
	AlbumID    int              `json:"album_id"`    // Album ID
	CanComment base.BoolInt     `json:"can_comment"` // Information whether current user can comment the photo
	Comments   base.ObjectCount `json:"comments"`
	Date       int              `json:"date"`   // Date when uploaded
	Height     int              `json:"height"` // Original photo height
	ID         int              `json:"id"`     // Photo ID
	Images     []Image          `json:"images"`
	Latitude   float64          `json:"lat"` // Latitude
	Likes      base.Likes       `json:"likes"`
	Longitude  float64          `json:"long"`     // Longitude
	OwnerID    int              `json:"owner_id"` // Photo owner's ID
	PostID     int              `json:"post_id"`  // Post ID
	Reposts    base.ObjectCount `json:"reposts"`
	Tags       base.Object      `json:"tags"`
	Text       string           `json:"text"`    // Photo caption
	UserID     int              `json:"user_id"` // ID of the user who have uploaded the photo
	Width      int              `json:"width"`   // Original photo width
}

// PhotoFullXtrRealOffset represents `photos_photo_full_xtr_real_offset` API object
type PhotoFullXtrRealOffset struct {
	AccessKey  string              `json:"access_key"`  // Access key for the photo
	AlbumID    int                 `json:"album_id"`    // Album ID
	CanComment base.BoolInt        `json:"can_comment"` // Information whether current user can comment the photo
	Comments   base.ObjectCount    `json:"comments"`
	Date       int                 `json:"date"`   // Date when uploaded
	Height     int                 `json:"height"` // Original photo height
	Hidden     base.PropertyExists `json:"hidden"` // Returns if the photo is hidden above the wall
	ID         int                 `json:"id"`     // Photo ID
	Latitude   float64             `json:"lat"`    // Latitude
	Likes      base.Likes          `json:"likes"`
	Longitude  float64             `json:"long"`        // Longitude
	OwnerID    int                 `json:"owner_id"`    // Photo owner's ID
	Photo1280  string              `json:"photo_1280"`  // URL of image with 1280 px width
	Photo130   string              `json:"photo_130"`   // URL of image with 130 px width
	Photo2560  string              `json:"photo_2560"`  // URL of image with 2560 px width
	Photo604   string              `json:"photo_604"`   // URL of image with 604 px width
	Photo75    string              `json:"photo_75"`    // URL of image with 75 px width
	Photo807   string              `json:"photo_807"`   // URL of image with 807 px width
	PostID     int                 `json:"post_id"`     // Post ID
	RealOffset int                 `json:"real_offset"` // Real position of the photo
	Reposts    base.ObjectCount    `json:"reposts"`
	Sizes      []PhotoSizes        `json:"sizes"`
	Tags       base.Object         `json:"tags"`
	Text       string              `json:"text"`    // Photo caption
	UserID     int                 `json:"user_id"` // ID of the user who have uploaded the photo
	Width      int                 `json:"width"`   // Original photo width
}

// PhotoSizes represents `photos_photo_sizes` API object
type PhotoSizes struct {
	Height int            `json:"height"` // Height in px
	Width  int            `json:"width"`  // Width in px
	Src    string         `json:"src"`    // URL of the image
	Type   PhotoSizesType `json:"type"`
}

// PhotoSizesType represents `photos_photo_sizes_type` API object
type PhotoSizesType string

func (p *PhotoSizesType) MarshalJSON() ([]byte, error) {
	switch *p {
	case "s", "m", "x", "o", "p", "q", "r", "y", "z", "w":
		return []byte(*p), nil
	}

	return []byte{}, fmt.Errorf("value is not in valid range")
}

func (p *PhotoSizesType) GetName() string {
	return string(*p)
}

// Tag represents `photos_photo_tag` API object
type Tag struct {
	Date       int          `json:"date"`        // Date when tag has been added in Unixtime
	ID         int          `json:"id"`          // Tag ID
	OwnerID    int          `json:"placer_id"`   // ID of the tag creator
	TaggedName string       `json:"tagged_name"` // Tag description
	UserID     int          `json:"user_id"`     // Tagged user ID
	Viewed     base.BoolInt `json:"viewed"`      // Information whether the tag was reviewed
	X          int          `json:"x"`           // Coordinate X of the left upper corner
	X2         int          `json:"x2"`          // Coordinate X of the right lower corner
	Y          int          `json:"y"`           // Coordinate Y of the left upper corner
	Y2         int          `json:"y2"`          // Coordinate Y of the right lower corner
}

// Upload represents `photos_photo_upload` API response
type Upload struct {
	AlbumID   int    `json:"album_id"`   // Album ID
	UploadURL string `json:"upload_url"` // URL to upload photo
	UserID    int    `json:"user_id"`    // User ID
}

// UploadResponse represents `photos_photo_upload_response` API object
type UploadResponse struct {
	AlbumID    int    `json:"aid"`         // Album ID
	Hash       string `json:"hash"`        // Uploading hash
	PhotosList string `json:"photos_list"` // Uploaded photos data
	Server     int    `json:"server"`      // Upload server number
}

// PhotoXtrRealOffset represents `photos_photo_xtr_real_offset` API object
type PhotoXtrRealOffset struct {
	AccessKey  string              `json:"access_key"`  // Access key for the photo
	AlbumID    int                 `json:"album_id"`    // Album ID
	Date       int                 `json:"date"`        // Date when uploaded
	Height     int                 `json:"height"`      // Original photo height
	Hidden     base.PropertyExists `json:"hidden"`      // Returns if the photo is hidden above the wall
	ID         int                 `json:"id"`          // Photo ID
	Latitude   float64             `json:"lat"`         // Latitude
	Longitude  float64             `json:"long"`        // Longitude
	OwnerID    int                 `json:"owner_id"`    // Photo owner's ID
	Photo1280  string              `json:"photo_1280"`  // URL of image with 1280 px width
	Photo130   string              `json:"photo_130"`   // URL of image with 130 px width
	Photo2560  string              `json:"photo_2560"`  // URL of image with 2560 px width
	Photo604   string              `json:"photo_604"`   // URL of image with 604 px width
	Photo75    string              `json:"photo_75"`    // URL of image with 75 px width
	Photo807   string              `json:"photo_807"`   // URL of image with 807 px width
	PostID     int                 `json:"post_id"`     // Post ID
	RealOffset int                 `json:"real_offset"` // Real position of the photos
	Sizes      []PhotoSizes        `json:"sizes"`
	Text       string              `json:"text"`    // Photo caption
	UserID     int                 `json:"user_id"` // ID of the user who have uploaded the photo
	Width      int                 `json:"width"`   // Original photo width
}

// PhotoXtrTagInfo represents `photos_photo_xtr_tag_info` API object
type PhotoXtrTagInfo struct {
	AccessKey  string       `json:"access_key"`  // Access key for the photo
	AlbumID    int          `json:"album_id"`    // Album ID
	Date       int          `json:"date"`        // Date when uploaded
	Height     int          `json:"height"`      // Original photo height
	ID         int          `json:"id"`          // Photo ID
	Latitude   float64      `json:"lat"`         // Latitude
	Longitude  float64      `json:"long"`        // Longitude
	OwnerID    int          `json:"owner_id"`    // Photo owner's ID
	Photo1280  string       `json:"photo_1280"`  // URL of image with 1280 px width
	Photo130   string       `json:"photo_130"`   // URL of image with 130 px width
	Photo2560  string       `json:"photo_2560"`  // URL of image with 2560 px width
	Photo604   string       `json:"photo_604"`   // URL of image with 604 px width
	Photo75    string       `json:"photo_75"`    // URL of image with 75 px width
	Photo807   string       `json:"photo_807"`   // URL of image with 807 px width
	PlacerID   int          `json:"placer_id"`   // ID of the tag creator
	PostID     int          `json:"post_id"`     // Post ID
	RealOffset int          `json:"real_offset"` // Real position of the photos
	Sizes      []PhotoSizes `json:"sizes"`
	TagCreated int          `json:"tag_created"` // Date when tag has been added in Unixtime
	TagID      int          `json:"tag_id"`      // Tag ID
	Text       string       `json:"text"`        // Photo caption
	UserID     int          `json:"user_id"`     // ID of the user who have uploaded the photo
	Width      int          `json:"width"`       // Original photo width
}

// WallUploadResponse represents `photos_wall_upload_response` API object
type WallUploadResponse struct {
	Hash   string `json:"hash"`   // Uploading hash
	Photo  string `json:"photo"`  // Uploaded photo data
	Server int    `json:"server"` // Upload server number
}
