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
// `wall` group of responses
/////////////////////////////////////////////////////////////

// WallCreateComment type represents `wall_createComment_response` API response object
type WallCreateComment struct {
	CommentId int `json:"comment_id"` // Created comment ID
}

// WallEdit type represents `wall_edit_response` API response object
type WallEdit struct {
	PostId int `json:"post_id"` // Edited post ID
}

// WallGetByIdExtended type represents `wall_getById_extended_response` API response object
type WallGetByIdExtended struct {
	Groups   []objects.GroupsGroupFull  `json:"groups"`
	Items    []objects.WallWallpostFull `json:"items"`
	Profiles []objects.UsersUserFull    `json:"profiles"`
}

// WallGetById type represents `wall_getById_response` API response object
type WallGetById objects.WallWallpostFull

// WallGetCommentsExtended type represents `wall_getComments_extended_response` API response object
type WallGetCommentsExtended struct {
	CanPost           bool                      `json:"can_post"`            // Information whether current user can comment the post
	Count             int                       `json:"count"`               // Total number
	CurrentLevelCount int                       `json:"current_level_count"` // Count of replies of current level
	Groups            []objects.GroupsGroup     `json:"groups"`
	GroupsCanPost     bool                      `json:"groups_can_post"` // Information whether groups can comment the post
	Items             []objects.WallWallComment `json:"items"`
	Profiles          []objects.UsersUser       `json:"profiles"`
	ShowReplyButton   bool                      `json:"show_reply_button"`
}

// WallGetComments type represents `wall_getComments_response` API response object
type WallGetComments struct {
	CanPost           bool                      `json:"can_post"`            // Information whether current user can comment the post
	Count             int                       `json:"count"`               // Total number
	CurrentLevelCount int                       `json:"current_level_count"` // Count of replies of current level
	GroupsCanPost     bool                      `json:"groups_can_post"`     // Information whether groups can comment the post
	Items             []objects.WallWallComment `json:"items"`
}

// WallGetReposts type represents `wall_getReposts_response` API response object
type WallGetReposts struct {
	Groups   []objects.GroupsGroup      `json:"groups"`
	Items    []objects.WallWallpostFull `json:"items"`
	Profiles []objects.UsersUser        `json:"profiles"`
}

// WallGetExtended type represents `wall_get_extended_response` API response object
type WallGetExtended struct {
	Count    int                        `json:"count"` // Total number
	Groups   []objects.GroupsGroupFull  `json:"groups"`
	Items    []objects.WallWallpostFull `json:"items"`
	Profiles []objects.UsersUserFull    `json:"profiles"`
}

// WallGet type represents `wall_get_response` API response object
type WallGet struct {
	Count int                        `json:"count"` // Total number
	Items []objects.WallWallpostFull `json:"items"`
}

// WallPostAdsStealth type represents `wall_postAdsStealth_response` API response object
type WallPostAdsStealth struct {
	PostId int `json:"post_id"` // Created post ID
}

// WallPost type represents `wall_post_response` API response object
type WallPost struct {
	PostId int `json:"post_id"` // Created post ID
}

// WallRepost type represents `wall_repost_response` API response object
type WallRepost struct {
	LikesCount   int            `json:"likes_count"`   // Reposts number
	PostId       int            `json:"post_id"`       // Created post ID
	RepostsCount int            `json:"reposts_count"` // Reposts number
	Success      objects.BaseOk `json:"success"`
}

// WallSearchExtended type represents `wall_search_extended_response` API response object
type WallSearchExtended struct {
	Count    int                        `json:"count"` // Total number
	Groups   []objects.GroupsGroupFull  `json:"groups"`
	Items    []objects.WallWallpostFull `json:"items"`
	Profiles []objects.UsersUserFull    `json:"profiles"`
}

// WallSearch type represents `wall_search_response` API response object
type WallSearch struct {
	Count int                        `json:"count"` // Total number
	Items []objects.WallWallpostFull `json:"items"`
}
