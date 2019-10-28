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
// `fave` group of objects
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

// FaveBookmark type represents `fave_bookmark` API object
type FaveBookmark struct {
	AddedDate int              `json:"added_date"` // Timestamp, when this item was bookmarked
	Link      BaseLink         `json:"link"`
	Post      WallWallpostFull `json:"post"`
	Product   MarketMarketItem `json:"product"`
	Seen      bool             `json:"seen"` // Has user seen this item
	Tags      []FaveTag        `json:"tags"`
	Type      FaveBookmarkType `json:"type"` // Item type
	Video     VideoVideo       `json:"video"`
}

// FaveBookmarkType type represents `fave_bookmark_type` API object
type FaveBookmarkType string

// FavePage type represents `fave_page` API object
type FavePage struct {
	Description string          `json:"description"` // Some info about user or group
	Group       GroupsGroupFull `json:"group"`
	Tags        []FaveTag       `json:"tags"`
	Type        FavePageType    `json:"type"`         // Item type
	UpdatedDate int             `json:"updated_date"` // Timestamp, when this page was bookmarked
	User        UsersUserFull   `json:"user"`
}

// FavePageType type represents `fave_page_type` API object
type FavePageType string

// FaveTag type represents `fave_tag` API object
type FaveTag struct {
	Id   int    `json:"id"`   // Tag id
	Name string `json:"name"` // Tag name
}
