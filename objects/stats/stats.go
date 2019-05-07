package stats

/////////////////////////////////////////////////////////////
// Stats related API objects	                           //
/////////////////////////////////////////////////////////////

// Activity represents `stats_activity` API object
type Activity struct {
	Comments     int `json:"comments"`     // Comments number
	Copies       int `json:"copies"`       // Reposts number
	Hidden       int `json:"hidden"`       // Hidden from news count
	Likes        int `json:"likes"`        // Likes number
	Subscribed   int `json:"subscribed"`   // New subscribers count
	Unsubscribed int `json:"unsubscribed"` // Unsubscribed count
}

// City represents `stats_city` API object
type City struct {
	Count int    `json:"count"` // Visitors number
	Name  string `json:"name"`  // City name
	ID    int    `json:"value"` // City ID
}

// Country represents `stats_country` API object
type Country struct {
	Count int    `json:"count"` // Visitors number
	Name  string `json:"name"`  // Country name
	ID    int    `json:"value"` // Country ID
	Code  string `json:"code"`  // Country code
}

// Period represents `stats_period` API object
type Period struct {
	Activity Activity `json:"activity"`
	From     int      `json:"period_from"` // Unix timestamp
	To       int      `json:"period_to"`   // Unix timestamp
	Reach    Reach    `json:"reach"`
	Visitors Views    `json:"visitors"`
}

// Reach represents `stats_reach` API object
type Reach struct {
	Age              []SexAge  `json:"age"`
	Cities           []City    `json:"cities"`
	Countries        []Country `json:"countries"`
	MobileReach      int       `json:"mobile_reach"`      // Reach count from mobile devices
	Reach            int       `json:"reach"`             // Reach count
	ReachSubscribers int       `json:"reach_subscribers"` // Subscribers reach count
	Sex              []SexAge  `json:"sex"`
	SexAge           []SexAge  `json:"sex_age"`
}

// SexAge represents `stats_sex_age` API object
type SexAge struct {
	Count int `json:"count"` // Visitors number
	Value int `json:"value"` // Sex/age value
}

// Views represents `stats_views` API object
type Views struct {
	Age         []SexAge  `json:"age"`
	Cities      []City    `json:"cities"`
	Countries   []Country `json:"countries"`
	MobileViews int       `json:"mobile_views"` // Views count from mobile devices
	Views       int       `json:"views"`        // Views number
	Visitors    int       `json:"visitors"`     // Visitors number
	Sex         []SexAge  `json:"sex"`
	SexAge      []SexAge  `json:"sex_age"`
}

// WallpostStat represents `stats_wallpost_stat` API object
type WallpostStat struct {
	Hide             int `json:"hide"`              // Hidings number
	JoinGroup        int `json:"join_group"`        // People have joined the group
	Links            int `json:"links"`             // Link clickthrough
	ReachSubscribers int `json:"reach_subscribers"` // Subscribers reach
	ReachTotal       int `json:"reach_total"`       // Total reach
	Report           int `json:"report"`            // Reports number
	ToGroup          int `json:"to_group"`          // Clickthrough to community
	Unsubscribe      int `json:"unsubscribe"`       // Unsubscribed members
}
