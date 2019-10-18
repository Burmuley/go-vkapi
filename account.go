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
// Source schema can be found at https://github.com/VKCOM/vk-api-schema/blob/master/methods.json           //
// Code generator location: https://gitlab.com/Burmuley/go-vkapi-gen                                       //
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

package go_vkapi

import (
	"gitlab.com/Burmuley/go-vkapi/responses"
)

type Account struct {
	*VKApi
}

/////////////////////////////////////////////////////////////
// `Account` methods
/////////////////////////////////////////////////////////////

// Ban - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * ownerId - NO DESCRIPTION IN JSON SCHEMA
func (a Account) Ban(ownerId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	err = a.SendObjRequest("account.ban", params, &resp)

	return
}

// ChangePassword - Changes a user password after access is successfully restored with the [vk.com/dev/auth.restore|auth.restore] method.
// Parameters:
//   * restoreSid - Session id received after the [vk.com/dev/auth.restore|auth.restore] method is executed. (If the password is changed right after the access was restored)
//   * changePasswordHash - Hash received after a successful OAuth authorization with a code got by SMS. (If the password is changed right after the access was restored)
//   * oldPassword - Current user password.
//   * newPassword - New password that will be set as a current
func (a Account) ChangePassword(restoreSid string, changePasswordHash string, oldPassword string, newPassword string) (resp responses.AccountChangePassword, err error) {
	params := map[string]interface{}{}

	if restoreSid != "" {
		params["restore_sid"] = restoreSid
	}

	if changePasswordHash != "" {
		params["change_password_hash"] = changePasswordHash
	}

	if oldPassword != "" {
		params["old_password"] = oldPassword
	}

	params["new_password"] = newPassword

	err = a.SendObjRequest("account.changePassword", params, &resp)

	return
}

// GetActiveOffers - Returns a list of active ads (offers) which executed by the user will bring him/her respective number of votes to his balance in the application.
// Parameters:
//   * offset - NO DESCRIPTION IN JSON SCHEMA
//   * count - Number of results to return.
func (a Account) GetActiveOffers(offset int, count int) (resp responses.AccountGetActiveOffers, err error) {
	params := map[string]interface{}{}

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	err = a.SendObjRequest("account.getActiveOffers", params, &resp)

	return
}

// GetAppPermissions - Gets settings of the user in this application.
// Parameters:
//   * userId - User ID whose settings information shall be got. By default: current user.
func (a Account) GetAppPermissions(userId int) (resp responses.AccountGetAppPermissions, err error) {
	params := map[string]interface{}{}

	params["user_id"] = userId

	err = a.SendObjRequest("account.getAppPermissions", params, &resp)

	return
}

// GetBanned - Returns a user's blacklist.
// Parameters:
//   * offset - Offset needed to return a specific subset of results.
//   * count - Number of results to return.
func (a Account) GetBanned(offset int, count int) (resp responses.AccountGetBanned, err error) {
	params := map[string]interface{}{}

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	err = a.SendObjRequest("account.getBanned", params, &resp)

	return
}

// GetCounters - Returns non-null values of user counters.
// Parameters:
//   * filter - Counters to be returned.
func (a Account) GetCounters(filter []string) (resp responses.AccountGetCounters, err error) {
	params := map[string]interface{}{}

	if len(filter) > 0 {
		params["filter"] = SliceToString(filter)
	}

	err = a.SendObjRequest("account.getCounters", params, &resp)

	return
}

// GetInfo - Returns current account info.
// Parameters:
//   * fields - Fields to return. Possible values: *'country' — user country,, *'https_required' — is "HTTPS only" option enabled,, *'own_posts_default' — is "Show my posts only" option is enabled,, *'no_wall_replies' — are wall replies disabled or not,, *'intro' — is intro passed by user or not,, *'lang' — user language. By default: all.
func (a Account) GetInfo(fields []string) (resp responses.AccountGetInfo, err error) {
	params := map[string]interface{}{}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	err = a.SendObjRequest("account.getInfo", params, &resp)

	return
}

