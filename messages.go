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

type Messages struct {
	*VKApi
}

/////////////////////////////////////////////////////////////
// `Messages` methods
/////////////////////////////////////////////////////////////

// Addchatuser - Adds a new user to a chat.
// Parameters:
//   * chatId - Chat ID.
//   * userId - ID of the user to be added to the chat.
func (m Messages) Addchatuser(chatId int, userId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["chat_id"] = chatId

	if userId > 0 {
		params["user_id"] = userId
	}

	err = m.SendObjRequest("messages.addChatUser", params, &resp)

	return
}

// Allowmessagesfromgroup - Allows sending messages from community to the current user.
// Parameters:
//   * groupId - Group ID.
//   * key - NO DESCRIPTION IN JSON SCHEMA
func (m Messages) Allowmessagesfromgroup(groupId int, key string) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	if key != "" {
		params["key"] = key
	}

	err = m.SendObjRequest("messages.allowMessagesFromGroup", params, &resp)

	return
}

// Createchat - Creates a chat with several participants.
// Parameters:
//   * userIds - IDs of the users to be added to the chat.
//   * title - Chat title.
func (m Messages) Createchat(userIds []int, title string) (resp responses.MessagesCreatechat, err error) {
	params := map[string]interface{}{}

	if len(userIds) > 0 {
		params["user_ids"] = SliceToString(userIds)
	}

	if title != "" {
		params["title"] = title
	}

	err = m.SendObjRequest("messages.createChat", params, &resp)

	return
}

// Delete - Deletes one or more messages.
// Parameters:
//   * messageIds - Message IDs.
//   * spam - '1' — to mark message as spam.
//   * groupId - Group ID (for group messages with user access token)
//   * deleteForAll - '1' — delete message for for all.
func (m Messages) Delete(messageIds []int, spam bool, groupId int, deleteForAll bool) (resp responses.MessagesDelet, err error) {
	params := map[string]interface{}{}

	if len(messageIds) > 0 {
		params["message_ids"] = SliceToString(messageIds)
	}

	params["spam"] = spam

	if groupId > 0 {
		params["group_id"] = groupId
	}

	params["delete_for_all"] = deleteForAll

	err = m.SendObjRequest("messages.delete", params, &resp)

	return
}

// Deletechatphoto - Deletes a chat's cover picture.
// Parameters:
//   * chatId - Chat ID.
//   * groupId - NO DESCRIPTION IN JSON SCHEMA
func (m Messages) Deletechatphoto(chatId int, groupId int) (resp responses.MessagesDeletechatphot, err error) {
	params := map[string]interface{}{}

	params["chat_id"] = chatId

	if groupId > 0 {
		params["group_id"] = groupId
	}

	err = m.SendObjRequest("messages.deleteChatPhoto", params, &resp)

	return
}

// Deleteconversation - Deletes all private messages in a conversation.
// Parameters:
//   * userId - User ID. To clear a chat history use 'chat_id'
//   * peerId - Destination ID. "For user: 'User ID', e.g. '12345'. For chat: '2000000000' + 'chat_id', e.g. '2000000001'. For community: '- community ID', e.g. '-12345'. "
//   * groupId - Group ID (for group messages with user access token)
func (m Messages) Deleteconversation(userId int, peerId int, groupId int) (resp responses.MessagesDeleteconversati, err error) {
	params := map[string]interface{}{}

	if userId > 0 {
		params["user_id"] = userId
	}

	if peerId > 0 {
		params["peer_id"] = peerId
	}

	if groupId > 0 {
		params["group_id"] = groupId
	}

	err = m.SendObjRequest("messages.deleteConversation", params, &resp)

	return
}

// Denymessagesfromgroup - Denies sending message from community to the current user.
// Parameters:
//   * groupId - Group ID.
func (m Messages) Denymessagesfromgroup(groupId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	err = m.SendObjRequest("messages.denyMessagesFromGroup", params, &resp)

	return
}

