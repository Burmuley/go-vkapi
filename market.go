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

type Market struct {
	*VKApi
}

/////////////////////////////////////////////////////////////
// `Market` methods
/////////////////////////////////////////////////////////////

// Add - Ads a new item to the market.
// Parameters:
//   * ownerId - ID of an item owner community.
//   * name - Item name.
//   * description - Item description.
//   * categoryId - Item category ID.
//   * price - Item price.
//   * oldPrice - NO DESCRIPTION IN JSON SCHEMA
//   * deleted - Item status ('1' — deleted, '0' — not deleted).
//   * mainPhotoId - Cover photo ID.
//   * photoIds - IDs of additional photos.
//   * url - Url for button in market item.
func (m Market) Add(ownerId int, name string, description string, categoryId int, price float64, oldPrice float64, deleted bool, mainPhotoId int, photoIds []int, url string) (resp responses.MarketAdd, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	params["name"] = name

	params["description"] = description

	params["category_id"] = categoryId

	if price > 0 {
		params["price"] = price
	}

	if oldPrice > 0 {
		params["old_price"] = oldPrice
	}

	params["deleted"] = deleted

	params["main_photo_id"] = mainPhotoId

	if len(photoIds) > 0 {
		params["photo_ids"] = SliceToString(photoIds)
	}

	if url != "" {
		params["url"] = url
	}

	err = m.SendObjRequest("market.add", params, &resp)

	return
}

// AddAlbum - Creates new collection of items
// Parameters:
//   * ownerId - ID of an item owner community.
//   * title - Collection title.
//   * photoId - Cover photo ID.
//   * mainAlbum - Set as main ('1' – set, '0' – no).
func (m Market) AddAlbum(ownerId int, title string, photoId int, mainAlbum bool) (resp responses.MarketAddAlbum, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	params["title"] = title

	if photoId > 0 {
		params["photo_id"] = photoId
	}

	params["main_album"] = mainAlbum

	err = m.SendObjRequest("market.addAlbum", params, &resp)

	return
}

// AddToAlbum - Adds an item to one or multiple collections.
// Parameters:
//   * ownerId - ID of an item owner community.
//   * itemId - Item ID.
//   * albumIds - Collections IDs to add item to.
func (m Market) AddToAlbum(ownerId int, itemId int, albumIds []int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	params["item_id"] = itemId

	params["album_ids"] = SliceToString(albumIds)

	err = m.SendObjRequest("market.addToAlbum", params, &resp)

	return
}

// CreateComment - Creates a new comment for an item.
// Parameters:
//   * ownerId - ID of an item owner community.
//   * itemId - Item ID.
//   * message - Comment text (required if 'attachments' parameter is not specified)
//   * attachments - Comma-separated list of objects attached to a comment. The field is submitted the following way: , "'<owner_id>_<media_id>,<owner_id>_<media_id>'", , '' - media attachment type: "'photo' - photo, 'video' - video, 'audio' - audio, 'doc' - document", , '<owner_id>' - media owner id, '<media_id>' - media attachment id, , For example: "photo100172_166443618,photo66748_265827614",
//   * fromGroup - '1' - comment will be published on behalf of a community, '0' - on behalf of a user (by default).
//   * replyToComment - ID of a comment to reply with current comment to.
//   * stickerId - Sticker ID.
//   * guid - Random value to avoid resending one comment.
func (m Market) CreateComment(ownerId int, itemId int, message string, attachments []string, fromGroup bool, replyToComment int, stickerId int, guid string) (resp responses.MarketCreateComment, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	params["item_id"] = itemId

	if message != "" {
		params["message"] = message
	}

	if len(attachments) > 0 {
		params["attachments"] = SliceToString(attachments)
	}

	params["from_group"] = fromGroup

	if replyToComment > 0 {
		params["reply_to_comment"] = replyToComment
	}

	if stickerId > 0 {
		params["sticker_id"] = stickerId
	}

	if guid != "" {
		params["guid"] = guid
	}

	err = m.SendObjRequest("market.createComment", params, &resp)

	return
}

