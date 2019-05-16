package audio

import "gitlab.com/Burmuley/go-vkapi/objects/base"

/////////////////////////////////////////////////////////////
// Audio related API objects	                           //
/////////////////////////////////////////////////////////////

// Audio represents `audio_audio` API object
type Audio struct {
	AccessKey  string `json:"access_key"`
	Artist     string `json:"artist"`
	ID         int    `json:"id"`
	Explicit   bool   `json:"is_explicit"`
	FocusTrack bool   `json:"is_focus_track"`
	Licensed   bool   `json:"is_licensed"`
	OwnerID    int    `json:"owner_id"`
	Title      string `json:"title"`
	Url        string `json:"url"`
}

// AudioFull represents `audio_audio_full` API object
type AudioFull struct {
	Audio
	Duration    int          `json:"duration"`
	Date        int          `json:"date"`
	AlbumID     int          `json:"album_id"`
	LyricsID    int          `json:"lyrics_id"`
	GenreID     int          `json:"genre_id"`
	Searchable  base.BoolInt `json:"no_search"`
	HighQuality bool         `json:"is_hq"`
}
