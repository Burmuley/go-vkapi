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

type Stats struct {
	*VKApi
}

/////////////////////////////////////////////////////////////
// `Stats` methods
/////////////////////////////////////////////////////////////

// Get - Returns statistics of a community or an application.
// Parameters:
//   * groupId - Community ID.
//   * appId - Application ID.
//   * timestampFrom - NO DESCRIPTION IN JSON SCHEMA
//   * timestampTo - NO DESCRIPTION IN JSON SCHEMA
//   * interval - NO DESCRIPTION IN JSON SCHEMA
//   * intervalsCount - NO DESCRIPTION IN JSON SCHEMA
//   * filters - NO DESCRIPTION IN JSON SCHEMA
//   * statsGroups - NO DESCRIPTION IN JSON SCHEMA
//   * extended - NO DESCRIPTION IN JSON SCHEMA
func (s Stats) Get(groupId int, appId int, timestampFrom int, timestampTo int, interval string, intervalsCount int, filters []string, statsGroups []string) (resp responses.StatsGet, err error) {
	params := map[string]interface{}{}

	if groupId > 0 {
		params["group_id"] = groupId
	}

	if appId > 0 {
		params["app_id"] = appId
	}

	if timestampFrom > 0 {
		params["timestamp_from"] = timestampFrom
	}

	if timestampTo > 0 {
		params["timestamp_to"] = timestampTo
	}

	if interval != "" {
		params["interval"] = interval
	}

	if intervalsCount > 0 {
		params["intervals_count"] = intervalsCount
	}

	if len(filters) > 0 {
		params["filters"] = SliceToString(filters)
	}

	if len(statsGroups) > 0 {
		params["stats_groups"] = SliceToString(statsGroups)
	}

	err = s.SendObjRequest("stats.get", params, &resp)

	return
}

// Getpostreach - Returns stats for a wall post.
// Parameters:
//   * ownerId - post owner community id. Specify with "-" sign.
//   * postId - wall post id. Note that stats are available only for '300' last (newest) posts on a community wall.
func (s Stats) Getpostreach(ownerId string, postId int) (resp responses.StatsGetpostreach, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	params["post_id"] = postId

	err = s.SendObjRequest("stats.getPostReach", params, &resp)

	return
}

// Trackvisitor - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * id - NO DESCRIPTION IN JSON SCHEMA
func (s Stats) Trackvisitor(id string) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["id"] = id

	err = s.SendObjRequest("stats.trackVisitor", params, &resp)

	return
}
