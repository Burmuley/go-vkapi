package objects

/////////////////////////////////////////////////////////////
// Notes related API objects	                           //
/////////////////////////////////////////////////////////////

// Note represents `notes_note` API object
type NotesNote struct {
	CanComment BaseBoolInt `json:"can_comment"` // Information whether current user can comment the note
	Comments   int         `json:"comments"`    // Comments number
	Date       int         `json:"date"`        // Date when the note has been created in Unixtime
	ID         int         `json:"id"`          // Note ID
	OwnerID    int         `json:"owner_id"`    // Note owner's ID
	Text       string      `json:"text"`        // Note text
	TextWiki   string      `json:"text_wiki"`
	Title      string      `json:"title"`    // Note title
	ViewUrl    string      `json:"view_url"` // URL of the page with note preview
}

// NoteComment represents `notes_note_comment` API object
type NotesNoteComment struct {
	Date     int    `json:"date"`     // Date when the comment has beed added in Unixtime
	ID       int    `json:"id"`       // Comment ID
	Msg      string `json:"message"`  // Comment text
	NoteID   int    `json:"nid"`      // Note ID
	OwnerID  int    `json:"oid"`      // Note ID
	ReplyTo  int    `json:"reply_to"` // ID of replied comment
	AuthorID int    `json:"uid"`      // Comment author's ID
}
