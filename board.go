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

type Board struct {
	*VKApi
}

/////////////////////////////////////////////////////////////
// `Board` methods
/////////////////////////////////////////////////////////////

// AddTopic - Creates a new topic on a community's discussion board.
// Parameters:
//   * groupId - ID of the community that owns the discussion board.
//   * title - Topic title.
//   * text - Text of the topic.
//   * fromGroup - For a community: '1' — to post the topic as by the community, '0' — to post the topic as by the user (default)
//   * attachments - List of media objects attached to the topic, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media object: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, '<owner_id>' — ID of the media owner. '<media_id>' — Media ID. Example: "photo100172_166443618,photo66748_265827614", , "NOTE: If you try to attach more than one reference, an error will be thrown.",
func (b Board) AddTopic(groupId int, title string, text string, fromGroup bool, attachments []string) (resp responses.BoardAddTopic, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	params["title"] = title

	if text != "" {
		params["text"] = text
	}

	params["from_group"] = fromGroup

	if len(attachments) > 0 {
		params["attachments"] = SliceToString(attachments)
	}

	err = b.SendObjRequest("board.addTopic", params, &resp)

	return
}

// CloseTopic - Closes a topic on a community's discussion board so that comments cannot be posted.
// Parameters:
//   * groupId - ID of the community that owns the discussion board.
//   * topicId - Topic ID.
func (b Board) CloseTopic(groupId int, topicId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	params["topic_id"] = topicId

	err = b.SendObjRequest("board.closeTopic", params, &resp)

	return
}

// CreateComment - Adds a comment on a topic on a community's discussion board.
// Parameters:
//   * groupId - ID of the community that owns the discussion board.
//   * topicId - ID of the topic to be commented on.
//   * message - (Required if 'attachments' is not set.) Text of the comment.
//   * attachments - (Required if 'text' is not set.) List of media objects attached to the comment, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media object: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, '<owner_id>' — ID of the media owner. '<media_id>' — Media ID.
//   * fromGroup - '1' — to post the comment as by the community, '0' — to post the comment as by the user (default)
//   * stickerId - Sticker ID.
//   * guid - Unique identifier to avoid repeated comments.
func (b Board) CreateComment(groupId int, topicId int, message string, attachments []string, fromGroup bool, stickerId int, guid string) (resp responses.BoardCreateComment, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	params["topic_id"] = topicId

	if message != "" {
		params["message"] = message
	}

	if len(attachments) > 0 {
		params["attachments"] = SliceToString(attachments)
	}

	params["from_group"] = fromGroup

	if stickerId > 0 {
		params["sticker_id"] = stickerId
	}

	if guid != "" {
		params["guid"] = guid
	}

	err = b.SendObjRequest("board.createComment", params, &resp)

	return
}

// DeleteComment - Deletes a comment on a topic on a community's discussion board.
// Parameters:
//   * groupId - ID of the community that owns the discussion board.
//   * topicId - Topic ID.
//   * commentId - Comment ID.
func (b Board) DeleteComment(groupId int, topicId int, commentId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	params["topic_id"] = topicId

	params["comment_id"] = commentId

	err = b.SendObjRequest("board.deleteComment", params, &resp)

	return
}

// DeleteTopic - Deletes a topic from a community's discussion board.
// Parameters:
//   * groupId - ID of the community that owns the discussion board.
//   * topicId - Topic ID.
func (b Board) DeleteTopic(groupId int, topicId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	params["topic_id"] = topicId

	err = b.SendObjRequest("board.deleteTopic", params, &resp)

	return
}

