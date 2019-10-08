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
// `account` group of objects
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

// AccountPushSettings type represents `account_push_settings` API object
type AccountPushSettings struct {
	Conversations AccountPushConversations `json:"conversations"`
	Disabled      BaseBoolInt              `json:"disabled"`       // Information whether notifications are disabled
	DisabledUntil int                      `json:"disabled_until"` // Time until that notifications are disabled in Unixtime
	Settings      AccountPushParams        `json:"settings"`
}

// AccountPushConversations type represents `account_push_conversations` API object
type AccountPushConversations struct {
	Count int                            `json:"count"` // Items count
	Items []AccountPushConversationsItem `json:"items"`
}

// AccountNameRequest type represents `account_name_request` API object
type AccountNameRequest struct {
	FirstName string                   `json:"first_name"` // First name in request
	Id        int                      `json:"id"`         // Request ID needed to cancel the request
	Lang      string                   `json:"lang"`
	LastName  string                   `json:"last_name"` // Last name in request
	Status    AccountNameRequestStatus `json:"status"`
}

// AccountNameRequestStatus type represents `account_name_request_status` API object
type AccountNameRequestStatus string // Request status

// AccountPushParams type represents `account_push_params` API object
type AccountPushParams struct {
	AppRequest     []AccountPushParamsOnoff    `json:"app_request"`
	Birthday       []AccountPushParamsOnoff    `json:"birthday"`
	Chat           []AccountPushParamsMode     `json:"chat"`
	Comment        []AccountPushParamsSettings `json:"comment"`
	EventSoon      []AccountPushParamsOnoff    `json:"event_soon"`
	Friend         []AccountPushParamsOnoff    `json:"friend"`
	FriendAccepted []AccountPushParamsOnoff    `json:"friend_accepted"`
	FriendFound    []AccountPushParamsOnoff    `json:"friend_found"`
	GroupAccepted  []AccountPushParamsOnoff    `json:"group_accepted"`
	GroupInvite    []AccountPushParamsOnoff    `json:"group_invite"`
	Like           []AccountPushParamsSettings `json:"like"`
	Mention        []AccountPushParamsSettings `json:"mention"`
	Msg            []AccountPushParamsMode     `json:"msg"`
	NewPost        []AccountPushParamsOnoff    `json:"new_post"`
	Reply          []AccountPushParamsOnoff    `json:"reply"`
	Repost         []AccountPushParamsSettings `json:"repost"`
	SdkOpen        []AccountPushParamsOnoff    `json:"sdk_open"`
	WallPost       []AccountPushParamsOnoff    `json:"wall_post"`
	WallPublish    []AccountPushParamsOnoff    `json:"wall_publish"`
}

// AccountPushParamsMode type represents `account_push_params_mode` API object
type AccountPushParamsMode string // Settings parameters

// AccountPushParamsSettings type represents `account_push_params_settings` API object
type AccountPushParamsSettings string // Settings parameters

// AccountAccountCounters type represents `account_account_counters` API object
type AccountAccountCounters struct {
	AppRequests        int `json:"app_requests"`        // New app requests number
	Events             int `json:"events"`              // New events number
	Friends            int `json:"friends"`             // New friends requests number
	FriendsSuggestions int `json:"friends_suggestions"` // New friends suggestions number
	Gifts              int `json:"gifts"`               // New gifts number
	Groups             int `json:"groups"`              // New groups number
	Messages           int `json:"messages"`            // New messages number
	Notifications      int `json:"notifications"`       // New notifications number
	Photos             int `json:"photos"`              // New photo tags number
	Videos             int `json:"videos"`              // New video tags number
}

// AccountUserSettingsInterests type represents `account_user_settings_interests` API object
type AccountUserSettingsInterests struct {
	About      AccountUserSettingsInterest `json:"about"`
	Activities AccountUserSettingsInterest `json:"activities"`
	Books      AccountUserSettingsInterest `json:"books"`
	Games      AccountUserSettingsInterest `json:"games"`
	Interests  AccountUserSettingsInterest `json:"interests"`
	Movies     AccountUserSettingsInterest `json:"movies"`
	Music      AccountUserSettingsInterest `json:"music"`
	Quotes     AccountUserSettingsInterest `json:"quotes"`
	Tv         AccountUserSettingsInterest `json:"tv"`
}

// AccountPushConversationsItem type represents `account_push_conversations_item` API object
type AccountPushConversationsItem struct {
	DisabledUntil int         `json:"disabled_until"` // Time until that notifications are disabled in seconds
	PeerId        int         `json:"peer_id"`        // Peer ID
	Sound         BaseBoolInt `json:"sound"`          // Information whether the sound are enabled
}

// AccountUserSettingsInterest type represents `account_user_settings_interest` API object
type AccountUserSettingsInterest struct {
	Title string `json:"title"`
	Value string `json:"value"`
}

// AccountUserSettings type represents `account_user_settings` API object
type AccountUserSettings struct {
	UsersUserMin         UsersUserMin         `json:"UsersUserMin"`
	UsersUserSettingsXtr UsersUserSettingsXtr `json:"UsersUserSettingsXtr"`
}

// AccountOffer type represents `account_offer` API object
type AccountOffer struct {
	Description      string `json:"description"`       // Offer description
	Id               int    `json:"id"`                // Offer ID
	Img              string `json:"img"`               // URL of the preview image
	Instruction      string `json:"instruction"`       // Instruction how to process the offer
	InstructionHtml  string `json:"instruction_html"`  // Instruction how to process the offer (HTML format)
	Price            int    `json:"price"`             // Offer price
	ShortDescription string `json:"short_description"` // Offer short description
	Tag              string `json:"tag"`               // Offer tag
	Title            string `json:"title"`             // Offer title
}

// AccountPushParamsOnoff type represents `account_push_params_onoff` API object
type AccountPushParamsOnoff string // Settings parameters

// AccountInfo type represents `account_info` API object
type AccountInfo struct {
	TwofaRequired   BaseBoolInt `json:"2fa_required"`      // Two factor authentication is enabled
	Country         string      `json:"country"`           // Country code
	HttpsRequired   BaseBoolInt `json:"https_required"`    // Information whether HTTPS-only is enabled
	Intro           BaseBoolInt `json:"intro"`             // Information whether user has been processed intro
	Lang            int         `json:"lang"`              // Language ID
	NoWallReplies   BaseBoolInt `json:"no_wall_replies"`   // Information whether wall comments should be hidden
	OwnPostsDefault BaseBoolInt `json:"own_posts_default"` // Information whether only owners posts should be shown
}
