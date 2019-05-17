package groups

import (
	"fmt"
	"gitlab.com/Burmuley/go-vkapi/objects"
	"gitlab.com/Burmuley/go-vkapi/objects/base"
	"gitlab.com/Burmuley/go-vkapi/objects/market"
	"gitlab.com/Burmuley/go-vkapi/objects/users"
	"strconv"
)

/////////////////////////////////////////////////////////////
// Groups related API objects	                           //
/////////////////////////////////////////////////////////////

// Address represents `groups_address` API object
type Address struct {
	AdditionalAddress string                `json:"additional_address"` // Additional address to the place (6 floor, left door)
	Address           string                `json:"address"`            // String address to the place (Nevsky, 28)
	CityID            int                   `json:"city_id"`            // City id of address
	CountryID         int                   `json:"country_id"`         // Country id of address
	Distance          int                   `json:"distance"`           // Distance from the point
	ID                int                   `json:"id"`                 // Address id
	Latitude          float64               `json:"latitude"`           // Address latitude
	Longitude         float64               `json:"longitude"`          // Address longitude
	MetroID           int                   `json:"metro_id"`           // Metro id of address
	Phone             string                `json:"phone"`              // Address phone
	TimeOffset        int                   `json:"time_offset"`        // Time offset int minutes from utc time
	Timetable         AddressTimetable      `json:"timetable"`          // Week timetable for the address
	Title             string                `json:"title"`              // Title of the place (Zinger, etc)
	WorkInfoStatus    AddressWorkInfoStatus `json:"work_info_status"`   // Status of information about timetable
}

// AddressTimetable represents `groups_address_timetable` API object
type AddressTimetable struct {
	Monday    AddressTimetableDay `json:"mon"` // Timetable for Monday
	Tuesday   AddressTimetableDay `json:"tue"` // Timetable for Tuesday
	Wednesday AddressTimetableDay `json:"wed"` // Timetable for Wednesday
	Thursday  AddressTimetableDay `json:"thu"` // Timetable for Thursday
	Friday    AddressTimetableDay `json:"fri"` // Timetable for Friday
	Saturday  AddressTimetableDay `json:"sat"` // Timetable for Saturday
	Sunday    AddressTimetableDay `json:"sun"` // Timetable for Sunday
}

// AddressTimetableDay represents `groups_address_timetable_day` API object
type AddressTimetableDay struct {
	BreakCloseTime int `json:"break_close_time"` // Close time of the break in minutes
	BreakOpenTime  int `json:"break_open_time"`  // Start time of the break in minutes
	CloseTime      int `json:"close_time"`       // Close time in minutes
	OpenTime       int `json:"open_time"`        // Open time in minutes
}

// AddressWorkInfoStatus represents `groups_address_work_info_status` API object
type AddressWorkInfoStatus string

func (a *AddressWorkInfoStatus) MarshalJSON() ([]byte, error) {
	return objects.GetStringFromRange(string(*a), "no_information", "temporarily_closed",
		"always_opened", "timetable", "forever_closed")
}

func (a *AddressWorkInfoStatus) GetName() string {
	return string(*a)
}

// AddressesInfo represents `groups_addresses_info` API object
type AddressesInfo struct {
	Enabled       bool `json:"is_enabled"`      // Information whether addresses is enabled
	MainAddressID int  `json:"main_address_id"` // Main address id for group
}

// BanInfo represents `groups_ban_info` APi object
type BanInfo struct {
	AdminID int           `json:"admin_id"` // Administrator ID
	Comment string        `json:"comment"`  // Comment for a ban
	Date    int           `json:"date"`     // Date when user has been added to blacklist in Unixtime
	EndDate int           `json:"end_date"` // Date when user will be removed from blacklist in Unixtime
	Reason  BanInfoReason `json:"reason"`
}

// BanInfoReason represents `groups_ban_info_reason` API object
type BanInfoReason int

func (b *BanInfoReason) GetName() string {
	switch *b {
	case 0:
		return "other"
	case 1:
		return "spam"
	case 2:
		return "verbal abuse"
	case 3:
		return "strong language"
	case 4:
		return "flood"
	default:
		return "unknown"
	}
}

