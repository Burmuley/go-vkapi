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
	"encoding/json"
	"github.com/Burmuley/go-vkapi/objects"
	"github.com/Burmuley/go-vkapi/responses"
)

type Photos struct {
	*VKApi
}

/////////////////////////////////////////////////////////////
// `Photos` methods
/////////////////////////////////////////////////////////////

// Confirmtag - Confirms a tag on a photo.
// Parameters:
//   * ownerId - ID of the user or community that owns the photo.
//   * photoId - Photo ID.
//   * tagId - Tag ID.
func (p *Photos) Confirmtag(ownerId int, photoId string, tagId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	params["photo_id"] = photoId

	params["tag_id"] = tagId

	err = p.SendObjRequest("photos.confirmTag", params, &resp)

	return
}

// Copy - Allows to copy a photo to the "Saved photos" album
// Parameters:
//   * ownerId - photo's owner ID
//   * photoId - photo ID
//   * accessKey - for private photos
func (p *Photos) Copy(ownerId int, photoId int, accessKey string) (resp responses.PhotosCopy, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	params["photo_id"] = photoId

	if accessKey != "" {
		params["access_key"] = accessKey
	}

	err = p.SendObjRequest("photos.copy", params, &resp)

	return
}

// Createalbum - Creates an empty photo album.
// Parameters:
//   * title - Album title.
//   * groupId - ID of the community in which the album will be created.
//   * description - Album description.
//   * privacyView - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * privacyComment - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * uploadByAdminsOnly - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * commentsDisabled - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (p *Photos) Createalbum(title string, groupId int, description string, privacyView []string, privacyComment []string, uploadByAdminsOnly bool, commentsDisabled bool) (resp responses.PhotosCreatealbum, err error) {
	params := map[string]interface{}{}

	params["title"] = title

	if groupId > 0 {
		params["group_id"] = groupId
	}

	if description != "" {
		params["description"] = description
	}

	if len(privacyView) > 0 {
		params["privacy_view"] = SliceToString(privacyView)
	}

	if len(privacyComment) > 0 {
		params["privacy_comment"] = SliceToString(privacyComment)
	}

	params["upload_by_admins_only"] = uploadByAdminsOnly

	params["comments_disabled"] = commentsDisabled

	err = p.SendObjRequest("photos.createAlbum", params, &resp)

	return
}

// Createcomment - Adds a new comment on the photo.
// Parameters:
//   * ownerId - ID of the user or community that owns the photo.
//   * photoId - Photo ID.
//   * message - Comment text.
//   * attachments - (Required if 'message' is not set.) List of objects attached to the post, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, '<owner_id>' — Media attachment owner ID. '<media_id>' — Media attachment ID. Example: "photo100172_166443618,photo66748_265827614"
//   * fromGroup - '1' — to post a comment from the community
//   * replyToComment - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * stickerId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * accessKey - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * guid - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (p *Photos) Createcomment(ownerId int, photoId int, message string, attachments []string, fromGroup bool, replyToComment int, stickerId int, accessKey string, guid string) (resp responses.PhotosCreatecomment, err error) {
	params := map[string]interface{}{}

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	params["photo_id"] = photoId

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

	if accessKey != "" {
		params["access_key"] = accessKey
	}

	if guid != "" {
		params["guid"] = guid
	}

	err = p.SendObjRequest("photos.createComment", params, &resp)

	return
}

// Delete - Deletes a photo.
// Parameters:
//   * ownerId - ID of the user or community that owns the photo.
//   * photoId - Photo ID.
func (p *Photos) Delete(ownerId int, photoId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	params["photo_id"] = photoId

	err = p.SendObjRequest("photos.delete", params, &resp)

	return
}

// Deletealbum - Deletes a photo album belonging to the current user.
// Parameters:
//   * albumId - Album ID.
//   * groupId - ID of the community that owns the album.
func (p *Photos) Deletealbum(albumId int, groupId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["album_id"] = albumId

	if groupId > 0 {
		params["group_id"] = groupId
	}

	err = p.SendObjRequest("photos.deleteAlbum", params, &resp)

	return
}

// Deletecomment - Deletes a comment on the photo.
// Parameters:
//   * ownerId - ID of the user or community that owns the photo.
//   * commentId - Comment ID.
func (p *Photos) Deletecomment(ownerId int, commentId int) (resp responses.PhotosDeletecomment, err error) {
	params := map[string]interface{}{}

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	params["comment_id"] = commentId

	err = p.SendObjRequest("photos.deleteComment", params, &resp)

	return
}

// Edit - Edits the caption of a photo.
// Parameters:
//   * ownerId - ID of the user or community that owns the photo.
//   * photoId - Photo ID.
//   * caption - New caption for the photo. If this parameter is not set, it is considered to be equal to an empty string.
//   * latitude - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * longitude - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * placeStr - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * foursquareId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * deletePlace - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (p *Photos) Edit(ownerId int, photoId int, caption string, latitude json.Number, longitude json.Number, placeStr string, foursquareId string, deletePlace bool) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	params["photo_id"] = photoId

	if caption != "" {
		params["caption"] = caption
	}

	if v, err := latitude.Int64(); err == nil && v > 0 {
		params["latitude"] = v
	} else if v := latitude.String(); v != "" {
		params["latitude"] = v
	}

	if v, err := longitude.Int64(); err == nil && v > 0 {
		params["longitude"] = v
	} else if v := longitude.String(); v != "" {
		params["longitude"] = v
	}

	if placeStr != "" {
		params["place_str"] = placeStr
	}

	if foursquareId != "" {
		params["foursquare_id"] = foursquareId
	}

	params["delete_place"] = deletePlace

	err = p.SendObjRequest("photos.edit", params, &resp)

	return
}

