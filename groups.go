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
	"gitlab.com/Burmuley/go-vkapi/objects"
	"gitlab.com/Burmuley/go-vkapi/responses"
)

type Groups struct {
	*VKApi
}

/////////////////////////////////////////////////////////////
// `Groups` methods
/////////////////////////////////////////////////////////////

// Addaddress - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * groupId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * title - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * address - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * additionalAddress - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * countryId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * cityId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * metroId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * latitude - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * longitude - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * phone - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * workInfoStatus - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * timetable - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * isMainAddress - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (g Groups) Addaddress(groupId int, title string, address string, additionalAddress string, countryId int, cityId int, metroId int, latitude float64, longitude float64, phone string, workInfoStatus objects.GroupsAddressWorkInfoStatus, timetable string, isMainAddress bool) (resp responses.GroupsAddAddress, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	params["title"] = title

	params["address"] = address

	if additionalAddress != "" {
		params["additional_address"] = additionalAddress
	}

	params["country_id"] = countryId

	params["city_id"] = cityId

	if metroId > 0 {
		params["metro_id"] = metroId
	}

	params["latitude"] = latitude

	params["longitude"] = longitude

	if phone != "" {
		params["phone"] = phone
	}

	if timetable != "" {
		params["timetable"] = timetable
	}

	params["is_main_address"] = isMainAddress

	err = g.SendObjRequest("groups.addAddress", params, &resp)

	return
}

// Addcallbackserver - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * groupId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * url - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * title - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * secretKey - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (g Groups) Addcallbackserver(groupId int, url string, title string, secretKey string) (resp responses.GroupsAddcallbackserver, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	params["url"] = url

	params["title"] = title

	if secretKey != "" {
		params["secret_key"] = secretKey
	}

	err = g.SendObjRequest("groups.addCallbackServer", params, &resp)

	return
}

// Addlink - Allows to add a link to the community.
// Parameters:
//   * groupId - Community ID.
//   * link - Link URL.
//   * text - Description text for the link.
func (g Groups) Addlink(groupId int, link string, text string) (resp responses.GroupsAddlink, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	params["link"] = link

	if text != "" {
		params["text"] = text
	}

	err = g.SendObjRequest("groups.addLink", params, &resp)

	return
}

// Approverequest - Allows to approve join request to the community.
// Parameters:
//   * groupId - Community ID.
//   * userId - User ID.
func (g Groups) Approverequest(groupId int, userId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	params["user_id"] = userId

	err = g.SendObjRequest("groups.approveRequest", params, &resp)

	return
}

// Ban - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * groupId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * ownerId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * endDate - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * reason - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * comment - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * commentVisible - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (g Groups) Ban(groupId int, ownerId int, endDate int, reason int, comment string, commentVisible bool) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	if endDate > 0 {
		params["end_date"] = endDate
	}

	if reason > 0 {
		params["reason"] = reason
	}

	if comment != "" {
		params["comment"] = comment
	}

	params["comment_visible"] = commentVisible

	err = g.SendObjRequest("groups.ban", params, &resp)

	return
}

// Create - Creates a new community.
// Parameters:
//   * title - Community title.
//   * description - Community description (ignored for 'type' = 'public').
//   * pType - Community type. Possible values: *'group' – group,, *'event' – event,, *'public' – public page
//   * publicCategory - Category ID (for 'type' = 'public' only).
//   * subtype - Public page subtype. Possible values: *'1' – place or small business,, *'2' – company, organization or website,, *'3' – famous person or group of people,, *'4' – product or work of art.
func (g Groups) Create(title string, description string, pType string, publicCategory int, subtype int) (resp responses.GroupsCreate, err error) {
	params := map[string]interface{}{}

	params["title"] = title

	if description != "" {
		params["description"] = description
	}

	if pType != "" {
		params["type"] = pType
	}

	if publicCategory > 0 {
		params["public_category"] = publicCategory
	}

	if subtype > 0 {
		params["subtype"] = subtype
	}

	err = g.SendObjRequest("groups.create", params, &resp)

	return
}

// Deletecallbackserver - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * groupId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * serverId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (g Groups) Deletecallbackserver(groupId int, serverId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	params["server_id"] = serverId

	err = g.SendObjRequest("groups.deleteCallbackServer", params, &resp)

	return
}

// Deletelink - Allows to delete a link from the community.
// Parameters:
//   * groupId - Community ID.
//   * linkId - Link ID.
func (g Groups) Deletelink(groupId int, linkId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	params["link_id"] = linkId

	err = g.SendObjRequest("groups.deleteLink", params, &resp)

	return
}

// Disableonline - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * groupId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (g Groups) Disableonline(groupId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	err = g.SendObjRequest("groups.disableOnline", params, &resp)

	return
}

