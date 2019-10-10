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

type Video struct {
	*VKApi
}

/////////////////////////////////////////////////////////////
// `Video` methods
/////////////////////////////////////////////////////////////

// Add - Adds a video to a user or community page.
// Parameters:
//   * targetId - identifier of a user or community to add a video to. Use a negative value to designate a community ID.
//   * videoId - Video ID.
//   * ownerId - ID of the user or community that owns the video. Use a negative value to designate a community ID.
func (v Video) Add(targetId int, videoId int, ownerId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	if targetId > 0 {
		params["target_id"] = targetId
	}

	params["video_id"] = videoId

	params["owner_id"] = ownerId

	err = v.SendObjRequest("video.add", params, &resp)

	return
}

// AddAlbum - Creates an empty album for videos.
// Parameters:
//   * groupId - Community ID (if the album will be created in a community).
//   * title - Album title.
//   * privacy - new access permissions for the album. Possible values: , *'0' – all users,, *'1' – friends only,, *'2' – friends and friends of friends,, *'3' – "only me".
func (v Video) AddAlbum(groupId int, title string, privacy []string) (resp responses.VideoAddAlbum, err error) {
	params := map[string]interface{}{}

	if groupId > 0 {
		params["group_id"] = groupId
	}

	if title != "" {
		params["title"] = title
	}

	if len(privacy) > 0 {
		params["privacy"] = SliceToString(privacy)
	}

	err = v.SendObjRequest("video.addAlbum", params, &resp)

	return
}

// AddToAlbum - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * targetId - NO DESCRIPTION IN JSON SCHEMA
//   * albumId - NO DESCRIPTION IN JSON SCHEMA
//   * albumIds - NO DESCRIPTION IN JSON SCHEMA
//   * ownerId - NO DESCRIPTION IN JSON SCHEMA
//   * videoId - NO DESCRIPTION IN JSON SCHEMA
func (v Video) AddToAlbum(targetId int, albumId int, albumIds []int, ownerId int, videoId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	if targetId > 0 {
		params["target_id"] = targetId
	}

	if albumId > 0 {
		params["album_id"] = albumId
	}

	if len(albumIds) > 0 {
		params["album_ids"] = SliceToString(albumIds)
	}

	params["owner_id"] = ownerId

	params["video_id"] = videoId

	err = v.SendObjRequest("video.addToAlbum", params, &resp)

	return
}

// CreateComment - Adds a new comment on a video.
// Parameters:
//   * ownerId - ID of the user or community that owns the video.
//   * videoId - Video ID.
//   * message - New comment text.
//   * attachments - List of objects attached to the comment, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, '<owner_id>' — ID of the media attachment owner. '<media_id>' — Media attachment ID. Example: "photo100172_166443618,photo66748_265827614"
//   * fromGroup - '1' — to post the comment from a community name (only if 'owner_id'<0)
//   * replyToComment - NO DESCRIPTION IN JSON SCHEMA
//   * stickerId - NO DESCRIPTION IN JSON SCHEMA
//   * guid - NO DESCRIPTION IN JSON SCHEMA
func (v Video) CreateComment(ownerId int, videoId int, message string, attachments []string, fromGroup bool, replyToComment int, stickerId int, guid string) (resp responses.VideoCreateComment, err error) {
	params := map[string]interface{}{}

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	params["video_id"] = videoId

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

	err = v.SendObjRequest("video.createComment", params, &resp)

	return
}

// Delete - Deletes a video from a user or community page.
// Parameters:
//   * videoId - Video ID.
//   * ownerId - ID of the user or community that owns the video.
//   * targetId - NO DESCRIPTION IN JSON SCHEMA
func (v Video) Delete(videoId int, ownerId int, targetId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["video_id"] = videoId

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	if targetId > 0 {
		params["target_id"] = targetId
	}

	err = v.SendObjRequest("video.delete", params, &resp)

	return
}

// DeleteAlbum - Deletes a video album.
// Parameters:
//   * groupId - Community ID (if the album is owned by a community).
//   * albumId - Album ID.
func (v Video) DeleteAlbum(groupId int, albumId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	if groupId > 0 {
		params["group_id"] = groupId
	}

	params["album_id"] = albumId

	err = v.SendObjRequest("video.deleteAlbum", params, &resp)

	return
}

