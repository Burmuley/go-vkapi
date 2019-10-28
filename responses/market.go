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

package responses

import (
	"gitlab.com/Burmuley/go-vkapi/objects"
)

/////////////////////////////////////////////////////////////
// `market` group of responses
/////////////////////////////////////////////////////////////

// MarketAddalbum type represents `market_addAlbum_response` API response object
type MarketAddalbum struct {
	MarketAlbumId int `json:"market_album_id"` // Album ID
}

// MarketAdd type represents `market_add_response` API response object
type MarketAdd struct {
	MarketItemId int `json:"market_item_id"` // Item ID
}

// MarketCreatecomment type represents `market_createComment_response` API response object
type MarketCreatecomment int // Comment ID

// MarketDeletecomment type represents `market_deleteComment_response` API response object
type MarketDeletecomment objects.BaseBoolInt // Returns 1 if request has been processed successfully (0 if the comment is not found)

// MarketGetalbumbyid type represents `market_getAlbumById_response` API response object
type MarketGetalbumbyid struct {
	Count int                         `json:"count"` // Total number
	Items []objects.MarketMarketAlbum `json:"items"`
}

// MarketGetalbums type represents `market_getAlbums_response` API response object
type MarketGetalbums struct {
	Count int                         `json:"count"` // Total number
	Items []objects.MarketMarketAlbum `json:"items"`
}

// MarketGetbyidExtended type represents `market_getById_extended_response` API response object
type MarketGetbyidExtended struct {
	Count int                            `json:"count"` // Total number
	Items []objects.MarketMarketItemFull `json:"items"`
}

// MarketGetbyid type represents `market_getById_response` API response object
type MarketGetbyid struct {
	Count int                        `json:"count"` // Total number
	Items []objects.MarketMarketItem `json:"items"`
}

// MarketGetcategories type represents `market_getCategories_response` API response object
type MarketGetcategories struct {
	Count int                            `json:"count"` // Total number
	Items []objects.MarketMarketCategory `json:"items"`
}

// MarketGetcomments type represents `market_getComments_response` API response object
type MarketGetcomments struct {
	Count int                       `json:"count"` // Total number
	Items []objects.WallWallComment `json:"items"`
}

// MarketGetExtended type represents `market_get_extended_response` API response object
type MarketGetExtended struct {
	Count int                            `json:"count"` // Total number
	Items []objects.MarketMarketItemFull `json:"items"`
}

// MarketGet type represents `market_get_response` API response object
type MarketGet struct {
	Count int                        `json:"count"` // Total number
	Items []objects.MarketMarketItem `json:"items"`
}

// MarketRestorecomment type represents `market_restoreComment_response` API response object
type MarketRestorecomment objects.BaseBoolInt // Returns 1 if request has been processed successfully (0 if the comment is not found)

// MarketSearchExtended type represents `market_search_extended_response` API response object
type MarketSearchExtended struct {
	Count int                            `json:"count"` // Total number
	Items []objects.MarketMarketItemFull `json:"items"`
}

// MarketSearch type represents `market_search_response` API response object
type MarketSearch struct {
	Count int                        `json:"count"` // Total number
	Items []objects.MarketMarketItem `json:"items"`
}
