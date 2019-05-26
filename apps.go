package go_vkapi

import "gitlab.com/Burmuley/go-vkapi/responses"

// Apps represents collection of methods related to VK Accounts
type Apps struct {
	*VKApi
}

// DeleteAppRequests deletes all request notifications from the current app
func (a *Apps) DeleteAppRequests() (responses.OkResponse, error) {
	panic("implement me!")
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
//  * type - List type. Possible values:
//      'invite' — available for invites (don't play the game)
//      'request' — available for request (play the game).
//      By default: 'invite'.
//  * fields - Additional profile fields, see [vk.com/dev/fields|description].
func (a *Apps) GetFriendsList() (responses.AppsGetFriendsList, error) {
	panic("implement me!")
}

// GetLeaderboard returns players rating in the game
//
// Parameters:
//  * type - Leaderboard type. Possible values:
//      'level' — by level
//      'points' — by mission points
//      'score' — by score ().
//  * global - Rating type. Possible values:
//      '1' — global rating among all players,, *'0' — rating among user friends.
func (a *Apps) GetLeaderboard() (responses.AppsGetLeaderboard, error) {
	params := map[string]string{"extended": "0"}
	panic("implement me!")
}

// GetLeaderboardExt returns extended players rating in the game
//
// Parameters:
//  * type - Leaderboard type. Possible values:
//      'level' — by level
//      'points' — by mission points
//      'score' — by score ().
//  * global - Rating type. Possible values:
//      '1' — global rating among all players,, *'0' — rating among user friends.
func (a *Apps) GetLeaderboardExt() (responses.AppsGetLeaderboardExt, error) {
	params := map[string]string{"extended": "1"}
	panic("implement me!")
}

// GetScopes returns scopes for auth
//
// Parameters:
//  * type -
func (a *Apps) GetScopes() (responses.AppsGetScopes, error) {
	panic("implement me!")
}

// GetScore returns user score in app
//
// Parameters:
//  * user_id -
func (a *Apps) GetScore() (responses.AppsGetScore, error) {
	panic("implement me!")
}

// SendRequest Sends a request to another user in an app that uses VK authorization
//
// Parameters:
//  * user_id - id of the user to send a request
//  * text - request text
//  * type - request type. Values: 'invite' – if the request is sent to a user who does not have the app installed,, 'request' – if a user has already installed the app
//  * name -
//  * key - special string key to be sent with the request
//  * separate -
func (a *Apps) SendRequest() (responses.AppsSendRequest, error) {
	panic("implement me!")
}
