package objects

/////////////////////////////////////////////////////////////
// Polls related API objects	                           //
/////////////////////////////////////////////////////////////

// Answer represents `polls_answer` API object
type PollsAnswer struct {
	ID    int     `json:"id"`    // Answer ID
	Rate  float64 `json:"rate"`  // Answer rate in percents
	Text  string  `json:"text"`  // Answer text
	Votes int     `json:"votes"` // Votes number
}

// Poll represents `polls_poll` API object
type PollsPoll struct {
	Anonymous BaseBoolInt   `json:"anonymous"` // Information whether the pole is anonymous
	AnswerID  int           `json:"answer_id"` // Current user's answer ID
	Answers   []PollsAnswer `json:"answers"`
	Created   int           `json:"created"`  // Date when poll has been created in Unixtime
	ID        int           `json:"id"`       // Poll ID
	OwnerID   int           `json:"owner_id"` // Poll owner's ID
	Question  string        `json:"question"` // Poll question
	Votes     int           `json:"votes"`    // Votes number
}

// Voters represents `polls_voters` API object
type PollsVoters struct {
	AnswerID int              `json:"answer_id"` // Answer ID
	Users    PollsVotersUsers `json:"users"`
}

// VotersUsers represents `polls_voters_users` API object
type PollsVotersUsers struct {
	Count int   `json:"count"` // Votes number
	Items []int `json:"items"` // User IDs
}
