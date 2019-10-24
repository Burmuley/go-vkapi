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

type Ads struct {
	*VKApi
}

/////////////////////////////////////////////////////////////
// `Ads` methods
/////////////////////////////////////////////////////////////

// Addofficeusers - Adds managers and/or supervisors to advertising account.
// Parameters:
//   * accountId - Advertising account ID.
//   * data - Serialized JSON array of objects that describe added managers. Description of 'user_specification' objects see below.
func (a Ads) Addofficeusers(accountId int, data string) (resp responses.AdsAddofficeuser, err error) {
	params := map[string]interface{}{}

	params["account_id"] = accountId

	params["data"] = data

	err = a.SendObjRequest("ads.addOfficeUsers", params, &resp)

	return
}

// Checklink - Allows to check the ad link.
// Parameters:
//   * accountId - Advertising account ID.
//   * linkType - Object type: *'community' — community,, *'post' — community post,, *'application' — VK application,, *'video' — video,, *'site' — external site.
//   * linkUrl - Object URL.
//   * campaignId - Campaign ID
func (a Ads) Checklink(accountId int, linkType string, linkUrl string, campaignId int) (resp responses.AdsChecklink, err error) {
	params := map[string]interface{}{}

	params["account_id"] = accountId

	params["link_type"] = linkType

	params["link_url"] = linkUrl

	if campaignId > 0 {
		params["campaign_id"] = campaignId
	}

	err = a.SendObjRequest("ads.checkLink", params, &resp)

	return
}

// Createads - Creates ads.
// Parameters:
//   * accountId - Advertising account ID.
//   * data - Serialized JSON array of objects that describe created ads. Description of 'ad_specification' objects see below.
func (a Ads) Createads(accountId int, data string) (resp responses.AdsCreatead, err error) {
	params := map[string]interface{}{}

	params["account_id"] = accountId

	params["data"] = data

	err = a.SendObjRequest("ads.createAds", params, &resp)

	return
}

// Createcampaigns - Creates advertising campaigns.
// Parameters:
//   * accountId - Advertising account ID.
//   * data - Serialized JSON array of objects that describe created campaigns. Description of 'campaign_specification' objects see below.
func (a Ads) Createcampaigns(accountId int, data string) (resp responses.AdsCreatecampaig, err error) {
	params := map[string]interface{}{}

	params["account_id"] = accountId

	params["data"] = data

	err = a.SendObjRequest("ads.createCampaigns", params, &resp)

	return
}

// Createclients - Creates clients of an advertising agency.
// Parameters:
//   * accountId - Advertising account ID.
//   * data - Serialized JSON array of objects that describe created campaigns. Description of 'client_specification' objects see below.
func (a Ads) Createclients(accountId int, data string) (resp responses.AdsCreateclient, err error) {
	params := map[string]interface{}{}

	params["account_id"] = accountId

	params["data"] = data

	err = a.SendObjRequest("ads.createClients", params, &resp)

	return
}

// Createtargetgroup - Creates a group to re-target ads for users who visited advertiser's site (viewed information about the product, registered, etc.).
// Parameters:
//   * accountId - Advertising account ID.
//   * clientId - 'Only for advertising agencies.', ID of the client with the advertising account where the group will be created.
//   * name - Name of the target group — a string up to 64 characters long.
//   * lifetime - 'For groups with auditory created with pixel code only.', , Number of days after that users will be automatically removed from the group.
//   * targetPixelId - NO DESCRIPTION IN JSON SCHEMA
//   * targetPixelRules - NO DESCRIPTION IN JSON SCHEMA
func (a Ads) Createtargetgroup(accountId int, clientId int, name string, lifetime int, targetPixelId int, targetPixelRules string) (resp responses.AdsCreatetargetgrou, err error) {
	params := map[string]interface{}{}

	params["account_id"] = accountId

	if clientId > 0 {
		params["client_id"] = clientId
	}

	params["name"] = name

	if lifetime > 0 {
		params["lifetime"] = lifetime
	}

	if targetPixelId > 0 {
		params["target_pixel_id"] = targetPixelId
	}

	if targetPixelRules != "" {
		params["target_pixel_rules"] = targetPixelRules
	}

	err = a.SendObjRequest("ads.createTargetGroup", params, &resp)

	return
}

