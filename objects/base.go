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
// `base` group of objects
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

// BaseBoolInt type represents `base_bool_int` API object
type BaseBoolInt int

// BaseCity type represents `base_city` API object
type BaseCity struct {
	Id    int    `json:"id"`    // City ID
	Title string `json:"title"` // City title
}

// BaseCommentsInfo type represents `base_comments_info` API object
type BaseCommentsInfo struct {
	CanPost       BaseBoolInt `json:"can_post"`        // Information whether current user can comment the post
	Count         int         `json:"count"`           // Comments number
	GroupsCanPost bool        `json:"groups_can_post"` // Information whether groups can comment the post
}

// BaseCountry type represents `base_country` API object
type BaseCountry struct {
	Id    int    `json:"id"`    // Country ID
	Title string `json:"title"` // Country title
}

// BaseError type represents `base_error` API object
type BaseError struct {
	ErrorCode     int                `json:"error_code"` // Error code
	ErrorMsg      string             `json:"error_msg"`  // Error message
	RequestParams []BaseRequestParam `json:"request_params"`
}

// BaseGeo type represents `base_geo` API object
type BaseGeo struct {
	Coordinates BaseGeoCoordinates `json:"coordinates"`
	Place       BasePlace          `json:"place"`
	Showmap     int                `json:"showmap"` // Information whether a map is showed
	Type        string             `json:"type"`    // Place type
}

// BaseGeoCoordinates type represents `base_geo_coordinates` API object
type BaseGeoCoordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// BaseImage type represents `base_image` API object
type BaseImage struct {
	Height int    `json:"height"` // Image height
	Url    string `json:"url"`    // Image url
	Width  int    `json:"width"`  // Image width
}

// BaseLikes type represents `base_likes` API object
type BaseLikes struct {
	Count     int         `json:"count"`      // Likes number
	UserLikes BaseBoolInt `json:"user_likes"` // Information whether current user likes the photo
}

// BaseLikesInfo type represents `base_likes_info` API object
type BaseLikesInfo struct {
	CanLike    BaseBoolInt `json:"can_like"`    // Information whether current user can like the post
	CanPublish BaseBoolInt `json:"can_publish"` // Information whether current user can repost
	Count      int         `json:"count"`       // Likes number
	UserLikes  int         `json:"user_likes"`  // Information whether current uer has liked the post
}

// BaseLink type represents `base_link` API object
type BaseLink struct {
	Application BaseLinkApplication `json:"application"`
	Button      BaseLinkButton      `json:"button"`
	Caption     string              `json:"caption"`     // Link caption
	Description string              `json:"description"` // Link description
	Id          string              `json:"id"`          // Link ID
	IsFavorite  bool                `json:"is_favorite"`
	Photo       PhotosPhoto         `json:"photo"`
	PreviewPage string              `json:"preview_page"` // String ID of the page with article preview
	PreviewUrl  string              `json:"preview_url"`  // URL of the page with article preview
	Product     BaseLinkProduct     `json:"product"`
	Rating      BaseLinkRating      `json:"rating"`
	Title       string              `json:"title"` // Link title
	Url         string              `json:"url"`   // Link URL
}

// BaseLinkApplication type represents `base_link_application` API object
type BaseLinkApplication struct {
	AppId float64                  `json:"app_id"` // Application Id
	Store BaseLinkApplicationStore `json:"store"`
}

// BaseLinkApplicationStore type represents `base_link_application_store` API object
type BaseLinkApplicationStore struct {
	Id   float64 `json:"id"`   // Store Id
	Name string  `json:"name"` // Store name
}

// BaseLinkButton type represents `base_link_button` API object
type BaseLinkButton struct {
	Action BaseLinkButtonAction `json:"action"` // Button action
	Title  string               `json:"title"`  // Button title
}

// BaseLinkButtonAction type represents `base_link_button_action` API object
type BaseLinkButtonAction struct {
	Type BaseLinkButtonActionType `json:"type"`
	Url  string                   `json:"url"` // Action URL
}

// BaseLinkButtonActionType type represents `base_link_button_action_type` API object
type BaseLinkButtonActionType string // Action type

// BaseLinkProduct type represents `base_link_product` API object
type BaseLinkProduct struct {
	Merchant    string      `json:"merchant"`
	OrdersCount int         `json:"orders_count"`
	Price       MarketPrice `json:"price"`
}

// BaseLinkRating type represents `base_link_rating` API object
type BaseLinkRating struct {
	ReviewsCount int     `json:"reviews_count"` // Count of reviews
	Stars        float64 `json:"stars"`         // Count of stars
}

// BaseMessageError type represents `base_message_error` API object
type BaseMessageError struct {
	Code        int    `json:"code"`        // Error code
	Description string `json:"description"` // Error message
}

// BaseObject type represents `base_object` API object
type BaseObject struct {
	Id    int    `json:"id"`    // Object ID
	Title string `json:"title"` // Object title
}

// BaseObjectCount type represents `base_object_count` API object
type BaseObjectCount struct {
	Count int `json:"count"` // Items count
}

// BaseObjectWithName type represents `base_object_with_name` API object
type BaseObjectWithName struct {
	Id   int    `json:"id"`   // Object ID
	Name string `json:"name"` // Object name
}

// BaseOkResponse type represents `base_ok_response` API object
type BaseOkResponse int // Returns 1 if request has been processed successfully

// BasePlace type represents `base_place` API object
type BasePlace struct {
	Address   string  `json:"address"`   // Place address
	Checkins  int     `json:"checkins"`  // Checkins number
	City      string  `json:"city"`      // City name
	Country   string  `json:"country"`   // Country name
	Created   int     `json:"created"`   // Date of the place creation in Unixtime
	Icon      string  `json:"icon"`      // URL of the place's icon
	Id        int     `json:"id"`        // Place ID
	Latitude  float64 `json:"latitude"`  // Place latitude
	Longitude float64 `json:"longitude"` // Place longitude
	Title     string  `json:"title"`     // Place title
	Type      string  `json:"type"`      // Place type
}

// BasePropertyExists type represents `base_property_exists` API object
type BasePropertyExists int

// BaseRepostsInfo type represents `base_reposts_info` API object
type BaseRepostsInfo struct {
	Count        int `json:"count"`         // Reposts number
	UserReposted int `json:"user_reposted"` // Information whether current user has reposted the post
}

// BaseRequestParam type represents `base_request_param` API object
type BaseRequestParam struct {
	Key   string `json:"key"`   // Parameter name
	Value string `json:"value"` // Parameter value
}

// BaseSex type represents `base_sex` API object
type BaseSex int

// BaseSticker type represents `base_sticker` API object
type BaseSticker struct {
	Images               []BaseImage `json:"images"`
	ImagesWithBackground []BaseImage `json:"images_with_background"`
	ProductId            int         `json:"product_id"` // Collection ID
	StickerId            int         `json:"sticker_id"` // Sticker ID
}

// BaseUploadServer type represents `base_upload_server` API object
type BaseUploadServer struct {
	UploadUrl string `json:"upload_url"` // Upload URL
}

// BaseUserGroupFields type represents `base_user_group_fields` API object
type BaseUserGroupFields string

// BaseUserId type represents `base_user_id` API object
type BaseUserId struct {
	UserId int `json:"user_id"` // User ID
}
