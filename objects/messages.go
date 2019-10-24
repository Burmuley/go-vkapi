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

package objects

/////////////////////////////////////////////////////////////////////////////////////////////////////////////
// `messages` group of objects
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

// MessagesAudioMessage type represents `messages_audio_message` API object
type MessagesAudioMessage struct {
	AccessKey string `json:"access_key"` // Access key for audio message
	Duration  int    `json:"duration"`   // Audio message duration in seconds
	Id        int    `json:"id"`         // Audio message ID
	LinkMp3   string `json:"link_mp3"`   // MP3 file URL
	LinkOgg   string `json:"link_ogg"`   // OGG file URL
	OwnerId   int    `json:"owner_id"`   // Audio message owner ID
	Waveform  []int  `json:"waveform"`
}

// MessagesChat type represents `messages_chat` API object
type MessagesChat struct {
	AdminId      int                       `json:"admin_id"`  // Chat creator ID
	Id           int                       `json:"id"`        // Chat ID
	Kicked       *BaseBoolInt              `json:"kicked"`    // Shows that user has been kicked from the chat
	Left         *BaseBoolInt              `json:"left"`      // Shows that user has been left the chat
	Photo100     string                    `json:"photo_100"` // URL of the preview image with 100 px in width
	Photo200     string                    `json:"photo_200"` // URL of the preview image with 200 px in width
	Photo50      string                    `json:"photo_50"`  // URL of the preview image with 50 px in width
	PushSettings *MessagesChatPushSettings `json:"push_settings"`
	Title        string                    `json:"title"` // Chat title
	Type         string                    `json:"type"`  // Chat type
	Users        []int                     `json:"users"`
}

// MessagesChatFull type represents `messages_chat_full` API object
type MessagesChatFull struct {
	AdminId      int                         `json:"admin_id"`  // Chat creator ID
	Id           int                         `json:"id"`        // Chat ID
	Kicked       *BaseBoolInt                `json:"kicked"`    // Shows that user has been kicked from the chat
	Left         *BaseBoolInt                `json:"left"`      // Shows that user has been left the chat
	Photo100     string                      `json:"photo_100"` // URL of the preview image with 100 px in width
	Photo200     string                      `json:"photo_200"` // URL of the preview image with 200 px in width
	Photo50      string                      `json:"photo_50"`  // URL of the preview image with 50 px in width
	PushSettings *MessagesChatPushSettings   `json:"push_settings"`
	Title        string                      `json:"title"` // Chat title
	Type         string                      `json:"type"`  // Chat type
	Users        []*MessagesUserXtrInvitedBy `json:"users"`
}

// MessagesChatPushSettings type represents `messages_chat_push_settings` API object
type MessagesChatPushSettings struct {
	DisabledUntil int          `json:"disabled_until"` // Time until that notifications are disabled
	Sound         *BaseBoolInt `json:"sound"`          // Information whether the sound is on
}

// MessagesChatRestrictions type represents `messages_chat_restrictions` API object
type MessagesChatRestrictions struct {
	AdminsPromoteUsers bool `json:"admins_promote_users"`  // Only admins can promote users to admins
	OnlyAdminsEditInfo bool `json:"only_admins_edit_info"` // Only admins can change chat info
	OnlyAdminsEditPin  bool `json:"only_admins_edit_pin"`  // Only admins can edit pinned message
	OnlyAdminsInvite   bool `json:"only_admins_invite"`    // Only admins can invite users to this chat
	OnlyAdminsKick     bool `json:"only_admins_kick"`      // Only admins can kick users from this chat
}

