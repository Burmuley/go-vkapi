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
// Source schema can be found at https://github.com/VKCOM/vk-api-schema/blob/master/methods.json           //
// Code generator location: https://gitlab.com/Burmuley/go-vkapi-gen                                       //
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

package go_vkapi

import (
	"gitlab.com/Burmuley/go-vkapi/objects"
	"gitlab.com/Burmuley/go-vkapi/responses"
)

type Polls struct {
	*VKApi
}

/////////////////////////////////////////////////////////////
// `Polls` methods
/////////////////////////////////////////////////////////////

// Addvote - Adds the current user's vote to the selected answer in the poll.
// Parameters:
//   * ownerId - ID of the user or community that owns the poll. Use a negative value to designate a community ID.
//   * pollId - Poll ID.
//   * answerIds - NO DESCRIPTION IN JSON SCHEMA
//   * isBoard - NO DESCRIPTION IN JSON SCHEMA
func (p Polls) Addvote(ownerId int, pollId int, answerIds []int, isBoard bool) (resp responses.PollsAddvot, err error) {
	params := map[string]interface{}{}

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	params["poll_id"] = pollId

	params["answer_ids"] = SliceToString(answerIds)

	params["is_board"] = isBoard

	err = p.SendObjRequest("polls.addVote", params, &resp)

	return
}

// Create - Creates polls that can be attached to the users' or communities' posts.
// Parameters:
//   * question - question text
//   * isAnonymous - '1' – anonymous poll, participants list is hidden,, '0' – public poll, participants list is available,, Default value is '0'.
//   * isMultiple - NO DESCRIPTION IN JSON SCHEMA
//   * endDate - NO DESCRIPTION IN JSON SCHEMA
//   * ownerId - If a poll will be added to a communty it is required to send a negative group identifier. Current user by default.
//   * addAnswers - available answers list, for example: " ["yes","no","maybe"]", There can be from 1 to 10 answers.
//   * photoId - NO DESCRIPTION IN JSON SCHEMA
//   * backgroundId - NO DESCRIPTION IN JSON SCHEMA
func (p Polls) Create(question string, isAnonymous bool, isMultiple bool, endDate int, ownerId int, addAnswers string, photoId int, backgroundId string) (resp responses.PollsCreat, err error) {
	params := map[string]interface{}{}

	if question != "" {
		params["question"] = question
	}

	params["is_anonymous"] = isAnonymous

	params["is_multiple"] = isMultiple

	if endDate > 0 {
		params["end_date"] = endDate
	}

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	if addAnswers != "" {
		params["add_answers"] = addAnswers
	}

	if photoId > 0 {
		params["photo_id"] = photoId
	}

	if backgroundId != "" {
		params["background_id"] = backgroundId
	}

	err = p.SendObjRequest("polls.create", params, &resp)

	return
}

// Deletevote - Deletes the current user's vote from the selected answer in the poll.
// Parameters:
//   * ownerId - ID of the user or community that owns the poll. Use a negative value to designate a community ID.
//   * pollId - Poll ID.
//   * answerId - Answer ID.
//   * isBoard - NO DESCRIPTION IN JSON SCHEMA
func (p Polls) Deletevote(ownerId int, pollId int, answerId int, isBoard bool) (resp responses.PollsDeletevot, err error) {
	params := map[string]interface{}{}

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	params["poll_id"] = pollId

	params["answer_id"] = answerId

	params["is_board"] = isBoard

	err = p.SendObjRequest("polls.deleteVote", params, &resp)

	return
}