// Deleteads - Archives ads.
// Parameters:
//   * accountId - Advertising account ID.
//   * ids - Serialized JSON array with ad IDs.
func (a Ads) Deleteads(accountId int, ids string) (resp responses.AdsDeletead, err error) {
	params := map[string]interface{}{}

	params["account_id"] = accountId

	params["ids"] = ids

	err = a.SendObjRequest("ads.deleteAds", params, &resp)

	return
}

// Deletecampaigns - Archives advertising campaigns.
// Parameters:
//   * accountId - Advertising account ID.
//   * ids - Serialized JSON array with IDs of deleted campaigns.
func (a Ads) Deletecampaigns(accountId int, ids string) (resp responses.AdsDeletecampaig, err error) {
	params := map[string]interface{}{}

	params["account_id"] = accountId

	params["ids"] = ids

	err = a.SendObjRequest("ads.deleteCampaigns", params, &resp)

	return
}

// Deleteclients - Archives clients of an advertising agency.
// Parameters:
//   * accountId - Advertising account ID.
//   * ids - Serialized JSON array with IDs of deleted clients.
func (a Ads) Deleteclients(accountId int, ids string) (resp responses.AdsDeleteclient, err error) {
	params := map[string]interface{}{}

	params["account_id"] = accountId

	params["ids"] = ids

	err = a.SendObjRequest("ads.deleteClients", params, &resp)

	return
}

// Deletetargetgroup - Deletes a retarget group.
// Parameters:
//   * accountId - Advertising account ID.
//   * clientId - 'Only for advertising agencies.' , ID of the client with the advertising account where the group will be created.
//   * targetGroupId - Group ID.
func (a Ads) Deletetargetgroup(accountId int, clientId int, targetGroupId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["account_id"] = accountId

	if clientId > 0 {
		params["client_id"] = clientId
	}

	params["target_group_id"] = targetGroupId

	err = a.SendObjRequest("ads.deleteTargetGroup", params, &resp)

	return
}

// Getaccounts - Returns a list of advertising accounts.
func (a Ads) Getaccounts() (resp responses.AdsGetaccount, err error) {
	params := map[string]interface{}{}

	err = a.SendObjRequest("ads.getAccounts", params, &resp)

	return
}

// Getads - Returns number of ads.
// Parameters:
//   * accountId - Advertising account ID.
//   * adIds - Filter by ads. Serialized JSON array with ad IDs. If the parameter is null, all ads will be shown.
//   * campaignIds - Filter by advertising campaigns. Serialized JSON array with campaign IDs. If the parameter is null, ads of all campaigns will be shown.
//   * clientId - 'Available and required for advertising agencies.' ID of the client ads are retrieved from.
//   * includeDeleted - Flag that specifies whether archived ads shall be shown: *0 — show only active ads,, *1 — show all ads.
//   * limit - Limit of number of returned ads. Used only if ad_ids parameter is null, and 'campaign_ids' parameter contains ID of only one campaign.
//   * offset - Offset. Used in the same cases as 'limit' parameter.
func (a Ads) Getads(accountId int, adIds string, campaignIds string, clientId int, includeDeleted bool, limit int, offset int) (resp responses.AdsGetad, err error) {
	params := map[string]interface{}{}

	params["account_id"] = accountId

	if adIds != "" {
		params["ad_ids"] = adIds
	}

	if campaignIds != "" {
		params["campaign_ids"] = campaignIds
	}

	if clientId > 0 {
		params["client_id"] = clientId
	}

	params["include_deleted"] = includeDeleted

	if limit > 0 {
		params["limit"] = limit
	}

	if offset > 0 {
		params["offset"] = offset
	}

	err = a.SendObjRequest("ads.getAds", params, &resp)

	return
}

