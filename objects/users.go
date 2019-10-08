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
// `users` group of objects
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

// UsersUniversity type represents `users_university` API object
type UsersUniversity struct {
	Chair           int    `json:"chair"`            // Chair ID
	ChairName       string `json:"chair_name"`       // Chair name
	City            int    `json:"city"`             // City ID
	Country         int    `json:"country"`          // Country ID
	EducationForm   string `json:"education_form"`   // Education form
	EducationStatus string `json:"education_status"` // Education status
	Faculty         int    `json:"faculty"`          // Faculty ID
	FacultyName     string `json:"faculty_name"`     // Faculty name
	Graduation      int    `json:"graduation"`       // Graduation year
	Id              int    `json:"id"`               // University ID
	Name            string `json:"name"`             // University name
}

// UsersUsersArray type represents `users_users_array` API object
type UsersUsersArray struct {
	Count int   `json:"count"` // Users number
	Items []int `json:"items"`
}

// UsersFields type represents `users_fields` API object
type UsersFields string

// UsersCropPhotoCrop type represents `users_crop_photo_crop` API object
type UsersCropPhotoCrop struct {
	X  float64 `json:"x"`  // Coordinate X of the left upper corner
	X2 float64 `json:"x2"` // Coordinate X of the right lower corner
	Y  float64 `json:"y"`  // Coordinate Y of the left upper corner
	Y2 float64 `json:"y2"` // Coordinate Y of the right lower corner
}

// UsersSchool type represents `users_school` API object
type UsersSchool struct {
	City          int    `json:"city"`           // City ID
	Class         string `json:"class"`          // School class letter
	Country       int    `json:"country"`        // Country ID
	Id            string `json:"id"`             // School ID
	Name          string `json:"name"`           // School name
	Type          int    `json:"type"`           // School type ID
	TypeStr       string `json:"type_str"`       // School type name
	YearFrom      int    `json:"year_from"`      // Year the user started to study
	YearGraduated int    `json:"year_graduated"` // Graduation year
	YearTo        int    `json:"year_to"`        // Year the user finished to study
}

// UsersUserXtrType type represents `users_user_xtr_type` API object
type UsersUserXtrType struct {
	UsersUser UsersUser     `json:"UsersUser"`
	Type      UsersUserType `json:"type"`
}

// UsersOccupation type represents `users_occupation` API object
type UsersOccupation struct {
	Id   int    `json:"id"`   // ID of school, university, company group
	Name string `json:"name"` // Name of occupation
	Type string `json:"type"` // Type of occupation
}

// UsersRelative type represents `users_relative` API object
type UsersRelative struct {
	BirthDate string `json:"birth_date"` // Date of child birthday (format dd.mm.yyyy)
	Id        int    `json:"id"`         // Relative ID
	Name      string `json:"name"`       // Name of relative
	Type      string `json:"type"`       // Relative type
}

// UsersUserConnections type represents `users_user_connections` API object
type UsersUserConnections struct {
	Facebook     string `json:"facebook"`      // User's Facebook account
	FacebookName string `json:"facebook_name"` // User's Facebook name
	Instagram    string `json:"instagram"`     // User's Instagram account
	Livejournal  string `json:"livejournal"`   // User's Livejournal account
	Skype        string `json:"skype"`         // User's Skype nickname
	Twitter      string `json:"twitter"`       // User's Twitter account
}

// UsersCropPhotoRect type represents `users_crop_photo_rect` API object
type UsersCropPhotoRect struct {
	X  float64 `json:"x"`  // Coordinate X of the left upper corner
	X2 float64 `json:"x2"` // Coordinate X of the right lower corner
	Y  float64 `json:"y"`  // Coordinate Y of the left upper corner
	Y2 float64 `json:"y2"` // Coordinate Y of the right lower corner
}

// UsersLastSeen type represents `users_last_seen` API object
type UsersLastSeen struct {
	Platform int `json:"platform"` // Type of the platform that used for the last authorization
	Time     int `json:"time"`     // Last visit date (in Unix time)
}

