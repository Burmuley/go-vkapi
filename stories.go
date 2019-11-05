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

type Stories struct {
	*VKApi
}

/////////////////////////////////////////////////////////////
// `Stories` methods
/////////////////////////////////////////////////////////////

// Banowner - Allows to hide stories from chosen sources from current user's feed.
// Parameters:
//   * ownersIds - List of sources IDs
func (s *Stories) Banowner(ownersIds []int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["owners_ids"] = SliceToString(ownersIds)

	err = s.SendObjRequest("stories.banOwner", params, &resp)

	return
}

// Delete - Allows to delete story.
// Parameters:
//   * ownerId - Story owner's ID. Current user id is used by default.
//   * storyId - Story ID.
func (s *Stories) Delete(ownerId int, storyId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	params["story_id"] = storyId

	err = s.SendObjRequest("stories.delete", params, &resp)

	return
}

// Get - Returns stories available for current user.
// Parameters:
//   * ownerId - Owner ID.
//   * extended - '1' — to return additional fields for users and communities. Default value is 0.
func (s *Stories) Get(ownerId int) (resp responses.StoriesGet, err error) {
	params := map[string]interface{}{}
	params["extended"] = "0"

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	err = s.SendObjRequest("stories.get", params, &resp)

	return
}

// GetExtended - Returns stories available for current user.
// Parameters:
//   * ownerId - Owner ID.
//   * extended - '1' — to return additional fields for users and communities. Default value is 0.
func (s *Stories) GetExtended(ownerId int) (resp responses.StoriesGetExtended, err error) {
	params := map[string]interface{}{}
	params["extended"] = "1"

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	err = s.SendObjRequest("stories.get", params, &resp)

	return
}

// Getbanned - Returns list of sources hidden from current user's feed.
// Parameters:
//   * extended - '1' — to return additional fields for users and communities. Default value is 0.
//   * fields - Additional fields to return
func (s *Stories) Getbanned(fields []objects.BaseUserGroupFields) (resp responses.StoriesGetbanned, err error) {
	params := map[string]interface{}{}
	params["extended"] = "0"

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	err = s.SendObjRequest("stories.getBanned", params, &resp)

	return
}

// GetbannedExtended - Returns list of sources hidden from current user's feed.
// Parameters:
//   * extended - '1' — to return additional fields for users and communities. Default value is 0.
//   * fields - Additional fields to return
func (s *Stories) GetbannedExtended(fields []objects.BaseUserGroupFields) (resp responses.StoriesGetbannedExtended, err error) {
	params := map[string]interface{}{}
	params["extended"] = "1"

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	err = s.SendObjRequest("stories.getBanned", params, &resp)

	return
}

// Getbyid - Returns story by its ID.
// Parameters:
//   * stories - Stories IDs separated by commas. Use format {owner_id}+'_'+{story_id}, for example, 12345_54331.
//   * extended - '1' — to return additional fields for users and communities. Default value is 0.
//   * fields - Additional fields to return
func (s *Stories) Getbyid(stories []string, fields []objects.BaseUserGroupFields) (resp responses.StoriesGetbyid, err error) {
	params := map[string]interface{}{}
	params["extended"] = "0"

	params["stories"] = SliceToString(stories)

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	err = s.SendObjRequest("stories.getById", params, &resp)

	return
}

// GetbyidExtended - Returns story by its ID.
// Parameters:
//   * stories - Stories IDs separated by commas. Use format {owner_id}+'_'+{story_id}, for example, 12345_54331.
//   * extended - '1' — to return additional fields for users and communities. Default value is 0.
//   * fields - Additional fields to return
func (s *Stories) GetbyidExtended(stories []string, fields []objects.BaseUserGroupFields) (resp responses.StoriesGetbyidExtended, err error) {
	params := map[string]interface{}{}
	params["extended"] = "1"

	params["stories"] = SliceToString(stories)

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	err = s.SendObjRequest("stories.getById", params, &resp)

	return
}

// Getphotouploadserver - Returns URL for uploading a story with photo.
// Parameters:
//   * addToNews - 1 — to add the story to friend's feed.
//   * userIds - List of users IDs who can see the story.
//   * replyToStory - ID of the story to reply with the current.
//   * linkText - Link text (for community's stories only).
//   * linkUrl - Link URL. Internal links on https://vk.com only.
//   * groupId - ID of the community to upload the story (should be verified or with the "fire" icon).
func (s *Stories) Getphotouploadserver(addToNews bool, userIds []int, replyToStory string, linkText objects.StoriesUploadLinkText, linkUrl string, groupId int) (resp responses.StoriesGetphotouploadserver, err error) {
	params := map[string]interface{}{}

	params["add_to_news"] = addToNews

	if len(userIds) > 0 {
		params["user_ids"] = SliceToString(userIds)
	}

	if replyToStory != "" {
		params["reply_to_story"] = replyToStory
	}

	if linkUrl != "" {
		params["link_url"] = linkUrl
	}

	if groupId > 0 {
		params["group_id"] = groupId
	}

	err = s.SendObjRequest("stories.getPhotoUploadServer", params, &resp)

	return
}

