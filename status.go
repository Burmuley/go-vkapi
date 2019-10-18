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

type Status struct {
	*VKApi
}

/////////////////////////////////////////////////////////////
// `Status` methods
/////////////////////////////////////////////////////////////

// Get - Returns data required to show the status of a user or community.
// Parameters:
//   * userId - User ID or community ID. Use a negative value to designate a community ID.
//   * groupId - NO DESCRIPTION IN JSON SCHEMA
func (s Status) Get(userId int, groupId int) (resp responses.StatusGet, err error) {
	params := map[string]interface{}{}

	if userId > 0 {
		params["user_id"] = userId
	}

	if groupId > 0 {
		params["group_id"] = groupId
	}

	err = s.SendObjRequest("status.get", params, &resp)

	return
}

// Set - Sets a new status for the current user.
// Parameters:
//   * text - Text of the new status.
//   * groupId - Identifier of a community to set a status in. If left blank the status is set to current user.
func (s Status) Set(text string, groupId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	if text != "" {
		params["text"] = text
	}

	if groupId > 0 {
		params["group_id"] = groupId
	}

	err = s.SendObjRequest("status.set", params, &resp)

	return
}
