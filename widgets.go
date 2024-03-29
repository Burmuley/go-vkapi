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
	"github.com/Burmuley/go-vkapi/objects"
	"github.com/Burmuley/go-vkapi/responses"
)

type Widgets struct {
	*VKApi
}

/////////////////////////////////////////////////////////////
// `Widgets` methods
/////////////////////////////////////////////////////////////

// Getcomments - Gets a list of comments for the page added through the [vk.com/dev/Comments|Comments widget].
// Parameters:
//   * widgetApiId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * url - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * pageId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * order - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * fields - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * offset - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * count - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (w *Widgets) Getcomments(widgetApiId int, url string, pageId string, order string, fields []objects.UsersFields, offset int, count int) (resp responses.WidgetsGetcomments, err error) {
	params := map[string]interface{}{}

	if widgetApiId > 0 {
		params["widget_api_id"] = widgetApiId
	}

	if url != "" {
		params["url"] = url
	}

	if pageId != "" {
		params["page_id"] = pageId
	}

	if order != "" {
		params["order"] = order
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	err = w.SendObjRequest("widgets.getComments", params, &resp)

	return
}

// Getpages - Gets a list of application/site pages where the [vk.com/dev/Comments|Comments widget] or [vk.com/dev/Like|Like widget] is installed.
// Parameters:
//   * widgetApiId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * order - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * period - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * offset - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * count - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (w *Widgets) Getpages(widgetApiId int, order string, period string, offset int, count int) (resp responses.WidgetsGetpages, err error) {
	params := map[string]interface{}{}

	if widgetApiId > 0 {
		params["widget_api_id"] = widgetApiId
	}

	if order != "" {
		params["order"] = order
	}

	if period != "" {
		params["period"] = period
	}

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	err = w.SendObjRequest("widgets.getPages", params, &resp)

	return
}
