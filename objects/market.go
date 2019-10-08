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
// `market` group of objects
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

// MarketMarketItemAvailability type represents `market_market_item_availability` API object
type MarketMarketItemAvailability int // Information whether the item is available

// MarketMarketCategory type represents `market_market_category` API object
type MarketMarketCategory struct {
	Id      int           `json:"id"`   // Category ID
	Name    string        `json:"name"` // Category name
	Section MarketSection `json:"section"`
}

// MarketMarketAlbum type represents `market_market_album` API object
type MarketMarketAlbum struct {
	Count       int         `json:"count"`    // Items number
	Id          int         `json:"id"`       // Market album ID
	OwnerId     int         `json:"owner_id"` // Market album owner's ID
	Photo       PhotosPhoto `json:"photo"`
	Title       string      `json:"title"`        // Market album title
	UpdatedTime int         `json:"updated_time"` // Date when album has been updated last time in Unixtime
}

// MarketMarketItemFull type represents `market_market_item_full` API object
type MarketMarketItemFull struct {
	MarketMarketItem MarketMarketItem `json:"MarketMarketItem"`
	AlbumsIds        []int            `json:"albums_ids"`
	CanComment       BaseBoolInt      `json:"can_comment"` // Information whether current use can comment the item
	CanRepost        BaseBoolInt      `json:"can_repost"`  // Information whether current use can repost the item
	Likes            BaseLikes        `json:"likes"`
	Photos           []PhotosPhoto    `json:"photos"`
	Reposts          BaseRepostsInfo  `json:"reposts"`
	ViewsCount       int              `json:"views_count"` // Views number
}

// MarketMarketItem type represents `market_market_item` API object
type MarketMarketItem struct {
	AccessKey    string                       `json:"access_key"` // Access key for the market item
	Availability MarketMarketItemAvailability `json:"availability"`
	ButtonTitle  string                       `json:"button_title"` // Title for button for url
	Category     MarketMarketCategory         `json:"category"`
	Date         int                          `json:"date"`        // Date when the item has been created in Unixtime
	Description  string                       `json:"description"` // Item description
	ExternalId   string                       `json:"external_id"`
	Id           int                          `json:"id"` // Item ID
	IsFavorite   bool                         `json:"is_favorite"`
	OwnerId      int                          `json:"owner_id"` // Item owner's ID
	Price        MarketPrice                  `json:"price"`
	ThumbPhoto   string                       `json:"thumb_photo"` // URL of the preview image
	Title        string                       `json:"title"`       // Item title
	Url          string                       `json:"url"`         // URL to item
}

// MarketCurrency type represents `market_currency` API object
type MarketCurrency struct {
	Id   int    `json:"id"`   // Currency ID
	Name string `json:"name"` // Currency sign
}

// MarketSection type represents `market_section` API object
type MarketSection struct {
	Id   int    `json:"id"`   // Section ID
	Name string `json:"name"` // Section name
}

// MarketPrice type represents `market_price` API object
type MarketPrice struct {
	Amount       string         `json:"amount"` // Amount
	Currency     MarketCurrency `json:"currency"`
	DiscountRate int            `json:"discount_rate"`
	OldAmount    string         `json:"old_amount"`
	Text         string         `json:"text"` // Text
}
