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
// `secure` group of responses
/////////////////////////////////////////////////////////////

// SecureChecktoken type represents `secure_checkToken_response` API response object
type SecureChecktoken objects.SecureTokenChecked

// SecureGetappbalance type represents `secure_getAppBalance_response` API response object
type SecureGetappbalance int // App balance

// SecureGetsmshistory type represents `secure_getSMSHistory_response` API response object
type SecureGetsmshistory objects.SecureSmsNotification

// SecureGettransactionshistory type represents `secure_getTransactionsHistory_response` API response object
type SecureGettransactionshistory objects.SecureTransaction

// SecureGetuserlevel type represents `secure_getUserLevel_response` API response object
type SecureGetuserlevel objects.SecureLevel

// SecureGiveeventsticker type represents `secure_giveEventSticker_response` API response object
type SecureGiveeventsticker interface{}

// SecureSendnotification type represents `secure_sendNotification_response` API response object
type SecureSendnotification int
