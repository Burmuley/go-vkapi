package market

/////////////////////////////////////////////////////////////
// Market related API objects	                           //
/////////////////////////////////////////////////////////////

// Currency represents `market_currency` API object
type Currency struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Price represents `market_price` API object
type Price struct {
	Amount   string   `json:"amount"`
	Text     string   `json:"text"`
	Currency Currency `json:"currency"`
}

// Section represents `market_section` API object
type Section struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// MarketItem represents `market_market_item` API object
type MarketItem struct {
}

type MarketAlbum struct {
}
