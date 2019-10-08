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
// `ads` group of objects
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

// AdsAccountType type represents `ads_account_type` API object
type AdsAccountType string // Account type

// AdsAd type represents `ads_ad` API object
type AdsAd struct {
	AdFormat              int           `json:"ad_format"`   // Ad format
	AdPlatform            multiple      `json:"ad_platform"` // Ad platform
	AllLimit              int           `json:"all_limit"`   // Total limit
	Approved              AdsAdApproved `json:"approved"`
	CampaignId            int           `json:"campaign_id"`  // Campaign ID
	Category1Id           int           `json:"category1_id"` // Category ID
	Category2Id           int           `json:"category2_id"` // Additional category ID
	CostType              AdsAdCostType `json:"cost_type"`
	Cpa                   int           `json:"cpa"`                    // Cost of an action, kopecks
	Cpc                   int           `json:"cpc"`                    // Cost of a click, kopecks
	Cpm                   int           `json:"cpm"`                    // Cost of 1000 impressions, kopecks
	DisclaimerMedical     BaseBoolInt   `json:"disclaimer_medical"`     // Information whether disclaimer is enabled
	DisclaimerSpecialist  BaseBoolInt   `json:"disclaimer_specialist"`  // Information whether disclaimer is enabled
	DisclaimerSupplements BaseBoolInt   `json:"disclaimer_supplements"` // Information whether disclaimer is enabled
	Id                    int           `json:"id"`                     // Ad ID
	ImpressionsLimit      int           `json:"impressions_limit"`      // Impressions limit
	ImpressionsLimited    BaseBoolInt   `json:"impressions_limited"`    // Information whether impressions are limited
	Name                  string        `json:"name"`                   // Ad title
	Status                AdsAdStatus   `json:"status"`
	Video                 BaseBoolInt   `json:"video"` // Information whether the ad is a video
}

// AdsCampaignStatus type represents `ads_campaign_status` API object
type AdsCampaignStatus int // Campaign status

// AdsCategory type represents `ads_category` API object
type AdsCategory struct {
	Id            int                  `json:"id"`   // Category ID
	Name          string               `json:"name"` // Category name
	Subcategories []BaseObjectWithName `json:"subcategories"`
}

// AdsAdCostType type represents `ads_ad_cost_type` API object
type AdsAdCostType int // Cost type

// AdsAdStatus type represents `ads_ad_status` API object
type AdsAdStatus int // Ad atatus

// AdsCampaign type represents `ads_campaign` API object
type AdsCampaign struct {
	AllLimit  string            `json:"all_limit"`  // Campaign's total limit, rubles
	DayLimit  string            `json:"day_limit"`  // Campaign's day limit, rubles
	Id        int               `json:"id"`         // Campaign ID
	Name      string            `json:"name"`       // Campaign title
	StartTime int               `json:"start_time"` // Campaign start time, as Unixtime
	Status    AdsCampaignStatus `json:"status"`
	StopTime  int               `json:"stop_time"` // Campaign stop time, as Unixtime
	Type      AdsCampaignType   `json:"type"`
}

// AdsLinkStatus type represents `ads_link_status` API object
type AdsLinkStatus struct {
	Description string `json:"description"`  // Reject reason
	RedirectUrl string `json:"redirect_url"` // URL
	Status      string `json:"status"`       // Link status
}

// AdsRejectReason type represents `ads_reject_reason` API object
type AdsRejectReason struct {
	Comment string     `json:"comment"` // Comment text
	Rules   []AdsRules `json:"rules"`
}

// AdsParagraphs type represents `ads_paragraphs` API object
type AdsParagraphs struct {
	Paragraph string `json:"paragraph"` // Rules paragraph
}

// AdsAccesses type represents `ads_accesses` API object
type AdsAccesses struct {
	ClientId string        `json:"client_id"` // Client ID
	Role     AdsAccessRole `json:"role"`
}

// AdsCriteriaSex type represents `ads_criteria_sex` API object
type AdsCriteriaSex int // Sex

