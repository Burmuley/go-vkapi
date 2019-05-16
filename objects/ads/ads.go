package ads

import (
	"gitlab.com/Burmuley/go-vkapi/objects"
	"gitlab.com/Burmuley/go-vkapi/objects/base"
)

/////////////////////////////////////////////////////////////
// ADS related API objects	                               //
/////////////////////////////////////////////////////////////

// AccessRole represents `ads_access_role` API object.
// Contains current user's role.
type AccessRole string

func (a *AccessRole) MarshalJSON() ([]byte, error) {
	return objects.GetStringFromRange(string(*a), "admin", "manager", "reports")
}

func (a *AccessRole) GetName() string {
	return string(*a)
}

// Accesses represents `ads_accesses` API object
type Accesses struct {
	ClientID int        `json:"client_id"` // Client ID
	Role     AccessRole `json:"role"`
}

// Account represents `ads_account` API object
type Account struct {
	AccessRole    AccessRole   `json:"access_role"`
	AccountID     int          `json:"account_id"`
	AccountStatus base.BoolInt `json:"account_status"` // Information whether account is active
	AccountType   AccountType  `json:"account_type"`
}

// AccountType represents `ads_account_type` API object
type AccountType string

func (a *AccountType) MarshalJSON() ([]byte, error) {
	return objects.GetStringFromRange(string(*a), "general", "agency")
}

func (a *AccountType) GetName() string {
	return string(*a)
}

// Ad represents `ads_ad` API object
type Ad struct {
	AdFormat              int          `json:"ad_format"` // Ad format
	AdPlatform            string       `json:"ad_platform"`
	AllLimit              int          `json:"all_limit"` // Total limit
	Approved              AdApproved   `json:"approved"`
	CampaignID            int          `json:"campaign_id"`  // Campaign ID
	Category1ID           int          `json:"category1_id"` // Category ID
	Category2ID           int          `json:"category2_id"` // Additional category ID
	CostType              AdCostType   `json:"cost_type"`
	CPC                   int          `json:"cpc"`                    // Cost of a click, kopecks
	CPM                   int          `json:"cpm"`                    // Cost of 1000 impressions, kopecks
	DisclaimerMedical     base.BoolInt `json:"disclaimer_medical"`     // Information whether disclaimer is enabled
	DisclaimerSpecialist  base.BoolInt `json:"disclaimer_specialist"`  // Information whether disclaimer is enabled
	DisclaimerSupplements base.BoolInt `json:"disclaimer_supplements"` // Information whether disclaimer is enabled
	ID                    int          `json:"id"`                     // Ad ID
	ImpressionsLimit      int          `json:"impressions_limit"`      // Impressions limit
	ImpressionsLimited    base.BoolInt `json:"impressions_limited"`    // Information whether impressions are limited
	Name                  string       `json:"name"`                   // Ad title
	Status                AdStatus     `json:"status"`
	Video                 base.BoolInt `json:"video"` // Information whether the ad is a video
}

// AdApproved represents `ads_ad_approved` API object.
// Contains review status.
type AdApproved int

func (a *AdApproved) MarshalJSON() ([]byte, error) {
	return objects.GetIntFromRange(int(*a), 0, 3, true)
}

func (a *AdApproved) GetName() string {
	switch *a {
	case 0:
		return "not moderated"
	case 1:
		return "pending moderation"
	case 2:
		return "approved"
	case 3:
		return "rejected"
	default:
		return "unknown"
	}
}

// AdCostType represents `ads_ad_cost_type` API object
// Contains cost type.
type AdCostType int

func (a *AdCostType) MarshalJSON() ([]byte, error) {
	return objects.GetIntFromRange(int(*a), 0, 1, true)
}

func (a *AdCostType) GetName() string {
	switch *a {
	case 0:
		return "per clicks"
	case 1:
		return "per impressions"
	default:
		return "unknown"
	}
}