// Editalbum - Edits information about a photo album.
// Parameters:
//   * albumId - ID of the photo album to be edited.
//   * title - New album title.
//   * description - New album description.
//   * ownerId - ID of the user or community that owns the album.
//   * privacyView - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * privacyComment - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * uploadByAdminsOnly - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * commentsDisabled - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (p *Photos) Editalbum(albumId int, title string, description string, ownerId int, privacyView []string, privacyComment []string, uploadByAdminsOnly bool, commentsDisabled bool) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["album_id"] = albumId

	if title != "" {
		params["title"] = title
	}

	if description != "" {
		params["description"] = description
	}

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	if len(privacyView) > 0 {
		params["privacy_view"] = SliceToString(privacyView)
	}

	if len(privacyComment) > 0 {
		params["privacy_comment"] = SliceToString(privacyComment)
	}

	params["upload_by_admins_only"] = uploadByAdminsOnly

	params["comments_disabled"] = commentsDisabled

	err = p.SendObjRequest("photos.editAlbum", params, &resp)

	return
}

// Editcomment - Edits a comment on a photo.
// Parameters:
//   * ownerId - ID of the user or community that owns the photo.
//   * commentId - Comment ID.
//   * message - New text of the comment.
//   * attachments - (Required if 'message' is not set.) List of objects attached to the post, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, '<owner_id>' — Media attachment owner ID. '<media_id>' — Media attachment ID. Example: "photo100172_166443618,photo66748_265827614"
func (p *Photos) Editcomment(ownerId int, commentId int, message string, attachments []string) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	params["comment_id"] = commentId

	if message != "" {
		params["message"] = message
	}

	if len(attachments) > 0 {
		params["attachments"] = SliceToString(attachments)
	}

	err = p.SendObjRequest("photos.editComment", params, &resp)

	return
}

// Get - Returns a list of a user's or community's photos.
// Parameters:
//   * ownerId - ID of the user or community that owns the photos. Use a negative value to designate a community ID.
//   * albumId - Photo album ID. To return information about photos from service albums, use the following string values: 'profile, wall, saved'.
//   * photoIds - Photo IDs.
//   * rev - Sort order: '1' — reverse chronological, '0' — chronological
//   * extended - '1' — to return additional 'likes', 'comments', and 'tags' fields, '0' — (default)
//   * feedType - Type of feed obtained in 'feed' field of the method.
//   * feed - unixtime, that can be obtained with [vk.com/dev/newsfeed.get|newsfeed.get] method in date field to get all photos uploaded by the user on a specific day, or photos the user has been tagged on. Also, 'uid' parameter of the user the event happened with shall be specified.
//   * photoSizes - '1' — to return photo sizes in a [vk.com/dev/photo_sizes|special format]
//   * offset - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * count - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (p *Photos) Get(ownerId int, albumId string, photoIds []string, rev bool, feedType string, feed int, photoSizes bool, offset int, count int) (resp responses.PhotosGet, err error) {
	params := map[string]interface{}{}
	params["extended"] = "0"

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	if albumId != "" {
		params["album_id"] = albumId
	}

	if len(photoIds) > 0 {
		params["photo_ids"] = SliceToString(photoIds)
	}

	params["rev"] = rev

	if feedType != "" {
		params["feed_type"] = feedType
	}

	if feed > 0 {
		params["feed"] = feed
	}

	params["photo_sizes"] = photoSizes

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	err = p.SendObjRequest("photos.get", params, &resp)

	return
}

// GetExtended - Returns a list of a user's or community's photos.
// Parameters:
//   * ownerId - ID of the user or community that owns the photos. Use a negative value to designate a community ID.
//   * albumId - Photo album ID. To return information about photos from service albums, use the following string values: 'profile, wall, saved'.
//   * photoIds - Photo IDs.
//   * rev - Sort order: '1' — reverse chronological, '0' — chronological
//   * extended - '1' — to return additional 'likes', 'comments', and 'tags' fields, '0' — (default)
//   * feedType - Type of feed obtained in 'feed' field of the method.
//   * feed - unixtime, that can be obtained with [vk.com/dev/newsfeed.get|newsfeed.get] method in date field to get all photos uploaded by the user on a specific day, or photos the user has been tagged on. Also, 'uid' parameter of the user the event happened with shall be specified.
//   * photoSizes - '1' — to return photo sizes in a [vk.com/dev/photo_sizes|special format]
//   * offset - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * count - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (p *Photos) GetExtended(ownerId int, albumId string, photoIds []string, rev bool, feedType string, feed int, photoSizes bool, offset int, count int) (resp responses.PhotosGetExtended, err error) {
	params := map[string]interface{}{}
	params["extended"] = "1"

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	if albumId != "" {
		params["album_id"] = albumId
	}

	if len(photoIds) > 0 {
		params["photo_ids"] = SliceToString(photoIds)
	}

	params["rev"] = rev

	if feedType != "" {
		params["feed_type"] = feedType
	}

	if feed > 0 {
		params["feed"] = feed
	}

	params["photo_sizes"] = photoSizes

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	err = p.SendObjRequest("photos.get", params, &resp)

	return
}

