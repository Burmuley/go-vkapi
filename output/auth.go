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

type Auth struct {
	*VKApi
}

/////////////////////////////////////////////////////////////
// `Auth` methods
/////////////////////////////////////////////////////////////

// Checkphone - Checks a user's phone number for correctness.
// Parameters:
//   * phone - Phone number.
//   * clientId - User ID.
//   * clientSecret - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * authByPhone - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (a *Auth) Checkphone(phone string, clientId int, clientSecret string, authByPhone bool) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["phone"] = phone

	if clientId > 0 {
		params["client_id"] = clientId
	}

	if clientSecret != "" {
		params["client_secret"] = clientSecret
	}

	params["auth_by_phone"] = authByPhone

	err = a.SendObjRequest("auth.checkPhone", params, &resp)

	return
}

// Restore - Allows to restore account access using a code received via SMS. " This method is only available for apps with [vk.com/dev/auth_direct|Direct authorization] access. "
// Parameters:
//   * phone - User phone number.
//   * lastName - User last name.
func (a *Auth) Restore(phone string, lastName string) (resp responses.AuthRestore, err error) {
	params := map[string]interface{}{}

	params["phone"] = phone

	params["last_name"] = lastName

	err = a.SendObjRequest("auth.restore", params, &resp)

	return
}
