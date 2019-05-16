package users

import (
	"gitlab.com/Burmuley/go-vkapi/objects/audio"
	"gitlab.com/Burmuley/go-vkapi/objects/base"
	"gitlab.com/Burmuley/go-vkapi/objects/photos"
)

/////////////////////////////////////////////////////////////
// Users related API objects	                           //
/////////////////////////////////////////////////////////////

// Career represents `users_career` API object
type Career struct {
	CityID    int    `json:"city_id"`
	Company   string `json:"company"`
	CountryID int    `json:"country_id"`
	FromYear  int    `json:"from"`
	GroupID   int    `json:"group_id"`
	ID        int    `json:"id"`
	Position  string `json:"position"`
	UntilYear int    `json:"until"`
}

// CropPhoto represents `users_crop_photo` APi object
type CropPhoto struct {
	Crop  CropPhotoCrop `json:"crop"`
	Photo photos.Photo  `json:"photo"`
	Rect  CropPhotoRect `json:"rect"`
}

// CropPhotoCrop represents `users_crop_photo_crop` API object
type CropPhotoCrop struct {
	X  int `json:"x"`
	X2 int `json:"x2"`
	Y  int `json:"y"`
	Y2 int `json:"y2"`
}

// RectPhotoRect represents `users_rect_photo_rect` API object
type CropPhotoRect struct {
	X  int `json:"x"`
	X2 int `json:"x2"`
	Y  int `json:"y"`
	Y2 int `json:"y2"`
}

// Exports represents `users_exports` API object
type Exports struct {
	Facebook    int `json:"facebook"`
	LiveJournal int `json:"livejournal"`
	Twitter     int `json:"twitter"`
}

// Don't think we need this
//type Fields string

// LastSeen represents `users_last_seen` API object
type LastSeen struct {
	Platform int `json:"platform"`
	Time     int `json:"platform"`
}

// Military represents `users_military` API object
type Military struct {
	CountryID int    `json:"country_id"`
	From      int    `json:"from"`
	ID        int    `json:"id"`
	Unit      string `json:"unit"`
	UnitID    int    `json:"unit_id"`
	Until     int    `json:"until"`
}

// Occupation represents `users_occupation` API object
type Occupation struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// Personal represents `users_personal` API object
type Personal struct {
	Alcohol    int      `json:"alcohol"`
	InspiredBy string   `json:"inspired_by"`
	Langs      []string `json:"langs"`
	LifeMain   int      `json:"life_main"`
	PeopleMain int      `json:"people_main"`
	Political  int      `json:"political"`
	Religion   string   `json:"religion"`
	Smoking    int      `json:"smoking"`
}

// Relative represents `users_relative` API object
type Relative struct {
	BirthDate string `json:"birth_date"`
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
}

// School represents `users_school` API object
type School struct {
	City          int    `json:"city"`
	Class         string `json:"class"`
	Country       int    `json:"country"`
	ID            string `json:"id"`
	Name          string `json:"name"`
	Type          int    `json:"type"`
	TypeStr       string `json:"type_str"`
	YearFrom      int    `json:"year_from"`
	YearTo        int    `json:"year_to"`
	YearGraduated int    `json:"year_graduated"`
}

// University represents `users_university` API object
type University struct {
	Chair           int    `json:"chair"`
	ChairName       string `json:"chair_name"`
	City            int    `json:"city"`
	Country         int    `json:"country"`
	EducationFrom   string `json:"education_from"`
	EducationStatus string `json:"education_status"`
	Faculty         int    `json:"faculty"`
	FacultyName     string `json:"faculty_name"`
	Graduation      int    `json:"graduation"`
	ID              int    `json:"id"`
	Name            string `json:"name"`
}

// User represents `users_user` API object
type User struct {
	UserMin
	Sex          base.Sex     `json:"sex"`
	ScreenName   string       `json:"screen_name"`
	Photo50      string       `json:"photo_50"`
	Photo100     string       `json:"photo_100"`
	Online       base.BoolInt `json:"online"`
	OnlineMobile base.BoolInt `json:"online_mobile"`
	OnlineApp    base.BoolInt `json:"online_app"`
}

// UserCounters represents `users_user_counters` API object
type UserCounters struct {
	Albums        int `json:"albums"`
	Audios        int `json:"audios"`
	Followers     int `json:"followers"`
	Friends       int `json:"friends"`
	Gifts         int `json:"gifts"`
	Groups        int `json:"groups"`
	Notes         int `json:"notes"`
	OnlineFriends int `json:"online_friends"`
	Pages         int `json:"pages"`
	Photos        int `json:"photos"`
	Subscriptions int `json:"subscriptions"`
	UserPhotos    int `json:"user_photos"`
	UserVideos    int `json:"user_videos"`
	Videos        int `json:"videos"`
}

