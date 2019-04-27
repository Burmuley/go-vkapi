package objects

//###########################################################
// Account related API objects
//###########################################################

// AccountInfo type represents 'account_info' API objects
type AccountInfo struct {
	TwoFARequired   BaseBoolInt `json:"2fa_required"`
	Country         string      `json:"country"`
	HTTPSRequired   BaseBoolInt `json:"https_required"`
	Intro           BaseBoolInt `json:"intro"`
	Lang            int         `json:"lang"`
	NoWallReplies   BaseBoolInt `json:"no_wall_replies"`
	OwnPostsDefault BaseBoolInt `json:"own_posts_default"`
}

// A AccountUserSettings represents 'account_user_settings` API object
type AccountUserSettings struct {
	FirstName       string      `json:"first_name"`
	LastName        string      `json:"last_name"`
	ScreenName      string      `json:"screen_name"`
	Sex             int         `json:"sex"`
	Relation        int         `json:"relation"`
	Birthday        string      `json:"bdate"`
	BirthVisibility int         `json:"bdate_visibility"`
	Hometown        string      `json:"home_town"`
	Status          string      `json:"status"`
	Phone           string      `json:"phone"`
	Country         BaseCountry `json:"country"`
	City            BaseObject  `json:"city"`
}
