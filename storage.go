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

type Storage struct {
	*VKApi
}

/////////////////////////////////////////////////////////////
// `Storage` methods
/////////////////////////////////////////////////////////////

// GetKeys - Returns the names of all variables.
// Parameters:
//   * userId - user id, whose variables names are returned if they were requested with a server method.
//   * global - NO DESCRIPTION IN JSON SCHEMA
//   * offset - NO DESCRIPTION IN JSON SCHEMA
//   * count - amount of variable names the info needs to be collected from.
func (s *Storage) GetKeys(userId int, global bool, offset int, count int) (resp responses.StorageGetKeys, err error) {
	params := map[string]interface{}{}

	if userId > 0 {
		params["user_id"] = userId
	}

	params["global"] = global

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	err = s.SendObjRequest("storage.getKeys", params, &resp)

	return
}

// Set - Saves a value of variable with the name set by 'key' parameter.
// Parameters:
//   * key - NO DESCRIPTION IN JSON SCHEMA
//   * value - NO DESCRIPTION IN JSON SCHEMA
//   * userId - NO DESCRIPTION IN JSON SCHEMA
//   * global - NO DESCRIPTION IN JSON SCHEMA
func (s *Storage) Set(key string, value string, userId int, global bool) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["key"] = key

	if value != "" {
		params["value"] = value
	}

	if userId > 0 {
		params["user_id"] = userId
	}

	params["global"] = global

	err = s.SendObjRequest("storage.set", params, &resp)

	return
}
