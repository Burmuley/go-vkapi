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

type Leads struct {
	*VKApi
}

/////////////////////////////////////////////////////////////
// `Leads` methods
/////////////////////////////////////////////////////////////

// CheckUser - Checks if the user can start the lead.
// Parameters:
//   * leadId - Lead ID.
//   * testResult - Value to be return in 'result' field when test mode is used.
//   * testMode - NO DESCRIPTION IN JSON SCHEMA
//   * autoStart - NO DESCRIPTION IN JSON SCHEMA
//   * age - User age.
//   * country - User country code.
func (l Leads) CheckUser(leadId int, testResult int, testMode bool, autoStart bool, age int, country string) (resp responses.LeadsCheckUser, err error) {
	params := map[string]interface{}{}

	params["lead_id"] = leadId

	if testResult > 0 {
		params["test_result"] = testResult
	}

	params["test_mode"] = testMode

	params["auto_start"] = autoStart

	if age > 0 {
		params["age"] = age
	}

	if country != "" {
		params["country"] = country
	}

	err = l.SendObjRequest("leads.checkUser", params, &resp)

	return
}

// Complete - Completes the lead started by user.
// Parameters:
//   * vkSid - Session obtained as GET parameter when session started.
//   * secret - Secret key from the lead testing interface.
//   * comment - Comment text.
func (l Leads) Complete(vkSid string, secret string, comment string) (resp responses.LeadsComplete, err error) {
	params := map[string]interface{}{}

	params["vk_sid"] = vkSid

	params["secret"] = secret

	if comment != "" {
		params["comment"] = comment
	}

	err = l.SendObjRequest("leads.complete", params, &resp)

	return
}

// GetStats - Returns lead stats data.
// Parameters:
//   * leadId - Lead ID.
//   * secret - Secret key obtained from the lead testing interface.
//   * dateStart - Day to start stats from (YYYY_MM_DD, e.g.2011-09-17).
//   * dateEnd - Day to finish stats (YYYY_MM_DD, e.g.2011-09-17).
func (l Leads) GetStats(leadId int, secret string, dateStart string, dateEnd string) (resp responses.LeadsGetStats, err error) {
	params := map[string]interface{}{}

	params["lead_id"] = leadId

	if secret != "" {
		params["secret"] = secret
	}

	if dateStart != "" {
		params["date_start"] = dateStart
	}

	if dateEnd != "" {
		params["date_end"] = dateEnd
	}

	err = l.SendObjRequest("leads.getStats", params, &resp)

	return
}

// GetUsers - Returns a list of last user actions for the offer.
// Parameters:
//   * offerId - Offer ID.
//   * secret - Secret key obtained in the lead testing interface.
//   * offset - Offset needed to return a specific subset of results.
//   * count - Number of results to return.
//   * status - Action type. Possible values: *'0' — start,, *'1' — finish,, *'2' — blocking users,, *'3' — start in a test mode,, *'4' — finish in a test mode.
//   * reverse - Sort order. Possible values: *'1' — chronological,, *'0' — reverse chronological.
func (l Leads) GetUsers(offerId int, secret string, offset int, count int, status int, reverse bool) (resp responses.LeadsGetUsers, err error) {
	params := map[string]interface{}{}

	params["offer_id"] = offerId

	params["secret"] = secret

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	if status > 0 {
		params["status"] = status
	}

	params["reverse"] = reverse

	err = l.SendObjRequest("leads.getUsers", params, &resp)

	return
}

// MetricHit - Counts the metric event.
// Parameters:
//   * data - Metric data obtained in the lead interface.
func (l Leads) MetricHit(data string) (resp responses.LeadsMetricHit, err error) {
	params := map[string]interface{}{}

	params["data"] = data

	err = l.SendObjRequest("leads.metricHit", params, &resp)

	return
}

// Start - Creates new session for the user passing the offer.
// Parameters:
//   * leadId - Lead ID.
//   * secret - Secret key from the lead testing interface.
//   * uid - NO DESCRIPTION IN JSON SCHEMA
//   * aid - NO DESCRIPTION IN JSON SCHEMA
//   * testMode - NO DESCRIPTION IN JSON SCHEMA
//   * force - NO DESCRIPTION IN JSON SCHEMA
func (l Leads) Start(leadId int, secret string, uid int, aid int, testMode bool, force bool) (resp responses.LeadsStart, err error) {
	params := map[string]interface{}{}

	params["lead_id"] = leadId

	params["secret"] = secret

	if uid > 0 {
		params["uid"] = uid
	}

	if aid > 0 {
		params["aid"] = aid
	}

	params["test_mode"] = testMode

	params["force"] = force

	err = l.SendObjRequest("leads.start", params, &resp)

	return
}