// Edit - Edits the message.
// Parameters:
//   * peerId - Destination ID. "For user: 'User ID', e.g. '12345'. For chat: '2000000000' + 'chat_id', e.g. '2000000001'. For community: '- community ID', e.g. '-12345'. "
//   * message - (Required if 'attachments' is not set.) Text of the message.
//   * messageId - NO DESCRIPTION IN JSON SCHEMA
//   * lat - Geographical latitude of a check-in, in degrees (from -90 to 90).
//   * long - Geographical longitude of a check-in, in degrees (from -180 to 180).
//   * attachment - (Required if 'message' is not set.) List of objects attached to the message, separated by commas, in the following format: "<owner_id>_<media_id>", '' — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, 'wall' — wall post, '<owner_id>' — ID of the media attachment owner. '<media_id>' — media attachment ID. Example: "photo100172_166443618"
//   * keepForwardMessages - '1' — to keep forwarded, messages.
//   * keepSnippets - '1' — to keep attached snippets.
//   * groupId - Group ID (for group messages with user access token)
//   * dontParseLinks - NO DESCRIPTION IN JSON SCHEMA
func (m Messages) Edit(peerId int, message string, messageId int, lat float64, long float64, attachment string, keepForwardMessages bool, keepSnippets bool, groupId int, dontParseLinks bool) (resp responses.MessagesEdit, err error) {
	params := map[string]interface{}{}

	params["peer_id"] = peerId

	if message != "" {
		params["message"] = message
	}

	params["message_id"] = messageId

	if lat > 0 {
		params["lat"] = lat
	}

	if long > 0 {
		params["long"] = long
	}

	if attachment != "" {
		params["attachment"] = attachment
	}

	params["keep_forward_messages"] = keepForwardMessages

	params["keep_snippets"] = keepSnippets

	if groupId > 0 {
		params["group_id"] = groupId
	}

	params["dont_parse_links"] = dontParseLinks

	err = m.SendObjRequest("messages.edit", params, &resp)

	return
}

// Editchat - Edits the title of a chat.
// Parameters:
//   * chatId - Chat ID.
//   * title - New title of the chat.
func (m Messages) Editchat(chatId int, title string) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["chat_id"] = chatId

	params["title"] = title

	err = m.SendObjRequest("messages.editChat", params, &resp)

	return
}

// Getbyconversationmessageid - Returns messages by their IDs within the conversation.
// Parameters:
//   * peerId - Destination ID. "For user: 'User ID', e.g. '12345'. For chat: '2000000000' + 'chat_id', e.g. '2000000001'. For community: '- community ID', e.g. '-12345'. "
//   * conversationMessageIds - Conversation message IDs.
//   * extended - Information whether the response should be extended
//   * fields - Profile fields to return.
//   * groupId - Group ID (for group messages with group access token)
func (m Messages) Getbyconversationmessageid(peerId int, conversationMessageIds []int, fields []objects.UsersFields, groupId int) (resp responses.MessagesGetbyconversationmessageid, err error) {
	params := map[string]interface{}{}

	params["peer_id"] = peerId

	params["conversation_message_ids"] = SliceToString(conversationMessageIds)

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	if groupId > 0 {
		params["group_id"] = groupId
	}

	err = m.SendObjRequest("messages.getByConversationMessageId", params, &resp)

	return
}

