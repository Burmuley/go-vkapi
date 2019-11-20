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
// Code generator location: https://github.com/Burmuley/go-vkapi-gen                                       //
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

package responses

import (
	"github.com/Burmuley/go-vkapi/objects"
)

/////////////////////////////////////////////////////////////
// `wall` group of responses
/////////////////////////////////////////////////////////////

// WallCreatecomment type represents `wall_createComment_response` API response object
type WallCreatecomment struct {
	CommentId int `json:"comment_id"` // Created comment ID
}

// WallEdit type represents `wall_edit_response` API response object
type WallEdit struct {
	PostId int `json:"post_id"` // Edited post ID
}

// WallGetbyidExtended type represents `wall_getById_extended_response` API response object
type WallGetbyidExtended struct {
	Groups   []objects.GroupsGroupFull  `json:"groups"`
	Items    []objects.WallWallpostFull `json:"items"`
	Profiles []objects.UsersUserFull    `json:"profiles"`
}

// WallGetbyid type represents `wall_getById_response` API response object
type WallGetbyid objects.WallWallpostFull

// WallGetcommentsExtended type represents `wall_getComments_extended_response` API response object
type WallGetcommentsExtended struct {
	CanPost           bool                      `json:"can_post"`            // Information whether current user can comment the post
	Count             int                       `json:"count"`               // Total number
	CurrentLevelCount int                       `json:"current_level_count"` // Count of replies of current level
	Groups            []objects.GroupsGroup     `json:"groups"`
	GroupsCanPost     bool                      `json:"groups_can_post"` // Information whether groups can comment the post
	Items             []objects.WallWallComment `json:"items"`
	Profiles          []objects.UsersUser       `json:"profiles"`
	ShowReplyButton   bool                      `json:"show_reply_button"`
}

// WallGetcomments type represents `wall_getComments_response` API response object
type WallGetcomments struct {
	CanPost           bool                      `json:"can_post"`            // Information whether current user can comment the post
	Count             int                       `json:"count"`               // Total number
	CurrentLevelCount int                       `json:"current_level_count"` // Count of replies of current level
	GroupsCanPost     bool                      `json:"groups_can_post"`     // Information whether groups can comment the post
	Items             []objects.WallWallComment `json:"items"`
}

// WallGetreposts type represents `wall_getReposts_response` API response object
type WallGetreposts struct {
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

// WallPostadsstealth type represents `wall_postAdsStealth_response` API response object
type WallPostadsstealth struct {
	PostId int `json:"post_id"` // Created post ID
}

// WallPost type represents `wall_post_response` API response object
type WallPost struct {
	PostId int `json:"post_id"` // Created post ID
}

// WallRepost type represents `wall_repost_response` API response object
type WallRepost struct {
	LikesCount   int                    `json:"likes_count"`   // Reposts number
	PostId       int                    `json:"post_id"`       // Created post ID
	RepostsCount int                    `json:"reposts_count"` // Reposts number
	Success      objects.BaseOkResponse `json:"success"`
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
