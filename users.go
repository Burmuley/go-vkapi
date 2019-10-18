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

type Users struct {
	*VKApi
}

/////////////////////////////////////////////////////////////
// `Users` methods
/////////////////////////////////////////////////////////////

// Get - Returns detailed information on users.
// Parameters:
//   * userIds - User IDs or screen names ('screen_name'). By default, current user ID.
//   * fields - Profile fields to return. Sample values: 'nickname', 'screen_name', 'sex', 'bdate' (birthdate), 'city', 'country', 'timezone', 'photo', 'photo_medium', 'photo_big', 'has_mobile', 'contacts', 'education', 'online', 'counters', 'relation', 'last_seen', 'activity', 'can_write_private_message', 'can_see_all_posts', 'can_post', 'universities',
//   * nameCase - Case for declension of user name and surname: 'nom' — nominative (default), 'gen' — genitive , 'dat' — dative, 'acc' — accusative , 'ins' — instrumental , 'abl' — prepositional
func (u Users) Get(userIds []string, fields []objects.UsersFields, nameCase string) (resp responses.UsersGet, err error) {
	params := map[string]interface{}{}

	if len(userIds) > 0 {
		params["user_ids"] = SliceToString(userIds)
	}

	if len(fields) > 0 {
		params["fields"] = SliceToString(fields)
	}

	if nameCase != "" {
		params["name_case"] = nameCase
	}

	err = u.SendObjRequest("users.get", params, &resp)

	return
}

