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

package objects

import (
	"encoding/json"
)

/////////////////////////////////////////////////////////////////////////////////////////////////////////////
// `polls` group of objects
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

// PollsAnswer type represents `polls_answer` API object
type PollsAnswer struct {
	Id    int         `json:"id"`    // Answer ID
	Rate  json.Number `json:"rate"`  // Answer rate in percents
	Text  string      `json:"text"`  // Answer text
	Votes int         `json:"votes"` // Votes number
}

// PollsPoll type represents `polls_poll` API object
type PollsPoll struct {
	Anonymous bool          `json:"anonymous"` // Information whether the field is anonymous
	AnswerId  int           `json:"answer_id"` // Current user's answer ID
	Answers   []PollsAnswer `json:"answers"`
	Created   int           `json:"created"`  // Date when poll has been created in Unixtime
	Id        int           `json:"id"`       // Poll ID
	OwnerId   int           `json:"owner_id"` // Poll owner's ID
	Question  string        `json:"question"` // Poll question
	Votes     string        `json:"votes"`    // Votes number
}

// PollsVoters type represents `polls_voters` API object
type PollsVoters struct {
	AnswerId int              `json:"answer_id"` // Answer ID
	Users    PollsVotersUsers `json:"users"`
}

// PollsVotersUsers type represents `polls_voters_users` API object
type PollsVotersUsers struct {
	Count int   `json:"count"` // Votes number
	Items []int `json:"items"`
}