// UsersSubscriptionsItem type represents `users_subscriptions_item` API object
type UsersSubscriptionsItem struct {
	GroupsGroupFull  GroupsGroupFull  `json:"GroupsGroupFull"`
	UsersUserXtrType UsersUserXtrType `json:"UsersUserXtrType"`
}

// UsersUserXtrCounters type represents `users_user_xtr_counters` API object
type UsersUserXtrCounters struct {
	UsersUserFull UsersUserFull     `json:"UsersUserFull"`
	Counters      UsersUserCounters `json:"counters"`
}

// UsersMilitary type represents `users_military` API object
type UsersMilitary struct {
	CountryId int    `json:"country_id"` // Country ID
	From      int    `json:"from"`       // From year
	Id        int    `json:"id"`         // Military ID
	Unit      string `json:"unit"`       // Unit name
	UnitId    int    `json:"unit_id"`    // Unit ID
	Until     int    `json:"until"`      // Till year
}

// UsersPersonal type represents `users_personal` API object
type UsersPersonal struct {
	Alcohol    int      `json:"alcohol"`     // User's views on alcohol
	InspiredBy string   `json:"inspired_by"` // User's inspired by
	Langs      []string `json:"langs"`
	LifeMain   int      `json:"life_main"`   // User's personal priority in life
	PeopleMain int      `json:"people_main"` // User's personal priority in people
	Political  int      `json:"political"`   // User's political views
	Religion   string   `json:"religion"`    // User's religion
	ReligionId int      `json:"religion_id"` // User's religion id
	Smoking    int      `json:"smoking"`     // User's views on smoking
}

// UsersUserCounters type represents `users_user_counters` API object
type UsersUserCounters struct {
	Albums        int `json:"albums"`         // Albums number
	Audios        int `json:"audios"`         // Audios number
	Followers     int `json:"followers"`      // Followers number
	Friends       int `json:"friends"`        // Friends number
	Gifts         int `json:"gifts"`          // Gifts number
	Groups        int `json:"groups"`         // Communities number
	Notes         int `json:"notes"`          // Notes number
	OnlineFriends int `json:"online_friends"` // Online friends number
	Pages         int `json:"pages"`          // Public pages number
	Photos        int `json:"photos"`         // Photos number
	Subscriptions int `json:"subscriptions"`  // Subscriptions number
	UserPhotos    int `json:"user_photos"`    // Number of photos with user
	UserVideos    int `json:"user_videos"`    // Number of videos with user
	Videos        int `json:"videos"`         // Videos number
}

// UsersUserRelation type represents `users_user_relation` API object
type UsersUserRelation int

// UsersExports type represents `users_exports` API object
type UsersExports struct {
	Facebook    int `json:"facebook"`
	Livejournal int `json:"livejournal"`
	Twitter     int `json:"twitter"`
}

// UsersUserMin type represents `users_user_min` API object
type UsersUserMin struct {
	CanAccessClosed bool   `json:"can_access_closed"`
	Deactivated     string `json:"deactivated"` // Returns if a profile is deleted or blocked
	FirstName       string `json:"first_name"`  // User first name
	Hidden          int    `json:"hidden"`      // Returns if a profile is hidden.
	Id              int    `json:"id"`          // User ID
	IsClosed        bool   `json:"is_closed"`
	LastName        string `json:"last_name"` // User last name
}

// UsersCareer type represents `users_career` API object
type UsersCareer struct {
	CityId    int    `json:"city_id"`    // City ID
	Company   string `json:"company"`    // Company name
	CountryId int    `json:"country_id"` // Country ID
	From      int    `json:"from"`       // From year
	GroupId   int    `json:"group_id"`   // Community ID
	Id        int    `json:"id"`         // Career ID
	Position  string `json:"position"`   // Position
	Until     int    `json:"until"`      // Till year
}

