package objects

/////////////////////////////////////////////////////////////
// Photos related API objects	                           //
/////////////////////////////////////////////////////////////

// CommentXtrPid represents `photos_comment_xtr_pid` API object
type PhotosCommentXtrPid struct {
	Attachments    []WallCommentAttachment `json:"attachments"`
	Date           int                     `json:"date"`    // Date when the comment has been added in Unixtime
	FromID         int                     `json:"from_id"` // Author ID
	ID             int                     `json:"id"`      // Comment ID
	Likes          BaseLikesInfo           `json:"likes"`
	PhotoID        int                     `json:"pid"`              // Photo ID
	ReplyToComment int                     `json:"reply_to_comment"` // Replied comment ID
	ReplyToUser    int                     `json:"reply_to_user"`    // Replied user ID
	Text           string                  `json:"text"`             // Comment text
}

// Image represents `photos_image` API object
type PhotosImage struct {
	Height int             `json:"height"` // Height of the photo in px.
	Url    string          `json:"url"`    // Photo URL
	Type   PhotosImageType `json:"type"`
	Width  int             `json:"width"` // Width of the photo in px.
}

// ImageType represents `photos_image_type` API object
type PhotosImageType string

func (i *PhotosImageType) MarshalJSON() ([]byte, error) {
	return GetStringFromRange(string(*i), "s", "m", "x", "o", "p", "q", "r", "y", "z", "w")
}

func (i *PhotosImageType) String() string {
	return string(*i)
}

// MarketAlbumUploadResponse represents `photos_market_album_upload_response` API object
type PhotosMarketAlbumUploadResponse struct {
	GroupID int    `json:"gid"`    // Community ID
	Hash    string `json:"hash"`   // Uploading hash
	Photo   string `json:"photo"`  // Uploaded photo data
	Server  int    `json:"server"` // Upload server number
}

// MarketUploadResponse represents `photos_market_upload_response` API object
type PhotosMarketUploadResponse struct {
	CropData string `json:"crop_data"` // Crop data
	CropHash string `json:"crop_hash"` // Crop hash
	GroupID  int    `json:"group_id"`  // Community ID
	Hash     string `json:"hash"`      // Uploading hash
	Photo    string `json:"photo"`     // Uploaded photo data
	Server   int    `json:"server"`    // Upload server number
}

// MessageUploadResponse represents `photos_message_upload_response` API object
type PhotosMessageUploadResponse struct {
	Hash   string `json:"hash"`   // Uploading hash
	Photo  string `json:"photo"`  // Uploaded photo data
	Server int    `json:"server"` // Upload server number
}

// OwnerUploadResponse represents `photos_owner_upload_response` API object
type PhotosOwnerUploadResponse struct {
	Hash   string `json:"hash"`   // Uploading hash
	Photo  string `json:"photo"`  // Uploaded photo data
	Server int    `json:"server"` // Upload server number
}

// Photo represents `photos_photo` APi object
type PhotosPhoto struct {
	AccessKey string        `json:"access_key"` // Access key for the photo
	AlbumID   int           `json:"album_id"`   // Album ID
	Date      int           `json:"date"`       // Date when uploaded
	Height    int           `json:"height"`     // Original photo height
	ID        int           `json:"id"`         // Photo ID
	Images    []PhotosImage `json:"images"`
	Latitude  float64       `json:"lat"`      // Latitude
	Longitude float64       `json:"long"`     // Longitude
	OwnerID   int           `json:"owner_id"` // Photo owner's ID
	PostID    int           `json:"post_id"`  // Post ID
	Text      string        `json:"text"`     // Photo caption
	UserID    int           `json:"user_id"`  // ID of the user who have uploaded the photo
	Width     int           `json:"width"`    // Original photo width
}

// Album represents `photos_photo_album` API object
type PhotosAlbum struct {
	Created     int         `json:"created"`     // Date when the album has been created in Unixtime
	Description string      `json:"description"` // Photo album description
	ID          int         `json:"id"`          // Photo album ID
	OwnerID     int         `json:"owner_id"`    // Album owner's ID
	Size        int         `json:"size"`        // Photos number
	Thumb       PhotosPhoto `json:"thumb"`
	Title       string      `json:"title"`   // Photo album title
	Updated     int         `json:"updated"` // Date when the album has been updated last time in Unixtime
}

