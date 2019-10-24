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

type Wall struct {
	*VKApi
}

/////////////////////////////////////////////////////////////
// `Wall` methods
/////////////////////////////////////////////////////////////

// Closecomments - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * ownerId - NO DESCRIPTION IN JSON SCHEMA
//   * postId - NO DESCRIPTION IN JSON SCHEMA
func (w Wall) Closecomments(ownerId int, postId int) (resp responses.BaseBool, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	params["post_id"] = postId

	err = w.SendObjRequest("wall.closeComments", params, &resp)

	return
}

// Createcomment - Adds a comment to a post on a user wall or community wall.
// Parameters:
//   * ownerId - User ID or community ID. Use a negative value to designate a community ID.
//   * postId - Post ID.
//   * fromGroup - Group ID.
//   * message - (Required if 'attachments' is not set.) Text of the comment.
//   * replyToComment - ID of comment to reply.
//   * attachments - (Required if 'message' is not set.) List of media objects attached to the comment, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media ojbect: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, '<owner_id>' — ID of the media owner. '<media_id>' — Media ID. For example: "photo100172_166443618,photo66748_265827614"
//   * stickerId - Sticker ID.
//   * guid - Unique identifier to avoid repeated comments.
func (w Wall) Createcomment(ownerId int, postId int, fromGroup int, message string, replyToComment int, attachments []string, stickerId int, guid string) (resp responses.WallCreatecomment, err error) {
	params := map[string]interface{}{}

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	params["post_id"] = postId

	if fromGroup > 0 {
		params["from_group"] = fromGroup
	}

	if message != "" {
		params["message"] = message
	}

	if replyToComment > 0 {
		params["reply_to_comment"] = replyToComment
	}

	if len(attachments) > 0 {
		params["attachments"] = SliceToString(attachments)
	}

	if stickerId > 0 {
		params["sticker_id"] = stickerId
	}

	if guid != "" {
		params["guid"] = guid
	}

	err = w.SendObjRequest("wall.createComment", params, &resp)

	return
}

// Delete - Deletes a post from a user wall or community wall.
// Parameters:
//   * ownerId - User ID or community ID. Use a negative value to designate a community ID.
//   * postId - ID of the post to be deleted.
func (w Wall) Delete(ownerId int, postId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	if postId > 0 {
		params["post_id"] = postId
	}

	err = w.SendObjRequest("wall.delete", params, &resp)

	return
}

// Deletecomment - Deletes a comment on a post on a user wall or community wall.
// Parameters:
//   * ownerId - User ID or community ID. Use a negative value to designate a community ID.
//   * commentId - Comment ID.
func (w Wall) Deletecomment(ownerId int, commentId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	params["comment_id"] = commentId

	err = w.SendObjRequest("wall.deleteComment", params, &resp)

	return
}

