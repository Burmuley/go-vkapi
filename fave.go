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

type Fave struct {
	*VKApi
}

/////////////////////////////////////////////////////////////
// `Fave` methods
/////////////////////////////////////////////////////////////

// AddArticle - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * url - NO DESCRIPTION IN JSON SCHEMA
func (f *Fave) AddArticle(url string) (resp responses.BaseBool, err error) {
	params := map[string]interface{}{}

	params["url"] = url

	err = f.SendObjRequest("fave.addArticle", params, &resp)

	return
}

// AddLink - Adds a link to user faves.
// Parameters:
//   * link - Link URL.
func (f *Fave) AddLink(link string) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["link"] = link

	err = f.SendObjRequest("fave.addLink", params, &resp)

	return
}

// AddPage - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * userId - NO DESCRIPTION IN JSON SCHEMA
//   * groupId - NO DESCRIPTION IN JSON SCHEMA
func (f *Fave) AddPage(userId int, groupId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	if userId > 0 {
		params["user_id"] = userId
	}

	if groupId > 0 {
		params["group_id"] = groupId
	}

	err = f.SendObjRequest("fave.addPage", params, &resp)

	return
}

// AddPost - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * ownerId - NO DESCRIPTION IN JSON SCHEMA
//   * id - NO DESCRIPTION IN JSON SCHEMA
//   * accessKey - NO DESCRIPTION IN JSON SCHEMA
func (f *Fave) AddPost(ownerId int, id int, accessKey string) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	params["id"] = id

	if accessKey != "" {
		params["access_key"] = accessKey
	}

	err = f.SendObjRequest("fave.addPost", params, &resp)

	return
}

// AddProduct - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * ownerId - NO DESCRIPTION IN JSON SCHEMA
//   * id - NO DESCRIPTION IN JSON SCHEMA
//   * accessKey - NO DESCRIPTION IN JSON SCHEMA
func (f *Fave) AddProduct(ownerId int, id int, accessKey string) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	params["id"] = id

	if accessKey != "" {
		params["access_key"] = accessKey
	}

	err = f.SendObjRequest("fave.addProduct", params, &resp)

	return
}

// AddTag - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * name - NO DESCRIPTION IN JSON SCHEMA
func (f *Fave) AddTag(name string) (resp responses.FaveAddTag, err error) {
	params := map[string]interface{}{}

	if name != "" {
		params["name"] = name
	}

	err = f.SendObjRequest("fave.addTag", params, &resp)

	return
}

// AddVideo - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * ownerId - NO DESCRIPTION IN JSON SCHEMA
//   * id - NO DESCRIPTION IN JSON SCHEMA
//   * accessKey - NO DESCRIPTION IN JSON SCHEMA
func (f *Fave) AddVideo(ownerId int, id int, accessKey string) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	params["id"] = id

	if accessKey != "" {
		params["access_key"] = accessKey
	}

	err = f.SendObjRequest("fave.addVideo", params, &resp)

	return
}

// EditTag - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * id - NO DESCRIPTION IN JSON SCHEMA
//   * name - NO DESCRIPTION IN JSON SCHEMA
func (f *Fave) EditTag(id int, name string) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["id"] = id

	params["name"] = name

	err = f.SendObjRequest("fave.editTag", params, &resp)

	return
}

// Get - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * extended - '1' — to return additional 'wall', 'profiles', and 'groups' fields. By default: '0'.
//   * itemType - NO DESCRIPTION IN JSON SCHEMA
//   * tagId - Tag ID.
//   * offset - Offset needed to return a specific subset of users.
//   * count - Number of users to return.
//   * fields - NO DESCRIPTION IN JSON SCHEMA
//   * isFromSnackbar - NO DESCRIPTION IN JSON SCHEMA
func (f *Fave) Get(itemType string, tagId int, offset int, count int, fields string, isFromSnackbar bool) (resp responses.FaveGet, err error) {
	params := map[string]interface{}{}
	params["extended"] = "0"

	if itemType != "" {
		params["item_type"] = itemType
	}

	if tagId > 0 {
		params["tag_id"] = tagId
	}

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	if fields != "" {
		params["fields"] = fields
	}

	params["is_from_snackbar"] = isFromSnackbar

	err = f.SendObjRequest("fave.get", params, &resp)

	return
}

// GetExtended - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * extended - '1' — to return additional 'wall', 'profiles', and 'groups' fields. By default: '0'.
//   * itemType - NO DESCRIPTION IN JSON SCHEMA
//   * tagId - Tag ID.
//   * offset - Offset needed to return a specific subset of users.
//   * count - Number of users to return.
//   * fields - NO DESCRIPTION IN JSON SCHEMA
//   * isFromSnackbar - NO DESCRIPTION IN JSON SCHEMA
func (f *Fave) GetExtended(itemType string, tagId int, offset int, count int, fields string, isFromSnackbar bool) (resp responses.FaveGetExtended, err error) {
	params := map[string]interface{}{}
	params["extended"] = "1"

	if itemType != "" {
		params["item_type"] = itemType
	}

	if tagId > 0 {
		params["tag_id"] = tagId
	}

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	if fields != "" {
		params["fields"] = fields
	}

	params["is_from_snackbar"] = isFromSnackbar

	err = f.SendObjRequest("fave.get", params, &resp)

	return
}