// DeleteComment - Deletes a comment on a video.
// Parameters:
//   * ownerId - ID of the user or community that owns the video.
//   * commentId - ID of the comment to be deleted.
func (v Video) DeleteComment(ownerId int, commentId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	params["comment_id"] = commentId

	err = v.SendObjRequest("video.deleteComment", params, &resp)

	return
}

// Edit - Edits information about a video on a user or community page.
// Parameters:
//   * ownerId - ID of the user or community that owns the video.
//   * videoId - Video ID.
//   * name - New video title.
//   * desc - New video description.
//   * privacyView - Privacy settings in a [vk.com/dev/privacy_setting|special format]. Privacy setting is available for videos uploaded to own profile by user.
//   * privacyComment - Privacy settings for comments in a [vk.com/dev/privacy_setting|special format].
//   * noComments - Disable comments for the group video.
//   * repeat - '1' — to repeat the playback of the video, '0' — to play the video once,
func (v Video) Edit(ownerId int, videoId int, name string, desc string, privacyView []string, privacyComment []string, noComments bool, repeat bool) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	params["video_id"] = videoId

	if name != "" {
		params["name"] = name
	}

	if desc != "" {
		params["desc"] = desc
	}

	if len(privacyView) > 0 {
		params["privacy_view"] = SliceToString(privacyView)
	}

	if len(privacyComment) > 0 {
		params["privacy_comment"] = SliceToString(privacyComment)
	}

	params["no_comments"] = noComments

	params["repeat"] = repeat

	err = v.SendObjRequest("video.edit", params, &resp)

	return
}

// EditAlbum - Edits the title of a video album.
// Parameters:
//   * groupId - Community ID (if the album edited is owned by a community).
//   * albumId - Album ID.
//   * title - New album title.
//   * privacy - new access permissions for the album. Possible values: , *'0' – all users,, *'1' – friends only,, *'2' – friends and friends of friends,, *'3' – "only me".
func (v Video) EditAlbum(groupId int, albumId int, title string, privacy []string) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	if groupId > 0 {
		params["group_id"] = groupId
	}

	params["album_id"] = albumId

	params["title"] = title

	if len(privacy) > 0 {
		params["privacy"] = SliceToString(privacy)
	}

	err = v.SendObjRequest("video.editAlbum", params, &resp)

	return
}

// EditComment - Edits the text of a comment on a video.
// Parameters:
//   * ownerId - ID of the user or community that owns the video.
//   * commentId - Comment ID.
//   * message - New comment text.
//   * attachments - List of objects attached to the comment, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", '' — Type of media attachment: 'photo' — photo, 'video' — video, 'audio' — audio, 'doc' — document, '<owner_id>' — ID of the media attachment owner. '<media_id>' — Media attachment ID. Example: "photo100172_166443618,photo66748_265827614"
func (v Video) EditComment(ownerId int, commentId int, message string, attachments []string) (resp responses.Ok, err error) {
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

	err = v.SendObjRequest("video.editComment", params, &resp)

	return
}

// Get - Returns detailed information about videos.
// Parameters:
//   * ownerId - ID of the user or community that owns the video(s).
//   * videos - Video IDs, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", Use a negative value to designate a community ID. Example: "-4363_136089719,13245770_137352259"
//   * albumId - ID of the album containing the video(s).
//   * count - Number of videos to return.
//   * offset - Offset needed to return a specific subset of videos.
//   * extended - '1' — to return an extended response with additional fields
func (v Video) Get(ownerId int, videos []string, albumId int, count int, offset int) (resp responses.VideoGet, err error) {
	params := map[string]interface{}{}
	params["extended"] = "0"

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	if len(videos) > 0 {
		params["videos"] = SliceToString(videos)
	}

	if albumId > 0 {
		params["album_id"] = albumId
	}

	if count > 0 {
		params["count"] = count
	}

	if offset > 0 {
		params["offset"] = offset
	}

	err = v.SendObjRequest("video.get", params, &resp)

	return
}