// MessagesConversation type represents `messages_conversation` API object
type MessagesConversation struct {
	CurrentKeyboard    *MessagesKeyboard         `json:"current_keyboard"`
	Important          bool                      `json:"important"`
	InRead             int                       `json:"in_read"`         // Last message user have read
	LastMessageId      int                       `json:"last_message_id"` // ID of the last message in conversation
	Mentions           []int                     `json:"mentions"`        // Ids of messages with mentions
	MessageRequest     string                    `json:"message_request"`
	OutRead            int                       `json:"out_read"` // Last outcoming message have been read by the opponent
	Peer               *MessagesConversationPeer `json:"peer"`
	SpecialServiceType string                    `json:"special_service_type"`
	Unanswered         bool                      `json:"unanswered"`
	UnreadCount        int                       `json:"unread_count"` // Unread messages number
}

// MessagesConversationMember type represents `messages_conversation_member` API object
type MessagesConversationMember struct {
	CanKick          bool `json:"can_kick"` // Is it possible for user to kick this member
	InvitedBy        int  `json:"invited_by"`
	IsAdmin          bool `json:"is_admin"`
	IsMessageRequest bool `json:"is_message_request"`
	IsOwner          bool `json:"is_owner"`
	JoinDate         int  `json:"join_date"`
	MemberId         int  `json:"member_id"`
	RequestDate      int  `json:"request_date"` // Message request date
}

// MessagesConversationPeer type represents `messages_conversation_peer` API object
type MessagesConversationPeer struct {
	Id      int                           `json:"id"`
	LocalId int                           `json:"local_id"`
	Type    *MessagesConversationPeerType `json:"type"`
}

// MessagesConversationPeerType type represents `messages_conversation_peer_type` API object
type MessagesConversationPeerType string // Peer type

// MessagesConversationWithMessage type represents `messages_conversation_with_message` API object
type MessagesConversationWithMessage struct {
	Conversation *MessagesConversation `json:"conversation"`
	LastMessage  *MessagesMessage      `json:"last_message"`
}

// MessagesForeignMessage type represents `messages_foreign_message` API object
type MessagesForeignMessage struct {
	Attachments           []*MessagesMessageAttachment `json:"attachments"`
	ConversationMessageId int                          `json:"conversation_message_id"` // Conversation message ID
	Date                  int                          `json:"date"`                    // Date when the message was created
	FromId                int                          `json:"from_id"`                 // Message author's ID
	FwdMessages           []*MessagesForeignMessage    `json:"fwd_messages"`
	Geo                   *BaseGeo                     `json:"geo"`
	Id                    int                          `json:"id"`      // Message ID
	PeerId                int                          `json:"peer_id"` // Peer ID
	ReplyMessage          *MessagesForeignMessage      `json:"reply_message"`
	Text                  string                       `json:"text"`        // Message text
	UpdateTime            int                          `json:"update_time"` // Date when the message has been updated in Unixtime
}

// MessagesGraffiti type represents `messages_graffiti` API object
type MessagesGraffiti struct {
	AccessKey string `json:"access_key"` // Access key for graffiti
	Height    int    `json:"height"`     // Graffiti height
	Id        int    `json:"id"`         // Graffiti ID
	OwnerId   int    `json:"owner_id"`   // Graffiti owner ID
	Url       string `json:"url"`        // Graffiti URL
	Width     int    `json:"width"`      // Graffiti width
}

// MessagesHistoryAttachment type represents `messages_history_attachment` API object
type MessagesHistoryAttachment struct {
	Attachment *MessagesHistoryMessageAttachment `json:"attachment"`
	FromId     int                               `json:"from_id"`    // Message author's ID
	MessageId  int                               `json:"message_id"` // Message ID
}

// MessagesHistoryMessageAttachment type represents `messages_history_message_attachment` API object
type MessagesHistoryMessageAttachment struct {
	Audio        *AudioAudio                           `json:"audio"`
	AudioMessage *MessagesAudioMessage                 `json:"audio_message"`
	Doc          *DocsDoc                              `json:"doc"`
	Graffiti     *MessagesGraffiti                     `json:"graffiti"`
	Link         *BaseLink                             `json:"link"`
	Market       *BaseLink                             `json:"market"`
	Photo        *PhotosPhoto                          `json:"photo"`
	Share        *BaseLink                             `json:"share"`
	Type         *MessagesHistoryMessageAttachmentType `json:"type"`
	Video        *VideoVideo                           `json:"video"`
	Wall         *BaseLink                             `json:"wall"`
}

