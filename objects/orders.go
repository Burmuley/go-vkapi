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

package objects

/////////////////////////////////////////////////////////////////////////////////////////////////////////////
// `orders` group of objects
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

// OrdersAmount type represents `orders_amount` API object
type OrdersAmount struct {
	Amounts  []*OrdersAmountItem `json:"amounts"`
	Currency string              `json:"currency"` // Currency name
}

// OrdersAmountItem type represents `orders_amount_item` API object
type OrdersAmountItem struct {
	Amount      int    `json:"amount"`      // Votes amount in user's currency
	Description string `json:"description"` // Amount description
	Votes       string `json:"votes"`       // Votes number
}

// OrdersOrder type represents `orders_order` API object
type OrdersOrder struct {
	Amount              int    `json:"amount"`                // Amount
	AppOrderId          int    `json:"app_order_id"`          // App order ID
	CancelTransactionId int    `json:"cancel_transaction_id"` // Cancel transaction ID
	Date                int    `json:"date"`                  // Date of creation in Unixtime
	Id                  int    `json:"id"`                    // Order ID
	Item                string `json:"item"`                  // Order item
	ReceiverId          int    `json:"receiver_id"`           // Receiver ID
	Status              string `json:"status"`                // Order status
	TransactionId       int    `json:"transaction_id"`        // Transaction ID
	UserId              int    `json:"user_id"`               // User ID
}

// OrdersSubscription type represents `orders_subscription` API object
type OrdersSubscription struct {
	CancelReason    string `json:"cancel_reason"`     // Cancel reason
	CreateTime      int    `json:"create_time"`       // Date of creation in Unixtime
	Id              int    `json:"id"`                // Subscription ID
	ItemId          string `json:"item_id"`           // Subscription order item
	NextBillTime    int    `json:"next_bill_time"`    // Date of next bill in Unixtime
	PendingCancel   bool   `json:"pending_cancel"`    // Pending cancel state
	Period          int    `json:"period"`            // Subscription period
	PeriodStartTime int    `json:"period_start_time"` // Date of last period start in Unixtime
	Price           int    `json:"price"`             // Subscription price
	Status          string `json:"status"`            // Subscription status
	TestMode        bool   `json:"test_mode"`         // Is test subscription
	TrialExpireTime int    `json:"trial_expire_time"` // Date of trial expire in Unixtime
	UpdateTime      int    `json:"update_time"`       // Date of last change in Unixtime
}
