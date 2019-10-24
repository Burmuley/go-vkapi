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
// `fave` group of responses
/////////////////////////////////////////////////////////////

// FaveAddtag type represents `fave_addTag_response` API response object
type FaveAddtag *objects.FaveTag

// FaveGetpag type represents `fave_getPages_response` API response object
type FaveGetpag struct {
	Count int                 `json:"count"`
	Items []*objects.FavePage `json:"items"`
}

// FaveGettag type represents `fave_getTags_response` API response object
type FaveGettag struct {
	Count int                `json:"count"`
	Items []*objects.FaveTag `json:"items"`
}

// FaveGetExtended type represents `fave_get_extended_response` API response object
type FaveGetExtended struct {
	Count    int                      `json:"count"` // Total number
	Groups   []*objects.GroupsGroup   `json:"groups"`
	Items    []*objects.FaveBookmark  `json:"items"`
	Profiles []*objects.UsersUserFull `json:"profiles"`
}

// FaveGet type represents `fave_get_response` API response object
type FaveGet struct {
	Count int                     `json:"count"` // Total number
	Items []*objects.FaveBookmark `json:"items"`
}
