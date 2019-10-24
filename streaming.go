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

type Streaming struct {
	*VKApi
}

/////////////////////////////////////////////////////////////
// `Streaming` methods
/////////////////////////////////////////////////////////////

// Getserverurl - Allows to receive data for the connection to Streaming API.
func (s Streaming) Getserverurl() (resp responses.StreamingGetserverurl, err error) {
	params := map[string]interface{}{}

	err = s.SendObjRequest("streaming.getServerUrl", params, &resp)

	return
}

// Setsettings - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * monthlyTier - NO DESCRIPTION IN JSON SCHEMA
func (s Streaming) Setsettings(monthlyTier string) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	if monthlyTier != "" {
		params["monthly_tier"] = monthlyTier
	}

	err = s.SendObjRequest("streaming.setSettings", params, &resp)

	return
}
