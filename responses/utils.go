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
// `utils` group of responses
/////////////////////////////////////////////////////////////

// UtilsGetServerTime type represents `utils_getServerTime_response` API response object
type UtilsGetServerTime int // Time as Unixtime

// UtilsGetShortLink type represents `utils_getShortLink_response` API response object
type UtilsGetShortLink objects.UtilsShortLink

// UtilsCheckLink type represents `utils_checkLink_response` API response object
type UtilsCheckLink objects.UtilsLinkChecked

// UtilsResolveScreenName type represents `utils_resolveScreenName_response` API response object
type UtilsResolveScreenName objects.UtilsDomainResolved

// UtilsGetLinkStats type represents `utils_getLinkStats_response` API response object
type UtilsGetLinkStats objects.UtilsLinkStats

// UtilsGetLastShortenedLinks type represents `utils_getLastShortenedLinks_response` API response object
type UtilsGetLastShortenedLinks struct {
	Count int                              `json:"count"` // Total number of available results
	Items []objects.UtilsLastShortenedLink `json:"items"`
}

// UtilsGetLinkStatsExtended type represents `utils_getLinkStats_extended_response` API response object
type UtilsGetLinkStatsExtended objects.UtilsLinkStatsExtended