// GetProfileInfo - Returns the current account info.
func (a Account) GetProfileInfo() (resp responses.AccountGetProfileInfo, err error) {
	params := map[string]interface{}{}

	err = a.SendObjRequest("account.getProfileInfo", params, &resp)

	return
}

// GetPushSettings - Gets settings of push notifications.
// Parameters:
//   * deviceId - Unique device ID.
func (a Account) GetPushSettings(deviceId string) (resp responses.AccountGetPushSettings, err error) {
	params := map[string]interface{}{}

	if deviceId != "" {
		params["device_id"] = deviceId
	}

	err = a.SendObjRequest("account.getPushSettings", params, &resp)

	return
}

// RegisterDevice - Subscribes an iOS/Android/Windows Phone-based device to receive push notifications
// Parameters:
//   * token - Device token used to send notifications. (for mpns, the token shall be URL for sending of notifications)
//   * deviceModel - String name of device model.
//   * deviceYear - Device year.
//   * deviceId - Unique device ID.
//   * systemVersion - String version of device operating system.
//   * settings - Push settings in a [vk.com/dev/push_settings|special format].
//   * sandbox - NO DESCRIPTION IN JSON SCHEMA
func (a Account) RegisterDevice(token string, deviceModel string, deviceYear int, deviceId string, systemVersion string, settings string, sandbox bool) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["token"] = token

	if deviceModel != "" {
		params["device_model"] = deviceModel
	}

	if deviceYear > 0 {
		params["device_year"] = deviceYear
	}

	params["device_id"] = deviceId

	if systemVersion != "" {
		params["system_version"] = systemVersion
	}

	if settings != "" {
		params["settings"] = settings
	}

	params["sandbox"] = sandbox

	err = a.SendObjRequest("account.registerDevice", params, &resp)

	return
}

// SaveProfileInfo - Edits current profile info.
// Parameters:
//   * firstName - User first name.
//   * lastName - User last name.
//   * maidenName - User maiden name (female only)
//   * screenName - User screen name.
//   * cancelRequestId - ID of the name change request to be canceled. If this parameter is sent, all the others are ignored.
//   * sex - User sex. Possible values: , * '1' – female,, * '2' – male.
//   * relation - User relationship status. Possible values: , * '1' – single,, * '2' – in a relationship,, * '3' – engaged,, * '4' – married,, * '5' – it's complicated,, * '6' – actively searching,, * '7' – in love,, * '0' – not specified.
//   * relationPartnerId - ID of the relationship partner.
//   * bdate - User birth date, format: DD.MM.YYYY.
//   * bdateVisibility - Birth date visibility. Returned values: , * '1' – show birth date,, * '2' – show only month and day,, * '0' – hide birth date.
//   * homeTown - User home town.
//   * countryId - User country.
//   * cityId - User city.
//   * status - Status text.
func (a Account) SaveProfileInfo(firstName string, lastName string, maidenName string, screenName string, cancelRequestId int, sex int, relation int, relationPartnerId int, bdate string, bdateVisibility int, homeTown string, countryId int, cityId int, status string) (resp responses.AccountSaveProfileInfo, err error) {
	params := map[string]interface{}{}

	if firstName != "" {
		params["first_name"] = firstName
	}

	if lastName != "" {
		params["last_name"] = lastName
	}

	if maidenName != "" {
		params["maiden_name"] = maidenName
	}

	if screenName != "" {
		params["screen_name"] = screenName
	}

	if cancelRequestId > 0 {
		params["cancel_request_id"] = cancelRequestId
	}

	if sex > 0 {
		params["sex"] = sex
	}

	if relation > 0 {
		params["relation"] = relation
	}

	if relationPartnerId > 0 {
		params["relation_partner_id"] = relationPartnerId
	}

	if bdate != "" {
		params["bdate"] = bdate
	}

	if bdateVisibility > 0 {
		params["bdate_visibility"] = bdateVisibility
	}

	if homeTown != "" {
		params["home_town"] = homeTown
	}

	if countryId > 0 {
		params["country_id"] = countryId
	}

	if cityId > 0 {
		params["city_id"] = cityId
	}

	if status != "" {
		params["status"] = status
	}

	err = a.SendObjRequest("account.saveProfileInfo", params, &resp)

	return
}

