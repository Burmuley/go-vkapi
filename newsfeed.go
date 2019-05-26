package go_vkapi

import (
	"fmt"
	"gitlab.com/Burmuley/go-vkapi/objects"
	"gitlab.com/Burmuley/go-vkapi/responses"
)

// Newsfeed represents collection of methods related to VK Accounts
type Newsfeed struct {
	*VKApi
}

// AddBan prevents news from specified users and communities from appearing in the current user's newsfeed
//
// Parameters:
//  * userIds - int array with user IDs
//  * groupIds - int array with group IDs
func (n *Newsfeed) AddBan(userIds []int, groupIds []int) (responses.OkResponse, error) {
	params := map[string]string{}

	if len(userIds) > 0 {
		params["user_ids"] = SliceToString(userIds)
	}

	if len(groupIds) > 0 {
		params["group_ids"] = SliceToString(groupIds)
	}

	var resp responses.OkResponse

	if err := n.SendObjRequest("newsfeed.addBan", params, &resp); err != nil {
		return responses.OkResponse(0), err
	}

	return resp, nil
}

// DeleteBan Allows news from previously banned users and communities to be shown in the current user's newsfeed
//
// Parameters:
//  * userIds - int array with user IDs
//  * groupIds - int array with group IDs
func (n *Newsfeed) DeleteBan(userIds []int, groupIds []int) (responses.OkResponse, error) {
	params := map[string]string{}

	if len(userIds) > 0 {
		params["user_ids"] = SliceToString(userIds)
	}

	if len(groupIds) > 0 {
		params["group_ids"] = SliceToString(groupIds)
	}

	var resp responses.OkResponse

	if err := n.SendObjRequest("newsfeed.deleteBan", params, &resp); err != nil {
		return responses.OkResponse(0), err
	}

	return resp, nil
}

// DeleteList deletes list
//
// Parameters:
//  * listId - int ID of the list to delete
func (n *Newsfeed) DeleteList(listId int) (responses.OkResponse, error) {
	params := map[string]string{"list_id": string(listId)}

	var resp responses.OkResponse

	if err := n.SendObjRequest("newsfeed.deleteList", params, &resp); err != nil {
		return responses.OkResponse(0), err
	}

	return resp, nil
}