// EditComment - Edits a comment on a topic on a community's discussion board.
// Parameters:
//   * groupId - ID of the community that owns the discussion board.
//   * topicId - Topic ID.
//   * commentId - ID of the comment on the topic.
//   * message - (Required if 'attachments' is not set). New comment text.
//   * attachments - (Required if 'message' is not set.) List of media objects attached to the comment, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media object: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, '<owner_id>' — ID of the media owner. '<media_id>' — Media ID. Example: "photo100172_166443618,photo66748_265827614"
func (b Board) EditComment(groupId int, topicId int, commentId int, message string, attachments []string) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	params["topic_id"] = topicId

	params["comment_id"] = commentId

	if message != "" {
		params["message"] = message
	}

	if len(attachments) > 0 {
		params["attachments"] = SliceToString(attachments)
	}

	err = b.SendObjRequest("board.editComment", params, &resp)

	return
}

// EditTopic - Edits the title of a topic on a community's discussion board.
// Parameters:
//   * groupId - ID of the community that owns the discussion board.
//   * topicId - Topic ID.
//   * title - New title of the topic.
func (b Board) EditTopic(groupId int, topicId int, title string) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	params["topic_id"] = topicId

	params["title"] = title

	err = b.SendObjRequest("board.editTopic", params, &resp)

	return
}

// FixTopic - Pins a topic (fixes its place) to the top of a community's discussion board.
// Parameters:
//   * groupId - ID of the community that owns the discussion board.
//   * topicId - Topic ID.
func (b Board) FixTopic(groupId int, topicId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	params["topic_id"] = topicId

	err = b.SendObjRequest("board.fixTopic", params, &resp)

	return
}

// GetComments - Returns a list of comments on a topic on a community's discussion board.
// Parameters:
//   * groupId - ID of the community that owns the discussion board.
//   * topicId - Topic ID.
//   * needLikes - '1' — to return the 'likes' field, '0' — not to return the 'likes' field (default)
//   * startCommentId - NO DESCRIPTION IN JSON SCHEMA
//   * offset - Offset needed to return a specific subset of comments.
//   * count - Number of comments to return.
//   * extended - '1' — to return information about users who posted comments, '0' — to return no additional fields (default)
//   * sort - Sort order: 'asc' — by creation date in chronological order, 'desc' — by creation date in reverse chronological order,
func (b Board) GetComments(groupId int, topicId int, needLikes bool, startCommentId int, offset int, count int, sort string) (resp responses.BoardGetComments, err error) {
	params := map[string]interface{}{}
	params["extended"] = "0"

	params["group_id"] = groupId

	params["topic_id"] = topicId

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

	err = b.SendObjRequest("board.getComments", params, &resp)

	return
}

// GetCommentsExtended - Returns a list of comments on a topic on a community's discussion board.
// Parameters:
//   * groupId - ID of the community that owns the discussion board.
//   * topicId - Topic ID.
//   * needLikes - '1' — to return the 'likes' field, '0' — not to return the 'likes' field (default)
//   * startCommentId - NO DESCRIPTION IN JSON SCHEMA
//   * offset - Offset needed to return a specific subset of comments.
//   * count - Number of comments to return.
//   * extended - '1' — to return information about users who posted comments, '0' — to return no additional fields (default)
//   * sort - Sort order: 'asc' — by creation date in chronological order, 'desc' — by creation date in reverse chronological order,
func (b Board) GetCommentsExtended(groupId int, topicId int, needLikes bool, startCommentId int, offset int, count int, sort string) (resp responses.BoardGetCommentsExtended, err error) {
	params := map[string]interface{}{}
	params["extended"] = "1"

	params["group_id"] = groupId

	params["topic_id"] = topicId

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

	err = b.SendObjRequest("board.getComments", params, &resp)

	return
}