// GetExtended - Returns detailed information about videos.
// Parameters:
//   * ownerId - ID of the user or community that owns the video(s).
//   * videos - Video IDs, in the following format: "<owner_id>_<media_id>,<owner_id>_<media_id>", Use a negative value to designate a community ID. Example: "-4363_136089719,13245770_137352259"
//   * albumId - ID of the album containing the video(s).
//   * count - Number of videos to return.
//   * offset - Offset needed to return a specific subset of videos.
//   * extended - '1' — to return an extended response with additional fields
func (v Video) GetExtended(ownerId int, videos []string, albumId int, count int, offset int) (resp responses.VideoGetExtended, err error) {
	params := map[string]interface{}{}
	params["extended"] = "1"

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	if len(videos) > 0 {
		params["videos"] = SliceToString(videos)
	}

	if albumId > 0 {
		params["album_id"] = albumId
	}

	if count > 0 {
		params["count"] = count
	}

	if offset > 0 {
		params["offset"] = offset
	}

	err = v.SendObjRequest("video.get", params, &resp)

	return
}

// GetAlbumById - Returns video album info
// Parameters:
//   * ownerId - identifier of a user or community to add a video to. Use a negative value to designate a community ID.
//   * albumId - Album ID.
func (v Video) GetAlbumById(ownerId int, albumId int) (resp responses.VideoGetAlbumById, err error) {
	params := map[string]interface{}{}

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	params["album_id"] = albumId

	err = v.SendObjRequest("video.getAlbumById", params, &resp)

	return
}

// GetAlbums - Returns a list of video albums owned by a user or community.
// Parameters:
//   * ownerId - ID of the user or community that owns the video album(s).
//   * offset - Offset needed to return a specific subset of video albums.
//   * count - Number of video albums to return.
//   * extended - '1' — to return additional information about album privacy settings for the current user
//   * needSystem - NO DESCRIPTION IN JSON SCHEMA
func (v Video) GetAlbums(ownerId int, offset int, count int, needSystem bool) (resp responses.VideoGetAlbums, err error) {
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

	params["need_system"] = needSystem

	err = v.SendObjRequest("video.getAlbums", params, &resp)

	return
}

// GetAlbumsExtended - Returns a list of video albums owned by a user or community.
// Parameters:
//   * ownerId - ID of the user or community that owns the video album(s).
//   * offset - Offset needed to return a specific subset of video albums.
//   * count - Number of video albums to return.
//   * extended - '1' — to return additional information about album privacy settings for the current user
//   * needSystem - NO DESCRIPTION IN JSON SCHEMA
func (v Video) GetAlbumsExtended(ownerId int, offset int, count int, needSystem bool) (resp responses.VideoGetAlbumsExtended, err error) {
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

	params["need_system"] = needSystem

	err = v.SendObjRequest("video.getAlbums", params, &resp)

	return
}

// GetAlbumsByVideo - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * targetId - NO DESCRIPTION IN JSON SCHEMA
//   * ownerId - NO DESCRIPTION IN JSON SCHEMA
//   * videoId - NO DESCRIPTION IN JSON SCHEMA
//   * extended - NO DESCRIPTION IN JSON SCHEMA
func (v Video) GetAlbumsByVideo(targetId int, ownerId int, videoId int) (resp responses.VideoGetAlbumsByVideo, err error) {
	params := map[string]interface{}{}
	params["extended"] = "0"

	if targetId > 0 {
		params["target_id"] = targetId
	}

	params["owner_id"] = ownerId

	params["video_id"] = videoId

	err = v.SendObjRequest("video.getAlbumsByVideo", params, &resp)

	return
}

// GetAlbumsByVideoExtended - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * targetId - NO DESCRIPTION IN JSON SCHEMA
//   * ownerId - NO DESCRIPTION IN JSON SCHEMA
//   * videoId - NO DESCRIPTION IN JSON SCHEMA
//   * extended - NO DESCRIPTION IN JSON SCHEMA
func (v Video) GetAlbumsByVideoExtended(targetId int, ownerId int, videoId int) (resp responses.VideoGetAlbumsByVideoExtended, err error) {
	params := map[string]interface{}{}
	params["extended"] = "1"

	if targetId > 0 {
		params["target_id"] = targetId
	}

	params["owner_id"] = ownerId

	params["video_id"] = videoId

	err = v.SendObjRequest("video.getAlbumsByVideo", params, &resp)

	return
}

