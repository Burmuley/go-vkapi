package responses

import (
	"gitlab.com/Burmuley/go-vkapi/objects"
)

/////////////////////////////////////////////////////////////
// Account related responses                               //
/////////////////////////////////////////////////////////////

// ChangePassword represents `account_changePassword_response` API response object.
type AccountChangePassword struct {
	Token  string `json:"token"`  // New token
	Secret string `json:"secret"` // New secret
}

// GetActiveOffers represents `account_getActiveOffers_response` API response object.
type AccountGetActiveOffers struct {
	Count int `json:"count"` // Total number
	Items []objects.AccountOffer
}

// GetAppPermissions represents `account_getAppPermissions_response` API response object.
// Contains permissions mask.
type AccountGetAppPermissions int

// GetBanned represents `account_getBanned_response` API response object.
type AccountGetBanned struct {
	Count    int                    `json:"count"` // Total number
	Items    []int                  `json:"items"`
	Profiles []objects.UsersUserMin `json:"profiles"`
	Groups   []objects.GroupsGroup  `json:"groups"`
}

// GetCounters represents `account_getCounters_response` API response object.
type AccountGetCounters struct {
	*objects.AccountCounters
}