// Edit - Edits a community.
// Parameters:
//   * groupId - Community ID.
//   * title - Community title.
//   * description - Community description.
//   * screenName - Community screen name.
//   * access - Community type. Possible values: *'0' – open,, *'1' – closed,, *'2' – private.
//   * website - Website that will be displayed in the community information field.
//   * subject - Community subject. Possible values: , *'1' – auto/moto,, *'2' – activity holidays,, *'3' – business,, *'4' – pets,, *'5' – health,, *'6' – dating and communication, , *'7' – games,, *'8' – IT (computers and software),, *'9' – cinema,, *'10' – beauty and fashion,, *'11' – cooking,, *'12' – art and culture,, *'13' – literature,, *'14' – mobile services and internet,, *'15' – music,, *'16' – science and technology,, *'17' – real estate,, *'18' – news and media,, *'19' – security,, *'20' – education,, *'21' – home and renovations,, *'22' – politics,, *'23' – food,, *'24' – industry,, *'25' – travel,, *'26' – work,, *'27' – entertainment,, *'28' – religion,, *'29' – family,, *'30' – sports,, *'31' – insurance,, *'32' – television,, *'33' – goods and services,, *'34' – hobbies,, *'35' – finance,, *'36' – photo,, *'37' – esoterics,, *'38' – electronics and appliances,, *'39' – erotic,, *'40' – humor,, *'41' – society, humanities,, *'42' – design and graphics.
//   * email - Organizer email (for events).
//   * phone - Organizer phone number (for events).
//   * rss - RSS feed address for import (available only to communities with special permission. Contact vk.com/support to get it.
//   * eventStartDate - Event start date in Unixtime format.
//   * eventFinishDate - Event finish date in Unixtime format.
//   * eventGroupId - Organizer community ID (for events only).
//   * publicCategory - Public page category ID.
//   * publicSubcategory - Public page subcategory ID.
//   * publicDate - Founding date of a company or organization owning the community in "dd.mm.YYYY" format.
//   * wall - Wall settings. Possible values: *'0' – disabled,, *'1' – open,, *'2' – limited (groups and events only),, *'3' – closed (groups and events only).
//   * topics - Board topics settings. Possbile values: , *'0' – disabled,, *'1' – open,, *'2' – limited (for groups and events only).
//   * photos - Photos settings. Possible values: *'0' – disabled,, *'1' – open,, *'2' – limited (for groups and events only).
//   * video - Video settings. Possible values: *'0' – disabled,, *'1' – open,, *'2' – limited (for groups and events only).
//   * audio - Audio settings. Possible values: *'0' – disabled,, *'1' – open,, *'2' – limited (for groups and events only).
//   * links - Links settings (for public pages only). Possible values: *'0' – disabled,, *'1' – enabled.
//   * events - Events settings (for public pages only). Possible values: *'0' – disabled,, *'1' – enabled.
//   * places - Places settings (for public pages only). Possible values: *'0' – disabled,, *'1' – enabled.
//   * contacts - Contacts settings (for public pages only). Possible values: *'0' – disabled,, *'1' – enabled.
//   * docs - Documents settings. Possible values: *'0' – disabled,, *'1' – open,, *'2' – limited (for groups and events only).
//   * wiki - Wiki pages settings. Possible values: *'0' – disabled,, *'1' – open,, *'2' – limited (for groups and events only).
//   * messages - Community messages. Possible values: *'0' — disabled,, *'1' — enabled.
//   * articles - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * addresses - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * ageLimits - Community age limits. Possible values: *'1' — no limits,, *'2' — 16+,, *'3' — 18+.
//   * market - Market settings. Possible values: *'0' – disabled,, *'1' – enabled.
//   * marketComments - market comments settings. Possible values: *'0' – disabled,, *'1' – enabled.
//   * marketCountry - Market delivery countries.
//   * marketCity - Market delivery cities (if only one country is specified).
//   * marketCurrency - Market currency settings. Possbile values: , *'643' – Russian rubles,, *'980' – Ukrainian hryvnia,, *'398' – Kazakh tenge,, *'978' – Euro,, *'840' – US dollars
//   * marketContact - Seller contact for market. Set '0' for community messages.
//   * marketWiki - ID of a wiki page with market description.
//   * obsceneFilter - Obscene expressions filter in comments. Possible values: , *'0' – disabled,, *'1' – enabled.
//   * obsceneStopwords - Stopwords filter in comments. Possible values: , *'0' – disabled,, *'1' – enabled.
//   * obsceneWords - Keywords for stopwords filter.
//   * mainSection - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * secondarySection - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * country - Country of the community.
//   * city - City of the community.
func (g Groups) Edit(groupId int, title string, description string, screenName string, access objects.GroupsGroupAccess, website string, subject objects.GroupsGroupSubject, email string, phone string, rss string, eventStartDate int, eventFinishDate int, eventGroupId int, publicCategory int, publicSubcategory int, publicDate string, wall objects.GroupsGroupWall, topics objects.GroupsGroupTopics, photos objects.GroupsGroupPhotos, video objects.GroupsGroupVideo, audio objects.GroupsGroupAudio, links bool, events bool, places bool, contacts bool, docs objects.GroupsGroupDocs, wiki objects.GroupsGroupWiki, messages bool, articles bool, addresses bool, ageLimits objects.GroupsGroupAgeLimits, market bool, marketComments bool, marketCountry []int, marketCity []int, marketCurrency objects.GroupsGroupMarketCurrency, marketContact int, marketWiki int, obsceneFilter bool, obsceneStopwords bool, obsceneWords []string, mainSection int, secondarySection int, country int, city int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	if title != "" {
		params["title"] = title
	}

	if description != "" {
		params["description"] = description
	}

	if screenName != "" {
		params["screen_name"] = screenName
	}

	if access > 0 {
		params["access"] = access
	}

	if website != "" {
		params["website"] = website
	}

	if email != "" {
		params["email"] = email
	}

	if phone != "" {
		params["phone"] = phone
	}

	if rss != "" {
		params["rss"] = rss
	}

	if eventStartDate > 0 {
		params["event_start_date"] = eventStartDate
	}

	if eventFinishDate > 0 {
		params["event_finish_date"] = eventFinishDate
	}

	if eventGroupId > 0 {
		params["event_group_id"] = eventGroupId
	}

	if publicCategory > 0 {
		params["public_category"] = publicCategory
	}

	if publicSubcategory > 0 {
		params["public_subcategory"] = publicSubcategory
	}

	if publicDate != "" {
		params["public_date"] = publicDate
	}

	if wall > 0 {
		params["wall"] = wall
	}

	if topics > 0 {
		params["topics"] = topics
	}

	if photos > 0 {
		params["photos"] = photos
	}

	if video > 0 {
		params["video"] = video
	}

	if audio > 0 {
		params["audio"] = audio
	}

	params["links"] = links

	params["events"] = events

	params["places"] = places

	params["contacts"] = contacts

	if docs > 0 {
		params["docs"] = docs
	}

	if wiki > 0 {
		params["wiki"] = wiki
	}

	params["messages"] = messages

	params["articles"] = articles

	params["addresses"] = addresses

	if ageLimits > 0 {
		params["age_limits"] = ageLimits
	}

	params["market"] = market

	params["market_comments"] = marketComments

	if len(marketCountry) > 0 {
		params["market_country"] = SliceToString(marketCountry)
	}

	if len(marketCity) > 0 {
		params["market_city"] = SliceToString(marketCity)
	}

	if marketCurrency > 0 {
		params["market_currency"] = marketCurrency
	}

	if marketContact > 0 {
		params["market_contact"] = marketContact
	}

	if marketWiki > 0 {
		params["market_wiki"] = marketWiki
	}

	params["obscene_filter"] = obsceneFilter

	params["obscene_stopwords"] = obsceneStopwords

	if len(obsceneWords) > 0 {
		params["obscene_words"] = SliceToString(obsceneWords)
	}

	if mainSection > 0 {
		params["main_section"] = mainSection
	}

	if secondarySection > 0 {
		params["secondary_section"] = secondarySection
	}

	if country > 0 {
		params["country"] = country
	}

	if city > 0 {
		params["city"] = city
	}

	err = g.SendObjRequest("groups.edit", params, &resp)

	return
}

