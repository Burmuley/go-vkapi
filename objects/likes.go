package objects

/////////////////////////////////////////////////////////////
// Likes related API objects	                           //
/////////////////////////////////////////////////////////////

// Type represents `likes_type` API object
type LikesType string

func (t *LikesType) MarshalJSON() ([]byte, error) {
	return GetStringFromRange(string(*t), "post", "comment", "photo", "audio", "video", "note",
		"market", "photo_comment", "video_comment", "topic_comment", "market_comment", "sitepage")
}

func (t *LikesType) GetName() string {
	return string(*t)
}
