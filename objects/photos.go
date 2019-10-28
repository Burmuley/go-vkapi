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
// `photos` group of objects
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

// PhotosCommentXtrPid type represents `photos_comment_xtr_pid` API object
type PhotosCommentXtrPid struct {
	Attachments    []WallCommentAttachment `json:"attachments"`
	Date           int                     `json:"date"`    // Date when the comment has been added in Unixtime
	FromId         int                     `json:"from_id"` // Author ID
	Id             int                     `json:"id"`      // Comment ID
	Likes          BaseLikesInfo           `json:"likes"`
	ParentsStack   []int                   `json:"parents_stack"`
	Pid            int                     `json:"pid"`              // Photo ID
	ReplyToComment int                     `json:"reply_to_comment"` // Replied comment ID
	ReplyToUser    int                     `json:"reply_to_user"`    // Replied user ID
	Text           string                  `json:"text"`             // Comment text
	Thread         CommentThread           `json:"thread"`
}

// PhotosImage type represents `photos_image` API object
type PhotosImage struct {
	Height int             `json:"height"` // Height of the photo in px.
	Type   PhotosImageType `json:"type"`
	Url    string          `json:"url"`   // Photo URL.
	Width  int             `json:"width"` // Width of the photo in px.
}

// PhotosImageType type represents `photos_image_type` API object
type PhotosImageType string // Photo's type.

// PhotosMarketAlbumUploadResponse type represents `photos_market_album_upload_response` API object
type PhotosMarketAlbumUploadResponse struct {
	Gid    int    `json:"gid"`    // Community ID
	Hash   string `json:"hash"`   // Uploading hash
	Photo  string `json:"photo"`  // Uploaded photo data
	Server int    `json:"server"` // Upload server number
}

// PhotosMarketUploadResponse type represents `photos_market_upload_response` API object
type PhotosMarketUploadResponse struct {
	CropData string `json:"crop_data"` // Crop data
	CropHash string `json:"crop_hash"` // Crop hash
	GroupId  int    `json:"group_id"`  // Community ID
	Hash     string `json:"hash"`      // Uploading hash
	Photo    string `json:"photo"`     // Uploaded photo data
	Server   int    `json:"server"`    // Upload server number
}

// PhotosMessageUploadResponse type represents `photos_message_upload_response` API object
type PhotosMessageUploadResponse struct {
	Hash   string `json:"hash"`   // Uploading hash
	Photo  string `json:"photo"`  // Uploaded photo data
	Server int    `json:"server"` // Upload server number
}

// PhotosOwnerUploadResponse type represents `photos_owner_upload_response` API object
type PhotosOwnerUploadResponse struct {
	Hash   string `json:"hash"`   // Uploading hash
	Photo  string `json:"photo"`  // Uploaded photo data
	Server int    `json:"server"` // Upload server number
}

// PhotosPhoto type represents `photos_photo` API object
type PhotosPhoto struct {
	AccessKey string             `json:"access_key"` // Access key for the photo
	AlbumId   int                `json:"album_id"`   // Album ID
	Date      int                `json:"date"`       // Date when uploaded
	Height    int                `json:"height"`     // Original photo height
	Id        int                `json:"id"`         // Photo ID
	Images    []PhotosImage      `json:"images"`
	Lat       float64            `json:"lat"`      // Latitude
	Long      float64            `json:"long"`     // Longitude
	OwnerId   int                `json:"owner_id"` // Photo owner's ID
	PostId    int                `json:"post_id"`  // Post ID
	Sizes     []PhotosPhotoSizes `json:"sizes"`
	Text      string             `json:"text"`    // Photo caption
	UserId    int                `json:"user_id"` // ID of the user who have uploaded the photo
	Width     int                `json:"width"`   // Original photo width
}

// PhotosPhotoAlbum type represents `photos_photo_album` API object
type PhotosPhotoAlbum struct {
	Created     int         `json:"created"`     // Date when the album has been created in Unixtime
	Description string      `json:"description"` // Photo album description
	Id          int         `json:"id"`          // Photo album ID
	OwnerId     int         `json:"owner_id"`    // Album owner's ID
	Size        int         `json:"size"`        // Photos number
	Thumb       PhotosPhoto `json:"thumb"`
	Title       string      `json:"title"`   // Photo album title
	Updated     int         `json:"updated"` // Date when the album has been updated last time in Unixtime
}