// AdsTargSuggestionsCities type represents `ads_targ_suggestions_cities` API object
type AdsTargSuggestionsCities struct {
	Id     int    `json:"id"`     // Object ID
	Name   string `json:"name"`   // Object name
	Parent string `json:"parent"` // Parent object
}

// AdsStatsSexValue type represents `ads_stats_sex_value` API object
type AdsStatsSexValue string // Sex

// AdsAdLayout type represents `ads_ad_layout` API object
type AdsAdLayout struct {
	AdFormat    int           `json:"ad_format"`   // Ad format
	CampaignId  int           `json:"campaign_id"` // Campaign ID
	CostType    AdsAdCostType `json:"cost_type"`
	Description string        `json:"description"`  // Ad description
	Id          int           `json:"id"`           // Ad ID
	ImageSrc    string        `json:"image_src"`    // Image URL
	ImageSrc2x  string        `json:"image_src_2x"` // URL of the preview image in double size
	LinkDomain  string        `json:"link_domain"`  // Domain of advertised object
	LinkUrl     string        `json:"link_url"`     // URL of advertised object
	PreviewLink multiple      `json:"preview_link"` // link to preview an ad as it is shown on the website
	Title       string        `json:"title"`        // Ad title
	Video       BaseBoolInt   `json:"video"`        // Information whether the ad is a video
}

// AdsStatsFormat type represents `ads_stats_format` API object
type AdsStatsFormat struct {
	Clicks          int    `json:"clicks"`            // Clicks number
	Day             string `json:"day"`               // Day as YYYY-MM-DD
	Impressions     int    `json:"impressions"`       // Impressions number
	JoinRate        int    `json:"join_rate"`         // Events number
	Month           string `json:"month"`             // Month as YYYY-MM
	Overall         int    `json:"overall"`           // 1 if period=overall
	Reach           int    `json:"reach"`             // Reach
	Spent           int    `json:"spent"`             // Spent funds
	VideoClicksSite int    `json:"video_clicks_site"` // Clickthoughs to the advertised site
	VideoViews      int    `json:"video_views"`       // Video views number
	VideoViewsFull  int    `json:"video_views_full"`  // Video views (full video)
	VideoViewsHalf  int    `json:"video_views_half"`  // Video views (half of video)
}

// AdsRules type represents `ads_rules` API object
type AdsRules struct {
	Paragraphs []AdsParagraphs `json:"paragraphs"`
	Title      string          `json:"title"` // Comment
}

// AdsAccessRole type represents `ads_access_role` API object
type AdsAccessRole string // Current user's role

// AdsAdApproved type represents `ads_ad_approved` API object
type AdsAdApproved int // Review status

// AdsDemoStats type represents `ads_demo_stats` API object
type AdsDemoStats struct {
	Id    int                `json:"id"` // Object ID
	Stats AdsDemostatsFormat `json:"stats"`
	Type  AdsObjectType      `json:"type"`
}

// AdsTargSettings type represents `ads_targ_settings` API object
type AdsTargSettings struct {
	AdsCriteria AdsCriteria `json:"AdsCriteria"`
	CampaignId  int         `json:"campaign_id"` // Campaign ID
	Id          int         `json:"id"`          // Ad ID
}

// AdsClient type represents `ads_client` API object
type AdsClient struct {
	AllLimit string `json:"all_limit"` // Client's total limit, rubles
	DayLimit string `json:"day_limit"` // Client's day limit, rubles
	Id       int    `json:"id"`        // Client ID
	Name     string `json:"name"`      // Client name
}

