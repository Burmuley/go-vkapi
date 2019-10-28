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

type Database struct {
	*VKApi
}

/////////////////////////////////////////////////////////////
// `Database` methods
/////////////////////////////////////////////////////////////

// Getchairs - Returns list of chairs on a specified faculty.
// Parameters:
//   * facultyId - id of the faculty to get chairs from
//   * offset - offset required to get a certain subset of chairs
//   * count - amount of chairs to get
func (d Database) Getchairs(facultyId int, offset int, count int) (resp responses.DatabaseGetchairs, err error) {
	params := map[string]interface{}{}

	params["faculty_id"] = facultyId

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	err = d.SendObjRequest("database.getChairs", params, &resp)

	return
}

// Getcities - Returns a list of cities.
// Parameters:
//   * countryId - Country ID.
//   * regionId - Region ID.
//   * q - Search query.
//   * needAll - '1' — to return all cities in the country, '0' — to return major cities in the country (default),
//   * offset - Offset needed to return a specific subset of cities.
//   * count - Number of cities to return.
func (d Database) Getcities(countryId int, regionId int, q string, needAll bool, offset int, count int) (resp responses.DatabaseGetcities, err error) {
	params := map[string]interface{}{}

	params["country_id"] = countryId

	if regionId > 0 {
		params["region_id"] = regionId
	}

	if q != "" {
		params["q"] = q
	}

	params["need_all"] = needAll

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	err = d.SendObjRequest("database.getCities", params, &resp)

	return
}

// Getcitiesbyid - Returns information about cities by their IDs.
// Parameters:
//   * cityIds - City IDs.
func (d Database) Getcitiesbyid(cityIds []int) (resp responses.DatabaseGetcitiesbyid, err error) {
	params := map[string]interface{}{}

	if len(cityIds) > 0 {
		params["city_ids"] = SliceToString(cityIds)
	}

	err = d.SendObjRequest("database.getCitiesById", params, &resp)

	return
}

// Getcountries - Returns a list of countries.
// Parameters:
//   * needAll - '1' — to return a full list of all countries, '0' — to return a list of countries near the current user's country (default).
//   * code - Country codes in [vk.com/dev/country_codes|ISO 3166-1 alpha-2] standard.
//   * offset - Offset needed to return a specific subset of countries.
//   * count - Number of countries to return.
func (d Database) Getcountries(needAll bool, code string, offset int, count int) (resp responses.DatabaseGetcountries, err error) {
	params := map[string]interface{}{}

	params["need_all"] = needAll

	if code != "" {
		params["code"] = code
	}

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	err = d.SendObjRequest("database.getCountries", params, &resp)

	return
}

// Getcountriesbyid - Returns information about countries by their IDs.
// Parameters:
//   * countryIds - Country IDs.
func (d Database) Getcountriesbyid(countryIds []int) (resp responses.DatabaseGetcountriesbyid, err error) {
	params := map[string]interface{}{}

	if len(countryIds) > 0 {
		params["country_ids"] = SliceToString(countryIds)
	}

	err = d.SendObjRequest("database.getCountriesById", params, &resp)

	return
}

// Getfaculties - Returns a list of faculties (i.e., university departments).
// Parameters:
//   * universityId - University ID.
//   * offset - Offset needed to return a specific subset of faculties.
//   * count - Number of faculties to return.
func (d Database) Getfaculties(universityId int, offset int, count int) (resp responses.DatabaseGetfaculties, err error) {
	params := map[string]interface{}{}

	params["university_id"] = universityId

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	err = d.SendObjRequest("database.getFaculties", params, &resp)

	return
}

// Getmetrostations - Get metro stations by city
// Parameters:
//   * cityId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * offset - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * count - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * extended - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (d Database) Getmetrostations(cityId int, offset int, count int) (resp responses.DatabaseGetmetrostations, err error) {
	params := map[string]interface{}{}

	params["city_id"] = cityId

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	err = d.SendObjRequest("database.getMetroStations", params, &resp)

	return
}

// Getmetrostationsbyid - Get metro station by his id
// Parameters:
//   * stationIds - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (d Database) Getmetrostationsbyid(stationIds []int) (resp responses.DatabaseGetmetrostationsbyid, err error) {
	params := map[string]interface{}{}

	if len(stationIds) > 0 {
		params["station_ids"] = SliceToString(stationIds)
	}

	err = d.SendObjRequest("database.getMetroStationsById", params, &resp)

	return
}

// Getregions - Returns a list of regions.
// Parameters:
//   * countryId - Country ID, received in [vk.com/dev/database.getCountries|database.getCountries] method.
//   * q - Search query.
//   * offset - Offset needed to return specific subset of regions.
//   * count - Number of regions to return.
func (d Database) Getregions(countryId int, q string, offset int, count int) (resp responses.DatabaseGetregions, err error) {
	params := map[string]interface{}{}

	params["country_id"] = countryId

	if q != "" {
		params["q"] = q
	}

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	err = d.SendObjRequest("database.getRegions", params, &resp)

	return
}

// Getschoolclasses - Returns a list of school classes specified for the country.
// Parameters:
//   * countryId - Country ID.
func (d Database) Getschoolclasses(countryId int) (resp responses.DatabaseGetschoolclasses, err error) {
	params := map[string]interface{}{}

	if countryId > 0 {
		params["country_id"] = countryId
	}

	err = d.SendObjRequest("database.getSchoolClasses", params, &resp)

	return
}

// Getschools - Returns a list of schools.
// Parameters:
//   * q - Search query.
//   * cityId - City ID.
//   * offset - Offset needed to return a specific subset of schools.
//   * count - Number of schools to return.
func (d Database) Getschools(q string, cityId int, offset int, count int) (resp responses.DatabaseGetschools, err error) {
	params := map[string]interface{}{}

	if q != "" {
		params["q"] = q
	}

	params["city_id"] = cityId

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	err = d.SendObjRequest("database.getSchools", params, &resp)

	return
}

// Getuniversities - Returns a list of higher education institutions.
// Parameters:
//   * q - Search query.
//   * countryId - Country ID.
//   * cityId - City ID.
//   * offset - Offset needed to return a specific subset of universities.
//   * count - Number of universities to return.
func (d Database) Getuniversities(q string, countryId int, cityId int, offset int, count int) (resp responses.DatabaseGetuniversities, err error) {
	params := map[string]interface{}{}

	if q != "" {
		params["q"] = q
	}

	if countryId > 0 {
		params["country_id"] = countryId
	}

	if cityId > 0 {
		params["city_id"] = cityId
	}

	if offset > 0 {
		params["offset"] = offset
	}

	if count > 0 {
		params["count"] = count
	}

	err = d.SendObjRequest("database.getUniversities", params, &resp)

	return
}
