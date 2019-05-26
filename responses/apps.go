package responses

import "gitlab.com/Burmuley/go-vkapi/objects"

/////////////////////////////////////////////////////////////
// Applications related responses                          //
/////////////////////////////////////////////////////////////

// AppsGetCatalog represents `apps_getCatalog_response` API response object
type AppsGetCatalog struct {
	Count int               `json:"count"` // Total number
	Items []objects.AppsApp `json:"items"`
}

// AppsGetFriendsList represents `apps_getFriendsList_response` API response object
type AppsGetFriendsList struct {
	Count int                     `json:"count"` // Total number
	Items []objects.UsersUserFull `json:"items"`
}

// AppsGetLeaderboardExt represents `apps_getLeaderboard_extended_response` API response object
type AppsGetLeaderboardExt struct {
	Count    int                       `json:"count"` // Total number
	Items    []objects.AppsLeaderboard `json:"items"`
	Profiles []objects.UsersUserMin    `json:"profiles"`
}

// AppsGetLeaderboard represents `apps_getLeaderboard_response` API response object
type AppsGetLeaderboard struct {
	Count int                       `json:"count"` // Total number
	Items []objects.AppsLeaderboard `json:"items"`
}

// AppsGetScopes represents `apps_getScopes_response` API response object
type AppsGetScopes struct {
	Count int                 `json:"count"` // Total number
	Items []objects.AppsScope `json:"items"`
}

// AppsGetScore represents `apps_getScore_response` API response object
type AppsGetScore int // Score number

// AppsGet represents `apps_get_response` API response object
type AppsGet struct {
	Count int               `json:"count"` // Total number
	Items []objects.AppsApp `json:"items"`
}

// AppsSendRequest represents `apps_sendRequest_response` API response object
type AppsSendRequest int // Request ID