// CallbackSettings represents `groups_callback_settings` API object
type CallbackSettings struct {
	ApiVersion string         `json:"api_version"` // API version used for the events
	Events     LongPollEvents `json:"events"`
}

// ContactsItem represents `groups_contacts_item` API object
type ContactsItem struct {
	Description string `json:"desc"`    // Contact description
	Email       string `json:"email"`   // Contact email
	Phone       string `json:"phone"`   // Contact phone
	UserID      int    `json:"user_id"` // User ID
}

// Counters represents `groups_counters_group` API object
type Counters struct {
	Addresses int `json:"addresses"` // Addresses number
	Albums    int `json:"albums"`    // Photo albums number
	Audios    int `json:"audios"`    // Audios number
	Docs      int `json:"docs"`      // Docs number
	Market    int `json:"market"`    // Market items number
	Photos    int `json:"photos"`    // Photos number
	Topics    int `json:"topics"`    // Topics number
	Videos    int `json:"videos"`    // Videos number
}

// Cover represents `group_cover` APi object
type Cover struct {
	Enabled base.BoolInt `json:"enabled"` // Information whether cover is enabled
	Images  base.Image   `json:"images"`
}

// Fields represents `groups_fields` API object
type Fields string

func (f *Fields) MarshalJSON() ([]byte, error) {
	return objects.GetStringFromRange(string(*f), "market", "member_status", "is_favorite", "is_subscribed",
		"city", "country", "verified", "description", "wiki_page", "members_count", "counters", "cover",
		"can_post", "can_see_all_posts", "activity", "fixed_post", "can_create_topic", "can_upload_video", "has_photo",
		"status", "main_album_id", "links", "contacts", "site", "main_section", "trending", "can_message",
		"is_messages_blocked", "can_send_notify", "online_status", "start_date", "finish_date", "age_limits",
		"ban_info", "action_button", "author_id", "phone", "has_market_app", "addresses", "live_covers", "is_adult",
		"can_subscribe_posts")
}

func (f *Fields) GetName() string {
	return string(*f)
}

// Filter represents `groups_filter` API object
type Filter string

func (f *Filter) MarshalJSON() ([]byte, error) {
	return objects.GetStringFromRange(string(*f), "admin", "editor", "moder", "groups", "publics",
		"events", "has_addresses")
}

func (f *Filter) GetName() string {
	return string(*f)
}

// Group represents `groups_group` API object
type Group struct {
	AdminLevel   GroupAdminLevel `json:"admin_level"`
	Deactivated  string          `json:"deactivated"`   // Information whether community is banned
	FinishDate   int             `json:"finish_date"`   // Finish date in Unixtime format
	ID           int             `json:"id"`            // Community ID
	IsAdmin      base.BoolInt    `json:"is_admin"`      // Information whether current user is administrator
	IsAdvertiser base.BoolInt    `json:"is_advertiser"` // Information whether current user is advertiser
	IsClosed     GroupClosed     `json:"is_closed"`
	IsMember     base.BoolInt    `json:"is_member"`   // Information whether current user is member
	Name         string          `json:"name"`        // Community name
	Photo100     string          `json:"photo_100"`   // URL of square photo of the community with 100 pixels in width
	Photo200     string          `json:"photo_200"`   // URL of square photo of the community with 200 pixels in width
	Photo50      string          `json:"photo_50"`    // URL of square photo of the community with 50 pixels in width
	ScreenName   string          `json:"screen_name"` // Domain of the community page
	StartDate    int             `json:"start_date"`  // Start date in Unixtime format
	Type         GroupType       `json:"type"`
}

// GroupAccess represents `groups_group_access` API object
type GroupAccess int

func (g *GroupAccess) MarshalJSON() ([]byte, error) {
	return objects.GetIntFromRange(int(*g), 0, 2, true)
}

func (g *GroupAccess) GetName() string {
	switch *g {
	case 0:
		return "open"
	case 1:
		return "closed"
	case 2:
		return "private"
	default:
		return "unknown"
	}
}

// GroupAdminLevel represents `groups_group_admin_level` API object
type GroupAdminLevel int

func (g *GroupAdminLevel) MarshalJSON() ([]byte, error) {
	return objects.GetIntFromRange(int(*g), 1, 3, true)
}

