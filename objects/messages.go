package objects

/////////////////////////////////////////////////////////////
// Messages related API objects	                           //
/////////////////////////////////////////////////////////////

// Chat represents `messages_chat` API object
type MessagesChat struct {
	AdminID      int                      `json:"admin_id"`  // Chat creator ID
	ID           int                      `json:"id"`        // Chat ID
	Kicked       BaseBoolInt              `json:"kicked"`    // Shows that user has been kicked from the chat
	Left         BaseBoolInt              `json:"left"`      // Shows that user has been left the chat
	Photo100     string                   `json:"photo_100"` // URL of the preview image with 100 px in width
	Photo200     string                   `json:"photo_200"` // URL of the preview image with 200 px in width
	Photo50      string                   `json:"photo_50"`  // URL of the preview image with 50 px in width
	PushSettings MessagesChatPushSettings `json:"push_settings"`
	Title        string                   `json:"title"` // Chat title
	Type         string                   `json:"type"`  // Chat type
	Users        []int                    `json:"users"` // User IDs
}

// ChatPreview represents `messages_chat_preview` API object
type MessagesChatPreview struct {
	AdminID      int   `json:"admin_id"`
	Joined       bool  `json:"joined"`
	LocalID      int   `json:"local_id"`
	Members      []int `json:"members"`
	MembersCount int   `json:"members_count"`
}

// ChatFull represents `messages_chat_full` API object
type MessagesChatFull struct {
	AdminID      int                        `json:"admin_id"`  // Chat creator ID
	ID           int                        `json:"id"`        // Chat ID
	Kicked       BaseBoolInt                `json:"kicked"`    // Shows that user has been kicked from the chat
	Left         BaseBoolInt                `json:"left"`      // Shows that user has been left the chat
	Photo100     string                     `json:"photo_100"` // URL of the preview image with 100 px in width
	Photo200     string                     `json:"photo_200"` // URL of the preview image with 200 px in width
	Photo50      string                     `json:"photo_50"`  // URL of the preview image with 50 px in width
	PushSettings MessagesChatPushSettings   `json:"push_settings"`
	Title        string                     `json:"title"` // Chat title
	Type         string                     `json:"type"`  // Chat type
	Users        []MessagesUserXtrInvitedBy `json:"users"`
}

// ChatPushSettings
type MessagesChatPushSettings struct {
	DisabledUntil int         `json:"disabled_until"` // Time until that notifications are disabled
	Sound         BaseBoolInt `json:"sound"`          // Information whether the sound is on
}

// ChatRestrictions represents `messages_chat_restrictions` API object
type MessagesChatRestrictions struct {
	AdminsPromoteUsers bool `json:"admins_promote_users"`  // Only admins can promote users to admins
	AdminsEditInfo     bool `json:"only_admins_edit_info"` // Only admins can change chat info
	AdminsEditPin      bool `json:"only_admins_edit_pin"`  // Only admins can edit pinned message
	AdminsInvite       bool `json:"only_admins_invite"`    // Only admins can invite users to this chat
	AdminsKick         bool `json:"only_admins_kick"`      // Only admins can kick users from this chat
}

// AudioMessage represents `messages_audio_message` API object
type MessagesAudioMessage struct {
	AccessKey string `json:"access_key"` // Access key for audio message
	Duration  int    `json:"duration"`   // Audio message duration in seconds
	ID        int    `json:"id"`         // Audio message ID
	LinkMp3   string `json:"link_mp3"`   // MP3 file URL
	LinkOgg   string `json:"link_ogg"`   // OGG file URL
	OwnerID   int    `json:"owner_id"`   // Audio message owner ID
	Waveform  []int  `json:"waveform"`   // Sound visualisation
}