// Getadslayout - Returns descriptions of ad layouts.
// Parameters:
//   * accountId - Advertising account ID.
//   * adIds - Filter by ads. Serialized JSON array with ad IDs. If the parameter is null, all ads will be shown.
//   * campaignIds - Filter by advertising campaigns. Serialized JSON array with campaign IDs. If the parameter is null, ads of all campaigns will be shown.
//   * clientId - 'For advertising agencies.' ID of the client ads are retrieved from.
//   * includeDeleted - Flag that specifies whether archived ads shall be shown. *0 — show only active ads,, *1 — show all ads.
//   * limit - Limit of number of returned ads. Used only if 'ad_ids' parameter is null, and 'campaign_ids' parameter contains ID of only one campaign.
//   * offset - Offset. Used in the same cases as 'limit' parameter.
func (a Ads) Getadslayout(accountId int, adIds string, campaignIds string, clientId int, includeDeleted bool, limit int, offset int) (resp responses.AdsGetadslayout, err error) {
	params := map[string]interface{}{}

	params["account_id"] = accountId

	if adIds != "" {
		params["ad_ids"] = adIds
	}

	if campaignIds != "" {
		params["campaign_ids"] = campaignIds
	}

	if clientId > 0 {
		params["client_id"] = clientId
	}

	params["include_deleted"] = includeDeleted

	if limit > 0 {
		params["limit"] = limit
	}

	if offset > 0 {
		params["offset"] = offset
	}

	err = a.SendObjRequest("ads.getAdsLayout", params, &resp)

	return
}

// Getadstargeting - Returns ad targeting parameters.
// Parameters:
//   * accountId - Advertising account ID.
//   * adIds - Filter by ads. Serialized JSON array with ad IDs. If the parameter is null, all ads will be shown.
//   * campaignIds - Filter by advertising campaigns. Serialized JSON array with campaign IDs. If the parameter is null, ads of all campaigns will be shown.
//   * clientId - 'For advertising agencies.' ID of the client ads are retrieved from.
//   * includeDeleted - flag that specifies whether archived ads shall be shown: *0 — show only active ads,, *1 — show all ads.
//   * limit - Limit of number of returned ads. Used only if 'ad_ids' parameter is null, and 'campaign_ids' parameter contains ID of only one campaign.
//   * offset - Offset needed to return a specific subset of results.
func (a Ads) Getadstargeting(accountId int, adIds string, campaignIds string, clientId int, includeDeleted bool, limit int, offset int) (resp responses.AdsGetadstargeting, err error) {
	params := map[string]interface{}{}

	params["account_id"] = accountId

	if adIds != "" {
		params["ad_ids"] = adIds
	}

	if campaignIds != "" {
		params["campaign_ids"] = campaignIds
	}

	if clientId > 0 {
		params["client_id"] = clientId
	}

	params["include_deleted"] = includeDeleted

	if limit > 0 {
		params["limit"] = limit
	}

	if offset > 0 {
		params["offset"] = offset
	}

	err = a.SendObjRequest("ads.getAdsTargeting", params, &resp)

	return
}

// Getbudget - Returns current budget of the advertising account.
// Parameters:
//   * accountId - Advertising account ID.
func (a Ads) Getbudget(accountId int) (resp responses.AdsGetbudget, err error) {
	params := map[string]interface{}{}

	params["account_id"] = accountId

	err = a.SendObjRequest("ads.getBudget", params, &resp)

	return
}

// Getcampaigns - Returns a list of campaigns in an advertising account.
// Parameters:
//   * accountId - Advertising account ID.
//   * clientId - 'For advertising agencies'. ID of the client advertising campaigns are retrieved from.
//   * includeDeleted - Flag that specifies whether archived ads shall be shown. *0 — show only active campaigns,, *1 — show all campaigns.
//   * campaignIds - Filter of advertising campaigns to show. Serialized JSON array with campaign IDs. Only campaigns that exist in 'campaign_ids' and belong to the specified advertising account will be shown. If the parameter is null, all campaigns will be shown.
func (a Ads) Getcampaigns(accountId int, clientId int, includeDeleted bool, campaignIds string) (resp responses.AdsGetcampaig, err error) {
	params := map[string]interface{}{}

	params["account_id"] = accountId

	if clientId > 0 {
		params["client_id"] = clientId
	}

	params["include_deleted"] = includeDeleted

	if campaignIds != "" {
		params["campaign_ids"] = campaignIds
	}

	err = a.SendObjRequest("ads.getCampaigns", params, &resp)

	return
}

