package objects

// MarketCurrency represents `market_currency` API object
type MarketCurrency struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// MarketPrice represents `market_price` API object
type MarketPrice struct {
	Amount   string         `json:"amount"`
	Text     string         `json:"text"`
	Currency MarketCurrency `json:"currency"`
}

// MarketSection represents `market_section` API object
type MarketSection struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
