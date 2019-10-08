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

import "gitlab.com/Burmuley/go-vkapi/objects"

/////////////////////////////////////////////////////////////
// `newsfeed` group of responses
/////////////////////////////////////////////////////////////

// NewsfeedGetListsExtended type represents `newsfeed_getLists_extended_response` API response object
type NewsfeedGetListsExtended struct {
	Count int                        `json:"count"` // Total number
	Items []objects.NewsfeedListFull `json:"items"`
}

// NewsfeedSearchExtended type represents `newsfeed_search_extended_response` API response object
type NewsfeedSearchExtended struct {
	Count      int                        `json:"count"` // Filtered number
	Groups     []objects.GroupsGroupFull  `json:"groups"`
	Items      []objects.WallWallpostFull `json:"items"`
	NextFrom   string                     `json:"next_from"`
	Profiles   []objects.UsersUserFull    `json:"profiles"`
	TotalCount int                        `json:"total_count"` // Total number
}

// NewsfeedGet type represents `newsfeed_get_response` API response object
type NewsfeedGet struct {
	Groups   []objects.GroupsGroupFull      `json:"groups"`
	Items    []objects.NewsfeedNewsfeedItem `json:"items"`
	NextFrom string                         `json:"next_from"` // New from value
	Profiles []objects.UsersUserFull        `json:"profiles"`
}

// NewsfeedSaveList type represents `newsfeed_saveList_response` API response object
type NewsfeedSaveList int // List ID

// NewsfeedGetLists type represents `newsfeed_getLists_response` API response object
type NewsfeedGetLists struct {
	Count int                    `json:"count"` // Total number
	Items []objects.NewsfeedList `json:"items"`
}

// NewsfeedGetSuggestedSources type represents `newsfeed_getSuggestedSources_response` API response object
type NewsfeedGetSuggestedSources struct {
	Count int `json:"count"` // Total number
	Items []struct {
		objects.GroupsGroupFull
		objects.UsersUserXtrType
	} `json:"items"`
}

// NewsfeedGetBanned type represents `newsfeed_getBanned_response` API response object
type NewsfeedGetBanned struct {
	Groups  []int `json:"groups"`
	Members []int `json:"members"`
}

// NewsfeedGetBannedExtended type represents `newsfeed_getBanned_extended_response` API response object
type NewsfeedGetBannedExtended struct {
	Groups   []objects.UsersUserFull   `json:"groups"`
	Profiles []objects.GroupsGroupFull `json:"profiles"`
}

// NewsfeedGetComments type represents `newsfeed_getComments_response` API response object
type NewsfeedGetComments struct {
	Groups   []objects.GroupsGroupFull      `json:"groups"`
	Items    []objects.NewsfeedNewsfeedItem `json:"items"`
	NextFrom string                         `json:"next_from"` // New from value
	Profiles []objects.UsersUserFull        `json:"profiles"`
}

// NewsfeedGetMentions type represents `newsfeed_getMentions_response` API response object
type NewsfeedGetMentions struct {
	Count int                        `json:"count"` // Total number
	Items []objects.WallWallpostToId `json:"items"`
}

// NewsfeedSearch type represents `newsfeed_search_response` API response object
type NewsfeedSearch struct {
	Items            []objects.WallWallpostFull `json:"items"`
	SuggestedQueries []string                   `json:"suggested_queries"`
}

// NewsfeedGetRecommended type represents `newsfeed_getRecommended_response` API response object
type NewsfeedGetRecommended struct {
	Groups    []objects.GroupsGroupFull      `json:"groups"`
	Items     []objects.NewsfeedNewsfeedItem `json:"items"`
	NewOffset string                         `json:"new_offset"` // New offset value
	NextFrom  string                         `json:"next_from"`  // Next from value
	Profiles  []objects.UsersUserFull        `json:"profiles"`
}
