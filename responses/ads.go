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

// AdsAddofficeuser type represents `ads_addOfficeUsers_response` API response object
type AdsAddofficeuser bool // true if success

// AdsChecklink type represents `ads_checkLink_response` API response object
type AdsChecklink *objects.AdsLinkStatus

// AdsCreatead type represents `ads_createAds_response` API response object
type AdsCreatead int

// AdsCreatecampaig type represents `ads_createCampaigns_response` API response object
type AdsCreatecampaig int

// AdsCreateclient type represents `ads_createClients_response` API response object
type AdsCreateclient int

// AdsCreatetargetgrou type represents `ads_createTargetGroup_response` API response object
type AdsCreatetargetgrou struct {
	Id    int    `json:"id"`    // Group ID
	Pixel string `json:"pixel"` // Pixel code
}

// AdsDeletead type represents `ads_deleteAds_response` API response object
type AdsDeletead int

// AdsDeletecampaig type represents `ads_deleteCampaigns_response` API response object
type AdsDeletecampaig int // 0 if success

// AdsDeleteclient type represents `ads_deleteClients_response` API response object
type AdsDeleteclient int // 0 if sucess

// AdsGetaccount type represents `ads_getAccounts_response` API response object
type AdsGetaccount *objects.AdsAccount

// AdsGetadslayout type represents `ads_getAdsLayout_response` API response object
type AdsGetadslayout *objects.AdsAdLayout

// AdsGetadstargeting type represents `ads_getAdsTargeting_response` API response object
type AdsGetadstargeting *objects.AdsTargSettings

// AdsGetad type represents `ads_getAds_response` API response object
type AdsGetad *objects.AdsAd

// AdsGetbudget type represents `ads_getBudget_response` API response object
type AdsGetbudget int // Budget

// AdsGetcampaig type represents `ads_getCampaigns_response` API response object
type AdsGetcampaig *objects.AdsCampaign

// AdsGetcategori type represents `ads_getCategories_response` API response object
type AdsGetcategori struct {
	V1 []*objects.AdsCategory `json:"v1"` // Old categories
	V2 []*objects.AdsCategory `json:"v2"` // Actual categories
}

// AdsGetclient type represents `ads_getClients_response` API response object
type AdsGetclient *objects.AdsClient

// AdsGetdemographic type represents `ads_getDemographics_response` API response object
type AdsGetdemographic *objects.AdsDemoStats

// AdsGetfloodstat type represents `ads_getFloodStats_response` API response object
type AdsGetfloodstat *objects.AdsFloodStats

// AdsGetofficeuser type represents `ads_getOfficeUsers_response` API response object
type AdsGetofficeuser *objects.AdsUsers

// AdsGetpostsreach type represents `ads_getPostsReach_response` API response object
type AdsGetpostsreach *objects.AdsPromotedPostReach

// AdsGetrejectionrea type represents `ads_getRejectionReason_response` API response object
type AdsGetrejectionrea *objects.AdsRejectReason

// AdsGetstatistic type represents `ads_getStatistics_response` API response object
type AdsGetstatistic *objects.AdsStats

// AdsGetsuggestionsCiti type represents `ads_getSuggestions_cities_response` API response object
type AdsGetsuggestionsCiti *objects.AdsTargSuggestionsCities

// AdsGetsuggestionsRegi type represents `ads_getSuggestions_regions_response` API response object
type AdsGetsuggestionsRegi *objects.AdsTargSuggestionsRegions

// AdsGetsuggesti type represents `ads_getSuggestions_response` API response object
type AdsGetsuggesti *objects.AdsTargSuggestions

// AdsGetsuggestionsSchool type represents `ads_getSuggestions_schools_response` API response object
type AdsGetsuggestionsSchool *objects.AdsTargSuggestionsSchools

// AdsGettargetgrou type represents `ads_getTargetGroups_response` API response object
type AdsGettargetgrou *objects.AdsTargetGroup

// AdsGettargetingstat type represents `ads_getTargetingStats_response` API response object
type AdsGettargetingstat *objects.AdsTargStats

// AdsGetuploadurl type represents `ads_getUploadURL_response` API response object
type AdsGetuploadurl string // Photo upload URL

// AdsGetvideouploadurl type represents `ads_getVideoUploadURL_response` API response object
type AdsGetvideouploadurl string // Video upload URL

// AdsImporttargetcontact type represents `ads_importTargetContacts_response` API response object
type AdsImporttargetcontact int // Imported contacts number

// AdsRemoveofficeuser type represents `ads_removeOfficeUsers_response` API response object
type AdsRemoveofficeuser bool // true if success

// AdsUpdatead type represents `ads_updateAds_response` API response object
type AdsUpdatead int

// AdsUpdatecampaig type represents `ads_updateCampaigns_response` API response object
type AdsUpdatecampaig int // Campaign ID

// AdsUpdateclient type represents `ads_updateClients_response` API response object
type AdsUpdateclient int // Client ID
