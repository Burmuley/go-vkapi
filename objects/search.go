package objects

/////////////////////////////////////////////////////////////
// Search related API objects	                           //
/////////////////////////////////////////////////////////////

// Hint represents `search_hint` API object
type SearchHint struct {
	Description string        `json:"description"` // Object description
	Global      BaseBoolInt   `json:"global"`      // Information whether the object has been found globally
	Group       GroupsGroup   `json:"group"`
	Profile     UsersUserMin  `json:"profile"`
	Section     SearchSection `json:"section"`
	Type        SearchType    `json:"type"`
}

// Section represents `search_hint_section` API object
type SearchSection string

func (s *SearchSection) MarshalJSON() ([]byte, error) {
	return GetStringFromRange(string(*s), "groups", "events", "publics", "correspondents", "people",
		"friends", "mutual_friends")
}

func (s *SearchSection) String() string {
	return string(*s)
}

// Type represents `search_hint_type` API object
type SearchType string

func (t *SearchType) MarshalJSON() ([]byte, error) {
	return GetStringFromRange(string(*t), "group", "profile")
}

func (t *SearchType) String() string {
	return string(*t)
}
