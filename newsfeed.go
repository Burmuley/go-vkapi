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
	"encoding/json"
	"gitlab.com/Burmuley/go-vkapi/objects"
	"gitlab.com/Burmuley/go-vkapi/responses"
)

type Newsfeed struct {
	*VKApi
}

/////////////////////////////////////////////////////////////
// `Newsfeed` methods
/////////////////////////////////////////////////////////////

// Addban - Prevents news from specified users and communities from appearing in the current user's newsfeed.
// Parameters:
//   * userIds - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * groupIds - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (n *Newsfeed) Addban(userIds []int, groupIds []int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	if len(userIds) > 0 {
		params["user_ids"] = SliceToString(userIds)
	}

	if len(groupIds) > 0 {
		params["group_ids"] = SliceToString(groupIds)
	}

	err = n.SendObjRequest("newsfeed.addBan", params, &resp)

	return
}

// Deleteban - Allows news from previously banned users and communities to be shown in the current user's newsfeed.
// Parameters:
//   * userIds - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * groupIds - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (n *Newsfeed) Deleteban(userIds []int, groupIds []int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	if len(userIds) > 0 {
		params["user_ids"] = SliceToString(userIds)
	}

	if len(groupIds) > 0 {
		params["group_ids"] = SliceToString(groupIds)
	}

	err = n.SendObjRequest("newsfeed.deleteBan", params, &resp)

	return
}

// Deletelist - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * listId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (n *Newsfeed) Deletelist(listId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["list_id"] = listId

	err = n.SendObjRequest("newsfeed.deleteList", params, &resp)

	return
}

// Get - Returns data required to show newsfeed for the current user.
// Parameters:
//   * filters - Filters to apply: 'post' — new wall posts, 'photo' — new photos, 'photo_tag' — new photo tags, 'wall_photo' — new wall photos, 'friend' — new friends, 'note' — new notes
//   * returnBanned - '1' — to return news items from banned sources
//   * startTime - Earliest timestamp (in Unix time) of a news item to return. By default, 24 hours ago.
//   * endTime - Latest timestamp (in Unix time) of a news item to return. By default, the current time.
//   * maxPhotos - Maximum number of photos to return. By default, '5'.
//   * sourceIds - Sources to obtain news from, separated by commas. User IDs can be specified in formats '' or 'u' , where '' is the user's friend ID. Community IDs can be specified in formats '-' or 'g' , where '' is the community ID. If the parameter is not set, all of the user's friends and communities are returned, except for banned sources, which can be obtained with the [vk.com/dev/newsfeed.getBanned|newsfeed.getBanned] method.
//   * startFrom - identifier required to get the next page of results. Value for this parameter is returned in 'next_from' field in a reply.
//   * count - Number of news items to return (default 50, maximum 100). For auto feed, you can use the 'new_offset' parameter returned by this method.
//   * fields - Additional fields of [vk.com/dev/fields|profiles] and [vk.com/dev/fields_groups|communities] to return.
//   * section - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (n *Newsfeed) Get(filters []objects.NewsfeedFilters, returnBanned bool, startTime int, endTime int, maxPhotos int, sourceIds string, startFrom string, count int, fields []objects.BaseUserGroupFields, section string) (resp responses.NewsfeedGet, err error) {
	params := map[string]interface{}{}

	if len(filters) > 0 {
		params["filters"] = SliceToString(filters)
	}

	params["return_banned"] = returnBanned

	if startTime > 0 {
		params["start_time"] = startTime
	}

	if endTime > 0 {
		params["end_time"] = endTime
	}

	if maxPhotos > 0 {
		params["max_photos"] = maxPhotos
	}

	if sourceIds != "" {
		params["source_ids"] = sourceIds
	}

	if startFrom != "" {
		params["start_from"] = startFrom
	}

	if count > 0 {
		params["count"] = count
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	if section != "" {
		params["section"] = section
	}

	err = n.SendObjRequest("newsfeed.get", params, &resp)

	return
}

// Getbanned - Returns a list of users and communities banned from the current user's newsfeed.
// Parameters:
//   * extended - '1' — return extra information about users and communities
//   * fields - Profile fields to return.
//   * nameCase - Case for declension of user name and surname: 'nom' — nominative (default), 'gen' — genitive , 'dat' — dative, 'acc' — accusative , 'ins' — instrumental , 'abl' — prepositional
func (n *Newsfeed) Getbanned(fields []objects.UsersFields, nameCase string) (resp responses.NewsfeedGetbanned, err error) {
	params := map[string]interface{}{}
	params["extended"] = "0"

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	if nameCase != "" {
		params["name_case"] = nameCase
	}

	err = n.SendObjRequest("newsfeed.getBanned", params, &resp)

	return
}