// Conversation represents `messages_conversation` API object
type MessagesConversation struct {
	CurrentKeyboard MessagesKeyboard         `json:"current_keyboard"`
	Important       bool                     `json:"important"`
	InRead          int                      `json:"in_read"`         // Last message user have read
	LastMsgID       int                      `json:"last_message_id"` // ID of the last message in conversation
	Mentions        []int                    `json:"mentions"`        // Ids of messages with mentions
	MessageRequest  string                   `json:"message_request"`
	OutRead         int                      `json:"out_read"` // Last outgoing message have been read by the opponent
	Peer            MessagesConversationPeer `json:"peer"`
	Unanswered      int                      `json:"unanswered"`
	UnreadCount     int                      `json:"unread_count"` // Unread messages number
}

// ConversationMember represents `messages_conversation_member` API object
type MessagesConversationMember struct {
	CanKick   bool `json:"can_kick"` // Is it possible for user to kick this member
	InvitedBy int  `json:"invited_by"`
	IsAdmin   bool `json:"is_admin"`
	IsOwner   bool `json:"is_owner"`
	JoinDate  int  `json:"join_date"`
	MemberID  int  `json:"member_id"`
}

// ConversationPeer represents `messages_conversation_peer` API object
type MessagesConversationPeer struct {
	ID      int                          `json:"id"`
	LocalID int                          `json:"local_id"`
	Type    MessagesConversationPeerType `json:"type"`
}

// ConversationPeerType represents `messages_conversation_peer_type` API object
type MessagesConversationPeerType string

func (c *MessagesConversationPeerType) MarshalJSON() ([]byte, error) {
	return GetStringFromRange(string(*c), "chat", "email", "user", "group")
}

func (c *MessagesConversationPeerType) String() string {
	return string(*c)
}

// ConversationWithMessage represents `messages_conversation_with_message` API object
type MessagesConversationWithMessage struct {
	Conversation MessagesConversation `json:"conversation"`
	LastMessage  MessagesMessage      `json:"last_message"`
}

// ForeignMessage represents `messages_foreign_message` API object
type MessagesForeignMessage struct {
	Attachments       []MessagesAttachment      `json:"attachments"`
	ConversationMsgID int                       `json:"conversation_message_id"` // Conversation message ID
	Date              int                       `json:"date"`                    // Date when the message was created
	FromID            int                       `json:"from_id"`                 // Message author's ID
	FwdMessages       []*MessagesForeignMessage `json:"fwd_messages"`
	Geo               BaseGeo                   `json:"geo"`
	ID                int                       `json:"id"`      // Message ID
	PeerID            int                       `json:"peer_id"` // Peer ID
	ReplyMessage      *MessagesForeignMessage   `json:"reply_message"`
	Text              string                    `json:"text"`        // Message text
	UpdateTime        int                       `json:"update_time"` // Date when the message has been updated in Unixtime
}

// Graffiti represents `messages_graffiti` API object
type MessagesGraffiti struct {
	AccessKey string `json:"access_key"` // Access key for graffiti
	Height    int    `json:"height"`     // Graffiti height
	ID        int    `json:"id"`         // Graffiti ID
	OwnerID   int    `json:"owner_id"`   // Graffiti owner ID
	Url       string `json:"url"`        // Graffiti owner ID
	Width     int    `json:"width"`      // Graffiti width
}

// HistoryAttachment represents `messages_history_attachment` API object
type MessagesHistoryAttachment struct {
	MsgID      int                              `json:"message_id"`
	Attachment MessagesHistoryMessageAttachment `json:"attachment"`
}

// HistoryMessageAttachment represents `messages_history_message_attachment` API object
type MessagesHistoryMessageAttachment struct {
	Audio    AudioAudioFull                       `json:"audio"`
	AudioMsg MessagesAudioMessage                 `json:"audio_message"`
	Doc      DocsDoc                              `json:"doc"`
	Graffiti MessagesGraffiti                     `json:"graffiti"`
	Link     BaseLink                             `json:"link"`
	Market   BaseLink                             `json:"market"`
	Photo    PhotosPhoto                          `json:"photo"`
	Share    BaseLink                             `json:"share"`
	Type     MessagesHistoryMessageAttachmentType `json:"type"`
	Video    VideosVideo                          `json:"video"`
	Wall     BaseLink                             `json:"wall"`
}

