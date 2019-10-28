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

type Prettycards struct {
	*VKApi
}

/////////////////////////////////////////////////////////////
// `Prettycards` methods
/////////////////////////////////////////////////////////////

// Create - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * ownerId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * photo - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * title - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * link - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * price - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * priceOld - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * button - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (p Prettycards) Create(ownerId int, photo string, title string, link string, price string, priceOld string, button string) (resp responses.PrettycardsCreate, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	params["photo"] = photo

	params["title"] = title

	params["link"] = link

	if price != "" {
		params["price"] = price
	}

	if priceOld != "" {
		params["price_old"] = priceOld
	}

	if button != "" {
		params["button"] = button
	}

	err = p.SendObjRequest("prettyCards.create", params, &resp)

	return
}

// Delete - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * ownerId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * cardId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (p Prettycards) Delete(ownerId int, cardId int) (resp responses.PrettycardsDelete, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	params["card_id"] = cardId

	err = p.SendObjRequest("prettyCards.delete", params, &resp)

	return
}

// Edit - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * ownerId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * cardId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * photo - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * title - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * link - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * price - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * priceOld - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * button - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (p Prettycards) Edit(ownerId int, cardId int, photo string, title string, link string, price string, priceOld string, button string) (resp responses.PrettycardsEdit, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	params["card_id"] = cardId

	if photo != "" {
		params["photo"] = photo
	}

	if title != "" {
		params["title"] = title
	}

	if link != "" {
		params["link"] = link
	}

	if price != "" {
		params["price"] = price
	}

	if priceOld != "" {
		params["price_old"] = priceOld
	}

	if button != "" {
		params["button"] = button
	}

	err = p.SendObjRequest("prettyCards.edit", params, &resp)

	return
}

// Get - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * ownerId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * offset - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * count - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (p Prettycards) Get(ownerId int, offset int, count int) (resp responses.PrettycardsGet, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	err = p.SendObjRequest("prettyCards.get", params, &resp)

	return
}

// Getbyid - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * ownerId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * cardIds - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (p Prettycards) Getbyid(ownerId int, cardIds []int) (resp responses.PrettycardsGetbyid, err error) {
	params := map[string]interface{}{}

	params["owner_id"] = ownerId

	params["card_ids"] = SliceToString(cardIds)

	err = p.SendObjRequest("prettyCards.getById", params, &resp)

	return
}

// Getuploadurl - NO DESCRIPTION IN JSON SCHEMA
func (p Prettycards) Getuploadurl() (resp responses.PrettycardsGetuploadurl, err error) {
	params := map[string]interface{}{}

	err = p.SendObjRequest("prettyCards.getUploadURL", params, &resp)

	return
}