// AdsCriteria type represents `ads_criteria` API object
type AdsCriteria struct {
	AgeFrom              int                `json:"age_from"`               // Age from
	AgeTo                int                `json:"age_to"`                 // Age to
	Apps                 string             `json:"apps"`                   // Apps IDs
	AppsNot              string             `json:"apps_not"`               // Apps IDs to except
	Birthday             int                `json:"birthday"`               // Days to birthday
	Cities               string             `json:"cities"`                 // Cities IDs
	CitiesNot            string             `json:"cities_not"`             // Cities IDs to except
	Country              int                `json:"country"`                // Country ID
	Districts            string             `json:"districts"`              // Districts IDs
	Groups               string             `json:"groups"`                 // Communities IDs
	InterestCategories   string             `json:"interest_categories"`    // Interests categories IDs
	Interests            string             `json:"interests"`              // Interests
	Paying               BaseBoolInt        `json:"paying"`                 // Information whether the user has proceeded VK payments before
	Positions            string             `json:"positions"`              // Positions IDs
	Religions            string             `json:"religions"`              // Religions IDs
	RetargetingGroups    string             `json:"retargeting_groups"`     // Retargeting groups IDs
	RetargetingGroupsNot string             `json:"retargeting_groups_not"` // Retargeting groups IDs to except
	SchoolFrom           int                `json:"school_from"`            // School graduation year from
	SchoolTo             int                `json:"school_to"`              // School graduation year to
	Schools              string             `json:"schools"`                // Schools IDs
	Sex                  AdsCriteriaSex     `json:"sex"`
	Stations             string             `json:"stations"`      // Stations IDs
	Statuses             string             `json:"statuses"`      // Relationship statuses
	Streets              string             `json:"streets"`       // Streets IDs
	Travellers           BasePropertyExists `json:"travellers"`    // Travellers only
	UniFrom              int                `json:"uni_from"`      // University graduation year from
	UniTo                int                `json:"uni_to"`        // University graduation year to
	UserBrowsers         string             `json:"user_browsers"` // Browsers
	UserDevices          string             `json:"user_devices"`  // Devices
	UserOs               string             `json:"user_os"`       // Operating systems
}

// AdsCampaignType type represents `ads_campaign_type` API object
type AdsCampaignType string // Campaign type

// AdsTargSuggestionsSchoolsType type represents `ads_targ_suggestions_schools_type` API object
type AdsTargSuggestionsSchoolsType string // School type

// AdsUsers type represents `ads_users` API object
type AdsUsers struct {
	Accesses []AdsAccesses `json:"accesses"`
	UserId   int           `json:"user_id"` // User ID
}

// AdsAccount type represents `ads_account` API object
type AdsAccount struct {
	AccessRole    AdsAccessRole  `json:"access_role"`
	AccountId     int            `json:"account_id"`     // Account ID
	AccountStatus BaseBoolInt    `json:"account_status"` // Information whether account is active
	AccountType   AdsAccountType `json:"account_type"`
}

// AdsStatsAge type represents `ads_stats_age` API object
type AdsStatsAge struct {
	ClicksRate      float64 `json:"clicks_rate"`      // Clicks rate
	ImpressionsRate float64 `json:"impressions_rate"` // Impressions rate
	Value           string  `json:"value"`            // Age interval
}

// AdsTargSuggestionsRegions type represents `ads_targ_suggestions_regions` API object
type AdsTargSuggestionsRegions struct {
	Id   int    `json:"id"`   // Object ID
	Name string `json:"name"` // Object name
	Type string `json:"type"` // Object type
}

// AdsStatsSexAge type represents `ads_stats_sex_age` API object
type AdsStatsSexAge struct {
	ClicksRate      float64 `json:"clicks_rate"`      // Clicks rate
	ImpressionsRate float64 `json:"impressions_rate"` // Impressions rate
	Value           string  `json:"value"`            // Sex and age interval
}

// AdsStats type represents `ads_stats` API object
type AdsStats struct {
	Id    int            `json:"id"` // Object ID
	Stats AdsStatsFormat `json:"stats"`
	Type  AdsObjectType  `json:"type"`
}