// GetFollowers - Returns a list of IDs of followers of the user in question, sorted by date added, most recent first.
// Parameters:
//   * userId - User ID.
//   * offset - Offset needed to return a specific subset of followers.
//   * count - Number of followers to return.
//   * fields - Profile fields to return. Sample values: 'nickname', 'screen_name', 'sex', 'bdate' (birthdate), 'city', 'country', 'timezone', 'photo', 'photo_medium', 'photo_big', 'has_mobile', 'rate', 'contacts', 'education', 'online'.
//   * nameCase - Case for declension of user name and surname: 'nom' — nominative (default), 'gen' — genitive , 'dat' — dative, 'acc' — accusative , 'ins' — instrumental , 'abl' — prepositional
func (u Users) GetFollowers(userId int, offset int, count int, fields []objects.UsersFields, nameCase string) (resp responses.UsersGetFollowers, err error) {
	params := map[string]interface{}{}

	if userId > 0 {
		params["user_id"] = userId
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

	if nameCase != "" {
		params["name_case"] = nameCase
	}

	err = u.SendObjRequest("users.getFollowers", params, &resp)

	return
}

// GetSubscriptions - Returns a list of IDs of users and communities followed by the user.
// Parameters:
//   * userId - User ID.
//   * extended - '1' — to return a combined list of users and communities, '0' — to return separate lists of users and communities (default)
//   * offset - Offset needed to return a specific subset of subscriptions.
//   * count - Number of users and communities to return.
//   * fields - NO DESCRIPTION IN JSON SCHEMA
func (u Users) GetSubscriptions(userId int, offset int, count int, fields []objects.UsersFields) (resp responses.UsersGetSubscriptions, err error) {
	params := map[string]interface{}{}
	params["extended"] = "0"

	if userId > 0 {
		params["user_id"] = userId
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

	err = u.SendObjRequest("users.getSubscriptions", params, &resp)

	return
}

// GetSubscriptionsExtended - Returns a list of IDs of users and communities followed by the user.
// Parameters:
//   * userId - User ID.
//   * extended - '1' — to return a combined list of users and communities, '0' — to return separate lists of users and communities (default)
//   * offset - Offset needed to return a specific subset of subscriptions.
//   * count - Number of users and communities to return.
//   * fields - NO DESCRIPTION IN JSON SCHEMA
func (u Users) GetSubscriptionsExtended(userId int, offset int, count int, fields []objects.UsersFields) (resp responses.UsersGetSubscriptionsExtended, err error) {
	params := map[string]interface{}{}
	params["extended"] = "1"

	if userId > 0 {
		params["user_id"] = userId
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

	err = u.SendObjRequest("users.getSubscriptions", params, &resp)

	return
}

// IsAppUser - Returns information whether a user installed the application.
// Parameters:
//   * userId - NO DESCRIPTION IN JSON SCHEMA
func (u Users) IsAppUser(userId int) (resp responses.UsersIsAppUser, err error) {
	params := map[string]interface{}{}

	if userId > 0 {
		params["user_id"] = userId
	}

	err = u.SendObjRequest("users.isAppUser", params, &resp)

	return
}

// Report - Reports (submits a complain about) a user.
// Parameters:
//   * userId - ID of the user about whom a complaint is being made.
//   * pType - Type of complaint: 'porn' – pornography, 'spam' – spamming, 'insult' – abusive behavior, 'advertisement' – disruptive advertisements
//   * comment - Comment describing the complaint.
func (u Users) Report(userId int, pType string, comment string) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["user_id"] = userId

	params["type"] = pType

	if comment != "" {
		params["comment"] = comment
	}

	err = u.SendObjRequest("users.report", params, &resp)

	return
}

// Search - Returns a list of users matching the search criteria.
// Parameters:
//   * q - Search query string (e.g., 'Vasya Babich').
//   * sort - Sort order: '1' — by date registered, '0' — by rating
//   * offset - Offset needed to return a specific subset of users.
//   * count - Number of users to return.
//   * fields - Profile fields to return. Sample values: 'nickname', 'screen_name', 'sex', 'bdate' (birthdate), 'city', 'country', 'timezone', 'photo', 'photo_medium', 'photo_big', 'has_mobile', 'rate', 'contacts', 'education', 'online',
//   * city - City ID.
//   * country - Country ID.
//   * hometown - City name in a string.
//   * universityCountry - ID of the country where the user graduated.
//   * university - ID of the institution of higher education.
//   * universityYear - Year of graduation from an institution of higher education.
//   * universityFaculty - Faculty ID.
//   * universityChair - Chair ID.
//   * sex - '1' — female, '2' — male, '0' — any (default)
//   * status - Relationship status: '1' — Not married, '2' — In a relationship, '3' — Engaged, '4' — Married, '5' — It's complicated, '6' — Actively searching, '7' — In love
//   * ageFrom - Minimum age.
//   * ageTo - Maximum age.
//   * birthDay - Day of birth.
//   * birthMonth - Month of birth.
//   * birthYear - Year of birth.
//   * online - '1' — online only, '0' — all users
//   * hasPhoto - '1' — with photo only, '0' — all users
//   * schoolCountry - ID of the country where users finished school.
//   * schoolCity - ID of the city where users finished school.
//   * schoolClass - NO DESCRIPTION IN JSON SCHEMA
//   * school - ID of the school.
//   * schoolYear - School graduation year.
//   * religion - Users' religious affiliation.
//   * interests - Users' interests.
//   * company - Name of the company where users work.
//   * position - Job position.
//   * groupId - ID of a community to search in communities.
//   * fromList - NO DESCRIPTION IN JSON SCHEMA
func (u Users) Search(q string, sort int, offset int, count int, fields []objects.UsersFields, city int, country int, hometown string, universityCountry int, university int, universityYear int, universityFaculty int, universityChair int, sex int, status int, ageFrom int, ageTo int, birthDay int, birthMonth int, birthYear int, online bool, hasPhoto bool, schoolCountry int, schoolCity int, schoolClass int, school int, schoolYear int, religion string, interests string, company string, position string, groupId int, fromList []string) (resp responses.UsersSearch, err error) {
	params := map[string]interface{}{}

	if q != "" {
		params["q"] = q
	}

	if sort > 0 {
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

	if city > 0 {
		params["city"] = city
	}

	if country > 0 {
		params["country"] = country
	}

	if hometown != "" {
		params["hometown"] = hometown
	}

	if universityCountry > 0 {
		params["university_country"] = universityCountry
	}

	if university > 0 {
		params["university"] = university
	}

	if universityYear > 0 {
		params["university_year"] = universityYear
	}

	if universityFaculty > 0 {
		params["university_faculty"] = universityFaculty
	}

	if universityChair > 0 {
		params["university_chair"] = universityChair
	}

	if sex > 0 {
		params["sex"] = sex
	}

	if status > 0 {
		params["status"] = status
	}

	if ageFrom > 0 {
		params["age_from"] = ageFrom
	}

	if ageTo > 0 {
		params["age_to"] = ageTo
	}

	if birthDay > 0 {
		params["birth_day"] = birthDay
	}

	if birthMonth > 0 {
		params["birth_month"] = birthMonth
	}

	if birthYear > 0 {
		params["birth_year"] = birthYear
	}

	params["online"] = online

	params["has_photo"] = hasPhoto

	if schoolCountry > 0 {
		params["school_country"] = schoolCountry
	}

	if schoolCity > 0 {
		params["school_city"] = schoolCity
	}

	if schoolClass > 0 {
		params["school_class"] = schoolClass
	}

	if school > 0 {
		params["school"] = school
	}

	if schoolYear > 0 {
		params["school_year"] = schoolYear
	}

	if religion != "" {
		params["religion"] = religion
	}

	if interests != "" {
		params["interests"] = interests
	}

	if company != "" {
		params["company"] = company
	}

	if position != "" {
		params["position"] = position
	}

	if groupId > 0 {
		params["group_id"] = groupId
	}

	if len(fromList) > 0 {
		params["from_list"] = SliceToString(fromList)
	}

	err = u.SendObjRequest("users.search", params, &resp)

	return
}
