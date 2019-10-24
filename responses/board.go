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
// `board` group of responses
/////////////////////////////////////////////////////////////

// BoardAddtopic type represents `board_addTopic_response` API response object
type BoardAddtopic int // Topic ID

// BoardCreatecomment type represents `board_createComment_response` API response object
type BoardCreatecomment int // Comment ID

// BoardGetcommentsExtended type represents `board_getComments_extended_response` API response object
type BoardGetcommentsExtended struct {
	Count    int                          `json:"count"` // Total number
	Groups   []*objects.GroupsGroup       `json:"groups"`
	Items    []*objects.BoardTopicComment `json:"items"`
	Poll     *objects.BoardTopicPoll      `json:"poll"`
	Profiles []*objects.UsersUser         `json:"profiles"`
}

// BoardGetcomment type represents `board_getComments_response` API response object
type BoardGetcomment struct {
	Count int                          `json:"count"` // Total number
	Items []*objects.BoardTopicComment `json:"items"`
	Poll  *objects.BoardTopicPoll      `json:"poll"`
}

// BoardGettopicsExtended type represents `board_getTopics_extended_response` API response object
type BoardGettopicsExtended struct {
	CanAddTopics *objects.BaseBoolInt       `json:"can_add_topics"` // Information whether current user can add topic
	Count        int                        `json:"count"`          // Total number
	DefaultOrder *objects.BoardDefaultOrder `json:"default_order"`
	Items        []*objects.BoardTopic      `json:"items"`
	Profiles     []*objects.UsersUserMin    `json:"profiles"`
}

// BoardGettopic type represents `board_getTopics_response` API response object
type BoardGettopic struct {
	CanAddTopics *objects.BaseBoolInt       `json:"can_add_topics"` // Information whether current user can add topic
	Count        int                        `json:"count"`          // Total number
	DefaultOrder *objects.BoardDefaultOrder `json:"default_order"`
	Items        []*objects.BoardTopic      `json:"items"`
}