// MessagesHistoryMessageAttachmentType type represents `messages_history_message_attachment_type` API object
type MessagesHistoryMessageAttachmentType string // Attachments type

// MessagesKeyboard type represents `messages_keyboard` API object
type MessagesKeyboard struct {
	AuthorId int                       `json:"author_id"` // Community or bot, which set this keyboard
	Buttons  []*MessagesKeyboardButton `json:"buttons"`
	Inline   bool                      `json:"inline"`
	OneTime  bool                      `json:"one_time"` // Should this keyboard disappear on first use
}

// MessagesKeyboardButton type represents `messages_keyboard_button` API object
type MessagesKeyboardButton struct {
	Action *MessagesKeyboardButtonAction `json:"action"`
	Color  string                        `json:"color"` // Button color
}

// MessagesKeyboardButtonAction type represents `messages_keyboard_button_action` API object
type MessagesKeyboardButtonAction struct {
	AppId   int    `json:"app_id"`   // Fragment value in app link like vk.com/app{app_id}_-654321#hash
	Hash    string `json:"hash"`     // Fragment value in app link like vk.com/app123456_-654321#{hash}
	Label   string `json:"label"`    // Label for button
	OwnerId int    `json:"owner_id"` // Fragment value in app link like vk.com/app123456_{owner_id}#hash
	Payload string `json:"payload"`  // Additional data sent along with message for developer convenience
	Type    string `json:"type"`     // Button type
}

// MessagesLastActivity type represents `messages_last_activity` API object
type MessagesLastActivity struct {
	Online *BaseBoolInt `json:"online"` // Information whether user is online
	Time   int          `json:"time"`   // Time when user was online in Unixtime
}

// MessagesLongpollMessages type represents `messages_longpoll_messages` API object
type MessagesLongpollMessages struct {
	Count int                `json:"count"` // Total number
	Items []*MessagesMessage `json:"items"`
}

// MessagesLongpollParams type represents `messages_longpoll_params` API object
type MessagesLongpollParams struct {
	Key    string `json:"key"`    // Key
	Pts    int    `json:"pts"`    // Persistent timestamp
	Server string `json:"server"` // Server URL
	Ts     int    `json:"ts"`     // Timestamp
}

// MessagesMessage type represents `messages_message` API object
type MessagesMessage struct {
	Action                *MessagesMessageAction       `json:"action"`
	AdminAuthorId         int                          `json:"admin_author_id"` // Only for messages from community. Contains user ID of community admin, who sent this message.
	Attachments           []*MessagesMessageAttachment `json:"attachments"`
	ConversationMessageId int                          `json:"conversation_message_id"` // Unique auto-incremented number for all messages with this peer
	Date                  int                          `json:"date"`                    // Date when the message has been sent in Unixtime
	Deleted               *BaseBoolInt                 `json:"deleted"`                 // Is it an deleted message
	FromId                int                          `json:"from_id"`                 // Message author's ID
	FwdMessages           []*MessagesForeignMessage    `json:"fwd_messages"`            // Forwarded messages
	Geo                   *BaseGeo                     `json:"geo"`
	Id                    int                          `json:"id"`        // Message ID
	Important             bool                         `json:"important"` // Is it an important message
	IsHidden              bool                         `json:"is_hidden"`
	Keyboard              *MessagesKeyboard            `json:"keyboard"`
	MembersCount          int                          `json:"members_count"` // Members number
	Out                   *BaseBoolInt                 `json:"out"`           // Information whether the message is outcoming
	Payload               string                       `json:"payload"`
	PeerId                int                          `json:"peer_id"`   // Peer ID
	RandomId              int                          `json:"random_id"` // ID used for sending messages. It returned only for outgoing messages
	Ref                   string                       `json:"ref"`
	RefSource             string                       `json:"ref_source"`
	ReplyMessage          *MessagesForeignMessage      `json:"reply_message"`
	Text                  string                       `json:"text"`        // Message text
	UpdateTime            int                          `json:"update_time"` // Date when the message has been updated in Unixtime
}

