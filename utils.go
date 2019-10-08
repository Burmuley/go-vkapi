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
// Source schema can be found at https://github.com/VKCOM/vk-api-schema/blob/master/methods.json           //
// Code generator location: https://gitlab.com/Burmuley/go-vkapi-gen                                       //
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

package go_vkapi

import (
	"gitlab.com/Burmuley/go-vkapi/objects"
	"gitlab.com/Burmuley/go-vkapi/responses"
)

type Utils struct {
	*VKApi
}

/////////////////////////////////////////////////////////////
// `Utils` methods
/////////////////////////////////////////////////////////////

// CheckLink - Checks whether a link is blocked in VK.
// Parameters:
//   * url - Link to check (e.g., 'http://google.com').
func (u *Utils) CheckLink(url string) (resp responses.UtilsCheckLink, err error) {
	params := map[string]interface{}{}

	params["url"] = url

	err = u.SendObjRequest("utils.checkLink", params, &resp)

	return
}

// DeleteFromLastShortened - Deletes shortened link from user's list.
// Parameters:
//   * key - Link key (characters after vk.cc/).
func (u *Utils) DeleteFromLastShortened(key string) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["key"] = key

	err = u.SendObjRequest("utils.deleteFromLastShortened", params, &resp)

	return
}

// GetLastShortenedLinks - Returns a list of user's shortened links.
// Parameters:
//   * count - Number of links to return.
//   * offset - Offset needed to return a specific subset of links.
func (u *Utils) GetLastShortenedLinks(count int, offset int) (resp responses.UtilsGetLastShortenedLinks, err error) {
	params := map[string]interface{}{}

	if count > 0 {
		params["count"] = count
	}

	if offset > 0 {
		params["offset"] = offset
	}

	err = u.SendObjRequest("utils.getLastShortenedLinks", params, &resp)

	return
}

// GetLinkStats - Returns stats data for shortened link.
// Parameters:
//   * key - Link key (characters after vk.cc/).
//   * source - Source of scope
//   * accessKey - Access key for private link stats.
//   * interval - Interval.
//   * intervalsCount - Number of intervals to return.
//   * extended - 1 — to return extended stats data (sex, age, geo). 0 — to return views number only.
func (u *Utils) GetLinkStats(key string, source string, accessKey string, interval string, intervalsCount int) (resp responses.UtilsGetLinkStats, err error) {
	params := map[string]interface{}{}
	params["extended"] = "0"

	params["key"] = key

	if source != "" {
		params["source"] = source
	}

	if accessKey != "" {
		params["access_key"] = accessKey
	}

	if interval != "" {
		params["interval"] = interval
	}

	if intervalsCount > 0 {
		params["intervals_count"] = intervalsCount
	}

	err = u.SendObjRequest("utils.getLinkStats", params, &resp)

	return
}

// GetLinkStatsExtended - Returns stats data for shortened link.
// Parameters:
//   * key - Link key (characters after vk.cc/).
//   * source - Source of scope
//   * accessKey - Access key for private link stats.
//   * interval - Interval.
//   * intervalsCount - Number of intervals to return.
//   * extended - 1 — to return extended stats data (sex, age, geo). 0 — to return views number only.
func (u *Utils) GetLinkStatsExtended(key string, source string, accessKey string, interval string, intervalsCount int) (resp responses.UtilsGetLinkStatsExtended, err error) {
	params := map[string]interface{}{}
	params["extended"] = "1"

	params["key"] = key

	if source != "" {
		params["source"] = source
	}

	if accessKey != "" {
		params["access_key"] = accessKey
	}

	if interval != "" {
		params["interval"] = interval
	}

	if intervalsCount > 0 {
		params["intervals_count"] = intervalsCount
	}

	err = u.SendObjRequest("utils.getLinkStats", params, &resp)

	return
}

// GetServerTime - Returns the current time of the VK server.
func (u *Utils) GetServerTime() (resp responses.UtilsGetServerTime, err error) {
	params := map[string]interface{}{}

	err = u.SendObjRequest("utils.getServerTime", params, &resp)

	return
}

// GetShortLink - Allows to receive a link shortened via vk.cc.
// Parameters:
//   * url - URL to be shortened.
//   * private - 1 — private stats, 0 — public stats.
func (u *Utils) GetShortLink(url string, private bool) (resp responses.UtilsGetShortLink, err error) {
	params := map[string]interface{}{}

	params["url"] = url

	params["private"] = private

	err = u.SendObjRequest("utils.getShortLink", params, &resp)

	return
}

// ResolveScreenName - Detects a type of object (e.g., user, community, application) and its ID by screen name.
// Parameters:
//   * screenName - Screen name of the user, community (e.g., 'apiclub,' 'andrew', or 'rules_of_war'), or application.
func (u *Utils) ResolveScreenName(screenName string) (resp responses.UtilsResolveScreenName, err error) {
	params := map[string]interface{}{}

	params["screen_name"] = screenName

	err = u.SendObjRequest("utils.resolveScreenName", params, &resp)

	return
}
