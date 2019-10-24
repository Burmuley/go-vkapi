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
// `search` group of objects
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

// SearchHint type represents `search_hint` API object
type SearchHint struct {
	App         *AppsApp           `json:"app"`
	Description string             `json:"description"` // Object description
	Global      *BaseBoolInt       `json:"global"`      // Information whether the object has been found globally
	Group       *GroupsGroup       `json:"group"`
	Profile     *UsersUserMin      `json:"profile"`
	Section     *SearchHintSection `json:"section"`
	Type        *SearchHintType    `json:"type"`
}

// SearchHintSection type represents `search_hint_section` API object
type SearchHintSection string // Section title

// SearchHintType type represents `search_hint_type` API object
type SearchHintType string // Object type