func (g *GroupAdminLevel) GetName() string {
	switch *g {
	case 1:
		return "moderator"
	case 2:
		return "editor"
	case 3:
		return "administrator"
	default:
		return "unknown"
	}
}

// GroupAgeLimits represents `groups_group_age_limits` API object
type GroupAgeLimits int

func (g *GroupAgeLimits) MarshalJSON() ([]byte, error) {
	return objects.GetIntFromRange(int(*g), 1, 3, true)
}

func (g *GroupAgeLimits) GetName() string {
	switch *g {
	case 1:
		return "unlimited"
	case 2:
		return "16 plus"
	case 3:
		return "18 plus"
	default:
		return "unknown"
	}
}

// GroupAudio represents `groups_group_audio` API object
type GroupAudio int

func (g *GroupAudio) MarshalJSON() ([]byte, error) {
	return objects.GetIntFromRange(int(*g), 0, 2, true)
}

func (g *GroupAudio) GetName() string {
	switch *g {
	case 0:
		return "disabled"
	case 1:
		return "open"
	case 2:
		return "limited"
	default:
		return "unknown"
	}
}

// GroupBanInfo represents `GroupBanInfo` API object
type GroupBanInfo struct {
	Comment string `json:"comment"`  // Ban comment
	EndDate int    `json:"end_date"` // End date of ban in Unixtime
}

// GroupCategory represents `groups_group_category` API object
type GroupCategory struct {
	ID            int                 `json:"id"`   // Category ID
	Name          string              `json:"name"` // Category name
	Subcategories base.ObjectWithName `json:"subcategories"`
}

// GroupCategoryFull represents `groups_group_category_full` AI object
type GroupCategoryFull struct {
	ID            int                 `json:"id"`         // Category ID
	Name          string              `json:"name"`       // Category name
	PageCount     int                 `json:"page_count"` // Pages number
	PagePreviews  []Group             `json:"page_previews"`
	Subcategories base.ObjectWithName `json:"subcategories"`
}

// CategoryType represents `groups_group_category_type` API object
type CategoryType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GroupDocs represents `groups_group_docs` API object
type GroupDocs int

func (g *GroupDocs) MarshalJSON() ([]byte, error) {
	return objects.GetIntFromRange(int(*g), 0, 2, true)
}

func (g *GroupDocs) GetName() string {
	switch *g {
	case 0:
		return "disabled"
	case 1:
		return "open"
	case 2:
		return "limited"
	default:
		return "unknown"
	}
}

// GroupFull represents `groups_group_full` API object
type GroupFull struct {
	Group
	Market               MarketInfo            `json:"market"`
	MemberStatus         GroupFullMemberStatus `json:"member_status"` // Current user's member status
	IsFavorite           base.BoolInt          `json:"is_favorite"`   // Information whether community is in faves
	IsSubscribed         base.BoolInt          `json:"is_subscribed"` // Information whether current user is subscribed
	City                 base.Object           `json:"city"`
	Country              base.Country          `json:"country"`
	Verified             base.BoolInt          `json:"verified"`      // Information whether community is verified
	Description          string                `json:"description"`   // Community description
	WikiPage             string                `json:"wiki_page"`     // Community's main wiki page title
	MembersCount         int                   `json:"members_count"` // Community members number
	Counters             Counters              `json:"counters"`
	Cover                Cover                 `json:"cover"`
	CanPost              base.BoolInt          `json:"can_post"`          // Information whether current user can post on community's wall
	CanSeeAllPosts       base.BoolInt          `json:"can_see_all_posts"` // Information whether current user can see all posts on community's wall
	Activity             string                `json:"activity"`          // Type of group, start date of event or category of public page
	FixedPost            int                   `json:"fixed_post"`        // Fixed post ID
	CanCreateTopic       base.BoolInt          `json:"can_create_topic"`  // Information whether current user can create topic
	CanUploadVideo       base.BoolInt          `json:"can_upload_video"`  // Information whether current user can upload video
	HasPhoto             base.BoolInt          `json:"has_photo"`         // Information whether community has photo
	Status               string                `json:"status"`            // Community status
	MainAlbumID          int                   `json:"main_album_id"`     // Community's main photo album ID
	Links                []LinksItem           `json:"links"`
	Contacts             []ContactsItem        `json:"contacts"`
	Site                 string                `json:"site"` // Community's website
	MainSection          GroupFullMainSection  `json:"main_section"`
	Trending             base.BoolInt          `json:"trending"`            // Information whether the community has a \"fire\" pictogram.
	CanMessage           base.BoolInt          `json:"can_message"`         // Information whether current user can send a message to community
	MessagesBlocked      base.BoolInt          `json:"is_messages_blocked"` // Information whether community can send a message to current user
	CanSendNotify        base.BoolInt          `json:"can_send_notify"`     // Information whether community can send notifications by phone number to current user
	OnlineStatus         OnlineStatus          `json:"online_status"`       // Status of replies in community messages
	AgeLimits            GroupFullAgeLimits    `json:"age_limits"`          // Information whether age limit
	BanInfo              GroupBanInfo          `json:"ban_info"`            // User ban info
	Addresses            AddressesInfo         `json:"addresses"`
	IsSubscribedPodcasts bool                  `json:"is_subscribed_podcasts"` // Information whether current user is subscribed to podcasts
	CanSubscribePodcasts bool                  `json:"can_subscribe_podcasts"` // Owner in whitelist or not
	CanSubscribeToWall   bool                  `json:"can_subscribe_to_wall"`  // Can subscribe to wall
}

