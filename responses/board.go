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
// `board` group of responses
/////////////////////////////////////////////////////////////

// BoardAddTopic type represents `board_addTopic_response` API response object
type BoardAddTopic int // Topic ID

// BoardCreateComment type represents `board_createComment_response` API response object
type BoardCreateComment int // Comment ID

// BoardGetCommentsExtended type represents `board_getComments_extended_response` API response object
type BoardGetCommentsExtended struct {
	Count    int                         `json:"count"` // Total number
	Groups   []objects.GroupsGroup       `json:"groups"`
	Items    []objects.BoardTopicComment `json:"items"`
	Poll     objects.BoardTopicPoll      `json:"poll"`
	Profiles []objects.UsersUser         `json:"profiles"`
}

// BoardGetComments type represents `board_getComments_response` API response object
type BoardGetComments struct {
	Count int                         `json:"count"` // Total number
	Items []objects.BoardTopicComment `json:"items"`
	Poll  objects.BoardTopicPoll      `json:"poll"`
}

// BoardGetTopicsExtended type represents `board_getTopics_extended_response` API response object
type BoardGetTopicsExtended struct {
	CanAddTopics objects.BaseBoolInt       `json:"can_add_topics"` // Information whether current user can add topic
	Count        int                       `json:"count"`          // Total number
	DefaultOrder objects.BoardDefaultOrder `json:"default_order"`
	Items        []objects.BoardTopic      `json:"items"`
	Profiles     []objects.UsersUserMin    `json:"profiles"`
}

// BoardGetTopics type represents `board_getTopics_response` API response object
type BoardGetTopics struct {
	CanAddTopics objects.BaseBoolInt       `json:"can_add_topics"` // Information whether current user can add topic
	Count        int                       `json:"count"`          // Total number
	DefaultOrder objects.BoardDefaultOrder `json:"default_order"`
	Items        []objects.BoardTopic      `json:"items"`
}
