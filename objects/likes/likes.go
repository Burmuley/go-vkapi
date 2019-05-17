package likes

import (
	"gitlab.com/Burmuley/go-vkapi/objects"
)

/////////////////////////////////////////////////////////////
// Likes related API objects	                           //
/////////////////////////////////////////////////////////////

// Type represents `likes_type` API object
type Type string

func (t *Type) MarshalJSON() ([]byte, error) {
	return objects.GetStringFromRange(string(*t), "post", "comment", "photo", "audio", "video", "note",
		"market", "photo_comment", "video_comment", "topic_comment", "market_comment", "sitepage")
}

func (t *Type) GetName() string {
	return string(*t)
}