// Getcategories - Returns a list of possible ad categories.
// Parameters:
//   * lang - Language. The full list of supported languages is [vk.com/dev/api_requests|here].
func (a Ads) Getcategories(lang string) (resp responses.AdsGetcategori, err error) {
	params := map[string]interface{}{}

	if lang != "" {
		params["lang"] = lang
	}

	err = a.SendObjRequest("ads.getCategories", params, &resp)

	return
}

// Getclients - Returns a list of advertising agency's clients.
// Parameters:
//   * accountId - Advertising account ID.
func (a Ads) Getclients(accountId int) (resp responses.AdsGetclient, err error) {
	params := map[string]interface{}{}

	params["account_id"] = accountId

	err = a.SendObjRequest("ads.getClients", params, &resp)

	return
}

// Getdemographics - Returns demographics for ads or campaigns.
// Parameters:
//   * accountId - Advertising account ID.
//   * idsType - Type of requested objects listed in 'ids' parameter: *ad — ads,, *campaign — campaigns.
//   * ids - IDs requested ads or campaigns, separated with a comma, depending on the value set in 'ids_type'. Maximum 2000 objects.
//   * period - Data grouping by dates: *day — statistics by days,, *month — statistics by months,, *overall — overall statistics. 'date_from' and 'date_to' parameters set temporary limits.
//   * dateFrom - Date to show statistics from. For different value of 'period' different date format is used: *day: YYYY-MM-DD, example: 2011-09-27 — September 27, 2011, **0 — day it was created on,, *month: YYYY-MM, example: 2011-09 — September 2011, **0 — month it was created in,, *overall: 0.
//   * dateTo - Date to show statistics to. For different value of 'period' different date format is used: *day: YYYY-MM-DD, example: 2011-09-27 — September 27, 2011, **0 — current day,, *month: YYYY-MM, example: 2011-09 — September 2011, **0 — current month,, *overall: 0.
func (a Ads) Getdemographics(accountId int, idsType string, ids string, period string, dateFrom string, dateTo string) (resp responses.AdsGetdemographic, err error) {
	params := map[string]interface{}{}

	params["account_id"] = accountId

	params["ids_type"] = idsType

	params["ids"] = ids

	params["period"] = period

	params["date_from"] = dateFrom

	params["date_to"] = dateTo

	err = a.SendObjRequest("ads.getDemographics", params, &resp)

	return
}

// Getfloodstats - Returns information about current state of a counter — number of remaining runs of methods and time to the next counter nulling in seconds.
// Parameters:
//   * accountId - Advertising account ID.
func (a Ads) Getfloodstats(accountId int) (resp responses.AdsGetfloodstat, err error) {
	params := map[string]interface{}{}

	params["account_id"] = accountId

	err = a.SendObjRequest("ads.getFloodStats", params, &resp)

	return
}

// Getofficeusers - Returns a list of managers and supervisors of advertising account.
// Parameters:
//   * accountId - Advertising account ID.
func (a Ads) Getofficeusers(accountId int) (resp responses.AdsGetofficeuser, err error) {
	params := map[string]interface{}{}

	params["account_id"] = accountId

	err = a.SendObjRequest("ads.getOfficeUsers", params, &resp)

	return
}

// Getpostsreach - Returns detailed statistics of promoted posts reach from campaigns and ads.
// Parameters:
//   * accountId - Advertising account ID.
//   * idsType - Type of requested objects listed in 'ids' parameter: *ad — ads,, *campaign — campaigns.
//   * ids - IDs requested ads or campaigns, separated with a comma, depending on the value set in 'ids_type'. Maximum 100 objects.
func (a Ads) Getpostsreach(accountId int, idsType string, ids string) (resp responses.AdsGetpostsreach, err error) {
	params := map[string]interface{}{}

	params["account_id"] = accountId

	params["ids_type"] = idsType

	params["ids"] = ids

	err = a.SendObjRequest("ads.getPostsReach", params, &resp)

	return
}