// GroupFullAgeLimits represents `groups_group_full_age_limits` API object
type GroupFullAgeLimits int

func (g *GroupFullAgeLimits) MarshalJSON() ([]byte, error) {
	return objects.GetIntFromRange(int(*g), 1, 3, true)
}

func (g *GroupFullAgeLimits) GetName() string {
	switch *g {
	case 1:
		return "no"
	case 2:
		return "over 16"
	case 3:
		return "over 18"
	default:
		return "unknown"
	}
}

// GroupFullMainSection represents `groups_group_full_main_section` API object
type GroupFullMainSection int

func (g *GroupFullMainSection) MarshalJSON() ([]byte, error) {
	return objects.GetIntFromRange(int(*g), 0, 5, true)
}

func (g *GroupFullMainSection) GetName() string {
	switch *g {
	case 0:
		return "absent"
	case 1:
		return "photos"
	case 2:
		return "topics"
	case 3:
		return "audio"
	case 4:
		return "video"
	case 5:
		return "market"
	default:
		return "unknown"
	}
}

// GroupFullMemberStatus represents `groups_group_full_main_section` API object
type GroupFullMemberStatus int

func (g *GroupFullMemberStatus) MarshalJSON() ([]byte, error) {
	return objects.GetIntFromRange(int(*g), 0, 5, true)
}

func (g *GroupFullMemberStatus) GetName() string {
	switch *g {
	case 0:
		return "not a member"
	case 1:
		return "member"
	case 2:
		return "not sure"
	case 3:
		return "declined"
	case 4:
		return "has sent a request"
	case 5:
		return "invited"
	default:
		return "unknown"
	}
}

// GroupClosed represents `groups_group_is_closed` API objects
type GroupClosed int

func (g *GroupClosed) MarshalJSON() ([]byte, error) {
	return objects.GetIntFromRange(int(*g), 0, 2, true)
}

func (g *GroupClosed) GetName() string {
	switch *g {
	case 0:
		return "open"
	case 1:
		return "closed"
	case 2:
		return "private"
	default:
		return "unknown"
	}
}

// GroupLink represents `groups_group_link` API object
type GroupLink struct {
	Description     string       `json:"desc"`             // Link description
	EditTitle       string       `json:"edit_title"`       // Information whether the title can be edited
	ID              int          `json:"id"`               // Link ID
	ImageProcessing base.BoolInt `json:"image_processing"` // Information whether the image on processing
	Url             string       `json:"url"`              // Link URL
}

// GroupMarketCurrency represents `groups_group_market_currency` API object
type GroupMarketCurrency int

func (g *GroupMarketCurrency) MarshalJSON() ([]byte, error) {
	switch *g {
	case 643, 980, 398, 978, 840:
		return []byte{byte(*g)}, nil
	}

	return []byte{}, fmt.Errorf("value should be within list [643, 980, 398, 978, 840]")
}

