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
// `newsfeed` group of responses
/////////////////////////////////////////////////////////////

// NewsfeedGetbannedExtended type represents `newsfeed_getBanned_extended_response` API response object
type NewsfeedGetbannedExtended struct {
	Groups   []objects.UsersUserFull   `json:"groups"`
	Profiles []objects.GroupsGroupFull `json:"profiles"`
}

// NewsfeedGetbanned type represents `newsfeed_getBanned_response` API response object
type NewsfeedGetbanned struct {
	Groups  []int `json:"groups"`
	Members []int `json:"members"`
}

// NewsfeedGetcomments type represents `newsfeed_getComments_response` API response object
type NewsfeedGetcomments struct {
	Groups   []objects.GroupsGroupFull      `json:"groups"`
	Items    []objects.NewsfeedNewsfeedItem `json:"items"`
	NextFrom string                         `json:"next_from"` // New from value
	Profiles []objects.UsersUserFull        `json:"profiles"`
}

// NewsfeedGetlistsExtended type represents `newsfeed_getLists_extended_response` API response object
type NewsfeedGetlistsExtended struct {
	Count int                        `json:"count"` // Total number
	Items []objects.NewsfeedListFull `json:"items"`
}

// NewsfeedGetlists type represents `newsfeed_getLists_response` API response object
type NewsfeedGetlists struct {
	Count int                    `json:"count"` // Total number
	Items []objects.NewsfeedList `json:"items"`
}

// NewsfeedGetmentions type represents `newsfeed_getMentions_response` API response object
type NewsfeedGetmentions struct {
	Count int                        `json:"count"` // Total number
	Items []objects.WallWallpostToId `json:"items"`
}

// NewsfeedGetrecommended type represents `newsfeed_getRecommended_response` API response object
type NewsfeedGetrecommended struct {
	Groups    []objects.GroupsGroupFull      `json:"groups"`
	Items     []objects.NewsfeedNewsfeedItem `json:"items"`
	NewOffset string                         `json:"new_offset"` // New offset value
	NextFrom  string                         `json:"next_from"`  // Next from value
	Profiles  []objects.UsersUserFull        `json:"profiles"`
}

// NewsfeedGetsuggestedsources type represents `newsfeed_getSuggestedSources_response` API response object
type NewsfeedGetsuggestedsources struct {
	Count                    int `json:"count"` // Total number
	Items                    []objects.GroupsGroupFull
	objects.UsersUserXtrType `json:"items"`
}

// NewsfeedGet type represents `newsfeed_get_response` API response object
type NewsfeedGet struct {
	Groups   []objects.GroupsGroupFull      `json:"groups"`
	Items    []objects.NewsfeedNewsfeedItem `json:"items"`
	NextFrom string                         `json:"next_from"` // New from value
	Profiles []objects.UsersUserFull        `json:"profiles"`
}

// NewsfeedSavelist type represents `newsfeed_saveList_response` API response object
type NewsfeedSavelist int // List ID

// NewsfeedSearchExtended type represents `newsfeed_search_extended_response` API response object
type NewsfeedSearchExtended struct {
	Count      int                        `json:"count"` // Filtered number
	Groups     []objects.GroupsGroupFull  `json:"groups"`
	Items      []objects.WallWallpostFull `json:"items"`
	NextFrom   string                     `json:"next_from"`
	Profiles   []objects.UsersUserFull    `json:"profiles"`
	TotalCount int                        `json:"total_count"` // Total number
}

// NewsfeedSearch type represents `newsfeed_search_response` API response object
type NewsfeedSearch struct {
	Items            []objects.WallWallpostFull `json:"items"`
	SuggestedQueries []string                   `json:"suggested_queries"`
}