// AlbumFull represents `photos_photo_album_full` API object
type PhotosAlbumFull struct {
	CanUpload        BaseBoolInt        `json:"can_upload"`        // Information whether current user can upload photo to the album
	CommentsDisabled BaseBoolInt        `json:"comments_disabled"` // Information whether album comments are disabled
	Created          int                `json:"created"`           // Date when the album has been created in Unixtime
	Description      string             `json:"description"`       // Photo album description
	ID               int                `json:"id"`                // Photo album ID
	OwnerID          int                `json:"owner_id"`          // Album owner's ID
	PrivacyComment   []string           `json:"privacy_comment"`   // Privacy comment
	PrivacyView      []string           `json:"privacy_view"`      // Privacy view
	Size             int                `json:"size"`              // Photos number
	Sizes            []PhotosPhotoSizes `json:"sizes"`
	ThumbID          int                `json:"thumb_id"`              // Thumb photo ID
	ThumbIsLast      BaseBoolInt        `json:"thumb_is_last"`         // Information whether the album thumb is last photo
	ThumbSrc         string             `json:"thumb_src"`             // URL of the thumb image
	Title            string             `json:"title"`                 // Photo album title
	Updated          int                `json:"updated"`               // Date when the album has been updated last time in Unixtime
	AdminsUpload     BaseBoolInt        `json:"upload_by_admins_only"` // Information whether only community administrators can upload photos
}

// PhotoFull represents `photos_photo_full` APi object
type PhotosPhotoFull struct {
	AccessKey  string          `json:"access_key"`  // Access key for the photo
	AlbumID    int             `json:"album_id"`    // Album ID
	CanComment BaseBoolInt     `json:"can_comment"` // Information whether current user can comment the photo
	Comments   BaseObjectCount `json:"comments"`
	Date       int             `json:"date"`   // Date when uploaded
	Height     int             `json:"height"` // Original photo height
	ID         int             `json:"id"`     // Photo ID
	Images     []PhotosImage   `json:"images"`
	Latitude   float64         `json:"lat"` // Latitude
	Likes      BaseLikes       `json:"likes"`
	Longitude  float64         `json:"long"`     // Longitude
	OwnerID    int             `json:"owner_id"` // Photo owner's ID
	PostID     int             `json:"post_id"`  // Post ID
	Reposts    BaseObjectCount `json:"reposts"`
	Tags       BaseObject      `json:"tags"`
	Text       string          `json:"text"`    // Photo caption
	UserID     int             `json:"user_id"` // ID of the user who have uploaded the photo
	Width      int             `json:"width"`   // Original photo width
}

// PhotoFullXtrRealOffset represents `photos_photo_full_xtr_real_offset` API object
type PhotosPhotoFullXtrRealOffset struct {
	AccessKey  string             `json:"access_key"`  // Access key for the photo
	AlbumID    int                `json:"album_id"`    // Album ID
	CanComment BaseBoolInt        `json:"can_comment"` // Information whether current user can comment the photo
	Comments   BaseObjectCount    `json:"comments"`
	Date       int                `json:"date"`   // Date when uploaded
	Height     int                `json:"height"` // Original photo height
	Hidden     BasePropertyExists `json:"hidden"` // Returns if the photo is hidden above the wall
	ID         int                `json:"id"`     // Photo ID
	Latitude   float64            `json:"lat"`    // Latitude
	Likes      BaseLikes          `json:"likes"`
	Longitude  float64            `json:"long"`        // Longitude
	OwnerID    int                `json:"owner_id"`    // Photo owner's ID
	Photo1280  string             `json:"photo_1280"`  // URL of image with 1280 px width
	Photo130   string             `json:"photo_130"`   // URL of image with 130 px width
	Photo2560  string             `json:"photo_2560"`  // URL of image with 2560 px width
	Photo604   string             `json:"photo_604"`   // URL of image with 604 px width
	Photo75    string             `json:"photo_75"`    // URL of image with 75 px width
	Photo807   string             `json:"photo_807"`   // URL of image with 807 px width
	PostID     int                `json:"post_id"`     // Post ID
	RealOffset int                `json:"real_offset"` // Real position of the photo
	Reposts    BaseObjectCount    `json:"reposts"`
	Sizes      []PhotosPhotoSizes `json:"sizes"`
	Tags       BaseObject         `json:"tags"`
	Text       string             `json:"text"`    // Photo caption
	UserID     int                `json:"user_id"` // ID of the user who have uploaded the photo
	Width      int                `json:"width"`   // Original photo width
}

// PhotoSizes represents `photos_photo_sizes` API object
type PhotosPhotoSizes struct {
	Height int                  `json:"height"` // Height in px
	Width  int                  `json:"width"`  // Width in px
	Src    string               `json:"src"`    // URL of the image
	Type   PhotosPhotoSizesType `json:"type"`
}

// PhotoSizesType represents `photos_photo_sizes_type` API object
type PhotosPhotoSizesType string

func (p *PhotosPhotoSizesType) MarshalJSON() ([]byte, error) {
	return GetStringFromRange(string(*p), "s", "m", "x", "o", "p", "q", "r", "y", "z", "w")
}

func (p *PhotosPhotoSizesType) String() string {
	return string(*p)
}