// PhotosPhotoAlbumFull type represents `photos_photo_album_full` API object
type PhotosPhotoAlbumFull struct {
	CanUpload          BaseBoolInt        `json:"can_upload"`        // Information whether current user can upload photo to the album
	CommentsDisabled   BaseBoolInt        `json:"comments_disabled"` // Information whether album comments are disabled
	Created            int                `json:"created"`           // Date when the album has been created in Unixtime
	Description        string             `json:"description"`       // Photo album description
	Id                 int                `json:"id"`                // Photo album ID
	OwnerId            int                `json:"owner_id"`          // Album owner's ID
	Size               int                `json:"size"`              // Photos number
	Sizes              []PhotosPhotoSizes `json:"sizes"`
	ThumbId            int                `json:"thumb_id"`              // Thumb photo ID
	ThumbIsLast        BaseBoolInt        `json:"thumb_is_last"`         // Information whether the album thumb is last photo
	ThumbSrc           string             `json:"thumb_src"`             // URL of the thumb image
	Title              string             `json:"title"`                 // Photo album title
	Updated            int                `json:"updated"`               // Date when the album has been updated last time in Unixtime
	UploadByAdminsOnly BaseBoolInt        `json:"upload_by_admins_only"` // Information whether only community administrators can upload photos
}

// PhotosPhotoFull type represents `photos_photo_full` API object
type PhotosPhotoFull struct {
	AccessKey  string          `json:"access_key"`  // Access key for the photo
	AlbumId    int             `json:"album_id"`    // Album ID
	CanComment BaseBoolInt     `json:"can_comment"` // Information whether current user can comment the photo
	Comments   BaseObjectCount `json:"comments"`
	Date       int             `json:"date"`   // Date when uploaded
	Height     int             `json:"height"` // Original photo height
	Id         int             `json:"id"`     // Photo ID
	Images     []PhotosImage   `json:"images"`
	Lat        float64         `json:"lat"` // Latitude
	Likes      BaseLikes       `json:"likes"`
	Long       float64         `json:"long"`     // Longitude
	OwnerId    int             `json:"owner_id"` // Photo owner's ID
	PostId     int             `json:"post_id"`  // Post ID
	Reposts    BaseObjectCount `json:"reposts"`
	Tags       BaseObjectCount `json:"tags"`
	Text       string          `json:"text"`    // Photo caption
	UserId     int             `json:"user_id"` // ID of the user who have uploaded the photo
	Width      int             `json:"width"`   // Original photo width
}

// PhotosPhotoFullXtrRealOffset type represents `photos_photo_full_xtr_real_offset` API object
type PhotosPhotoFullXtrRealOffset struct {
	AccessKey  string             `json:"access_key"` // Access key for the photo
	AlbumId    int                `json:"album_id"`   // Album ID
	CanComment BaseBoolInt        `json:"can_comment"`
	Comments   BaseObjectCount    `json:"comments"`
	Date       int                `json:"date"`   // Date when uploaded
	Height     int                `json:"height"` // Original photo height
	Hidden     BasePropertyExists `json:"hidden"` // Returns if the photo is hidden above the wall
	Id         int                `json:"id"`     // Photo ID
	Lat        float64            `json:"lat"`    // Latitude
	Likes      BaseLikes          `json:"likes"`
	Long       float64            `json:"long"`        // Longitude
	OwnerId    int                `json:"owner_id"`    // Photo owner's ID
	Photo1280  string             `json:"photo_1280"`  // URL of image with 1280 px width
	Photo130   string             `json:"photo_130"`   // URL of image with 130 px width
	Photo2560  string             `json:"photo_2560"`  // URL of image with 2560 px width
	Photo604   string             `json:"photo_604"`   // URL of image with 604 px width
	Photo75    string             `json:"photo_75"`    // URL of image with 75 px width
	Photo807   string             `json:"photo_807"`   // URL of image with 807 px width
	PostId     int                `json:"post_id"`     // Post ID
	RealOffset int                `json:"real_offset"` // Real position of the photo
	Reposts    BaseObjectCount    `json:"reposts"`
	Sizes      []PhotosPhotoSizes `json:"sizes"`
	Tags       BaseObjectCount    `json:"tags"`
	Text       string             `json:"text"`    // Photo caption
	UserId     int                `json:"user_id"` // ID of the user who have uploaded the photo
	Width      int                `json:"width"`   // Original photo width
}

// PhotosPhotoSizes type represents `photos_photo_sizes` API object
type PhotosPhotoSizes struct {
	Height int                  `json:"height"` // Height in px
	Type   PhotosPhotoSizesType `json:"type"`
	Url    string               `json:"url"`   // URL of the image
	Width  int                  `json:"width"` // Width in px
}

// PhotosPhotoSizesType type represents `photos_photo_sizes_type` API object
type PhotosPhotoSizesType string // Size type

