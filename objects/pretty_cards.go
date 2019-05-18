package objects

/////////////////////////////////////////////////////////////
// PrettyCards related API objects	                       //
/////////////////////////////////////////////////////////////

// PrettyCard represents `prettyCards_prettyCard` API object
type PrettyCardsPrettyCard struct {
	Button     string      `json:"button"`      // Button key
	ButtonText string      `json:"button_text"` // Button text in current language
	CardID     int         `json:"card_id"`     // Card ID (long int returned as string)
	Images     []BaseImage `json:"images"`
	LinkUrl    string      `json:"link_url"`  // Link URL
	Photo      string      `json:"photo"`     // Photo ID (format \"<owner_id>_<media_id>\")
	Price      float64     `json:"price"`     // Price if set (decimal number returned as string)
	PriceOld   float64     `json:"price_old"` // Old price if set (decimal number returned as string)
	Title      string      `json:"title"`     // Title
}