func (g *GroupMarketCurrency) GetName() string {
	switch *g {
	case 643:
		return "russian rubles"
	case 980:
		return "ukrainian hryvnia"
	case 398:
		return "kazakh tenge"
	case 978:
		return "euro"
	case 840:
		return "us dollars"
	default:
		return "unknown"
	}
}

// GroupPhotos represents `groups_group_photos` API object
type GroupPhotos int

func (g *GroupPhotos) MarshalJSON() ([]byte, error) {
	return objects.GetIntFromRange(int(*g), 0, 2, true)
}

func (g *GroupPhotos) GetName() string {
	switch *g {
	case 0:
		return "disabled"
	case 1:
		return "open"
	case 2:
		return "limited"
	default:
		return "unknown"
	}
}

// GroupPublicCategoryList represents `groups_group_public_category_list` API object
type GroupPublicCategoryList struct {
	ID           int            `json:"id"`
	Name         string         `json:"name"`
	SubtypesList []CategoryType `json:"subtypes_list"`
}

// GroupRole represents `groups_group_role` API object
type GroupRole string

func (g *GroupRole) MarshalJSON() ([]byte, error) {
	return objects.GetStringFromRange(string(*g), "moderator", "editor", "administrator")
}

func (g *GroupRole) GetName() string {
	return string(*g)
}

// GroupSettings represents `groups_group_settings` API object
type GroupSettings struct {
	Access             int                     `json:"access"`            // Community access settings
	Address            string                  `json:"address"`           // Community's page domain
	Audio              int                     `json:"audio"`             // Audio settings
	Description        string                  `json:"description"`       // Community description
	Docs               int                     `json:"docs"`              // Docs settings
	ObsceneFilter      base.BoolInt            `json:"obscene_filter"`    // Information whether the obscene filter is enabled
	ObsceneStopwords   base.BoolInt            `json:"obscene_stopwords"` // Information whether the stopwords filter is enabled
	ObsceneWords       string                  `json:"obscene_words"`     // The list of stop words
	PhotosSettings     int                     `json:"photos_settings"`   // Photos settings
	PublicCategory     int                     `json:"public_category"`   // Information about the group category
	PublicCategoryList GroupPublicCategoryList `json:"public_category_list"`
	PublicSubcategory  int                     `json:"public_subcategory"` // Information about the group subcategory
	RSS                string                  `json:"rss"`                // URL of the RSS feed
	Subject            int                     `json:"subject"`            // Community subject ID
	SubjectList        []SubjectItem           `json:"subject_list"`
	Title              string                  `json:"title"`         // Community title
	Topics             int                     `json:"topics"`        // Topics settings
	Video              int                     `json:"video"`         // Video settings
	Wall               int                     `json:"wall"`          // Wall settings
	Website            string                  `json:"website"`       // Community website
	WikiSettings       int                     `json:"wiki_settings"` // Wiki settings
}

// GroupSubject represents `groups_group_subject` API object
type GroupSubject string

func (g *GroupSubject) MarshalJSON() ([]byte, error) {
	intG, err := strconv.Atoi(string(*g))

	if err != nil {
		return []byte{}, err
	}

	return objects.GetIntFromRange(intG, 1, 42, true)
}

func (g *GroupSubject) GetName() string {
	intG, err := strconv.Atoi(string(*g))

	if err != nil {
		return string(*g)
	}

	switch intG {
	case 1:
		return "auto"
	case 2:
		return "activity holidays"
	case 3:
		return "business"
	case 4:
		return "pets"
	case 5:
		return "health"
	case 6:
		return "dating and communication"
	case 7:
		return "games"
	case 8:
		return "it"
	case 9:
		return "cinema"
	case 10:
		return "beauty and fashion"
	case 11:
		return "cooking"
	case 12:
		return "art and culture"
	case 13:
		return "literature"
	case 14:
		return "mobile services and internet"
	case 15:
		return "music"
	case 16:
		return "science and technology"
	case 17:
		return "real estate"
	case 18:
		return "news and media"
	case 19:
		return "security"
	case 20:
		return "education"
	case 21:
		return "home and renovations"
	case 22:
		return "politics"
	case 23:
		return "food"
	case 24:
		return "industry"
	case 25:
		return "travel"
	case 26:
		return "work"
	case 27:
		return "entertainment"
	case 28:
		return "religion"
	case 29:
		return "family"
	case 30:
		return "sports"
	case 31:
		return "insurance"
	case 32:
		return "television"
	case 33:
		return "goods and services"
	case 34:
		return "hobbies"
	case 35:
		return "finance"
	case 36:
		return "photo"
	case 37:
		return "esoterics"
	case 38:
		return "electronics and appliances"
	case 39:
		return "erotic"
	case 40:
		return "humor"
	case 41:
		return "society_humanities"
	case 42:
		return "design and graphics"
	default:
		return "unknown"
	}
}

