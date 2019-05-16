package search

import (
	"fmt"
	"gitlab.com/Burmuley/go-vkapi/objects/base"
	"gitlab.com/Burmuley/go-vkapi/objects/groups"
	"gitlab.com/Burmuley/go-vkapi/objects/users"
)

/////////////////////////////////////////////////////////////
// Search related API objects	                           //
/////////////////////////////////////////////////////////////

// Hint represents `search_hint` API object
type Hint struct {
	Description string        `json:"description"` // Object description
	Global      base.BoolInt  `json:"global"`      // Information whether the object has been found globally
	Group       groups.Group  `json:"group"`
	Profile     users.UserMin `json:"profile"`
	Section     Section       `json:"section"`
	Type        Type          `json:"type"`
}

// Section represents `search_hint_section` API object
type Section string

func (s *Section) MarshalJSON() ([]byte, error) {
	switch *s {
	case "groups",
		"events",
		"publics",
		"correspondents",
		"people",
		"friends",
		"mutual_friends":
		return []byte(*s), nil
	}

	return []byte{}, fmt.Errorf("value is not in valid range")
}

func (s *Section) GetName() string {
	return string(*s)
}

// Type represents `search_hint_type` API object
type Type string

func (t *Type) MarshalJSON() ([]byte, error) {
	switch *t {
	case "group", "profile":
		return []byte(*t), nil
	}

	return []byte{}, fmt.Errorf("value is not in valid range")
}

func (t *Type) GetName() string {
	return string(*t)
}
