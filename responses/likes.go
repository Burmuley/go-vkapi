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
// `likes` group of responses
/////////////////////////////////////////////////////////////

// LikesAdd type represents `likes_add_response` API response object
type LikesAdd struct {
	Likes int `json:"likes"` // Total likes number
}

// LikesDelete type represents `likes_delete_response` API response object
type LikesDelete struct {
	Likes int `json:"likes"` // Total likes number
}

// LikesGetListExtended type represents `likes_getList_extended_response` API response object
type LikesGetListExtended struct {
	Count int                    `json:"count"` // Total number
	Items []objects.UsersUserMin `json:"items"`
}

// LikesGetList type represents `likes_getList_response` API response object
type LikesGetList struct {
	Count int   `json:"count"` // Total number
	Items []int `json:"items"`
}

// LikesIsLiked type represents `likes_isLiked_response` API response object
type LikesIsLiked struct {
	Copied objects.BaseBoolInt `json:"copied"` // Information whether user reposted the object
	Liked  objects.BaseBoolInt `json:"liked"`  // Information whether user liked the object
}