// GroupTopics represents `groups_group_topics` API object
type GroupTopics int

func (g *GroupTopics) MarshalJSON() ([]byte, error) {
	return objects.GetIntFromRange(int(*g), 0, 2, true)
}

func (g *GroupTopics) GetName() string {
	switch *g {
	case 0:
		return "disabled"
	case 1:
		return "open"
	case 2:
		return "limited"
	default:
		return "unknown"
	}
}

// GroupType represents `groups_group_type` API object
type GroupType string

func (g *GroupType) MarshalJSON() ([]byte, error) {
	return objects.GetStringFromRange(string(*g), "group", "page", "event")
}

func (g *GroupType) GetName() string {
	return string(*g)
}

// GroupVideo represents `groups_group_video` API object
type GroupVideo int

func (g *GroupVideo) MarshalJSON() ([]byte, error) {
	return objects.GetIntFromRange(int(*g), 0, 2, true)
}

func (g *GroupVideo) GetName() string {
	switch *g {
	case 0:
		return "disabled"
	case 1:
		return "open"
	case 2:
		return "limited"
	default:
		return "unknown"
	}
}

// GroupWall represents `groups_group_wall` API object
type GroupWall int

func (g *GroupWall) MarshalJSON() ([]byte, error) {
	return objects.GetIntFromRange(int(*g), 0, 3, true)
}

func (g *GroupWall) GetName() string {
	switch *g {
	case 0:
		return "disabled"
	case 1:
		return "open"
	case 2:
		return "limited"
	case 3:
		return "closed"
	default:
		return "unknown"
	}
}

// GroupWiki represents `groups_group_wiki` API object
type GroupWiki int

func (g *GroupWiki) MarshalJSON() ([]byte, error) {
	return objects.GetIntFromRange(int(*g), 0, 2, true)
}

func (g *GroupWiki) GetName() string {
	switch *g {
	case 0:
		return "disabled"
	case 1:
		return "open"
	case 2:
		return "limited"
	default:
		return "unknown"
	}
}

// GroupXtrInvitedBy represents `groups_group_xtr_invited_by` API object
type GroupXtrInvitedBy struct {
	AdminLevel GroupXtrInvitedByAdminLevel `json:"admin_level"`
	ID         int                         `json:"id"`          // Community ID
	InvitedBy  int                         `json:"invited_by"`  // Inviter ID
	Admin      base.BoolInt                `json:"is_admin"`    // Information whether current user is manager
	Closed     base.BoolInt                `json:"is_closed"`   // Information whether community is closed
	IsMember   base.BoolInt                `json:"is_member"`   // Information whether current user is member
	Name       string                      `json:"name"`        // Community name
	Photo100   string                      `json:"photo_100"`   // URL of square photo of the community with 100 pixels in width
	Photo200   string                      `json:"photo_200"`   // URL of square photo of the community with 200 pixels in width
	Photo50    string                      `json:"photo_50"`    // URL of square photo of the community with 50 pixels in width
	ScreenName string                      `json:"screen_name"` // Domain of the community page
	Type       GroupXtrInvitedByType       `json:"type"`
}

// GroupXtrInvitedByAdminLevel represents `groups_group_xtr_invited_by_admin_level` API object
type GroupXtrInvitedByAdminLevel int

func (g *GroupXtrInvitedByAdminLevel) MarshalJSON() ([]byte, error) {
	return objects.GetIntFromRange(int(*g), 1, 3, true)
}