// GetComments - Returns a list of comments on a video.
// Parameters:
//   * ownerId - ID of the user or community that owns the video.
//   * videoId - Video ID.
//   * needLikes - '1' — to return an additional 'likes' field
//   * startCommentId - NO DESCRIPTION IN JSON SCHEMA
//   * offset - Offset needed to return a specific subset of comments.
//   * count - Number of comments to return.
//   * sort - Sort order: 'asc' — oldest comment first, 'desc' — newest comment first
//   * extended - NO DESCRIPTION IN JSON SCHEMA
//   * fields - NO DESCRIPTION IN JSON SCHEMA
func (v Video) GetComments(ownerId int, videoId int, needLikes bool, startCommentId int, offset int, count int, sort string, fields []string) (resp responses.VideoGetComments, err error) {
	params := map[string]interface{}{}
	params["extended"] = "0"

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	params["video_id"] = videoId

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

	err = v.SendObjRequest("video.getComments", params, &resp)

	return
}

// GetCommentsExtended - Returns a list of comments on a video.
// Parameters:
//   * ownerId - ID of the user or community that owns the video.
//   * videoId - Video ID.
//   * needLikes - '1' — to return an additional 'likes' field
//   * startCommentId - NO DESCRIPTION IN JSON SCHEMA
//   * offset - Offset needed to return a specific subset of comments.
//   * count - Number of comments to return.
//   * sort - Sort order: 'asc' — oldest comment first, 'desc' — newest comment first
//   * extended - NO DESCRIPTION IN JSON SCHEMA
//   * fields - NO DESCRIPTION IN JSON SCHEMA
func (v Video) GetCommentsExtended(ownerId int, videoId int, needLikes bool, startCommentId int, offset int, count int, sort string, fields []string) (resp responses.VideoGetCommentsExtended, err error) {
	params := map[string]interface{}{}
	params["extended"] = "1"

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	params["video_id"] = videoId

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

	err = v.SendObjRequest("video.getComments", params, &resp)

	return
}

// RemoveFromAlbum - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * targetId - NO DESCRIPTION IN JSON SCHEMA
//   * albumId - NO DESCRIPTION IN JSON SCHEMA
//   * albumIds - NO DESCRIPTION IN JSON SCHEMA
//   * ownerId - NO DESCRIPTION IN JSON SCHEMA
//   * videoId - NO DESCRIPTION IN JSON SCHEMA
func (v Video) RemoveFromAlbum(targetId int, albumId int, albumIds []int, ownerId int, videoId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	if targetId > 0 {
		params["target_id"] = targetId
	}

	if albumId > 0 {
		params["album_id"] = albumId
	}

	if len(albumIds) > 0 {
		params["album_ids"] = SliceToString(albumIds)
	}

	params["owner_id"] = ownerId

	params["video_id"] = videoId

	err = v.SendObjRequest("video.removeFromAlbum", params, &resp)

	return
}

// ReorderAlbums - Reorders the album in the list of user video albums.
// Parameters:
//   * ownerId - ID of the user or community that owns the albums..
//   * albumId - Album ID.
//   * before - ID of the album before which the album in question shall be placed.
//   * after - ID of the album after which the album in question shall be placed.
func (v Video) ReorderAlbums(ownerId int, albumId int, before int, after int) (resp responses.Ok, err error) {
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

	err = v.SendObjRequest("video.reorderAlbums", params, &resp)

	return
}

// ReorderVideos - Reorders the video in the video album.
// Parameters:
//   * targetId - ID of the user or community that owns the album with videos.
//   * albumId - ID of the video album.
//   * ownerId - ID of the user or community that owns the video.
//   * videoId - ID of the video.
//   * beforeOwnerId - ID of the user or community that owns the video before which the video in question shall be placed.
//   * beforeVideoId - ID of the video before which the video in question shall be placed.
//   * afterOwnerId - ID of the user or community that owns the video after which the photo in question shall be placed.
//   * afterVideoId - ID of the video after which the photo in question shall be placed.
func (v Video) ReorderVideos(targetId int, albumId int, ownerId int, videoId int, beforeOwnerId int, beforeVideoId int, afterOwnerId int, afterVideoId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	if targetId > 0 {
		params["target_id"] = targetId
	}

	if albumId > 0 {
		params["album_id"] = albumId
	}

	params["owner_id"] = ownerId

	params["video_id"] = videoId

	if beforeOwnerId > 0 {
		params["before_owner_id"] = beforeOwnerId
	}

	if beforeVideoId > 0 {
		params["before_video_id"] = beforeVideoId
	}

	if afterOwnerId > 0 {
		params["after_owner_id"] = afterOwnerId
	}

	if afterVideoId > 0 {
		params["after_video_id"] = afterVideoId
	}

	err = v.SendObjRequest("video.reorderVideos", params, &resp)

	return
}

