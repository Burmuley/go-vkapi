package go_vkapi

import (
	"encoding/json"
	"fmt"
	"strings"
)

//###########################################################
// Base API objects
//###########################################################

// BaseBoolInt represents `base_bool_int` API object
// May take values 'yes' or '1' for true and 'no' and '0' for false
type BaseBoolInt bool

func (*BaseBoolInt) ToJson(i interface{}) ([]byte, error) {
	return json.Marshal(i)
}

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

// BaseObject represents `base_object` API object
type BaseObject struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

//BaseCommentsInfo represents `base_comments_info` API object
type BaseCommentsInfo struct {
	CanPost       BaseBoolInt `json:"can_post"`
	Count         int         `json:"count"`
	GroupsCanPost bool        `json:"groups_can_post"`
}

// BaseCountry represents `base_country` API object
type BaseCountry struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

//BaseError represents `base_error` API object
type BaseError struct {
	ErrorCode     int              `json:"error_code"`
	ErrorMessage  string           `json:"error_msg"`
	RequestParams BaseRequestParam `json:"request_params"`
}

//BaseGeo represents `base_geo` API object
type BaseGeo struct {
	Coordinates BaseGeoCoordinates `json:"coordinates"`
	Place       BasePlace          `json:"base_place"`
	ShowMap     int                `json:"show_map"`
	Type        string             `json:"type"'`
}

//BaseGeoCoordinates represents `base_geo_coordinates` API object
type BaseGeoCoordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

//BaseImage represents `base_image` API object
type BaseImage struct {
	Height int    `json:"height"`
	Width  int    `json:"width"`
	Url    string `json:url`
}

//BaseLikes represents `base_likes` API object
type BaseLikes struct {
	Count     int         `json:"count"`
	UserLikes BaseBoolInt `json:"user_likes"`
}

//BaseLikesInfo represents `base_likes_info` API object
type BaseLikesInfo struct {
	CanLike    BaseBoolInt `json:"can_like"`
	CanPublish BaseBoolInt `json:"can_publish"`
	Count      int         `json:"count"`
	UserLikes  BaseBoolInt `json:"user_likes"`
}

//BaseLink represents `base_link` API object
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

//BaseLinkApplication represents `base_link_application` API object
type BaseLinkApplication struct {
}

type BaseLinkButton struct {
}

type BaseLinkRating struct {
}

type BaseLinkProduct struct {
}

//BasePlace represents `base_place` coordinate API object
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

//BaseRequestParams represents `base_request_params` API object
type BaseRequestParam struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

//###########################################################
// Account related API objects
//###########################################################

// AccountInfo type represents 'account_info' API objects
type AccountInfo struct {
	TwoFARequired   BaseBoolInt `json:"2fa_required"`
	Country         string      `json:"country"`
	HTTPSRequired   BaseBoolInt `json:"https_required"`
	Intro           BaseBoolInt `json:"intro"`
	Lang            int         `json:"lang"`
	NoWallReplies   BaseBoolInt `json:"no_wall_replies"`
	OwnPostsDefault BaseBoolInt `json:"own_posts_default"`
}

// A AccountUserSettings represents 'account_user_settings` API object
type AccountUserSettings struct {
	FirstName       string      `json:"first_name"`
	LastName        string      `json:"last_name"`
	ScreenName      string      `json:"screen_name"`
	Sex             int         `json:"sex"`
	Relation        int         `json:"relation"`
	Birthday        string      `json:"bdate"`
	BirthVisibility int         `json:"bdate_visibility"`
	Hometown        string      `json:"home_town"`
	Status          string      `json:"status"`
	Phone           string      `json:"phone"`
	Country         BaseCountry `json:"country"`
	City            BaseObject  `json:"city"`
}

//###########################################################
// Photos related API objects
//###########################################################

type PhotosPhoto struct {
}