// Get returns data required to show newsfeed for the current user
//
// Parameters:
//  * filters - Filters to apply: 'post' — new wall posts, 'photo' — new photos, 'photo_tag' — new photo tags, 'wall_photo' — new wall photos, 'friend' — new friends, 'note' — new notes
//  * returnBanned - 'true' — to return news items from banned sources
//  * startTime - earliest timestamp (in Unix time) of a news item to return. By default, 24 hours ago
//  * endTime - latest timestamp (in Unix time) of a news item to return. By default, the current time
//  * maxPhotos - maximum number of photos to return. By default, '5'
//  * sourceIds - sources to obtain news from, separated by commas.
//      User IDs can be specified in formats '' or 'u' , where '' is the user's friend ID.
//      Community IDs can be specified in formats '-' or 'g' , where '' is the community ID.
//      If the parameter is not set, all of the user's friends and communities are returned, except for banned sources, which can be obtained with the vk.com/dev/newsfeed.getBanned method
//  * startFrom - identifier required to get the next page of results. Value for this parameter is returned in 'next_from' field in a reply.
//  * count - Number of news items to return (default 50, maximum 100). For auto feed, you can use the 'new_offset' parameter returned by this method
//  * fields - additional fields of vk.com/dev/fields|profiles and vk.com/dev/fields_groups|communities to return
//  * section - NO DESCRIPTION
func (n *Newsfeed) Get(count, startTime, endTime, maxPhotos int, returnBanned bool,
	filters []objects.NewsfeedFilters, section, startFrom string, fields []objects.BaseUserGroupFields) (responses.NewsfeedGetResponse, error) {
	params := map[string]string{}

	if count > 0 {
		params["count"] = string(count)
	}

	if startTime > 0 {
		params["start_time"] = string(startTime)
	}

	if endTime > 0 {
		params["end_time"] = string(endTime)
	}

	if maxPhotos > 0 {
		params["max_photos"] = string(maxPhotos)
	}

	if len(startFrom) > 0 {
		params["start_from"] = startFrom
	}

	if len(section) > 0 {
		params["section"] = section
	}

	if len(filters) > 0 {
		params["filters"] = SliceToString(filters)
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	resp := responses.NewsfeedGetResponse{}

	if err := n.SendObjRequest("newsfeed.get", params, &resp); err != nil {
		return responses.NewsfeedGetResponse{}, err
	}

	return resp, nil
}

// GetBanned returns a list of users and communities banned from the current user's newsfeed
//
// Parameters:
//  * fields - profile fields to return
//  * nameCase - Case for declension of user name and surname:
//        'nom' — nominative (default)
//        'gen' — genitive
//        'dat' — dative
//        'acc' — accusative
//        'ins' — instrumental
//        'abl' — prepositional
func (n *Newsfeed) GetBanned(fields []objects.UsersFields, nameCase string) (responses.NewsfeedGetBanned, error) {
	params := map[string]string{"extended": "0"}

	if len(nameCase) > 0 {
		params["name_case"] = nameCase
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	resp := responses.NewsfeedGetBanned{}

	if err := n.SendObjRequest("newsfeed.getBanned", params, &resp); err != nil {
		return responses.NewsfeedGetBanned{}, err
	}

	return resp, nil
}

// GetBannedExt returns a list of users and communities banned from the current user's newsfeed
//
// Parameters:
//  * fields - profile fields to return
//  * nameCase - Case for declension of user name and surname:
//        'nom' — nominative (default)
//        'gen' — genitive
//        'dat' — dative
//        'acc' — accusative
//        'ins' — instrumental
//        'abl' — prepositional
func (n *Newsfeed) GetBannedExt(fields []objects.UsersFields, nameCase string) (responses.NewsfeedGetBannedExt, error) {
	params := map[string]string{"extended": "1"}

	if len(nameCase) > 0 {
		params["name_case"] = nameCase
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	resp := responses.NewsfeedGetBannedExt{}

	if err := n.SendObjRequest("newsfeed.getBanned", params, &resp); err != nil {
		return responses.NewsfeedGetBannedExt{}, err
	}

	return resp, nil
}

// GetComments returns a list of comments in the current user's newsfeed
//
// Parameters:
//  * count - number of comments to return. For auto feed, you can use the 'new_offset' parameter returned by this method
//  * filters - filters to apply:
//       'post' — new comments on wall posts
//       'photo' — new comments on photos
//       'video' — new comments on videos
//       'topic' — new comments on discussions
//       'note' — new comments on notes
//  * reposts - object ID, comments on repost of which shall be returned, e.g. 'wall1_45486'
//      If the parameter is set, the 'filters' parameter is optional.
//  * startTime - earliest timestamp (in Unix time) of a news item to return. By default, 24 hours ago
//  * endTime - latest timestamp (in Unix time) of a news item to return. By default, the current time
//  * lastCommentsCount - NO DESCRIPTION
//  * startFrom -  identificator needed to return the next page with results. Value for this parameter returns in 'next_from' field
//  * fields - additional fields of vk.com/dev/fields|profiles and vk.com/dev/fields_groups|communities to return
func (n *Newsfeed) GetComments(count, startTime, endTime, lastCommentCount int,
	filters []objects.NewsfeedCommentsFilters, reposts, startFrom string, fields []objects.BaseUserGroupFields) (responses.NewsfeedGetComments, error) {
	params := map[string]string{}

	if count > 0 {
		params["count"] = string(count)
	}

	if startTime > 0 {
		params["start_time"] = string(startTime)
	}

	if endTime > 0 {
		params["end_time"] = string(endTime)
	}

	if lastCommentCount > 0 {
		params["last_comments_count"] = string(lastCommentCount)
	}

	if len(startFrom) > 0 {
		params["start_from"] = startFrom
	}

	if len(reposts) > 0 {
		params["reposts"] = reposts
	}

	if len(filters) > 0 {
		params["filters"] = SliceToString(filters)
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	resp := responses.NewsfeedGetComments{}

	if err := n.SendObjRequest("newsfeed.getComments", params, &resp); err != nil {
		return responses.NewsfeedGetComments{}, err
	}

	return resp, nil
}

// GetLists returns a list of newsfeeds followed by the current user
//
//Parameters:
//  * listIds - numeric list identifiers.
func (n *Newsfeed) GetLists(listIds []int) (responses.NewsfeedGetLists, error) {
	params := map[string]string{"extended": "false"}

	if len(listIds) > 0 {
		params["list_ids"] = SliceToString(listIds)
	}

	resp := responses.NewsfeedGetLists{}

	if err := n.SendObjRequest("newsfeed.getLists", params, &resp); err != nil {
		return responses.NewsfeedGetLists{}, err
	}

	return resp, nil
}

// GetListsExt returns an extended list of newsfeeds followed by the current user
//
// Parameters:
//  * listIds - numeric list identifiers.
func (n *Newsfeed) GetListsExt(listIds []int) (responses.NewsfeedGetListsExt, error) {
	params := map[string]string{"extended": "true"}

	if len(listIds) > 0 {
		params["list_ids"] = SliceToString(listIds)
	}

	resp := responses.NewsfeedGetListsExt{}

	if err := n.SendObjRequest("newsfeed.getLists", params, &resp); err != nil {
		return responses.NewsfeedGetListsExt{}, err
	}

	return resp, nil
}

// GetMentions returns a list of posts on user walls in which the current user is mentioned.
//
// Parameters:
//  * ownerId - Owner ID.
//  * startTime - Earliest timestamp (in Unix time) of a post to return. By default, 24 hours ago.
//  * endTime - Latest timestamp (in Unix time) of a post to return. By default, the current time.
//  * offset - Offset needed to return a specific subset of posts.
//  * count - Number of posts to return.
func (n *Newsfeed) GetMentions(ownerId, startTime, endTime, offset, count int) (responses.NewsfeedGetMentions, error) {
	params := map[string]string{}

	if count > 0 {
		params["count"] = string(count)
	}

	if startTime > 0 {
		params["start_time"] = string(startTime)
	}

	if endTime > 0 {
		params["end_time"] = string(endTime)
	}

	if offset > 0 {
		params["offset"] = string(offset)
	}

	if ownerId > 0 {
		params["owner_id"] = string(ownerId)
	}

	resp := responses.NewsfeedGetMentions{}

	if err := n.SendObjRequest("newsfeed.getMentions", params, &resp); err != nil {
		return responses.NewsfeedGetMentions{}, err
	}

	return resp, nil
}

// GetRecommended returns a list of newsfeeds recommended to the current user
// Parameters:
//  * startTime - Earliest timestamp (in Unix time) of a news item to return. By default, 24 hours ago.
//  * endTime - Latest timestamp (in Unix time) of a news item to return. By default, the current time.
//  * maxPhotos - Maximum number of photos to return. By default, '5'.
//  * startFrom - 'new_from' value obtained in previous call.
//  * count - Number of news items to return.
//  * fields - Additional fields of [vk.com/dev/fields|profiles] and [vk.com/dev/fields_groups|communities] to return.
func (n *Newsfeed) GetRecommended(startTime, endTime, maxPhotos, count int,
	startFrom string, fields []objects.BaseUserGroupFields) (responses.NewsfeedGetRecommended, error) {
	params := map[string]string{}

	if len(startFrom) > 0 {
		params["start_from"] = startFrom
	}

	if count > 0 {
		params["count"] = string(count)
	}

	if startTime > 0 {
		params["start_time"] = string(startTime)
	}

	if endTime > 0 {
		params["end_time"] = string(endTime)
	}

	if maxPhotos > 0 {
		params["max_photos"] = string(maxPhotos)
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	resp := responses.NewsfeedGetRecommended{}

	if err := n.SendObjRequest("newsfeed.getRecommended", params, &resp); err != nil {
		return responses.NewsfeedGetRecommended{}, err
	}

	return resp, nil
}

// GetSuggestedSources returns communities and users that current user is suggested to follow
// Parameters:
//  * offset - offset required to choose a particular subset of communities or users.
//  * count - amount of communities or users to return.
//  * shuffle - shuffle the returned list or not.
//  * fields - list of extra fields to be returned. See available fields for [vk.com/dev/fields|users] and [vk.com/dev/fields_groups|communities].
func (n *Newsfeed) GetSuggestedSources(offset, count int, shuffle bool,
	fields []objects.BaseUserGroupFields) (responses.NewsfeedGetSuggestedSources, error) {
	params := map[string]string{"shuffle": fmt.Sprint(shuffle)}

	if offset > 0 {
		params["offset"] = string(offset)
	}

	if count > 0 {
		params["count"] = string(count)
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	var resp responses.NewsfeedGetSuggestedSources

	if err := n.SendObjRequest("newsfeed.getSuggestedSources", params, &resp); err != nil {
		return responses.NewsfeedGetSuggestedSources{}, err
	}

	return resp, nil

}

// IgnoreItem hides an item from the newsfeed
// Parameters:
//  * itemType - Item type. Possible values:
//      'wall' – post on the wall
//      'tag' – tag on a photo
//      'profilephoto' – profile photo
//      'video' – video
//      'audio' – audio.
//  * ownerId - Item owner's identifier (user or community). Note that community id must be negative. 'owner_id=1' – user , 'owner_id=-1' – community.
//  * itemId - Item identifier
func (n *Newsfeed) IgnoreItem(itemType string, ownerId, itemId int) (responses.OkResponse, error) {
	params := map[string]string{"type": itemType, "owner_id": string(ownerId), "item_id": string(itemId)}

	var resp responses.OkResponse

	if err := n.SendObjRequest("newsfeed.ignoreItem", params, &resp); err != nil {
		return responses.OkResponse(0), err
	}

	return resp, nil

}

// SaveList creates and edits user newsfeed lists
// Parameters:
//  * listId - numeric list identifier (if not sent, will be set automatically).
//  * title - list name.
//  * sourceIds - users and communities identifiers to be added to the list. Community identifiers must be negative numbers.
//  * noReposts - reposts display on and off ('1' is for off)
func (n *Newsfeed) SaveList(listId int, title string, sourceIds []int, noReposts bool) (responses.NewsfeedSaveList, error) {
	params := map[string]string{"title": title, "no_reposts": fmt.Sprint(noReposts)}

	if listId > 0 {
		params["list_id"] = string(listId)
	}

	if len(sourceIds) > 0 {
		params["source_ids"] = SliceToString(sourceIds)
	}

	var resp responses.NewsfeedSaveList

	if err := n.SendObjRequest("newsfeed.saveList", params, &resp); err != nil {
		return responses.NewsfeedSaveList(-1), err
	}

	return resp, nil
}

// Search returns search results by statuses.
//Parameters:
//  * query - Search query string (e.g., 'New Year').
//  * count - Number of posts to return.
//  * latitude - Geographical latitude point (in degrees, -90 to 90) within which to search.
//  * longitude - Geographical longitude point (in degrees, -180 to 180) within which to search.
//  * startTime - Earliest timestamp (in Unix time) of a news item to return. By default, 24 hours ago.
//  * endTime - Latest timestamp (in Unix time) of a news item to return. By default, the current time.
//  * startFrom -
//  * fields - Additional fields of [vk.com/dev/fields|profiles] and [vk.com/dev/fields_groups|communities] to return.
func (n *Newsfeed) Search(query, startFrom string, count, startTime, endTime int, latitude, longitude float64) (responses.NewsfeedSearch, error) {
	params := map[string]string{"extended": "0", "q": query}

	if latitude > 0 {
		params["latitude"] = fmt.Sprint(latitude)
	}

	if longitude > 0 {
		params["longitude"] = fmt.Sprint(longitude)
	}

	if len(startFrom) > 0 {
		params["start_from"] = startFrom
	}

	if count > 0 {
		params["count"] = string(count)
	}

	if startTime > 0 {
		params["start_time"] = string(startTime)
	}

	if endTime > 0 {
		params["end_time"] = string(endTime)
	}

	resp := responses.NewsfeedSearch{}

	if err := n.SendObjRequest("newsfeed.search", params, &resp); err != nil {
		return responses.NewsfeedSearch{}, err
	}

	return resp, nil
}

// SearchExt returns extended search results by statuses.
//Parameters:
//  * query - Search query string (e.g., 'New Year').
//  * count - Number of posts to return.
//  * latitude - Geographical latitude point (in degrees, -90 to 90) within which to search.
//  * longitude - Geographical longitude point (in degrees, -180 to 180) within which to search.
//  * start_time - Earliest timestamp (in Unix time) of a news item to return. By default, 24 hours ago.
//  * end_time - Latest timestamp (in Unix time) of a news item to return. By default, the current time.
//  * start_from -
//  * fields - Additional fields of [vk.com/dev/fields|profiles] and [vk.com/dev/fields_groups|communities] to return.
func (n *Newsfeed) SearchExt(query, startFrom string, fields []objects.BaseUserGroupFields, count, startTime, endTime int,
	latitude, longitude float64) (responses.NewsfeedSearchExt, error) {
	params := map[string]string{"extended": "1", "q": query}

	if latitude > 0 {
		params["latitude"] = fmt.Sprint(latitude)
	}

	if longitude > 0 {
		params["longitude"] = fmt.Sprint(longitude)
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	if len(startFrom) > 0 {
		params["start_from"] = startFrom
	}

	if count > 0 {
		params["count"] = string(count)
	}

	if startTime > 0 {
		params["start_time"] = string(startTime)
	}

	if endTime > 0 {
		params["end_time"] = string(endTime)
	}

	resp := responses.NewsfeedSearchExt{}

	if err := n.SendObjRequest("newsfeed.search", params, &resp); err != nil {
		return responses.NewsfeedSearchExt{}, err
	}

	return resp, nil
}

// UnignoreItem returns a hidden item to the newsfeed
// Parameters:
//  * itemType - Item type. Possible values:
//      'wall' – post on the wall
//      'tag' – tag on a photo
//      'profilephoto' – profile photo
//      'video' – video
//      'audio' – audio
//  * ownerId - Item owner's identifier (user or community), Note that community id must be negative. 'owner_id=1' – user , 'owner_id=-1' – community
//  * itemId - Item identifier
func (n *Newsfeed) UnignoreItem(itemType objects.NewsfeedIgnoreItemType, ownerId, itemId int) (responses.OkResponse, error) {
	params := map[string]string{"type": string(itemType), "owner_id": string(ownerId), "item_id": string(itemId)}
	var resp responses.OkResponse

	if err := n.SendObjRequest("newsfeed.unsubscribe", params, &resp); err != nil {
		return responses.OkResponse(0), err
	}

	return resp, nil
}

// Unsubscribe Unsubscribes the current user from specified newsfeeds.
//Parameters:
//  * objectType - Type of object from which to unsubscribe:
//      'note' — note
//      'photo' — photo
//      'post' — post on user wall or community wall
//      'topic' — topic
//      'video' — video
//  * ownerId - Object owner ID
//  * itemId - Object ID
func (n *Newsfeed) Unsubscribe(objectType string, ownerId, itemId int) (responses.OkResponse, error) {
	params := map[string]string{"type": objectType, "owner_id": string(ownerId), "item_id": string(itemId)}
	var resp responses.OkResponse

	if err := n.SendObjRequest("newsfeed.unsubscribe", params, &resp); err != nil {
		return responses.OkResponse(0), err
	}

	return resp, nil
}