// HistoryMessageAttachmentType represents `messages_history_message_attachment_type` API object
type MessagesHistoryMessageAttachmentType string

func (h *MessagesHistoryMessageAttachmentType) MarshalJSON() ([]byte, error) {
	return GetStringFromRange(string(*h), "photo", "video", "audio", "doc", "link", "market",
		"wall", "share", "graffiti", "audio_message")
}

func (h *MessagesHistoryMessageAttachmentType) String() string {
	return string(*h)
}

// Keyboard represents `messages_keyboard` API object
type MessagesKeyboard struct {
	AuthorID int                        `json:"author_id"` // Community or bot, which set this keyboard
	Buttons  [][]MessagesKeyboardButton `json:"buttons"`
	OneTime  bool                       `json:"one_time"` // Should this keyboard disappear on first use
}

// KeyboardButton represents `messages_keyboard_button` API object
type MessagesKeyboardButton struct {
	Action MessagesKeyboardButtonAction `json:"action"`
	Color  string                       `json:"color"` // Button color
}

// KeyboardButtonAction represents `messages_keyboard_button_action` API object.
// Contains description of the action, that should be performed on button click.
type MessagesKeyboardButtonAction struct {
	AppID   int    `json:"app_id"`   // Fragment value in app link like vk.com/app{app_id}_-654321#hash
	Hash    string `json:"hash"`     // Fragment value in app link like vk.com/app123456_-654321#{hash}
	Label   string `json:"label"`    // Label for button
	OwnerID int    `json:"owner_id"` // Fragment value in app link like vk.com/app123456_{owner_id}#hash
	Payload string `json:"payload"`  // Additional data sent along with message for developer convenience
	Type    string `json:"type"`     // Button type
}

// LastActivity represents `messages_last_activity` API object
type MessagesLastActivity struct {
	Online BaseBoolInt `json:"online"` // Information whether user is online
	Time   int         `json:"time"`   // Time when user was online in Unixtime
}

// LongpollMessages represents `messages_longpoll_messages` API object
type MessagesLongpollMessages struct {
	Count int               `json:"count"` // Total number
	Items []MessagesMessage `json:"items"`
}

// LongpollParams represents `messages_longpoll_params` API object
type MessagesLongpollParams struct {
	Key    string `json:"key"`    // Key
	Pts    int    `json:"pts"`    // Persistent timestamp
	Ts     int    `json:"ts"`     // Timestamp
	Server string `json:"server"` // Server URL
}

// Message represents `messages_message` API object
type MessagesMessage struct {
	Action            MessagesMessageAction    `json:"action"`
	AdminAuthorID     int                      `json:"admin_author_id"` // Only for messages from community. Contains user ID of community admin, who sent this message.
	Attachments       []MessagesAttachment     `json:"attachments"`
	ConversationMsgID int                      `json:"conversation_message_id"` // Unique auto-incremented number for all messages with this peer
	Date              int                      `json:"date"`                    // Date when the message has been sent in Unixtime
	Deleted           BaseBoolInt              `json:"deleted"`                 // Is it an deleted message
	FromID            int                      `json:"from_id"`                 // Message author's ID
	FwdMessages       []MessagesForeignMessage `json:"fwd_messages"`            // Forwarded messages
	Geo               BaseGeo                  `json:"geo"`
	ID                int                      `json:"id"`        // Message ID
	Important         bool                     `json:"important"` // Is it an important message
	Hidden            bool                     `json:"is_hidden"`
	Keyboard          MessagesKeyboard         `json:"keyboard"`
	MembersCount      int                      `json:"members_count"` // Members number
	Out               BaseBoolInt              `json:"out"`           // Information whether the message is outgoing
	Payload           string                   `json:"payload"`
	PeerID            int                      `json:"peer_id"`   // Peer ID
	RandomID          int                      `json:"random_id"` // ID used for sending messages. It returned only for outgoing messages
	ReplyMessage      MessagesForeignMessage   `json:"reply_message"`
	Text              string                   `json:"text"`        // Message text
	UpdateTime        int                      `json:"update_time"` // Date when the message has been updated in Unixtime
}

