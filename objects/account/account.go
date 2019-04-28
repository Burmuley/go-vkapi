package account

import "gitlab.com/Burmuley/go-vkapi/objects/base"

//###########################################################
// Account related API objects
//###########################################################

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

// A UserSettings represents 'account_user_settings` API object
type UserSettings struct {
	FirstName       string       `json:"first_name"`
	LastName        string       `json:"last_name"`
	ScreenName      string       `json:"screen_name"`
	Sex             int          `json:"sex"`
	Relation        int          `json:"relation"`
	Birthday        string       `json:"bdate"`
	BirthVisibility int          `json:"bdate_visibility"`
	Hometown        string       `json:"home_town"`
	Status          string       `json:"status"`
	Phone           string       `json:"phone"`
	Country         base.Country `json:"country"`
	City            base.Object  `json:"city"`
}
