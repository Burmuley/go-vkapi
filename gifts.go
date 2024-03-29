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

type Gifts struct {
	*VKApi
}

/////////////////////////////////////////////////////////////
// `Gifts` methods
/////////////////////////////////////////////////////////////

// Get - Returns a list of user gifts.
// Parameters:
//   * userId - User ID.
//   * count - Number of gifts to return.
//   * offset - Offset needed to return a specific subset of results.
func (g *Gifts) Get(userId int, count int, offset int) (resp responses.GiftsGet, err error) {
	params := map[string]interface{}{}

	if userId > 0 {
		params["user_id"] = userId
	}

	if count > 0 {
		params["count"] = count
	}

	if offset > 0 {
		params["offset"] = offset
	}

	err = g.SendObjRequest("gifts.get", params, &resp)

	return
}
