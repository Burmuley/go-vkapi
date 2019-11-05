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

type Orders struct {
	*VKApi
}

/////////////////////////////////////////////////////////////
// `Orders` methods
/////////////////////////////////////////////////////////////

// Cancelsubscription - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * userId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * subscriptionId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * pendingCancel - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (o *Orders) Cancelsubscription(userId int, subscriptionId int, pendingCancel bool) (resp responses.OrdersCancelsubscription, err error) {
	params := map[string]interface{}{}

	params["user_id"] = userId

	params["subscription_id"] = subscriptionId

	params["pending_cancel"] = pendingCancel

	err = o.SendObjRequest("orders.cancelSubscription", params, &resp)

	return
}

// Changestate - Changes order status.
// Parameters:
//   * orderId - order ID.
//   * action - action to be done with the order. Available actions: *cancel — to cancel unconfirmed order. *charge — to confirm unconfirmed order. Applies only if processing of [vk.com/dev/payments_status|order_change_state] notification failed. *refund — to cancel confirmed order.
//   * appOrderId - internal ID of the order in the application.
//   * testMode - if this parameter is set to 1, this method returns a list of test mode orders. By default — 0.
func (o *Orders) Changestate(orderId int, action string, appOrderId int, testMode bool) (resp responses.OrdersChangestate, err error) {
	params := map[string]interface{}{}

	params["order_id"] = orderId

	params["action"] = action

	if appOrderId > 0 {
		params["app_order_id"] = appOrderId
	}

	params["test_mode"] = testMode

	err = o.SendObjRequest("orders.changeState", params, &resp)

	return
}

// Get - Returns a list of orders.
// Parameters:
//   * offset - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * count - number of returned orders.
//   * testMode - if this parameter is set to 1, this method returns a list of test mode orders. By default — 0.
func (o *Orders) Get(offset int, count int, testMode bool) (resp responses.OrdersGet, err error) {
	params := map[string]interface{}{}

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	params["test_mode"] = testMode

	err = o.SendObjRequest("orders.get", params, &resp)

	return
}

// Getamount - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * userId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * votes - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (o *Orders) Getamount(userId int, votes []string) (resp responses.OrdersGetamount, err error) {
	params := map[string]interface{}{}

	params["user_id"] = userId

	params["votes"] = SliceToString(votes)

	err = o.SendObjRequest("orders.getAmount", params, &resp)

	return
}

// Getbyid - Returns information about orders by their IDs.
// Parameters:
//   * orderId - order ID.
//   * orderIds - order IDs (when information about several orders is requested).
//   * testMode - if this parameter is set to 1, this method returns a list of test mode orders. By default — 0.
func (o *Orders) Getbyid(orderId int, orderIds []int, testMode bool) (resp responses.OrdersGetbyid, err error) {
	params := map[string]interface{}{}

	if orderId > 0 {
		params["order_id"] = orderId
	}

	if len(orderIds) > 0 {
		params["order_ids"] = SliceToString(orderIds)
	}

	params["test_mode"] = testMode

	err = o.SendObjRequest("orders.getById", params, &resp)

	return
}

// Getusersubscriptionbyid - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * userId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * subscriptionId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (o *Orders) Getusersubscriptionbyid(userId int, subscriptionId int) (resp responses.OrdersGetusersubscriptionbyid, err error) {
	params := map[string]interface{}{}

	params["user_id"] = userId

	params["subscription_id"] = subscriptionId

	err = o.SendObjRequest("orders.getUserSubscriptionById", params, &resp)

	return
}

// Getusersubscriptions - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * userId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (o *Orders) Getusersubscriptions(userId int) (resp responses.OrdersGetusersubscriptions, err error) {
	params := map[string]interface{}{}

	params["user_id"] = userId

	err = o.SendObjRequest("orders.getUserSubscriptions", params, &resp)

	return
}

// Updatesubscription - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * userId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * subscriptionId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * price - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (o *Orders) Updatesubscription(userId int, subscriptionId int, price int) (resp responses.OrdersUpdatesubscription, err error) {
	params := map[string]interface{}{}

	params["user_id"] = userId

	params["subscription_id"] = subscriptionId

	params["price"] = price

	err = o.SendObjRequest("orders.updateSubscription", params, &resp)

	return
}
