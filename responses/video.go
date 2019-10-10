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

package responses

import "gitlab.com/Burmuley/go-vkapi/objects"

/////////////////////////////////////////////////////////////
// `video` group of responses
/////////////////////////////////////////////////////////////

// VideoAddAlbum type represents `video_addAlbum_response` API response object
type VideoAddAlbum struct {
	AlbumId int `json:"album_id"` // Created album ID
}

// VideoCreateComment type represents `video_createComment_response` API response object
type VideoCreateComment int // Created comment ID

// VideoGetAlbumById type represents `video_getAlbumById_response` API response object
type VideoGetAlbumById objects.VideoVideoAlbumFull

// VideoGetAlbumsByVideoExtended type represents `video_getAlbumsByVideo_extended_response` API response object
type VideoGetAlbumsByVideoExtended struct {
	Count int                           `json:"count"` // Total number
	Items []objects.VideoVideoAlbumFull `json:"items"`
}

// VideoGetAlbumsByVideo type represents `video_getAlbumsByVideo_response` API response object
type VideoGetAlbumsByVideo int

// VideoGetAlbumsExtended type represents `video_getAlbums_extended_response` API response object
type VideoGetAlbumsExtended struct {
	Count int                           `json:"count"` // Total number
	Items []objects.VideoVideoAlbumFull `json:"items"`
}

// VideoGetAlbums type represents `video_getAlbums_response` API response object
type VideoGetAlbums struct {
	Count int                           `json:"count"` // Total number
	Items []objects.VideoVideoAlbumFull `json:"items"`
}

// VideoGetCommentsExtended type represents `video_getComments_extended_response` API response object
type VideoGetCommentsExtended struct {
	Count    int                       `json:"count"` // Total number
	Groups   []objects.GroupsGroupFull `json:"groups"`
	Items    []objects.WallWallComment `json:"items"`
	Profiles []objects.UsersUserMin    `json:"profiles"`
}

// VideoGetComments type represents `video_getComments_response` API response object
type VideoGetComments struct {
	Count int                       `json:"count"` // Total number
	Items []objects.WallWallComment `json:"items"`
}

// VideoGetExtended type represents `video_get_extended_response` API response object
type VideoGetExtended struct {
	Count    int                       `json:"count"` // Total number
	Groups   []objects.GroupsGroupFull `json:"groups"`
	Items    []objects.VideoVideoFull  `json:"items"`
	Profiles []objects.UsersUserMin    `json:"profiles"`
}

// VideoGet type represents `video_get_response` API response object
type VideoGet struct {
	Count int                  `json:"count"` // Total number
	Items []objects.VideoVideo `json:"items"`
}

// VideoRestoreComment type represents `video_restoreComment_response` API response object
type VideoRestoreComment objects.BaseBoolInt // Returns 1 if request has been processed successfully, 0 if the comment is not found

// VideoSave type represents `video_save_response` API response object
type VideoSave objects.VideoSaveResult

// VideoSearchExtended type represents `video_search_extended_response` API response object
type VideoSearchExtended struct {
	Count    int                       `json:"count"` // Total number
	Groups   []objects.GroupsGroupFull `json:"groups"`
	Items    []objects.VideoVideo      `json:"items"`
	Profiles []objects.UsersUserMin    `json:"profiles"`
}

// VideoSearch type represents `video_search_response` API response object
type VideoSearch struct {
	Count int                  `json:"count"` // Total number
	Items []objects.VideoVideo `json:"items"`
}
