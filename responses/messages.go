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
// Source schema can be found at https://github.com/VKCOM/vk-api-schema/blob/master/responses.json         //
// Code generator location: https://gitlab.com/Burmuley/go-vkapi-gen                                       //
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

package responses

import "gitlab.com/Burmuley/go-vkapi/objects"

/////////////////////////////////////////////////////////////
// `messages` group of responses
/////////////////////////////////////////////////////////////

// MessagesJoinChatByInviteLink type represents `messages_joinChatByInviteLink_response` API response object
type MessagesJoinChatByInviteLink struct {
	ChatId int `json:"chat_id"`
}

// MessagesGetChatChatIdsFields type represents `messages_getChat_chat_ids_fields_response` API response object
type MessagesGetChatChatIdsFields objects.MessagesChatFull

// MessagesGetConversationsById type represents `messages_getConversationsById_response` API response object
type MessagesGetConversationsById struct {
	Count int                            `json:"count"` // Total number
	Items []objects.MessagesConversation `json:"items"`
}

// MessagesGetChatChatIds type represents `messages_getChat_chat_ids_response` API response object
type MessagesGetChatChatIds objects.MessagesChat

// MessagesSearch type represents `messages_search_response` API response object
type MessagesSearch struct {
	Count int                       `json:"count"` // Total number
	Items []objects.MessagesMessage `json:"items"`
}

// MessagesSetChatPhoto type represents `messages_setChatPhoto_response` API response object
type MessagesSetChatPhoto struct {
	Chat      objects.MessagesChat `json:"chat"`
	MessageId int                  `json:"message_id"` // Service message ID
}

// MessagesSearchConversations type represents `messages_searchConversations_response` API response object
type MessagesSearchConversations struct {
	Count    int                            `json:"count"` // Total results number
	Groups   []objects.GroupsGroupFull      `json:"groups"`
	Items    []objects.MessagesConversation `json:"items"`
	Profiles []objects.UsersUserFull        `json:"profiles"`
}

// MessagesGetChat type represents `messages_getChat_response` API response object
type MessagesGetChat objects.MessagesChat

// MessagesGetConversationsByIdExtended type represents `messages_getConversationsById_extended_response` API response object
type MessagesGetConversationsByIdExtended struct {
	Count    int                            `json:"count"` // Total number
	Items    []objects.MessagesConversation `json:"items"`
	Profiles []objects.UsersUser            `json:"profiles"`
}

// MessagesEdit type represents `messages_edit_response` API response object
type MessagesEdit objects.BaseBoolInt // Result

// MessagesGetInviteLink type represents `messages_getInviteLink_response` API response object
type MessagesGetInviteLink struct {
	Link string `json:"link"`
}

// MessagesPin type represents `messages_pin_response` API response object
type MessagesPin objects.MessagesPinnedMessage

// MessagesMarkAsImportant type represents `messages_markAsImportant_response` API response object
type MessagesMarkAsImportant int

// MessagesGetLongPollHistory type represents `messages_getLongPollHistory_response` API response object
type MessagesGetLongPollHistory struct {
	Chats         []objects.MessagesChat           `json:"chats"`
	Conversations []objects.MessagesConversation   `json:"conversations"`
	Groups        []objects.GroupsGroup            `json:"groups"`
	History       []int                            `json:"history"`
	Messages      objects.MessagesLongpollMessages `json:"messages"`
	More          bool                             `json:"more"`    // Has more
	NewPts        int                              `json:"new_pts"` // Persistence timestamp
	Profiles      []objects.UsersUserFull          `json:"profiles"`
}

// MessagesDeleteChatPhoto type represents `messages_deleteChatPhoto_response` API response object
type MessagesDeleteChatPhoto struct {
	Chat      objects.MessagesChat `json:"chat"`
	MessageId int                  `json:"message_id"` // Service message ID
}

// MessagesCreateChat type represents `messages_createChat_response` API response object
type MessagesCreateChat int // Chat ID