// AdLayout represents `ads_ad_layout` API object
type AdLayout struct {
	AdFormat    int              `json:"ad_format"`   // Ad format
	CampaignID  int              `json:"campaign_id"` // Campaign ID
	CostType    AdLayoutCostType `json:"cost_type"`
	Description string           `json:"description"`  // Ad description
	ID          int              `json:"id"`           // Ad ID
	ImageSrc    string           `json:"image_src"`    // Image URL
	ImageSrc2x  string           `json:"image_src_2x"` // URL of the preview image in double size
	LinkDomain  string           `json:"link_domain"`  // Domain of advertised object
	LinkUrl     string           `json:"link_url"`     // URL of advertised object
	PreviewLink string           `json:"preview_link"` // link to preview an ad as it is shown on the website
	Title       string           `json:"title"`        // Ad title
	Video       base.BoolInt     `json:"video"`        // Information whether the ad is a video
}

// AdLayoutCostType represents `ads_ad_layout_cost_type` API object
type AdLayoutCostType int

func (a *AdLayoutCostType) MarshalJSON() ([]byte, error) {
	return objects.GetIntFromRange(int(*a), 0, 1, true)
}

func (a *AdLayoutCostType) GetName() string {
	switch *a {
	case 0:
		return "per clicks"
	case 1:
		return "per impressions"
	default:
		return "unknown"
	}
}

// AdStatus represents `ads_ad_status` API object.
// Contains Ad status.
type AdStatus int

func (a *AdStatus) MarshalJSON() ([]byte, error) {
	return objects.GetIntFromRange(int(*a), 0, 2, true)
}

func (a *AdStatus) GetName() string {
	switch *a {
	case 0:
		return "stopped"
	case 1:
		return "started"
	case 3:
		return "deleted"
	default:
		return "unknown"
	}
}

// Campaign represents `ads_campaign` API object
type Campaign struct {
	AllLimit  string         `json:"all_limit"`  // Campaign's total limit, rubles
	DayLimit  string         `json:"day_limit"`  // Campaign's day limit, rubles
	ID        int            `json:"id"`         // Campaign ID
	Name      string         `json:"name"`       // Campaign title
	StartTime int            `json:"start_time"` // Campaign start time, as Unixtime
	StopTime  int            `json:"stop_time"`  // Campaign stop time, as Unixtime
	Status    CampaignStatus `json:"status"`
	Type      CampaignType   `json:"type"`
}

// CampaignStatus represents `ads_campaign_status` API object
type CampaignStatus int

func (c *CampaignStatus) MarshalJSON() ([]byte, error) {
	return objects.GetIntFromRange(int(*c), 0, 2, true)
}

func (c *CampaignStatus) GetName() string {
	switch *c {
	case 0:
		return "stopped"
	case 1:
		return "started"
	case 3:
		return "deleted"
	default:
		return "unknown"
	}
}

// CampaignType represents `ads_campaign_type` API object
type CampaignType string

func (c *CampaignType) GetName() string {
	return string(*c)
}

// Category represents `ads_category` API object
type Category struct {
	ID            int                   `json:"id"`   // Category ID
	Name          string                `json:"name"` // Category name
	Subcategories []base.ObjectWithName `json:"subcategories"`
}

// Client represents `ads_client` API object
type Client struct {
	AllLimit string `json:"all_limit"` // Client's total limit, rubles
	DayLimit string `json:"day_limit"` // Client's day limit, rubles
	ID       int    `json:"id"`        // Client ID
	Name     string `json:"name"`      // Client name
}

// Criteria represents `ads_criteria` API object
type Criteria struct {
}

// CriteriaSex represents `ads_criteria_sex` API object
type CriteriaSex int

func (c *CriteriaSex) MarshalJSON() ([]byte, error) {
	return objects.GetIntFromRange(int(*c), 0, 2, true)
}

func (c *CriteriaSex) GetName() string {
	switch *c {
	case 0:
		return "any"
	case 1:
		return "male"
	case 2:
		return "female"
	default:
		return "unknown"
	}
}