// AdsTargetGroup type represents `ads_target_group` API object
type AdsTargetGroup struct {
	AudienceCount int    `json:"audience_count"` // Audience
	Domain        string `json:"domain"`         // Site domain
	Id            int    `json:"id"`             // Group ID
	Lifetime      int    `json:"lifetime"`       // Number of days for user to be in group
	Name          string `json:"name"`           // Group name
	Pixel         string `json:"pixel"`          // Pixel code
}

// AdsStatsSex type represents `ads_stats_sex` API object
type AdsStatsSex struct {
	ClicksRate      float64          `json:"clicks_rate"`      // Clicks rate
	ImpressionsRate float64          `json:"impressions_rate"` // Impressions rate
	Value           AdsStatsSexValue `json:"value"`
}

// AdsTargSuggestionsSchools type represents `ads_targ_suggestions_schools` API object
type AdsTargSuggestionsSchools struct {
	Desc   string                        `json:"desc"`   // Full school title
	Id     int                           `json:"id"`     // School ID
	Name   string                        `json:"name"`   // School title
	Parent string                        `json:"parent"` // City name
	Type   AdsTargSuggestionsSchoolsType `json:"type"`
}

// AdsTargStats type represents `ads_targ_stats` API object
type AdsTargStats struct {
	AudienceCount  int     `json:"audience_count"`  // Audience
	RecommendedCpc float64 `json:"recommended_cpc"` // Recommended CPC value
	RecommendedCpm float64 `json:"recommended_cpm"` // Recommended CPM value
}

// AdsStatsCities type represents `ads_stats_cities` API object
type AdsStatsCities struct {
	ClicksRate      float64 `json:"clicks_rate"`      // Clicks rate
	ImpressionsRate float64 `json:"impressions_rate"` // Impressions rate
	Name            string  `json:"name"`             // City name
	Value           int     `json:"value"`            // City ID
}

// AdsPromotedPostReach type represents `ads_promoted_post_reach` API object
type AdsPromotedPostReach struct {
	Hide             int `json:"hide"`              // Hides amount
	Id               int `json:"id"`                // Object ID from 'ids' parameter
	JoinGroup        int `json:"join_group"`        // Community joins
	Links            int `json:"links"`             // Link clicks
	ReachSubscribers int `json:"reach_subscribers"` // Subscribers reach
	ReachTotal       int `json:"reach_total"`       // Total reach
	Report           int `json:"report"`            // Reports amount
	ToGroup          int `json:"to_group"`          // Community clicks
	Unsubscribe      int `json:"unsubscribe"`       // 'Unsubscribe' events amount
	VideoViews100p   int `json:"video_views_100p"`  // Video views for 100 percent
	VideoViews25p    int `json:"video_views_25p"`   // Video views for 25 percent
	VideoViews3s     int `json:"video_views_3s"`    // Video views for 3 seconds
	VideoViews50p    int `json:"video_views_50p"`   // Video views for 50 percent
	VideoViews75p    int `json:"video_views_75p"`   // Video views for 75 percent
	VideoViewsStart  int `json:"video_views_start"` // Video starts
}

// AdsObjectType type represents `ads_object_type` API object
type AdsObjectType string // Object type

// AdsTargSuggestions type represents `ads_targ_suggestions` API object
type AdsTargSuggestions struct {
	Id   int    `json:"id"`   // Object ID
	Name string `json:"name"` // Object name
}

// AdsDemostatsFormat type represents `ads_demostats_format` API object
type AdsDemostatsFormat struct {
	Age     []AdsStatsAge    `json:"age"`
	Cities  []AdsStatsCities `json:"cities"`
	Day     string           `json:"day"`     // Day as YYYY-MM-DD
	Month   string           `json:"month"`   // Month as YYYY-MM
	Overall int              `json:"overall"` // 1 if period=overall
	Sex     []AdsStatsSex    `json:"sex"`
	SexAge  []AdsStatsSexAge `json:"sex_age"`
}

// AdsFloodStats type represents `ads_flood_stats` API object
type AdsFloodStats struct {
	Left    int `json:"left"`    // Requests left
	Refresh int `json:"refresh"` // Time to refresh in seconds
}
