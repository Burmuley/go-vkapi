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
// Code generator location: https://github.com/Burmuley/go-vkapi-gen                                       //
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

package objects

/////////////////////////////////////////////////////////////////////////////////////////////////////////////
// `secure` group of objects
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

// SecureLevel type represents `secure_level` API object
type SecureLevel struct {
	Level int `json:"level"` // Level
	Uid   int `json:"uid"`   // User ID
}

// SecureSmsNotification type represents `secure_sms_notification` API object
type SecureSmsNotification struct {
	AppId   string `json:"app_id"`  // Application ID
	Date    string `json:"date"`    // Date when message has been sent in Unixtime
	Id      string `json:"id"`      // Notification ID
	Message string `json:"message"` // Messsage text
	UserId  string `json:"user_id"` // User ID
}

// SecureTokenChecked type represents `secure_token_checked` API object
type SecureTokenChecked struct {
	Date    int            `json:"date"`    // Date when access_token has been generated in Unixtime
	Expire  int            `json:"expire"`  // Date when access_token will expire in Unixtime
	Success BaseOkResponse `json:"success"` // Returns if successfully processed
	UserId  int            `json:"user_id"` // User ID
}

// SecureTransaction type represents `secure_transaction` API object
type SecureTransaction struct {
	Date    int `json:"date"`     // Transaction date in Unixtime
	Id      int `json:"id"`       // Transaction ID
	UidFrom int `json:"uid_from"` // From ID
	UidTo   int `json:"uid_to"`   // To ID
	Votes   int `json:"votes"`    // Votes number
}
