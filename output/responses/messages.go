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
// Code generator location: https://github.com/Burmuley/go-vkapi-gen                                       //
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

package responses

import (
	"github.com/Burmuley/go-vkapi/objects"
)

/////////////////////////////////////////////////////////////
// `messages` group of responses
/////////////////////////////////////////////////////////////

// MessagesCreatechat type represents `messages_createChat_response` API response object
type MessagesCreatechat int // Chat ID

// MessagesDeletechatphoto type represents `messages_deleteChatPhoto_response` API response object
type MessagesDeletechatphoto struct {
	Chat      objects.MessagesChat `json:"chat"`
	MessageId int                  `json:"message_id"` // Service message ID
}

// MessagesDeleteconversation type represents `messages_deleteConversation_response` API response object
type MessagesDeleteconversation struct {
	LastDeletedId int `json:"last_deleted_id"` // Id of the last message, that was deleted
}

// MessagesDelete type represents `messages_delete_response` API response object
type MessagesDelete struct {
}

// MessagesEdit type represents `messages_edit_response` API response object
type MessagesEdit objects.BaseBoolInt // Result

// MessagesGetbyconversationmessageid type represents `messages_getByConversationMessageId_response` API response object
type MessagesGetbyconversationmessageid struct {
	Count int                       `json:"count"` // Total number
	Items []objects.MessagesMessage `json:"items"`
}

// MessagesGetbyidExtended type represents `messages_getById_extended_response` API response object
type MessagesGetbyidExtended struct {
	Count    int                       `json:"count"` // Total number
	Groups   []objects.GroupsGroupFull `json:"groups"`
	Items    []objects.MessagesMessage `json:"items"`
	Profiles []objects.UsersUserFull   `json:"profiles"`
}

// MessagesGetbyid type represents `messages_getById_response` API response object
type MessagesGetbyid struct {
	Count int                       `json:"count"` // Total number
	Items []objects.MessagesMessage `json:"items"`
}

// MessagesGetchatpreview type represents `messages_getChatPreview_response` API response object
type MessagesGetchatpreview struct {
	Preview  objects.MessageChatPreview `json:"preview"`
	Profiles []objects.UsersUserFull    `json:"profiles"`
}

// MessagesGetchatChatIdsFields type represents `messages_getChat_chat_ids_fields_response` API response object
type MessagesGetchatChatIdsFields objects.MessagesChatFull

// MessagesGetchatChatIds type represents `messages_getChat_chat_ids_response` API response object
type MessagesGetchatChatIds objects.MessagesChat

// MessagesGetchatFields type represents `messages_getChat_fields_response` API response object
type MessagesGetchatFields objects.MessagesChatFull

// MessagesGetchat type represents `messages_getChat_response` API response object
type MessagesGetchat objects.MessagesChat

// MessagesGetconversationmembers type represents `messages_getConversationMembers_response` API response object
type MessagesGetconversationmembers struct {
	ChatRestrictions objects.MessagesChatRestrictions     `json:"chat_restrictions"`
	Count            int                                  `json:"count"` // Chat members count
	Groups           []objects.GroupsGroupFull            `json:"groups"`
	Items            []objects.MessagesConversationMember `json:"items"`
	Profiles         []objects.UsersUserFull              `json:"profiles"`
}

// MessagesGetconversationsbyidExtended type represents `messages_getConversationsById_extended_response` API response object
type MessagesGetconversationsbyidExtended struct {
	Count    int                            `json:"count"` // Total number
	Items    []objects.MessagesConversation `json:"items"`
	Profiles []objects.UsersUser            `json:"profiles"`
}

// MessagesGetconversationsbyid type represents `messages_getConversationsById_response` API response object
type MessagesGetconversationsbyid struct {
	Count int                            `json:"count"` // Total number
	Items []objects.MessagesConversation `json:"items"`
}

// MessagesGetconversations type represents `messages_getConversations_response` API response object
type MessagesGetconversations struct {
	Count       int                                       `json:"count"` // Total number
	Groups      []objects.GroupsGroupFull                 `json:"groups"`
	Items       []objects.MessagesConversationWithMessage `json:"items"`
	Profiles    []objects.UsersUserFull                   `json:"profiles"`
	UnreadCount int                                       `json:"unread_count"` // Unread dialogs number
}

// MessagesGethistoryattachments type represents `messages_getHistoryAttachments_response` API response object
type MessagesGethistoryattachments struct {
	Items    []objects.MessagesHistoryAttachment `json:"items"`
	NextFrom string                              `json:"next_from"` // Value for pagination
}

// MessagesGethistory type represents `messages_getHistory_response` API response object
type MessagesGethistory struct {
	Count    int                       `json:"count"` // Total number
	Groups   []objects.GroupsGroupFull `json:"groups"`
	Items    []objects.MessagesMessage `json:"items"`
	Profiles []objects.UsersUserFull   `json:"profiles"`
}

// MessagesGetinvitelink type represents `messages_getInviteLink_response` API response object
type MessagesGetinvitelink struct {
	Link string `json:"link"`
}

// MessagesGetlastactivity type represents `messages_getLastActivity_response` API response object
type MessagesGetlastactivity objects.MessagesLastActivity

// MessagesGetlongpollhistory type represents `messages_getLongPollHistory_response` API response object
type MessagesGetlongpollhistory struct {
	Chats         []objects.MessagesChat           `json:"chats"`
	Conversations []objects.MessagesConversation   `json:"conversations"`
	Groups        []objects.GroupsGroup            `json:"groups"`
	History       []int                            `json:"history"`
	Messages      objects.MessagesLongpollMessages `json:"messages"`
	More          bool                             `json:"more"`    // Has more
	NewPts        int                              `json:"new_pts"` // Persistence timestamp
	Profiles      []objects.UsersUserFull          `json:"profiles"`
}

// MessagesGetlongpollserver type represents `messages_getLongPollServer_response` API response object
type MessagesGetlongpollserver objects.MessagesLongpollParams

// MessagesIsmessagesfromgroupallowed type represents `messages_isMessagesFromGroupAllowed_response` API response object
type MessagesIsmessagesfromgroupallowed struct {
	IsAllowed objects.BaseBoolInt `json:"is_allowed"`
}

// MessagesJoinchatbyinvitelink type represents `messages_joinChatByInviteLink_response` API response object
type MessagesJoinchatbyinvitelink struct {
	ChatId int `json:"chat_id"`
}

// MessagesMarkasimportant type represents `messages_markAsImportant_response` API response object
type MessagesMarkasimportant int

// MessagesPin type represents `messages_pin_response` API response object
type MessagesPin objects.MessagesPinnedMessage

// MessagesSearchconversations type represents `messages_searchConversations_response` API response object
type MessagesSearchconversations struct {
	Count    int                            `json:"count"` // Total results number
	Groups   []objects.GroupsGroupFull      `json:"groups"`
	Items    []objects.MessagesConversation `json:"items"`
	Profiles []objects.UsersUserFull        `json:"profiles"`
}

// MessagesSearch type represents `messages_search_response` API response object
type MessagesSearch struct {
	Count int                       `json:"count"` // Total number
	Items []objects.MessagesMessage `json:"items"`
}

// MessagesSend type represents `messages_send_response` API response object
type MessagesSend int // Message ID

// MessagesSendUserIds type represents `messages_send_user_ids_response` API response object
type MessagesSendUserIds interface{}

// MessagesSetchatphoto type represents `messages_setChatPhoto_response` API response object
type MessagesSetchatphoto struct {
	Chat      objects.MessagesChat `json:"chat"`
	MessageId int                  `json:"message_id"` // Service message ID
}
