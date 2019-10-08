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
// `database` group of responses
/////////////////////////////////////////////////////////////

// DatabaseGetFaculties type represents `database_getFaculties_response` API response object
type DatabaseGetFaculties struct {
	Count int                       `json:"count"` // Total number
	Items []objects.DatabaseFaculty `json:"items"`
}

// DatabaseGetMetroStations type represents `database_getMetroStations_response` API response object
type DatabaseGetMetroStations struct {
	Count int                       `json:"count"` // Total number
	Items []objects.DatabaseStation `json:"items"`
}

// DatabaseGetCountriesById type represents `database_getCountriesById_response` API response object
type DatabaseGetCountriesById objects.BaseCountry

// DatabaseGetSchools type represents `database_getSchools_response` API response object
type DatabaseGetSchools struct {
	Count int                      `json:"count"` // Total number
	Items []objects.DatabaseSchool `json:"items"`
}

// DatabaseGetChairs type represents `database_getChairs_response` API response object
type DatabaseGetChairs struct {
	Count int                  `json:"count"` // Total number
	Items []objects.BaseObject `json:"items"`
}

// DatabaseGetRegions type represents `database_getRegions_response` API response object
type DatabaseGetRegions struct {
	Count int                      `json:"count"` // Total number
	Items []objects.DatabaseRegion `json:"items"`
}

// DatabaseGetCities type represents `database_getCities_response` API response object
type DatabaseGetCities struct {
	Count int                    `json:"count"` // Total number
	Items []objects.DatabaseCity `json:"items"`
}

// DatabaseGetUniversities type represents `database_getUniversities_response` API response object
type DatabaseGetUniversities struct {
	Count int                          `json:"count"` // Total number
	Items []objects.DatabaseUniversity `json:"items"`
}

// DatabaseGetCitiesById type represents `database_getCitiesById_response` API response object
type DatabaseGetCitiesById objects.BaseObject

// DatabaseGetMetroStationsById type represents `database_getMetroStationsById_response` API response object
type DatabaseGetMetroStationsById objects.DatabaseStation

// DatabaseGetCountries type represents `database_getCountries_response` API response object
type DatabaseGetCountries struct {
	Count int                   `json:"count"` // Total number
	Items []objects.BaseCountry `json:"items"`
}

// DatabaseGetSchoolClasses type represents `database_getSchoolClasses_response` API response object
type DatabaseGetSchoolClasses multiple
