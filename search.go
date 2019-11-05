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

type Search struct {
	*VKApi
}

/////////////////////////////////////////////////////////////
// `Search` methods
/////////////////////////////////////////////////////////////

// Gethints - Allows the programmer to do a quick search for any substring.
// Parameters:
//   * q - Search query string.
//   * offset - Offset for querying specific result subset
//   * limit - Maximum number of results to return.
//   * filters - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * fields - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * searchGlobal - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (s *Search) Gethints(q string, offset int, limit int, filters []string, fields []string, searchGlobal bool) (resp responses.SearchGethints, err error) {
	params := map[string]interface{}{}

	if q != "" {
		params["q"] = q
	}

	if offset > 0 {
		params["offset"] = offset
	}

	if limit > 0 {
		params["limit"] = limit
	}

	if len(filters) > 0 {
		params["filters"] = SliceToString(filters)
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	params["search_global"] = searchGlobal

	err = s.SendObjRequest("search.getHints", params, &resp)

	return
}