// Getrejectionreason - Returns a reason of ad rejection for pre-moderation.
// Parameters:
//   * accountId - Advertising account ID.
//   * adId - Ad ID.
func (a Ads) Getrejectionreason(accountId int, adId int) (resp responses.AdsGetrejectionrea, err error) {
	params := map[string]interface{}{}

	params["account_id"] = accountId

	params["ad_id"] = adId

	err = a.SendObjRequest("ads.getRejectionReason", params, &resp)

	return
}

// Getstatistics - Returns statistics of performance indicators for ads, campaigns, clients or the whole account.
// Parameters:
//   * accountId - Advertising account ID.
//   * idsType - Type of requested objects listed in 'ids' parameter: *ad — ads,, *campaign — campaigns,, *client — clients,, *office — account.
//   * ids - IDs requested ads, campaigns, clients or account, separated with a comma, depending on the value set in 'ids_type'. Maximum 2000 objects.
//   * period - Data grouping by dates: *day — statistics by days,, *month — statistics by months,, *overall — overall statistics. 'date_from' and 'date_to' parameters set temporary limits.
//   * dateFrom - Date to show statistics from. For different value of 'period' different date format is used: *day: YYYY-MM-DD, example: 2011-09-27 — September 27, 2011, **0 — day it was created on,, *month: YYYY-MM, example: 2011-09 — September 2011, **0 — month it was created in,, *overall: 0.
//   * dateTo - Date to show statistics to. For different value of 'period' different date format is used: *day: YYYY-MM-DD, example: 2011-09-27 — September 27, 2011, **0 — current day,, *month: YYYY-MM, example: 2011-09 — September 2011, **0 — current month,, *overall: 0.
func (a Ads) Getstatistics(accountId int, idsType string, ids string, period string, dateFrom string, dateTo string) (resp responses.AdsGetstatistic, err error) {
	params := map[string]interface{}{}

	params["account_id"] = accountId

	params["ids_type"] = idsType

	params["ids"] = ids

	params["period"] = period

	params["date_from"] = dateFrom

	params["date_to"] = dateTo

	err = a.SendObjRequest("ads.getStatistics", params, &resp)

	return
}

// Getsuggestions - Returns a set of auto-suggestions for various targeting parameters.
// Parameters:
//   * section - Section, suggestions are retrieved in. Available values: *countries — request of a list of countries. If q is not set or blank, a short list of countries is shown. Otherwise, a full list of countries is shown. *regions — requested list of regions. 'country' parameter is required. *cities — requested list of cities. 'country' parameter is required. *districts — requested list of districts. 'cities' parameter is required. *stations — requested list of subway stations. 'cities' parameter is required. *streets — requested list of streets. 'cities' parameter is required. *schools — requested list of educational organizations. 'cities' parameter is required. *interests — requested list of interests. *positions — requested list of positions (professions). *group_types — requested list of group types. *religions — requested list of religious commitments. *browsers — requested list of browsers and mobile devices.
//   * ids - Objects IDs separated by commas. If the parameter is passed, 'q, country, cities' should not be passed.
//   * q - Filter-line of the request (for countries, regions, cities, streets, schools, interests, positions).
//   * country - ID of the country objects are searched in.
//   * cities - IDs of cities where objects are searched in, separated with a comma.
//   * lang - Language of the returned string values. Supported languages: *ru — Russian,, *ua — Ukrainian,, *en — English.
func (a Ads) Getsuggestions(section string, ids string, q string, country int, cities string, lang string) (resp responses.AdsGetsuggesti, err error) {
	params := map[string]interface{}{}

	params["section"] = section

	if ids != "" {
		params["ids"] = ids
	}

	if q != "" {
		params["q"] = q
	}

	if country > 0 {
		params["country"] = country
	}

	if cities != "" {
		params["cities"] = cities
	}

	if lang != "" {
		params["lang"] = lang
	}

	err = a.SendObjRequest("ads.getSuggestions", params, &resp)

	return
}

