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
// `prettyCards` group of objects
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

// PrettyCardsPrettyCard type represents `prettyCards_prettyCard` API object
type PrettyCardsPrettyCard struct {
	Button     string      `json:"button"`      // Button key
	ButtonText string      `json:"button_text"` // Button text in current language
	CardId     string      `json:"card_id"`     // Card ID (long int returned as string)
	Images     []BaseImage `json:"images"`
	LinkUrl    string      `json:"link_url"`  // Link URL
	Photo      string      `json:"photo"`     // Photo ID (format "<owner_id>_<media_id>")
	Price      string      `json:"price"`     // Price if set (decimal number returned as string)
	PriceOld   string      `json:"price_old"` // Old price if set (decimal number returned as string)
	Title      string      `json:"title"`     // Title
}