// Tag represents `photos_photo_tag` API object
type PhotosTag struct {
	Date       int         `json:"date"`        // Date when tag has been added in Unixtime
	ID         int         `json:"id"`          // Tag ID
	OwnerID    int         `json:"placer_id"`   // ID of the tag creator
	TaggedName string      `json:"tagged_name"` // Tag description
	UserID     int         `json:"user_id"`     // Tagged user ID
	Viewed     BaseBoolInt `json:"viewed"`      // Information whether the tag was reviewed
	X          int         `json:"x"`           // Coordinate X of the left upper corner
	X2         int         `json:"x2"`          // Coordinate X of the right lower corner
	Y          int         `json:"y"`           // Coordinate Y of the left upper corner
	Y2         int         `json:"y2"`          // Coordinate Y of the right lower corner
}

// Upload represents `photos_photo_upload` API response
type PhotosUpload struct {
	AlbumID   int    `json:"album_id"`   // Album ID
	UploadURL string `json:"upload_url"` // URL to upload photo
	UserID    int    `json:"user_id"`    // User ID
}

// UploadResponse represents `photos_photo_upload_response` API object
type PhotosUploadResponse struct {
	AlbumID    int    `json:"aid"`         // Album ID
	Hash       string `json:"hash"`        // Uploading hash
	PhotosList string `json:"photos_list"` // Uploaded photos data
	Server     int    `json:"server"`      // Upload server number
}

// PhotoXtrRealOffset represents `photos_photo_xtr_real_offset` API object
type PhotosPhotoXtrRealOffset struct {
	AccessKey  string             `json:"access_key"`  // Access key for the photo
	AlbumID    int                `json:"album_id"`    // Album ID
	Date       int                `json:"date"`        // Date when uploaded
	Height     int                `json:"height"`      // Original photo height
	Hidden     BasePropertyExists `json:"hidden"`      // Returns if the photo is hidden above the wall
	ID         int                `json:"id"`          // Photo ID
	Latitude   float64            `json:"lat"`         // Latitude
	Longitude  float64            `json:"long"`        // Longitude
	OwnerID    int                `json:"owner_id"`    // Photo owner's ID
	Photo1280  string             `json:"photo_1280"`  // URL of image with 1280 px width
	Photo130   string             `json:"photo_130"`   // URL of image with 130 px width
	Photo2560  string             `json:"photo_2560"`  // URL of image with 2560 px width
	Photo604   string             `json:"photo_604"`   // URL of image with 604 px width
	Photo75    string             `json:"photo_75"`    // URL of image with 75 px width
	Photo807   string             `json:"photo_807"`   // URL of image with 807 px width
	PostID     int                `json:"post_id"`     // Post ID
	RealOffset int                `json:"real_offset"` // Real position of the photos
	Sizes      []PhotosPhotoSizes `json:"sizes"`
	Text       string             `json:"text"`    // Photo caption
	UserID     int                `json:"user_id"` // ID of the user who have uploaded the photo
	Width      int                `json:"width"`   // Original photo width
}

// PhotoXtrTagInfo represents `photos_photo_xtr_tag_info` API object
type PhotosPhotoXtrTagInfo struct {
	AccessKey  string             `json:"access_key"`  // Access key for the photo
	AlbumID    int                `json:"album_id"`    // Album ID
	Date       int                `json:"date"`        // Date when uploaded
	Height     int                `json:"height"`      // Original photo height
	ID         int                `json:"id"`          // Photo ID
	Latitude   float64            `json:"lat"`         // Latitude
	Longitude  float64            `json:"long"`        // Longitude
	OwnerID    int                `json:"owner_id"`    // Photo owner's ID
	Photo1280  string             `json:"photo_1280"`  // URL of image with 1280 px width
	Photo130   string             `json:"photo_130"`   // URL of image with 130 px width
	Photo2560  string             `json:"photo_2560"`  // URL of image with 2560 px width
	Photo604   string             `json:"photo_604"`   // URL of image with 604 px width
	Photo75    string             `json:"photo_75"`    // URL of image with 75 px width
	Photo807   string             `json:"photo_807"`   // URL of image with 807 px width
	PlacerID   int                `json:"placer_id"`   // ID of the tag creator
	PostID     int                `json:"post_id"`     // Post ID
	RealOffset int                `json:"real_offset"` // Real position of the photos
	Sizes      []PhotosPhotoSizes `json:"sizes"`
	TagCreated int                `json:"tag_created"` // Date when tag has been added in Unixtime
	TagID      int                `json:"tag_id"`      // Tag ID
	Text       string             `json:"text"`        // Photo caption
	UserID     int                `json:"user_id"`     // ID of the user who have uploaded the photo
	Width      int                `json:"width"`       // Original photo width
}

// WallUploadResponse represents `photos_wall_upload_response` API object
type PhotosWallUploadResponse struct {
	Hash   string `json:"hash"`   // Uploading hash
	Photo  string `json:"photo"`  // Uploaded photo data
	Server int    `json:"server"` // Upload server number
}
