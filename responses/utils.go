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
// `utils` group of responses
/////////////////////////////////////////////////////////////

// UtilsChecklink type represents `utils_checkLink_response` API response object
type UtilsChecklink objects.UtilsLinkChecked

// UtilsGetlastshortenedlinks type represents `utils_getLastShortenedLinks_response` API response object
type UtilsGetlastshortenedlinks struct {
	Count int                              `json:"count"` // Total number of available results
	Items []objects.UtilsLastShortenedLink `json:"items"`
}

// UtilsGetlinkstatsExtended type represents `utils_getLinkStats_extended_response` API response object
type UtilsGetlinkstatsExtended objects.UtilsLinkStatsExtended

// UtilsGetlinkstats type represents `utils_getLinkStats_response` API response object
type UtilsGetlinkstats objects.UtilsLinkStats

// UtilsGetservertime type represents `utils_getServerTime_response` API response object
type UtilsGetservertime int // Time as Unixtime

// UtilsGetshortlink type represents `utils_getShortLink_response` API response object
type UtilsGetshortlink objects.UtilsShortLink

// UtilsResolvescreenname type represents `utils_resolveScreenName_response` API response object
type UtilsResolvescreenname objects.UtilsDomainResolved
