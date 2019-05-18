package objects

import (
	"fmt"
	"strings"
)

/////////////////////////////////////////////////////////////
// Base API objects                                        //
/////////////////////////////////////////////////////////////

// BoolInt represents `base_bool_int` API object
// May take values 'yes' or '1' for true and 'no' and '0' for false
type BaseBoolInt bool

func (v *BaseBoolInt) UnmarshalJSON(b []byte) error {
	strValue := strings.ToLower(string(b))

	if strValue == "yes" || strValue == "1" {
		*v = true
	} else if strValue == "no" || strValue == "0" {
		*v = false
	} else {
		return fmt.Errorf("unknown value")
	}
	return nil
}

func (v *BaseBoolInt) MarshalJSON() ([]byte, error) {
	if *v {
		return []byte("1"), nil
	} else if !(*v) {
		return []byte("0"), nil
	}

	return []byte{0}, fmt.Errorf("unknown value")
}

// CommentsInfo represents `base_comments_info` API object
type BaseCommentsInfo struct {
	CanPost       BaseBoolInt `json:"can_post"`
	Count         int         `json:"count"`
	GroupsCanPost bool        `json:"groups_can_post"`
}

// Country represents `base_country` API object
type BaseCountry struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// Error represents `base_error` API object
type BaseError struct {
	ErrorCode     int              `json:"error_code"`
	ErrorMessage  string           `json:"error_msg"`
	RequestParams BaseRequestParam `json:"request_params"`
}

// Geo represents `base_geo` API object
type BaseGeo struct {
	Coordinates BaseGeoCoordinates `json:"coordinates"`
	Place       BasePlace          `json:"base_place"`
	ShowMap     int                `json:"show_map"`
	Type        string             `json:"type"`
}

// GeoCoordinates represents `base_geo_coordinates` API object
type BaseGeoCoordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Image represents `base_image` API object
type BaseImage struct {
	Height int    `json:"height"`
	Width  int    `json:"width"`
	Url    string `json:"url"`
}

// Likes represents `base_likes` API object
type BaseLikes struct {
	Count     int         `json:"count"`
	UserLikes BaseBoolInt `json:"user_likes"`
}

// LikesInfo represents `base_likes_info` API object
type BaseLikesInfo struct {
	CanLike    BaseBoolInt `json:"can_like"`
	CanPublish BaseBoolInt `json:"can_publish"`
	Count      int         `json:"count"`
	UserLikes  BaseBoolInt `json:"user_likes"`
}

// Link represents `base_link` API object
type BaseLink struct {
	Application BaseLinkApplication `json:"application"`
	Button      BaseLinkButton      `json:"button"`
	Caption     string              `json:"caption"`
	Description string              `json:"description"`
	ID          int                 `json:"id"`
	IsFavorite  bool                `json:"is_favorite"`
	Photo       PhotosPhoto         `json:"photo"`
	PreviewPage string              `json:"preview_page"`
	PreviewUrl  string              `json:"preview_url"`
	Product     BaseLinkProduct     `json:"product"`
	Rating      BaseLinkRating      `json:"rating"`
	Title       string              `json:"title"`
	Url         string              `json:"url"`
}

// LinkApplication represents `base_link_application` API object
type BaseLinkApplication struct {
	ID    float64                  `json:"app_id"`
	Store BaseLinkApplicationStore `json:"store"`
}

// LinkApplicationStore represents `base_link_application_store` API object
type BaseLinkApplicationStore struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// LinkButton represents `base_link_button` API object
type BaseLinkButton struct {
	Action BaseLinkButtonAction `json:"action"`
	Title  string               `json:"title"`
}

// LinkButtonAction represents `base_link_button_action` API object
type BaseLinkButtonAction struct {
	URL  string                   `json:"url"`
	Type BaseLinkButtonActionType `json:"type"`
}

// LinkButtonActionType represents `base_link_button_action` API object
type BaseLinkButtonActionType string

// LinkProduct represents `base_link_product` API object
type BaseLinkProduct struct {
	Price MarketPrice `json:"price"`
}

// LinkRating represents `base_link_rating` API object
type BaseLinkRating struct {
	Reviews int     `json:"reviews_count"`
	Stars   float64 `json:"stars"`
}

// MessageError represents `base_message_error` API object
type BaseMessageError struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

// Object represents `base_object` API object
type BaseObject struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// ObjectCount represents `base_object_count` API object
type BaseObjectCount struct {
	Count int `json:"count"`
}

// ObjectWithName represents `base_object_with_name` API object
type BaseObjectWithName struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// OkResponse represents `base_ok_response` API object
type BaseOkResponse int

func (b *BaseOkResponse) GetName() string {
	switch *b {
	case 1:
		return "ok"
	default:
		return ""
	}
}

//Place represents `base_place` coordinate API object
type BasePlace struct {
	Address   string  `json:"address"`
	Checkins  int     `json:"checkins"`
	City      string  `json:"city"`
	Country   string  `json:"country"`
	Created   int     `json:"created"`
	Icon      string  `json:"icon"`
	ID        int     `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Title     string  `json:"title"`
	Type      string  `json:"type"`
}

// PropertyExists represents `base_property_exists` API object
type BasePropertyExists int

func (b *BasePropertyExists) GetName() string {
	switch *b {
	case 1:
		return "Property exists"
	default:
		return ""
	}
}

// RepostsInfo represents `base_reposts_info` API object
type BaseRepostsInfo struct {
	Count    int `json:"count"`
	Reposted int `json:"user_reposted"`
}

//BaseRequestParams represents `base_request_params` API object
type BaseRequestParam struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Sex represents `base_sex` API object
type BaseSex int

func (b *BaseSex) GetName() string {
	switch *b {
	case 1:
		return "female"
	case 2:
		return "male"
	default:
		return "unknown"
	}
}

// Sticker represents `base_sticker` API object
type BaseSticker struct {
	Images    []BaseImage `json:"images"`
	ImagesBck []BaseImage `json:"images_with_background"`
	ProductID int         `json:"product_id"`
	StickerID int         `json:"sticker_id"`
}

// UploadServer represents `base_upload_server` API object
type BaseUploadServer struct {
	UploadURL string `json:"upload_url"`
}

// UserGroupFields represents `base_user_group_fields` API object
type BaseUserGroupFields string

// UserID represents `base_user_id` API object
type BaseUserID struct {
	ID int `json:"user_id"`
}