// Gettargetgroups - Returns a list of target groups.
// Parameters:
//   * accountId - Advertising account ID.
//   * clientId - 'Only for advertising agencies.', ID of the client with the advertising account where the group will be created.
//   * extended - '1' — to return pixel code.
func (a Ads) Gettargetgroups(accountId int, clientId int) (resp responses.AdsGettargetgrou, err error) {
	params := map[string]interface{}{}

	params["account_id"] = accountId

	if clientId > 0 {
		params["client_id"] = clientId
	}

	err = a.SendObjRequest("ads.getTargetGroups", params, &resp)

	return
}

// Gettargetingstats - Returns the size of targeting audience, and also recommended values for CPC and CPM.
// Parameters:
//   * accountId - Advertising account ID.
//   * clientId - NO DESCRIPTION IN JSON SCHEMA
//   * criteria - Serialized JSON object that describes targeting parameters. Description of 'criteria' object see below.
//   * adId - ID of an ad which targeting parameters shall be analyzed.
//   * adFormat - Ad format. Possible values: *'1' — image and text,, *'2' — big image,, *'3' — exclusive format,, *'4' — community, square image,, *'7' — special app format,, *'8' — special community format,, *'9' — post in community,, *'10' — app board.
//   * adPlatform - Platforms to use for ad showing. Possible values: (for 'ad_format' = '1'), *'0' — VK and partner sites,, *'1' — VK only. (for 'ad_format' = '9'), *'all' — all platforms,, *'desktop' — desktop version,, *'mobile' — mobile version and apps.
//   * adPlatformNoWall - NO DESCRIPTION IN JSON SCHEMA
//   * adPlatformNoAdNetwork - NO DESCRIPTION IN JSON SCHEMA
//   * linkUrl - URL for the advertised object.
//   * linkDomain - Domain of the advertised object.
func (a Ads) Gettargetingstats(accountId int, clientId int, criteria string, adId int, adFormat int, adPlatform string, adPlatformNoWall string, adPlatformNoAdNetwork string, linkUrl string, linkDomain string) (resp responses.AdsGettargetingstat, err error) {
	params := map[string]interface{}{}

	params["account_id"] = accountId

	if clientId > 0 {
		params["client_id"] = clientId
	}

	if criteria != "" {
		params["criteria"] = criteria
	}

	if adId > 0 {
		params["ad_id"] = adId
	}

	if adFormat > 0 {
		params["ad_format"] = adFormat
	}

	if adPlatform != "" {
		params["ad_platform"] = adPlatform
	}

	if adPlatformNoWall != "" {
		params["ad_platform_no_wall"] = adPlatformNoWall
	}

	if adPlatformNoAdNetwork != "" {
		params["ad_platform_no_ad_network"] = adPlatformNoAdNetwork
	}

	params["link_url"] = linkUrl

	if linkDomain != "" {
		params["link_domain"] = linkDomain
	}

	err = a.SendObjRequest("ads.getTargetingStats", params, &resp)

	return
}

// Getuploadurl - Returns URL to upload an ad photo to.
// Parameters:
//   * adFormat - Ad format: *1 — image and text,, *2 — big image,, *3 — exclusive format,, *4 — community, square image,, *7 — special app format.
//   * icon - NO DESCRIPTION IN JSON SCHEMA
func (a Ads) Getuploadurl(adFormat int, icon int) (resp responses.AdsGetuploadurl, err error) {
	params := map[string]interface{}{}

	params["ad_format"] = adFormat

	if icon > 0 {
		params["icon"] = icon
	}

	err = a.SendObjRequest("ads.getUploadURL", params, &resp)

	return
}

// Getvideouploadurl - Returns URL to upload an ad video to.
func (a Ads) Getvideouploadurl() (resp responses.AdsGetvideouploadurl, err error) {
	params := map[string]interface{}{}

	err = a.SendObjRequest("ads.getVideoUploadURL", params, &resp)

	return
}

