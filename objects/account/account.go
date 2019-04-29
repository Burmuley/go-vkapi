package account

import (
	"fmt"
	"gitlab.com/Burmuley/go-vkapi/objects/base"
	"gitlab.com/Burmuley/go-vkapi/objects/users"
)

//###########################################################
// Account related API objects
//###########################################################

// Counters represents `account_account_counters` API object
type Counters struct {
	AppRequests        int `json:"app_requests"`
	Events             int `json:"events"`
	Friends            int `json:"friends"`
	FriendsSuggestions int `json:"friends_suggestions"`
	Gifts              int `json:"gifts"`
	Groups             int `json:"groups"`
	Messages           int `json:"messages"`
	Notifications      int `json:"notifications"`
	Photos             int `json:"photos"`
	Videos             int `json:"videos"`
}

// Info type represents 'account_info' API objects
type Info struct {
	TwoFARequired   base.BoolInt `json:"2fa_required"`
	Country         string       `json:"country"`
	HTTPSRequired   base.BoolInt `json:"https_required"`
	Intro           base.BoolInt `json:"intro"`
	Lang            int          `json:"lang"`
	NoWallReplies   base.BoolInt `json:"no_wall_replies"`
	OwnPostsDefault base.BoolInt `json:"own_posts_default"`
}

// NameRequest represents `account_name_request` API object
type NameRequest struct {
	ID        int               `json:"id"`
	FirstName string            `json:"first_name"`
	LastName  string            `json:"last_name"`
	Status    NameRequestStatus `json:"status"`
}

// NameRequestStatus represents `account_name_request_status` API object
type NameRequestStatus string

func (n *NameRequestStatus) GetName() string {
	return string(*n)
}

// Offer represents `account_offer` API object
type Offer struct {
	ID               int    `json:"id"`
	Description      string `json:"description"`
	Image            string `json:"img"`
	Instruction      string `json:"instruction"`
	InstructionHTML  string `json:"instruction_html"`
	Price            int    `json:"price"`
	ShortDescription string `json:"short_description"`
	Tag              string `json:"tag"`
	Title            string `json:"title"`
}

// OnOffOptions represents `account_onoff_options` API object
type OnOffOptions string

func (o *OnOffOptions) MarshalJSON() ([]byte, error) {
	if *o != "on" && *o != "off" {
		return []byte{}, fmt.Errorf("value is not in range ['on', 'off']")
	}

	return []byte(*o), nil
}

func (o *OnOffOptions) GetName() string {
	return string(*o)
}

// PushConversations represents `account_push_conversations` API object
type PushConversations struct {
	Count int                     `json:"count"`
	Items []PushConversationsItem `json:"items"`
}

// PushConversationsItem represents `account_push_conversations_item` API object
type PushConversationsItem struct {
	DisabledUntil int          `json:"disabled_until"`
	PeerID        int          `json:"peer_id"`
	SoundEnabled  base.BoolInt `json:"sound"`
}

// PushParams represents `account_push_params` API object
type PushParams struct {
	AppRequest     []OnOffOptions       `json:"app_request"`
	Birthday       []OnOffOptions       `json:"birthday"`
	Chat           []PushParamsMode     `json:"chat"`
	Comment        []PushParamsSettings `json:"comment"`
	EventSoon      []OnOffOptions       `json:"event_soon"`
	Friend         []OnOffOptions       `json:"friend"`
	FriendAccepted []OnOffOptions       `json:"friend_accepted"`
	FriendFound    []OnOffOptions       `json:"friend_found"`
	GroupAccepted  []OnOffOptions       `json:"group_accepted"`
	GroupInvite    []OnOffOptions       `json:"group_invite"`
	Like           []PushParamsSettings `json:"like"`
	Mention        []PushParamsSettings `json:"mention"`
	Message        []PushParamsMode     `json:"msg"`
	Post           []OnOffOptions       `json:"new_post"`
	PhotosTag      []PushParamsSettings `json:"photos_tag"`
	Reply          []OnOffOptions       `json:"reply"`
	Repost         []PushParamsSettings `json:"repost"`
	SdkOpen        []OnOffOptions       `json:"sdk_open"`
	WallPost       []OnOffOptions       `json:"wall_post"`
	WallPublish    []OnOffOptions       `json:"wall_publish"`
}

// PushParamsMode represents `account_push_params_mode` API object
type PushParamsMode string

func (p *PushParamsMode) MarshalJSON() ([]byte, error) {
	switch *p {
	case "on", "off", "no_sound", "no_text":
		return []byte(*p), nil
	}

	return []byte{}, fmt.Errorf("value is not in range ['on', 'off', 'no_sound', 'no_text']")
}

func (p *PushParamsMode) GetName() string {
	return string(*p)
}

// PushParamsSettings represents `account_push_params_settings` API object
type PushParamsSettings string

func (p *PushParamsSettings) MarshalJSON() ([]byte, error) {
	switch *p {
	case "on", "off", "fr_of_fr":
		return []byte(*p), nil
	}

	return []byte{}, fmt.Errorf("value is not in range ['on', 'off', 'fr_of_fr']")
}

func (p *PushParamsSettings) GetName() string {
	return string(*p)
}

// PushSettings represents `account_push_settings` API object
type PushSettings struct {
	Conversations PushConversations `json:"conversations"`
	Disabled      base.BoolInt      `json:"disabled"`
	DisabledUntil int               `json:"disabled_until"`
	Settings      PushParams        `json:"settings"`
}

// A UserSettings represents 'account_user_settings` API object
type UserSettings struct {
	Birthday         string          `json:"bdate"`
	BirthVisibility  int             `json:"bdate_visibility"`
	City             base.Object     `json:"city"`
	Country          base.Country    `json:"country"`
	FirstName        string          `json:"first_name"`
	Hometown         string          `json:"home_town"`
	LastName         string          `json:"last_name"`
	MaidenName       string          `json:"maiden_name"`
	NameRequest      NameRequest     `json:"name_request"`
	Phone            string          `json:"phone"`
	Relation         int             `json:"relation"`
	RelationPartner  users.UserMin   `json:"relation_partner"`
	RelationPending  int             `json:"relation_pending"`
	RelationRequests []users.UserMin `json:"relation_requests"`
	ScreenName       string          `json:"screen_name"`
	Sex              base.Sex        `json:"sex"`
	Status           string          `json:"status"`
}
