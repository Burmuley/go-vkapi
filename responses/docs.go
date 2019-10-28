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
// `docs` group of responses
/////////////////////////////////////////////////////////////

// DocsAdd type represents `docs_add_response` API response object
type DocsAdd struct {
	Id int `json:"id"` // Doc ID
}

// DocsGetbyid type represents `docs_getById_response` API response object
type DocsGetbyid objects.DocsDoc

// DocsGettypes type represents `docs_getTypes_response` API response object
type DocsGettypes struct {
	Count int                    `json:"count"` // Total number
	Items []objects.DocsDocTypes `json:"items"`
}

// DocsGetuploadserver type represents `docs_getUploadServer` API response object
type DocsGetuploadserver objects.BaseUploadServer

// DocsGet type represents `docs_get_response` API response object
type DocsGet struct {
	Count int               `json:"count"` // Total number
	Items []objects.DocsDoc `json:"items"`
}

// DocsSave type represents `docs_save_response` API response object
type DocsSave struct {
	AudioMessage objects.MessagesAudioMessage  `json:"audio_message"`
	Doc          objects.DocsDoc               `json:"doc"`
	Graffiti     objects.MessagesGraffiti      `json:"graffiti"`
	Type         objects.DocsDocAttachmentType `json:"type"`
}

// DocsSearch type represents `docs_search_response` API response object
type DocsSearch struct {
	Count int               `json:"count"` // Total number
	Items []objects.DocsDoc `json:"items"`
}