// Edit - Edits created polls
// Parameters:
//   * ownerId - poll owner id
//   * pollId - edited poll's id
//   * question - new question text
//   * addAnswers - answers list, for example: , "["yes","no","maybe"]"
//   * editAnswers - object containing answers that need to be edited,, key – answer id, value – new answer text. Example: {"382967099":"option1", "382967103":"option2"}"
//   * deleteAnswers - list of answer ids to be deleted. For example: "[382967099, 382967103]"
//   * endDate - NO DESCRIPTION IN JSON SCHEMA
//   * photoId - NO DESCRIPTION IN JSON SCHEMA
//   * backgroundId - NO DESCRIPTION IN JSON SCHEMA
func (p Polls) Edit(ownerId int, pollId int, question string, addAnswers string, editAnswers string, deleteAnswers string, endDate int, photoId int, backgroundId string) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	params["poll_id"] = pollId

	if question != "" {
		params["question"] = question
	}

	if addAnswers != "" {
		params["add_answers"] = addAnswers
	}

	if editAnswers != "" {
		params["edit_answers"] = editAnswers
	}

	if deleteAnswers != "" {
		params["delete_answers"] = deleteAnswers
	}

	if endDate > 0 {
		params["end_date"] = endDate
	}

	if photoId > 0 {
		params["photo_id"] = photoId
	}

	if backgroundId != "" {
		params["background_id"] = backgroundId
	}

	err = p.SendObjRequest("polls.edit", params, &resp)

	return
}

// Getbyid - Returns detailed information about a poll by its ID.
// Parameters:
//   * ownerId - ID of the user or community that owns the poll. Use a negative value to designate a community ID.
//   * isBoard - '1' – poll is in a board, '0' – poll is on a wall. '0' by default.
//   * pollId - Poll ID.
//   * extended - NO DESCRIPTION IN JSON SCHEMA
//   * friendsCount - NO DESCRIPTION IN JSON SCHEMA
//   * fields - NO DESCRIPTION IN JSON SCHEMA
//   * nameCase - NO DESCRIPTION IN JSON SCHEMA
func (p Polls) Getbyid(ownerId int, isBoard bool, pollId int, friendsCount int, fields []string, nameCase string) (resp responses.PollsGetbyid, err error) {
	params := map[string]interface{}{}

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	params["is_board"] = isBoard

	params["poll_id"] = pollId

	if friendsCount > 0 {
		params["friends_count"] = friendsCount
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	if nameCase != "" {
		params["name_case"] = nameCase
	}

	err = p.SendObjRequest("polls.getById", params, &resp)

	return
}

// Getvoters - Returns a list of IDs of users who selected specific answers in the poll.
// Parameters:
//   * ownerId - ID of the user or community that owns the poll. Use a negative value to designate a community ID.
//   * pollId - Poll ID.
//   * answerIds - Answer IDs.
//   * isBoard - NO DESCRIPTION IN JSON SCHEMA
//   * friendsOnly - '1' — to return only current user's friends, '0' — to return all users (default),
//   * offset - Offset needed to return a specific subset of voters. '0' — (default)
//   * count - Number of user IDs to return (if the 'friends_only' parameter is not set, maximum '1000', otherwise '10'). '100' — (default)
//   * fields - Profile fields to return. Sample values: 'nickname', 'screen_name', 'sex', 'bdate (birthdate)', 'city', 'country', 'timezone', 'photo', 'photo_medium', 'photo_big', 'has_mobile', 'rate', 'contacts', 'education', 'online', 'counters'.
//   * nameCase - Case for declension of user name and surname: , 'nom' — nominative (default) , 'gen' — genitive , 'dat' — dative , 'acc' — accusative , 'ins' — instrumental , 'abl' — prepositional
func (p Polls) Getvoters(ownerId int, pollId int, answerIds []int, isBoard bool, friendsOnly bool, offset int, count int, fields []objects.UsersFields, nameCase string) (resp responses.PollsGetvoter, err error) {
	params := map[string]interface{}{}

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	params["poll_id"] = pollId

	params["answer_ids"] = SliceToString(answerIds)

	params["is_board"] = isBoard

	params["friends_only"] = friendsOnly

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	if nameCase != "" {
		params["name_case"] = nameCase
	}

	err = p.SendObjRequest("polls.getVoters", params, &resp)

	return
}
