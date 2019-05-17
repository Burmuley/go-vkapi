package gifts

import (
	"gitlab.com/Burmuley/go-vkapi/objects"
)

/////////////////////////////////////////////////////////////
// Gifts related API objects	                           //
/////////////////////////////////////////////////////////////

// Layout represents `gifts_layout` API object
type Layout struct {
	ID       int    `json:"id"`
	Thumb256 string `json:"thumb_256"`
	Thumb96  string `json:"thumb_96"`
	Thumb48  string `json:"thumb_48"`
}

// GiftPrivacy represents `gifts_gift_privacy` API object
type GiftPrivacy int

func (g *GiftPrivacy) MarshalJSON() ([]byte, error) {
	return objects.GetIntFromRange(int(*g), 0, 2, true)
}

func (g *GiftPrivacy) GetName() string {
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

// Gift represents `gifts_gift` API object
type Gift struct {
	Date     int         `json:"date"`
	From     int         `json:"from_id"`
	Gift     Layout      `json:"gift"`
	GiftHash string      `json:"gift_hash"`
	ID       int         `json:"id"`
	Message  string      `json:"message"`
	Privacy  GiftPrivacy `json:"privacy"`
}