// UserFull represents `users_user_full` API object
type UserFull struct {
	User
	Nickname                string          `json:"nickname"`
	MaidenName              string          `json:"maiden_name"`
	Domain                  string          `json:"domain"`
	Birthday                string          `json:"bdate"`
	City                    base.Object     `json:"city"`
	Country                 base.Country    `json:"country"`
	Timezone                int             `json:"timezone"`
	Photo200                string          `json:"photo_200"`
	PhotoMax                string          `json:"photo_max"`
	Photo200Orig            string          `json:"photo_200_orig"`
	Photo400Orig            string          `json:"photo_400_orig"`
	PhotoMaxOrig            string          `json:"photo_max_orig"`
	PhotoID                 string          `json:"photo_id"`
	HasPhoto                base.BoolInt    `json:"has_photo"`
	Trending                base.BoolInt    `json:"trending"`
	HasMobile               base.BoolInt    `json:"has_mobile"`
	IsFriend                base.BoolInt    `json:"is_friend"`
	FriendStatus            int             `json:"friend_status"`
	WallComments            base.BoolInt    `json:"wall_comments"`
	CanPost                 base.BoolInt    `json:"can_post"`
	CanSeeAllPosts          base.BoolInt    `json:"can_see_all_posts"`
	CanSeeAudio             base.BoolInt    `json:"can_see_audio"`
	CanWritePrivateMessages base.BoolInt    `json:"can_write_private_messages"`
	CanSendFriendRequest    base.BoolInt    `json:"can_send_friend_request"`
	MobilePhone             string          `json:"mobile_phone"`
	HomePhone               string          `json:"home_phone"`
	Skype                   string          `json:"skype"`
	Facebook                string          `json:"facebook"`
	FacebookName            string          `json:"facebook_name"`
	Twitter                 string          `json:"twitter"`
	Livejournal             string          `json:"livejournal"`
	Instagram               string          `json:"instagram"`
	Site                    string          `json:"site"`
	StatusAudio             audio.AudioFull `json:"status_audio"`
	Status                  string          `json:"status"`
	Activity                string          `json:"activity"`
	LastSeen                LastSeen        `json:"last_seen"`
	Exports                 Exports         `json:"exports"`
	CropPhoto               CropPhoto       `json:"crop_photo"`
	Verified                base.BoolInt    `json:"verified"`
	FollowersCount          int             `json:"followers_count"`
	Blacklisted             base.BoolInt    `json:"blacklisted"`
	BlacklistedByMe         base.BoolInt    `json:"blacklisted_by_me"`
	IsFavorite              base.BoolInt    `json:"is_favorite"`
	IsHiddenFromFeed        base.BoolInt    `json:"is_hidden_from_feed"`
	CommonCount             int             `json:"common_count"`
	Occupation              Occupation      `json:"occupation"`
	Career                  []Career        `json:"career"`
	Military                []Military      `json:"military"`
	University              int             `json:"university"`
	UniversityName          string          `json:"university_name"`
	Faculty                 int             `json:"faculty"`
	FacultyName             string          `json:"faculty_name"`
	Graduation              int             `json:"graduation"`
	EducationForm           string          `json:"education_form"`
	EducationStatus         string          `json:"education_status"`
	HomeTown                string          `json:"home_town"`
	Relation                int             `json:"relation"`
	RelationPartner         UserMin         `json:"relation_partner"`
	Personal                Personal        `json:"personal"`
	Interests               string          `json:"interests"`
	Music                   string          `json:"music"`
	Activities              string          `json:"activities"`
	Movies                  string          `json:"movies"`
	Tv                      string          `json:"tv"`
	Books                   string          `json:"books"`
	Games                   string          `json:"games"`
	Universities            []University    `json:"universities"`
	Schools                 []School        `json:"schools"`
	About                   string          `json:"about"`
	Relatives               []Relative      `json:"relatives"`
	Quotes                  string          `json:"quotes"`
	IsSubscribedPodcasts    bool            `json:"is_subscribed_podcasts"`
	CanSubscribePodcasts    bool            `json:"can_subscribe_podcasts"`
	CanSubscribePosts       bool            `json:"can_subscribe_posts"`
}

// UserMin represents `users_user_min` API object
type UserMin struct {
	Deactivated string `json:"deactivated"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Hidden      int    `json:"hidden"`
	ID          int    `json:"id"`
}

// UserType represents `users_user_type` API object
type UserType string

func (u *UserType) GetName() string {
	return string(*u)
}

// UserXtrCounters represents `users_user_xtr_counters` API object
type UserXtrCounters struct {
	UserFull
	Counters UserCounters `json:"counters"`
}

// UserXtrType represents `users_user_xtr_type` API object
type UserXtrType struct {
	User
	Type UserType `json:"type"`
}

// UsersArray represents `users_users_array` API object
type UsersArray struct {
	Count int   `json:"count"`
	Items []int `json:"items"`
}
