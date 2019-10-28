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
// `notes` group of objects
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

// NotesNote type represents `notes_note` API object
type NotesNote struct {
	CanComment   BaseBoolInt `json:"can_comment"` // Information whether current user can comment the note
	Comments     int         `json:"comments"`    // Comments number
	Date         int         `json:"date"`        // Date when the note has been created in Unixtime
	Id           int         `json:"id"`          // Note ID
	OwnerId      int         `json:"owner_id"`    // Note owner's ID
	ReadComments int         `json:"read_comments"`
	Text         string      `json:"text"`      // Note text
	TextWiki     string      `json:"text_wiki"` // Note text in wiki format
	Title        string      `json:"title"`     // Note title
	ViewUrl      string      `json:"view_url"`  // URL of the page with note preview
}

// NotesNoteComment type represents `notes_note_comment` API object
type NotesNoteComment struct {
	Date    int    `json:"date"`     // Date when the comment has beed added in Unixtime
	Id      int    `json:"id"`       // Comment ID
	Message string `json:"message"`  // Comment text
	Nid     int    `json:"nid"`      // Note ID
	Oid     int    `json:"oid"`      // Note ID
	ReplyTo int    `json:"reply_to"` // ID of replied comment
	Uid     int    `json:"uid"`      // Comment author's ID
}