// GetbannedExtended - Returns a list of users and communities banned from the current user's newsfeed.
// Parameters:
//   * extended - '1' — return extra information about users and communities
//   * fields - Profile fields to return.
//   * nameCase - Case for declension of user name and surname: 'nom' — nominative (default), 'gen' — genitive , 'dat' — dative, 'acc' — accusative , 'ins' — instrumental , 'abl' — prepositional
func (n *Newsfeed) GetbannedExtended(fields []objects.UsersFields, nameCase string) (resp responses.NewsfeedGetbannedExtended, err error) {
	params := map[string]interface{}{}
	params["extended"] = "1"

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	if nameCase != "" {
		params["name_case"] = nameCase
	}

	err = n.SendObjRequest("newsfeed.getBanned", params, &resp)

	return
}

// Getcomments - Returns a list of comments in the current user's newsfeed.
// Parameters:
//   * count - Number of comments to return. For auto feed, you can use the 'new_offset' parameter returned by this method.
//   * filters - Filters to apply: 'post' — new comments on wall posts, 'photo' — new comments on photos, 'video' — new comments on videos, 'topic' — new comments on discussions, 'note' — new comments on notes,
//   * reposts - Object ID, comments on repost of which shall be returned, e.g. 'wall1_45486'. (If the parameter is set, the 'filters' parameter is optional.),
//   * startTime - Earliest timestamp (in Unix time) of a comment to return. By default, 24 hours ago.
//   * endTime - Latest timestamp (in Unix time) of a comment to return. By default, the current time.
//   * lastCommentsCount - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * startFrom - Identificator needed to return the next page with results. Value for this parameter returns in 'next_from' field.
//   * fields - Additional fields of [vk.com/dev/fields|profiles] and [vk.com/dev/fields_groups|communities] to return.
func (n *Newsfeed) Getcomments(count int, filters []objects.NewsfeedCommentsFilters, reposts string, startTime int, endTime int, lastCommentsCount int, startFrom string, fields []objects.BaseUserGroupFields) (resp responses.NewsfeedGetcomments, err error) {
	params := map[string]interface{}{}

	if count > 0 {
		params["count"] = count
	}

	if len(filters) > 0 {
		params["filters"] = SliceToString(filters)
	}

	if reposts != "" {
		params["reposts"] = reposts
	}

	if startTime > 0 {
		params["start_time"] = startTime
	}

	if endTime > 0 {
		params["end_time"] = endTime
	}

	if lastCommentsCount > 0 {
		params["last_comments_count"] = lastCommentsCount
	}

	if startFrom != "" {
		params["start_from"] = startFrom
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	err = n.SendObjRequest("newsfeed.getComments", params, &resp)

	return
}

// Getlists - Returns a list of newsfeeds followed by the current user.
// Parameters:
//   * listIds - numeric list identifiers.
//   * extended - Return additional list info
func (n *Newsfeed) Getlists(listIds []int) (resp responses.NewsfeedGetlists, err error) {
	params := map[string]interface{}{}
	params["extended"] = "0"

	if len(listIds) > 0 {
		params["list_ids"] = SliceToString(listIds)
	}

	err = n.SendObjRequest("newsfeed.getLists", params, &resp)

	return
}

// GetlistsExtended - Returns a list of newsfeeds followed by the current user.
// Parameters:
//   * listIds - numeric list identifiers.
//   * extended - Return additional list info
func (n *Newsfeed) GetlistsExtended(listIds []int) (resp responses.NewsfeedGetlistsExtended, err error) {
	params := map[string]interface{}{}
	params["extended"] = "1"

	if len(listIds) > 0 {
		params["list_ids"] = SliceToString(listIds)
	}

	err = n.SendObjRequest("newsfeed.getLists", params, &resp)

	return
}

// Getmentions - Returns a list of posts on user walls in which the current user is mentioned.
// Parameters:
//   * ownerId - Owner ID.
//   * startTime - Earliest timestamp (in Unix time) of a post to return. By default, 24 hours ago.
//   * endTime - Latest timestamp (in Unix time) of a post to return. By default, the current time.
//   * offset - Offset needed to return a specific subset of posts.
//   * count - Number of posts to return.
func (n *Newsfeed) Getmentions(ownerId int, startTime int, endTime int, offset int, count int) (resp responses.NewsfeedGetmentions, err error) {
	params := map[string]interface{}{}

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	if startTime > 0 {
		params["start_time"] = startTime
	}

	if endTime > 0 {
		params["end_time"] = endTime
	}

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	err = n.SendObjRequest("newsfeed.getMentions", params, &resp)

	return
}

// Getrecommended - , Returns a list of newsfeeds recommended to the current user.
// Parameters:
//   * startTime - Earliest timestamp (in Unix time) of a news item to return. By default, 24 hours ago.
//   * endTime - Latest timestamp (in Unix time) of a news item to return. By default, the current time.
//   * maxPhotos - Maximum number of photos to return. By default, '5'.
//   * startFrom - 'new_from' value obtained in previous call.
//   * count - Number of news items to return.
//   * fields - Additional fields of [vk.com/dev/fields|profiles] and [vk.com/dev/fields_groups|communities] to return.
func (n *Newsfeed) Getrecommended(startTime int, endTime int, maxPhotos int, startFrom string, count int, fields []objects.BaseUserGroupFields) (resp responses.NewsfeedGetrecommended, err error) {
	params := map[string]interface{}{}

	if startTime > 0 {
		params["start_time"] = startTime
	}

	if endTime > 0 {
		params["end_time"] = endTime
	}

	if maxPhotos > 0 {
		params["max_photos"] = maxPhotos
	}

	if startFrom != "" {
		params["start_from"] = startFrom
	}

	if count > 0 {
		params["count"] = count
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	err = n.SendObjRequest("newsfeed.getRecommended", params, &resp)

	return
}

// Getsuggestedsources - Returns communities and users that current user is suggested to follow.
// Parameters:
//   * offset - offset required to choose a particular subset of communities or users.
//   * count - amount of communities or users to return.
//   * shuffle - shuffle the returned list or not.
//   * fields - list of extra fields to be returned. See available fields for [vk.com/dev/fields|users] and [vk.com/dev/fields_groups|communities].
func (n *Newsfeed) Getsuggestedsources(offset int, count int, shuffle bool, fields []objects.BaseUserGroupFields) (resp responses.NewsfeedGetsuggestedsources, err error) {
	params := map[string]interface{}{}

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	params["shuffle"] = shuffle

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	err = n.SendObjRequest("newsfeed.getSuggestedSources", params, &resp)

	return
}

// Ignoreitem - Hides an item from the newsfeed.
// Parameters:
//   * pType - Item type. Possible values: *'wall' – post on the wall,, *'tag' – tag on a photo,, *'profilephoto' – profile photo,, *'video' – video,, *'audio' – audio.
//   * ownerId - Item owner's identifier (user or community), "Note that community id must be negative. 'owner_id=1' – user , 'owner_id=-1' – community "
//   * itemId - Item identifier
func (n *Newsfeed) Ignoreitem(pType objects.NewsfeedIgnoreItemType, ownerId int, itemId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["type"] = pType

	params["owner_id"] = ownerId

	params["item_id"] = itemId

	err = n.SendObjRequest("newsfeed.ignoreItem", params, &resp)

	return
}

// Savelist - Creates and edits user newsfeed lists
// Parameters:
//   * listId - numeric list identifier (if not sent, will be set automatically).
//   * title - list name.
//   * sourceIds - users and communities identifiers to be added to the list. Community identifiers must be negative numbers.
//   * noReposts - reposts display on and off ('1' is for off).
func (n *Newsfeed) Savelist(listId int, title string, sourceIds []int, noReposts bool) (resp responses.NewsfeedSavelist, err error) {
	params := map[string]interface{}{}

	if listId > 0 {
		params["list_id"] = listId
	}

	params["title"] = title

	if len(sourceIds) > 0 {
		params["source_ids"] = SliceToString(sourceIds)
	}

	params["no_reposts"] = noReposts

	err = n.SendObjRequest("newsfeed.saveList", params, &resp)

	return
}

// Search - Returns search results by statuses.
// Parameters:
//   * q - Search query string (e.g., 'New Year').
//   * extended - '1' — to return additional information about the user or community that placed the post.
//   * count - Number of posts to return.
//   * latitude - Geographical latitude point (in degrees, -90 to 90) within which to search.
//   * longitude - Geographical longitude point (in degrees, -180 to 180) within which to search.
//   * startTime - Earliest timestamp (in Unix time) of a news item to return. By default, 24 hours ago.
//   * endTime - Latest timestamp (in Unix time) of a news item to return. By default, the current time.
//   * startFrom - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * fields - Additional fields of [vk.com/dev/fields|profiles] and [vk.com/dev/fields_groups|communities] to return.
func (n *Newsfeed) Search(q string, count int, latitude json.Number, longitude json.Number, startTime int, endTime int, startFrom string, fields []objects.BaseUserGroupFields) (resp responses.NewsfeedSearch, err error) {
	params := map[string]interface{}{}
	params["extended"] = "0"

	if q != "" {
		params["q"] = q
	}

	if count > 0 {
		params["count"] = count
	}

	if v, err := latitude.Int64(); err == nil && v > 0 {
		params["latitude"] = v
	} else if v := latitude.String(); v != "" {
		params["latitude"] = v
	}

	if v, err := longitude.Int64(); err == nil && v > 0 {
		params["longitude"] = v
	} else if v := longitude.String(); v != "" {
		params["longitude"] = v
	}

	if startTime > 0 {
		params["start_time"] = startTime
	}

	if endTime > 0 {
		params["end_time"] = endTime
	}

	if startFrom != "" {
		params["start_from"] = startFrom
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	err = n.SendObjRequest("newsfeed.search", params, &resp)

	return
}

// SearchExtended - Returns search results by statuses.
// Parameters:
//   * q - Search query string (e.g., 'New Year').
//   * extended - '1' — to return additional information about the user or community that placed the post.
//   * count - Number of posts to return.
//   * latitude - Geographical latitude point (in degrees, -90 to 90) within which to search.
//   * longitude - Geographical longitude point (in degrees, -180 to 180) within which to search.
//   * startTime - Earliest timestamp (in Unix time) of a news item to return. By default, 24 hours ago.
//   * endTime - Latest timestamp (in Unix time) of a news item to return. By default, the current time.
//   * startFrom - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * fields - Additional fields of [vk.com/dev/fields|profiles] and [vk.com/dev/fields_groups|communities] to return.
func (n *Newsfeed) SearchExtended(q string, count int, latitude json.Number, longitude json.Number, startTime int, endTime int, startFrom string, fields []objects.BaseUserGroupFields) (resp responses.NewsfeedSearchExtended, err error) {
	params := map[string]interface{}{}
	params["extended"] = "1"

	if q != "" {
		params["q"] = q
	}

	if count > 0 {
		params["count"] = count
	}

	if v, err := latitude.Int64(); err == nil && v > 0 {
		params["latitude"] = v
	} else if v := latitude.String(); v != "" {
		params["latitude"] = v
	}

	if v, err := longitude.Int64(); err == nil && v > 0 {
		params["longitude"] = v
	} else if v := longitude.String(); v != "" {
		params["longitude"] = v
	}

	if startTime > 0 {
		params["start_time"] = startTime
	}

	if endTime > 0 {
		params["end_time"] = endTime
	}

	if startFrom != "" {
		params["start_from"] = startFrom
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	err = n.SendObjRequest("newsfeed.search", params, &resp)

	return
}

// Unignoreitem - Returns a hidden item to the newsfeed.
// Parameters:
//   * pType - Item type. Possible values: *'wall' – post on the wall,, *'tag' – tag on a photo,, *'profilephoto' – profile photo,, *'video' – video,, *'audio' – audio.
//   * ownerId - Item owner's identifier (user or community), "Note that community id must be negative. 'owner_id=1' – user , 'owner_id=-1' – community "
//   * itemId - Item identifier
func (n *Newsfeed) Unignoreitem(pType objects.NewsfeedIgnoreItemType, ownerId int, itemId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["type"] = pType

	params["owner_id"] = ownerId

	params["item_id"] = itemId

	err = n.SendObjRequest("newsfeed.unignoreItem", params, &resp)

	return
}

// Unsubscribe - Unsubscribes the current user from specified newsfeeds.
// Parameters:
//   * pType - Type of object from which to unsubscribe: 'note' — note, 'photo' — photo, 'post' — post on user wall or community wall, 'topic' — topic, 'video' — video
//   * ownerId - Object owner ID.
//   * itemId - Object ID.
func (n *Newsfeed) Unsubscribe(pType string, ownerId int, itemId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["type"] = pType

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	params["item_id"] = itemId

	err = n.SendObjRequest("newsfeed.unsubscribe", params, &resp)

	return
}