// Editaddress - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * groupId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * addressId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * title - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * address - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * additionalAddress - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * countryId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * cityId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * metroId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * latitude - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * longitude - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * phone - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * workInfoStatus - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * timetable - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * isMainAddress - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (g Groups) Editaddress(groupId int, addressId int, title string, address string, additionalAddress string, countryId int, cityId int, metroId int, latitude float64, longitude float64, phone string, workInfoStatus objects.GroupsAddressWorkInfoStatus, timetable string, isMainAddress bool) (resp responses.GroupsEditaddress, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	params["address_id"] = addressId

	if title != "" {
		params["title"] = title
	}

	if address != "" {
		params["address"] = address
	}

	if additionalAddress != "" {
		params["additional_address"] = additionalAddress
	}

	if countryId > 0 {
		params["country_id"] = countryId
	}

	if cityId > 0 {
		params["city_id"] = cityId
	}

	if metroId > 0 {
		params["metro_id"] = metroId
	}

	if latitude > 0 {
		params["latitude"] = latitude
	}

	if longitude > 0 {
		params["longitude"] = longitude
	}

	if phone != "" {
		params["phone"] = phone
	}

	if timetable != "" {
		params["timetable"] = timetable
	}

	params["is_main_address"] = isMainAddress

	err = g.SendObjRequest("groups.editAddress", params, &resp)

	return
}

// Editcallbackserver - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * groupId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * serverId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * url - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * title - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * secretKey - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (g Groups) Editcallbackserver(groupId int, serverId int, url string, title string, secretKey string) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	params["server_id"] = serverId

	params["url"] = url

	params["title"] = title

	if secretKey != "" {
		params["secret_key"] = secretKey
	}

	err = g.SendObjRequest("groups.editCallbackServer", params, &resp)

	return
}

// Editlink - Allows to edit a link in the community.
// Parameters:
//   * groupId - Community ID.
//   * linkId - Link ID.
//   * text - New description text for the link.
func (g Groups) Editlink(groupId int, linkId int, text string) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	params["link_id"] = linkId

	if text != "" {
		params["text"] = text
	}

	err = g.SendObjRequest("groups.editLink", params, &resp)

	return
}

// Editmanager - Allows to add, remove or edit the community manager.
// Parameters:
//   * groupId - Community ID.
//   * userId - User ID.
//   * role - Manager role. Possible values: *'moderator',, *'editor',, *'administrator'.
//   * isContact - '1' — to show the manager in Contacts block of the community.
//   * contactPosition - Position to show in Contacts block.
//   * contactPhone - Contact phone.
//   * contactEmail - Contact e-mail.
func (g Groups) Editmanager(groupId int, userId int, role objects.GroupsGroupRole, isContact bool, contactPosition string, contactPhone string, contactEmail string) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	params["user_id"] = userId

	params["is_contact"] = isContact

	if contactPosition != "" {
		params["contact_position"] = contactPosition
	}

	if contactPhone != "" {
		params["contact_phone"] = contactPhone
	}

	if contactEmail != "" {
		params["contact_email"] = contactEmail
	}

	err = g.SendObjRequest("groups.editManager", params, &resp)

	return
}

// Enableonline - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * groupId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (g Groups) Enableonline(groupId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	err = g.SendObjRequest("groups.enableOnline", params, &resp)

	return
}

