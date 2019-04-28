package objects

// EventsEventAttach represents `events_event_attach` API object
type EventsEventAttach struct {
	Address      string                      `json:"address"`
	ButtonText   string                      `json:"button_text"`
	Friends      []int                       `json:"friends"`
	ID           int                         `json:"id"`
	Favorite     bool                        `json:"is_favorite"`
	MemberStatus GroupsGroupFullMemberStatus `json:"member_status"`
	Text         string                      `json:"text"`
	Type         int                         `json:"type"`
}

// FaveFavesLink represents `fave_faves_link` API object
type FaveFavesLink struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Photo100    string `json:"photo100"`
	Photo200    string `json:"photo200"`
	Photo50     string `json:"photo50"`
	Title       string `json:"title"`
	Url         string `json:"url"`
}
