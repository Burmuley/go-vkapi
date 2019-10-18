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

// GetChairs - Returns list of chairs on a specified faculty.
// Parameters:
//   * facultyId - id of the faculty to get chairs from
//   * offset - offset required to get a certain subset of chairs
//   * count - amount of chairs to get
func (d Database) GetChairs(facultyId int, offset int, count int) (resp responses.DatabaseGetChairs, err error) {
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

// GetCities - Returns a list of cities.
// Parameters:
//   * countryId - Country ID.
//   * regionId - Region ID.
//   * q - Search query.
//   * needAll - '1' — to return all cities in the country, '0' — to return major cities in the country (default),
//   * offset - Offset needed to return a specific subset of cities.
//   * count - Number of cities to return.
func (d Database) GetCities(countryId int, regionId int, q string, needAll bool, offset int, count int) (resp responses.DatabaseGetCities, err error) {
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

// GetCitiesById - Returns information about cities by their IDs.
// Parameters:
//   * cityIds - City IDs.
func (d Database) GetCitiesById(cityIds []int) (resp responses.DatabaseGetCitiesById, err error) {
	params := map[string]interface{}{}

	if len(cityIds) > 0 {
		params["city_ids"] = SliceToString(cityIds)
	}

	err = d.SendObjRequest("database.getCitiesById", params, &resp)

	return
}

// GetCountries - Returns a list of countries.
// Parameters:
//   * needAll - '1' — to return a full list of all countries, '0' — to return a list of countries near the current user's country (default).
//   * code - Country codes in [vk.com/dev/country_codes|ISO 3166-1 alpha-2] standard.
//   * offset - Offset needed to return a specific subset of countries.
//   * count - Number of countries to return.
func (d Database) GetCountries(needAll bool, code string, offset int, count int) (resp responses.DatabaseGetCountries, err error) {
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

// GetCountriesById - Returns information about countries by their IDs.
// Parameters:
//   * countryIds - Country IDs.
func (d Database) GetCountriesById(countryIds []int) (resp responses.DatabaseGetCountriesById, err error) {
	params := map[string]interface{}{}

	if len(countryIds) > 0 {
		params["country_ids"] = SliceToString(countryIds)
	}

	err = d.SendObjRequest("database.getCountriesById", params, &resp)

	return
}

// GetFaculties - Returns a list of faculties (i.e., university departments).
// Parameters:
//   * universityId - University ID.
//   * offset - Offset needed to return a specific subset of faculties.
//   * count - Number of faculties to return.
func (d Database) GetFaculties(universityId int, offset int, count int) (resp responses.DatabaseGetFaculties, err error) {
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

// GetMetroStations - Get metro stations by city
// Parameters:
//   * cityId - NO DESCRIPTION IN JSON SCHEMA
//   * offset - NO DESCRIPTION IN JSON SCHEMA
//   * count - NO DESCRIPTION IN JSON SCHEMA
//   * extended - NO DESCRIPTION IN JSON SCHEMA
func (d Database) GetMetroStations(cityId int, offset int, count int) (resp responses.DatabaseGetMetroStations, err error) {
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

// GetMetroStationsById - Get metro station by his id
// Parameters:
//   * stationIds - NO DESCRIPTION IN JSON SCHEMA
func (d Database) GetMetroStationsById(stationIds []int) (resp responses.DatabaseGetMetroStationsById, err error) {
	params := map[string]interface{}{}

	if len(stationIds) > 0 {
		params["station_ids"] = SliceToString(stationIds)
	}

	err = d.SendObjRequest("database.getMetroStationsById", params, &resp)

	return
}

// GetRegions - Returns a list of regions.
// Parameters:
//   * countryId - Country ID, received in [vk.com/dev/database.getCountries|database.getCountries] method.
//   * q - Search query.
//   * offset - Offset needed to return specific subset of regions.
//   * count - Number of regions to return.
func (d Database) GetRegions(countryId int, q string, offset int, count int) (resp responses.DatabaseGetRegions, err error) {
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

// GetSchoolClasses - Returns a list of school classes specified for the country.
// Parameters:
//   * countryId - Country ID.
func (d Database) GetSchoolClasses(countryId int) (resp responses.DatabaseGetSchoolClasses, err error) {
	params := map[string]interface{}{}

	if countryId > 0 {
		params["country_id"] = countryId
	}

	err = d.SendObjRequest("database.getSchoolClasses", params, &resp)

	return
}

// GetSchools - Returns a list of schools.
// Parameters:
//   * q - Search query.
//   * cityId - City ID.
//   * offset - Offset needed to return a specific subset of schools.
//   * count - Number of schools to return.
func (d Database) GetSchools(q string, cityId int, offset int, count int) (resp responses.DatabaseGetSchools, err error) {
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

// GetUniversities - Returns a list of higher education institutions.
// Parameters:
//   * q - Search query.
//   * countryId - Country ID.
//   * cityId - City ID.
//   * offset - Offset needed to return a specific subset of universities.
//   * count - Number of universities to return.
func (d Database) GetUniversities(q string, countryId int, cityId int, offset int, count int) (resp responses.DatabaseGetUniversities, err error) {
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