// DemoStats represents `ads_demo_stats` API object
type DemoStats struct {
	ID    int             `json:"id"` // Object ID
	Stats DemoStatsFormat `json:"stats"`
	Type  ObjectType      `json:"type"`
}

// DemoStatsFormat represents `ads_demostats_format` API object
type DemoStatsFormat struct {
	Age     []StatsAge    `json:"age"`
	Cities  []StatsCities `json:"cities"`
	Day     string        `json:"day"`     // Day as YYYY-MM-DD
	Month   string        `json:"month"`   // Month as YYYY-MM
	Overall int           `json:"overall"` // 1 if period=overall
	Sex     []StatsSex    `json:"sex"`
	SexAge  []StatsSexAge `json:"sex_age"`
}

// FloodStats represents `ads_flood_stats` API object
type FloodStats struct {
	Left    int `json:"left"`    // Requests left
	Refresh int `json:"refresh"` // Time to refresh in seconds
}

// LinkStatus represents `ads_link_status` API object
type LinkStatus struct {
	Description string `json:"description"`  // Reject reason
	RedirectUrl string `json:"redirect_url"` // URL
	Status      string `json:"status"`       // Link status
}

// ObjectType represents `ads_object_type` API object
type ObjectType string

func (o *ObjectType) MarshalJSON() ([]byte, error) {
	return objects.GetStringFromRange(string(*o), "ad", "campaign", "client", "office")
}

func (o *ObjectType) GetName() string {
	return string(*o)
}

// Paragraphs represents `ads_paragraphs` API object
type Paragraphs struct {
	Paragraph string `json:"paragraph"` // Rules paragraph
}

// PromotedPostReach represents `ads_promoted_post_reach` API object
type PromotedPostReach struct {
	Hide             int `json:"hide"`              // Hides amount
	ID               int `json:"id"`                // Object ID from 'ids' parameter
	JoinGroup        int `json:"join_group"`        // Community joins
	Links            int `json:"links"`             // Link clicks
	ReachSubscribers int `json:"reach_subscribers"` // Subscribers reach
	ReachTotal       int `json:"reach_total"`       // Total reach
	Report           int `json:"report"`            // Reports amount
	ToGroup          int `json:"to_group"`          // Community clicks
	Unsubscribe      int `json:"unsubscribe"`       // 'Unsubscribe' events amount
	VideoViews100p   int `json:"video_views_100p"`  // Video views for 100 percent
	VideoViews25p    int `json:"video_views_25p"`   // Video views for 25 percent
	VideoViews50p    int `json:"video_views_50p"`   // Video views for 50 percent
	VideoViews75p    int `json:"video_views_75p"`   // Video views for 75 percent
	VideoViews3s     int `json:"video_views_3s"`    // Video views for 3 seconds
	VideoViewsStart  int `json:"video_views_start"` // Video starts
}

// RejectReason represents `ads_reject_reason` API object
type RejectReason struct {
	Comment string  `json:"comment"` // Comment text
	Rules   []Rules `json:"rules"`
}

// Rules represents `ads_rules` API object
type Rules struct {
	Paragraphs []Paragraphs `json:"paragraphs"`
	Title      string       `json:"title"` // Comment
}

// Stats represents `ads_stats` API object
type Stats struct {
	ID    int         `json:"id"` // Object ID
	Stats StatsFormat `json:"stats"`
	Type  ObjectType  `json:"type"`
}

// StatsAge represents `ads_stats_age` API object
type StatsAge struct {
	ClicksRate      float64 `json:"clicks_rate"`      // Clicks rate
	ImpressionsRate float64 `json:"impressions_rate"` // Impressions rate
	Value           string  `json:"value"`            // Age interval
}

