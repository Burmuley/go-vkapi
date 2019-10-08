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
// `ads` group of responses
/////////////////////////////////////////////////////////////

// AdsImportTargetContacts type represents `ads_importTargetContacts_response` API response object
type AdsImportTargetContacts int // Imported contacts number

// AdsGetTargetGroups type represents `ads_getTargetGroups_response` API response object
type AdsGetTargetGroups objects.AdsTargetGroup

// AdsGetSuggestionsCities type represents `ads_getSuggestions_cities_response` API response object
type AdsGetSuggestionsCities objects.AdsTargSuggestionsCities

// AdsRemoveOfficeUsers type represents `ads_removeOfficeUsers_response` API response object

// AdsCreateAds type represents `ads_createAds_response` API response object
type AdsCreateAds int

// AdsGetClients type represents `ads_getClients_response` API response object
type AdsGetClients objects.AdsClient

// AdsGetCategories type represents `ads_getCategories_response` API response object
type AdsGetCategories struct {
	V1 []objects.AdsCategory `json:"v1"` // Old categories
	V2 []objects.AdsCategory `json:"v2"` // Actual categories
}

// AdsUpdateAds type represents `ads_updateAds_response` API response object
type AdsUpdateAds int

// AdsGetBudget type represents `ads_getBudget_response` API response object
type AdsGetBudget int // Budget

// AdsGetRejectionReason type represents `ads_getRejectionReason_response` API response object
type AdsGetRejectionReason objects.AdsRejectReason

// AdsGetDemographics type represents `ads_getDemographics_response` API response object
type AdsGetDemographics objects.AdsDemoStats

// AdsGetVideoUploadURL type represents `ads_getVideoUploadURL_response` API response object
type AdsGetVideoUploadURL string // Video upload URL

// AdsGetOfficeUsers type represents `ads_getOfficeUsers_response` API response object
type AdsGetOfficeUsers objects.AdsUsers

// AdsGetAdsTargeting type represents `ads_getAdsTargeting_response` API response object
type AdsGetAdsTargeting objects.AdsTargSettings

// AdsGetSuggestionsSchools type represents `ads_getSuggestions_schools_response` API response object
type AdsGetSuggestionsSchools objects.AdsTargSuggestionsSchools

// AdsGetTargetingStats type represents `ads_getTargetingStats_response` API response object
type AdsGetTargetingStats objects.AdsTargStats

// AdsCreateClients type represents `ads_createClients_response` API response object
type AdsCreateClients int

// AdsUpdateClients type represents `ads_updateClients_response` API response object
type AdsUpdateClients int // Client ID

// AdsCreateTargetGroup type represents `ads_createTargetGroup_response` API response object
type AdsCreateTargetGroup struct {
	Id    int    `json:"id"`    // Group ID
	Pixel string `json:"pixel"` // Pixel code
}

// AdsGetAds type represents `ads_getAds_response` API response object
type AdsGetAds objects.AdsAd

// AdsGetStatistics type represents `ads_getStatistics_response` API response object
type AdsGetStatistics objects.AdsStats

// AdsGetSuggestions type represents `ads_getSuggestions_response` API response object
type AdsGetSuggestions objects.AdsTargSuggestions

// AdsDeleteAds type represents `ads_deleteAds_response` API response object
type AdsDeleteAds int

// AdsCheckLink type represents `ads_checkLink_response` API response object
type AdsCheckLink objects.AdsLinkStatus

// AdsGetAdsLayout type represents `ads_getAdsLayout_response` API response object
type AdsGetAdsLayout objects.AdsAdLayout

// AdsCreateCampaigns type represents `ads_createCampaigns_response` API response object
type AdsCreateCampaigns int

// AdsGetAccounts type represents `ads_getAccounts_response` API response object
type AdsGetAccounts objects.AdsAccount

// AdsDeleteCampaigns type represents `ads_deleteCampaigns_response` API response object
type AdsDeleteCampaigns int // 0 if success

// AdsGetCampaigns type represents `ads_getCampaigns_response` API response object
type AdsGetCampaigns objects.AdsCampaign

// AdsGetFloodStats type represents `ads_getFloodStats_response` API response object
type AdsGetFloodStats objects.AdsFloodStats

// AdsGetPostsReach type represents `ads_getPostsReach_response` API response object
type AdsGetPostsReach objects.AdsPromotedPostReach

// AdsUpdateCampaigns type represents `ads_updateCampaigns_response` API response object
type AdsUpdateCampaigns int // Campaign ID

// AdsAddOfficeUsers type represents `ads_addOfficeUsers_response` API response object

// AdsGetUploadURL type represents `ads_getUploadURL_response` API response object
type AdsGetUploadURL string // Photo upload URL

// AdsDeleteClients type represents `ads_deleteClients_response` API response object
type AdsDeleteClients int // 0 if sucess

// AdsGetSuggestionsRegions type represents `ads_getSuggestions_regions_response` API response object
type AdsGetSuggestionsRegions objects.AdsTargSuggestionsRegions
