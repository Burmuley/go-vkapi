package go_vkapi

import (
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
	panic("implement me!")
}

// DeleteBan Allows news from previously banned users and communities to be shown in the current user's newsfeed
//
// Parameters:
//  * userIds - int array with user IDs
//  * groupIds - int array with group IDs
func (n *Newsfeed) DeleteBan(userIds []int, groupIds []int) (responses.OkResponse, error) {
	panic("implement me!")
}

// DeleteList deletes list
//
// Parameters:
//  * listId - int ID of the list to delete
func (n *Newsfeed) DeleteList(listId int) (responses.OkResponse, error) {
	panic("implement me!")
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
func (n *Newsfeed) Get() (responses.NewsfeedGetResponse, error) {
	panic("implement me!")
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
func (n *Newsfeed) GetBanned(fields []objects.UsersFields, nameCase string) {
	params := map[string]string{"extended": "1"}

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
func (n *Newsfeed) GetBannedExt(fields []objects.UsersFields, nameCase string) {
	params := map[string]string{"extended": "0"}

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
	panic("implement me!")
}

// GetLists returns a list of newsfeeds followed by the current user
//
//Parameters:
//  * listIds - numeric list identifiers.
func (n *Newsfeed) GetLists(listIds []int) (responses.NewsfeedGetLists, error) {
	params := map[string]string{"extended": "false"}
	panic("implement me!")
}

// GetListsExt returns an extended list of newsfeeds followed by the current user
//
// Parameters:
//  * listIds - numeric list identifiers.
func (n *Newsfeed) GetListsExt(listIds []int) (responses.NewsfeedGetLists, error) {
	params := map[string]string{"extended": "true"}
	panic("implement me!")
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
	panic("implement me!")
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
	panic("implement me!")
}

// GetSuggestedSources returns communities and users that current user is suggested to follow
// Parameters:
//  * offset - offset required to choose a particular subset of communities or users.
//  * count - amount of communities or users to return.
//  * shuffle - shuffle the returned list or not.
//  * fields - list of extra fields to be returned. See available fields for [vk.com/dev/fields|users] and [vk.com/dev/fields_groups|communities].
func (n *Newsfeed) GetSuggestedSources(offset, count int, shuffle bool,
	fields []objects.BaseUserGroupFields) (responses.NewsfeedGetSuggestedSources, error) {
	panic("implement me!")
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
	panic("implement me!")
}

// SaveList creates and edits user newsfeed lists
// Parameters:
//  * listId - numeric list identifier (if not sent, will be set automatically).
//  * title - list name.
//  * sourceIds - users and communities identifiers to be added to the list. Community identifiers must be negative numbers.
//  * noReposts - reposts display on and off ('1' is for off)
func (n *Newsfeed) SaveList(listId int, title string, sourceIds []int, noReposts bool) (responses.NewsfeedSaveList, error) {
	panic("implement me!")
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
func (n *Newsfeed) Search(query, startFrom string, fields []objects.BaseUserGroupFields, count, startTime, endTime int,
	latitude, longitude float64) (responses.NewsfeedSearch, error) {
	params := map[string]string{"extended": "0"}
	panic("implement me!")
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
func (n *Newsfeed) SearchExt() (responses.NewsfeedSearchExt, error) {
	params := map[string]string{"extended": "1"}
	panic("implement me!")
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
	panic("implement me!")
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
	panic("implement me!")
}
