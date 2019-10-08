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
// Source schema can be found at https://github.com/VKCOM/vk-api-schema/blob/master/responses.json         //
// Code generator location: https://gitlab.com/Burmuley/go-vkapi-gen                                       //
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

package responses

import "gitlab.com/Burmuley/go-vkapi/objects"

/////////////////////////////////////////////////////////////
// `orders` group of responses
/////////////////////////////////////////////////////////////

// OrdersGetAmount type represents `orders_getAmount_response` API response object
type OrdersGetAmount objects.OrdersAmount

// OrdersGetUserSubscriptions type represents `orders_getUserSubscriptions_response` API response object
type OrdersGetUserSubscriptions struct {
	Count int                          `json:"count"` // Total number
	Items []objects.OrdersSubscription `json:"items"`
}

// OrdersGetUserSubscriptionById type represents `orders_getUserSubscriptionById_response` API response object
type OrdersGetUserSubscriptionById objects.OrdersSubscription

// OrdersGetById type represents `orders_getById_response` API response object
type OrdersGetById objects.OrdersOrder

// OrdersUpdateSubscription type represents `orders_updateSubscription_response` API response object
type OrdersUpdateSubscription objects.BaseBoolInt // Result

// OrdersCancelSubscription type represents `orders_cancelSubscription_response` API response object
type OrdersCancelSubscription objects.BaseBoolInt // Result

// OrdersChangeState type represents `orders_changeState_response` API response object
type OrdersChangeState string // New state

// OrdersGet type represents `orders_get_response` API response object
type OrdersGet objects.OrdersOrder
