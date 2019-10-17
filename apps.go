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

type Apps struct {
	*VKApi
}

/////////////////////////////////////////////////////////////
// `Apps` methods
/////////////////////////////////////////////////////////////

// DeleteAppRequests - Deletes all request notifications from the current app.
func (a Apps) DeleteAppRequests() (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	err = a.SendObjRequest("apps.deleteAppRequests", params, &resp)

	return
}

// Get - Returns applications data.
// Parameters:
//   * appId - Application ID
//   * appIds - List of application ID
//   * platform - platform. Possible values: *'ios' — iOS,, *'android' — Android,, *'winphone' — Windows Phone,, *'web' — приложения на vk.com. By default: 'web'.
//   * extended - NO DESCRIPTION IN JSON SCHEMA
//   * returnFriends - NO DESCRIPTION IN JSON SCHEMA
//   * fields - Profile fields to return. Sample values: 'nickname', 'screen_name', 'sex', 'bdate' (birthdate), 'city', 'country', 'timezone', 'photo', 'photo_medium', 'photo_big', 'has_mobile', 'contacts', 'education', 'online', 'counters', 'relation', 'last_seen', 'activity', 'can_write_private_message', 'can_see_all_posts', 'can_post', 'universities', (only if return_friends - 1)
//   * nameCase - Case for declension of user name and surname: 'nom' — nominative (default),, 'gen' — genitive,, 'dat' — dative,, 'acc' — accusative,, 'ins' — instrumental,, 'abl' — prepositional. (only if 'return_friends' = '1')
func (a Apps) Get(appId int, appIds []string, platform string, returnFriends bool, fields []objects.UsersFields, nameCase string) (resp responses.AppsGet, err error) {
	params := map[string]interface{}{}

	if appId > 0 {
		params["app_id"] = appId
	}

	if len(appIds) > 0 {
		params["app_ids"] = SliceToString(appIds)
	}

	if platform != "" {
		params["platform"] = platform
	}

	params["return_friends"] = returnFriends

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	if nameCase != "" {
		params["name_case"] = nameCase
	}

	err = a.SendObjRequest("apps.get", params, &resp)

	return
}

// GetCatalog - Returns a list of applications (apps) available to users in the App Catalog.
// Parameters:
//   * sort - Sort order: 'popular_today' — popular for one day (default), 'visitors' — by visitors number , 'create_date' — by creation date, 'growth_rate' — by growth rate, 'popular_week' — popular for one week
//   * offset - Offset required to return a specific subset of apps.
//   * count - Number of apps to return.
//   * platform - NO DESCRIPTION IN JSON SCHEMA
//   * extended - '1' — to return additional fields 'screenshots', 'MAU', 'catalog_position', and 'international'. If set, 'count' must be less than or equal to '100'. '0' — not to return additional fields (default).
//   * returnFriends - NO DESCRIPTION IN JSON SCHEMA
//   * fields - NO DESCRIPTION IN JSON SCHEMA
//   * nameCase - NO DESCRIPTION IN JSON SCHEMA
//   * q - Search query string.
//   * genreId - NO DESCRIPTION IN JSON SCHEMA
//   * filter - 'installed' — to return list of installed apps (only for mobile platform).
func (a Apps) GetCatalog(sort string, offset int, count int, platform string, returnFriends bool, fields []objects.UsersFields, nameCase string, q string, genreId int, filter string) (resp responses.AppsGetCatalog, err error) {
	params := map[string]interface{}{}

	if sort != "" {
		params["sort"] = sort
	}

	if offset > 0 {
		params["offset"] = offset
	}

	params["count"] = count

	if platform != "" {
		params["platform"] = platform
	}

	params["return_friends"] = returnFriends

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	if nameCase != "" {
		params["name_case"] = nameCase
	}

	if q != "" {
		params["q"] = q
	}

	if genreId > 0 {
		params["genre_id"] = genreId
	}

	if filter != "" {
		params["filter"] = filter
	}

	err = a.SendObjRequest("apps.getCatalog", params, &resp)

	return
}