// Getreplies - Returns replies to the story.
// Parameters:
//   * ownerId - Story owner ID.
//   * storyId - Story ID.
//   * accessKey - Access key for the private object.
//   * extended - '1' — to return additional fields for users and communities. Default value is 0.
//   * fields - Additional fields to return
func (s *Stories) Getreplies(ownerId int, storyId int, accessKey string, fields []objects.BaseUserGroupFields) (resp responses.StoriesGetreplies, err error) {
	params := map[string]interface{}{}
	params["extended"] = "0"

	params["owner_id"] = ownerId

	params["story_id"] = storyId

	if accessKey != "" {
		params["access_key"] = accessKey
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	err = s.SendObjRequest("stories.getReplies", params, &resp)

	return
}

// GetrepliesExtended - Returns replies to the story.
// Parameters:
//   * ownerId - Story owner ID.
//   * storyId - Story ID.
//   * accessKey - Access key for the private object.
//   * extended - '1' — to return additional fields for users and communities. Default value is 0.
//   * fields - Additional fields to return
func (s *Stories) GetrepliesExtended(ownerId int, storyId int, accessKey string, fields []objects.BaseUserGroupFields) (resp responses.StoriesGetrepliesExtended, err error) {
	params := map[string]interface{}{}
	params["extended"] = "1"

	params["owner_id"] = ownerId

	params["story_id"] = storyId

	if accessKey != "" {
		params["access_key"] = accessKey
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	err = s.SendObjRequest("stories.getReplies", params, &resp)

	return
}

// Getstats - Returns stories available for current user.
// Parameters:
//   * ownerId - Story owner ID.
//   * storyId - Story ID.
func (s *Stories) Getstats(ownerId int, storyId int) (resp responses.StoriesGetstats, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	params["story_id"] = storyId

	err = s.SendObjRequest("stories.getStats", params, &resp)

	return
}

// Getvideouploadserver - Allows to receive URL for uploading story with video.
// Parameters:
//   * addToNews - 1 — to add the story to friend's feed.
//   * userIds - List of users IDs who can see the story.
//   * replyToStory - ID of the story to reply with the current.
//   * linkText - Link text (for community's stories only).
//   * linkUrl - Link URL. Internal links on https://vk.com only.
//   * groupId - ID of the community to upload the story (should be verified or with the "fire" icon).
func (s *Stories) Getvideouploadserver(addToNews bool, userIds []int, replyToStory string, linkText objects.StoriesUploadLinkText, linkUrl string, groupId int) (resp responses.StoriesGetvideouploadserver, err error) {
	params := map[string]interface{}{}

	params["add_to_news"] = addToNews

	if len(userIds) > 0 {
		params["user_ids"] = SliceToString(userIds)
	}

	if replyToStory != "" {
		params["reply_to_story"] = replyToStory
	}

	if linkUrl != "" {
		params["link_url"] = linkUrl
	}

	if groupId > 0 {
		params["group_id"] = groupId
	}

	err = s.SendObjRequest("stories.getVideoUploadServer", params, &resp)

	return
}

// Getviewers - Returns a list of story viewers.
// Parameters:
//   * ownerId - Story owner ID.
//   * storyId - Story ID.
//   * count - Maximum number of results.
//   * offset - Offset needed to return a specific subset of results.
//   * extended - '1' — to return detailed information about photos
func (s *Stories) Getviewers(ownerId int, storyId int, count int, offset int) (resp responses.StoriesGetviewers, err error) {
	params := map[string]interface{}{}
	params["extended"] = "0"

	params["owner_id"] = ownerId

	params["story_id"] = storyId

	if count > 0 {
		params["count"] = count
	}

	if offset > 0 {
		params["offset"] = offset
	}

	err = s.SendObjRequest("stories.getViewers", params, &resp)

	return
}

// GetviewersExtended - Returns a list of story viewers.
// Parameters:
//   * ownerId - Story owner ID.
//   * storyId - Story ID.
//   * count - Maximum number of results.
//   * offset - Offset needed to return a specific subset of results.
//   * extended - '1' — to return detailed information about photos
func (s *Stories) GetviewersExtended(ownerId int, storyId int, count int, offset int) (resp responses.StoriesGetviewersExtended, err error) {
	params := map[string]interface{}{}
	params["extended"] = "1"

	params["owner_id"] = ownerId

	params["story_id"] = storyId

	if count > 0 {
		params["count"] = count
	}

	if offset > 0 {
		params["offset"] = offset
	}

	err = s.SendObjRequest("stories.getViewers", params, &resp)

	return
}

// Hideallreplies - Hides all replies in the last 24 hours from the user to current user's stories.
// Parameters:
//   * ownerId - ID of the user whose replies should be hidden.
//   * groupId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (s *Stories) Hideallreplies(ownerId int, groupId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	if groupId > 0 {
		params["group_id"] = groupId
	}

	err = s.SendObjRequest("stories.hideAllReplies", params, &resp)

	return
}

// Hidereply - Hides the reply to the current user's story.
// Parameters:
//   * ownerId - ID of the user whose replies should be hidden.
//   * storyId - Story ID.
func (s *Stories) Hidereply(ownerId int, storyId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	params["story_id"] = storyId

	err = s.SendObjRequest("stories.hideReply", params, &resp)

	return
}

// Unbanowner - Allows to show stories from hidden sources in current user's feed.
// Parameters:
//   * ownersIds - List of hidden sources to show stories from.
func (s *Stories) Unbanowner(ownersIds []int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["owners_ids"] = SliceToString(ownersIds)

	err = s.SendObjRequest("stories.unbanOwner", params, &resp)

	return
}
