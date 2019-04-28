package objects

import "fmt"

// GiftsLayout represents `gifts_layout` API object
type GiftsLayout struct {
	ID       int    `json:"id"`
	Thumb256 string `json:"thumb_256"`
	Thumb96  string `json:"thumb_96"`
	Thumb48  string `json:"thumb_48"`
}

// GiftsGiftPrivacy represents `gifts_gift_privacy` API object
type GiftsGiftPrivacy int

func (g *GiftsGiftPrivacy) MarshalJSON() ([]byte, error) {
	if *g < 0 || *g > 2 {
		return []byte{}, fmt.Errorf("value not in range 0..2")
	}

	return []byte{byte(*g)}, nil
}

func (g *GiftsGiftPrivacy) GetName() string {
	switch *g {
	case 0:
		return "name and message for all"
	case 1:
		return "name for all"
	case 2:
		return "name and message for recipient only"
	default:
		return "unknown"
	}
}

// GiftsGift represents `gifts_gift` API object
type GiftsGift struct {
	Date     int              `json:"date"`
	From     int              `json:"from_id"`
	Gift     GiftsLayout      `json:"gift"`
	GiftHash string           `json:"gift_hash"`
	ID       int              `json:"id"`
	Message  string           `json:"message"`
	Privacy  GiftsGiftPrivacy `json:"privacy"`
}
