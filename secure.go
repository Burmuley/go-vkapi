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
	"gitlab.com/Burmuley/go-vkapi/objects"
	"gitlab.com/Burmuley/go-vkapi/responses"
)

type Secure struct {
	*VKApi
}

/////////////////////////////////////////////////////////////
// `Secure` methods
/////////////////////////////////////////////////////////////

// AddAppEvent - Adds user activity information to an application
// Parameters:
//   * userId - ID of a user to save the data
//   * activityId - there are 2 default activities: , * 1 – level. Works similar to ,, * 2 – points, saves points amount, Any other value is for saving completed missions
//   * value - depends on activity_id: * 1 – number, current level number,, * 2 – number, current user's points amount, , Any other value is ignored
func (s *Secure) AddAppEvent(userId int, activityId int, value int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["user_id"] = userId

	params["activity_id"] = activityId

	if value > 0 {
		params["value"] = value
	}

	err = s.SendObjRequest("secure.addAppEvent", params, &resp)

	return
}

// CheckToken - Checks the user authentication in 'IFrame' and 'Flash' apps using the 'access_token' parameter.
// Parameters:
//   * token - client 'access_token'
//   * ip - user 'ip address'. Note that user may access using the 'ipv6' address, in this case it is required to transmit the 'ipv6' address. If not transmitted, the address will not be checked.
func (s *Secure) CheckToken(token string, ip string) (resp responses.SecureCheckToken, err error) {
	params := map[string]interface{}{}

	if token != "" {
		params["token"] = token
	}

	if ip != "" {
		params["ip"] = ip
	}

	err = s.SendObjRequest("secure.checkToken", params, &resp)

	return
}

// GetAppBalance - Returns payment balance of the application in hundredth of a vote.
func (s *Secure) GetAppBalance() (resp responses.SecureGetAppBalance, err error) {
	params := map[string]interface{}{}

	err = s.SendObjRequest("secure.getAppBalance", params, &resp)

	return
}

// GetSMSHistory - Shows a list of SMS notifications sent by the application using [vk.com/dev/secure.sendSMSNotification|secure.sendSMSNotification] method.
// Parameters:
//   * userId - NO DESCRIPTION IN JSON SCHEMA
//   * dateFrom - filter by start date. It is set as UNIX-time.
//   * dateTo - filter by end date. It is set as UNIX-time.
//   * limit - number of returned posts. By default — 1000.
func (s *Secure) GetSMSHistory(userId int, dateFrom int, dateTo int, limit int) (resp responses.SecureGetSMSHistory, err error) {
	params := map[string]interface{}{}

	if userId > 0 {
		params["user_id"] = userId
	}

	if dateFrom > 0 {
		params["date_from"] = dateFrom
	}

	if dateTo > 0 {
		params["date_to"] = dateTo
	}

	if limit > 0 {
		params["limit"] = limit
	}

	err = s.SendObjRequest("secure.getSMSHistory", params, &resp)

	return
}

// GetTransactionsHistory - Shows history of votes transaction between users and the application.
// Parameters:
//   * pType - NO DESCRIPTION IN JSON SCHEMA
//   * uidFrom - NO DESCRIPTION IN JSON SCHEMA
//   * uidTo - NO DESCRIPTION IN JSON SCHEMA
//   * dateFrom - NO DESCRIPTION IN JSON SCHEMA
//   * dateTo - NO DESCRIPTION IN JSON SCHEMA
//   * limit - NO DESCRIPTION IN JSON SCHEMA
func (s *Secure) GetTransactionsHistory(pType int, uidFrom int, uidTo int, dateFrom int, dateTo int, limit int) (resp responses.SecureGetTransactionsHistory, err error) {
	params := map[string]interface{}{}

	if pType > 0 {
		params["type"] = pType
	}

	if uidFrom > 0 {
		params["uid_from"] = uidFrom
	}

	if uidTo > 0 {
		params["uid_to"] = uidTo
	}

	if dateFrom > 0 {
		params["date_from"] = dateFrom
	}

	if dateTo > 0 {
		params["date_to"] = dateTo
	}

	if limit > 0 {
		params["limit"] = limit
	}

	err = s.SendObjRequest("secure.getTransactionsHistory", params, &resp)

	return
}

// GetUserLevel - Returns one of the previously set game levels of one or more users in the application.
// Parameters:
//   * userIds - NO DESCRIPTION IN JSON SCHEMA
func (s *Secure) GetUserLevel(userIds []int) (resp responses.SecureGetUserLevel, err error) {
	params := map[string]interface{}{}

	params["user_ids"] = SliceToString(userIds)

	err = s.SendObjRequest("secure.getUserLevel", params, &resp)

	return
}

// GiveEventSticker - Opens the game achievement and gives the user a sticker
// Parameters:
//   * userIds - NO DESCRIPTION IN JSON SCHEMA
//   * achievementId - NO DESCRIPTION IN JSON SCHEMA
func (s *Secure) GiveEventSticker(userIds []int, achievementId int) (resp responses.SecureGiveEventSticker, err error) {
	params := map[string]interface{}{}

	params["user_ids"] = SliceToString(userIds)

	params["achievement_id"] = achievementId

	err = s.SendObjRequest("secure.giveEventSticker", params, &resp)

	return
}

// SendNotification - Sends notification to the user.
// Parameters:
//   * userIds - NO DESCRIPTION IN JSON SCHEMA
//   * userId - NO DESCRIPTION IN JSON SCHEMA
//   * message - notification text which should be sent in 'UTF-8' encoding ('254' characters maximum).
func (s *Secure) SendNotification(userIds []int, userId int, message string) (resp responses.SecureSendNotification, err error) {
	params := map[string]interface{}{}

	if len(userIds) > 0 {
		params["user_ids"] = SliceToString(userIds)
	}

	if userId > 0 {
		params["user_id"] = userId
	}

	params["message"] = message

	err = s.SendObjRequest("secure.sendNotification", params, &resp)

	return
}

// SendSMSNotification - Sends 'SMS' notification to a user's mobile device.
// Parameters:
//   * userId - ID of the user to whom SMS notification is sent. The user shall allow the application to send him/her notifications (, +1).
//   * message - 'SMS' text to be sent in 'UTF-8' encoding. Only Latin letters and numbers are allowed. Maximum size is '160' characters.
func (s *Secure) SendSMSNotification(userId int, message string) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["user_id"] = userId

	params["message"] = message

	err = s.SendObjRequest("secure.sendSMSNotification", params, &resp)

	return
}

// SetCounter - Sets a counter which is shown to the user in bold in the left menu.
// Parameters:
//   * counters - NO DESCRIPTION IN JSON SCHEMA
//   * userId - NO DESCRIPTION IN JSON SCHEMA
//   * counter - counter value.
//   * increment - NO DESCRIPTION IN JSON SCHEMA
func (s *Secure) SetCounter(counters []string, userId int, counter int, increment bool) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	if len(counters) > 0 {
		params["counters"] = SliceToString(counters)
	}

	if userId > 0 {
		params["user_id"] = userId
	}

	if counter > 0 {
		params["counter"] = counter
	}

	params["increment"] = increment

	err = s.SendObjRequest("secure.setCounter", params, &resp)

	return
}
