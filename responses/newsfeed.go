package responses

import "gitlab.com/Burmuley/go-vkapi/objects"

/////////////////////////////////////////////////////////////
// Newsfeed related responses                              //
/////////////////////////////////////////////////////////////

// NewsfeedGetBannedExt represents `newsfeed_getBanned_extended_response` API response object
type NewsfeedGetBannedExt struct {
	Groups  []objects.GroupsGroupFull `json:"groups"`
	Members []objects.UsersUserFull   `json:"members"`
}

// NewsfeedGetBanned represents `newsfeed_getBanned_response` API response object
type NewsfeedGetBanned struct {
	Groups  []int `json:"groups"`  // Community IDs
	Members []int `json:"members"` // User IDs
}

// NewsfeedGetComments represents `newsfeed_getComments_response` API response object
type NewsfeedGetComments struct {
	Items    []objects.NewsfeedItem    `json:"items"`
	Profiles []objects.UsersUserFull   `json:"profiles"`
	Groups   []objects.GroupsGroupFull `json:"groups"`
	NextFrom int                       `json:"next_from"` // New from value
}

// NewsfeedGetListsExt represents `newsfeed_getLists_extended_response` API response object
type NewsfeedGetListsExt struct {
	Count int                        `json:"count"` // Total number
	Items []objects.NewsfeedListFull `json:"items"`
}

// NewsfeedGetLists represents `newsfeed_getLists_response` API response object
type NewsfeedGetLists struct {
	Count int                    `json:"count"` // Total number
	Items []objects.NewsfeedList `json:"items"`
}

// NewsfeedGetMentions represents `newsfeed_getMentions_response` API response object
type NewsfeedGetMentions struct {
	Count int                        `json:"count"` // Total number
	Items []objects.WallWallpostToId `json:"items"`
}

// NewsfeedGetRecommended represents `newsfeed_getRecommended_response` API response object
type NewsfeedGetRecommended struct {
	Items     []objects.NewsfeedItem    `json:"items"`
	Profiles  []objects.UsersUserFull   `json:"profiles"`
	Groups    []objects.GroupsGroupFull `json:"groups"`
	NewOffset int                       `json:"new_offset"` // New offset value
	NextFrom  int                       `json:"next_from"`  // New from value
}

// NewsfeedGetSuggestedSources represents `newsfeed_getSuggestedSources_response` API response object
type NewsfeedGetSuggestedSources struct {
	Count int `json:"count"` // Total number
	Items []struct {
		*objects.GroupsGroupFull
		*objects.UsersUserXtrType
	}
}

// NewsfeedGetResponse represents `newsfeed_get_response` API response object
type NewsfeedGetResponse struct {
	Items    []objects.NewsfeedItem    `json:"items"`
	Profiles []objects.UsersUserFull   `json:"profiles"`
	Groups   []objects.GroupsGroupFull `json:"groups"`
}

// NewsfeedSaveList represents `newsfeed_saveList_response` API response object
type NewsfeedSaveList int

// NewsfeedSearchExt represents `newsfeed_search_extended_response` API response object
type NewsfeedSearchExt struct {
	Items    []objects.WallWallpostFull `json:"items"`
	Profiles []objects.UsersUserFull    `json:"profiles"`
	Groups   []objects.GroupsGroupFull  `json:"groups"`
}

// NewsfeedSearch represents `newsfeed_search_response` API response object
type NewsfeedSearch struct {
	Items            []objects.WallWallpostFull `json:"items"`
	SuggestedQueries []string                   `json:"suggested_queries"`
}