func (g *GroupXtrInvitedByAdminLevel) GetName() string {
	switch *g {
	case 1:
		return "moderator"
	case 2:
		return "editor"
	case 3:
		return "administrator"
	default:
		return "unknown"
	}
}

// GroupXtrInvitedByType represents `groups_group_xtr_invited_by_type` API object
type GroupXtrInvitedByType string

func (g *GroupXtrInvitedByType) MarshalJSON() ([]byte, error) {
	return objects.GetStringFromRange(string(*g), "group", "page", "event")
}

func (g *GroupXtrInvitedByType) GetName() string {
	return string(*g)
}

// GroupsArray represents `groups_groups_array` API object
type GroupsArray struct {
	Count int   `json:"count"` // Communities number
	Items []int `json:"items"` // Community ID
}

// LinksItem represents `groups_links_item` API object
type LinksItem struct {
	Description string `json:"desc"`       // Link description
	EditTitle   string `json:"edit_title"` // Information whether the link title can be edited
	ID          int    `json:"id"`         // Link ID
	Name        string `json:"name"`       // Link title
	Photo100    string `json:"photo_100"`  // URL of square photo of the community with 100 pixels in width
	Photo50     string `json:"photo_50"`   // URL of square photo of the community with 50 pixels in width
	Url         string `json:"url"`        // Link URL
}

// LongPollEvents represents `groups_long_poll_events` API object
type LongPollEvents struct {
	AudioNew             base.BoolInt `json:"audio_new"`
	BoardPostDelete      base.BoolInt `json:"board_post_delete"`
	BoardPostEdit        base.BoolInt `json:"board_post_edit"`
	BoardPostNew         base.BoolInt `json:"board_post_new"`
	BoardPostRestore     base.BoolInt `json:"board_post_restore"`
	GroupChangePhoto     base.BoolInt `json:"group_change_photo"`
	GroupChangeSettings  base.BoolInt `json:"group_change_settings"`
	GroupJoin            base.BoolInt `json:"group_join"`
	GroupLeave           base.BoolInt `json:"group_leave"`
	GroupOfficersEdit    base.BoolInt `json:"group_officers_edit"`
	LeadFormsNew         base.BoolInt `json:"lead_forms_new"`
	MarketCommentDelete  base.BoolInt `json:"market_comment_delete"`
	MarketCommentEdit    base.BoolInt `json:"market_comment_edit"`
	MarketCommentNew     base.BoolInt `json:"market_comment_new"`
	MarketCommentRestore base.BoolInt `json:"market_comment_restore"`
	MessageAllow         base.BoolInt `json:"message_allow"`
	MessageDeny          base.BoolInt `json:"message_deny"`
	MessageNew           base.BoolInt `json:"message_new"`
	MessageRead          base.BoolInt `json:"message_read"`
	MessageReply         base.BoolInt `json:"message_reply"`
	MessageTypingState   base.BoolInt `json:"message_typing_state"`
	MessagesEdit         base.BoolInt `json:"messages_edit"`
	PhotoCommentDelete   base.BoolInt `json:"photo_comment_delete"`
	PhotoCommentEdit     base.BoolInt `json:"photo_comment_edit"`
	PhotoCommentNew      base.BoolInt `json:"photo_comment_new"`
	PhotoCommentRestore  base.BoolInt `json:"photo_comment_restore"`
	PhotoNew             base.BoolInt `json:"photo_new"`
	PollVoteNew          base.BoolInt `json:"poll_vote_new"`
	UserBlock            base.BoolInt `json:"user_block"`
	UserUnblock          base.BoolInt `json:"user_unblock"`
	VideoCommentDelete   base.BoolInt `json:"video_comment_delete"`
	VideoCommentEdit     base.BoolInt `json:"video_comment_edit"`
	VideoCommentNew      base.BoolInt `json:"video_comment_new"`
	VideoCommentRestore  base.BoolInt `json:"video_comment_restore"`
	VideoNew             base.BoolInt `json:"video_new"`
	WallpostNew          base.BoolInt `json:"wall_post_new"`
	WallReplyDelete      base.BoolInt `json:"wall_reply_delete"`
	WallReplyEdit        base.BoolInt `json:"wall_reply_edit"`
	WallReplyNew         base.BoolInt `json:"wall_reply_new"`
	WallReplyRestore     base.BoolInt `json:"wall_reply_restore"`
	WallRepost           base.BoolInt `json:"wall_repost"`
}