// Getalbums - Returns a list of a user's or community's photo albums.
// Parameters:
//   * ownerId - ID of the user or community that owns the albums.
//   * albumIds - Album IDs.
//   * offset - Offset needed to return a specific subset of albums.
//   * count - Number of albums to return.
//   * needSystem - '1' — to return system albums with negative IDs
//   * needCovers - '1' — to return an additional 'thumb_src' field, '0' — (default)
//   * photoSizes - '1' — to return photo sizes in a
func (p *Photos) Getalbums(ownerId int, albumIds []int, offset int, count int, needSystem bool, needCovers bool, photoSizes bool) (resp responses.PhotosGetalbums, err error) {
	params := map[string]interface{}{}

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	if len(albumIds) > 0 {
		params["album_ids"] = SliceToString(albumIds)
	}

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	params["need_system"] = needSystem

	params["need_covers"] = needCovers

	params["photo_sizes"] = photoSizes

	err = p.SendObjRequest("photos.getAlbums", params, &resp)

	return
}

// Getalbumscount - Returns the number of photo albums belonging to a user or community.
// Parameters:
//   * userId - User ID.
//   * groupId - Community ID.
func (p *Photos) Getalbumscount(userId int, groupId int) (resp responses.PhotosGetalbumscount, err error) {
	params := map[string]interface{}{}

	if userId > 0 {
		params["user_id"] = userId
	}

	if groupId > 0 {
		params["group_id"] = groupId
	}

	err = p.SendObjRequest("photos.getAlbumsCount", params, &resp)

	return
}

// Getall - Returns a list of photos belonging to a user or community, in reverse chronological order.
// Parameters:
//   * ownerId - ID of a user or community that owns the photos. Use a negative value to designate a community ID.
//   * extended - '1' — to return detailed information about photos
//   * offset - Offset needed to return a specific subset of photos. By default, '0'.
//   * count - Number of photos to return.
//   * photoSizes - '1' – to return image sizes in [vk.com/dev/photo_sizes|special format].
//   * noServiceAlbums - '1' – to return photos only from standard albums, '0' – to return all photos including those in service albums, e.g., 'My wall photos' (default)
//   * needHidden - '1' – to show information about photos being hidden from the block above the wall.
//   * skipHidden - '1' – not to return photos being hidden from the block above the wall. Works only with owner_id>0, no_service_albums is ignored.
func (p *Photos) Getall(ownerId int, offset int, count int, photoSizes bool, noServiceAlbums bool, needHidden bool, skipHidden bool) (resp responses.PhotosGetall, err error) {
	params := map[string]interface{}{}
	params["extended"] = "0"

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	params["photo_sizes"] = photoSizes

	params["no_service_albums"] = noServiceAlbums

	params["need_hidden"] = needHidden

	params["skip_hidden"] = skipHidden

	err = p.SendObjRequest("photos.getAll", params, &resp)

	return
}

// GetallExtended - Returns a list of photos belonging to a user or community, in reverse chronological order.
// Parameters:
//   * ownerId - ID of a user or community that owns the photos. Use a negative value to designate a community ID.
//   * extended - '1' — to return detailed information about photos
//   * offset - Offset needed to return a specific subset of photos. By default, '0'.
//   * count - Number of photos to return.
//   * photoSizes - '1' – to return image sizes in [vk.com/dev/photo_sizes|special format].
//   * noServiceAlbums - '1' – to return photos only from standard albums, '0' – to return all photos including those in service albums, e.g., 'My wall photos' (default)
//   * needHidden - '1' – to show information about photos being hidden from the block above the wall.
//   * skipHidden - '1' – not to return photos being hidden from the block above the wall. Works only with owner_id>0, no_service_albums is ignored.
func (p *Photos) GetallExtended(ownerId int, offset int, count int, photoSizes bool, noServiceAlbums bool, needHidden bool, skipHidden bool) (resp responses.PhotosGetallExtended, err error) {
	params := map[string]interface{}{}
	params["extended"] = "1"

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	params["photo_sizes"] = photoSizes

	params["no_service_albums"] = noServiceAlbums

	params["need_hidden"] = needHidden

	params["skip_hidden"] = skipHidden

	err = p.SendObjRequest("photos.getAll", params, &resp)

	return
}

