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

type Likes struct {
	*VKApi
}

/////////////////////////////////////////////////////////////
// `Likes` methods
/////////////////////////////////////////////////////////////

// Add - Adds the specified object to the 'Likes' list of the current user.
// Parameters:
//   * pType - Object type: 'post' — post on user or community wall, 'comment' — comment on a wall post, 'photo' — photo, 'audio' — audio, 'video' — video, 'note' — note, 'photo_comment' — comment on the photo, 'video_comment' — comment on the video, 'topic_comment' — comment in the discussion, 'sitepage' — page of the site where the [vk.com/dev/Like|Like widget] is installed
//   * ownerId - ID of the user or community that owns the object.
//   * itemId - Object ID.
//   * accessKey - Access key required for an object owned by a private entity.
func (l *Likes) Add(pType objects.LikesType, ownerId int, itemId int, accessKey string) (resp responses.LikesAdd, err error) {
	params := map[string]interface{}{}

	params["type"] = pType

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	params["item_id"] = itemId

	if accessKey != "" {
		params["access_key"] = accessKey
	}

	err = l.SendObjRequest("likes.add", params, &resp)

	return
}

// Delete - Deletes the specified object from the 'Likes' list of the current user.
// Parameters:
//   * pType - Object type: 'post' — post on user or community wall, 'comment' — comment on a wall post, 'photo' — photo, 'audio' — audio, 'video' — video, 'note' — note, 'photo_comment' — comment on the photo, 'video_comment' — comment on the video, 'topic_comment' — comment in the discussion, 'sitepage' — page of the site where the [vk.com/dev/Like|Like widget] is installed
//   * ownerId - ID of the user or community that owns the object.
//   * itemId - Object ID.
func (l *Likes) Delete(pType objects.LikesType, ownerId int, itemId int) (resp responses.LikesDelete, err error) {
	params := map[string]interface{}{}

	params["type"] = pType

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	params["item_id"] = itemId

	err = l.SendObjRequest("likes.delete", params, &resp)

	return
}

// Getlist - Returns a list of IDs of users who added the specified object to their 'Likes' list.
// Parameters:
//   * pType - , Object type: 'post' — post on user or community wall, 'comment' — comment on a wall post, 'photo' — photo, 'audio' — audio, 'video' — video, 'note' — note, 'photo_comment' — comment on the photo, 'video_comment' — comment on the video, 'topic_comment' — comment in the discussion, 'sitepage' — page of the site where the [vk.com/dev/Like|Like widget] is installed
//   * ownerId - ID of the user, community, or application that owns the object. If the 'type' parameter is set as 'sitepage', the application ID is passed as 'owner_id'. Use negative value for a community id. If the 'type' parameter is not set, the 'owner_id' is assumed to be either the current user or the same application ID as if the 'type' parameter was set to 'sitepage'.
//   * itemId - Object ID. If 'type' is set as 'sitepage', 'item_id' can include the 'page_id' parameter value used during initialization of the [vk.com/dev/Like|Like widget].
//   * pageUrl - URL of the page where the [vk.com/dev/Like|Like widget] is installed. Used instead of the 'item_id' parameter.
//   * filter - Filters to apply: 'likes' — returns information about all users who liked the object (default), 'copies' — returns information only about users who told their friends about the object
//   * friendsOnly - Specifies which users are returned: '1' — to return only the current user's friends, '0' — to return all users (default)
//   * extended - Specifies whether extended information will be returned. '1' — to return extended information about users and communities from the 'Likes' list, '0' — to return no additional information (default)
//   * offset - Offset needed to select a specific subset of users.
//   * count - Number of user IDs to return (maximum '1000'). Default is '100' if 'friends_only' is set to '0', otherwise, the default is '10' if 'friends_only' is set to '1'.
//   * skipOwn - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (l *Likes) Getlist(pType objects.LikesType, ownerId int, itemId int, pageUrl string, filter string, friendsOnly int, offset int, count int, skipOwn bool) (resp responses.LikesGetlist, err error) {
	params := map[string]interface{}{}
	params["extended"] = "0"

	params["type"] = pType

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	if itemId > 0 {
		params["item_id"] = itemId
	}

	if pageUrl != "" {
		params["page_url"] = pageUrl
	}

	if filter != "" {
		params["filter"] = filter
	}

	if friendsOnly > 0 {
		params["friends_only"] = friendsOnly
	}

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	params["skip_own"] = skipOwn

	err = l.SendObjRequest("likes.getList", params, &resp)

	return
}

