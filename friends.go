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

type Friends struct {
	*VKApi
}

/////////////////////////////////////////////////////////////
// `Friends` methods
/////////////////////////////////////////////////////////////

// Add - Approves or creates a friend request.
// Parameters:
//   * userId - ID of the user whose friend request will be approved or to whom a friend request will be sent.
//   * text - Text of the message (up to 500 characters) for the friend request, if any.
//   * follow - '1' to pass an incoming request to followers list.
func (f Friends) Add(userId int, text string, follow bool) (resp responses.FriendsAdd, err error) {
	params := map[string]interface{}{}

	if userId > 0 {
		params["user_id"] = userId
	}

	if text != "" {
		params["text"] = text
	}

	params["follow"] = follow

	err = f.SendObjRequest("friends.add", params, &resp)

	return
}

// AddList - Creates a new friend list for the current user.
// Parameters:
//   * name - Name of the friend list.
//   * userIds - IDs of users to be added to the friend list.
func (f Friends) AddList(name string, userIds []int) (resp responses.FriendsAddList, err error) {
	params := map[string]interface{}{}

	params["name"] = name

	if len(userIds) > 0 {
		params["user_ids"] = SliceToString(userIds)
	}

	err = f.SendObjRequest("friends.addList", params, &resp)

	return
}

// AreFriends - Checks the current user's friendship status with other specified users.
// Parameters:
//   * userIds - IDs of the users whose friendship status to check.
//   * needSign - '1' — to return 'sign' field. 'sign' is md5("{id}_{user_id}_{friends_status}_{application_secret}"), where id is current user ID. This field allows to check that data has not been modified by the client. By default: '0'.
func (f Friends) AreFriends(userIds []int, needSign bool) (resp responses.FriendsAreFriends, err error) {
	params := map[string]interface{}{}

	params["user_ids"] = SliceToString(userIds)

	params["need_sign"] = needSign

	err = f.SendObjRequest("friends.areFriends", params, &resp)

	return
}

// Delete - Declines a friend request or deletes a user from the current user's friend list.
// Parameters:
//   * userId - ID of the user whose friend request is to be declined or who is to be deleted from the current user's friend list.
func (f Friends) Delete(userId int) (resp responses.FriendsDelete, err error) {
	params := map[string]interface{}{}

	if userId > 0 {
		params["user_id"] = userId
	}

	err = f.SendObjRequest("friends.delete", params, &resp)

	return
}

// DeleteAllRequests - Marks all incoming friend requests as viewed.
func (f Friends) DeleteAllRequests() (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	err = f.SendObjRequest("friends.deleteAllRequests", params, &resp)

	return
}

// DeleteList - Deletes a friend list of the current user.
// Parameters:
//   * listId - ID of the friend list to delete.
func (f Friends) DeleteList(listId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["list_id"] = listId

	err = f.SendObjRequest("friends.deleteList", params, &resp)

	return
}

// Edit - Edits the friend lists of the selected user.
// Parameters:
//   * userId - ID of the user whose friend list is to be edited.
//   * listIds - IDs of the friend lists to which to add the user.
func (f Friends) Edit(userId int, listIds []int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["user_id"] = userId

	if len(listIds) > 0 {
		params["list_ids"] = SliceToString(listIds)
	}

	err = f.SendObjRequest("friends.edit", params, &resp)

	return
}

// EditList - Edits a friend list of the current user.
// Parameters:
//   * name - Name of the friend list.
//   * listId - Friend list ID.
//   * userIds - IDs of users in the friend list.
//   * addUserIds - (Applies if 'user_ids' parameter is not set.), User IDs to add to the friend list.
//   * deleteUserIds - (Applies if 'user_ids' parameter is not set.), User IDs to delete from the friend list.
func (f Friends) EditList(name string, listId int, userIds []int, addUserIds []int, deleteUserIds []int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	if name != "" {
		params["name"] = name
	}

	params["list_id"] = listId

	if len(userIds) > 0 {
		params["user_ids"] = SliceToString(userIds)
	}

	if len(addUserIds) > 0 {
		params["add_user_ids"] = SliceToString(addUserIds)
	}

	if len(deleteUserIds) > 0 {
		params["delete_user_ids"] = SliceToString(deleteUserIds)
	}

	err = f.SendObjRequest("friends.editList", params, &resp)

	return
}