// Getbyid - Returns messages by their IDs.
// Parameters:
//   * messageIds - Message IDs.
//   * previewLength - Number of characters after which to truncate a previewed message. To preview the full message, specify '0'. "NOTE: Messages are not truncated by default. Messages are truncated by words."
//   * extended - Information whether the response should be extended
//   * fields - Profile fields to return.
//   * groupId - Group ID (for group messages with group access token)
func (m Messages) Getbyid(messageIds []int, previewLength int, fields []objects.UsersFields, groupId int) (resp responses.MessagesGetbyid, err error) {
	params := map[string]interface{}{}
	params["extended"] = "0"

	params["message_ids"] = SliceToString(messageIds)

	if previewLength > 0 {
		params["preview_length"] = previewLength
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	if groupId > 0 {
		params["group_id"] = groupId
	}

	err = m.SendObjRequest("messages.getById", params, &resp)

	return
}

// GetbyidExtended - Returns messages by their IDs.
// Parameters:
//   * messageIds - Message IDs.
//   * previewLength - Number of characters after which to truncate a previewed message. To preview the full message, specify '0'. "NOTE: Messages are not truncated by default. Messages are truncated by words."
//   * extended - Information whether the response should be extended
//   * fields - Profile fields to return.
//   * groupId - Group ID (for group messages with group access token)
func (m Messages) GetbyidExtended(messageIds []int, previewLength int, fields []objects.UsersFields, groupId int) (resp responses.MessagesGetbyidExtended, err error) {
	params := map[string]interface{}{}
	params["extended"] = "1"

	params["message_ids"] = SliceToString(messageIds)

	if previewLength > 0 {
		params["preview_length"] = previewLength
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	if groupId > 0 {
		params["group_id"] = groupId
	}

	err = m.SendObjRequest("messages.getById", params, &resp)

	return
}

// Getchatpreview - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * peerId - NO DESCRIPTION IN JSON SCHEMA
//   * link - Invitation link.
//   * fields - Profile fields to return.
func (m Messages) Getchatpreview(peerId int, link string, fields []objects.UsersFields) (resp responses.MessagesGetchatpreview, err error) {
	params := map[string]interface{}{}

	if peerId > 0 {
		params["peer_id"] = peerId
	}

	if link != "" {
		params["link"] = link
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	err = m.SendObjRequest("messages.getChatPreview", params, &resp)

	return
}

// Getconversationmembers - Returns a list of IDs of users participating in a chat.
// Parameters:
//   * peerId - Peer ID.
//   * fields - Profile fields to return.
//   * groupId - Group ID (for group messages with group access token)
func (m Messages) Getconversationmembers(peerId int, fields []objects.UsersFields, groupId int) (resp responses.MessagesGetconversationmember, err error) {
	params := map[string]interface{}{}

	params["peer_id"] = peerId

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	if groupId > 0 {
		params["group_id"] = groupId
	}

	err = m.SendObjRequest("messages.getConversationMembers", params, &resp)

	return
}

// Getconversations - Returns a list of the current user's conversations.
// Parameters:
//   * offset - Offset needed to return a specific subset of conversations.
//   * count - Number of conversations to return.
//   * filter - Filter to apply: 'all' — all conversations, 'unread' — conversations with unread messages, 'important' — conversations, marked as important (only for community messages), 'unanswered' — conversations, marked as unanswered (only for community messages)
//   * extended - '1' — return extra information about users and communities
//   * startMessageId - ID of the message from what to return dialogs.
//   * fields - Profile and communities fields to return.
//   * groupId - Group ID (for group messages with group access token)
func (m Messages) Getconversations(offset int, count int, filter string, startMessageId int, fields []objects.BaseUserGroupFields, groupId int) (resp responses.MessagesGetconversati, err error) {
	params := map[string]interface{}{}

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	if filter != "" {
		params["filter"] = filter
	}

	if startMessageId > 0 {
		params["start_message_id"] = startMessageId
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	if groupId > 0 {
		params["group_id"] = groupId
	}

	err = m.SendObjRequest("messages.getConversations", params, &resp)

	return
}

// Getconversationsbyid - Returns conversations by their IDs
// Parameters:
//   * peerIds - Destination IDs. "For user: 'User ID', e.g. '12345'. For chat: '2000000000' + 'chat_id', e.g. '2000000001'. For community: '- community ID', e.g. '-12345'. "
//   * extended - Return extended properties
//   * fields - Profile and communities fields to return.
//   * groupId - Group ID (for group messages with group access token)
func (m Messages) Getconversationsbyid(peerIds []int, fields []objects.BaseUserGroupFields, groupId int) (resp responses.MessagesGetconversationsbyid, err error) {
	params := map[string]interface{}{}
	params["extended"] = "0"

	params["peer_ids"] = SliceToString(peerIds)

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	if groupId > 0 {
		params["group_id"] = groupId
	}

	err = m.SendObjRequest("messages.getConversationsById", params, &resp)

	return
}

// GetconversationsbyidExtended - Returns conversations by their IDs
// Parameters:
//   * peerIds - Destination IDs. "For user: 'User ID', e.g. '12345'. For chat: '2000000000' + 'chat_id', e.g. '2000000001'. For community: '- community ID', e.g. '-12345'. "
//   * extended - Return extended properties
//   * fields - Profile and communities fields to return.
//   * groupId - Group ID (for group messages with group access token)
func (m Messages) GetconversationsbyidExtended(peerIds []int, fields []objects.BaseUserGroupFields, groupId int) (resp responses.MessagesGetconversationsbyidExtended, err error) {
	params := map[string]interface{}{}
	params["extended"] = "1"

	params["peer_ids"] = SliceToString(peerIds)

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	if groupId > 0 {
		params["group_id"] = groupId
	}

	err = m.SendObjRequest("messages.getConversationsById", params, &resp)

	return
}

// Gethistory - Returns message history for the specified user or group chat.
// Parameters:
//   * offset - Offset needed to return a specific subset of messages.
//   * count - Number of messages to return.
//   * userId - ID of the user whose message history you want to return.
//   * peerId - NO DESCRIPTION IN JSON SCHEMA
//   * startMessageId - Starting message ID from which to return history.
//   * rev - Sort order: '1' — return messages in chronological order. '0' — return messages in reverse chronological order.
//   * extended - Information whether the response should be extended
//   * fields - Profile fields to return.
//   * groupId - Group ID (for group messages with group access token)
func (m Messages) Gethistory(offset int, count int, userId int, peerId int, startMessageId int, rev int, fields []objects.UsersFields, groupId int) (resp responses.MessagesGethistory, err error) {
	params := map[string]interface{}{}

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	if userId > 0 {
		params["user_id"] = userId
	}

	if peerId > 0 {
		params["peer_id"] = peerId
	}

	if startMessageId > 0 {
		params["start_message_id"] = startMessageId
	}

	if rev > 0 {
		params["rev"] = rev
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	if groupId > 0 {
		params["group_id"] = groupId
	}

	err = m.SendObjRequest("messages.getHistory", params, &resp)

	return
}

// Gethistoryattachments - Returns media files from the dialog or group chat.
// Parameters:
//   * peerId - Peer ID. ", For group chat: '2000000000 + chat ID' , , For community: '-community ID'"
//   * mediaType - Type of media files to return: *'photo',, *'video',, *'audio',, *'doc',, *'link'.,*'market'.,*'wall'.,*'share'
//   * startFrom - Message ID to start return results from.
//   * count - Number of objects to return.
//   * photoSizes - '1' — to return photo sizes in a
//   * fields - Additional profile [vk.com/dev/fields|fields] to return.
//   * groupId - Group ID (for group messages with group access token)
//   * preserveOrder - NO DESCRIPTION IN JSON SCHEMA
//   * maxForwardsLevel - NO DESCRIPTION IN JSON SCHEMA
func (m Messages) Gethistoryattachments(peerId int, mediaType string, startFrom string, count int, photoSizes bool, fields []objects.UsersFields, groupId int, preserveOrder bool, maxForwardsLevel int) (resp responses.MessagesGethistoryattachment, err error) {
	params := map[string]interface{}{}

	params["peer_id"] = peerId

	if mediaType != "" {
		params["media_type"] = mediaType
	}

	if startFrom != "" {
		params["start_from"] = startFrom
	}

	if count > 0 {
		params["count"] = count
	}

	params["photo_sizes"] = photoSizes

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	if groupId > 0 {
		params["group_id"] = groupId
	}

	params["preserve_order"] = preserveOrder

	if maxForwardsLevel > 0 {
		params["max_forwards_level"] = maxForwardsLevel
	}

	err = m.SendObjRequest("messages.getHistoryAttachments", params, &resp)

	return
}

// Getinvitelink - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * peerId - Destination ID.
//   * reset - 1 — to generate new link (revoke previous), 0 — to return previous link.
//   * groupId - Group ID
func (m Messages) Getinvitelink(peerId int, reset bool, groupId int) (resp responses.MessagesGetinvitelink, err error) {
	params := map[string]interface{}{}

	params["peer_id"] = peerId

	params["reset"] = reset

	if groupId > 0 {
		params["group_id"] = groupId
	}

	err = m.SendObjRequest("messages.getInviteLink", params, &resp)

	return
}

// Getlastactivity - Returns a user's current status and date of last activity.
// Parameters:
//   * userId - User ID.
func (m Messages) Getlastactivity(userId int) (resp responses.MessagesGetlastactivity, err error) {
	params := map[string]interface{}{}

	params["user_id"] = userId

	err = m.SendObjRequest("messages.getLastActivity", params, &resp)

	return
}

// Getlongpollhistory - Returns updates in user's private messages.
// Parameters:
//   * ts - Last value of the 'ts' parameter returned from the Long Poll server or by using [vk.com/dev/messages.getLongPollHistory|messages.getLongPollHistory] method.
//   * pts - Lsat value of 'pts' parameter returned from the Long Poll server or by using [vk.com/dev/messages.getLongPollHistory|messages.getLongPollHistory] method.
//   * previewLength - Number of characters after which to truncate a previewed message. To preview the full message, specify '0'. "NOTE: Messages are not truncated by default. Messages are truncated by words."
//   * onlines - '1' — to return history with online users only.
//   * fields - Additional profile [vk.com/dev/fields|fields] to return.
//   * eventsLimit - Maximum number of events to return.
//   * msgsLimit - Maximum number of messages to return.
//   * maxMsgId - Maximum ID of the message among existing ones in the local copy. Both messages received with API methods (for example, , ), and data received from a Long Poll server (events with code 4) are taken into account.
//   * groupId - Group ID (for group messages with user access token)
//   * lpVersion - NO DESCRIPTION IN JSON SCHEMA
//   * lastN - NO DESCRIPTION IN JSON SCHEMA
//   * credentials - NO DESCRIPTION IN JSON SCHEMA
func (m Messages) Getlongpollhistory(ts int, pts int, previewLength int, onlines bool, fields []objects.UsersFields, eventsLimit int, msgsLimit int, maxMsgId int, groupId int, lpVersion int, lastN int, credentials bool) (resp responses.MessagesGetlongpollhistory, err error) {
	params := map[string]interface{}{}

	if ts > 0 {
		params["ts"] = ts
	}

	if pts > 0 {
		params["pts"] = pts
	}

	if previewLength > 0 {
		params["preview_length"] = previewLength
	}

	params["onlines"] = onlines

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	if eventsLimit > 0 {
		params["events_limit"] = eventsLimit
	}

	if msgsLimit > 0 {
		params["msgs_limit"] = msgsLimit
	}

	if maxMsgId > 0 {
		params["max_msg_id"] = maxMsgId
	}

	if groupId > 0 {
		params["group_id"] = groupId
	}

	if lpVersion > 0 {
		params["lp_version"] = lpVersion
	}

	if lastN > 0 {
		params["last_n"] = lastN
	}

	params["credentials"] = credentials

	err = m.SendObjRequest("messages.getLongPollHistory", params, &resp)

	return
}

// Getlongpollserver - Returns data required for connection to a Long Poll server.
// Parameters:
//   * needPts - '1' — to return the 'pts' field, needed for the [vk.com/dev/messages.getLongPollHistory|messages.getLongPollHistory] method.
//   * groupId - Group ID (for group messages with user access token)
//   * lpVersion - Long poll version
func (m Messages) Getlongpollserver(needPts bool, groupId int, lpVersion int) (resp responses.MessagesGetlongpollserver, err error) {
	params := map[string]interface{}{}

	params["need_pts"] = needPts

	if groupId > 0 {
		params["group_id"] = groupId
	}

	if lpVersion > 0 {
		params["lp_version"] = lpVersion
	}

	err = m.SendObjRequest("messages.getLongPollServer", params, &resp)

	return
}

// Ismessagesfromgroupallowed - Returns information whether sending messages from the community to current user is allowed.
// Parameters:
//   * groupId - Group ID.
//   * userId - User ID.
func (m Messages) Ismessagesfromgroupallowed(groupId int, userId int) (resp responses.MessagesIsmessagesfromgroupallowed, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	params["user_id"] = userId

	err = m.SendObjRequest("messages.isMessagesFromGroupAllowed", params, &resp)

	return
}

// Joinchatbyinvitelink - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * link - Invitation link.
func (m Messages) Joinchatbyinvitelink(link string) (resp responses.MessagesJoinchatbyinvitelink, err error) {
	params := map[string]interface{}{}

	params["link"] = link

	err = m.SendObjRequest("messages.joinChatByInviteLink", params, &resp)

	return
}

// Markasansweredconversation - Marks and unmarks conversations as unanswered.
// Parameters:
//   * peerId - ID of conversation to mark as important.
//   * answered - '1' — to mark as answered, '0' — to remove the mark
//   * groupId - Group ID (for group messages with group access token)
func (m Messages) Markasansweredconversation(peerId int, answered bool, groupId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["peer_id"] = peerId

	params["answered"] = answered

	if groupId > 0 {
		params["group_id"] = groupId
	}

	err = m.SendObjRequest("messages.markAsAnsweredConversation", params, &resp)

	return
}

// Markasimportant - Marks and unmarks messages as important (starred).
// Parameters:
//   * messageIds - IDs of messages to mark as important.
//   * important - '1' — to add a star (mark as important), '0' — to remove the star
func (m Messages) Markasimportant(messageIds []int, important int) (resp responses.MessagesMarkasimportant, err error) {
	params := map[string]interface{}{}

	if len(messageIds) > 0 {
		params["message_ids"] = SliceToString(messageIds)
	}

	if important > 0 {
		params["important"] = important
	}

	err = m.SendObjRequest("messages.markAsImportant", params, &resp)

	return
}

// Markasimportantconversation - Marks and unmarks conversations as important.
// Parameters:
//   * peerId - ID of conversation to mark as important.
//   * important - '1' — to add a star (mark as important), '0' — to remove the star
//   * groupId - Group ID (for group messages with group access token)
func (m Messages) Markasimportantconversation(peerId int, important bool, groupId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["peer_id"] = peerId

	params["important"] = important

	if groupId > 0 {
		params["group_id"] = groupId
	}

	err = m.SendObjRequest("messages.markAsImportantConversation", params, &resp)

	return
}

// Markasread - Marks messages as read.
// Parameters:
//   * messageIds - IDs of messages to mark as read.
//   * peerId - Destination ID. "For user: 'User ID', e.g. '12345'. For chat: '2000000000' + 'chat_id', e.g. '2000000001'. For community: '- community ID', e.g. '-12345'. "
//   * startMessageId - Message ID to start from.
//   * groupId - Group ID (for group messages with user access token)
func (m Messages) Markasread(messageIds []int, peerId int, startMessageId int, groupId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	if len(messageIds) > 0 {
		params["message_ids"] = SliceToString(messageIds)
	}

	if peerId > 0 {
		params["peer_id"] = peerId
	}

	if startMessageId > 0 {
		params["start_message_id"] = startMessageId
	}

	if groupId > 0 {
		params["group_id"] = groupId
	}

	err = m.SendObjRequest("messages.markAsRead", params, &resp)

	return
}

// Pin - Pin a message.
// Parameters:
//   * peerId - Destination ID. "For user: 'User ID', e.g. '12345'. For chat: '2000000000' + 'Chat ID', e.g. '2000000001'. For community: '- Community ID', e.g. '-12345'. "
//   * messageId - NO DESCRIPTION IN JSON SCHEMA
func (m Messages) Pin(peerId int, messageId int) (resp responses.MessagesPi, err error) {
	params := map[string]interface{}{}

	params["peer_id"] = peerId

	if messageId > 0 {
		params["message_id"] = messageId
	}

	err = m.SendObjRequest("messages.pin", params, &resp)

	return
}

// Removechatuser - Allows the current user to leave a chat or, if the current user started the chat, allows the user to remove another user from the chat.
// Parameters:
//   * chatId - Chat ID.
//   * userId - ID of the user to be removed from the chat.
//   * memberId - NO DESCRIPTION IN JSON SCHEMA
func (m Messages) Removechatuser(chatId int, userId int, memberId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["chat_id"] = chatId

	if userId > 0 {
		params["user_id"] = userId
	}

	if memberId > 0 {
		params["member_id"] = memberId
	}

	err = m.SendObjRequest("messages.removeChatUser", params, &resp)

	return
}

// Restore - Restores a deleted message.
// Parameters:
//   * messageId - ID of a previously-deleted message to restore.
//   * groupId - Group ID (for group messages with user access token)
func (m Messages) Restore(messageId int, groupId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["message_id"] = messageId

	if groupId > 0 {
		params["group_id"] = groupId
	}

	err = m.SendObjRequest("messages.restore", params, &resp)

	return
}

// Search - Returns a list of the current user's private messages that match search criteria.
// Parameters:
//   * q - Search query string.
//   * peerId - Destination ID. "For user: 'User ID', e.g. '12345'. For chat: '2000000000' + 'chat_id', e.g. '2000000001'. For community: '- community ID', e.g. '-12345'. "
//   * date - Date to search message before in Unixtime.
//   * previewLength - Number of characters after which to truncate a previewed message. To preview the full message, specify '0'. "NOTE: Messages are not truncated by default. Messages are truncated by words."
//   * offset - Offset needed to return a specific subset of messages.
//   * count - Number of messages to return.
//   * extended - NO DESCRIPTION IN JSON SCHEMA
//   * fields - NO DESCRIPTION IN JSON SCHEMA
//   * groupId - Group ID (for group messages with group access token)
func (m Messages) Search(q string, peerId int, date int, previewLength int, offset int, count int, fields []string, groupId int) (resp responses.MessagesSearch, err error) {
	params := map[string]interface{}{}

	if q != "" {
		params["q"] = q
	}

	if peerId > 0 {
		params["peer_id"] = peerId
	}

	if date > 0 {
		params["date"] = date
	}

	if previewLength > 0 {
		params["preview_length"] = previewLength
	}

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	if groupId > 0 {
		params["group_id"] = groupId
	}

	err = m.SendObjRequest("messages.search", params, &resp)

	return
}

// Searchconversations - Returns a list of the current user's conversations that match search criteria.
// Parameters:
//   * q - Search query string.
//   * count - Maximum number of results.
//   * extended - '1' — return extra information about users and communities
//   * fields - Profile fields to return.
//   * groupId - Group ID (for group messages with user access token)
func (m Messages) Searchconversations(q string, count int, fields []objects.UsersFields, groupId int) (resp responses.MessagesSearchconversati, err error) {
	params := map[string]interface{}{}

	if q != "" {
		params["q"] = q
	}

	if count > 0 {
		params["count"] = count
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	if groupId > 0 {
		params["group_id"] = groupId
	}

	err = m.SendObjRequest("messages.searchConversations", params, &resp)

	return
}

// Send - Sends a message.
// Parameters:
//   * userId - User ID (by default — current user).
//   * randomId - Unique identifier to avoid resending the message.
//   * peerId - Destination ID. "For user: 'User ID', e.g. '12345'. For chat: '2000000000' + 'chat_id', e.g. '2000000001'. For community: '- community ID', e.g. '-12345'. "
//   * domain - User's short address (for example, 'illarionov').
//   * chatId - ID of conversation the message will relate to.
//   * userIds - IDs of message recipients (if new conversation shall be started).
//   * message - (Required if 'attachments' is not set.) Text of the message.
//   * lat - Geographical latitude of a check-in, in degrees (from -90 to 90).
//   * long - Geographical longitude of a check-in, in degrees (from -180 to 180).
//   * attachment - (Required if 'message' is not set.) List of objects attached to the message, separated by commas, in the following format: "<owner_id>_<media_id>", '' — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, 'wall' — wall post, '<owner_id>' — ID of the media attachment owner. '<media_id>' — media attachment ID. Example: "photo100172_166443618"
//   * replyTo - NO DESCRIPTION IN JSON SCHEMA
//   * forwardMessages - ID of forwarded messages, separated with a comma. Listed messages of the sender will be shown in the message body at the recipient's. Example: "123,431,544"
//   * forward - NO DESCRIPTION IN JSON SCHEMA
//   * stickerId - Sticker id.
//   * groupId - Group ID (for group messages with group access token)
//   * keyboard - NO DESCRIPTION IN JSON SCHEMA
//   * payload - NO DESCRIPTION IN JSON SCHEMA
//   * dontParseLinks - NO DESCRIPTION IN JSON SCHEMA
//   * disableMentions - NO DESCRIPTION IN JSON SCHEMA
func (m Messages) Send(userId int, randomId int, peerId int, domain string, chatId int, userIds []int, message string, lat float64, long float64, attachment string, replyTo int, forwardMessages []int, forward string, stickerId int, groupId int, keyboard objects.MessagesKeyboard, payload string, dontParseLinks bool, disableMentions bool) (resp responses.MessagesSend, err error) {
	params := map[string]interface{}{}

	if userId > 0 {
		params["user_id"] = userId
	}

	if randomId > 0 {
		params["random_id"] = randomId
	}

	if peerId > 0 {
		params["peer_id"] = peerId
	}

	if domain != "" {
		params["domain"] = domain
	}

	if chatId > 0 {
		params["chat_id"] = chatId
	}

	if len(userIds) > 0 {
		params["user_ids"] = SliceToString(userIds)
	}

	if message != "" {
		params["message"] = message
	}

	if lat > 0 {
		params["lat"] = lat
	}

	if long > 0 {
		params["long"] = long
	}

	if attachment != "" {
		params["attachment"] = attachment
	}

	if replyTo > 0 {
		params["reply_to"] = replyTo
	}

	if len(forwardMessages) > 0 {
		params["forward_messages"] = SliceToString(forwardMessages)
	}

	if forward != "" {
		params["forward"] = forward
	}

	if stickerId > 0 {
		params["sticker_id"] = stickerId
	}

	if groupId > 0 {
		params["group_id"] = groupId
	}

	if payload != "" {
		params["payload"] = payload
	}

	params["dont_parse_links"] = dontParseLinks

	params["disable_mentions"] = disableMentions

	err = m.SendObjRequest("messages.send", params, &resp)

	return
}

// Setactivity - Changes the status of a user as typing in a conversation.
// Parameters:
//   * userId - User ID.
//   * pType - 'typing' — user has started to type.
//   * peerId - Destination ID. "For user: 'User ID', e.g. '12345'. For chat: '2000000000' + 'chat_id', e.g. '2000000001'. For community: '- community ID', e.g. '-12345'. "
//   * groupId - Group ID (for group messages with group access token)
func (m Messages) Setactivity(userId int, pType string, peerId int, groupId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	if userId > 0 {
		params["user_id"] = userId
	}

	if pType != "" {
		params["type"] = pType
	}

	if peerId > 0 {
		params["peer_id"] = peerId
	}

	if groupId > 0 {
		params["group_id"] = groupId
	}

	err = m.SendObjRequest("messages.setActivity", params, &resp)

	return
}

// Setchatphoto - Sets a previously-uploaded picture as the cover picture of a chat.
// Parameters:
//   * file - Upload URL from the 'response' field returned by the [vk.com/dev/photos.getChatUploadServer|photos.getChatUploadServer] method upon successfully uploading an image.
func (m Messages) Setchatphoto(file string) (resp responses.MessagesSetchatphot, err error) {
	params := map[string]interface{}{}

	params["file"] = file

	err = m.SendObjRequest("messages.setChatPhoto", params, &resp)

	return
}

// Unpin - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * peerId - NO DESCRIPTION IN JSON SCHEMA
//   * groupId - NO DESCRIPTION IN JSON SCHEMA
func (m Messages) Unpin(peerId int, groupId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["peer_id"] = peerId

	if groupId > 0 {
		params["group_id"] = groupId
	}

	err = m.SendObjRequest("messages.unpin", params, &resp)

	return
}