// GetlistExtended - Returns a list of IDs of users who added the specified object to their 'Likes' list.
// Parameters:
//   * pType - , Object type: 'post' — post on user or community wall, 'comment' — comment on a wall post, 'photo' — photo, 'audio' — audio, 'video' — video, 'note' — note, 'photo_comment' — comment on the photo, 'video_comment' — comment on the video, 'topic_comment' — comment in the discussion, 'sitepage' — page of the site where the [vk.com/dev/Like|Like widget] is installed
//   * ownerId - ID of the user, community, or application that owns the object. If the 'type' parameter is set as 'sitepage', the application ID is passed as 'owner_id'. Use negative value for a community id. If the 'type' parameter is not set, the 'owner_id' is assumed to be either the current user or the same application ID as if the 'type' parameter was set to 'sitepage'.
//   * itemId - Object ID. If 'type' is set as 'sitepage', 'item_id' can include the 'page_id' parameter value used during initialization of the [vk.com/dev/Like|Like widget].
//   * pageUrl - URL of the page where the [vk.com/dev/Like|Like widget] is installed. Used instead of the 'item_id' parameter.
//   * filter - Filters to apply: 'likes' — returns information about all users who liked the object (default), 'copies' — returns information only about users who told their friends about the object
//   * friendsOnly - Specifies which users are returned: '1' — to return only the current user's friends, '0' — to return all users (default)
//   * extended - Specifies whether extended information will be returned. '1' — to return extended information about users and communities from the 'Likes' list, '0' — to return no additional information (default)
//   * offset - Offset needed to select a specific subset of users.
//   * count - Number of user IDs to return (maximum '1000'). Default is '100' if 'friends_only' is set to '0', otherwise, the default is '10' if 'friends_only' is set to '1'.
//   * skipOwn - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (l *Likes) GetlistExtended(pType objects.LikesType, ownerId int, itemId int, pageUrl string, filter string, friendsOnly int, offset int, count int, skipOwn bool) (resp responses.LikesGetlistExtended, err error) {
	params := map[string]interface{}{}
	params["extended"] = "1"

	params["type"] = pType

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	if itemId > 0 {
		params["item_id"] = itemId
	}

	if pageUrl != "" {
		params["page_url"] = pageUrl
	}

	if filter != "" {
		params["filter"] = filter
	}

	if friendsOnly > 0 {
		params["friends_only"] = friendsOnly
	}

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	params["skip_own"] = skipOwn

	err = l.SendObjRequest("likes.getList", params, &resp)

	return
}

// Isliked - Checks for the object in the 'Likes' list of the specified user.
// Parameters:
//   * userId - User ID.
//   * pType - Object type: 'post' — post on user or community wall, 'comment' — comment on a wall post, 'photo' — photo, 'audio' — audio, 'video' — video, 'note' — note, 'photo_comment' — comment on the photo, 'video_comment' — comment on the video, 'topic_comment' — comment in the discussion
//   * ownerId - ID of the user or community that owns the object.
//   * itemId - Object ID.
func (l *Likes) Isliked(userId int, pType objects.LikesType, ownerId int, itemId int) (resp responses.LikesIsliked, err error) {
	params := map[string]interface{}{}

	if userId > 0 {
		params["user_id"] = userId
	}

	params["type"] = pType

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	params["item_id"] = itemId

	err = l.SendObjRequest("likes.isLiked", params, &resp)

	return
}
