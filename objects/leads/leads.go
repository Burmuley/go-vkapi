package leads

import "gitlab.com/Burmuley/go-vkapi/objects/base"

// Checked represents `leads_checked` API objects
type Checked struct {
	Reason    string        `json:"reason"` // Reason why user can't start the lead
	Result    CheckedResult `json:"result"`
	SessionID int           `json:"sid"`        // Session ID
	StartLink string        `json:"start_link"` // URL user should open to start the lead
}

// CheckedResult represents `leads_checked_result` API object.
// Contains information whether user can start the lead
type CheckedResult bool

func (c *CheckedResult) GetName() string {
	if *c {
		return "true"
	}

	return "false"
}

// Complete represents `leads_complete` API object
type Complete struct {
	Cost     int `json:"cost"`  // Offer cost
	Limit    int `json:"limit"` // Offer limit
	Spent    int `json:"spent"` // Amount of spent votes
	Success  base.OkResponse
	TestMode base.BoolInt `json:"test_mode"` // Information whether test mode is enabled
}

// Entry represents `leads_entry` API object
type Entry struct {
	AppID     int          `json:"aid"`        // Application ID
	Comment   string       `json:"comment"`    // Comment text
	Date      int          `json:"date"`       // Date when the action has been started in Unixtime
	SessionID int          `json:"sid"`        // Session string ID
	StartDate int          `json:"start_date"` // Start date in Unixtime (for status=2)
	Status    int          `json:"status"`     // Action type
	TestMode  base.BoolInt `json:"test_mode"`  // Information whether test mode is enabled
	UserID    int          `json:"uid"`        // User ID
}

// Lead represents `leads_lead` API object
type Lead struct {
	Completed   int      `json:"completed"` // Completed offers number
	Cost        int      `json:"cost"`      // Offer cost
	Days        LeadDays `json:"days"`
	Impressions int      `json:"impressions"` // Impressions number
	Limit       int      `json:"limit"`       // Lead limit
	Spent       int      `json:"spent"`       // Amount of spent votes
	Started     int      `json:"started"`     // Started offers number
}

// LeadDays represents `leads_lead_days` API object
type LeadDays struct {
	Completed   int `json:"completed"`   // Completed offers number
	Impressions int `json:"impressions"` // Impressions number
	Spent       int `json:"spent"`       // Amount of spent votes
	Started     int `json:"started"`     // Started offers number
}

// Start represents `leads_start` API object
type Start struct {
	TestMode  base.BoolInt `json:"test_mode"` // Information whether test mode is enabled
	SessionID string       `json:"vk_sid"`    // Session data
}
