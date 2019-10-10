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
// `photos` group of responses
/////////////////////////////////////////////////////////////

// PhotosCopy type represents `photos_copy_response` API response object
type PhotosCopy int // Photo ID

// PhotosCreateAlbum type represents `photos_createAlbum_response` API response object
type PhotosCreateAlbum objects.PhotosPhotoAlbumFull

// PhotosCreateComment type represents `photos_createComment_response` API response object
type PhotosCreateComment int // Created comment ID

// PhotosDeleteComment type represents `photos_deleteComment_response` API response object
type PhotosDeleteComment objects.BaseBoolInt // Returns 1 if request has been processed successfully, 0 if the comment is not found

// PhotosGetAlbumsCount type represents `photos_getAlbumsCount_response` API response object
type PhotosGetAlbumsCount int // Albums number

// PhotosGetAlbums type represents `photos_getAlbums_response` API response object
type PhotosGetAlbums struct {
	Count int                            `json:"count"` // Total number
	Items []objects.PhotosPhotoAlbumFull `json:"items"`
}

// PhotosGetAllComments type represents `photos_getAllComments_response` API response object
type PhotosGetAllComments struct {
	Count int                           `json:"count"` // Total number
	Items []objects.PhotosCommentXtrPid `json:"items"`
}

// PhotosGetAllExtended type represents `photos_getAll_extended_response` API response object
type PhotosGetAllExtended struct {
	Count int                                    `json:"count"` // Total number
	Items []objects.PhotosPhotoFullXtrRealOffset `json:"items"`
	More  objects.BaseBoolInt                    `json:"more"` // Information whether next page is presented
}

// PhotosGetAll type represents `photos_getAll_response` API response object
type PhotosGetAll struct {
	Count int                                `json:"count"` // Total number
	Items []objects.PhotosPhotoXtrRealOffset `json:"items"`
	More  objects.BaseBoolInt                `json:"more"` // Information whether next page is presented
}

// PhotosGetByIdExtended type represents `photos_getById_extended_response` API response object
type PhotosGetByIdExtended objects.PhotosPhotoFull

// PhotosGetById type represents `photos_getById_response` API response object
type PhotosGetById objects.PhotosPhoto

// PhotosGetCommentsExtended type represents `photos_getComments_extended_response` API response object
type PhotosGetCommentsExtended struct {
	Count      int                       `json:"count"` // Total number
	Groups     []objects.GroupsGroupFull `json:"groups"`
	Items      []objects.WallWallComment `json:"items"`
	Profiles   []objects.UsersUserFull   `json:"profiles"`
	RealOffset int                       `json:"real_offset"` // Real offset of the comments
}

// PhotosGetComments type represents `photos_getComments_response` API response object
type PhotosGetComments struct {
	Count      int                       `json:"count"` // Total number
	Items      []objects.WallWallComment `json:"items"`
	RealOffset int                       `json:"real_offset"` // Real offset of the comments
}

// PhotosGetMarketUploadServer type represents `photos_getMarketUploadServer_response` API response object
type PhotosGetMarketUploadServer objects.BaseUploadServer

// PhotosGetMessagesUploadServer type represents `photos_getMessagesUploadServer_response` API response object
type PhotosGetMessagesUploadServer objects.PhotosPhotoUpload

// PhotosGetNewTags type represents `photos_getNewTags_response` API response object
type PhotosGetNewTags struct {
	Count int                             `json:"count"` // Total number
	Items []objects.PhotosPhotoXtrTagInfo `json:"items"`
}

// PhotosGetTags type represents `photos_getTags_response` API response object
type PhotosGetTags objects.PhotosPhotoTag

// PhotosGetUploadServer type represents `photos_getUploadServer_response` API response object
type PhotosGetUploadServer objects.PhotosPhotoUpload

// PhotosGetUserPhotosExtended type represents `photos_getUserPhotos_extended_response` API response object
type PhotosGetUserPhotosExtended struct {
	Count int                       `json:"count"` // Total number
	Items []objects.PhotosPhotoFull `json:"items"`
}

// PhotosGetUserPhotos type represents `photos_getUserPhotos_response` API response object
type PhotosGetUserPhotos struct {
	Count int                   `json:"count"` // Total number
	Items []objects.PhotosPhoto `json:"items"`
}

// PhotosGetWallUploadServer type represents `photos_getWallUploadServer_response` API response object
type PhotosGetWallUploadServer objects.PhotosPhotoUpload

// PhotosGetExtended type represents `photos_get_extended_response` API response object
type PhotosGetExtended struct {
	Count int                       `json:"count"` // Total number
	Items []objects.PhotosPhotoFull `json:"items"`
}

// PhotosGet type represents `photos_get_response` API response object
type PhotosGet struct {
	Count int                   `json:"count"` // Total number
	Items []objects.PhotosPhoto `json:"items"`
}

// PhotosPutTag type represents `photos_putTag_response` API response object
type PhotosPutTag int // Created tag ID

// PhotosRestoreComment type represents `photos_restoreComment_response` API response object
type PhotosRestoreComment objects.BaseBoolInt // Returns 1 if request has been processed successfully, 0 if the comment is not found

// PhotosSaveMarketAlbumPhoto type represents `photos_saveMarketAlbumPhoto_response` API response object
type PhotosSaveMarketAlbumPhoto objects.PhotosPhoto

// PhotosSaveMarketPhoto type represents `photos_saveMarketPhoto_response` API response object
type PhotosSaveMarketPhoto objects.PhotosPhoto

// PhotosSaveMessagesPhoto type represents `photos_saveMessagesPhoto_response` API response object
type PhotosSaveMessagesPhoto objects.PhotosPhoto

// PhotosSaveOwnerCoverPhoto type represents `photos_saveOwnerCoverPhoto_response` API response object
type PhotosSaveOwnerCoverPhoto objects.BaseImage

// PhotosSaveOwnerPhoto type represents `photos_saveOwnerPhoto_response` API response object
type PhotosSaveOwnerPhoto struct {
	PhotoHash     string `json:"photo_hash"`      // Photo hash
	PhotoSrc      string `json:"photo_src"`       // Uploaded image url
	PhotoSrcBig   string `json:"photo_src_big"`   // Uploaded image url
	PhotoSrcSmall string `json:"photo_src_small"` // Uploaded image url
	PostId        int    `json:"post_id"`         // Created post ID
	Saved         int    `json:"saved"`           // Returns 1 if profile photo is saved
}

// PhotosSaveWallPhoto type represents `photos_saveWallPhoto_response` API response object
type PhotosSaveWallPhoto objects.PhotosPhoto

// PhotosSave type represents `photos_save_response` API response object
type PhotosSave objects.PhotosPhoto

// PhotosSearch type represents `photos_search_response` API response object
type PhotosSearch struct {
	Count int                   `json:"count"` // Total number
	Items []objects.PhotosPhoto `json:"items"`
}
