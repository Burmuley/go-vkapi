package likes

import "fmt"

// Type represents `likes_type` API object
type Type string

func (t *Type) MarshalJSON() ([]byte, error) {
	switch *t {
	case "post",
		"comment",
		"photo",
		"audio",
		"video",
		"note",
		"market",
		"photo_comment",
		"video_comment",
		"topic_comment",
		"market_comment",
		"sitepage":
		return []byte(*t), nil
	}

	return []byte{}, fmt.Errorf("value is not in allowed range")
}

func (t *Type) GetName() string {
	return string(*t)
}
