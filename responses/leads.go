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

import (
	"gitlab.com/Burmuley/go-vkapi/objects"
)

/////////////////////////////////////////////////////////////
// `leads` group of responses
/////////////////////////////////////////////////////////////

// LeadsCheckuser type represents `leads_checkUser_response` API response object
type LeadsCheckuser objects.LeadsChecked

// LeadsComplete type represents `leads_complete_response` API response object
type LeadsComplete objects.LeadsComplete

// LeadsGetstats type represents `leads_getStats_response` API response object
type LeadsGetstats objects.LeadsLead

// LeadsGetusers type represents `leads_getUsers_response` API response object
type LeadsGetusers objects.LeadsEntry

// LeadsMetrichit type represents `leads_metricHit_response` API response object
type LeadsMetrichit struct {
	RedirectLink string `json:"redirect_link"` // Redirect link
	Result       bool   `json:"result"`        // Information whether request has been processed successfully
}

// LeadsStart type represents `leads_start_response` API response object
type LeadsStart objects.LeadsStart
