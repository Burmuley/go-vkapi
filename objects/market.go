package objects

/////////////////////////////////////////////////////////////
// Market related API objects	                           //
/////////////////////////////////////////////////////////////

// Currency represents `market_currency` API object
type MarketCurrency struct {
	ID   int    `json:"id"`   // Currency ID
	Name string `json:"name"` // Currency sign
}

// Price represents `market_price` API object
type MarketPrice struct {
	Amount   string         `json:"amount"` // Amount
	Text     string         `json:"text"`
	Currency MarketCurrency `json:"currency"` // Text
}

// Section represents `market_section` API object
type MarketSection struct {
	ID   int    `json:"id"`   // Section ID
	Name string `json:"name"` // Section name
}

// Item represents `market_market_item` API object
type MarketItem struct {
	AccessKey    string                 `json:"access_key"` // Access key for the market item
	Availability MarketItemAvailability `json:"availability"`
	ButtonTitle  string                 `json:"button_title"` // Title for button for url
	Category     MarketCategory         `json:"category"`
	Date         int                    `json:"date"`        // Date when the item has been created in Unixtime
	Description  string                 `json:"description"` // Item description
	ID           int                    `json:"id"`          // Item ID
	IsFavorite   bool                   `json:"is_favorite"`
	OwnerID      int                    `json:"owner_id"` // Item owner's ID
	Price        MarketPrice            `json:"price"`
	ThumbPhoto   string                 `json:"thumb_photo"` // URL of the preview image
	Title        string                 `json:"title"`       // Item title
	Url          string                 `json:"url"`         // URL to item
}

// ItemFull represents `market_market_item_full` API object
type MarketItemFull struct {
	*MarketItem
	Photos     []PhotosPhoto `json:"photos"`
	CanComment BaseBoolInt   `json:"can_comment"` // Information whether current use can comment the item
	CanRepost  BaseBoolInt   `json:"can_repost"`  // Information whether current use can repost the item
	Likes      BaseLikes     `json:"likes"`
	Views      int           `json:"views_count"` // Views number
}

// ItemAvailability represents `market_market_item_availability` API object
type MarketItemAvailability int

func (m *MarketItemAvailability) String() string {
	switch *m {
	case 0:
		return "available"
	case 1:
		return "removed"
	case 2:
		return "unavailable"
	default:
		return "unknown"
	}
}

// Album represents `market_market_album` API object
type MarketAlbum struct {
	Count       int         `json:"count"`    // Items number
	ID          int         `json:"id"`       // Market album ID
	OwnerID     int         `json:"owner_id"` // Market album owner's ID
	Photo       PhotosPhoto `json:"photo"`
	Title       string      `json:"title"`        // Market album title
	UpdatedTime int         `json:"updated_time"` // Date when album has been updated last time in Unixtime
}

// Category represents `market_market_cetegory` API object
type MarketCategory struct {
	ID      int           `json:"id"`   // Category ID
	Name    string        `json:"name"` // Category name
	Section MarketSection `json:"section"`
}
