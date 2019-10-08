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

type AppWidgets struct {
	*VKApi
}

/////////////////////////////////////////////////////////////
// `AppWidgets` methods
/////////////////////////////////////////////////////////////

// Update - Allows to update community app widget
// Parameters:
//   * code - NO DESCRIPTION IN JSON SCHEMA
//   * pType - NO DESCRIPTION IN JSON SCHEMA
func (a *AppWidgets) Update(code string, pType string) (resp objects.BaseOk, err error) {
	params := map[string]interface{}{}

	params["code"] = code

	params["type"] = pType

	err = a.SendObjRequest("appWidgets.update", params, &resp)

	return
}