// Get - Returns a list of the communities to which a user belongs.
// Parameters:
//   * userId - User ID.
//   * extended - '1' — to return complete information about a user's communities, '0' — to return a list of community IDs without any additional fields (default),
//   * filter - Types of communities to return: 'admin' — to return communities administered by the user , 'editor' — to return communities where the user is an administrator or editor, 'moder' — to return communities where the user is an administrator, editor, or moderator, 'groups' — to return only groups, 'publics' — to return only public pages, 'events' — to return only events
//   * fields - Profile fields to return.
//   * offset - Offset needed to return a specific subset of communities.
//   * count - Number of communities to return.
func (g Groups) Get(userId int, filter []objects.GroupsFilter, fields []objects.GroupsFields, offset int, count int) (resp responses.GroupsGet, err error) {
	params := map[string]interface{}{}
	params["extended"] = "0"

	if userId > 0 {
		params["user_id"] = userId
	}

	if len(filter) > 0 {
		params["filter"] = SliceToString(filter)
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	err = g.SendObjRequest("groups.get", params, &resp)

	return
}

// GetExtended - Returns a list of the communities to which a user belongs.
// Parameters:
//   * userId - User ID.
//   * extended - '1' — to return complete information about a user's communities, '0' — to return a list of community IDs without any additional fields (default),
//   * filter - Types of communities to return: 'admin' — to return communities administered by the user , 'editor' — to return communities where the user is an administrator or editor, 'moder' — to return communities where the user is an administrator, editor, or moderator, 'groups' — to return only groups, 'publics' — to return only public pages, 'events' — to return only events
//   * fields - Profile fields to return.
//   * offset - Offset needed to return a specific subset of communities.
//   * count - Number of communities to return.
func (g Groups) GetExtended(userId int, filter []objects.GroupsFilter, fields []objects.GroupsFields, offset int, count int) (resp responses.GroupsGetExtended, err error) {
	params := map[string]interface{}{}
	params["extended"] = "1"

	if userId > 0 {
		params["user_id"] = userId
	}

	if len(filter) > 0 {
		params["filter"] = SliceToString(filter)
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	err = g.SendObjRequest("groups.get", params, &resp)

	return
}

// Getaddresses - Returns a list of community addresses.
// Parameters:
//   * groupId - ID or screen name of the community.
//   * addressIds - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * latitude - Latitude of  the user geo position.
//   * longitude - Longitude of the user geo position.
//   * offset - Offset needed to return a specific subset of community addresses.
//   * count - Number of community addresses to return.
//   * fields - Address fields
func (g Groups) Getaddresses(groupId int, addressIds []int, latitude float64, longitude float64, offset int, count int, fields []objects.AddressesFields) (resp responses.GroupsGetaddresses, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	if len(addressIds) > 0 {
		params["address_ids"] = SliceToString(addressIds)
	}

	if latitude > 0 {
		params["latitude"] = latitude
	}

	if longitude > 0 {
		params["longitude"] = longitude
	}

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	err = g.SendObjRequest("groups.getAddresses", params, &resp)

	return
}

// Getbanned - Returns a list of users on a community blacklist.
// Parameters:
//   * groupId - Community ID.
//   * offset - Offset needed to return a specific subset of users.
//   * count - Number of users to return.
//   * fields - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * ownerId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (g Groups) Getbanned(groupId int, offset int, count int, fields []objects.BaseUserGroupFields, ownerId int) (resp responses.GroupsGetbanned, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	err = g.SendObjRequest("groups.getBanned", params, &resp)

	return
}

// Getbyid - Returns information about communities by their IDs.
// Parameters:
//   * groupIds - IDs or screen names of communities.
//   * groupId - ID or screen name of the community.
//   * fields - Group fields to return.
func (g Groups) Getbyid(groupIds []string, groupId string, fields []objects.GroupsFields) (resp responses.GroupsGetbyid, err error) {
	params := map[string]interface{}{}

	if len(groupIds) > 0 {
		params["group_ids"] = SliceToString(groupIds)
	}

	if groupId != "" {
		params["group_id"] = groupId
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	err = g.SendObjRequest("groups.getById", params, &resp)

	return
}

// Getcallbackconfirmationcode - Returns Callback API confirmation code for the community.
// Parameters:
//   * groupId - Community ID.
func (g Groups) Getcallbackconfirmationcode(groupId int) (resp responses.GroupsGetcallbackconfirmationcode, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	err = g.SendObjRequest("groups.getCallbackConfirmationCode", params, &resp)

	return
}

// Getcallbackservers - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * groupId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * serverIds - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (g Groups) Getcallbackservers(groupId int, serverIds []int) (resp responses.GroupsGetcallbackservers, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	if len(serverIds) > 0 {
		params["server_ids"] = SliceToString(serverIds)
	}

	err = g.SendObjRequest("groups.getCallbackServers", params, &resp)

	return
}

// Getcallbacksettings - Returns [vk.com/dev/callback_api|Callback API] notifications settings.
// Parameters:
//   * groupId - Community ID.
//   * serverId - Server ID.
func (g Groups) Getcallbacksettings(groupId int, serverId int) (resp responses.GroupsGetcallbacksettings, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	if serverId > 0 {
		params["server_id"] = serverId
	}

	err = g.SendObjRequest("groups.getCallbackSettings", params, &resp)

	return
}

// Getcatalog - Returns communities list for a catalog category.
// Parameters:
//   * categoryId - Category id received from [vk.com/dev/groups.getCatalogInfo|groups.getCatalogInfo].
//   * subcategoryId - Subcategory id received from [vk.com/dev/groups.getCatalogInfo|groups.getCatalogInfo].
func (g Groups) Getcatalog(categoryId int, subcategoryId int) (resp responses.GroupsGetcatalog, err error) {
	params := map[string]interface{}{}

	if categoryId > 0 {
		params["category_id"] = categoryId
	}

	if subcategoryId > 0 {
		params["subcategory_id"] = subcategoryId
	}

	err = g.SendObjRequest("groups.getCatalog", params, &resp)

	return
}

// Getcataloginfo - Returns categories list for communities catalog
// Parameters:
//   * extended - 1 – to return communities count and three communities for preview. By default: 0.
//   * subcategories - 1 – to return subcategories info. By default: 0.
func (g Groups) Getcataloginfo(subcategories bool) (resp responses.GroupsGetcataloginfo, err error) {
	params := map[string]interface{}{}
	params["extended"] = "0"

	params["subcategories"] = subcategories

	err = g.SendObjRequest("groups.getCatalogInfo", params, &resp)

	return
}

// GetcataloginfoExtended - Returns categories list for communities catalog
// Parameters:
//   * extended - 1 – to return communities count and three communities for preview. By default: 0.
//   * subcategories - 1 – to return subcategories info. By default: 0.
func (g Groups) GetcataloginfoExtended(subcategories bool) (resp responses.GroupsGetcataloginfoExtended, err error) {
	params := map[string]interface{}{}
	params["extended"] = "1"

	params["subcategories"] = subcategories

	err = g.SendObjRequest("groups.getCatalogInfo", params, &resp)

	return
}

