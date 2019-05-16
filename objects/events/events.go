package events

import "gitlab.com/Burmuley/go-vkapi/objects/groups"

/////////////////////////////////////////////////////////////
// Events related API objects	                           //
/////////////////////////////////////////////////////////////

// EventAttach represents `events_event_attach` API object
type EventAttach struct {
	Address      string                       `json:"address"`
	ButtonText   string                       `json:"button_text"`
	Friends      []int                        `json:"friends"`
	ID           int                          `json:"id"`
	Favorite     bool                         `json:"is_favorite"`
	MemberStatus groups.GroupFullMemberStatus `json:"member_status"`
	Text         string                       `json:"text"`
	Type         int                          `json:"type"`
}
