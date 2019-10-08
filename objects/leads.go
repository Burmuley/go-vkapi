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
// `leads` group of objects
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

// LeadsChecked type represents `leads_checked` API object
type LeadsChecked struct {
	Reason    string             `json:"reason"` // Reason why user can't start the lead
	Result    LeadsCheckedResult `json:"result"`
	Sid       string             `json:"sid"`        // Session ID
	StartLink string             `json:"start_link"` // URL user should open to start the lead
}

// LeadsStart type represents `leads_start` API object
type LeadsStart struct {
	TestMode BaseBoolInt `json:"test_mode"` // Information whether test mode is enabled
	VkSid    string      `json:"vk_sid"`    // Session data
}

// LeadsCheckedResult type represents `leads_checked_result` API object
type LeadsCheckedResult string // Information whether user can start the lead

// LeadsEntry type represents `leads_entry` API object
type LeadsEntry struct {
	Aid       int         `json:"aid"`        // Application ID
	Comment   string      `json:"comment"`    // Comment text
	Date      int         `json:"date"`       // Date when the action has been started in Unixtime
	Sid       string      `json:"sid"`        // Session string ID
	StartDate int         `json:"start_date"` // Start date in Unixtime (for status=2)
	Status    int         `json:"status"`     // Action type
	TestMode  BaseBoolInt `json:"test_mode"`  // Information whether test mode is enabled
	Uid       int         `json:"uid"`        // User ID
}

// LeadsComplete type represents `leads_complete` API object
type LeadsComplete struct {
	Cost     int         `json:"cost"`  // Offer cost
	Limit    int         `json:"limit"` // Offer limit
	Spent    int         `json:"spent"` // Amount of spent votes
	Success  BaseOk      `json:"success"`
	TestMode BaseBoolInt `json:"test_mode"` // Information whether test mode is enabled
}

// LeadsLead type represents `leads_lead` API object
type LeadsLead struct {
	Completed   int           `json:"completed"` // Completed offers number
	Cost        int           `json:"cost"`      // Offer cost
	Days        LeadsLeadDays `json:"days"`
	Impressions int           `json:"impressions"` // Impressions number
	Limit       int           `json:"limit"`       // Lead limit
	Spent       int           `json:"spent"`       // Amount of spent votes
	Started     int           `json:"started"`     // Started offers number
}

// LeadsLeadDays type represents `leads_lead_days` API object
type LeadsLeadDays struct {
	Completed   int `json:"completed"`   // Completed offers number
	Impressions int `json:"impressions"` // Impressions number
	Spent       int `json:"spent"`       // Amount of spent votes
	Started     int `json:"started"`     // Started offers number
}