// GetPages - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * offset - NO DESCRIPTION IN JSON SCHEMA
//   * count - NO DESCRIPTION IN JSON SCHEMA
//   * pType - NO DESCRIPTION IN JSON SCHEMA
//   * fields - NO DESCRIPTION IN JSON SCHEMA
//   * tagId - NO DESCRIPTION IN JSON SCHEMA
func (f *Fave) GetPages(offset int, count int, pType string, fields []objects.BaseUserGroupFields, tagId int) (resp responses.FaveGetPages, err error) {
	params := map[string]interface{}{}

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	if pType != "" {
		params["type"] = pType
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	if tagId > 0 {
		params["tag_id"] = tagId
	}

	err = f.SendObjRequest("fave.getPages", params, &resp)

	return
}

// GetTags - NO DESCRIPTION IN JSON SCHEMA
func (f *Fave) GetTags() (resp responses.FaveGetTags, err error) {
	params := map[string]interface{}{}

	err = f.SendObjRequest("fave.getTags", params, &resp)

	return
}

// MarkSeen - NO DESCRIPTION IN JSON SCHEMA
func (f *Fave) MarkSeen() (resp responses.BaseBool, err error) {
	params := map[string]interface{}{}

	err = f.SendObjRequest("fave.markSeen", params, &resp)

	return
}

// RemoveArticle - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * ownerId - NO DESCRIPTION IN JSON SCHEMA
//   * articleId - NO DESCRIPTION IN JSON SCHEMA
func (f *Fave) RemoveArticle(ownerId int, articleId int) (resp responses.BaseBool, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	params["article_id"] = articleId

	err = f.SendObjRequest("fave.removeArticle", params, &resp)

	return
}

// RemoveLink - Removes link from the user's faves.
// Parameters:
//   * linkId - Link ID (can be obtained by [vk.com/dev/faves.getLinks|faves.getLinks] method).
//   * link - Link URL
func (f *Fave) RemoveLink(linkId string, link string) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	if linkId != "" {
		params["link_id"] = linkId
	}

	if link != "" {
		params["link"] = link
	}

	err = f.SendObjRequest("fave.removeLink", params, &resp)

	return
}

// RemovePage - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * userId - NO DESCRIPTION IN JSON SCHEMA
//   * groupId - NO DESCRIPTION IN JSON SCHEMA
func (f *Fave) RemovePage(userId int, groupId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	if userId > 0 {
		params["user_id"] = userId
	}

	if groupId > 0 {
		params["group_id"] = groupId
	}

	err = f.SendObjRequest("fave.removePage", params, &resp)

	return
}

// RemovePost - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * ownerId - NO DESCRIPTION IN JSON SCHEMA
//   * id - NO DESCRIPTION IN JSON SCHEMA
func (f *Fave) RemovePost(ownerId int, id int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	params["id"] = id

	err = f.SendObjRequest("fave.removePost", params, &resp)

	return
}

// RemoveProduct - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * ownerId - NO DESCRIPTION IN JSON SCHEMA
//   * id - NO DESCRIPTION IN JSON SCHEMA
func (f *Fave) RemoveProduct(ownerId int, id int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	params["id"] = id

	err = f.SendObjRequest("fave.removeProduct", params, &resp)

	return
}

// RemoveTag - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * id - NO DESCRIPTION IN JSON SCHEMA
func (f *Fave) RemoveTag(id int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["id"] = id

	err = f.SendObjRequest("fave.removeTag", params, &resp)

	return
}

// ReorderTags - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * ids - NO DESCRIPTION IN JSON SCHEMA
func (f *Fave) ReorderTags(ids []int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["ids"] = SliceToString(ids)

	err = f.SendObjRequest("fave.reorderTags", params, &resp)

	return
}

// SetPageTags - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * userId - NO DESCRIPTION IN JSON SCHEMA
//   * groupId - NO DESCRIPTION IN JSON SCHEMA
//   * tagIds - NO DESCRIPTION IN JSON SCHEMA
func (f *Fave) SetPageTags(userId int, groupId int, tagIds []int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	if userId > 0 {
		params["user_id"] = userId
	}

	if groupId > 0 {
		params["group_id"] = groupId
	}

	if len(tagIds) > 0 {
		params["tag_ids"] = SliceToString(tagIds)
	}

	err = f.SendObjRequest("fave.setPageTags", params, &resp)

	return
}

// SetTags - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * itemType - NO DESCRIPTION IN JSON SCHEMA
//   * itemOwnerId - NO DESCRIPTION IN JSON SCHEMA
//   * itemId - NO DESCRIPTION IN JSON SCHEMA
//   * tagIds - NO DESCRIPTION IN JSON SCHEMA
//   * linkId - NO DESCRIPTION IN JSON SCHEMA
//   * linkUrl - NO DESCRIPTION IN JSON SCHEMA
func (f *Fave) SetTags(itemType string, itemOwnerId int, itemId int, tagIds []int, linkId string, linkUrl string) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	if itemType != "" {
		params["item_type"] = itemType
	}

	if itemOwnerId > 0 {
		params["item_owner_id"] = itemOwnerId
	}

	if itemId > 0 {
		params["item_id"] = itemId
	}

	if len(tagIds) > 0 {
		params["tag_ids"] = SliceToString(tagIds)
	}

	if linkId != "" {
		params["link_id"] = linkId
	}

	if linkUrl != "" {
		params["link_url"] = linkUrl
	}

	err = f.SendObjRequest("fave.setTags", params, &resp)

	return
}

// TrackPageInteraction - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * userId - NO DESCRIPTION IN JSON SCHEMA
//   * groupId - NO DESCRIPTION IN JSON SCHEMA
func (f *Fave) TrackPageInteraction(userId int, groupId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	if userId > 0 {
		params["user_id"] = userId
	}

	if groupId > 0 {
		params["group_id"] = groupId
	}

	err = f.SendObjRequest("fave.trackPageInteraction", params, &resp)

	return
}
