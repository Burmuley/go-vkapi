package objects

/////////////////////////////////////////////////////////////
// Videos related API objects	                           //
/////////////////////////////////////////////////////////////

// SaveResult represents `video_save_result` API objects
type VideosSaveResult struct {
	Description string `json:"description"` // Video description
	OwnerID     int    `json:"owner_id"`    // Video owner ID
	Title       string `json:"title"`       // Video title
	UploadURL   string `json:"upload_url"`  // URL for the video uploading
	VideoID     int    `json:"video_id"`    // Video ID
}

// Video represents `videos_video` API object
type VideosVideo struct {
	AccessKey     string             `json:"access_key"`  // Video access key
	AddingDate    int                `json:"adding_date"` // Date when the video has been added in Unixtime
	CanAdd        BaseBoolInt        `json:"can_add"`     // Information whether current user can add the video
	CanComment    BaseBoolInt        `json:"can_comment"` // Information whether current user can comment the video
	CanEdit       BaseBoolInt        `json:"can_edit"`    // Information whether current user can edit the video
	CanLike       BaseBoolInt        `json:"can_like"`    // Information whether current user can like the video
	CanRepost     BaseBoolInt        `json:"can_repost"`  // Information whether current user can repost this video
	Comments      int                `json:"comments"`    // Number of comments
	Date          int                `json:"date"`        // Date when video has been uploaded in Unixtime
	Description   string             `json:"description"` // Video description
	Duration      int                `json:"duration"`    // Video duration in seconds
	Files         VideosVideoFiles   `json:"files"`
	FirstFrame130 string             `json:"first_frame_130"` // URL of the first frame for the corresponding width
	FirstFrame160 string             `json:"first_frame_160"` // URL of the first frame for the corresponding width
	FirstFrame320 string             `json:"first_frame_320"` // URL of the first frame for the corresponding width
	FirstFrame800 string             `json:"first_frame_800"` // URL of the first frame for the corresponding width
	Height        int                `json:"height"`          // Video height
	ID            int                `json:"id"`              // Video ID
	IsFavorite    bool               `json:"is_favorite"`
	Live          BasePropertyExists `json:"live"`       // Returns if the video is a live stream
	OwnerID       int                `json:"owner_id"`   // Video owner ID
	Photo130      string             `json:"photo_130"`  // URL of the preview image with 130 px in width
	Photo320      string             `json:"photo_320"`  // URL of the preview image with 320 px in width
	Photo800      string             `json:"photo_800"`  // URL of the preview image with 800 px in width
	Player        string             `json:"player"`     // URL of the page with a player that can be used to play the video in the browser
	Processing    BasePropertyExists `json:"processing"` // Returns if the video is processing
	Title         string             `json:"title"`      // Video title
	Type          string             `json:"type"`       // Content type (video)
	Views         int                `json:"views"`      // Number of views
	Width         int                `json:"width"`      // Video width
}

//VideoAlbumFull represents `video_video_album_full` API object
type VideosVideoAlbumFull struct {
	ID          int                `json:"id"`           // Album ID
	Count       int                `json:"count"`        // Total number of videos in album
	Image       []VideosVideoImage `json:"image"`        // Album cover image in different sizes
	IsSystem    int                `json:"is_system"`    // Information whether album is system
	OwnerID     int                `json:"owner_id"`     // Album owner's ID
	Photo160    string             `json:"photo_160"`    // URL of the preview image with 160px in width
	Photo320    string             `json:"photo_320"`    // URL of the preview image with 320spx in width
	Privacy     string             `json:"privacy"`      // Album privacy settings
	Title       string             `json:"title"`        // Album title
	UpdatedTime int                `json:"updated_time"` // Date when the album has been updated last time in Unixtime
}

// VideoFiles represents `video_video_files` API object
type VideosVideoFiles struct {
	External string `json:"external"` // URL of the external player
	Mp1080   string `json:"mp_1080"`  // URL of the mpeg4 file with 1080p quality
	Mp720    string `json:"mp_720"`   // URL of the mpeg4 file with 720p quality
	Mp240    string `json:"mp_240"`   // URL of the mpeg4 file with 240p quality
	Mp360    string `json:"mp_360"`   // URL of the mpeg4 file with 360p quality
	Mp480    string `json:"mp_480"`   // URL of the mpeg4 file with 480p quality
}

// VideoFull represents `videos_video_full` API object
type VideosVideoFull struct {
	AccessKey      string             `json:"access_key"`  // Video access key
	AddingDate     int                `json:"adding_date"` // Date when the video has been added in Unixtime
	CanAdd         BaseBoolInt        `json:"can_add"`     // Information whether current user can add the video
	CanComment     BaseBoolInt        `json:"can_comment"` // Information whether current user can comment the video
	CanEdit        BaseBoolInt        `json:"can_edit"`    // Information whether current user can edit the video
	CanRepost      BaseBoolInt        `json:"can_repost"`  // Information whether current user can repost this video
	Comments       int                `json:"comments"`    // Number of comments
	Date           int                `json:"date"`        // Date when video has been uploaded in Unixtime
	Description    string             `json:"description"` // Video description
	Duration       int                `json:"duration"`    // Video duration in seconds
	Files          VideosVideoFiles   `json:"files"`
	FirstFrame130  string             `json:"first_frame_130"` // URL of the first frame for the corresponding width
	FirstFrame160  string             `json:"first_frame_160"` // URL of the first frame for the corresponding width
	FirstFrame320  string             `json:"first_frame_320"` // URL of the first frame for the corresponding width
	FirstFrame800  string             `json:"first_frame_800"` // URL of the first frame for the corresponding width
	ID             int                `json:"id"`              // Video ID
	Likes          BaseLikes          `json:"likes"`           // Number of likes for the video
	Live           BasePropertyExists `json:"live"`            // Returns if the video is a live stream
	OwnerID        int                `json:"owner_id"`        // Video owner ID
	Photo130       string             `json:"photo_130"`       // URL of the preview image with 130 px in width
	Photo320       string             `json:"photo_320"`       // URL of the preview image with 320 px in width
	Photo800       string             `json:"photo_800"`       // URL of the preview image with 800 px in width
	Player         string             `json:"player"`          // URL of the page with a player that can be used to play the video in the browser
	PrivacyComment []string           `json:"privacy_comment"` // Privacy comment
	PrivacyView    []string           `json:"privacy_view"`    // Privacy view
	Processing     BasePropertyExists `json:"processing"`      // Returns if the video is processing
	Repeat         BaseBoolInt        `json:"repeat"`          // Information whether the video is repeated
	Title          string             `json:"title"`           // Video title
	Views          int                `json:"views"`           // Number of views
}

// VideoImage represents `video_video_image` API object
type VideosVideoImage struct {
	*BaseImage
	WithPadding BaseBoolInt `json:"with_padding"`
}
