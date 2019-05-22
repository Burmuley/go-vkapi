package go_vkapi

import (
	"fmt"
	"gitlab.com/Burmuley/go-vkapi/objects"
	"gitlab.com/Burmuley/go-vkapi/responses"
	"strings"
)

// Account represents collection of methods related to VK Accounts
type Account struct {
	*VKApi
}

// Ban performs blocking operation for user
//
// Parameters:
//  * ownerId - ID of the user to ban
func (a *Account) Ban(ownerId int) (responses.OkResponse, error) {
	panic("implement me!")
}

// ChangePassword changes a user password after access is successfully restored with the [vk.com/dev/auth.restore|auth.restore] method.
//
// Parameters:
// 	* restoreSid - session id received after the [vk.com/dev/auth.restore|auth.restore] method is executed. (If the password is changed right after the access was restored)
// 	* changePwdhash - hash received after a successful OAuth authorization with a code got by SMS. (If the password is changed right after the access was restored)
// 	* oldPwd - Current user password
// 	* newPwd - New password that will be set as a current
func (a *Account) ChangePassword(restoreSid, changePwdHash, oldPwd, newPwd string) (responses.AccountChangePassword, error) {
	panic("implement me!")
}

// GetActiveOffers returns a list of active ads (offers) which executed by the user will bring him/her respective number of votes to his balance in the application.
//
// Parameters:
//  * offset - NO DESCRIPTION
//  * count - Number of results to return
func (a *Account) GetActiveOffers(offset, count int) (responses.AccountGetActiveOffers, error) {
	panic("implement me!")
}

// GetAppPermissions gets settings of the user in this application.
//
// Parameters:
//  * userId - User ID whose settings information shall be got. By default: current user (0).
func (a *Account) GetAppPermissions(userId int) (responses.AccountGetAppPermissions, error) {
	panic("implement me!")
}

// GetBanned Returns a user's blacklist
//
// Parameters:
//  * offset - Offset needed to return a specific subset of results
//  * count - Number of results to return
func (a *Account) GetBanned(offset, count int) (responses.AccountGetBanned, error) {
	panic("implement me!")
}

// GetCounters returns non-null values of user counters
//
// Parameters:
//  * filter - Counters to be returned. Possible values:
//              "friends",
//              "messages",
//              "photos",
//              "videos",
//              "notes",
//              "gifts",
//              "events",
//              "groups",
//              "sdk",
//              "friends_suggestions"
func (a *Account) GetCounters(filters ...string) (responses.AccountGetCounters, error) {
	resp := responses.AccountGetCounters{}

	if err := a.SendObjRequest("account.getCounters", map[string]string{}, resp); err != nil {
		return responses.AccountGetCounters{}, err
	}

	return resp, nil
}

// GetInfo returns current account info
//
// Parameters:
//  * fields - Fields to return. Possible values:
//     'country' — user country,
//     'https_required' — is "HTTPS only" option enabled,
//     'own_posts_default' — is "Show my posts only" option is enabled,
//     'no_wall_replies' — are wall replies disabled or not,
//     'intro' — is intro passed by user or not,
//     'lang' — user language.
//     By default: all.
func (a *Account) GetInfo(fields ...string) (responses.AccountGetInfo, error) {
	params := map[string]string{"fields": strings.Join(fields, ",")}
	resp := responses.AccountGetInfo{}

	if err := a.SendObjRequest("account.getInfo", params, resp); err != nil {
		return responses.AccountGetInfo{}, err
	}

	return resp, nil
}

// GetProfileInfo returns the current account info
//
// Parameters: no
func (a *Account) GetProfileInfo() (responses.AccountGetProfileInfo, error) {
	resp := responses.AccountGetProfileInfo{}

	if err := a.SendObjRequest("account.getProfileInfo", map[string]string{}, resp); err != nil {
		return responses.AccountGetProfileInfo{}, err
	}

	return resp, nil
}

// GetPushSettings return push notifications settings
//
// Parameters:
//  * deviceId - Unique device ID
func (a *Account) GetPushSettings(deviceId string) (responses.AccountGetPushSettings, error) {
	params := map[string]string{"device_id": deviceId}
	resp := responses.AccountGetPushSettings{}

	if err := a.SendObjRequest("account.getPushSettings", params, resp); err != nil {
		return responses.AccountGetPushSettings{}, err
	}

	return resp, nil
}