// UsersUserFull type represents `users_user_full` API object
type UsersUserFull struct {
	UsersUser              UsersUser         `json:"UsersUser"`
	Activity               string            `json:"activity"`                  // User's status
	Bdate                  string            `json:"bdate"`                     // User's date of birth
	Blacklisted            BaseBoolInt       `json:"blacklisted"`               // Information whether current user is in the requested user's blacklist.
	BlacklistedByMe        BaseBoolInt       `json:"blacklisted_by_me"`         // Information whether the requested user is in current user's blacklist
	CanPost                BaseBoolInt       `json:"can_post"`                  // Information whether current user can post on the user's wall
	CanSeeAllPosts         BaseBoolInt       `json:"can_see_all_posts"`         // Information whether current user can see other users' audio on the wall
	CanSeeAudio            BaseBoolInt       `json:"can_see_audio"`             // Information whether current user can see the user's audio
	CanSendFriendRequest   BaseBoolInt       `json:"can_send_friend_request"`   // Information whether current user can send a friend request
	CanSubscribePodcasts   bool              `json:"can_subscribe_podcasts"`    // Owner in whitelist or not
	CanSubscribePosts      bool              `json:"can_subscribe_posts"`       // Can subscribe to wall
	CanWritePrivateMessage BaseBoolInt       `json:"can_write_private_message"` // Information whether current user can write private message
	Career                 []UsersCareer     `json:"career"`
	City                   BaseObject        `json:"city"`
	CommonCount            int               `json:"common_count"` // Number of common friends with current user
	Country                BaseCountry       `json:"country"`
	CropPhoto              UsersCropPhoto    `json:"crop_photo"`
	Domain                 string            `json:"domain"`           // Domain name of the user's page
	EducationForm          string            `json:"education_form"`   // Education form
	EducationStatus        string            `json:"education_status"` // User's education status
	Exports                UsersExports      `json:"exports"`
	Faculty                int               `json:"faculty"`                // Faculty ID
	FacultyName            string            `json:"faculty_name"`           // Faculty name
	FollowersCount         int               `json:"followers_count"`        // Number of user's followers
	Graduation             int               `json:"graduation"`             // Graduation year
	HasMobile              BaseBoolInt       `json:"has_mobile"`             // Information whether the user specified his phone number
	HasPhoto               BaseBoolInt       `json:"has_photo"`              // Information whether the user has main photo
	HomePhone              string            `json:"home_phone"`             // User's mobile phone number
	HomeTown               string            `json:"home_town"`              // User hometown
	IsFavorite             BaseBoolInt       `json:"is_favorite"`            // Information whether the requested user is in faves of current user
	IsFriend               BaseBoolInt       `json:"is_friend"`              // Information whether the user is a friend of current user
	IsHiddenFromFeed       BaseBoolInt       `json:"is_hidden_from_feed"`    // Information whether the requested user is hidden from current user's newsfeed
	IsSubscribedPodcasts   bool              `json:"is_subscribed_podcasts"` // Information whether current user is subscribed to podcasts
	LastSeen               UsersLastSeen     `json:"last_seen"`
	MaidenName             string            `json:"maiden_name"` // User maiden name
	Military               []UsersMilitary   `json:"military"`
	MobilePhone            string            `json:"mobile_phone"` // Information whether current user can see
	Nickname               string            `json:"nickname"`     // User nickname
	Occupation             UsersOccupation   `json:"occupation"`
	Personal               UsersPersonal     `json:"personal"`
	Photo200               string            `json:"photo_200"`      // URL of square photo of the user with 200 pixels in width
	Photo200Orig           string            `json:"photo_200_orig"` // URL of user's photo with 200 pixels in width
	Photo400Orig           string            `json:"photo_400_orig"` // URL of user's photo with 400 pixels in width
	PhotoId                string            `json:"photo_id"`       // ID of the user's main photo
	PhotoMax               string            `json:"photo_max"`      // URL of square photo of the user with maximum width
	PhotoMaxOrig           string            `json:"photo_max_orig"` // URL of user's photo of maximum size
	Relation               UsersUserRelation `json:"relation"`       // User relationship status
	RelationPartner        UsersUserMin      `json:"relation_partner"`
	Relatives              []UsersRelative   `json:"relatives"`
	Schools                []UsersSchool     `json:"schools"`
	Site                   string            `json:"site"`   // User's website
	Status                 string            `json:"status"` // User's status
	StatusAudio            AudioAudio        `json:"status_audio"`
	Timezone               int               `json:"timezone"` // User's timezone
	Universities           []UsersUniversity `json:"universities"`
	University             int               `json:"university"`      // University ID
	UniversityName         string            `json:"university_name"` // University name
	WallComments           BaseBoolInt       `json:"wall_comments"`   // Information whether current user can comment wall posts
}