// PhotosPhotoTag type represents `photos_photo_tag` API object
type PhotosPhotoTag struct {
	Date       int         `json:"date"`        // Date when tag has been added in Unixtime
	Id         int         `json:"id"`          // Tag ID
	PlacerId   int         `json:"placer_id"`   // ID of the tag creator
	TaggedName string      `json:"tagged_name"` // Tag description
	UserId     int         `json:"user_id"`     // Tagged user ID
	Viewed     BaseBoolInt `json:"viewed"`      // Information whether the tag is reviewed
	X          float64     `json:"x"`           // Coordinate X of the left upper corner
	X2         float64     `json:"x2"`          // Coordinate X of the right lower corner
	Y          float64     `json:"y"`           // Coordinate Y of the left upper corner
	Y2         float64     `json:"y2"`          // Coordinate Y of the right lower corner
}

// PhotosPhotoUpload type represents `photos_photo_upload` API object
type PhotosPhotoUpload struct {
	AlbumId   int    `json:"album_id"`   // Album ID
	UploadUrl string `json:"upload_url"` // URL to upload photo
	UserId    int    `json:"user_id"`    // User ID
}

// PhotosPhotoUploadResponse type represents `photos_photo_upload_response` API object
type PhotosPhotoUploadResponse struct {
	Aid        int    `json:"aid"`         // Album ID
	Hash       string `json:"hash"`        // Uploading hash
	PhotosList string `json:"photos_list"` // Uploaded photos data
	Server     int    `json:"server"`      // Upload server number
}

// PhotosPhotoXtrRealOffset type represents `photos_photo_xtr_real_offset` API object
type PhotosPhotoXtrRealOffset struct {
	AccessKey  string             `json:"access_key"`  // Access key for the photo
	AlbumId    int                `json:"album_id"`    // Album ID
	Date       int                `json:"date"`        // Date when uploaded
	Height     int                `json:"height"`      // Original photo height
	Hidden     BasePropertyExists `json:"hidden"`      // Returns if the photo is hidden above the wall
	Id         int                `json:"id"`          // Photo ID
	Lat        float64            `json:"lat"`         // Latitude
	Long       float64            `json:"long"`        // Longitude
	OwnerId    int                `json:"owner_id"`    // Photo owner's ID
	Photo1280  string             `json:"photo_1280"`  // URL of image with 1280 px width
	Photo130   string             `json:"photo_130"`   // URL of image with 130 px width
	Photo2560  string             `json:"photo_2560"`  // URL of image with 2560 px width
	Photo604   string             `json:"photo_604"`   // URL of image with 604 px width
	Photo75    string             `json:"photo_75"`    // URL of image with 75 px width
	Photo807   string             `json:"photo_807"`   // URL of image with 807 px width
	PostId     int                `json:"post_id"`     // Post ID
	RealOffset int                `json:"real_offset"` // Real position of the photo
	Sizes      []PhotosPhotoSizes `json:"sizes"`
	Text       string             `json:"text"`    // Photo caption
	UserId     int                `json:"user_id"` // ID of the user who have uploaded the photo
	Width      int                `json:"width"`   // Original photo width
}

// PhotosPhotoXtrTagInfo type represents `photos_photo_xtr_tag_info` API object
type PhotosPhotoXtrTagInfo struct {
	AccessKey  string             `json:"access_key"` // Access key for the photo
	AlbumId    int                `json:"album_id"`   // Album ID
	Date       int                `json:"date"`       // Date when uploaded
	Height     int                `json:"height"`     // Original photo height
	Id         int                `json:"id"`         // Photo ID
	Lat        float64            `json:"lat"`        // Latitude
	Long       float64            `json:"long"`       // Longitude
	OwnerId    int                `json:"owner_id"`   // Photo owner's ID
	Photo1280  string             `json:"photo_1280"` // URL of image with 1280 px width
	Photo130   string             `json:"photo_130"`  // URL of image with 130 px width
	Photo2560  string             `json:"photo_2560"` // URL of image with 2560 px width
	Photo604   string             `json:"photo_604"`  // URL of image with 604 px width
	Photo75    string             `json:"photo_75"`   // URL of image with 75 px width
	Photo807   string             `json:"photo_807"`  // URL of image with 807 px width
	PlacerId   int                `json:"placer_id"`  // ID of the tag creator
	PostId     int                `json:"post_id"`    // Post ID
	Sizes      []PhotosPhotoSizes `json:"sizes"`
	TagCreated int                `json:"tag_created"` // Date when tag has been added in Unixtime
	TagId      int                `json:"tag_id"`      // Tag ID
	Text       string             `json:"text"`        // Photo caption
	UserId     int                `json:"user_id"`     // ID of the user who have uploaded the photo
	Width      int                `json:"width"`       // Original photo width
}

// PhotosWallUploadResponse type represents `photos_wall_upload_response` API object
type PhotosWallUploadResponse struct {
	Hash   string `json:"hash"`   // Uploading hash
	Photo  string `json:"photo"`  // Uploaded photo data
	Server int    `json:"server"` // Upload server number
}
