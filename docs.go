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
// Code generator location: https://github.com/Burmuley/go-vkapi-gen                                       //
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

package go_vkapi

import (
	"github.com/Burmuley/go-vkapi/responses"
)

type Docs struct {
	*VKApi
}

/////////////////////////////////////////////////////////////
// `Docs` methods
/////////////////////////////////////////////////////////////

// Add - Copies a document to a user's or community's document list.
// Parameters:
//   * ownerId - ID of the user or community that owns the document. Use a negative value to designate a community ID.
//   * docId - Document ID.
//   * accessKey - Access key. This parameter is required if 'access_key' was returned with the document's data.
func (d *Docs) Add(ownerId int, docId int, accessKey string) (resp responses.DocsAdd, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	params["doc_id"] = docId

	if accessKey != "" {
		params["access_key"] = accessKey
	}

	err = d.SendObjRequest("docs.add", params, &resp)

	return
}

// Delete - Deletes a user or community document.
// Parameters:
//   * ownerId - ID of the user or community that owns the document. Use a negative value to designate a community ID.
//   * docId - Document ID.
func (d *Docs) Delete(ownerId int, docId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	params["doc_id"] = docId

	err = d.SendObjRequest("docs.delete", params, &resp)

	return
}

// Edit - Edits a document.
// Parameters:
//   * ownerId - User ID or community ID. Use a negative value to designate a community ID.
//   * docId - Document ID.
//   * title - Document title.
//   * tags - Document tags.
func (d *Docs) Edit(ownerId int, docId int, title string, tags []string) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	params["doc_id"] = docId

	if title != "" {
		params["title"] = title
	}

	if len(tags) > 0 {
		params["tags"] = SliceToString(tags)
	}

	err = d.SendObjRequest("docs.edit", params, &resp)

	return
}

// Get - Returns detailed information about user or community documents.
// Parameters:
//   * count - Number of documents to return. By default, all documents.
//   * offset - Offset needed to return a specific subset of documents.
//   * pType - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * ownerId - ID of the user or community that owns the documents. Use a negative value to designate a community ID.
func (d *Docs) Get(count int, offset int, pType int, ownerId int) (resp responses.DocsGet, err error) {
	params := map[string]interface{}{}

	if count > 0 {
		params["count"] = count
	}

	if offset > 0 {
		params["offset"] = offset
	}

	if pType > 0 {
		params["type"] = pType
	}

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	err = d.SendObjRequest("docs.get", params, &resp)

	return
}

// Getbyid - Returns information about documents by their IDs.
// Parameters:
//   * docs - Document IDs. Example: , "66748_91488,66748_91455",
func (d *Docs) Getbyid(docs []string) (resp responses.DocsGetbyid, err error) {
	params := map[string]interface{}{}

	params["docs"] = SliceToString(docs)

	err = d.SendObjRequest("docs.getById", params, &resp)

	return
}

// Getmessagesuploadserver - Returns the server address for document upload.
// Parameters:
//   * pType - Document type.
//   * peerId - Destination ID. "For user: 'User ID', e.g. '12345'. For chat: '2000000000' + 'Chat ID', e.g. '2000000001'. For community: '- Community ID', e.g. '-12345'. "
func (d *Docs) Getmessagesuploadserver(pType string, peerId int) (resp responses.BaseGetuploadserver, err error) {
	params := map[string]interface{}{}

	if pType != "" {
		params["type"] = pType
	}

	if peerId > 0 {
		params["peer_id"] = peerId
	}

	err = d.SendObjRequest("docs.getMessagesUploadServer", params, &resp)

	return
}

// Gettypes - Returns documents types available for current user.
// Parameters:
//   * ownerId - ID of the user or community that owns the documents. Use a negative value to designate a community ID.
func (d *Docs) Gettypes(ownerId int) (resp responses.DocsGettypes, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	err = d.SendObjRequest("docs.getTypes", params, &resp)

	return
}

// Getuploadserver - Returns the server address for document upload.
// Parameters:
//   * groupId - Community ID (if the document will be uploaded to the community).
func (d *Docs) Getuploadserver(groupId int) (resp responses.DocsGetuploadserver, err error) {
	params := map[string]interface{}{}

	if groupId > 0 {
		params["group_id"] = groupId
	}

	err = d.SendObjRequest("docs.getUploadServer", params, &resp)

	return
}

// Getwalluploadserver - Returns the server address for document upload onto a user's or community's wall.
// Parameters:
//   * groupId - Community ID (if the document will be uploaded to the community).
func (d *Docs) Getwalluploadserver(groupId int) (resp responses.BaseGetuploadserver, err error) {
	params := map[string]interface{}{}

	if groupId > 0 {
		params["group_id"] = groupId
	}

	err = d.SendObjRequest("docs.getWallUploadServer", params, &resp)

	return
}

// Save - Saves a document after [vk.com/dev/upload_files_2|uploading it to a server].
// Parameters:
//   * file - This parameter is returned when the file is [vk.com/dev/upload_files_2|uploaded to the server].
//   * title - Document title.
//   * tags - Document tags.
func (d *Docs) Save(file string, title string, tags string) (resp responses.DocsSave, err error) {
	params := map[string]interface{}{}

	params["file"] = file

	if title != "" {
		params["title"] = title
	}

	if tags != "" {
		params["tags"] = tags
	}

	err = d.SendObjRequest("docs.save", params, &resp)

	return
}

// Search - Returns a list of documents matching the search criteria.
// Parameters:
//   * q - Search query string.
//   * searchOwn - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * count - Number of results to return.
//   * offset - Offset needed to return a specific subset of results.
func (d *Docs) Search(q string, searchOwn bool, count int, offset int) (resp responses.DocsSearch, err error) {
	params := map[string]interface{}{}

	params["q"] = q

	params["search_own"] = searchOwn

	if count > 0 {
		params["count"] = count
	}

	if offset > 0 {
		params["offset"] = offset
	}

	err = d.SendObjRequest("docs.search", params, &resp)

	return
}