// Getallcomments - Returns a list of comments on a specific photo album or all albums of the user sorted in reverse chronological order.
// Parameters:
//   * ownerId - ID of the user or community that owns the album(s).
//   * albumId - Album ID. If the parameter is not set, comments on all of the user's albums will be returned.
//   * needLikes - '1' — to return an additional 'likes' field, '0' — (default)
//   * offset - Offset needed to return a specific subset of comments. By default, '0'.
//   * count - Number of comments to return. By default, '20'. Maximum value, '100'.
func (p *Photos) Getallcomments(ownerId int, albumId int, needLikes bool, offset int, count int) (resp responses.PhotosGetallcomments, err error) {
	params := map[string]interface{}{}

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	if albumId > 0 {
		params["album_id"] = albumId
	}

	params["need_likes"] = needLikes

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	err = p.SendObjRequest("photos.getAllComments", params, &resp)

	return
}

// Getbyid - Returns information about photos by their IDs.
// Parameters:
//   * photos - IDs separated with a comma, that are IDs of users who posted photos and IDs of photos themselves with an underscore character between such IDs. To get information about a photo in the group album, you shall specify group ID instead of user ID. Example: "1_129207899,6492_135055734, , -20629724_271945303"
//   * extended - '1' — to return additional fields, '0' — (default)
//   * photoSizes - '1' — to return photo sizes in a
func (p *Photos) Getbyid(photos []string, photoSizes bool) (resp responses.PhotosGetbyid, err error) {
	params := map[string]interface{}{}
	params["extended"] = "0"

	params["photos"] = SliceToString(photos)

	params["photo_sizes"] = photoSizes

	err = p.SendObjRequest("photos.getById", params, &resp)

	return
}

// GetbyidExtended - Returns information about photos by their IDs.
// Parameters:
//   * photos - IDs separated with a comma, that are IDs of users who posted photos and IDs of photos themselves with an underscore character between such IDs. To get information about a photo in the group album, you shall specify group ID instead of user ID. Example: "1_129207899,6492_135055734, , -20629724_271945303"
//   * extended - '1' — to return additional fields, '0' — (default)
//   * photoSizes - '1' — to return photo sizes in a
func (p *Photos) GetbyidExtended(photos []string, photoSizes bool) (resp responses.PhotosGetbyidExtended, err error) {
	params := map[string]interface{}{}
	params["extended"] = "1"

	params["photos"] = SliceToString(photos)

	params["photo_sizes"] = photoSizes

	err = p.SendObjRequest("photos.getById", params, &resp)

	return
}

// Getchatuploadserver - Returns an upload link for chat cover pictures.
// Parameters:
//   * chatId - ID of the chat for which you want to upload a cover photo.
//   * cropX - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * cropY - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * cropWidth - Width (in pixels) of the photo after cropping.
func (p *Photos) Getchatuploadserver(chatId int, cropX int, cropY int, cropWidth int) (resp responses.BaseGetuploadserver, err error) {
	params := map[string]interface{}{}

	params["chat_id"] = chatId

	if cropX > 0 {
		params["crop_x"] = cropX
	}

	if cropY > 0 {
		params["crop_y"] = cropY
	}

	if cropWidth > 0 {
		params["crop_width"] = cropWidth
	}

	err = p.SendObjRequest("photos.getChatUploadServer", params, &resp)

	return
}

// Getcomments - Returns a list of comments on a photo.
// Parameters:
//   * ownerId - ID of the user or community that owns the photo.
//   * photoId - Photo ID.
//   * needLikes - '1' — to return an additional 'likes' field, '0' — (default)
//   * startCommentId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * offset - Offset needed to return a specific subset of comments. By default, '0'.
//   * count - Number of comments to return.
//   * sort - Sort order: 'asc' — old first, 'desc' — new first
//   * accessKey - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * extended - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * fields - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (p *Photos) Getcomments(ownerId int, photoId int, needLikes bool, startCommentId int, offset int, count int, sort string, accessKey string, fields []objects.UsersFields) (resp responses.PhotosGetcomments, err error) {
	params := map[string]interface{}{}
	params["extended"] = "0"

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	params["photo_id"] = photoId

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

	if accessKey != "" {
		params["access_key"] = accessKey
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	err = p.SendObjRequest("photos.getComments", params, &resp)

	return
}

// GetcommentsExtended - Returns a list of comments on a photo.
// Parameters:
//   * ownerId - ID of the user or community that owns the photo.
//   * photoId - Photo ID.
//   * needLikes - '1' — to return an additional 'likes' field, '0' — (default)
//   * startCommentId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * offset - Offset needed to return a specific subset of comments. By default, '0'.
//   * count - Number of comments to return.
//   * sort - Sort order: 'asc' — old first, 'desc' — new first
//   * accessKey - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * extended - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * fields - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (p *Photos) GetcommentsExtended(ownerId int, photoId int, needLikes bool, startCommentId int, offset int, count int, sort string, accessKey string, fields []objects.UsersFields) (resp responses.PhotosGetcommentsExtended, err error) {
	params := map[string]interface{}{}
	params["extended"] = "1"

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	params["photo_id"] = photoId

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

	if accessKey != "" {
		params["access_key"] = accessKey
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	err = p.SendObjRequest("photos.getComments", params, &resp)

	return
}

// Getmarketalbumuploadserver - Returns the server address for market album photo upload.
// Parameters:
//   * groupId - Community ID.
func (p *Photos) Getmarketalbumuploadserver(groupId int) (resp responses.BaseGetuploadserver, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	err = p.SendObjRequest("photos.getMarketAlbumUploadServer", params, &resp)

	return
}

