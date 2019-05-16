package pages

import "gitlab.com/Burmuley/go-vkapi/objects/base"

/////////////////////////////////////////////////////////////
// Pages related API objects	                           //
/////////////////////////////////////////////////////////////

// PrivacySettings represents `pages_privacy_settings` API object
type PrivacySettings int

func (p *PrivacySettings) GetName() string {
	switch *p {
	case 0:
		return "community managers only"
	case 1:
		return "community members only"
	case 2:
		return "everyone"
	default:
		return "unknown"
	}
}

// Wikipage represents `pages_wikipage` API object
type Wikipage struct {
	CreatorID   int             `json:"creator_id"`   // Page creator ID
	CreatorName string          `json:"creator_name"` // Page creator name
	EditorID    int             `json:"editor_id"`    // Last editor ID
	EditorName  string          `json:"editor_name"`  // Last editor name
	GroupID     int             `json:"group_id"`     // Community ID
	ID          int             `json:"id"`           // Page ID
	Title       string          `json:"title"`        // Page title
	Views       int             `json:"views"`        // Views number
	WhoCanEdit  PrivacySettings `json:"who_can_edit"` // Edit settings of the page
	WhoCanView  PrivacySettings `json:"who_can_view"` // View settings of the page
}

type WikipageFull struct {
	Created              int             `json:"created"`                      // Date when the page has been created in Unixtime
	CreatorID            int             `json:"creator_id"`                   // Page creator ID
	CurUserCanEdit       base.BoolInt    `json:"current_user_can_edit"`        // Information whether current user can edit the page
	CurUserCanEditAccess base.BoolInt    `json:"current_user_can_edit_access"` // Information whether current user can edit the page access settings
	EditorID             int             `json:"editor_id"`                    // Last editor ID
	Edited               int             `json:"edited"`                       // Date when the page has been edited in Unixtime
	GroupID              int             `json:"group_id"`                     // Community ID
	HTML                 string          `json:"html"`                         // Page content, HTML
	ID                   int             `json:"id"`                           // Page ID
	Source               string          `json:"source"`                       // Page content, wiki
	Title                string          `json:"title"`                        // Page title
	Views                int             `json:"views"`                        // Views number
	ViewUrl              string          `json:"view_url"`                     // URL of the page preview
	WhoCanEdit           PrivacySettings `json:"who_can_edit"`                 // Edit settings of the page
	WhoCanView           PrivacySettings `json:"who_can_view"`                 // View settings of the page
}

// WikipageVersion represents `pages_wikipage_version` API object
type WikipageVersion struct {
	Edited     int    `json:"edited"`      // Date when the page has been edited in Unixtime
	EditorID   int    `json:"editor_id"`   // Last editor ID
	EditorName string `json:"editor_name"` // Last editor name
	ID         int    `json:"id"`          // Version ID
	Length     int    `json:"length"`      // Page size in bytes
}
