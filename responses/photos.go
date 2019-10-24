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

import (
	"gitlab.com/Burmuley/go-vkapi/objects"
)

/////////////////////////////////////////////////////////////
// `photos` group of responses
/////////////////////////////////////////////////////////////

// PhotosCopy type represents `photos_copy_response` API response object
type PhotosCopy int // Photo ID

// PhotosCreatealbum type represents `photos_createAlbum_response` API response object
type PhotosCreatealbum *objects.PhotosPhotoAlbumFull

// PhotosCreatecomment type represents `photos_createComment_response` API response object
type PhotosCreatecomment int // Created comment ID

// PhotosDeletecomment type represents `photos_deleteComment_response` API response object
type PhotosDeletecomment *objects.BaseBoolInt // Returns 1 if request has been processed successfully, 0 if the comment is not found

// PhotosGetalbumscount type represents `photos_getAlbumsCount_response` API response object
type PhotosGetalbumscount int // Albums number

// PhotosGetalbum type represents `photos_getAlbums_response` API response object
type PhotosGetalbum struct {
	Count int                             `json:"count"` // Total number
	Items []*objects.PhotosPhotoAlbumFull `json:"items"`
}

// PhotosGetallcomment type represents `photos_getAllComments_response` API response object
type PhotosGetallcomment struct {
	Count int                            `json:"count"` // Total number
	Items []*objects.PhotosCommentXtrPid `json:"items"`
}

// PhotosGetallExtended type represents `photos_getAll_extended_response` API response object
type PhotosGetallExtended struct {
	Count int                                     `json:"count"` // Total number
	Items []*objects.PhotosPhotoFullXtrRealOffset `json:"items"`
	More  *objects.BaseBoolInt                    `json:"more"` // Information whether next page is presented
}

// PhotosGetall type represents `photos_getAll_response` API response object
type PhotosGetall struct {
	Count int                                 `json:"count"` // Total number
	Items []*objects.PhotosPhotoXtrRealOffset `json:"items"`
	More  *objects.BaseBoolInt                `json:"more"` // Information whether next page is presented
}

// PhotosGetbyidExtended type represents `photos_getById_extended_response` API response object
type PhotosGetbyidExtended *objects.PhotosPhotoFull

// PhotosGetbyid type represents `photos_getById_response` API response object
type PhotosGetbyid *objects.PhotosPhoto

// PhotosGetcommentsExtended type represents `photos_getComments_extended_response` API response object
type PhotosGetcommentsExtended struct {
	Count      int                        `json:"count"` // Total number
	Groups     []*objects.GroupsGroupFull `json:"groups"`
	Items      []*objects.WallWallComment `json:"items"`
	Profiles   []*objects.UsersUserFull   `json:"profiles"`
	RealOffset int                        `json:"real_offset"` // Real offset of the comments
}

// PhotosGetcomment type represents `photos_getComments_response` API response object
type PhotosGetcomment struct {
	Count      int                        `json:"count"` // Total number
	Items      []*objects.WallWallComment `json:"items"`
	RealOffset int                        `json:"real_offset"` // Real offset of the comments
}

// PhotosGetmarketuploadserver type represents `photos_getMarketUploadServer_response` API response object
type PhotosGetmarketuploadserver *objects.BaseUploadServer

// PhotosGetmessagesuploadserver type represents `photos_getMessagesUploadServer_response` API response object
type PhotosGetmessagesuploadserver *objects.PhotosPhotoUpload

// PhotosGetnewtag type represents `photos_getNewTags_response` API response object
type PhotosGetnewtag struct {
	Count int                              `json:"count"` // Total number
	Items []*objects.PhotosPhotoXtrTagInfo `json:"items"`
}

// PhotosGettag type represents `photos_getTags_response` API response object
type PhotosGettag *objects.PhotosPhotoTag

// PhotosGetuploadserver type represents `photos_getUploadServer_response` API response object
type PhotosGetuploadserver *objects.PhotosPhotoUpload

// PhotosGetuserphotosExtended type represents `photos_getUserPhotos_extended_response` API response object
type PhotosGetuserphotosExtended struct {
	Count int                        `json:"count"` // Total number
	Items []*objects.PhotosPhotoFull `json:"items"`
}

// PhotosGetuserphot type represents `photos_getUserPhotos_response` API response object
type PhotosGetuserphot struct {
	Count int                    `json:"count"` // Total number
	Items []*objects.PhotosPhoto `json:"items"`
}

// PhotosGetwalluploadserver type represents `photos_getWallUploadServer_response` API response object
type PhotosGetwalluploadserver *objects.PhotosPhotoUpload

// PhotosGetExtended type represents `photos_get_extended_response` API response object
type PhotosGetExtended struct {
	Count int                        `json:"count"` // Total number
	Items []*objects.PhotosPhotoFull `json:"items"`
}

// PhotosGet type represents `photos_get_response` API response object
type PhotosGet struct {
	Count int                    `json:"count"` // Total number
	Items []*objects.PhotosPhoto `json:"items"`
}

// PhotosPuttag type represents `photos_putTag_response` API response object
type PhotosPuttag int // Created tag ID

// PhotosRestorecomment type represents `photos_restoreComment_response` API response object
type PhotosRestorecomment *objects.BaseBoolInt // Returns 1 if request has been processed successfully, 0 if the comment is not found

// PhotosSavemarketalbumphot type represents `photos_saveMarketAlbumPhoto_response` API response object
type PhotosSavemarketalbumphot *objects.PhotosPhoto

// PhotosSavemarketphot type represents `photos_saveMarketPhoto_response` API response object
type PhotosSavemarketphot *objects.PhotosPhoto

// PhotosSavemessagesphot type represents `photos_saveMessagesPhoto_response` API response object
type PhotosSavemessagesphot *objects.PhotosPhoto

// PhotosSaveownercoverphot type represents `photos_saveOwnerCoverPhoto_response` API response object
type PhotosSaveownercoverphot *objects.BaseImage

// PhotosSaveownerphot type represents `photos_saveOwnerPhoto_response` API response object
type PhotosSaveownerphot struct {
	PhotoHash     string `json:"photo_hash"`      // Photo hash
	PhotoSrc      string `json:"photo_src"`       // Uploaded image url
	PhotoSrcBig   string `json:"photo_src_big"`   // Uploaded image url
	PhotoSrcSmall string `json:"photo_src_small"` // Uploaded image url
	PostId        int    `json:"post_id"`         // Created post ID
	Saved         int    `json:"saved"`           // Returns 1 if profile photo is saved
}

// PhotosSavewallphot type represents `photos_saveWallPhoto_response` API response object
type PhotosSavewallphot *objects.PhotosPhoto

// PhotosSav type represents `photos_save_response` API response object
type PhotosSav *objects.PhotosPhoto

// PhotosSearch type represents `photos_search_response` API response object
type PhotosSearch struct {
	Count int                    `json:"count"` // Total number
	Items []*objects.PhotosPhoto `json:"items"`
}