// Edit - Edits a post on a user wall or community wall.
// Parameters:
//   * ownerId - User ID or community ID. Use a negative value to designate a community ID.
//   * postId - NO DESCRIPTION IN JSON SCHEMA
//   * friendsOnly - NO DESCRIPTION IN JSON SCHEMA
//   * message - (Required if 'attachments' is not set.) Text of the post.
//   * attachments - (Required if 'message' is not set.) List of objects attached to the post, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, '<owner_id>' — ID of the media application owner. '<media_id>' — Media application ID. Example: "photo100172_166443618,photo66748_265827614", May contain a link to an external page to include in the post. Example: "photo66748_265827614,http://habrahabr.ru", "NOTE: If more than one link is being attached, an error is thrown."
//   * services - NO DESCRIPTION IN JSON SCHEMA
//   * signed - NO DESCRIPTION IN JSON SCHEMA
//   * publishDate - NO DESCRIPTION IN JSON SCHEMA
//   * lat - NO DESCRIPTION IN JSON SCHEMA
//   * long - NO DESCRIPTION IN JSON SCHEMA
//   * placeId - NO DESCRIPTION IN JSON SCHEMA
//   * markAsAds - NO DESCRIPTION IN JSON SCHEMA
//   * closeComments - NO DESCRIPTION IN JSON SCHEMA
//   * posterBkgId - NO DESCRIPTION IN JSON SCHEMA
//   * posterBkgOwnerId - NO DESCRIPTION IN JSON SCHEMA
//   * posterBkgAccessHash - NO DESCRIPTION IN JSON SCHEMA
func (w Wall) Edit(ownerId int, postId int, friendsOnly bool, message string, attachments []string, services string, signed bool, publishDate int, lat float64, long float64, placeId int, markAsAds bool, closeComments bool, posterBkgId int, posterBkgOwnerId int, posterBkgAccessHash string) (resp responses.WallEdit, err error) {
	params := map[string]interface{}{}

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	params["post_id"] = postId

	params["friends_only"] = friendsOnly

	if message != "" {
		params["message"] = message
	}

	if len(attachments) > 0 {
		params["attachments"] = SliceToString(attachments)
	}

	if services != "" {
		params["services"] = services
	}

	params["signed"] = signed

	if publishDate > 0 {
		params["publish_date"] = publishDate
	}

	if lat > 0 {
		params["lat"] = lat
	}

	if long > 0 {
		params["long"] = long
	}

	if placeId > 0 {
		params["place_id"] = placeId
	}

	params["mark_as_ads"] = markAsAds

	params["close_comments"] = closeComments

	if posterBkgId > 0 {
		params["poster_bkg_id"] = posterBkgId
	}

	if posterBkgOwnerId > 0 {
		params["poster_bkg_owner_id"] = posterBkgOwnerId
	}

	if posterBkgAccessHash != "" {
		params["poster_bkg_access_hash"] = posterBkgAccessHash
	}

	err = w.SendObjRequest("wall.edit", params, &resp)

	return
}

// Editadsstealth - Allows to edit hidden post.
// Parameters:
//   * ownerId - User ID or community ID. Use a negative value to designate a community ID.
//   * postId - Post ID. Used for publishing of scheduled and suggested posts.
//   * message - (Required if 'attachments' is not set.) Text of the post.
//   * attachments - (Required if 'message' is not set.) List of objects attached to the post, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, 'page' — wiki-page, 'note' — note, 'poll' — poll, 'album' — photo album, '<owner_id>' — ID of the media application owner. '<media_id>' — Media application ID. Example: "photo100172_166443618,photo66748_265827614", May contain a link to an external page to include in the post. Example: "photo66748_265827614,http://habrahabr.ru", "NOTE: If more than one link is being attached, an error will be thrown."
//   * signed - Only for posts in communities with 'from_group' set to '1': '1' — post will be signed with the name of the posting user, '0' — post will not be signed (default)
//   * lat - Geographical latitude of a check-in, in degrees (from -90 to 90).
//   * long - Geographical longitude of a check-in, in degrees (from -180 to 180).
//   * placeId - ID of the location where the user was tagged.
//   * linkButton - Link button ID
//   * linkTitle - Link title
//   * linkImage - Link image url
//   * linkVideo - Link video ID in format "<owner_id>_<media_id>"
func (w Wall) Editadsstealth(ownerId int, postId int, message string, attachments []string, signed bool, lat float64, long float64, placeId int, linkButton string, linkTitle string, linkImage string, linkVideo string) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	params["post_id"] = postId

	if message != "" {
		params["message"] = message
	}

	if len(attachments) > 0 {
		params["attachments"] = SliceToString(attachments)
	}

	params["signed"] = signed

	if lat > 0 {
		params["lat"] = lat
	}

	if long > 0 {
		params["long"] = long
	}

	if placeId > 0 {
		params["place_id"] = placeId
	}

	if linkButton != "" {
		params["link_button"] = linkButton
	}

	if linkTitle != "" {
		params["link_title"] = linkTitle
	}

	if linkImage != "" {
		params["link_image"] = linkImage
	}

	if linkVideo != "" {
		params["link_video"] = linkVideo
	}

	err = w.SendObjRequest("wall.editAdsStealth", params, &resp)

	return
}