// Getinvitedusers - Returns invited users list of a community
// Parameters:
//   * groupId - Group ID to return invited users for.
//   * offset - Offset needed to return a specific subset of results.
//   * count - Number of results to return.
//   * fields - List of additional fields to be returned. Available values: 'sex, bdate, city, country, photo_50, photo_100, photo_200_orig, photo_200, photo_400_orig, photo_max, photo_max_orig, online, online_mobile, lists, domain, has_mobile, contacts, connections, site, education, universities, schools, can_post, can_see_all_posts, can_see_audio, can_write_private_message, status, last_seen, common_count, relation, relatives, counters'.
//   * nameCase - Case for declension of user name and surname. Possible values: *'nom' — nominative (default),, *'gen' — genitive,, *'dat' — dative,, *'acc' — accusative, , *'ins' — instrumental,, *'abl' — prepositional.
func (g Groups) Getinvitedusers(groupId int, offset int, count int, fields []objects.UsersFields, nameCase string) (resp responses.GroupsGetinvitedusers, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	if nameCase != "" {
		params["name_case"] = nameCase
	}

	err = g.SendObjRequest("groups.getInvitedUsers", params, &resp)

	return
}

// Getinvites - Returns a list of invitations to join communities and events.
// Parameters:
//   * offset - Offset needed to return a specific subset of invitations.
//   * count - Number of invitations to return.
//   * extended - '1' — to return additional [vk.com/dev/fields_groups|fields] for communities..
func (g Groups) Getinvites(offset int, count int) (resp responses.GroupsGetinvites, err error) {
	params := map[string]interface{}{}
	params["extended"] = "0"

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	err = g.SendObjRequest("groups.getInvites", params, &resp)

	return
}

// GetinvitesExtended - Returns a list of invitations to join communities and events.
// Parameters:
//   * offset - Offset needed to return a specific subset of invitations.
//   * count - Number of invitations to return.
//   * extended - '1' — to return additional [vk.com/dev/fields_groups|fields] for communities..
func (g Groups) GetinvitesExtended(offset int, count int) (resp responses.GroupsGetinvitesExtended, err error) {
	params := map[string]interface{}{}
	params["extended"] = "1"

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	err = g.SendObjRequest("groups.getInvites", params, &resp)

	return
}

// Getlongpollserver - Returns the data needed to query a Long Poll server for events
// Parameters:
//   * groupId - Community ID
func (g Groups) Getlongpollserver(groupId int) (resp responses.GroupsGetlongpollserver, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	err = g.SendObjRequest("groups.getLongPollServer", params, &resp)

	return
}

// Getlongpollsettings - Returns Long Poll notification settings
// Parameters:
//   * groupId - Community ID.
func (g Groups) Getlongpollsettings(groupId int) (resp responses.GroupsGetlongpollsettings, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	err = g.SendObjRequest("groups.getLongPollSettings", params, &resp)

	return
}

// Getmembers - Returns a list of community members.
// Parameters:
//   * groupId - ID or screen name of the community.
//   * sort - Sort order. Available values: 'id_asc', 'id_desc', 'time_asc', 'time_desc'. 'time_asc' and 'time_desc' are availavle only if the method is called by the group's 'moderator'.
//   * offset - Offset needed to return a specific subset of community members.
//   * count - Number of community members to return.
//   * fields - List of additional fields to be returned. Available values: 'sex, bdate, city, country, photo_50, photo_100, photo_200_orig, photo_200, photo_400_orig, photo_max, photo_max_orig, online, online_mobile, lists, domain, has_mobile, contacts, connections, site, education, universities, schools, can_post, can_see_all_posts, can_see_audio, can_write_private_message, status, last_seen, common_count, relation, relatives, counters'.
//   * filter - *'friends' – only friends in this community will be returned,, *'unsure' – only those who pressed 'I may attend' will be returned (if it's an event).
func (g Groups) Getmembers(groupId string, sort string, offset int, count int, fields []objects.UsersFields, filter string) (resp responses.GroupsGetmembers, err error) {
	params := map[string]interface{}{}

	if groupId != "" {
		params["group_id"] = groupId
	}

	if sort != "" {
		params["sort"] = sort
	}

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	if filter != "" {
		params["filter"] = filter
	}

	err = g.SendObjRequest("groups.getMembers", params, &resp)

	return
}

// Getrequests - Returns a list of requests to the community.
// Parameters:
//   * groupId - Community ID.
//   * offset - Offset needed to return a specific subset of results.
//   * count - Number of results to return.
//   * fields - Profile fields to return.
func (g Groups) Getrequests(groupId int, offset int, count int, fields []objects.UsersFields) (resp responses.GroupsGetrequests, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	err = g.SendObjRequest("groups.getRequests", params, &resp)

	return
}

// Getsettings - Returns community settings.
// Parameters:
//   * groupId - Community ID.
func (g Groups) Getsettings(groupId int) (resp responses.GroupsGetsettings, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	err = g.SendObjRequest("groups.getSettings", params, &resp)

	return
}

// Gettokenpermissions - NO DESCRIPTION IN JSON SCHEMA
func (g Groups) Gettokenpermissions() (resp responses.GroupsGettokenpermissions, err error) {
	params := map[string]interface{}{}

	err = g.SendObjRequest("groups.getTokenPermissions", params, &resp)

	return
}

// Invite - Allows to invite friends to the community.
// Parameters:
//   * groupId - Community ID.
//   * userId - User ID.
func (g Groups) Invite(groupId int, userId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	params["user_id"] = userId

	err = g.SendObjRequest("groups.invite", params, &resp)

	return
}