// Getmarketuploadserver - Returns the server address for market photo upload.
// Parameters:
//   * groupId - Community ID.
//   * mainPhoto - '1' if you want to upload the main item photo.
//   * cropX - X coordinate of the crop left upper corner.
//   * cropY - Y coordinate of the crop left upper corner.
//   * cropWidth - Width of the cropped photo in px.
func (p *Photos) Getmarketuploadserver(groupId int, mainPhoto bool, cropX int, cropY int, cropWidth int) (resp responses.PhotosGetmarketuploadserver, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	params["main_photo"] = mainPhoto

	if cropX > 0 {
		params["crop_x"] = cropX
	}

	if cropY > 0 {
		params["crop_y"] = cropY
	}

	if cropWidth > 0 {
		params["crop_width"] = cropWidth
	}

	err = p.SendObjRequest("photos.getMarketUploadServer", params, &resp)

	return
}

// Getmessagesuploadserver - Returns the server address for photo upload in a private message for a user.
// Parameters:
//   * peerId - Destination ID. "For user: 'User ID', e.g. '12345'. For chat: '2000000000' + 'Chat ID', e.g. '2000000001'. For community: '- Community ID', e.g. '-12345'. "
func (p *Photos) Getmessagesuploadserver(peerId int) (resp responses.PhotosGetmessagesuploadserver, err error) {
	params := map[string]interface{}{}

	if peerId > 0 {
		params["peer_id"] = peerId
	}

	err = p.SendObjRequest("photos.getMessagesUploadServer", params, &resp)

	return
}

// Getnewtags - Returns a list of photos with tags that have not been viewed.
// Parameters:
//   * offset - Offset needed to return a specific subset of photos.
//   * count - Number of photos to return.
func (p *Photos) Getnewtags(offset int, count int) (resp responses.PhotosGetnewtags, err error) {
	params := map[string]interface{}{}

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	err = p.SendObjRequest("photos.getNewTags", params, &resp)

	return
}

// Getownercoverphotouploadserver - Returns the server address for owner cover upload.
// Parameters:
//   * groupId - ID of community that owns the album (if the photo will be uploaded to a community album).
//   * cropX - X coordinate of the left-upper corner
//   * cropY - Y coordinate of the left-upper corner
//   * cropX2 - X coordinate of the right-bottom corner
//   * cropY2 - Y coordinate of the right-bottom corner
func (p *Photos) Getownercoverphotouploadserver(groupId int, cropX int, cropY int, cropX2 int, cropY2 int) (resp responses.BaseGetuploadserver, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	if cropX > 0 {
		params["crop_x"] = cropX
	}

	if cropY > 0 {
		params["crop_y"] = cropY
	}

	if cropX2 > 0 {
		params["crop_x2"] = cropX2
	}

	if cropY2 > 0 {
		params["crop_y2"] = cropY2
	}

	err = p.SendObjRequest("photos.getOwnerCoverPhotoUploadServer", params, &resp)

	return
}

// Getownerphotouploadserver - Returns an upload server address for a profile or community photo.
// Parameters:
//   * ownerId - identifier of a community or current user. "Note that community id must be negative. 'owner_id=1' – user, 'owner_id=-1' – community, "
func (p *Photos) Getownerphotouploadserver(ownerId int) (resp responses.BaseGetuploadserver, err error) {
	params := map[string]interface{}{}

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	err = p.SendObjRequest("photos.getOwnerPhotoUploadServer", params, &resp)

	return
}

// Gettags - Returns a list of tags on a photo.
// Parameters:
//   * ownerId - ID of the user or community that owns the photo.
//   * photoId - Photo ID.
//   * accessKey - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (p *Photos) Gettags(ownerId int, photoId int, accessKey string) (resp responses.PhotosGettags, err error) {
	params := map[string]interface{}{}

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	params["photo_id"] = photoId

	if accessKey != "" {
		params["access_key"] = accessKey
	}

	err = p.SendObjRequest("photos.getTags", params, &resp)

	return
}

// Getuploadserver - Returns the server address for photo upload.
// Parameters:
//   * groupId - ID of community that owns the album (if the photo will be uploaded to a community album).
//   * albumId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (p *Photos) Getuploadserver(groupId int, albumId int) (resp responses.PhotosGetuploadserver, err error) {
	params := map[string]interface{}{}

	if groupId > 0 {
		params["group_id"] = groupId
	}

	if albumId > 0 {
		params["album_id"] = albumId
	}

	err = p.SendObjRequest("photos.getUploadServer", params, &resp)

	return
}

