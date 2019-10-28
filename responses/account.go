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
// `account` group of responses
/////////////////////////////////////////////////////////////

// AccountChangepassword type represents `account_changePassword_response` API response object
type AccountChangepassword struct {
	Secret string `json:"secret"` // New secret
	Token  string `json:"token"`  // New token
}

// AccountGetactiveoffers type represents `account_getActiveOffers_response` API response object
type AccountGetactiveoffers struct {
	Count int                    `json:"count"` // Total number
	Items []objects.AccountOffer `json:"items"`
}

// AccountGetapppermissions type represents `account_getAppPermissions_response` API response object
type AccountGetapppermissions int // Permissions mask

// AccountGetbanned type represents `account_getBanned_response` API response object
type AccountGetbanned struct {
	Count    int                    `json:"count"` // Total number
	Groups   []objects.GroupsGroup  `json:"groups"`
	Items    []int                  `json:"items"`
	Profiles []objects.UsersUserMin `json:"profiles"`
}

// AccountGetcounters type represents `account_getCounters_response` API response object
type AccountGetcounters objects.AccountAccountCounters

// AccountGetinfo type represents `account_getInfo_response` API response object
type AccountGetinfo objects.AccountInfo

// AccountGetprofileinfo type represents `account_getProfileInfo_response` API response object
type AccountGetprofileinfo objects.AccountUserSettings

// AccountGetpushsettings type represents `account_getPushSettings_response` API response object
type AccountGetpushsettings objects.AccountPushSettings

// AccountSaveprofileinfo type represents `account_saveProfileInfo_response` API response object
type AccountSaveprofileinfo struct {
	Changed     objects.BaseBoolInt        `json:"changed"` // 1 if changes has been processed
	NameRequest objects.AccountNameRequest `json:"name_request"`
}
