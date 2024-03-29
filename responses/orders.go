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
// Code generator location: https://github.com/Burmuley/go-vkapi-gen                                       //
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

package responses

import (
	"github.com/Burmuley/go-vkapi/objects"
)

/////////////////////////////////////////////////////////////
// `orders` group of responses
/////////////////////////////////////////////////////////////

// OrdersCancelsubscription type represents `orders_cancelSubscription_response` API response object
type OrdersCancelsubscription objects.BaseBoolInt // Result

// OrdersChangestate type represents `orders_changeState_response` API response object
type OrdersChangestate string // New state

// OrdersGetamount type represents `orders_getAmount_response` API response object
type OrdersGetamount objects.OrdersAmount

// OrdersGetbyid type represents `orders_getById_response` API response object
type OrdersGetbyid objects.OrdersOrder

// OrdersGetusersubscriptionbyid type represents `orders_getUserSubscriptionById_response` API response object
type OrdersGetusersubscriptionbyid objects.OrdersSubscription

// OrdersGetusersubscriptions type represents `orders_getUserSubscriptions_response` API response object
type OrdersGetusersubscriptions struct {
	Count int                          `json:"count"` // Total number
	Items []objects.OrdersSubscription `json:"items"`
}

// OrdersGet type represents `orders_get_response` API response object
type OrdersGet objects.OrdersOrder

// OrdersUpdatesubscription type represents `orders_updateSubscription_response` API response object
type OrdersUpdatesubscription objects.BaseBoolInt // Result