// MessageAction represents `messages_message_action` APi object
type MessagesMessageAction struct {
	ConversationMsgID int                         `json:"conversation_message_id"` // Message ID
	Email             string                      `json:"email"`                   // Email address for chat_invite_user or chat_kick_user actions
	MemberID          int                         `json:"member_id"`               // User or email peer ID
	Message           string                      `json:"message"`                 // Message body of related message
	Photo             MessagesMessageActionPhoto  `json:"photo"`
	Text              string                      `json:"text"` // New chat title for chat_create and chat_title_update actions
	Type              MessagesMessageActionStatus `json:"type"`
}

// MessageActionPhoto represents `messages_message_action_photo` API object
type MessagesMessageActionPhoto struct {
	Photo100 string `json:"photo_100"` // URL of the preview image with 100px in width
	Photo200 string `json:"photo_200"` // URL of the preview image with 200px in width
	Photo50  string `json:"photo_50"`  // URL of the preview image with 5s0px in width
}

// MessageActionStatus represents `messages_message_action_status` API object
type MessagesMessageActionStatus string

func (m *MessagesMessageActionStatus) MarshalJSON() ([]byte, error) {
	return GetStringFromRange(string(*m), "chat_photo_update", "chat_photo_remove", "chat_create",
		"chat_title_update", "chat_invite_user", "chat_kick_user", "chat_pin_message", "chat_unpin_message",
		"chat_invite_user_by_link")
}

func (m *MessagesMessageActionStatus) String() string {
	return string(*m)
}

// Attachment represents `messages_message_attachment` API object
type MessagesAttachment struct {
	Audio        AudioAudioFull         `json:"audio"`
	AudioMessage MessagesAudioMessage   `json:"audio_message"`
	Doc          DocsDoc                `json:"doc"`
	Gift         GiftsLayout            `json:"gift"`
	Graffiti     MessagesGraffiti       `json:"graffiti"`
	Link         BaseLink               `json:"link"`
	Market       MarketItem             `json:"market"`
	MarketAlbum  MarketAlbum            `json:"market_market_album"`
	Photo        PhotosPhoto            `json:"photo"`
	Sticker      BaseSticker            `json:"sticker"`
	Type         MessagesAttachmentType `json:"type"`
	Video        VideosVideo            `json:"video"`
	Wall         WallWallpostFull       `json:"wall"`
	WallReply    WallWallComment        `json:"wall_reply"`
}

// AttachmentType represents `messages_message_attachment_type` API object
type MessagesAttachmentType string

func (a *MessagesAttachmentType) MarshalJSON() ([]byte, error) {
	return GetStringFromRange(string(*a), "photo", "audio", "video", "doc", "link", "market",
		"market_album", "gift", "sticker", "wall", "wall_reply", "graffiti", "audio_message")
}

func (a *MessagesAttachmentType) String() string {
	return string(*a)
}

// PinnedMsg represents `messages_pinned_message` API object
type MessagesPinnedMsg struct {
	Attachments       []MessagesAttachment   `json:"attachments"`
	ConversationMsgID int                    `json:"conversation_message_id"` // Unique auto-incremented number for all messages with this peer
	Date              int                    `json:"date"`                    // Date when the message has been sent in Unixtime
	FromID            int                    `json:"from_id"`                 // Message author's ID
	FwdMessages       MessagesForeignMessage `json:"fwd_messages"`
	Geo               BaseGeo                `json:"geo"`
	ID                int                    `json:"id"`      // Message ID
	PeerID            int                    `json:"peer_id"` // Peer ID
	ReplyMsg          MessagesForeignMessage `json:"reply_message"`
	Text              string                 `json:"text"` // Message text
}

// UserXtrInvitedBy represents `messages_user_xtr_invited_by` API object
type MessagesUserXtrInvitedBy struct {
	*UsersUserXtrType
	InvitedBy int `json:"invited_by"` // ID of the inviter
}
