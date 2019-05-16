package fave

/////////////////////////////////////////////////////////////
// Fave related API objects	                               //
/////////////////////////////////////////////////////////////

// FavesLink represents `fave_faves_link` API object
type FavesLink struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Photo100    string `json:"photo100"`
	Photo200    string `json:"photo200"`
	Photo50     string `json:"photo50"`
	Title       string `json:"title"`
	Url         string `json:"url"`
}