// Importtargetcontacts - Imports a list of advertiser's contacts to count VK registered users against the target group.
// Parameters:
//   * accountId - Advertising account ID.
//   * clientId - 'Only for advertising agencies.' , ID of the client with the advertising account where the group will be created.
//   * targetGroupId - Target group ID.
//   * contacts - List of phone numbers, emails or user IDs separated with a comma.
func (a Ads) Importtargetcontacts(accountId int, clientId int, targetGroupId int, contacts string) (resp responses.AdsImporttargetcontact, err error) {
	params := map[string]interface{}{}

	params["account_id"] = accountId

	if clientId > 0 {
		params["client_id"] = clientId
	}

	params["target_group_id"] = targetGroupId

	params["contacts"] = contacts

	err = a.SendObjRequest("ads.importTargetContacts", params, &resp)

	return
}

// Removeofficeusers - Removes managers and/or supervisors from advertising account.
// Parameters:
//   * accountId - Advertising account ID.
//   * ids - Serialized JSON array with IDs of deleted managers.
func (a Ads) Removeofficeusers(accountId int, ids string) (resp responses.AdsRemoveofficeuser, err error) {
	params := map[string]interface{}{}

	params["account_id"] = accountId

	params["ids"] = ids

	err = a.SendObjRequest("ads.removeOfficeUsers", params, &resp)

	return
}

// Updateads - Edits ads.
// Parameters:
//   * accountId - Advertising account ID.
//   * data - Serialized JSON array of objects that describe changes in ads. Description of 'ad_edit_specification' objects see below.
func (a Ads) Updateads(accountId int, data string) (resp responses.AdsUpdatead, err error) {
	params := map[string]interface{}{}

	params["account_id"] = accountId

	params["data"] = data

	err = a.SendObjRequest("ads.updateAds", params, &resp)

	return
}

// Updatecampaigns - Edits advertising campaigns.
// Parameters:
//   * accountId - Advertising account ID.
//   * data - Serialized JSON array of objects that describe changes in campaigns. Description of 'campaign_mod' objects see below.
func (a Ads) Updatecampaigns(accountId int, data string) (resp responses.AdsUpdatecampaig, err error) {
	params := map[string]interface{}{}

	params["account_id"] = accountId

	params["data"] = data

	err = a.SendObjRequest("ads.updateCampaigns", params, &resp)

	return
}

// Updateclients - Edits clients of an advertising agency.
// Parameters:
//   * accountId - Advertising account ID.
//   * data - Serialized JSON array of objects that describe changes in clients. Description of 'client_mod' objects see below.
func (a Ads) Updateclients(accountId int, data string) (resp responses.AdsUpdateclient, err error) {
	params := map[string]interface{}{}

	params["account_id"] = accountId

	params["data"] = data

	err = a.SendObjRequest("ads.updateClients", params, &resp)

	return
}

// Updatetargetgroup - Edits a retarget group.
// Parameters:
//   * accountId - Advertising account ID.
//   * clientId - 'Only for advertising agencies.' , ID of the client with the advertising account where the group will be created.
//   * targetGroupId - Group ID.
//   * name - New name of the target group — a string up to 64 characters long.
//   * domain - Domain of the site where user accounting code will be placed.
//   * lifetime - 'Only for the groups that get audience from sites with user accounting code.', Time in days when users added to a retarget group will be automatically excluded from it. '0' – automatic exclusion is off.
//   * targetPixelId - NO DESCRIPTION IN JSON SCHEMA
//   * targetPixelRules - NO DESCRIPTION IN JSON SCHEMA
func (a Ads) Updatetargetgroup(accountId int, clientId int, targetGroupId int, name string, domain string, lifetime int, targetPixelId int, targetPixelRules string) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["account_id"] = accountId

	if clientId > 0 {
		params["client_id"] = clientId
	}

	params["target_group_id"] = targetGroupId

	params["name"] = name

	if domain != "" {
		params["domain"] = domain
	}

	if lifetime > 0 {
		params["lifetime"] = lifetime
	}

	if targetPixelId > 0 {
		params["target_pixel_id"] = targetPixelId
	}

	if targetPixelRules != "" {
		params["target_pixel_rules"] = targetPixelRules
	}

	err = a.SendObjRequest("ads.updateTargetGroup", params, &resp)

	return
}