// SetInfo - Allows to edit the current account info.
// Parameters:
//   * name - Setting name.
//   * value - Setting value.
func (a Account) SetInfo(name string, value string) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	if name != "" {
		params["name"] = name
	}

	if value != "" {
		params["value"] = value
	}

	err = a.SendObjRequest("account.setInfo", params, &resp)

	return
}

// SetNameInMenu - Sets an application screen name (up to 17 characters), that is shown to the user in the left menu.
// Parameters:
//   * userId - User ID.
//   * name - Application screen name.
func (a Account) SetNameInMenu(userId int, name string) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["user_id"] = userId

	if name != "" {
		params["name"] = name
	}

	err = a.SendObjRequest("account.setNameInMenu", params, &resp)

	return
}

// SetOffline - Marks a current user as offline.
func (a Account) SetOffline() (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	err = a.SendObjRequest("account.setOffline", params, &resp)

	return
}

// SetOnline - Marks the current user as online for 15 minutes.
// Parameters:
//   * voip - '1' if videocalls are available for current device.
func (a Account) SetOnline(voip bool) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["voip"] = voip

	err = a.SendObjRequest("account.setOnline", params, &resp)

	return
}

// SetPushSettings - Change push settings.
// Parameters:
//   * deviceId - Unique device ID.
//   * settings - Push settings in a [vk.com/dev/push_settings|special format].
//   * key - Notification key.
//   * value - New value for the key in a [vk.com/dev/push_settings|special format].
func (a Account) SetPushSettings(deviceId string, settings string, key string, value []string) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["device_id"] = deviceId

	if settings != "" {
		params["settings"] = settings
	}

	if key != "" {
		params["key"] = key
	}

	if len(value) > 0 {
		params["value"] = SliceToString(value)
	}

	err = a.SendObjRequest("account.setPushSettings", params, &resp)

	return
}

// SetSilenceMode - Mutes push notifications for the set period of time.
// Parameters:
//   * deviceId - Unique device ID.
//   * time - Time in seconds for what notifications should be disabled. '-1' to disable forever.
//   * peerId - Destination ID. "For user: 'User ID', e.g. '12345'. For chat: '2000000000' + 'Chat ID', e.g. '2000000001'. For community: '- Community ID', e.g. '-12345'. "
//   * sound - '1' — to enable sound in this dialog, '0' — to disable sound. Only if 'peer_id' contains user or community ID.
func (a Account) SetSilenceMode(deviceId string, time int, peerId int, sound int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	if deviceId != "" {
		params["device_id"] = deviceId
	}

	if time > 0 {
		params["time"] = time
	}

	if peerId > 0 {
		params["peer_id"] = peerId
	}

	if sound > 0 {
		params["sound"] = sound
	}

	err = a.SendObjRequest("account.setSilenceMode", params, &resp)

	return
}

// Unban - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * ownerId - NO DESCRIPTION IN JSON SCHEMA
func (a Account) Unban(ownerId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	err = a.SendObjRequest("account.unban", params, &resp)

	return
}

// UnregisterDevice - Unsubscribes a device from push notifications.
// Parameters:
//   * deviceId - Unique device ID.
//   * sandbox - NO DESCRIPTION IN JSON SCHEMA
func (a Account) UnregisterDevice(deviceId string, sandbox bool) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	if deviceId != "" {
		params["device_id"] = deviceId
	}

	params["sandbox"] = sandbox

	err = a.SendObjRequest("account.unregisterDevice", params, &resp)

	return
}
