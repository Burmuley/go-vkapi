/*
Copyright 2019 Konstantin Vasilev (burmuley@gmail.com)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

/////////////////////////////////////////////////////////////////////////////////////////////////////////////
// WARNING! AUTOMATICALLY GENERATED CONTENT! DON'T CHANGE IT MANUALLY!                                     //
// Source schema can be found at https://github.com/VKCOM/vk-api-schema/blob/master/responses.json         //
// Code generator location: https://gitlab.com/Burmuley/go-vkapi-gen                                       //
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

package objects

/////////////////////////////////////////////////////////////////////////////////////////////////////////////
// `video` group of objects
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

// VideoSaveResult type represents `video_save_result` API object
type VideoSaveResult struct {
	AccessKey   string `json:"access_key"`  // Video access key
	Description string `json:"description"` // Video description
	OwnerId     int    `json:"owner_id"`    // Video owner ID
	Title       string `json:"title"`       // Video title
	UploadUrl   string `json:"upload_url"`  // URL for the video uploading
	VideoId     int    `json:"video_id"`    // Video ID
}

// VideoVideo type represents `video_video` API object
type VideoVideo struct {
	AccessKey   string             `json:"access_key"`  // Video access key
	AddingDate  int                `json:"adding_date"` // Date when the video has been added in Unixtime
	CanAdd      BaseBoolInt        `json:"can_add"`     // Information whether current user can add the video
	CanComment  BaseBoolInt        `json:"can_comment"` // Information whether current user can comment the video
	CanEdit     BaseBoolInt        `json:"can_edit"`    // Information whether current user can edit the video
	CanLike     BaseBoolInt        `json:"can_like"`    // Information whether current user can like the video
	CanRepost   BaseBoolInt        `json:"can_repost"`  // Information whether current user can repost this video
	Comments    int                `json:"comments"`    // Number of comments
	Date        int                `json:"date"`        // Date when video has been uploaded in Unixtime
	Description string             `json:"description"` // Video description
	Duration    int                `json:"duration"`    // Video duration in seconds
	Files       VideoVideoFiles    `json:"files"`
	FirstFrame  []VideoVideoImage  `json:"first_frame"`
	Height      int                `json:"height"` // Video height
	Id          int                `json:"id"`     // Video ID
	Image       []VideoVideoImage  `json:"image"`
	IsFavorite  bool               `json:"is_favorite"`
	Live        BasePropertyExists `json:"live"`       // Returns if the video is a live stream
	OwnerId     int                `json:"owner_id"`   // Video owner ID
	Player      string             `json:"player"`     // URL of the page with a player that can be used to play the video in the browser.
	Processing  BasePropertyExists `json:"processing"` // Returns if the video is processing
	Title       string             `json:"title"`      // Video title
	Type        string             `json:"type"`
	Views       int                `json:"views"` // Number of views
	Width       int                `json:"width"` // Video width
}

// VideoVideoAlbumFull type represents `video_video_album_full` API object
type VideoVideoAlbumFull struct {
	Count       int               `json:"count"`        // Total number of videos in album
	Id          int               `json:"id"`           // Album ID
	Image       []VideoVideoImage `json:"image"`        // Album cover image in different sizes
	IsSystem    int               `json:"is_system"`    // Information whether album is system
	OwnerId     int               `json:"owner_id"`     // Album owner's ID
	Title       string            `json:"title"`        // Album title
	UpdatedTime int               `json:"updated_time"` // Date when the album has been updated last time in Unixtime
}

// VideoVideoFiles type represents `video_video_files` API object
type VideoVideoFiles struct {
	External string `json:"external"` // URL of the external player
	Mp41080  string `json:"mp4_1080"` // URL of the mpeg4 file with 1080p quality
	Mp4240   string `json:"mp4_240"`  // URL of the mpeg4 file with 240p quality
	Mp4360   string `json:"mp4_360"`  // URL of the mpeg4 file with 360p quality
	Mp4480   string `json:"mp4_480"`  // URL of the mpeg4 file with 480p quality
	Mp4720   string `json:"mp4_720"`  // URL of the mpeg4 file with 720p quality
}

// VideoVideoFull type represents `video_video_full` API object
type VideoVideoFull struct {
	AccessKey      string             `json:"access_key"`       // Video access key
	AddingDate     int                `json:"adding_date"`      // Date when the video has been added in Unixtime
	CanAdd         BaseBoolInt        `json:"can_add"`          // Information whether current user can add the video
	CanAddToFaves  BaseBoolInt        `json:"can_add_to_faves"` // Information whether current user can add the video to favourites
	CanComment     BaseBoolInt        `json:"can_comment"`      // Information whether current user can comment the video
	CanEdit        BaseBoolInt        `json:"can_edit"`         // Information whether current user can edit the video
	CanRepost      BaseBoolInt        `json:"can_repost"`       // Information whether current user can comment the video
	Comments       int                `json:"comments"`         // Number of comments
	Date           int                `json:"date"`             // Date when video has been uploaded in Unixtime
	Description    string             `json:"description"`      // Video description
	Duration       int                `json:"duration"`         // Video duration in seconds
	Files          VideoVideoFiles    `json:"files"`
	FirstFrame     []VideoVideoImage  `json:"first_frame"`
	FirstFrame1280 string             `json:"first_frame_1280"` // URL of the first frame for the corresponding width.
	FirstFrame640  string             `json:"first_frame_640"`  // URL of the first frame for the corresponding width.
	Id             int                `json:"id"`               // Video ID
	Image          []VideoVideoImage  `json:"image"`
	Likes          BaseLikes          `json:"likes"`
	Live           BasePropertyExists `json:"live"`       // Returns if the video is live translation
	OwnerId        int                `json:"owner_id"`   // Video owner ID
	Player         string             `json:"player"`     // URL of the page with a player that can be used to play the video in the browser.
	Processing     BasePropertyExists `json:"processing"` // Returns if the video is processing
	Repeat         BaseBoolInt        `json:"repeat"`     // Information whether the video is repeated
	Title          string             `json:"title"`      // Video title
	Views          int                `json:"views"`      // Number of views
}

// VideoVideoImage type represents `video_video_image` API object
type VideoVideoImage struct {
	BaseImage   BaseImage   `json:"BaseImage"`
	WithPadding BaseBoolInt `json:"with_padding"`
}
