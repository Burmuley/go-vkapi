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
	"gitlab.com/Burmuley/go-vkapi/responses"
)

type Utils struct {
	*VKApi
}

/////////////////////////////////////////////////////////////
// `Utils` methods
/////////////////////////////////////////////////////////////

// Checklink - Checks whether a link is blocked in VK.
// Parameters:
//   * url - Link to check (e.g., 'http://google.com').
func (u *Utils) Checklink(url string) (resp responses.UtilsChecklink, err error) {
	params := map[string]interface{}{}

	params["url"] = url

	err = u.SendObjRequest("utils.checkLink", params, &resp)

	return
}

// Deletefromlastshortened - Deletes shortened link from user's list.
// Parameters:
//   * key - Link key (characters after vk.cc/).
func (u *Utils) Deletefromlastshortened(key string) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["key"] = key

	err = u.SendObjRequest("utils.deleteFromLastShortened", params, &resp)

	return
}

// Getlastshortenedlinks - Returns a list of user's shortened links.
// Parameters:
//   * count - Number of links to return.
//   * offset - Offset needed to return a specific subset of links.
func (u *Utils) Getlastshortenedlinks(count int, offset int) (resp responses.UtilsGetlastshortenedlinks, err error) {
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

// Getlinkstats - Returns stats data for shortened link.
// Parameters:
//   * key - Link key (characters after vk.cc/).
//   * source - Source of scope
//   * accessKey - Access key for private link stats.
//   * interval - Interval.
//   * intervalsCount - Number of intervals to return.
//   * extended - 1 — to return extended stats data (sex, age, geo). 0 — to return views number only.
func (u *Utils) Getlinkstats(key string, source string, accessKey string, interval string, intervalsCount int) (resp responses.UtilsGetlinkstats, err error) {
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

// GetlinkstatsExtended - Returns stats data for shortened link.
// Parameters:
//   * key - Link key (characters after vk.cc/).
//   * source - Source of scope
//   * accessKey - Access key for private link stats.
//   * interval - Interval.
//   * intervalsCount - Number of intervals to return.
//   * extended - 1 — to return extended stats data (sex, age, geo). 0 — to return views number only.
func (u *Utils) GetlinkstatsExtended(key string, source string, accessKey string, interval string, intervalsCount int) (resp responses.UtilsGetlinkstatsExtended, err error) {
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

// Getservertime - Returns the current time of the VK server.
func (u *Utils) Getservertime() (resp responses.UtilsGetservertime, err error) {
	params := map[string]interface{}{}

	err = u.SendObjRequest("utils.getServerTime", params, &resp)

	return
}

// Getshortlink - Allows to receive a link shortened via vk.cc.
// Parameters:
//   * url - URL to be shortened.
//   * private - 1 — private stats, 0 — public stats.
func (u *Utils) Getshortlink(url string, private bool) (resp responses.UtilsGetshortlink, err error) {
	params := map[string]interface{}{}

	params["url"] = url

	params["private"] = private

	err = u.SendObjRequest("utils.getShortLink", params, &resp)

	return
}

// Resolvescreenname - Detects a type of object (e.g., user, community, application) and its ID by screen name.
// Parameters:
//   * screenName - Screen name of the user, community (e.g., 'apiclub,' 'andrew', or 'rules_of_war'), or application.
func (u *Utils) Resolvescreenname(screenName string) (resp responses.UtilsResolvescreenname, err error) {
	params := map[string]interface{}{}

	params["screen_name"] = screenName

	err = u.SendObjRequest("utils.resolveScreenName", params, &resp)

	return
}