// Report - Reports (submits a complaint about) a video.
// Parameters:
//   * ownerId - ID of the user or community that owns the video.
//   * videoId - Video ID.
//   * reason - Reason for the complaint: '0' – spam, '1' – child pornography, '2' – extremism, '3' – violence, '4' – drug propaganda, '5' – adult material, '6' – insult, abuse
//   * comment - Comment describing the complaint.
//   * searchQuery - (If the video was found in search results.) Search query string.
func (v Video) Report(ownerId int, videoId int, reason int, comment string, searchQuery string) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	params["video_id"] = videoId

	if reason > 0 {
		params["reason"] = reason
	}

	if comment != "" {
		params["comment"] = comment
	}

	if searchQuery != "" {
		params["search_query"] = searchQuery
	}

	err = v.SendObjRequest("video.report", params, &resp)

	return
}

// ReportComment - Reports (submits a complaint about) a comment on a video.
// Parameters:
//   * ownerId - ID of the user or community that owns the video.
//   * commentId - ID of the comment being reported.
//   * reason - Reason for the complaint: , 0 – spam , 1 – child pornography , 2 – extremism , 3 – violence , 4 – drug propaganda , 5 – adult material , 6 – insult, abuse
func (v Video) ReportComment(ownerId int, commentId int, reason int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	params["comment_id"] = commentId

	if reason > 0 {
		params["reason"] = reason
	}

	err = v.SendObjRequest("video.reportComment", params, &resp)

	return
}

// Restore - Restores a previously deleted video.
// Parameters:
//   * videoId - Video ID.
//   * ownerId - ID of the user or community that owns the video.
func (v Video) Restore(videoId int, ownerId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["video_id"] = videoId

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	err = v.SendObjRequest("video.restore", params, &resp)

	return
}

// RestoreComment - Restores a previously deleted comment on a video.
// Parameters:
//   * ownerId - ID of the user or community that owns the video.
//   * commentId - ID of the deleted comment.
func (v Video) RestoreComment(ownerId int, commentId int) (resp responses.VideoRestoreComment, err error) {
	params := map[string]interface{}{}

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	params["comment_id"] = commentId

	err = v.SendObjRequest("video.restoreComment", params, &resp)

	return
}

// Save - Returns a server address (required for upload) and video data.
// Parameters:
//   * name - Name of the video.
//   * description - Description of the video.
//   * isPrivate - '1' — to designate the video as private (send it via a private message), the video will not appear on the user's video list and will not be available by ID for other users, '0' — not to designate the video as private
//   * wallpost - '1' — to post the saved video on a user's wall, '0' — not to post the saved video on a user's wall
//   * link - URL for embedding the video from an external website.
//   * groupId - ID of the community in which the video will be saved. By default, the current user's page.
//   * albumId - ID of the album to which the saved video will be added.
//   * privacyView - NO DESCRIPTION IN JSON SCHEMA
//   * privacyComment - NO DESCRIPTION IN JSON SCHEMA
//   * noComments - NO DESCRIPTION IN JSON SCHEMA
//   * repeat - '1' — to repeat the playback of the video, '0' — to play the video once,
//   * compression - NO DESCRIPTION IN JSON SCHEMA
func (v Video) Save(name string, description string, isPrivate bool, wallpost bool, link string, groupId int, albumId int, privacyView []string, privacyComment []string, noComments bool, repeat bool, compression bool) (resp responses.VideoSave, err error) {
	params := map[string]interface{}{}

	if name != "" {
		params["name"] = name
	}

	if description != "" {
		params["description"] = description
	}

	params["is_private"] = isPrivate

	params["wallpost"] = wallpost

	if link != "" {
		params["link"] = link
	}

	if groupId > 0 {
		params["group_id"] = groupId
	}

	if albumId > 0 {
		params["album_id"] = albumId
	}

	if len(privacyView) > 0 {
		params["privacy_view"] = SliceToString(privacyView)
	}

	if len(privacyComment) > 0 {
		params["privacy_comment"] = SliceToString(privacyComment)
	}

	params["no_comments"] = noComments

	params["repeat"] = repeat

	params["compression"] = compression

	err = v.SendObjRequest("video.save", params, &resp)

	return
}

