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

import "gitlab.com/Burmuley/go-vkapi/objects"

/////////////////////////////////////////////////////////////
// `prettyCards` group of responses
/////////////////////////////////////////////////////////////

// PrettyCardsDelete type represents `prettyCards_delete_response` API response object
type PrettyCardsDelete struct {
	CardId  string `json:"card_id"`  // Card ID of deleted pretty card
	Error   string `json:"error"`    // Error reason if error happened
	OwnerId int    `json:"owner_id"` // Owner ID of deleted pretty card
}

// PrettyCardsGetById type represents `prettyCards_getById_response` API response object
type PrettyCardsGetById objects.PrettyCardsPrettyCard

// PrettyCardsGet type represents `prettyCards_get_response` API response object
type PrettyCardsGet struct {
	Count int                             `json:"count"` // Total number
	Items []objects.PrettyCardsPrettyCard `json:"items"`
}

// PrettyCardsGetUploadURL type represents `prettyCards_getUploadURL_response` API response object
type PrettyCardsGetUploadURL string // Upload URL

// PrettyCardsCreate type represents `prettyCards_create_response` API response object
type PrettyCardsCreate struct {
	CardId  string `json:"card_id"`  // Card ID of created pretty card
	OwnerId int    `json:"owner_id"` // Owner ID of created pretty card
}

// PrettyCardsEdit type represents `prettyCards_edit_response` API response object
type PrettyCardsEdit struct {
	CardId  string `json:"card_id"`  // Card ID of edited pretty card
	OwnerId int    `json:"owner_id"` // Owner ID of edited pretty card
}