// UsersCropPhoto type represents `users_crop_photo` API object
type UsersCropPhoto struct {
	Crop  UsersCropPhotoCrop `json:"crop"`
	Photo PhotosPhoto        `json:"photo"`
	Rect  UsersCropPhotoRect `json:"rect"`
}

// UsersUserType type represents `users_user_type` API object
type UsersUserType string // Object type

// UsersUser type represents `users_user` API object
type UsersUser struct {
	UsersUserMin UsersUserMin              `json:"UsersUserMin"`
	FriendStatus FriendsFriendStatusStatus `json:"friend_status"`
	Mutual       FriendsRequestsMutual     `json:"mutual"`
	Online       BaseBoolInt               `json:"online"`        // Information whether the user is online
	OnlineApp    int                       `json:"online_app"`    // Application ID
	OnlineMobile BaseBoolInt               `json:"online_mobile"` // Information whether the user is online in mobile site or application
	Photo100     string                    `json:"photo_100"`     // URL of square photo of the user with 100 pixels in width
	Photo50      string                    `json:"photo_50"`      // URL of square photo of the user with 50 pixels in width
	ScreenName   string                    `json:"screen_name"`   // Domain name of the user's page
	Sex          BaseSex                   `json:"sex"`           // User sex
	Trending     BaseBoolInt               `json:"trending"`      // Information whether the user has a "fire" pictogram.
	Verified     BaseBoolInt               `json:"verified"`      // Information whether the user is verified
}

// UsersUserSettingsXtr type represents `users_user_settings_xtr` API object
type UsersUserSettingsXtr struct {
	Bdate            string                       `json:"bdate"`            // User's date of birth
	BdateVisibility  int                          `json:"bdate_visibility"` // Information whether user's birthdate are hidden
	City             BaseCity                     `json:"city"`
	Connections      UsersUserConnections         `json:"connections"`
	Country          BaseCountry                  `json:"country"`
	FirstName        string                       `json:"first_name"` // User first name
	HomeTown         string                       `json:"home_town"`  // User's hometown
	Interests        AccountUserSettingsInterests `json:"interests"`
	Languages        []string                     `json:"languages"`
	LastName         string                       `json:"last_name"`   // User last name
	MaidenName       string                       `json:"maiden_name"` // User maiden name
	NameRequest      AccountNameRequest           `json:"name_request"`
	Personal         UsersPersonal                `json:"personal"`
	Phone            string                       `json:"phone"`    // User phone number with some hidden digits
	Relation         UsersUserRelation            `json:"relation"` // User relationship status
	RelationPartner  UsersUserMin                 `json:"relation_partner"`
	RelationPending  BaseBoolInt                  `json:"relation_pending"` // Information whether relation status is pending
	RelationRequests []UsersUserMin               `json:"relation_requests"`
	ScreenName       string                       `json:"screen_name"` // Domain name of the user's page
	Sex              BaseSex                      `json:"sex"`         // User sex
	Status           string                       `json:"status"`      // User status
	StatusAudio      AudioAudio                   `json:"status_audio"`
}