// GetFriendsList - Creates friends list for requests and invites in current app.
// Parameters:
//   * extended - NO DESCRIPTION IN JSON SCHEMA
//   * count - List size.
//   * offset - NO DESCRIPTION IN JSON SCHEMA
//   * pType - List type. Possible values: * 'invite' — available for invites (don't play the game),, * 'request' — available for request (play the game). By default: 'invite'.
//   * fields - Additional profile fields, see [vk.com/dev/fields|description].
func (a Apps) GetFriendsList(count int, offset int, pType string, fields []objects.UsersFields) (resp responses.AppsGetFriendsList, err error) {
	params := map[string]interface{}{}

	if count > 0 {
		params["count"] = count
	}

	if offset > 0 {
		params["offset"] = offset
	}

	if pType != "" {
		params["type"] = pType
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	err = a.SendObjRequest("apps.getFriendsList", params, &resp)

	return
}

// GetLeaderboard - Returns players rating in the game.
// Parameters:
//   * pType - Leaderboard type. Possible values: *'level' — by level,, *'points' — by mission points,, *'score' — by score ().
//   * global - Rating type. Possible values: *'1' — global rating among all players,, *'0' — rating among user friends.
//   * extended - 1 — to return additional info about users
func (a Apps) GetLeaderboard(pType string, global bool) (resp responses.AppsGetLeaderboard, err error) {
	params := map[string]interface{}{}
	params["extended"] = "0"

	params["type"] = pType

	params["global"] = global

	err = a.SendObjRequest("apps.getLeaderboard", params, &resp)

	return
}

// GetLeaderboardExtended - Returns players rating in the game.
// Parameters:
//   * pType - Leaderboard type. Possible values: *'level' — by level,, *'points' — by mission points,, *'score' — by score ().
//   * global - Rating type. Possible values: *'1' — global rating among all players,, *'0' — rating among user friends.
//   * extended - 1 — to return additional info about users
func (a Apps) GetLeaderboardExtended(pType string, global bool) (resp responses.AppsGetLeaderboardExtended, err error) {
	params := map[string]interface{}{}
	params["extended"] = "1"

	params["type"] = pType

	params["global"] = global

	err = a.SendObjRequest("apps.getLeaderboard", params, &resp)

	return
}

// GetScopes - Returns scopes for auth
// Parameters:
//   * pType - NO DESCRIPTION IN JSON SCHEMA
func (a Apps) GetScopes(pType string) (resp responses.AppsGetScopes, err error) {
	params := map[string]interface{}{}

	if pType != "" {
		params["type"] = pType
	}

	err = a.SendObjRequest("apps.getScopes", params, &resp)

	return
}

// GetScore - Returns user score in app
// Parameters:
//   * userId - NO DESCRIPTION IN JSON SCHEMA
func (a Apps) GetScore(userId int) (resp responses.AppsGetScore, err error) {
	params := map[string]interface{}{}

	params["user_id"] = userId

	err = a.SendObjRequest("apps.getScore", params, &resp)

	return
}

// SendRequest - Sends a request to another user in an app that uses VK authorization.
// Parameters:
//   * userId - id of the user to send a request
//   * text - request text
//   * pType - request type. Values: 'invite' – if the request is sent to a user who does not have the app installed,, 'request' – if a user has already installed the app
//   * name - NO DESCRIPTION IN JSON SCHEMA
//   * key - special string key to be sent with the request
//   * separate - NO DESCRIPTION IN JSON SCHEMA
func (a Apps) SendRequest(userId int, text string, pType string, name string, key string, separate bool) (resp responses.AppsSendRequest, err error) {
	params := map[string]interface{}{}

	params["user_id"] = userId

	if text != "" {
		params["text"] = text
	}

	if pType != "" {
		params["type"] = pType
	}

	if name != "" {
		params["name"] = name
	}

	if key != "" {
		params["key"] = key
	}

	params["separate"] = separate

	err = a.SendObjRequest("apps.sendRequest", params, &resp)

	return
}
