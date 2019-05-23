package responses

import (
	"gitlab.com/Burmuley/go-vkapi/objects"
)

/////////////////////////////////////////////////////////////
// Account related responses                               //
/////////////////////////////////////////////////////////////

// AccountChangePassword represents `account_changePassword_response` API response object.
type AccountChangePassword struct {
	Token  string `json:"token"`  // New token
	Secret string `json:"secret"` // New secret
}

// AccountGetActiveOffers represents `account_getActiveOffers_response` API response object.
type AccountGetActiveOffers struct {
	Count int `json:"count"` // Total number
	Items []objects.AccountOffer
}

// AccountGetAppPermissions represents `account_getAppPermissions_response` API response object.
// Contains permissions mask.
type AccountGetAppPermissions int

// AccountGetBanned represents `account_getBanned_response` API response object.
type AccountGetBanned struct {
	Count    int                    `json:"count"` // Total number
	Items    []int                  `json:"items"`
	Profiles []objects.UsersUserMin `json:"profiles"`
	Groups   []objects.GroupsGroup  `json:"groups"`
}

// AccountGetCounters represents `account_getCounters_response` API response object.
type AccountGetCounters struct {
	*objects.AccountCounters
}

// AccountGetInfo represents `account_getInfo_response` API response object.
type AccountGetInfo struct {
	*objects.AccountInfo
}

// AccountGetProfileInfo represents `account_getProfileInfo_response` API response object.
type AccountGetProfileInfo struct {
	*objects.AccountUserSettings
}

// AccountGetPushSettings represents `account_getPushSettings_response` API response object.
type AccountGetPushSettings struct {
	*objects.AccountPushSettings
}

// AccountSaveProfileInfo represents `account_saveProfileInfo_response` API response object.
type AccountSaveProfileInfo struct {
	Changed     objects.BaseBoolInt        `json:"changed"` // 1 if changes has been processed
	NameRequest objects.AccountNameRequest `json:"name_request"`
}
