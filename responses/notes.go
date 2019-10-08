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
// `notes` group of responses
/////////////////////////////////////////////////////////////

// NotesGetComments type represents `notes_getComments_response` API response object
type NotesGetComments struct {
	Count int                        `json:"count"` // Total number
	Items []objects.NotesNoteComment `json:"items"`
}

// NotesCreateComment type represents `notes_createComment_response` API response object
type NotesCreateComment int // Comment ID

// NotesGet type represents `notes_get_response` API response object
type NotesGet struct {
	Count int                 `json:"count"` // Total number
	Items []objects.NotesNote `json:"items"`
}

// NotesGetById type represents `notes_getById_response` API response object
type NotesGetById objects.NotesNote

// NotesAdd type represents `notes_add_response` API response object
type NotesAdd int // Note ID