// GetTopics - Returns a list of topics on a community's discussion board.
// Parameters:
//   * groupId - ID of the community that owns the discussion board.
//   * topicIds - IDs of topics to be returned (100 maximum). By default, all topics are returned. If this parameter is set, the 'order', 'offset', and 'count' parameters are ignored.
//   * order - Sort order: '1' — by date updated in reverse chronological order. '2' — by date created in reverse chronological order. '-1' — by date updated in chronological order. '-2' — by date created in chronological order. If no sort order is specified, topics are returned in the order specified by the group administrator. Pinned topics are returned first, regardless of the sorting.
//   * offset - Offset needed to return a specific subset of topics.
//   * count - Number of topics to return.
//   * extended - '1' — to return information about users who created topics or who posted there last, '0' — to return no additional fields (default)
//   * preview - '1' — to return the first comment in each topic,, '2' — to return the last comment in each topic,, '0' — to return no comments. By default: '0'.
//   * previewLength - Number of characters after which to truncate the previewed comment. To preview the full comment, specify '0'.
func (b Board) GetTopics(groupId int, topicIds []int, order int, offset int, count int, preview int, previewLength int) (resp responses.BoardGetTopics, err error) {
	params := map[string]interface{}{}
	params["extended"] = "0"

	params["group_id"] = groupId

	if len(topicIds) > 0 {
		params["topic_ids"] = SliceToString(topicIds)
	}

	if order > 0 {
		params["order"] = order
	}

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	if preview > 0 {
		params["preview"] = preview
	}

	if previewLength > 0 {
		params["preview_length"] = previewLength
	}

	err = b.SendObjRequest("board.getTopics", params, &resp)

	return
}

// GetTopicsExtended - Returns a list of topics on a community's discussion board.
// Parameters:
//   * groupId - ID of the community that owns the discussion board.
//   * topicIds - IDs of topics to be returned (100 maximum). By default, all topics are returned. If this parameter is set, the 'order', 'offset', and 'count' parameters are ignored.
//   * order - Sort order: '1' — by date updated in reverse chronological order. '2' — by date created in reverse chronological order. '-1' — by date updated in chronological order. '-2' — by date created in chronological order. If no sort order is specified, topics are returned in the order specified by the group administrator. Pinned topics are returned first, regardless of the sorting.
//   * offset - Offset needed to return a specific subset of topics.
//   * count - Number of topics to return.
//   * extended - '1' — to return information about users who created topics or who posted there last, '0' — to return no additional fields (default)
//   * preview - '1' — to return the first comment in each topic,, '2' — to return the last comment in each topic,, '0' — to return no comments. By default: '0'.
//   * previewLength - Number of characters after which to truncate the previewed comment. To preview the full comment, specify '0'.
func (b Board) GetTopicsExtended(groupId int, topicIds []int, order int, offset int, count int, preview int, previewLength int) (resp responses.BoardGetTopicsExtended, err error) {
	params := map[string]interface{}{}
	params["extended"] = "1"

	params["group_id"] = groupId

	if len(topicIds) > 0 {
		params["topic_ids"] = SliceToString(topicIds)
	}

	if order > 0 {
		params["order"] = order
	}

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	if preview > 0 {
		params["preview"] = preview
	}

	if previewLength > 0 {
		params["preview_length"] = previewLength
	}

	err = b.SendObjRequest("board.getTopics", params, &resp)

	return
}

// OpenTopic - Re-opens a previously closed topic on a community's discussion board.
// Parameters:
//   * groupId - ID of the community that owns the discussion board.
//   * topicId - Topic ID.
func (b Board) OpenTopic(groupId int, topicId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	params["topic_id"] = topicId

	err = b.SendObjRequest("board.openTopic", params, &resp)

	return
}

// RestoreComment - Restores a comment deleted from a topic on a community's discussion board.
// Parameters:
//   * groupId - ID of the community that owns the discussion board.
//   * topicId - Topic ID.
//   * commentId - Comment ID.
func (b Board) RestoreComment(groupId int, topicId int, commentId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	params["topic_id"] = topicId

	params["comment_id"] = commentId

	err = b.SendObjRequest("board.restoreComment", params, &resp)

	return
}

// UnfixTopic - Unpins a pinned topic from the top of a community's discussion board.
// Parameters:
//   * groupId - ID of the community that owns the discussion board.
//   * topicId - Topic ID.
func (b Board) UnfixTopic(groupId int, topicId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	params["topic_id"] = topicId

	err = b.SendObjRequest("board.unfixTopic", params, &resp)

	return
}
