package objects

/////////////////////////////////////////////////////////////
// Account related API objects                             //
/////////////////////////////////////////////////////////////

// Counters represents `account_account_counters` API object
type AccountCounters struct {
	AppRequests            int `json:"app_requests"`
	SDKRequests            int `json:"sdk"`
	Events                 int `json:"events"`
	Friends                int `json:"friends"`
	FriendsSuggestions     int `json:"friends_suggestions"`
	FriendsRecommendations int `json:"friends_recommendations"`
	Gifts                  int `json:"gifts"`
	Groups                 int `json:"groups"`
	Messages               int `json:"messages"`
	Notifications          int `json:"notifications"`
	Photos                 int `json:"photos"`
	Videos                 int `json:"videos"`
	MenuBadge              int `json:"menu_discover_badge"`
}

// Info type represents 'account_info' API objects
type AccountInfo struct {
	TwoFARequired   BaseBoolInt `json:"2fa_required"`
	Country         string      `json:"country"`
	HTTPSRequired   BaseBoolInt `json:"https_required"`
	Intro           BaseBoolInt `json:"intro"`
	Lang            int         `json:"lang"`
	NoWallReplies   BaseBoolInt `json:"no_wall_replies"`
	OwnPostsDefault BaseBoolInt `json:"own_posts_default"`
}

// NameRequest represents `account_name_request` API object
type AccountNameRequest struct {
	ID        int                      `json:"id"`
	FirstName string                   `json:"first_name"`
	LastName  string                   `json:"last_name"`
	Status    AccountNameRequestStatus `json:"status"`
}

// NameRequestStatus represents `account_name_request_status` API object
type AccountNameRequestStatus string

func (n *AccountNameRequestStatus) String() string {
	return string(*n)
}

// Offer represents `account_offer` API object
type AccountOffer struct {
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
type AccountOnOffOptions string

func (o *AccountOnOffOptions) MarshalJSON() ([]byte, error) {
	return GetStringFromRange(string(*o), "on", "off")
}

func (o *AccountOnOffOptions) String() string {
	return string(*o)
}

// PushConversations represents `account_push_conversations` API object
type AccountPushConversations struct {
	Count int                            `json:"count"`
	Items []AccountPushConversationsItem `json:"items"`
}

// PushConversationsItem represents `account_push_conversations_item` API object
type AccountPushConversationsItem struct {
	DisabledUntil int         `json:"disabled_until"`
	PeerID        int         `json:"peer_id"`
	SoundEnabled  BaseBoolInt `json:"sound"`
}

// PushParams represents `account_push_params` API object
type AccountPushParams struct {
	AppRequest     []AccountOnOffOptions       `json:"app_request"`
	Birthday       []AccountOnOffOptions       `json:"birthday"`
	Chat           []AccountPushParamsMode     `json:"chat"`
	Comment        []AccountPushParamsSettings `json:"comment"`
	EventSoon      []AccountOnOffOptions       `json:"event_soon"`
	Friend         []AccountOnOffOptions       `json:"friend"`
	FriendAccepted []AccountOnOffOptions       `json:"friend_accepted"`
	FriendFound    []AccountOnOffOptions       `json:"friend_found"`
	GroupAccepted  []AccountOnOffOptions       `json:"group_accepted"`
	GroupInvite    []AccountOnOffOptions       `json:"group_invite"`
	Like           []AccountPushParamsSettings `json:"like"`
	Mention        []AccountPushParamsSettings `json:"mention"`
	Message        []AccountPushParamsMode     `json:"msg"`
	Post           []AccountOnOffOptions       `json:"new_post"`
	PhotosTag      []AccountPushParamsSettings `json:"photos_tag"`
	Reply          []AccountOnOffOptions       `json:"reply"`
	Repost         []AccountPushParamsSettings `json:"repost"`
	SdkOpen        []AccountOnOffOptions       `json:"sdk_open"`
	WallPost       []AccountOnOffOptions       `json:"wall_post"`
	WallPublish    []AccountOnOffOptions       `json:"wall_publish"`
}

// PushParamsMode represents `account_push_params_mode` API object
type AccountPushParamsMode string

func (p *AccountPushParamsMode) MarshalJSON() ([]byte, error) {
	return GetStringFromRange(string(*p), "on", "off", "no_sound", "no_text")
}

func (p *AccountPushParamsMode) String() string {
	return string(*p)
}

// PushParamsSettings represents `account_push_params_settings` API object
type AccountPushParamsSettings string

func (p *AccountPushParamsSettings) MarshalJSON() ([]byte, error) {
	return GetStringFromRange(string(*p), "on", "off", "fr_of_fr")
}

func (p *AccountPushParamsSettings) String() string {
	return string(*p)
}

// PushSettings represents `account_push_settings` API object
type AccountPushSettings struct {
	Conversations AccountPushConversations `json:"conversations"`
	Disabled      BaseBoolInt              `json:"disabled"`
	DisabledUntil int                      `json:"disabled_until"`
	Settings      AccountPushParams        `json:"settings"`
}

// A UserSettings represents 'account_user_settings` API object
type AccountUserSettings struct {
	Birthday         string             `json:"bdate"`
	BirthVisibility  int                `json:"bdate_visibility"`
	City             BaseObject         `json:"city"`
	Country          BaseCountry        `json:"country"`
	FirstName        string             `json:"first_name"`
	Hometown         string             `json:"home_town"`
	LastName         string             `json:"last_name"`
	MaidenName       string             `json:"maiden_name"`
	NameRequest      AccountNameRequest `json:"name_request"`
	Phone            string             `json:"phone"`
	Relation         int                `json:"relation"`
	RelationPartner  UsersUserMin       `json:"relation_partner"`
	RelationPending  int                `json:"relation_pending"`
	RelationRequests []UsersUserMin     `json:"relation_requests"`
	ScreenName       string             `json:"screen_name"`
	Sex              BaseSex            `json:"sex"`
	Status           string             `json:"status"`
}