// RegisterDevice subscribes an iOS/Android/Windows Phone-based device to receive push notifications
//
// Parameters:
//  * token - device token used to send notifications. (for mpns, the token shall be URL for sending of notifications)
//  * deviceModel - string name of device model
//  * deviceYear - device year
//  * deviceId - unique device ID
//  * systemVer - string version of device operating system
//  * settings - Push settings in a special format (see https://vk.com/dev/push_settings)
//  * sandbox - test run or not
func (a *Account) RegisterDevice(token, deviceModel, deviceId, systemVer, settings objects.AccountPushSettings,
	deviceYear int, sandbox objects.BaseBoolInt) (responses.OkResponse, error) {
	fmt.Println(sandbox)
	panic("implement me!")
}

// SaveProfileInfo edits current profile info
//
// Parameters:
//  * firstName, lastName, maidenName, screenName - corresponding user's first name, last name,  maiden name (for females only) and screen name (nickname)
//  * cancelReqId - ID of the name change request to be canceled. If this parameter is sent, all the others are ignored
//  * sex - user's sex. Possible values:
//    1 – 'female'
//    2 – 'male'
//  * relation - user relationship status. Possible values:
//    1 – single
//    2 – in a relationship
//    3 – engaged
//    4 – married
//    5 – it's complicated
//    6 – actively searching
//    7 – in love
//    0 – not specified
//  * relationPartner - ID of the relationship partner
//  * bdate - user's birth date, format: DD.MM.YYYY
//  * bdateVisible - Birth date visibility. Possible values:
//    1 – show birth date
//    2 – show only month and day
//    0 – hide birth date.
//  * homeTown - user's home town
//  * countryId - user's country ID
//  * cityId - user's city ID
//  * status - status text
func (a *Account) SaveProfileInfo(firstName, lastName, maidenName, screenName string,
	cancelReqId int, sex, relation, relationPartner int, bdate string, bdateVisible int,
	homeTown string, countryId, cityId int, status string) (responses.AccountSaveProfileInfo, error) {
	panic("implement me!")
}

// SetInfo allows to edit the current account info
//
// Parameters:
//  * settings - slice of settings in format {name: "name", value: "value"}
func (a *Account) SetInfo(settings ...struct{ name, value string }) (responses.OkResponse, error) {
	panic("implement me!")
}

// SetNameInMenu sets an application screen name (up to 17 characters), that is shown to the user in the left menu
//
// Parameters:
//  * userId - user's ID
//  * name - application screen name
func (a *Account) SetNameInMenu(userId int, name string) (responses.OkResponse, error) {
	panic("implement me!")
}

// SetOffline marks a current user as offline
//
// Parameters: none
func (a *Account) SetOffline() (responses.OkResponse, error) {
	panic("implement me!")
}

// SetOnline marks a current user as online
//
// Parameters:
//  * voip - '1' if video calls are available for current device
func (a *Account) SetOnline(voip bool) (responses.OkResponse, error) {
	panic("implement me!")
}

// SetPushSettings changes push settings
//
// Parameters:
//  * deviceId - Unique device ID
//  * settings - Push settings in a special format (see https://vk.com/dev/push_settings)
//  * key - notification key
//  * value - new value for the key in a special format (see https://vk.com/dev/push_settings)
func (a *Account) SetPushSettings(deviceId, key string, value []string, settings objects.AccountPushSettings) (responses.OkResponse, error) {
	panic("implement me!")
}

// SetSilenceMode mutes push notifications for the set period of time
//
// Parameters:
//  * deviceId - unique device ID
//  * time - time in seconds for what notifications should be disabled. '-1' to disable forever.
//  * peerId - destination ID. For user: 'User ID', e.g. '12345'.
//                             For chat: '2000000000' + 'Chat ID', e.g. '2000000001'
//                             For community: '- Community ID', e.g. '-12345'
//  * sound - '1' — to enable sound in this dialog, '0' — to disable sound.
//    Only if 'peer_id' contains user or community ID
func (a *Account) SetSilenceMode(deviceId string, time, peerId, sound int) (responses.OkResponse, error) {
	panic("implement me!")
}

// Unban unblocks user
//
// Parameters:
//  * ownerId - ID of the user to ban
func (a *Account) Unban(ownerId int) (responses.OkResponse, error) {
	panic("implement me!")
}

// UnregisterDevice unsubscribes a device from push notifications
//
// Parameters:
//  * deviceId - unique device ID
//  * sandbox - test request or not
func (a *Account) UnregisterDevice(deviceId int, sandbox bool) (responses.OkResponse, error) {
	panic("implement me!")
}