// Ismember - Returns information specifying whether a user is a member of a community.
// Parameters:
//   * groupId - ID or screen name of the community.
//   * userId - User ID.
//   * userIds - User IDs.
//   * extended - '1' — to return an extended response with additional fields. By default: '0'.
func (g Groups) Ismember(groupId string, userId int, userIds []int) (resp responses.GroupsIsmember, err error) {
	params := map[string]interface{}{}
	params["extended"] = "0"

	params["group_id"] = groupId

	if userId > 0 {
		params["user_id"] = userId
	}

	if len(userIds) > 0 {
		params["user_ids"] = SliceToString(userIds)
	}

	err = g.SendObjRequest("groups.isMember", params, &resp)

	return
}

// IsmemberExtended - Returns information specifying whether a user is a member of a community.
// Parameters:
//   * groupId - ID or screen name of the community.
//   * userId - User ID.
//   * userIds - User IDs.
//   * extended - '1' — to return an extended response with additional fields. By default: '0'.
func (g Groups) IsmemberExtended(groupId string, userId int, userIds []int) (resp responses.GroupsIsmemberExtended, err error) {
	params := map[string]interface{}{}
	params["extended"] = "1"

	params["group_id"] = groupId

	if userId > 0 {
		params["user_id"] = userId
	}

	if len(userIds) > 0 {
		params["user_ids"] = SliceToString(userIds)
	}

	err = g.SendObjRequest("groups.isMember", params, &resp)

	return
}

// Join - With this method you can join the group or public page, and also confirm your participation in an event.
// Parameters:
//   * groupId - ID or screen name of the community.
//   * notSure - Optional parameter which is taken into account when 'gid' belongs to the event: '1' — Perhaps I will attend, '0' — I will be there for sure (default), ,
func (g Groups) Join(groupId int, notSure string) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	if groupId > 0 {
		params["group_id"] = groupId
	}

	if notSure != "" {
		params["not_sure"] = notSure
	}

	err = g.SendObjRequest("groups.join", params, &resp)

	return
}

// Leave - With this method you can leave a group, public page, or event.
// Parameters:
//   * groupId - ID or screen name of the community.
func (g Groups) Leave(groupId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	err = g.SendObjRequest("groups.leave", params, &resp)

	return
}

// Removeuser - Removes a user from the community.
// Parameters:
//   * groupId - Community ID.
//   * userId - User ID.
func (g Groups) Removeuser(groupId int, userId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	params["user_id"] = userId

	err = g.SendObjRequest("groups.removeUser", params, &resp)

	return
}

// Reorderlink - Allows to reorder links in the community.
// Parameters:
//   * groupId - Community ID.
//   * linkId - Link ID.
//   * after - ID of the link after which to place the link with 'link_id'.
func (g Groups) Reorderlink(groupId int, linkId int, after int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	params["link_id"] = linkId

	if after > 0 {
		params["after"] = after
	}

	err = g.SendObjRequest("groups.reorderLink", params, &resp)

	return
}

// Search - Returns a list of communities matching the search criteria.
// Parameters:
//   * q - Search query string.
//   * pType - Community type. Possible values: 'group, page, event.'
//   * countryId - Country ID.
//   * cityId - City ID. If this parameter is transmitted, country_id is ignored.
//   * future - '1' — to return only upcoming events. Works with the 'type' = 'event' only.
//   * market - '1' — to return communities with enabled market only.
//   * sort - Sort order. Possible values: *'0' — default sorting (similar the full version of the site),, *'1' — by growth speed,, *'2'— by the "day attendance/members number" ratio,, *'3' — by the "Likes number/members number" ratio,, *'4' — by the "comments number/members number" ratio,, *'5' — by the "boards entries number/members number" ratio.
//   * offset - Offset needed to return a specific subset of results.
//   * count - Number of communities to return. "Note that you can not receive more than first thousand of results, regardless of 'count' and 'offset' values."
func (g Groups) Search(q string, pType string, countryId int, cityId int, future bool, market bool, sort int, offset int, count int) (resp responses.GroupsSearch, err error) {
	params := map[string]interface{}{}

	params["q"] = q

	if pType != "" {
		params["type"] = pType
	}

	if countryId > 0 {
		params["country_id"] = countryId
	}

	if cityId > 0 {
		params["city_id"] = cityId
	}

	params["future"] = future

	params["market"] = market

	if sort > 0 {
		params["sort"] = sort
	}

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	err = g.SendObjRequest("groups.search", params, &resp)

	return
}