// Delete - Deletes an item.
// Parameters:
//   * ownerId - ID of an item owner community.
//   * itemId - Item ID.
func (m Market) Delete(ownerId int, itemId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	params["item_id"] = itemId

	err = m.SendObjRequest("market.delete", params, &resp)

	return
}

// DeleteAlbum - Deletes a collection of items.
// Parameters:
//   * ownerId - ID of an collection owner community.
//   * albumId - Collection ID.
func (m Market) DeleteAlbum(ownerId int, albumId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	params["album_id"] = albumId

	err = m.SendObjRequest("market.deleteAlbum", params, &resp)

	return
}

// DeleteComment - Deletes an item's comment
// Parameters:
//   * ownerId - identifier of an item owner community, "Note that community id in the 'owner_id' parameter should be negative number. For example 'owner_id'=-1 matches the [vk.com/apiclub|VK API] community "
//   * commentId - comment id
func (m Market) DeleteComment(ownerId int, commentId int) (resp responses.MarketDeleteComment, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	params["comment_id"] = commentId

	err = m.SendObjRequest("market.deleteComment", params, &resp)

	return
}

// Edit - Edits an item.
// Parameters:
//   * ownerId - ID of an item owner community.
//   * itemId - Item ID.
//   * name - Item name.
//   * description - Item description.
//   * categoryId - Item category ID.
//   * price - Item price.
//   * deleted - Item status ('1' — deleted, '0' — not deleted).
//   * mainPhotoId - Cover photo ID.
//   * photoIds - IDs of additional photos.
//   * url - Url for button in market item.
func (m Market) Edit(ownerId int, itemId int, name string, description string, categoryId int, price float64, deleted bool, mainPhotoId int, photoIds []int, url string) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	params["item_id"] = itemId

	params["name"] = name

	params["description"] = description

	params["category_id"] = categoryId

	params["price"] = price

	params["deleted"] = deleted

	params["main_photo_id"] = mainPhotoId

	if len(photoIds) > 0 {
		params["photo_ids"] = SliceToString(photoIds)
	}

	if url != "" {
		params["url"] = url
	}

	err = m.SendObjRequest("market.edit", params, &resp)

	return
}

// EditAlbum - Edits a collection of items
// Parameters:
//   * ownerId - ID of an collection owner community.
//   * albumId - Collection ID.
//   * title - Collection title.
//   * photoId - Cover photo id
//   * mainAlbum - Set as main ('1' – set, '0' – no).
func (m Market) EditAlbum(ownerId int, albumId int, title string, photoId int, mainAlbum bool) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	params["album_id"] = albumId

	params["title"] = title

	if photoId > 0 {
		params["photo_id"] = photoId
	}

	params["main_album"] = mainAlbum

	err = m.SendObjRequest("market.editAlbum", params, &resp)

	return
}

// EditComment - Chages item comment's text
// Parameters:
//   * ownerId - ID of an item owner community.
//   * commentId - Comment ID.
//   * message - New comment text (required if 'attachments' are not specified), , 2048 symbols maximum.
//   * attachments - Comma-separated list of objects attached to a comment. The field is submitted the following way: , "'<owner_id>_<media_id>,<owner_id>_<media_id>'", , '' - media attachment type: "'photo' - photo, 'video' - video, 'audio' - audio, 'doc' - document", , '<owner_id>' - media owner id, '<media_id>' - media attachment id, , For example: "photo100172_166443618,photo66748_265827614",
func (m Market) EditComment(ownerId int, commentId int, message string, attachments []string) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	params["comment_id"] = commentId

	if message != "" {
		params["message"] = message
	}

	if len(attachments) > 0 {
		params["attachments"] = SliceToString(attachments)
	}

	err = m.SendObjRequest("market.editComment", params, &resp)

	return
}

// Get - Returns items list for a community.
// Parameters:
//   * ownerId - ID of an item owner community, "Note that community id in the 'owner_id' parameter should be negative number. For example 'owner_id'=-1 matches the [vk.com/apiclub|VK API] community "
//   * albumId - NO DESCRIPTION IN JSON SCHEMA
//   * count - Number of items to return.
//   * offset - Offset needed to return a specific subset of results.
//   * extended - '1' – method will return additional fields: 'likes, can_comment, car_repost, photos'. These parameters are not returned by default.
func (m Market) Get(ownerId int, albumId int, count int, offset int) (resp responses.MarketGet, err error) {
	params := map[string]interface{}{}
	params["extended"] = "0"

	params["owner_id"] = ownerId

	if albumId > 0 {
		params["album_id"] = albumId
	}

	if count > 0 {
		params["count"] = count
	}

	if offset > 0 {
		params["offset"] = offset
	}

	err = m.SendObjRequest("market.get", params, &resp)

	return
}