// LongPollServer represents `groups_long_poll_server` API object
type LongPollServer struct {
	Key    string `json:"key"`    // Long Poll key
	Server string `json:"server"` // Long Poll server address
	TS     int    `json:"ts"`     // Number of the last event
}

// LongPollSettings represents `groups_long_poll_settings` API object
type LongPollSettings struct {
	ApiVersion string         `json:"api_version"` // API version used for the events
	Events     LongPollEvents `json:"events"`
	Enabled    bool           `json:"is_enabled"` // Shows whether Long Poll is enabled
}

// MarketInfo represents `groups_market_info` API object
type MarketInfo struct {
	ContactID    int             `json:"contact_id"` // Contact person ID
	Currency     market.Currency `json:"currency"`
	CurrencyText string          `json:"currency_text"` // Currency name
	Enabled      base.BoolInt    `json:"enabled"`
	MainAlbumID  int             `json:"main_album_id"` // Main market album ID
	PriceMax     int             `json:"price_max"`     // Maximum price
	PriceMin     int             `json:"price_min"`     // Minimum price
}

// MemberRole represents `groups_member_role` API object
type MemberRole struct {
	ID   int              `json:"id"` // User ID
	Role MemberRoleStatus `json:"role"`
}

// MemberRoleStatus represents `groups_member_role_status` API object.
// Contains user's credentials as community admin
type MemberRoleStatus string

func (m *MemberRoleStatus) MarshalJSON() ([]byte, error) {
	return objects.GetStringFromRange(string(*m), "moderator", "editor", "administrator", "creator")
}

func (m *MemberRoleStatus) GetName() string {
	return string(*m)
}

// MemberStatus represents `groups_member_status` API object
type MemberStatus struct {
	Member base.BoolInt `json:"member"`  // Information whether user is a member of the group
	UserID int          `json:"user_id"` // User ID
}

// MemberStatusFull represents `groups_member_status_full` API object
type MemberStatusFull struct {
	Member     base.BoolInt `json:"member"`     // Information whether user is a member of the group
	UserID     int          `json:"user_id"`    // User ID
	Invitation base.BoolInt `json:"invitation"` // Information whether user has been invited to the group
	Request    base.BoolInt `json:"request"`    // Information whether user has send request to the group
}

// OnlineStatus represents `groups_online_status` API object
type OnlineStatus struct {
	Minutes int              `json:"minutes"` // Estimated time of answer (for status = answer_mark)
	Status  OnlineStatusType `json:"status"`
}

// OnlineStatusType represents `groups_online_status_type` API object.
// Contains type of online status of group
type OnlineStatusType string

func (o *OnlineStatusType) MarshalJSON() ([]byte, error) {
	return objects.GetStringFromRange(string(*o), "none", "online", "answer_mark")
}

func (o *OnlineStatusType) GetName() string {
	return string(*o)
}

// OwnerXtrBanInfo represents `groups_owner_xtr_ban_info` API object
type OwnerXtrBanInfo struct {
	BanInfo BanInfo        `json:"ban_info"`
	Group   Group          `json:"group"` // Information about group if type = group
	Profile users.User     `json:"profile"`
	Type    XtrBanInfoType `json:"type"`
}

// XtrBanInfoType represents `groups_owner_xtr_ban_info_type` API object
type XtrBanInfoType string

func (x *XtrBanInfoType) MarshalJSON() ([]byte, error) {
	return objects.GetStringFromRange(string(*x), "group", "profile")
}

func (x *XtrBanInfoType) GetName() string {
	return string(*x)
}

// RoleOptions represents `groups_role_options` API object
type RoleOptions string

func (r *RoleOptions) MarshalJSON() ([]byte, error) {
	return objects.GetStringFromRange(string(*r), "moderator", "editor", "administrator", "creator")
}

func (r *RoleOptions) GetName() string {
	return string(*r)
}

// SubjectItem represents `groups_subject_item` API object
type SubjectItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// UserXtrRole represents `groups_user_xtr_role` API object
type UserXtrRole struct {
	users.UserFull
	Role RoleOptions
}