// Setcallbacksettings - Allow to set notifications settings for group.
// Parameters:
//   * groupId - Community ID.
//   * serverId - Server ID.
//   * apiVersion - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * messageNew - A new incoming message has been received ('0' — disabled, '1' — enabled).
//   * messageReply - A new outcoming message has been received ('0' — disabled, '1' — enabled).
//   * messageAllow - Allowed messages notifications ('0' — disabled, '1' — enabled).
//   * messageEdit - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * messageDeny - Denied messages notifications ('0' — disabled, '1' — enabled).
//   * messageTypingState - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * photoNew - New photos notifications ('0' — disabled, '1' — enabled).
//   * audioNew - New audios notifications ('0' — disabled, '1' — enabled).
//   * videoNew - New videos notifications ('0' — disabled, '1' — enabled).
//   * wallReplyNew - New wall replies notifications ('0' — disabled, '1' — enabled).
//   * wallReplyEdit - Wall replies edited notifications ('0' — disabled, '1' — enabled).
//   * wallReplyDelete - A wall comment has been deleted ('0' — disabled, '1' — enabled).
//   * wallReplyRestore - A wall comment has been restored ('0' — disabled, '1' — enabled).
//   * wallPostNew - New wall posts notifications ('0' — disabled, '1' — enabled).
//   * wallRepost - New wall posts notifications ('0' — disabled, '1' — enabled).
//   * boardPostNew - New board posts notifications ('0' — disabled, '1' — enabled).
//   * boardPostEdit - Board posts edited notifications ('0' — disabled, '1' — enabled).
//   * boardPostRestore - Board posts restored notifications ('0' — disabled, '1' — enabled).
//   * boardPostDelete - Board posts deleted notifications ('0' — disabled, '1' — enabled).
//   * photoCommentNew - New comment to photo notifications ('0' — disabled, '1' — enabled).
//   * photoCommentEdit - A photo comment has been edited ('0' — disabled, '1' — enabled).
//   * photoCommentDelete - A photo comment has been deleted ('0' — disabled, '1' — enabled).
//   * photoCommentRestore - A photo comment has been restored ('0' — disabled, '1' — enabled).
//   * videoCommentNew - New comment to video notifications ('0' — disabled, '1' — enabled).
//   * videoCommentEdit - A video comment has been edited ('0' — disabled, '1' — enabled).
//   * videoCommentDelete - A video comment has been deleted ('0' — disabled, '1' — enabled).
//   * videoCommentRestore - A video comment has been restored ('0' — disabled, '1' — enabled).
//   * marketCommentNew - New comment to market item notifications ('0' — disabled, '1' — enabled).
//   * marketCommentEdit - A market comment has been edited ('0' — disabled, '1' — enabled).
//   * marketCommentDelete - A market comment has been deleted ('0' — disabled, '1' — enabled).
//   * marketCommentRestore - A market comment has been restored ('0' — disabled, '1' — enabled).
//   * pollVoteNew - A vote in a public poll has been added ('0' — disabled, '1' — enabled).
//   * groupJoin - Joined community notifications ('0' — disabled, '1' — enabled).
//   * groupLeave - Left community notifications ('0' — disabled, '1' — enabled).
//   * groupChangeSettings - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * groupChangePhoto - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * groupOfficersEdit - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * userBlock - User added to community blacklist
//   * userUnblock - User removed from community blacklist
//   * leadFormsNew - New form in lead forms
func (g Groups) Setcallbacksettings(groupId int, serverId int, apiVersion string, messageNew bool, messageReply bool, messageAllow bool, messageEdit bool, messageDeny bool, messageTypingState bool, photoNew bool, audioNew bool, videoNew bool, wallReplyNew bool, wallReplyEdit bool, wallReplyDelete bool, wallReplyRestore bool, wallPostNew bool, wallRepost bool, boardPostNew bool, boardPostEdit bool, boardPostRestore bool, boardPostDelete bool, photoCommentNew bool, photoCommentEdit bool, photoCommentDelete bool, photoCommentRestore bool, videoCommentNew bool, videoCommentEdit bool, videoCommentDelete bool, videoCommentRestore bool, marketCommentNew bool, marketCommentEdit bool, marketCommentDelete bool, marketCommentRestore bool, pollVoteNew bool, groupJoin bool, groupLeave bool, groupChangeSettings bool, groupChangePhoto bool, groupOfficersEdit bool, userBlock bool, userUnblock bool, leadFormsNew bool) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	if serverId > 0 {
		params["server_id"] = serverId
	}

	if apiVersion != "" {
		params["api_version"] = apiVersion
	}

	params["message_new"] = messageNew

	params["message_reply"] = messageReply

	params["message_allow"] = messageAllow

	params["message_edit"] = messageEdit

	params["message_deny"] = messageDeny

	params["message_typing_state"] = messageTypingState

	params["photo_new"] = photoNew

	params["audio_new"] = audioNew

	params["video_new"] = videoNew

	params["wall_reply_new"] = wallReplyNew

	params["wall_reply_edit"] = wallReplyEdit

	params["wall_reply_delete"] = wallReplyDelete

	params["wall_reply_restore"] = wallReplyRestore

	params["wall_post_new"] = wallPostNew

	params["wall_repost"] = wallRepost

	params["board_post_new"] = boardPostNew

	params["board_post_edit"] = boardPostEdit

	params["board_post_restore"] = boardPostRestore

	params["board_post_delete"] = boardPostDelete

	params["photo_comment_new"] = photoCommentNew

	params["photo_comment_edit"] = photoCommentEdit

	params["photo_comment_delete"] = photoCommentDelete

	params["photo_comment_restore"] = photoCommentRestore

	params["video_comment_new"] = videoCommentNew

	params["video_comment_edit"] = videoCommentEdit

	params["video_comment_delete"] = videoCommentDelete

	params["video_comment_restore"] = videoCommentRestore

	params["market_comment_new"] = marketCommentNew

	params["market_comment_edit"] = marketCommentEdit

	params["market_comment_delete"] = marketCommentDelete

	params["market_comment_restore"] = marketCommentRestore

	params["poll_vote_new"] = pollVoteNew

	params["group_join"] = groupJoin

	params["group_leave"] = groupLeave

	params["group_change_settings"] = groupChangeSettings

	params["group_change_photo"] = groupChangePhoto

	params["group_officers_edit"] = groupOfficersEdit

	params["user_block"] = userBlock

	params["user_unblock"] = userUnblock

	params["lead_forms_new"] = leadFormsNew

	err = g.SendObjRequest("groups.setCallbackSettings", params, &resp)

	return
}

