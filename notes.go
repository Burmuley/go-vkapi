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
	"gitlab.com/Burmuley/go-vkapi/responses"
)

type Notes struct {
	*VKApi
}

/////////////////////////////////////////////////////////////
// `Notes` methods
/////////////////////////////////////////////////////////////

// Add - Creates a new note for the current user.
// Parameters:
//   * title - Note title.
//   * text - Note text.
//   * privacyView - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * privacyComment - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (n Notes) Add(title string, text string, privacyView []string, privacyComment []string) (resp responses.NotesAdd, err error) {
	params := map[string]interface{}{}

	params["title"] = title

	params["text"] = text

	if len(privacyView) > 0 {
		params["privacy_view"] = SliceToString(privacyView)
	}

	if len(privacyComment) > 0 {
		params["privacy_comment"] = SliceToString(privacyComment)
	}

	err = n.SendObjRequest("notes.add", params, &resp)

	return
}

// Createcomment - Adds a new comment on a note.
// Parameters:
//   * noteId - Note ID.
//   * ownerId - Note owner ID.
//   * replyTo - ID of the user to whom the reply is addressed (if the comment is a reply to another comment).
//   * message - Comment text.
//   * guid - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (n Notes) Createcomment(noteId int, ownerId int, replyTo int, message string, guid string) (resp responses.NotesCreatecomment, err error) {
	params := map[string]interface{}{}

	params["note_id"] = noteId

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	if replyTo > 0 {
		params["reply_to"] = replyTo
	}

	params["message"] = message

	if guid != "" {
		params["guid"] = guid
	}

	err = n.SendObjRequest("notes.createComment", params, &resp)

	return
}

// Delete - Deletes a note of the current user.
// Parameters:
//   * noteId - Note ID.
func (n Notes) Delete(noteId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["note_id"] = noteId

	err = n.SendObjRequest("notes.delete", params, &resp)

	return
}

// Deletecomment - Deletes a comment on a note.
// Parameters:
//   * commentId - Comment ID.
//   * ownerId - Note owner ID.
func (n Notes) Deletecomment(commentId int, ownerId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["comment_id"] = commentId

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	err = n.SendObjRequest("notes.deleteComment", params, &resp)

	return
}

// Edit - Edits a note of the current user.
// Parameters:
//   * noteId - Note ID.
//   * title - Note title.
//   * text - Note text.
//   * privacyView - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * privacyComment - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (n Notes) Edit(noteId int, title string, text string, privacyView []string, privacyComment []string) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["note_id"] = noteId

	params["title"] = title

	params["text"] = text

	if len(privacyView) > 0 {
		params["privacy_view"] = SliceToString(privacyView)
	}

	if len(privacyComment) > 0 {
		params["privacy_comment"] = SliceToString(privacyComment)
	}

	err = n.SendObjRequest("notes.edit", params, &resp)

	return
}

// Editcomment - Edits a comment on a note.
// Parameters:
//   * commentId - Comment ID.
//   * ownerId - Note owner ID.
//   * message - New comment text.
func (n Notes) Editcomment(commentId int, ownerId int, message string) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["comment_id"] = commentId

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	params["message"] = message

	err = n.SendObjRequest("notes.editComment", params, &resp)

	return
}

// Get - Returns a list of notes created by a user.
// Parameters:
//   * noteIds - Note IDs.
//   * userId - Note owner ID.
//   * offset - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * count - Number of notes to return.
//   * sort - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (n Notes) Get(noteIds []int, userId int, offset int, count int, sort int) (resp responses.NotesGet, err error) {
	params := map[string]interface{}{}

	if len(noteIds) > 0 {
		params["note_ids"] = SliceToString(noteIds)
	}

	if userId > 0 {
		params["user_id"] = userId
	}

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	if sort > 0 {
		params["sort"] = sort
	}

	err = n.SendObjRequest("notes.get", params, &resp)

	return
}

// Getbyid - Returns a note by its ID.
// Parameters:
//   * noteId - Note ID.
//   * ownerId - Note owner ID.
//   * needWiki - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (n Notes) Getbyid(noteId int, ownerId int, needWiki bool) (resp responses.NotesGetbyid, err error) {
	params := map[string]interface{}{}

	params["note_id"] = noteId

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	params["need_wiki"] = needWiki

	err = n.SendObjRequest("notes.getById", params, &resp)

	return
}

// Getcomments - Returns a list of comments on a note.
// Parameters:
//   * noteId - Note ID.
//   * ownerId - Note owner ID.
//   * sort - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * offset - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * count - Number of comments to return.
func (n Notes) Getcomments(noteId int, ownerId int, sort int, offset int, count int) (resp responses.NotesGetcomments, err error) {
	params := map[string]interface{}{}

	params["note_id"] = noteId

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	if sort > 0 {
		params["sort"] = sort
	}

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	err = n.SendObjRequest("notes.getComments", params, &resp)

	return
}

// Restorecomment - Restores a deleted comment on a note.
// Parameters:
//   * commentId - Comment ID.
//   * ownerId - Note owner ID.
func (n Notes) Restorecomment(commentId int, ownerId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["comment_id"] = commentId

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	err = n.SendObjRequest("notes.restoreComment", params, &resp)

	return
}