// MessagesMessageAction type represents `messages_message_action` API object
type MessagesMessageAction struct {
	ConversationMessageId int                          `json:"conversation_message_id"` // Message ID
	Email                 string                       `json:"email"`                   // Email address for chat_invite_user or chat_kick_user actions
	MemberId              int                          `json:"member_id"`               // User or email peer ID
	Message               string                       `json:"message"`                 // Message body of related message
	Photo                 *MessagesMessageActionPhoto  `json:"photo"`
	Text                  string                       `json:"text"` // New chat title for chat_create and chat_title_update actions
	Type                  *MessagesMessageActionStatus `json:"type"`
}

// MessagesMessageActionPhoto type represents `messages_message_action_photo` API object
type MessagesMessageActionPhoto struct {
	Photo100 string `json:"photo_100"` // URL of the preview image with 100px in width
	Photo200 string `json:"photo_200"` // URL of the preview image with 200px in width
	Photo50  string `json:"photo_50"`  // URL of the preview image with 50px in width
}

// MessagesMessageActionStatus type represents `messages_message_action_status` API object
type MessagesMessageActionStatus string // Action status

// MessagesMessageAttachment type represents `messages_message_attachment` API object
type MessagesMessageAttachment struct {
	Audio             *AudioAudio                    `json:"audio"`
	AudioMessage      *MessagesAudioMessage          `json:"audio_message"`
	Doc               *DocsDoc                       `json:"doc"`
	Gift              *GiftsLayout                   `json:"gift"`
	Graffiti          *MessagesGraffiti              `json:"graffiti"`
	Link              *BaseLink                      `json:"link"`
	Market            *MarketMarketItem              `json:"market"`
	MarketMarketAlbum *MarketMarketAlbum             `json:"market_market_album"`
	Photo             *PhotosPhoto                   `json:"photo"`
	Sticker           *BaseSticker                   `json:"sticker"`
	Type              *MessagesMessageAttachmentType `json:"type"`
	Video             *VideoVideo                    `json:"video"`
	Wall              *WallWallpostFull              `json:"wall"`
	WallReply         *WallWallComment               `json:"wall_reply"`
}

// MessagesMessageAttachmentType type represents `messages_message_attachment_type` API object
type MessagesMessageAttachmentType string // Attachment type

// MessagesPinnedMessage type represents `messages_pinned_message` API object
type MessagesPinnedMessage struct {
	Attachments           []*MessagesMessageAttachment `json:"attachments"`
	ConversationMessageId int                          `json:"conversation_message_id"` // Unique auto-incremented number for all messages with this peer
	Date                  int                          `json:"date"`                    // Date when the message has been sent in Unixtime
	FromId                int                          `json:"from_id"`                 // Message author's ID
	FwdMessages           []*MessagesForeignMessage    `json:"fwd_messages"`            // Forwarded messages
	Geo                   *BaseGeo                     `json:"geo"`
	Id                    int                          `json:"id"` // Message ID
	Keyboard              *MessagesKeyboard            `json:"keyboard"`
	PeerId                int                          `json:"peer_id"` // Peer ID
	ReplyMessage          *MessagesForeignMessage      `json:"reply_message"`
	Text                  string                       `json:"text"` // Message text
}

// MessagesUserXtrInvitedBy type represents `messages_user_xtr_invited_by` API object
type MessagesUserXtrInvitedBy struct {
	Usersuserxtrtype *UsersUserXtrType `json:"Usersuserxtrtype"`
	InvitedBy        int               `json:"invited_by"` // ID of the inviter
}