// Get - Returns a list of user IDs or detailed information about a user's friends.
// Parameters:
//   * userId - User ID. By default, the current user ID.
//   * order - Sort order: , 'name' — by name (enabled only if the 'fields' parameter is used), 'hints' — by rating, similar to how friends are sorted in My friends section, , This parameter is available only for [vk.com/dev/standalone|desktop applications].
//   * listId - ID of the friend list returned by the [vk.com/dev/friends.getLists|friends.getLists] method to be used as the source. This parameter is taken into account only when the uid parameter is set to the current user ID. This parameter is available only for [vk.com/dev/standalone|desktop applications].
//   * count - Number of friends to return.
//   * offset - Offset needed to return a specific subset of friends.
//   * fields - Profile fields to return. Sample values: 'uid', 'first_name', 'last_name', 'nickname', 'sex', 'bdate' (birthdate), 'city', 'country', 'timezone', 'photo', 'photo_medium', 'photo_big', 'domain', 'has_mobile', 'rate', 'contacts', 'education'.
//   * nameCase - Case for declension of user name and surname: , 'nom' — nominative (default) , 'gen' — genitive , 'dat' — dative , 'acc' — accusative , 'ins' — instrumental , 'abl' — prepositional
//   * ref - NO DESCRIPTION IN JSON SCHEMA
func (f Friends) Get(userId int, order string, listId int, count int, offset int, fields []objects.UsersFields, nameCase string, ref string) (resp responses.FriendsGet, err error) {
	params := map[string]interface{}{}

	if userId > 0 {
		params["user_id"] = userId
	}

	if order != "" {
		params["order"] = order
	}

	if listId > 0 {
		params["list_id"] = listId
	}

	if count > 0 {
		params["count"] = count
	}

	if offset > 0 {
		params["offset"] = offset
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	if nameCase != "" {
		params["name_case"] = nameCase
	}

	if ref != "" {
		params["ref"] = ref
	}

	err = f.SendObjRequest("friends.get", params, &resp)

	return
}

// GetAppUsers - Returns a list of IDs of the current user's friends who installed the application.
func (f Friends) GetAppUsers() (resp responses.FriendsGetAppUsers, err error) {
	params := map[string]interface{}{}

	err = f.SendObjRequest("friends.getAppUsers", params, &resp)

	return
}

// GetByPhones - Returns a list of the current user's friends whose phone numbers, validated or specified in a profile, are in a given list.
// Parameters:
//   * phones - List of phone numbers in MSISDN format (maximum 1000). Example: "+79219876543,+79111234567"
//   * fields - Profile fields to return. Sample values: 'nickname', 'screen_name', 'sex', 'bdate' (birthdate), 'city', 'country', 'timezone', 'photo', 'photo_medium', 'photo_big', 'has_mobile', 'rate', 'contacts', 'education', 'online, counters'.
func (f Friends) GetByPhones(phones []string, fields []objects.UsersFields) (resp responses.FriendsGetByPhones, err error) {
	params := map[string]interface{}{}

	if len(phones) > 0 {
		params["phones"] = SliceToString(phones)
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	err = f.SendObjRequest("friends.getByPhones", params, &resp)

	return
}

// GetLists - Returns a list of the user's friend lists.
// Parameters:
//   * userId - User ID.
//   * returnSystem - '1' — to return system friend lists. By default: '0'.
func (f Friends) GetLists(userId int, returnSystem bool) (resp responses.FriendsGetLists, err error) {
	params := map[string]interface{}{}

	if userId > 0 {
		params["user_id"] = userId
	}

	params["return_system"] = returnSystem

	err = f.SendObjRequest("friends.getLists", params, &resp)

	return
}

// GetMutual - Returns a list of user IDs of the mutual friends of two users.
// Parameters:
//   * sourceUid - ID of the user whose friends will be checked against the friends of the user specified in 'target_uid'.
//   * targetUid - ID of the user whose friends will be checked against the friends of the user specified in 'source_uid'.
//   * targetUids - IDs of the users whose friends will be checked against the friends of the user specified in 'source_uid'.
//   * order - Sort order: 'random' — random order
//   * count - Number of mutual friends to return.
//   * offset - Offset needed to return a specific subset of mutual friends.
func (f Friends) GetMutual(sourceUid int, targetUid int, targetUids []int, order string, count int, offset int) (resp responses.FriendsGetMutual, err error) {
	params := map[string]interface{}{}

	if sourceUid > 0 {
		params["source_uid"] = sourceUid
	}

	if targetUid > 0 {
		params["target_uid"] = targetUid
	}

	if len(targetUids) > 0 {
		params["target_uids"] = SliceToString(targetUids)
	}

	if order != "" {
		params["order"] = order
	}

	if count > 0 {
		params["count"] = count
	}

	if offset > 0 {
		params["offset"] = offset
	}

	err = f.SendObjRequest("friends.getMutual", params, &resp)

	return
}

// GetOnline - Returns a list of user IDs of a user's friends who are online.
// Parameters:
//   * userId - User ID.
//   * listId - Friend list ID. If this parameter is not set, information about all online friends is returned.
//   * onlineMobile - '1' — to return an additional 'online_mobile' field, '0' — (default),
//   * order - Sort order: 'random' — random order
//   * count - Number of friends to return.
//   * offset - Offset needed to return a specific subset of friends.
func (f Friends) GetOnline(userId int, listId int, onlineMobile bool, order string, count int, offset int) (resp responses.FriendsGetOnline, err error) {
	params := map[string]interface{}{}

	if userId > 0 {
		params["user_id"] = userId
	}

	if listId > 0 {
		params["list_id"] = listId
	}

	params["online_mobile"] = onlineMobile

	if order != "" {
		params["order"] = order
	}

	if count > 0 {
		params["count"] = count
	}

	if offset > 0 {
		params["offset"] = offset
	}

	err = f.SendObjRequest("friends.getOnline", params, &resp)

	return
}

// GetRecent - Returns a list of user IDs of the current user's recently added friends.
// Parameters:
//   * count - Number of recently added friends to return.
func (f Friends) GetRecent(count int) (resp responses.FriendsGetRecent, err error) {
	params := map[string]interface{}{}

	if count > 0 {
		params["count"] = count
	}

	err = f.SendObjRequest("friends.getRecent", params, &resp)

	return
}

// GetRequests - Returns information about the current user's incoming and outgoing friend requests.
// Parameters:
//   * offset - Offset needed to return a specific subset of friend requests.
//   * count - Number of friend requests to return (default 100, maximum 1000).
//   * extended - '1' — to return response messages from users who have sent a friend request or, if 'suggested' is set to '1', to return a list of suggested friends
//   * needMutual - '1' — to return a list of mutual friends (up to 20), if any
//   * out - '1' — to return outgoing requests, '0' — to return incoming requests (default)
//   * sort - Sort order: '1' — by number of mutual friends, '0' — by date
//   * needViewed - NO DESCRIPTION IN JSON SCHEMA
//   * suggested - '1' — to return a list of suggested friends, '0' — to return friend requests (default)
//   * ref - NO DESCRIPTION IN JSON SCHEMA
//   * fields - NO DESCRIPTION IN JSON SCHEMA
func (f Friends) GetRequests(offset int, count int, needMutual bool, out bool, sort int, needViewed bool, suggested bool, ref string, fields []objects.UsersFields) (resp responses.FriendsGetRequests, err error) {
	params := map[string]interface{}{}
	params["extended"] = "0"

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	params["need_mutual"] = needMutual

	params["out"] = out

	if sort > 0 {
		params["sort"] = sort
	}

	params["need_viewed"] = needViewed

	params["suggested"] = suggested

	if ref != "" {
		params["ref"] = ref
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	err = f.SendObjRequest("friends.getRequests", params, &resp)

	return
}

// GetRequestsExtended - Returns information about the current user's incoming and outgoing friend requests.
// Parameters:
//   * offset - Offset needed to return a specific subset of friend requests.
//   * count - Number of friend requests to return (default 100, maximum 1000).
//   * extended - '1' — to return response messages from users who have sent a friend request or, if 'suggested' is set to '1', to return a list of suggested friends
//   * needMutual - '1' — to return a list of mutual friends (up to 20), if any
//   * out - '1' — to return outgoing requests, '0' — to return incoming requests (default)
//   * sort - Sort order: '1' — by number of mutual friends, '0' — by date
//   * needViewed - NO DESCRIPTION IN JSON SCHEMA
//   * suggested - '1' — to return a list of suggested friends, '0' — to return friend requests (default)
//   * ref - NO DESCRIPTION IN JSON SCHEMA
//   * fields - NO DESCRIPTION IN JSON SCHEMA
func (f Friends) GetRequestsExtended(offset int, count int, needMutual bool, out bool, sort int, needViewed bool, suggested bool, ref string, fields []objects.UsersFields) (resp responses.FriendsGetRequestsExtended, err error) {
	params := map[string]interface{}{}
	params["extended"] = "1"

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	params["need_mutual"] = needMutual

	params["out"] = out

	if sort > 0 {
		params["sort"] = sort
	}

	params["need_viewed"] = needViewed

	params["suggested"] = suggested

	if ref != "" {
		params["ref"] = ref
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	err = f.SendObjRequest("friends.getRequests", params, &resp)

	return
}

// GetSuggestions - Returns a list of profiles of users whom the current user may know.
// Parameters:
//   * filter - Types of potential friends to return: 'mutual' — users with many mutual friends , 'contacts' — users found with the [vk.com/dev/account.importContacts|account.importContacts] method , 'mutual_contacts' — users who imported the same contacts as the current user with the [vk.com/dev/account.importContacts|account.importContacts] method
//   * count - Number of suggestions to return.
//   * offset - Offset needed to return a specific subset of suggestions.
//   * fields - Profile fields to return. Sample values: 'nickname', 'screen_name', 'sex', 'bdate' (birthdate), 'city', 'country', 'timezone', 'photo', 'photo_medium', 'photo_big', 'has_mobile', 'rate', 'contacts', 'education', 'online', 'counters'.
//   * nameCase - Case for declension of user name and surname: , 'nom' — nominative (default) , 'gen' — genitive , 'dat' — dative , 'acc' — accusative , 'ins' — instrumental , 'abl' — prepositional
func (f Friends) GetSuggestions(filter []string, count int, offset int, fields []objects.UsersFields, nameCase string) (resp responses.FriendsGetSuggestions, err error) {
	params := map[string]interface{}{}

	if len(filter) > 0 {
		params["filter"] = SliceToString(filter)
	}

	if count > 0 {
		params["count"] = count
	}

	if offset > 0 {
		params["offset"] = offset
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	if nameCase != "" {
		params["name_case"] = nameCase
	}

	err = f.SendObjRequest("friends.getSuggestions", params, &resp)

	return
}

// Search - Returns a list of friends matching the search criteria.
// Parameters:
//   * userId - User ID.
//   * q - Search query string (e.g., 'Vasya Babich').
//   * fields - Profile fields to return. Sample values: 'nickname', 'screen_name', 'sex', 'bdate' (birthdate), 'city', 'country', 'timezone', 'photo', 'photo_medium', 'photo_big', 'has_mobile', 'rate', 'contacts', 'education', 'online',
//   * nameCase - Case for declension of user name and surname: 'nom' — nominative (default), 'gen' — genitive , 'dat' — dative, 'acc' — accusative , 'ins' — instrumental , 'abl' — prepositional
//   * offset - Offset needed to return a specific subset of friends.
//   * count - Number of friends to return.
func (f Friends) Search(userId int, q string, fields []objects.UsersFields, nameCase string, offset int, count int) (resp responses.FriendsSearch, err error) {
	params := map[string]interface{}{}

	params["user_id"] = userId

	if q != "" {
		params["q"] = q
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	if nameCase != "" {
		params["name_case"] = nameCase
	}

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	err = f.SendObjRequest("friends.search", params, &resp)

	return
}