// Getuserphotos - Returns a list of photos in which a user is tagged.
// Parameters:
//   * userId - User ID.
//   * offset - Offset needed to return a specific subset of photos. By default, '0'.
//   * count - Number of photos to return. Maximum value is 1000.
//   * extended - '1' — to return an additional 'likes' field, '0' — (default)
//   * sort - Sort order: '1' — by date the tag was added in ascending order, '0' — by date the tag was added in descending order
func (p *Photos) Getuserphotos(userId int, offset int, count int, sort string) (resp responses.PhotosGetuserphotos, err error) {
	params := map[string]interface{}{}
	params["extended"] = "0"

	if userId > 0 {
		params["user_id"] = userId
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

	err = p.SendObjRequest("photos.getUserPhotos", params, &resp)

	return
}

// GetuserphotosExtended - Returns a list of photos in which a user is tagged.
// Parameters:
//   * userId - User ID.
//   * offset - Offset needed to return a specific subset of photos. By default, '0'.
//   * count - Number of photos to return. Maximum value is 1000.
//   * extended - '1' — to return an additional 'likes' field, '0' — (default)
//   * sort - Sort order: '1' — by date the tag was added in ascending order, '0' — by date the tag was added in descending order
func (p *Photos) GetuserphotosExtended(userId int, offset int, count int, sort string) (resp responses.PhotosGetuserphotosExtended, err error) {
	params := map[string]interface{}{}
	params["extended"] = "1"

	if userId > 0 {
		params["user_id"] = userId
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

	err = p.SendObjRequest("photos.getUserPhotos", params, &resp)

	return
}

// Getwalluploadserver - Returns the server address for photo upload onto a user's wall.
// Parameters:
//   * groupId - ID of community to whose wall the photo will be uploaded.
func (p *Photos) Getwalluploadserver(groupId int) (resp responses.PhotosGetwalluploadserver, err error) {
	params := map[string]interface{}{}

	if groupId > 0 {
		params["group_id"] = groupId
	}

	err = p.SendObjRequest("photos.getWallUploadServer", params, &resp)

	return
}

// Makecover - Makes a photo into an album cover.
// Parameters:
//   * ownerId - ID of the user or community that owns the photo.
//   * photoId - Photo ID.
//   * albumId - Album ID.
func (p *Photos) Makecover(ownerId int, photoId int, albumId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	params["photo_id"] = photoId

	if albumId > 0 {
		params["album_id"] = albumId
	}

	err = p.SendObjRequest("photos.makeCover", params, &resp)

	return
}

// Move - Moves a photo from one album to another.
// Parameters:
//   * ownerId - ID of the user or community that owns the photo.
//   * targetAlbumId - ID of the album to which the photo will be moved.
//   * photoId - Photo ID.
func (p *Photos) Move(ownerId int, targetAlbumId int, photoId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	params["target_album_id"] = targetAlbumId

	params["photo_id"] = photoId

	err = p.SendObjRequest("photos.move", params, &resp)

	return
}

// Puttag - Adds a tag on the photo.
// Parameters:
//   * ownerId - ID of the user or community that owns the photo.
//   * photoId - Photo ID.
//   * userId - ID of the user to be tagged.
//   * x - Upper left-corner coordinate of the tagged area (as a percentage of the photo's width).
//   * y - Upper left-corner coordinate of the tagged area (as a percentage of the photo's height).
//   * x2 - Lower right-corner coordinate of the tagged area (as a percentage of the photo's width).
//   * y2 - Lower right-corner coordinate of the tagged area (as a percentage of the photo's height).
func (p *Photos) Puttag(ownerId int, photoId int, userId int, x json.Number, y json.Number, x2 json.Number, y2 json.Number) (resp responses.PhotosPuttag, err error) {
	params := map[string]interface{}{}

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	params["photo_id"] = photoId

	params["user_id"] = userId

	if v, err := x.Int64(); err == nil && v > 0 {
		params["x"] = v
	} else if v := x.String(); v != "" {
		params["x"] = v
	}

	if v, err := y.Int64(); err == nil && v > 0 {
		params["y"] = v
	} else if v := y.String(); v != "" {
		params["y"] = v
	}

	if v, err := x2.Int64(); err == nil && v > 0 {
		params["x2"] = v
	} else if v := x2.String(); v != "" {
		params["x2"] = v
	}

	if v, err := y2.Int64(); err == nil && v > 0 {
		params["y2"] = v
	} else if v := y2.String(); v != "" {
		params["y2"] = v
	}

	err = p.SendObjRequest("photos.putTag", params, &resp)

	return
}

// Removetag - Removes a tag from a photo.
// Parameters:
//   * ownerId - ID of the user or community that owns the photo.
//   * photoId - Photo ID.
//   * tagId - Tag ID.
func (p *Photos) Removetag(ownerId int, photoId int, tagId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	params["photo_id"] = photoId

	params["tag_id"] = tagId

	err = p.SendObjRequest("photos.removeTag", params, &resp)

	return
}

// Reorderalbums - Reorders the album in the list of user albums.
// Parameters:
//   * ownerId - ID of the user or community that owns the album.
//   * albumId - Album ID.
//   * before - ID of the album before which the album in question shall be placed.
//   * after - ID of the album after which the album in question shall be placed.
func (p *Photos) Reorderalbums(ownerId int, albumId int, before int, after int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	params["album_id"] = albumId

	if before > 0 {
		params["before"] = before
	}

	if after > 0 {
		params["after"] = after
	}

	err = p.SendObjRequest("photos.reorderAlbums", params, &resp)

	return
}

// Reorderphotos - Reorders the photo in the list of photos of the user album.
// Parameters:
//   * ownerId - ID of the user or community that owns the photo.
//   * photoId - Photo ID.
//   * before - ID of the photo before which the photo in question shall be placed.
//   * after - ID of the photo after which the photo in question shall be placed.
func (p *Photos) Reorderphotos(ownerId int, photoId int, before int, after int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	params["photo_id"] = photoId

	if before > 0 {
		params["before"] = before
	}

	if after > 0 {
		params["after"] = after
	}

	err = p.SendObjRequest("photos.reorderPhotos", params, &resp)

	return
}

// Report - Reports (submits a complaint about) a photo.
// Parameters:
//   * ownerId - ID of the user or community that owns the photo.
//   * photoId - Photo ID.
//   * reason - Reason for the complaint: '0' – spam, '1' – child pornography, '2' – extremism, '3' – violence, '4' – drug propaganda, '5' – adult material, '6' – insult, abuse
func (p *Photos) Report(ownerId int, photoId int, reason int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	params["photo_id"] = photoId

	if reason > 0 {
		params["reason"] = reason
	}

	err = p.SendObjRequest("photos.report", params, &resp)

	return
}

// Reportcomment - Reports (submits a complaint about) a comment on a photo.
// Parameters:
//   * ownerId - ID of the user or community that owns the photo.
//   * commentId - ID of the comment being reported.
//   * reason - Reason for the complaint: '0' – spam, '1' – child pornography, '2' – extremism, '3' – violence, '4' – drug propaganda, '5' – adult material, '6' – insult, abuse
func (p *Photos) Reportcomment(ownerId int, commentId int, reason int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	params["comment_id"] = commentId

	if reason > 0 {
		params["reason"] = reason
	}

	err = p.SendObjRequest("photos.reportComment", params, &resp)

	return
}

// Restore - Restores a deleted photo.
// Parameters:
//   * ownerId - ID of the user or community that owns the photo.
//   * photoId - Photo ID.
func (p *Photos) Restore(ownerId int, photoId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	params["photo_id"] = photoId

	err = p.SendObjRequest("photos.restore", params, &resp)

	return
}

// Restorecomment - Restores a deleted comment on a photo.
// Parameters:
//   * ownerId - ID of the user or community that owns the photo.
//   * commentId - ID of the deleted comment.
func (p *Photos) Restorecomment(ownerId int, commentId int) (resp responses.PhotosRestorecomment, err error) {
	params := map[string]interface{}{}

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	params["comment_id"] = commentId

	err = p.SendObjRequest("photos.restoreComment", params, &resp)

	return
}

// Save - Saves photos after successful uploading.
// Parameters:
//   * albumId - ID of the album to save photos to.
//   * groupId - ID of the community to save photos to.
//   * server - Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
//   * photosList - Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
//   * hash - Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
//   * latitude - Geographical latitude, in degrees (from '-90' to '90').
//   * longitude - Geographical longitude, in degrees (from '-180' to '180').
//   * caption - Text describing the photo. 2048 digits max.
func (p *Photos) Save(albumId int, groupId int, server int, photosList string, hash string, latitude json.Number, longitude json.Number, caption string) (resp responses.PhotosSave, err error) {
	params := map[string]interface{}{}

	if albumId > 0 {
		params["album_id"] = albumId
	}

	if groupId > 0 {
		params["group_id"] = groupId
	}

	if server > 0 {
		params["server"] = server
	}

	if photosList != "" {
		params["photos_list"] = photosList
	}

	if hash != "" {
		params["hash"] = hash
	}

	if v, err := latitude.Int64(); err == nil && v > 0 {
		params["latitude"] = v
	} else if v := latitude.String(); v != "" {
		params["latitude"] = v
	}

	if v, err := longitude.Int64(); err == nil && v > 0 {
		params["longitude"] = v
	} else if v := longitude.String(); v != "" {
		params["longitude"] = v
	}

	if caption != "" {
		params["caption"] = caption
	}

	err = p.SendObjRequest("photos.save", params, &resp)

	return
}

// Savemarketalbumphoto - Saves market album photos after successful uploading.
// Parameters:
//   * groupId - Community ID.
//   * photo - Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
//   * server - Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
//   * hash - Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
func (p *Photos) Savemarketalbumphoto(groupId int, photo string, server int, hash string) (resp responses.PhotosSavemarketalbumphoto, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	params["photo"] = photo

	params["server"] = server

	params["hash"] = hash

	err = p.SendObjRequest("photos.saveMarketAlbumPhoto", params, &resp)

	return
}

// Savemarketphoto - Saves market photos after successful uploading.
// Parameters:
//   * groupId - Community ID.
//   * photo - Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
//   * server - Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
//   * hash - Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
//   * cropData - Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
//   * cropHash - Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
func (p *Photos) Savemarketphoto(groupId int, photo string, server int, hash string, cropData string, cropHash string) (resp responses.PhotosSavemarketphoto, err error) {
	params := map[string]interface{}{}

	if groupId > 0 {
		params["group_id"] = groupId
	}

	params["photo"] = photo

	params["server"] = server

	params["hash"] = hash

	if cropData != "" {
		params["crop_data"] = cropData
	}

	if cropHash != "" {
		params["crop_hash"] = cropHash
	}

	err = p.SendObjRequest("photos.saveMarketPhoto", params, &resp)

	return
}

// Savemessagesphoto - Saves a photo after being successfully uploaded. URL obtained with [vk.com/dev/photos.getMessagesUploadServer|photos.getMessagesUploadServer] method.
// Parameters:
//   * photo - Parameter returned when the photo is [vk.com/dev/upload_files|uploaded to the server].
//   * server - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * hash - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (p *Photos) Savemessagesphoto(photo string, server int, hash string) (resp responses.PhotosSavemessagesphoto, err error) {
	params := map[string]interface{}{}

	params["photo"] = photo

	if server > 0 {
		params["server"] = server
	}

	if hash != "" {
		params["hash"] = hash
	}

	err = p.SendObjRequest("photos.saveMessagesPhoto", params, &resp)

	return
}

// Saveownercoverphoto - Saves cover photo after successful uploading.
// Parameters:
//   * hash - Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
//   * photo - Parameter returned when photos are [vk.com/dev/upload_files|uploaded to server].
func (p *Photos) Saveownercoverphoto(hash string, photo string) (resp responses.PhotosSaveownercoverphoto, err error) {
	params := map[string]interface{}{}

	params["hash"] = hash

	params["photo"] = photo

	err = p.SendObjRequest("photos.saveOwnerCoverPhoto", params, &resp)

	return
}

// Saveownerphoto - Saves a profile or community photo. Upload URL can be got with the [vk.com/dev/photos.getOwnerPhotoUploadServer|photos.getOwnerPhotoUploadServer] method.
// Parameters:
//   * server - parameter returned after [vk.com/dev/upload_files|photo upload].
//   * hash - parameter returned after [vk.com/dev/upload_files|photo upload].
//   * photo - parameter returned after [vk.com/dev/upload_files|photo upload].
func (p *Photos) Saveownerphoto(server string, hash string, photo string) (resp responses.PhotosSaveownerphoto, err error) {
	params := map[string]interface{}{}

	if server != "" {
		params["server"] = server
	}

	if hash != "" {
		params["hash"] = hash
	}

	if photo != "" {
		params["photo"] = photo
	}

	err = p.SendObjRequest("photos.saveOwnerPhoto", params, &resp)

	return
}

// Savewallphoto - Saves a photo to a user's or community's wall after being uploaded.
// Parameters:
//   * userId - ID of the user on whose wall the photo will be saved.
//   * groupId - ID of community on whose wall the photo will be saved.
//   * photo - Parameter returned when the the photo is [vk.com/dev/upload_files|uploaded to the server].
//   * server - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * hash - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * latitude - Geographical latitude, in degrees (from '-90' to '90').
//   * longitude - Geographical longitude, in degrees (from '-180' to '180').
//   * caption - Text describing the photo. 2048 digits max.
func (p *Photos) Savewallphoto(userId int, groupId int, photo string, server int, hash string, latitude json.Number, longitude json.Number, caption string) (resp responses.PhotosSavewallphoto, err error) {
	params := map[string]interface{}{}

	if userId > 0 {
		params["user_id"] = userId
	}

	if groupId > 0 {
		params["group_id"] = groupId
	}

	params["photo"] = photo

	if server > 0 {
		params["server"] = server
	}

	if hash != "" {
		params["hash"] = hash
	}

	if v, err := latitude.Int64(); err == nil && v > 0 {
		params["latitude"] = v
	} else if v := latitude.String(); v != "" {
		params["latitude"] = v
	}

	if v, err := longitude.Int64(); err == nil && v > 0 {
		params["longitude"] = v
	} else if v := longitude.String(); v != "" {
		params["longitude"] = v
	}

	if caption != "" {
		params["caption"] = caption
	}

	err = p.SendObjRequest("photos.saveWallPhoto", params, &resp)

	return
}

// Search - Returns a list of photos.
// Parameters:
//   * q - Search query string.
//   * lat - Geographical latitude, in degrees (from '-90' to '90').
//   * long - Geographical longitude, in degrees (from '-180' to '180').
//   * startTime - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * endTime - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * sort - Sort order:
//   * offset - Offset needed to return a specific subset of photos.
//   * count - Number of photos to return.
//   * radius - Radius of search in meters (works very approximately). Available values: '10', '100', '800', '6000', '50000'.
func (p *Photos) Search(q string, lat json.Number, long json.Number, startTime int, endTime int, sort int, offset int, count int, radius int) (resp responses.PhotosSearch, err error) {
	params := map[string]interface{}{}

	if q != "" {
		params["q"] = q
	}

	if v, err := lat.Int64(); err == nil && v > 0 {
		params["lat"] = v
	} else if v := lat.String(); v != "" {
		params["lat"] = v
	}

	if v, err := long.Int64(); err == nil && v > 0 {
		params["long"] = v
	} else if v := long.String(); v != "" {
		params["long"] = v
	}

	if startTime > 0 {
		params["start_time"] = startTime
	}

	if endTime > 0 {
		params["end_time"] = endTime
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

	if radius > 0 {
		params["radius"] = radius
	}

	err = p.SendObjRequest("photos.search", params, &resp)

	return
}
