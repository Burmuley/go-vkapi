package base

import (
	"fmt"
	"gitlab.com/Burmuley/go-vkapi/objects/market"
	"gitlab.com/Burmuley/go-vkapi/objects/photos"
	"strings"
)

/////////////////////////////////////////////////////////////
// Base API objects                                        //
/////////////////////////////////////////////////////////////

// BoolInt represents `base_bool_int` API object
// May take values 'yes' or '1' for true and 'no' and '0' for false
type BoolInt bool

func (v *BoolInt) UnmarshalJSON(b []byte) error {
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

func (v *BoolInt) MarshalJSON() ([]byte, error) {
	if *v {
		return []byte("1"), nil
	} else if !(*v) {
		return []byte("0"), nil
	}

	return []byte{0}, fmt.Errorf("unknown value")
}

// CommentsInfo represents `base_comments_info` API object
type CommentsInfo struct {
	CanPost       BoolInt `json:"can_post"`
	Count         int     `json:"count"`
	GroupsCanPost bool    `json:"groups_can_post"`
}

// Country represents `base_country` API object
type Country struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// Error represents `base_error` API object
type Error struct {
	ErrorCode     int          `json:"error_code"`
	ErrorMessage  string       `json:"error_msg"`
	RequestParams RequestParam `json:"request_params"`
}

// Geo represents `base_geo` API object
type Geo struct {
	Coordinates GeoCoordinates `json:"coordinates"`
	Place       Place          `json:"base_place"`
	ShowMap     int            `json:"show_map"`
	Type        string         `json:"type"`
}

// GeoCoordinates represents `base_geo_coordinates` API object
type GeoCoordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Image represents `base_image` API object
type Image struct {
	Height int    `json:"height"`
	Width  int    `json:"width"`
	Url    string `json:"url"`
}

// Likes represents `base_likes` API object
type Likes struct {
	Count     int     `json:"count"`
	UserLikes BoolInt `json:"user_likes"`
}

// LikesInfo represents `base_likes_info` API object
type LikesInfo struct {
	CanLike    BoolInt `json:"can_like"`
	CanPublish BoolInt `json:"can_publish"`
	Count      int     `json:"count"`
	UserLikes  BoolInt `json:"user_likes"`
}

// Link represents `base_link` API object
type Link struct {
	Application LinkApplication `json:"application"`
	Button      LinkButton      `json:"button"`
	Caption     string          `json:"caption"`
	Description string          `json:"description"`
	ID          int             `json:"id"`
	IsFavorite  bool            `json:"is_favorite"`
	Photo       photos.Photo    `json:"photo"`
	PreviewPage string          `json:"preview_page"`
	PreviewUrl  string          `json:"preview_url"`
	Product     LinkProduct     `json:"product"`
	Rating      LinkRating      `json:"rating"`
	Title       string          `json:"title"`
	Url         string          `json:"url"`
}

// LinkApplication represents `base_link_application` API object
type LinkApplication struct {
	ID    float64              `json:"app_id"`
	Store LinkApplicationStore `json:"store"`
}

// LinkApplicationStore represents `base_link_application_store` API object
type LinkApplicationStore struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// LinkButton represents `base_link_button` API object
type LinkButton struct {
	Action LinkButtonAction `json:"action"`
	Title  string           `json:"title"`
}

// LinkButtonAction represents `base_link_button_action` API object
type LinkButtonAction struct {
	URL  string               `json:"url"`
	Type LinkButtonActionType `json:"type"`
}

// LinkButtonActionType represents `base_link_button_action` API object
type LinkButtonActionType string

// LinkProduct represents `base_link_product` API object
type LinkProduct struct {
	Price market.Price `json:"price"`
}

// LinkRating represents `base_link_rating` API object
type LinkRating struct {
	Reviews int     `json:"reviews_count"`
	Stars   float64 `json:"stars"`
}

// MessageError represents `base_message_error` API object
type MessageError struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

// Object represents `base_object` API object
type Object struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// ObjectCount represents `base_object_count` API object
type ObjectCount struct {
	Count int `json:"count"`
}

// ObjectWithName represents `base_object_with_name` API object
type ObjectWithName struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// OkResponse represents `base_ok_response` API object
type OkResponse int

func (b *OkResponse) GetName() string {
	switch *b {
	case 1:
		return "ok"
	default:
		return ""
	}
}

//Place represents `base_place` coordinate API object
type Place struct {
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
type PropertyExists int

func (b *PropertyExists) GetName() string {
	switch *b {
	case 1:
		return "Property exists"
	default:
		return ""
	}
}

// RepostsInfo represents `base_reposts_info` API object
type RepostsInfo struct {
	Count    int `json:"count"`
	Reposted int `json:"user_reposted"`
}

//BaseRequestParams represents `base_request_params` API object
type RequestParam struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Sex represents `base_sex` API object
type Sex int

func (b *Sex) GetName() string {
	switch *b {
	case 1:
		return "female"
	case 2:
		return "male"
	default:
		return "unknown"
	}
}

// BaseSticker represents `base_sticker` API object
type BaseSticker struct {
	Images    []Image `json:"images"`
	ImagesBck []Image `json:"images_with_background"`
	ProductID int     `json:"product_id"`
	StickerID int     `json:"sticker_id"`
}

// UploadServer represents `base_upload_server` API object
type UploadServer struct {
	UploadURL string `json:"upload_url"`
}

// UserGroupFields represents `base_user_group_fields` API object
type UserGroupFields string

// UserID represents `base_user_id` API object
type UserID struct {
	ID int `json:"user_id"`
}
