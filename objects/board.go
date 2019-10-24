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
// `board` group of objects
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

// BoardDefaultOrder type represents `board_default_order` API object
type BoardDefaultOrder int // Sort type

// BoardTopic type represents `board_topic` API object
type BoardTopic struct {
	Comments  int          `json:"comments"`   // Comments number
	Created   int          `json:"created"`    // Date when the topic has been created in Unixtime
	CreatedBy int          `json:"created_by"` // Creator ID
	Id        int          `json:"id"`         // Topic ID
	IsClosed  *BaseBoolInt `json:"is_closed"`  // Information whether the topic is closed
	IsFixed   *BaseBoolInt `json:"is_fixed"`   // Information whether the topic is fixed
	Title     string       `json:"title"`      // Topic title
	Updated   int          `json:"updated"`    // Date when the topic has been updated in Unixtime
	UpdatedBy int          `json:"updated_by"` // ID of user who updated the topic
}

// BoardTopicComment type represents `board_topic_comment` API object
type BoardTopicComment struct {
	Attachments []*WallCommentAttachment `json:"attachments"`
	Date        int                      `json:"date"`        // Date when the comment has been added in Unixtime
	FromId      int                      `json:"from_id"`     // Author ID
	Id          int                      `json:"id"`          // Comment ID
	RealOffset  int                      `json:"real_offset"` // Real position of the comment
	Text        string                   `json:"text"`        // Comment text
}

// BoardTopicPoll type represents `board_topic_poll` API object
type BoardTopicPoll struct {
	AnswerId int            `json:"answer_id"` // Current user's answer ID
	Answers  []*PollsAnswer `json:"answers"`
	Created  int            `json:"created"`   // Date when poll has been created in Unixtime
	IsClosed *BaseBoolInt   `json:"is_closed"` // Information whether the poll is closed
	OwnerId  int            `json:"owner_id"`  // Poll owner's ID
	PollId   int            `json:"poll_id"`   // Poll ID
	Question string         `json:"question"`  // Poll question
	Votes    string         `json:"votes"`     // Votes number
}
