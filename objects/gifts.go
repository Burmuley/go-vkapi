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

package objects

/////////////////////////////////////////////////////////////////////////////////////////////////////////////
// `gifts` group of objects
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

// GiftsGift type represents `gifts_gift` API object
type GiftsGift struct {
	Date     int              `json:"date"`    // Date when gist has been sent in Unixtime
	FromId   int              `json:"from_id"` // Gift sender ID
	Gift     GiftsLayout      `json:"gift"`
	GiftHash string           `json:"gift_hash"` // Hash
	Id       int              `json:"id"`        // Gift ID
	Message  string           `json:"message"`   // Comment text
	Privacy  GiftsGiftPrivacy `json:"privacy"`
}

// GiftsLayout type represents `gifts_layout` API object
type GiftsLayout struct {
	Id       int    `json:"id"`        // Gift ID
	Thumb256 string `json:"thumb_256"` // URL of the preview image with 256 px in width
	Thumb48  string `json:"thumb_48"`  // URL of the preview image with 48 px in width
	Thumb96  string `json:"thumb_96"`  // URL of the preview image with 96 px in width
}

// GiftsGiftPrivacy type represents `gifts_gift_privacy` API object
type GiftsGiftPrivacy int // Gift privacy
