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
// `docs` group of objects
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

// DocsDoc type represents `docs_doc` API object
type DocsDoc struct {
	AccessKey  string          `json:"access_key"` // Access key for the document
	Date       int             `json:"date"`       // Date when file has been uploaded in Unixtime
	Ext        string          `json:"ext"`        // File extension
	Id         int             `json:"id"`         // Document ID
	IsLicensed *BaseBoolInt    `json:"is_licensed"`
	OwnerId    int             `json:"owner_id"` // Document owner ID
	Preview    *DocsDocPreview `json:"preview"`
	Size       int             `json:"size"`  // File size in bites
	Title      string          `json:"title"` // Document title
	Type       int             `json:"type"`  // Document type
	Url        string          `json:"url"`   // File URL
}

// DocsDocAttachmentType type represents `docs_doc_attachment_type` API object
type DocsDocAttachmentType string // Doc attachment type

// DocsDocPreview type represents `docs_doc_preview` API object
type DocsDocPreview struct {
	Photo *DocsDocPreviewPhoto `json:"photo"`
	Video *DocsDocPreviewVideo `json:"video"`
}

// DocsDocPreviewPhoto type represents `docs_doc_preview_photo` API object
type DocsDocPreviewPhoto struct {
	Sizes []*PhotosPhotoSizes `json:"sizes"`
}

// DocsDocPreviewVideo type represents `docs_doc_preview_video` API object
type DocsDocPreviewVideo struct {
	Filesize int    `json:"filesize"` // Video file size in bites
	Height   int    `json:"height"`   // Video's height in pixels
	Src      string `json:"src"`      // Video URL
	Width    int    `json:"width"`    // Video's width in pixels
}

// DocsDocTypes type represents `docs_doc_types` API object
type DocsDocTypes struct {
	Count int    `json:"count"` // Number of docs
	Id    int    `json:"id"`    // Doc type ID
	Title string `json:"title"` // Doc type title
}

// DocsDocUploadResponse type represents `docs_doc_upload_response` API object
type DocsDocUploadResponse struct {
	File string `json:"file"` // Uploaded file data
}
