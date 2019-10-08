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

type Notifications struct {
	*VKApi
}

/////////////////////////////////////////////////////////////
// `Notifications` methods
/////////////////////////////////////////////////////////////

// Get - Returns a list of notifications about other users' feedback to the current user's wall posts.
// Parameters:
//   * count - Number of notifications to return.
//   * startFrom - NO DESCRIPTION IN JSON SCHEMA
//   * filters - Type of notifications to return: 'wall' — wall posts, 'mentions' — mentions in wall posts, comments, or topics, 'comments' — comments to wall posts, photos, and videos, 'likes' — likes, 'reposted' — wall posts that are copied from the current user's wall, 'followers' — new followers, 'friends' — accepted friend requests
//   * startTime - Earliest timestamp (in Unix time) of a notification to return. By default, 24 hours ago.
//   * endTime - Latest timestamp (in Unix time) of a notification to return. By default, the current time.
func (n *Notifications) Get(count int, startFrom string, filters []string, startTime int, endTime int) (resp responses.NotificationsGet, err error) {
	params := map[string]interface{}{}

	if count > 0 {
		params["count"] = count
	}

	if startFrom != "" {
		params["start_from"] = startFrom
	}

	if len(filters) > 0 {
		params["filters"] = SliceToString(filters)
	}

	if startTime > 0 {
		params["start_time"] = startTime
	}

	if endTime > 0 {
		params["end_time"] = endTime
	}

	err = n.SendObjRequest("notifications.get", params, &resp)

	return
}

// MarkAsViewed - Resets the counter of new notifications about other users' feedback to the current user's wall posts.
func (n *Notifications) MarkAsViewed() (resp responses.NotificationsMarkAsViewed, err error) {
	params := map[string]interface{}{}

	err = n.SendObjRequest("notifications.markAsViewed", params, &resp)

	return
}

// SendMessage - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * userIds - NO DESCRIPTION IN JSON SCHEMA
//   * message - NO DESCRIPTION IN JSON SCHEMA
//   * fragment - NO DESCRIPTION IN JSON SCHEMA
//   * groupId - NO DESCRIPTION IN JSON SCHEMA
func (n *Notifications) SendMessage(userIds []int, message string, fragment string, groupId int) (resp responses.NotificationsSendMessage, err error) {
	params := map[string]interface{}{}

	params["user_ids"] = SliceToString(userIds)

	params["message"] = message

	if fragment != "" {
		params["fragment"] = fragment
	}

	if groupId > 0 {
		params["group_id"] = groupId
	}

	err = n.SendObjRequest("notifications.sendMessage", params, &resp)

	return
}
