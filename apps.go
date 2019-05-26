package go_vkapi

import (
	"gitlab.com/Burmuley/go-vkapi/objects"
	"gitlab.com/Burmuley/go-vkapi/responses"
)

// Apps represents collection of methods related to VK Accounts
type Apps struct {
	*VKApi
}

// DeleteAppRequests deletes all request notifications from the current app
func (a *Apps) DeleteAppRequests() (responses.OkResponse, error) {
	var resp responses.OkResponse

	if err := a.SendObjRequest("apps.getFriendsList", map[string]string{}, &resp); err != nil {
		return responses.OkResponse(0), err
	}

	return resp, nil
}

// Get returns applications data
//
// Parameters:
//  * app_id - Application ID
//  * app_ids - List of application ID
//  * platform - platform. Possible values:
//      'ios' — iOS
//      'android' — Android
//      'winphone' — Windows Phone
//      'web' — приложения на vk.com.
//      By default: 'web'.
//  * extended -
//  * return_friends -
//  * fields - Profile fields to return. Sample values:
//      'nickname', 'screen_name''sex', 'bdate' (birthdate), 'city', 'country', 'timezone',
//      'photo', 'photo_medium', 'photo_big', 'has_mobile', 'contacts', 'education', 'online',
//      'counters', 'relation', 'last_seen', 'activity', 'can_write_private_message', 'can_see_all_posts',
//      'can_post', 'universities', (only if return_friends - 1)
//  * name_case - Case for declension of user name and surname:
//      'nom' — nominative (default)
//      'gen' — genitive
//      'dat' — dative
//      'acc' — accusative
//      'ins' — instrumental
//      'abl' — prepositional. (only if 'return_friends' = '1')
func (a *Apps) Get() (responses.AppsGet, error) {
	panic("implement me!")
}

// GetCatalog returns a list of applications (apps) available to users in the App Catalog
//
// Parameters:
//  * sort - Sort order:
//      'popular_today' — popular for one day (default)
//      'visitors' — by visitors number
//      'create_date' — by creation date
//      'growth_rate' — by growth rate
//      'popular_week' — popular for one week
//  * offset - Offset required to return a specific subset of apps.
//  * count - Number of apps to return.
//  * platform -
//  * extended - '1' — to return additional fields 'screenshots', 'MAU', 'catalog_position', and 'international'.
//    If set, 'count' must be less than or equal to '100'. '0' — not to return additional fields (default).
//  * return_friends - NO DESCRIPTION (bool)
//  * fields - fields to return
//  * name_case -
//  * q - Search query string.
//  * genre_id - NO DESCRIPTION (int)
//  * filter - 'installed' — to return list of installed apps (only for mobile platform).
func (a *Apps) GetCatalog() (responses.AppsGetCatalog, error) {
	panic("implement me!")
}

// GetFriendsList creates friends list for requests and invites in current app
//
// Parameters:
//  * extended - NO DESCRIPTION (bool)
//  * count - List size
//  * offset - NO DESCRIPTION (int)
//  * listType - List type. Possible values:
//      'invite' — available for invites (don't play the game)
//      'request' — available for request (play the game).
//      By default: 'invite'.
//  * fields - Additional profile fields, see [vk.com/dev/fields|description].
func (a *Apps) GetFriendsList(extended, count, offset int, listType string, fields []objects.UsersFields) (responses.AppsGetFriendsList, error) {
	params := map[string]string{"extended": string(extended)}

	if count > 0 {
		params["count"] = string(count)
	}

	if offset > 0 {
		params["offset"] = string(offset)
	}

	if len(listType) > 0 {
		params["type"] = listType
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	resp := responses.AppsGetFriendsList{}

	if err := a.SendObjRequest("apps.getFriendsList", params, &resp); err != nil {
		return responses.AppsGetFriendsList{}, err
	}

	return resp, nil
}

// GetLeaderboard returns players rating in the game
//
// Parameters:
//  * boardType - Leaderboard type. Possible values:
//      'level' — by level
//      'points' — by mission points
//      'score' — by score ().
//  * global - Rating type. Possible values:
//      '1' — global rating among all players
//      '0' — rating among user friends.
func (a *Apps) GetLeaderboard(boardType string, global int) (responses.AppsGetLeaderboard, error) {
	params := map[string]string{"extended": "0", "type": boardType}

	if global > -1 || global <= 1 {
		params["global"] = string(global)
	}

	resp := responses.AppsGetLeaderboard{}

	if err := a.SendObjRequest("apps.getLeaderboard", params, &resp); err != nil {
		return responses.AppsGetLeaderboard{}, err
	}

	return resp, nil
}

// GetLeaderboardExt returns extended players rating in the game
//
// Parameters:
//  * type - Leaderboard type. Possible values:
//      'level' — by level
//      'points' — by mission points
//      'score' — by score ().
//  * global - Rating type. Possible values:
//      '1' — global rating among all players
//      '0' — rating among user friends.
func (a *Apps) GetLeaderboardExt(boardType string, global int) (responses.AppsGetLeaderboardExt, error) {
	params := map[string]string{"extended": "1", "type": boardType}

	if global > -1 || global <= 1 {
		params["global"] = string(global)
	}

	resp := responses.AppsGetLeaderboardExt{}

	if err := a.SendObjRequest("apps.getLeaderboard", params, &resp); err != nil {
		return responses.AppsGetLeaderboardExt{}, err
	}

	return resp, nil
}

// GetScopes returns scopes for auth
//
// Parameters:
//  * type - type of the scope. By default 'user'
func (a *Apps) GetScopes(scopeType string) (responses.AppsGetScopes, error) {
	if len(scopeType) == 0 {
		scopeType = "user"
	}

	params := map[string]string{"type": scopeType}

	resp := responses.AppsGetScopes{}

	if err := a.SendObjRequest("apps.getScopes", params, &resp); err != nil {
		return responses.AppsGetScopes{}, err
	}

	return resp, nil
}

// GetScore returns user score in app
//
// Parameters:
//  * userId - User's ID
func (a *Apps) GetScore(userId int) (responses.AppsGetScore, error) {
	params := map[string]string{"user_id": string(userId)}

	var resp responses.AppsGetScore

	if err := a.SendObjRequest("apps.getScore", params, &resp); err != nil {
		return responses.AppsGetScore(0), err
	}

	return resp, nil
}

// SendRequest Sends a request to another user in an app that uses VK authorization
//
// Parameters:
//  * userId - id of the user to send a request
//  * text - request text
//  * requestType - request type. Values: 'invite' – if the request is sent to a user who does not have the app installed,, 'request' – if a user has already installed the app
//  * name - unique request name
//  * key - special string key to be sent with the request
//  * separate - controls if request can be grouped with other ones with the same name
func (a *Apps) SendRequest(userId, separate int, text, requestType, name, key string) (responses.AppsSendRequest, error) {
	params := map[string]string{"user_id": string(userId), "separate": string(separate)}

	if len(text) > 0 {
		params["text"] = text
	}

	if len(requestType) > 0 {
		params["type"] = requestType
	}

	if len(name) > 0 {
		params["name"] = name
	}

	if len(key) > 0 {
		params["key"] = key
	}

	var resp responses.AppsSendRequest

	if err := a.SendObjRequest("apps.sendRequest", params, &resp); err != nil {
		return responses.AppsSendRequest(0), err
	}

	return resp, nil
}
