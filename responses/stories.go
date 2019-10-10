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
// `stories` group of responses
/////////////////////////////////////////////////////////////

// StoriesGetBannedExtended type represents `stories_getBanned_extended_response` API response object
type StoriesGetBannedExtended struct {
	Count    int                       `json:"count"` // Stories count
	Groups   []objects.GroupsGroupFull `json:"groups"`
	Items    []int                     `json:"items"`
	Profiles []objects.UsersUserFull   `json:"profiles"`
}

// StoriesGetBanned type represents `stories_getBanned_response` API response object
type StoriesGetBanned struct {
	Count int   `json:"count"` // Stories count
	Items []int `json:"items"`
}

// StoriesGetByIdExtended type represents `stories_getById_extended_response` API response object
type StoriesGetByIdExtended struct {
	Count    int                       `json:"count"` // Stories count
	Groups   []objects.GroupsGroupFull `json:"groups"`
	Items    []objects.StoriesStory    `json:"items"`
	Profiles []objects.UsersUserFull   `json:"profiles"`
}

// StoriesGetById type represents `stories_getById_response` API response object
type StoriesGetById struct {
	Count int                    `json:"count"` // Stories count
	Items []objects.StoriesStory `json:"items"`
}

// StoriesGetPhotoUploadServer type represents `stories_getPhotoUploadServer_response` API response object
type StoriesGetPhotoUploadServer struct {
	UploadUrl string `json:"upload_url"` // Upload URL
	UserIds   []int  `json:"user_ids"`   // Users ID who can to see story.
}

// StoriesGetRepliesExtended type represents `stories_getReplies_extended_response` API response object
type StoriesGetRepliesExtended struct {
	Count    int                       `json:"count"` // Stories count
	Groups   []objects.GroupsGroupFull `json:"groups"`
	Items    []objects.StoriesStory    `json:"items"`
	Profiles []objects.UsersUserFull   `json:"profiles"`
}

// StoriesGetReplies type represents `stories_getReplies_response` API response object
type StoriesGetReplies struct {
	Count int                    `json:"count"` // Stories count
	Items []objects.StoriesStory `json:"items"`
}

// StoriesGetStats type represents `stories_getStats_response` API response object
type StoriesGetStats objects.StoriesStoryStats

// StoriesGetVideoUploadServer type represents `stories_getVideoUploadServer_response` API response object
type StoriesGetVideoUploadServer struct {
	UploadUrl string `json:"upload_url"` // Upload URL
	UserIds   []int  `json:"user_ids"`   // Users ID who can to see story.
}

// StoriesGetViewersExtended type represents `stories_getViewers_extended_response` API response object
type StoriesGetViewersExtended struct {
	Count int                     `json:"count"` // Viewers count
	Items []objects.UsersUserFull `json:"items"`
}

// StoriesGetViewers type represents `stories_getViewers_response` API response object
type StoriesGetViewers struct {
	Count int   `json:"count"` // Viewers count
	Items []int `json:"items"`
}

// StoriesGetExtended type represents `stories_get_extended_response` API response object
type StoriesGetExtended struct {
	Count    int                    `json:"count"` // Stories count
	Groups   []objects.GroupsGroup  `json:"groups"`
	Items    []objects.StoriesStory `json:"items"`
	Profiles []objects.UsersUser    `json:"profiles"`
}

// StoriesGet type represents `stories_get_response` API response object
type StoriesGet struct {
	Count     int                       `json:"count"` // Stories count
	Items     []objects.StoriesStory    `json:"items"`
	PromoData objects.StoriesPromoBlock `json:"promo_data"`
}

// StoriesUpload type represents `stories_upload_response` API response object
type StoriesUpload struct {
	Story objects.StoriesStory `json:"story"`
}