// GetExtended - Returns items list for a community.
// Parameters:
//   * ownerId - ID of an item owner community, "Note that community id in the 'owner_id' parameter should be negative number. For example 'owner_id'=-1 matches the [vk.com/apiclub|VK API] community "
//   * albumId - NO DESCRIPTION IN JSON SCHEMA
//   * count - Number of items to return.
//   * offset - Offset needed to return a specific subset of results.
//   * extended - '1' – method will return additional fields: 'likes, can_comment, car_repost, photos'. These parameters are not returned by default.
func (m Market) GetExtended(ownerId int, albumId int, count int, offset int) (resp responses.MarketGetExtended, err error) {
	params := map[string]interface{}{}
	params["extended"] = "1"

	params["owner_id"] = ownerId

	if albumId > 0 {
		params["album_id"] = albumId
	}

	if count > 0 {
		params["count"] = count
	}

	if offset > 0 {
		params["offset"] = offset
	}

	err = m.SendObjRequest("market.get", params, &resp)

	return
}

// GetAlbumById - Returns items album's data
// Parameters:
//   * ownerId - identifier of an album owner community, "Note that community id in the 'owner_id' parameter should be negative number. For example 'owner_id'=-1 matches the [vk.com/apiclub|VK API] community "
//   * albumIds - collections identifiers to obtain data from
func (m Market) GetAlbumById(ownerId int, albumIds []int) (resp responses.MarketGetAlbumById, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	params["album_ids"] = SliceToString(albumIds)

	err = m.SendObjRequest("market.getAlbumById", params, &resp)

	return
}

// GetAlbums - Returns community's collections list.
// Parameters:
//   * ownerId - ID of an items owner community.
//   * offset - Offset needed to return a specific subset of results.
//   * count - Number of items to return.
func (m Market) GetAlbums(ownerId int, offset int, count int) (resp responses.MarketGetAlbums, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	err = m.SendObjRequest("market.getAlbums", params, &resp)

	return
}

// GetById - Returns information about market items by their ids.
// Parameters:
//   * itemIds - Comma-separated ids list: {user id}_{item id}. If an item belongs to a community -{community id} is used. " 'Videos' value example: , '-4363_136089719,13245770_137352259'"
//   * extended - '1' – to return additional fields: 'likes, can_comment, car_repost, photos'. By default: '0'.
func (m Market) GetById(itemIds []string) (resp responses.MarketGetById, err error) {
	params := map[string]interface{}{}
	params["extended"] = "0"

	params["item_ids"] = SliceToString(itemIds)

	err = m.SendObjRequest("market.getById", params, &resp)

	return
}

// GetByIdExtended - Returns information about market items by their ids.
// Parameters:
//   * itemIds - Comma-separated ids list: {user id}_{item id}. If an item belongs to a community -{community id} is used. " 'Videos' value example: , '-4363_136089719,13245770_137352259'"
//   * extended - '1' – to return additional fields: 'likes, can_comment, car_repost, photos'. By default: '0'.
func (m Market) GetByIdExtended(itemIds []string) (resp responses.MarketGetByIdExtended, err error) {
	params := map[string]interface{}{}
	params["extended"] = "1"

	params["item_ids"] = SliceToString(itemIds)

	err = m.SendObjRequest("market.getById", params, &resp)

	return
}

// GetCategories - Returns a list of market categories.
// Parameters:
//   * count - Number of results to return.
//   * offset - Offset needed to return a specific subset of results.
func (m Market) GetCategories(count int, offset int) (resp responses.MarketGetCategories, err error) {
	params := map[string]interface{}{}

	if count > 0 {
		params["count"] = count
	}

	if offset > 0 {
		params["offset"] = offset
	}

	err = m.SendObjRequest("market.getCategories", params, &resp)

	return
}