// MessagesGetLastActivity type represents `messages_getLastActivity_response` API response object
type MessagesGetLastActivity objects.MessagesLastActivity

// MessagesGetHistory type represents `messages_getHistory_response` API response object
type MessagesGetHistory struct {
	Count    int                       `json:"count"` // Total number
	Groups   []objects.GroupsGroupFull `json:"groups"`
	Items    []objects.MessagesMessage `json:"items"`
	Profiles []objects.UsersUserFull   `json:"profiles"`
}

// MessagesGetByConversationMessageId type represents `messages_getByConversationMessageId_response` API response object
type MessagesGetByConversationMessageId struct {
	Count int                       `json:"count"` // Total number
	Items []objects.MessagesMessage `json:"items"`
}

// MessagesIsMessagesFromGroupAllowed type represents `messages_isMessagesFromGroupAllowed_response` API response object
type MessagesIsMessagesFromGroupAllowed struct {
	IsAllowed objects.BaseBoolInt `json:"is_allowed"`
}

// MessagesGetById type represents `messages_getById_response` API response object
type MessagesGetById struct {
	Count int                       `json:"count"` // Total number
	Items []objects.MessagesMessage `json:"items"`
}

// MessagesSend type represents `messages_send_response` API response object
type MessagesSend int // Message ID

// MessagesGetChatPreview type represents `messages_getChatPreview_response` API response object
type MessagesGetChatPreview struct {
	Preview  objects.MessageChatPreview `json:"preview"`
	Profiles []objects.UsersUserFull    `json:"profiles"`
}

// MessagesGetByIdExtended type represents `messages_getById_extended_response` API response object
type MessagesGetByIdExtended struct {
	Count    int                       `json:"count"` // Total number
	Groups   []objects.GroupsGroupFull `json:"groups"`
	Items    []objects.MessagesMessage `json:"items"`
	Profiles []objects.UsersUserFull   `json:"profiles"`
}

// MessagesGetConversations type represents `messages_getConversations_response` API response object
type MessagesGetConversations struct {
	Count       int                                       `json:"count"` // Total number
	Groups      []objects.GroupsGroupFull                 `json:"groups"`
	Items       []objects.MessagesConversationWithMessage `json:"items"`
	Profiles    []objects.UsersUserFull                   `json:"profiles"`
	UnreadCount int                                       `json:"unread_count"` // Unread dialogs number
}

// MessagesGetHistoryAttachments type represents `messages_getHistoryAttachments_response` API response object
type MessagesGetHistoryAttachments struct {
	Items    []objects.MessagesHistoryAttachment `json:"items"`
	NextFrom string                              `json:"next_from"` // Value for pagination
}

// MessagesGetChatFields type represents `messages_getChat_fields_response` API response object
type MessagesGetChatFields objects.MessagesChatFull

// MessagesGetConversationMembers type represents `messages_getConversationMembers_response` API response object
type MessagesGetConversationMembers struct {
	ChatRestrictions objects.MessagesChatRestrictions     `json:"chat_restrictions"`
	Count            int                                  `json:"count"` // Chat members count
	Groups           []objects.GroupsGroupFull            `json:"groups"`
	Items            []objects.MessagesConversationMember `json:"items"`
	Profiles         []objects.UsersUserFull              `json:"profiles"`
}

// MessagesSendUserIds type represents `messages_send_user_ids_response` API response object
type MessagesSendUserIds object

// MessagesDeleteConversation type represents `messages_deleteConversation_response` API response object
type MessagesDeleteConversation struct {
	LastDeletedId int `json:"last_deleted_id"` // Id of the last message, that was deleted
}

// MessagesGetLongPollServer type represents `messages_getLongPollServer_response` API response object
type MessagesGetLongPollServer objects.MessagesLongpollParams

// MessagesDelete type represents `messages_delete_response` API response object
type MessagesDelete struct {
}