// StatsCities represents `ads_stats_cities` API object
type StatsCities struct {
	ClicksRate      float64 `json:"clicks_rate"`      // Clicks rate
	ImpressionsRate float64 `json:"impressions_rate"` // Impressions rate
	Value           string  `json:"value"`            // Age interval
	Name            string  `json:"name"`             // City name
}

// StatsFormat represents `ads_stats_format` API object
type StatsFormat struct {
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

// StatsSex represents `ads_stats_sex` API object
type StatsSex struct {
	ClicksRate      float64 `json:"clicks_rate"`      // Clicks rate
	ImpressionsRate float64 `json:"impressions_rate"` // Impressions rate
	Value           string  `json:"value"`            // Age interval
}

// StatsSexAge represents `ads_stats_sex_age` API object
type StatsSexAge struct {
	ClicksRate      float64 `json:"clicks_rate"`      // Clicks rate
	ImpressionsRate float64 `json:"impressions_rate"` // Impressions rate
	Value           string  `json:"value"`            // Age interval
}

// StatsSexValue represents `ads_stats_sex_value` API object
type StatsSexValue string

func (s *StatsSexValue) MarshalJSON() ([]byte, error) {
	return objects.GetStringFromRange(string(*s), "f", "m")
}

func (s *StatsSexValue) GetName() string {
	switch *s {
	case "f":
		return "female"
	case "m":
		return "male"
	default:
		return "unknown"
	}
}

// TargSettings represents `ads_targ_settings` API object
type TargSettings struct {
	Criteria
	ID         int `json:"id"`          // Ad ID
	CampaignID int `json:"campaign_id"` // Campaign ID
}

// TargStats represents `ads_targ_stats` API object
type TargStats struct {
	AudienceCount  int `json:"audience_count"`  // Audience
	RecommendedCPC int `json:"recommended_cpc"` // Recommended CPC value
	RecommendedCPM int `json:"recommended_cpm"` // Recommended CPM value
}

// TargSuggestions represents `ads_targ_suggestions` API object
type TargSuggestions struct {
	ID   int    `json:"id"`   // Object ID
	Name string `json:"name"` // Object name
}

// TargSuggestionsCities represents `ads_targ_suggestions_cities` API object
type TargSuggestionsCities struct {
	ID     int    `json:"id"`     // Object ID
	Name   string `json:"name"`   // Object name
	Parent string `json:"parent"` // Parent object
}

// TargSuggestionsRegions represents `ads_targ_suggestions_regions` API object
type TargSuggestionsRegions struct {
	ID   int    `json:"id"`   // Object ID
	Name string `json:"name"` // Object name
	Type string `json:"type"` // Object type
}

// TargSuggestionsSchools represents `ads_targ_suggestions_schools` API object
type TargSuggestionsSchools struct {
	Description string                     `json:"desc"`   // Full school title
	ID          int                        `json:"id"`     // School ID
	Name        string                     `json:"name"`   // School name
	Parent      string                     `json:"parent"` // City name
	Type        TargSuggestionsSchoolsType `json:"type"`   // School type
}

// TargSuggestionsSchoolsType represents `ads_targ_suggestions_schools_type` API object
type TargSuggestionsSchoolsType string

func (t *TargSuggestionsSchoolsType) MarshalJSON() ([]byte, error) {
	return objects.GetStringFromRange(string(*t), "school", "university", "faculty", "chair")
}

func (t *TargSuggestionsSchoolsType) GetName() string {
	return string(*t)
}

// TargetGroup represents `ads_target_group` API object
type TargetGroup struct {
	AudienceCount int    `json:"audience_count"` // Audience
	Domain        string `json:"domain"`         // Site domain
	ID            int    `json:"id"`             // Group ID
	Lifetime      int    `json:"lifetime"`       // Number of days for user to be in group
	Name          string `json:"name"`           // Group name
	Pixel         string `json:"pixel"`          // Pixel code
}

// Users represents `ads_users` API object
type Users struct {
	Accesses []Accesses `json:"accesses"`
	ID       int        `json:"user_id"` // User ID
}