// Setlongpollsettings - Sets Long Poll notification settings
// Parameters:
//   * groupId - Community ID.
//   * enabled - Sets whether Long Poll is enabled ('0' — disabled, '1' — enabled).
//   * apiVersion - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * messageNew - A new incoming message has been received ('0' — disabled, '1' — enabled).
//   * messageReply - A new outcoming message has been received ('0' — disabled, '1' — enabled).
//   * messageAllow - Allowed messages notifications ('0' — disabled, '1' — enabled).
//   * messageDeny - Denied messages notifications ('0' — disabled, '1' — enabled).
//   * messageEdit - A message has been edited ('0' — disabled, '1' — enabled).
//   * messageTypingState - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * photoNew - New photos notifications ('0' — disabled, '1' — enabled).
//   * audioNew - New audios notifications ('0' — disabled, '1' — enabled).
//   * videoNew - New videos notifications ('0' — disabled, '1' — enabled).
//   * wallReplyNew - New wall replies notifications ('0' — disabled, '1' — enabled).
//   * wallReplyEdit - Wall replies edited notifications ('0' — disabled, '1' — enabled).
//   * wallReplyDelete - A wall comment has been deleted ('0' — disabled, '1' — enabled).
//   * wallReplyRestore - A wall comment has been restored ('0' — disabled, '1' — enabled).
//   * wallPostNew - New wall posts notifications ('0' — disabled, '1' — enabled).
//   * wallRepost - New wall posts notifications ('0' — disabled, '1' — enabled).
//   * boardPostNew - New board posts notifications ('0' — disabled, '1' — enabled).
//   * boardPostEdit - Board posts edited notifications ('0' — disabled, '1' — enabled).
//   * boardPostRestore - Board posts restored notifications ('0' — disabled, '1' — enabled).
//   * boardPostDelete - Board posts deleted notifications ('0' — disabled, '1' — enabled).
//   * photoCommentNew - New comment to photo notifications ('0' — disabled, '1' — enabled).
//   * photoCommentEdit - A photo comment has been edited ('0' — disabled, '1' — enabled).
//   * photoCommentDelete - A photo comment has been deleted ('0' — disabled, '1' — enabled).
//   * photoCommentRestore - A photo comment has been restored ('0' — disabled, '1' — enabled).
//   * videoCommentNew - New comment to video notifications ('0' — disabled, '1' — enabled).
//   * videoCommentEdit - A video comment has been edited ('0' — disabled, '1' — enabled).
//   * videoCommentDelete - A video comment has been deleted ('0' — disabled, '1' — enabled).
//   * videoCommentRestore - A video comment has been restored ('0' — disabled, '1' — enabled).
//   * marketCommentNew - New comment to market item notifications ('0' — disabled, '1' — enabled).
//   * marketCommentEdit - A market comment has been edited ('0' — disabled, '1' — enabled).
//   * marketCommentDelete - A market comment has been deleted ('0' — disabled, '1' — enabled).
//   * marketCommentRestore - A market comment has been restored ('0' — disabled, '1' — enabled).
//   * pollVoteNew - A vote in a public poll has been added ('0' — disabled, '1' — enabled).
//   * groupJoin - Joined community notifications ('0' — disabled, '1' — enabled).
//   * groupLeave - Left community notifications ('0' — disabled, '1' — enabled).
//   * groupChangeSettings - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * groupChangePhoto - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * groupOfficersEdit - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * userBlock - User added to community blacklist
//   * userUnblock - User removed from community blacklist
func (g Groups) Setlongpollsettings(groupId int, enabled bool, apiVersion string, messageNew bool, messageReply bool, messageAllow bool, messageDeny bool, messageEdit bool, messageTypingState bool, photoNew bool, audioNew bool, videoNew bool, wallReplyNew bool, wallReplyEdit bool, wallReplyDelete bool, wallReplyRestore bool, wallPostNew bool, wallRepost bool, boardPostNew bool, boardPostEdit bool, boardPostRestore bool, boardPostDelete bool, photoCommentNew bool, photoCommentEdit bool, photoCommentDelete bool, photoCommentRestore bool, videoCommentNew bool, videoCommentEdit bool, videoCommentDelete bool, videoCommentRestore bool, marketCommentNew bool, marketCommentEdit bool, marketCommentDelete bool, marketCommentRestore bool, pollVoteNew bool, groupJoin bool, groupLeave bool, groupChangeSettings bool, groupChangePhoto bool, groupOfficersEdit bool, userBlock bool, userUnblock bool) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	params["enabled"] = enabled

	if apiVersion != "" {
		params["api_version"] = apiVersion
	}

	params["message_new"] = messageNew

	params["message_reply"] = messageReply

	params["message_allow"] = messageAllow

	params["message_deny"] = messageDeny

	params["message_edit"] = messageEdit

	params["message_typing_state"] = messageTypingState

	params["photo_new"] = photoNew

	params["audio_new"] = audioNew

	params["video_new"] = videoNew

	params["wall_reply_new"] = wallReplyNew

	params["wall_reply_edit"] = wallReplyEdit

	params["wall_reply_delete"] = wallReplyDelete

	params["wall_reply_restore"] = wallReplyRestore

	params["wall_post_new"] = wallPostNew

	params["wall_repost"] = wallRepost

	params["board_post_new"] = boardPostNew

	params["board_post_edit"] = boardPostEdit

	params["board_post_restore"] = boardPostRestore

	params["board_post_delete"] = boardPostDelete

	params["photo_comment_new"] = photoCommentNew

	params["photo_comment_edit"] = photoCommentEdit

	params["photo_comment_delete"] = photoCommentDelete

	params["photo_comment_restore"] = photoCommentRestore

	params["video_comment_new"] = videoCommentNew

	params["video_comment_edit"] = videoCommentEdit

	params["video_comment_delete"] = videoCommentDelete

	params["video_comment_restore"] = videoCommentRestore

	params["market_comment_new"] = marketCommentNew

	params["market_comment_edit"] = marketCommentEdit

	params["market_comment_delete"] = marketCommentDelete

	params["market_comment_restore"] = marketCommentRestore

	params["poll_vote_new"] = pollVoteNew

	params["group_join"] = groupJoin

	params["group_leave"] = groupLeave

	params["group_change_settings"] = groupChangeSettings

	params["group_change_photo"] = groupChangePhoto

	params["group_officers_edit"] = groupOfficersEdit

	params["user_block"] = userBlock

	params["user_unblock"] = userUnblock

	err = g.SendObjRequest("groups.setLongPollSettings", params, &resp)

	return
}

// Unban - NO DESCRIPTION IN JSON SCHEMA
// Parameters:
//   * groupId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * ownerId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (g Groups) Unban(groupId int, ownerId int) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["group_id"] = groupId

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	err = g.SendObjRequest("groups.unban", params, &resp)

	return
}
