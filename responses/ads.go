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
// `ads` group of responses
/////////////////////////////////////////////////////////////

// AdsAddofficeusers type represents `ads_addOfficeUsers_response` API response object
type AdsAddofficeusers bool // true if success

// AdsChecklink type represents `ads_checkLink_response` API response object
type AdsChecklink objects.AdsLinkStatus

// AdsCreateads type represents `ads_createAds_response` API response object
type AdsCreateads int

// AdsCreatecampaigns type represents `ads_createCampaigns_response` API response object
type AdsCreatecampaigns int

// AdsCreateclients type represents `ads_createClients_response` API response object
type AdsCreateclients int

// AdsCreatetargetgroup type represents `ads_createTargetGroup_response` API response object
type AdsCreatetargetgroup struct {
	Id    int    `json:"id"`    // Group ID
	Pixel string `json:"pixel"` // Pixel code
}

// AdsDeleteads type represents `ads_deleteAds_response` API response object
type AdsDeleteads int

// AdsDeletecampaigns type represents `ads_deleteCampaigns_response` API response object
type AdsDeletecampaigns int // 0 if success

// AdsDeleteclients type represents `ads_deleteClients_response` API response object
type AdsDeleteclients int // 0 if sucess

// AdsGetaccounts type represents `ads_getAccounts_response` API response object
type AdsGetaccounts objects.AdsAccount

// AdsGetadslayout type represents `ads_getAdsLayout_response` API response object
type AdsGetadslayout objects.AdsAdLayout

// AdsGetadstargeting type represents `ads_getAdsTargeting_response` API response object
type AdsGetadstargeting objects.AdsTargSettings

// AdsGetads type represents `ads_getAds_response` API response object
type AdsGetads objects.AdsAd

// AdsGetbudget type represents `ads_getBudget_response` API response object
type AdsGetbudget int // Budget

// AdsGetcampaigns type represents `ads_getCampaigns_response` API response object
type AdsGetcampaigns objects.AdsCampaign

// AdsGetcategories type represents `ads_getCategories_response` API response object
type AdsGetcategories struct {
	V1 []objects.AdsCategory `json:"v1"` // Old categories
	V2 []objects.AdsCategory `json:"v2"` // Actual categories
}

// AdsGetclients type represents `ads_getClients_response` API response object
type AdsGetclients objects.AdsClient

// AdsGetdemographics type represents `ads_getDemographics_response` API response object
type AdsGetdemographics objects.AdsDemoStats

// AdsGetfloodstats type represents `ads_getFloodStats_response` API response object
type AdsGetfloodstats objects.AdsFloodStats

// AdsGetofficeusers type represents `ads_getOfficeUsers_response` API response object
type AdsGetofficeusers objects.AdsUsers

// AdsGetpostsreach type represents `ads_getPostsReach_response` API response object
type AdsGetpostsreach objects.AdsPromotedPostReach

// AdsGetrejectionreason type represents `ads_getRejectionReason_response` API response object
type AdsGetrejectionreason objects.AdsRejectReason

// AdsGetstatistics type represents `ads_getStatistics_response` API response object
type AdsGetstatistics objects.AdsStats

// AdsGetsuggestionsCities type represents `ads_getSuggestions_cities_response` API response object
type AdsGetsuggestionsCities objects.AdsTargSuggestionsCities

// AdsGetsuggestionsRegions type represents `ads_getSuggestions_regions_response` API response object
type AdsGetsuggestionsRegions objects.AdsTargSuggestionsRegions

// AdsGetsuggestions type represents `ads_getSuggestions_response` API response object
type AdsGetsuggestions objects.AdsTargSuggestions

// AdsGetsuggestionsSchools type represents `ads_getSuggestions_schools_response` API response object
type AdsGetsuggestionsSchools objects.AdsTargSuggestionsSchools

// AdsGettargetgroups type represents `ads_getTargetGroups_response` API response object
type AdsGettargetgroups objects.AdsTargetGroup

// AdsGettargetingstats type represents `ads_getTargetingStats_response` API response object
type AdsGettargetingstats objects.AdsTargStats

// AdsGetuploadurl type represents `ads_getUploadURL_response` API response object
type AdsGetuploadurl string // Photo upload URL

// AdsGetvideouploadurl type represents `ads_getVideoUploadURL_response` API response object
type AdsGetvideouploadurl string // Video upload URL

// AdsImporttargetcontacts type represents `ads_importTargetContacts_response` API response object
type AdsImporttargetcontacts int // Imported contacts number

// AdsRemoveofficeusers type represents `ads_removeOfficeUsers_response` API response object
type AdsRemoveofficeusers bool // true if success

// AdsUpdateads type represents `ads_updateAds_response` API response object
type AdsUpdateads int

// AdsUpdatecampaigns type represents `ads_updateCampaigns_response` API response object
type AdsUpdatecampaigns int // Campaign ID

// AdsUpdateclients type represents `ads_updateClients_response` API response object
type AdsUpdateclients int // Client ID