// GetComments - Returns comments list for an item.
// Parameters:
//   * ownerId - ID of an item owner community
//   * itemId - Item ID.
//   * needLikes - '1' — to return likes info.
//   * startCommentId - ID of a comment to start a list from (details below).
//   * offset - NO DESCRIPTION IN JSON SCHEMA
//   * count - Number of results to return.
//   * sort - Sort order ('asc' — from old to new, 'desc' — from new to old)
//   * extended - '1' — comments will be returned as numbered objects, in addition lists of 'profiles' and 'groups' objects will be returned.
//   * fields - List of additional profile fields to return. See the [vk.com/dev/fields|details]
func (m Market) GetComments(ownerId int, itemId int, needLikes bool, startCommentId int, offset int, count int, sort string, fields []objects.UsersFields) (resp responses.MarketGetComments, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	params["item_id"] = itemId

	params["need_likes"] = needLikes

	if startCommentId > 0 {
		params["start_comment_id"] = startCommentId
	}

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	if sort != "" {
		params["sort"] = sort
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	err = m.SendObjRequest("market.getComments", params, &resp)

	return
}

// RemoveFromAlbum - Removes an item from one or multiple collections.
// Parameters:
//   * ownerId - ID of an item owner community.
//   * itemId - Item ID.
//   * albumIds - Collections IDs to remove item from.
func (m Market) RemoveFromAlbum(ownerId int, itemId int, albumIds []int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	params["item_id"] = itemId

	params["album_ids"] = SliceToString(albumIds)

	err = m.SendObjRequest("market.removeFromAlbum", params, &resp)

	return
}

// ReorderAlbums - Reorders the collections list.
// Parameters:
//   * ownerId - ID of an item owner community.
//   * albumId - Collection ID.
//   * before - ID of a collection to place current collection before it.
//   * after - ID of a collection to place current collection after it.
func (m Market) ReorderAlbums(ownerId int, albumId int, before int, after int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	params["album_id"] = albumId

	if before > 0 {
		params["before"] = before
	}

	if after > 0 {
		params["after"] = after
	}

	err = m.SendObjRequest("market.reorderAlbums", params, &resp)

	return
}

// ReorderItems - Changes item place in a collection.
// Parameters:
//   * ownerId - ID of an item owner community.
//   * albumId - ID of a collection to reorder items in. Set 0 to reorder full items list.
//   * itemId - Item ID.
//   * before - ID of an item to place current item before it.
//   * after - ID of an item to place current item after it.
func (m Market) ReorderItems(ownerId int, albumId int, itemId int, before int, after int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	if albumId > 0 {
		params["album_id"] = albumId
	}

	params["item_id"] = itemId

	if before > 0 {
		params["before"] = before
	}

	if after > 0 {
		params["after"] = after
	}

	err = m.SendObjRequest("market.reorderItems", params, &resp)

	return
}

// Report - Sends a complaint to the item.
// Parameters:
//   * ownerId - ID of an item owner community.
//   * itemId - Item ID.
//   * reason - Complaint reason. Possible values: *'0' — spam,, *'1' — child porn,, *'2' — extremism,, *'3' — violence,, *'4' — drugs propaganda,, *'5' — adult materials,, *'6' — insult.
func (m Market) Report(ownerId int, itemId int, reason int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	params["item_id"] = itemId

	if reason > 0 {
		params["reason"] = reason
	}

	err = m.SendObjRequest("market.report", params, &resp)

	return
}

// ReportComment - Sends a complaint to the item's comment.
// Parameters:
//   * ownerId - ID of an item owner community.
//   * commentId - Comment ID.
//   * reason - Complaint reason. Possible values: *'0' — spam,, *'1' — child porn,, *'2' — extremism,, *'3' — violence,, *'4' — drugs propaganda,, *'5' — adult materials,, *'6' — insult.
func (m Market) ReportComment(ownerId int, commentId int, reason int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	params["comment_id"] = commentId

	params["reason"] = reason

	err = m.SendObjRequest("market.reportComment", params, &resp)

	return
}

// Restore - Restores recently deleted item
// Parameters:
//   * ownerId - ID of an item owner community.
//   * itemId - Deleted item ID.
func (m Market) Restore(ownerId int, itemId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	params["item_id"] = itemId

	err = m.SendObjRequest("market.restore", params, &resp)

	return
}

// RestoreComment - Restores a recently deleted comment
// Parameters:
//   * ownerId - identifier of an item owner community, "Note that community id in the 'owner_id' parameter should be negative number. For example 'owner_id'=-1 matches the [vk.com/apiclub|VK API] community "
//   * commentId - deleted comment id
func (m Market) RestoreComment(ownerId int, commentId int) (resp responses.MarketRestoreComment, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	params["comment_id"] = commentId

	err = m.SendObjRequest("market.restoreComment", params, &resp)

	return
}

// Search - Searches market items in a community's catalog
// Parameters:
//   * ownerId - ID of an items owner community.
//   * albumId - NO DESCRIPTION IN JSON SCHEMA
//   * q - Search query, for example "pink slippers".
//   * priceFrom - Minimum item price value.
//   * priceTo - Maximum item price value.
//   * tags - Comma-separated tag IDs list.
//   * sort - NO DESCRIPTION IN JSON SCHEMA
//   * rev - '0' — do not use reverse order, '1' — use reverse order
//   * offset - Offset needed to return a specific subset of results.
//   * count - Number of items to return.
//   * extended - '1' – to return additional fields: 'likes, can_comment, car_repost, photos'. By default: '0'.
//   * status - NO DESCRIPTION IN JSON SCHEMA
func (m Market) Search(ownerId int, albumId int, q string, priceFrom int, priceTo int, tags []int, sort int, rev int, offset int, count int, status int) (resp responses.MarketSearch, err error) {
	params := map[string]interface{}{}
	params["extended"] = "0"

	params["owner_id"] = ownerId

	if albumId > 0 {
		params["album_id"] = albumId
	}

	if q != "" {
		params["q"] = q
	}

	if priceFrom > 0 {
		params["price_from"] = priceFrom
	}

	if priceTo > 0 {
		params["price_to"] = priceTo
	}

	if len(tags) > 0 {
		params["tags"] = SliceToString(tags)
	}

	if sort > 0 {
		params["sort"] = sort
	}

	if rev > 0 {
		params["rev"] = rev
	}

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	if status > 0 {
		params["status"] = status
	}

	err = m.SendObjRequest("market.search", params, &resp)

	return
}

// SearchExtended - Searches market items in a community's catalog
// Parameters:
//   * ownerId - ID of an items owner community.
//   * albumId - NO DESCRIPTION IN JSON SCHEMA
//   * q - Search query, for example "pink slippers".
//   * priceFrom - Minimum item price value.
//   * priceTo - Maximum item price value.
//   * tags - Comma-separated tag IDs list.
//   * sort - NO DESCRIPTION IN JSON SCHEMA
//   * rev - '0' — do not use reverse order, '1' — use reverse order
//   * offset - Offset needed to return a specific subset of results.
//   * count - Number of items to return.
//   * extended - '1' – to return additional fields: 'likes, can_comment, car_repost, photos'. By default: '0'.
//   * status - NO DESCRIPTION IN JSON SCHEMA
func (m Market) SearchExtended(ownerId int, albumId int, q string, priceFrom int, priceTo int, tags []int, sort int, rev int, offset int, count int, status int) (resp responses.MarketSearchExtended, err error) {
	params := map[string]interface{}{}
	params["extended"] = "1"

	params["owner_id"] = ownerId

	if albumId > 0 {
		params["album_id"] = albumId
	}

	if q != "" {
		params["q"] = q
	}

	if priceFrom > 0 {
		params["price_from"] = priceFrom
	}

	if priceTo > 0 {
		params["price_to"] = priceTo
	}

	if len(tags) > 0 {
		params["tags"] = SliceToString(tags)
	}

	if sort > 0 {
		params["sort"] = sort
	}

	if rev > 0 {
		params["rev"] = rev
	}

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	if status > 0 {
		params["status"] = status
	}

	err = m.SendObjRequest("market.search", params, &resp)

	return
}