// Editcomment - Edits a comment on a user wall or community wall.
// Parameters:
//   * ownerId - User ID or community ID. Use a negative value to designate a community ID.
//   * commentId - Comment ID.
//   * message - New comment text.
//   * attachments - List of objects attached to the comment, in the following format: , "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, '<owner_id>' — ID of the media attachment owner. '<media_id>' — Media attachment ID. For example: "photo100172_166443618,photo66748_265827614"
func (w Wall) Editcomment(ownerId int, commentId int, message string, attachments []string) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	params["comment_id"] = commentId

	if message != "" {
		params["message"] = message
	}

	if len(attachments) > 0 {
		params["attachments"] = SliceToString(attachments)
	}

	err = w.SendObjRequest("wall.editComment", params, &resp)

	return
}

// Get - Returns a list of posts on a user wall or community wall.
// Parameters:
//   * ownerId - ID of the user or community that owns the wall. By default, current user ID. Use a negative value to designate a community ID.
//   * domain - User or community short address.
//   * offset - Offset needed to return a specific subset of posts.
//   * count - Number of posts to return (maximum 100).
//   * filter - Filter to apply: 'owner' — posts by the wall owner, 'others' — posts by someone else, 'all' — posts by the wall owner and others (default), 'postponed' — timed posts (only available for calls with an 'access_token'), 'suggests' — suggested posts on a community wall
//   * extended - '1' — to return 'wall', 'profiles', and 'groups' fields, '0' — to return no additional fields (default)
//   * fields - NO DESCRIPTION IN JSON SCHEMA
func (w Wall) Get(ownerId int, domain string, offset int, count int, filter string, fields []objects.BaseUserGroupFields) (resp responses.WallGet, err error) {
	params := map[string]interface{}{}
	params["extended"] = "0"

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	if domain != "" {
		params["domain"] = domain
	}

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	if filter != "" {
		params["filter"] = filter
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	err = w.SendObjRequest("wall.get", params, &resp)

	return
}

// GetExtended - Returns a list of posts on a user wall or community wall.
// Parameters:
//   * ownerId - ID of the user or community that owns the wall. By default, current user ID. Use a negative value to designate a community ID.
//   * domain - User or community short address.
//   * offset - Offset needed to return a specific subset of posts.
//   * count - Number of posts to return (maximum 100).
//   * filter - Filter to apply: 'owner' — posts by the wall owner, 'others' — posts by someone else, 'all' — posts by the wall owner and others (default), 'postponed' — timed posts (only available for calls with an 'access_token'), 'suggests' — suggested posts on a community wall
//   * extended - '1' — to return 'wall', 'profiles', and 'groups' fields, '0' — to return no additional fields (default)
//   * fields - NO DESCRIPTION IN JSON SCHEMA
func (w Wall) GetExtended(ownerId int, domain string, offset int, count int, filter string, fields []objects.BaseUserGroupFields) (resp responses.WallGetExtended, err error) {
	params := map[string]interface{}{}
	params["extended"] = "1"

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	if domain != "" {
		params["domain"] = domain
	}

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	if filter != "" {
		params["filter"] = filter
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	err = w.SendObjRequest("wall.get", params, &resp)

	return
}

// Getbyid - Returns a list of posts from user or community walls by their IDs.
// Parameters:
//   * posts - User or community IDs and post IDs, separated by underscores. Use a negative value to designate a community ID. Example: "93388_21539,93388_20904,2943_4276,-1_1"
//   * extended - '1' — to return user and community objects needed to display posts, '0' — no additional fields are returned (default)
//   * copyHistoryDepth - Sets the number of parent elements to include in the array 'copy_history' that is returned if the post is a repost from another wall.
//   * fields - NO DESCRIPTION IN JSON SCHEMA
func (w Wall) Getbyid(posts []string, copyHistoryDepth int, fields []objects.BaseUserGroupFields) (resp responses.WallGetbyid, err error) {
	params := map[string]interface{}{}
	params["extended"] = "0"

	params["posts"] = SliceToString(posts)

	if copyHistoryDepth > 0 {
		params["copy_history_depth"] = copyHistoryDepth
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	err = w.SendObjRequest("wall.getById", params, &resp)

	return
}

// GetbyidExtended - Returns a list of posts from user or community walls by their IDs.
// Parameters:
//   * posts - User or community IDs and post IDs, separated by underscores. Use a negative value to designate a community ID. Example: "93388_21539,93388_20904,2943_4276,-1_1"
//   * extended - '1' — to return user and community objects needed to display posts, '0' — no additional fields are returned (default)
//   * copyHistoryDepth - Sets the number of parent elements to include in the array 'copy_history' that is returned if the post is a repost from another wall.
//   * fields - NO DESCRIPTION IN JSON SCHEMA
func (w Wall) GetbyidExtended(posts []string, copyHistoryDepth int, fields []objects.BaseUserGroupFields) (resp responses.WallGetbyidExtended, err error) {
	params := map[string]interface{}{}
	params["extended"] = "1"

	params["posts"] = SliceToString(posts)

	if copyHistoryDepth > 0 {
		params["copy_history_depth"] = copyHistoryDepth
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	err = w.SendObjRequest("wall.getById", params, &resp)

	return
}

// Getcomments - Returns a list of comments on a post on a user wall or community wall.
// Parameters:
//   * ownerId - User ID or community ID. Use a negative value to designate a community ID.
//   * postId - Post ID.
//   * needLikes - '1' — to return the 'likes' field, '0' — not to return the 'likes' field (default)
//   * startCommentId - NO DESCRIPTION IN JSON SCHEMA
//   * offset - Offset needed to return a specific subset of comments.
//   * count - Number of comments to return (maximum 100).
//   * sort - Sort order: 'asc' — chronological, 'desc' — reverse chronological
//   * previewLength - Number of characters at which to truncate comments when previewed. By default, '90'. Specify '0' if you do not want to truncate comments.
//   * extended - NO DESCRIPTION IN JSON SCHEMA
//   * fields - NO DESCRIPTION IN JSON SCHEMA
//   * commentId - Comment ID.
//   * threadItemsCount - Count items in threads.
func (w Wall) Getcomments(ownerId int, postId int, needLikes bool, startCommentId int, offset int, count int, sort string, previewLength int, fields []objects.BaseUserGroupFields, commentId int, threadItemsCount int) (resp responses.WallGetcomment, err error) {
	params := map[string]interface{}{}
	params["extended"] = "0"

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	if postId > 0 {
		params["post_id"] = postId
	}

	params["need_likes"] = needLikes

	if startCommentId > 0 {
		params["start_comment_id"] = startCommentId
	}

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	if sort != "" {
		params["sort"] = sort
	}

	if previewLength > 0 {
		params["preview_length"] = previewLength
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	if commentId > 0 {
		params["comment_id"] = commentId
	}

	if threadItemsCount > 0 {
		params["thread_items_count"] = threadItemsCount
	}

	err = w.SendObjRequest("wall.getComments", params, &resp)

	return
}

// GetcommentsExtended - Returns a list of comments on a post on a user wall or community wall.
// Parameters:
//   * ownerId - User ID or community ID. Use a negative value to designate a community ID.
//   * postId - Post ID.
//   * needLikes - '1' — to return the 'likes' field, '0' — not to return the 'likes' field (default)
//   * startCommentId - NO DESCRIPTION IN JSON SCHEMA
//   * offset - Offset needed to return a specific subset of comments.
//   * count - Number of comments to return (maximum 100).
//   * sort - Sort order: 'asc' — chronological, 'desc' — reverse chronological
//   * previewLength - Number of characters at which to truncate comments when previewed. By default, '90'. Specify '0' if you do not want to truncate comments.
//   * extended - NO DESCRIPTION IN JSON SCHEMA
//   * fields - NO DESCRIPTION IN JSON SCHEMA
//   * commentId - Comment ID.
//   * threadItemsCount - Count items in threads.
func (w Wall) GetcommentsExtended(ownerId int, postId int, needLikes bool, startCommentId int, offset int, count int, sort string, previewLength int, fields []objects.BaseUserGroupFields, commentId int, threadItemsCount int) (resp responses.WallGetcommentsExtended, err error) {
	params := map[string]interface{}{}
	params["extended"] = "1"

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	if postId > 0 {
		params["post_id"] = postId
	}

	params["need_likes"] = needLikes

	if startCommentId > 0 {
		params["start_comment_id"] = startCommentId
	}

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	if sort != "" {
		params["sort"] = sort
	}

	if previewLength > 0 {
		params["preview_length"] = previewLength
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	if commentId > 0 {
		params["comment_id"] = commentId
	}

	if threadItemsCount > 0 {
		params["thread_items_count"] = threadItemsCount
	}

	err = w.SendObjRequest("wall.getComments", params, &resp)

	return
}

// Getreposts - Returns information about reposts of a post on user wall or community wall.
// Parameters:
//   * ownerId - User ID or community ID. By default, current user ID. Use a negative value to designate a community ID.
//   * postId - Post ID.
//   * offset - Offset needed to return a specific subset of reposts.
//   * count - Number of reposts to return.
func (w Wall) Getreposts(ownerId int, postId int, offset int, count int) (resp responses.WallGetrepost, err error) {
	params := map[string]interface{}{}

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	if postId > 0 {
		params["post_id"] = postId
	}

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	err = w.SendObjRequest("wall.getReposts", params, &resp)

	return
}

// Opencomments - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * ownerId - NO DESCRIPTION IN JSON SCHEMA
//   * postId - NO DESCRIPTION IN JSON SCHEMA
func (w Wall) Opencomments(ownerId int, postId int) (resp responses.BaseBool, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	params["post_id"] = postId

	err = w.SendObjRequest("wall.openComments", params, &resp)

	return
}

// Pin - Pins the post on wall.
// Parameters:
//   * ownerId - ID of the user or community that owns the wall. By default, current user ID. Use a negative value to designate a community ID.
//   * postId - Post ID.
func (w Wall) Pin(ownerId int, postId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	params["post_id"] = postId

	err = w.SendObjRequest("wall.pin", params, &resp)

	return
}

// Post - Adds a new post on a user wall or community wall. Can also be used to publish suggested or scheduled posts.
// Parameters:
//   * ownerId - User ID or community ID. Use a negative value to designate a community ID.
//   * friendsOnly - '1' — post will be available to friends only, '0' — post will be available to all users (default)
//   * fromGroup - For a community: '1' — post will be published by the community, '0' — post will be published by the user (default)
//   * message - (Required if 'attachments' is not set.) Text of the post.
//   * attachments - (Required if 'message' is not set.) List of objects attached to the post, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, 'page' — wiki-page, 'note' — note, 'poll' — poll, 'album' — photo album, '<owner_id>' — ID of the media application owner. '<media_id>' — Media application ID. Example: "photo100172_166443618,photo66748_265827614", May contain a link to an external page to include in the post. Example: "photo66748_265827614,http://habrahabr.ru", "NOTE: If more than one link is being attached, an error will be thrown."
//   * services - List of services or websites the update will be exported to, if the user has so requested. Sample values: 'twitter', 'facebook'.
//   * signed - Only for posts in communities with 'from_group' set to '1': '1' — post will be signed with the name of the posting user, '0' — post will not be signed (default)
//   * publishDate - Publication date (in Unix time). If used, posting will be delayed until the set time.
//   * lat - Geographical latitude of a check-in, in degrees (from -90 to 90).
//   * long - Geographical longitude of a check-in, in degrees (from -180 to 180).
//   * placeId - ID of the location where the user was tagged.
//   * postId - Post ID. Used for publishing of scheduled and suggested posts.
//   * guid - NO DESCRIPTION IN JSON SCHEMA
//   * markAsAds - NO DESCRIPTION IN JSON SCHEMA
//   * closeComments - NO DESCRIPTION IN JSON SCHEMA
//   * muteNotifications - NO DESCRIPTION IN JSON SCHEMA
func (w Wall) Post(ownerId int, friendsOnly bool, fromGroup bool, message string, attachments []string, services string, signed bool, publishDate int, lat float64, long float64, placeId int, postId int, guid string, markAsAds bool, closeComments bool, muteNotifications bool) (resp responses.WallPost, err error) {
	params := map[string]interface{}{}

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	params["friends_only"] = friendsOnly

	params["from_group"] = fromGroup

	if message != "" {
		params["message"] = message
	}

	if len(attachments) > 0 {
		params["attachments"] = SliceToString(attachments)
	}

	if services != "" {
		params["services"] = services
	}

	params["signed"] = signed

	if publishDate > 0 {
		params["publish_date"] = publishDate
	}

	if lat > 0 {
		params["lat"] = lat
	}

	if long > 0 {
		params["long"] = long
	}

	if placeId > 0 {
		params["place_id"] = placeId
	}

	if postId > 0 {
		params["post_id"] = postId
	}

	if guid != "" {
		params["guid"] = guid
	}

	params["mark_as_ads"] = markAsAds

	params["close_comments"] = closeComments

	params["mute_notifications"] = muteNotifications

	err = w.SendObjRequest("wall.post", params, &resp)

	return
}

// Postadsstealth - Allows to create hidden post which will not be shown on the community's wall and can be used for creating an ad with type "Community post".
// Parameters:
//   * ownerId - User ID or community ID. Use a negative value to designate a community ID.
//   * message - (Required if 'attachments' is not set.) Text of the post.
//   * attachments - (Required if 'message' is not set.) List of objects attached to the post, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, 'page' — wiki-page, 'note' — note, 'poll' — poll, 'album' — photo album, '<owner_id>' — ID of the media application owner. '<media_id>' — Media application ID. Example: "photo100172_166443618,photo66748_265827614", May contain a link to an external page to include in the post. Example: "photo66748_265827614,http://habrahabr.ru", "NOTE: If more than one link is being attached, an error will be thrown."
//   * signed - Only for posts in communities with 'from_group' set to '1': '1' — post will be signed with the name of the posting user, '0' — post will not be signed (default)
//   * lat - Geographical latitude of a check-in, in degrees (from -90 to 90).
//   * long - Geographical longitude of a check-in, in degrees (from -180 to 180).
//   * placeId - ID of the location where the user was tagged.
//   * guid - Unique identifier to avoid duplication the same post.
//   * linkButton - Link button ID
//   * linkTitle - Link title
//   * linkImage - Link image url
//   * linkVideo - Link video ID in format "<owner_id>_<media_id>"
func (w Wall) Postadsstealth(ownerId int, message string, attachments []string, signed bool, lat float64, long float64, placeId int, guid string, linkButton string, linkTitle string, linkImage string, linkVideo string) (resp responses.WallPostadsstealth, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	if message != "" {
		params["message"] = message
	}

	if len(attachments) > 0 {
		params["attachments"] = SliceToString(attachments)
	}

	params["signed"] = signed

	if lat > 0 {
		params["lat"] = lat
	}

	if long > 0 {
		params["long"] = long
	}

	if placeId > 0 {
		params["place_id"] = placeId
	}

	if guid != "" {
		params["guid"] = guid
	}

	if linkButton != "" {
		params["link_button"] = linkButton
	}

	if linkTitle != "" {
		params["link_title"] = linkTitle
	}

	if linkImage != "" {
		params["link_image"] = linkImage
	}

	if linkVideo != "" {
		params["link_video"] = linkVideo
	}

	err = w.SendObjRequest("wall.postAdsStealth", params, &resp)

	return
}

// Reportcomment - Reports (submits a complaint about) a comment on a post on a user wall or community wall.
// Parameters:
//   * ownerId - ID of the user or community that owns the wall.
//   * commentId - Comment ID.
//   * reason - Reason for the complaint: '0' – spam, '1' – child pornography, '2' – extremism, '3' – violence, '4' – drug propaganda, '5' – adult material, '6' – insult, abuse
func (w Wall) Reportcomment(ownerId int, commentId int, reason int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	params["comment_id"] = commentId

	if reason > 0 {
		params["reason"] = reason
	}

	err = w.SendObjRequest("wall.reportComment", params, &resp)

	return
}

// Reportpost - Reports (submits a complaint about) a post on a user wall or community wall.
// Parameters:
//   * ownerId - ID of the user or community that owns the wall.
//   * postId - Post ID.
//   * reason - Reason for the complaint: '0' – spam, '1' – child pornography, '2' – extremism, '3' – violence, '4' – drug propaganda, '5' – adult material, '6' – insult, abuse
func (w Wall) Reportpost(ownerId int, postId int, reason int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	params["post_id"] = postId

	if reason > 0 {
		params["reason"] = reason
	}

	err = w.SendObjRequest("wall.reportPost", params, &resp)

	return
}

// Repost - Reposts (copies) an object to a user wall or community wall.
// Parameters:
//   * object - ID of the object to be reposted on the wall. Example: "wall66748_3675"
//   * message - Comment to be added along with the reposted object.
//   * groupId - Target community ID when reposting to a community.
//   * markAsAds - NO DESCRIPTION IN JSON SCHEMA
//   * muteNotifications - NO DESCRIPTION IN JSON SCHEMA
func (w Wall) Repost(object string, message string, groupId int, markAsAds bool, muteNotifications bool) (resp responses.WallRepost, err error) {
	params := map[string]interface{}{}

	params["object"] = object

	if message != "" {
		params["message"] = message
	}

	if groupId > 0 {
		params["group_id"] = groupId
	}

	params["mark_as_ads"] = markAsAds

	params["mute_notifications"] = muteNotifications

	err = w.SendObjRequest("wall.repost", params, &resp)

	return
}

// Restore - Restores a post deleted from a user wall or community wall.
// Parameters:
//   * ownerId - User ID or community ID from whose wall the post was deleted. Use a negative value to designate a community ID.
//   * postId - ID of the post to be restored.
func (w Wall) Restore(ownerId int, postId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	if postId > 0 {
		params["post_id"] = postId
	}

	err = w.SendObjRequest("wall.restore", params, &resp)

	return
}

// Restorecomment - Restores a comment deleted from a user wall or community wall.
// Parameters:
//   * ownerId - User ID or community ID. Use a negative value to designate a community ID.
//   * commentId - Comment ID.
func (w Wall) Restorecomment(ownerId int, commentId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	params["comment_id"] = commentId

	err = w.SendObjRequest("wall.restoreComment", params, &resp)

	return
}

// Search - Allows to search posts on user or community walls.
// Parameters:
//   * ownerId - user or community id. "Remember that for a community 'owner_id' must be negative."
//   * domain - user or community screen name.
//   * query - search query string.
//   * ownersOnly - '1' – returns only page owner's posts.
//   * count - count of posts to return.
//   * offset - Offset needed to return a specific subset of posts.
//   * extended - show extended post info.
//   * fields - NO DESCRIPTION IN JSON SCHEMA
func (w Wall) Search(ownerId int, domain string, query string, ownersOnly bool, count int, offset int, fields []objects.BaseUserGroupFields) (resp responses.WallSearch, err error) {
	params := map[string]interface{}{}
	params["extended"] = "0"

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	if domain != "" {
		params["domain"] = domain
	}

	if query != "" {
		params["query"] = query
	}

	params["owners_only"] = ownersOnly

	if count > 0 {
		params["count"] = count
	}

	if offset > 0 {
		params["offset"] = offset
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	err = w.SendObjRequest("wall.search", params, &resp)

	return
}

// SearchExtended - Allows to search posts on user or community walls.
// Parameters:
//   * ownerId - user or community id. "Remember that for a community 'owner_id' must be negative."
//   * domain - user or community screen name.
//   * query - search query string.
//   * ownersOnly - '1' – returns only page owner's posts.
//   * count - count of posts to return.
//   * offset - Offset needed to return a specific subset of posts.
//   * extended - show extended post info.
//   * fields - NO DESCRIPTION IN JSON SCHEMA
func (w Wall) SearchExtended(ownerId int, domain string, query string, ownersOnly bool, count int, offset int, fields []objects.BaseUserGroupFields) (resp responses.WallSearchExtended, err error) {
	params := map[string]interface{}{}
	params["extended"] = "1"

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	if domain != "" {
		params["domain"] = domain
	}

	if query != "" {
		params["query"] = query
	}

	params["owners_only"] = ownersOnly

	if count > 0 {
		params["count"] = count
	}

	if offset > 0 {
		params["offset"] = offset
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	err = w.SendObjRequest("wall.search", params, &resp)

	return
}

// Unpin - Unpins the post on wall.
// Parameters:
//   * ownerId - ID of the user or community that owns the wall. By default, current user ID. Use a negative value to designate a community ID.
//   * postId - Post ID.
func (w Wall) Unpin(ownerId int, postId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	params["post_id"] = postId

	err = w.SendObjRequest("wall.unpin", params, &resp)

	return
}