// Search - Returns a list of videos under the set search criterion.
// Parameters:
//   * q - Search query string (e.g., 'The Beatles').
//   * sort - Sort order: '1' — by duration, '2' — by relevance, '0' — by date added
//   * hd - If not null, only searches for high-definition videos.
//   * adult - '1' — to disable the Safe Search filter, '0' — to enable the Safe Search filter
//   * filters - Filters to apply: 'youtube' — return YouTube videos only, 'vimeo' — return Vimeo videos only, 'short' — return short videos only, 'long' — return long videos only
//   * searchOwn - NO DESCRIPTION IN JSON SCHEMA
//   * offset - Offset needed to return a specific subset of videos.
//   * longer - NO DESCRIPTION IN JSON SCHEMA
//   * shorter - NO DESCRIPTION IN JSON SCHEMA
//   * count - Number of videos to return.
//   * extended - NO DESCRIPTION IN JSON SCHEMA
func (v Video) Search(q string, sort int, hd int, adult bool, filters []string, searchOwn bool, offset int, longer int, shorter int, count int) (resp responses.VideoSearch, err error) {
	params := map[string]interface{}{}
	params["extended"] = "0"

	params["q"] = q

	if sort > 0 {
		params["sort"] = sort
	}

	if hd > 0 {
		params["hd"] = hd
	}

	params["adult"] = adult

	if len(filters) > 0 {
		params["filters"] = SliceToString(filters)
	}

	params["search_own"] = searchOwn

	if offset > 0 {
		params["offset"] = offset
	}

	if longer > 0 {
		params["longer"] = longer
	}

	if shorter > 0 {
		params["shorter"] = shorter
	}

	if count > 0 {
		params["count"] = count
	}

	err = v.SendObjRequest("video.search", params, &resp)

	return
}

// SearchExtended - Returns a list of videos under the set search criterion.
// Parameters:
//   * q - Search query string (e.g., 'The Beatles').
//   * sort - Sort order: '1' — by duration, '2' — by relevance, '0' — by date added
//   * hd - If not null, only searches for high-definition videos.
//   * adult - '1' — to disable the Safe Search filter, '0' — to enable the Safe Search filter
//   * filters - Filters to apply: 'youtube' — return YouTube videos only, 'vimeo' — return Vimeo videos only, 'short' — return short videos only, 'long' — return long videos only
//   * searchOwn - NO DESCRIPTION IN JSON SCHEMA
//   * offset - Offset needed to return a specific subset of videos.
//   * longer - NO DESCRIPTION IN JSON SCHEMA
//   * shorter - NO DESCRIPTION IN JSON SCHEMA
//   * count - Number of videos to return.
//   * extended - NO DESCRIPTION IN JSON SCHEMA
func (v Video) SearchExtended(q string, sort int, hd int, adult bool, filters []string, searchOwn bool, offset int, longer int, shorter int, count int) (resp responses.VideoSearchExtended, err error) {
	params := map[string]interface{}{}
	params["extended"] = "1"

	params["q"] = q

	if sort > 0 {
		params["sort"] = sort
	}

	if hd > 0 {
		params["hd"] = hd
	}

	params["adult"] = adult

	if len(filters) > 0 {
		params["filters"] = SliceToString(filters)
	}

	params["search_own"] = searchOwn

	if offset > 0 {
		params["offset"] = offset
	}

	if longer > 0 {
		params["longer"] = longer
	}

	if shorter > 0 {
		params["shorter"] = shorter
	}

	if count > 0 {
		params["count"] = count
	}

	err = v.SendObjRequest("video.search", params, &resp)

	return
}
