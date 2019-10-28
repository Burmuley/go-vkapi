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
// `database` group of responses
/////////////////////////////////////////////////////////////

// DatabaseGetchairs type represents `database_getChairs_response` API response object
type DatabaseGetchairs struct {
	Count int                  `json:"count"` // Total number
	Items []objects.BaseObject `json:"items"`
}

// DatabaseGetcitiesbyid type represents `database_getCitiesById_response` API response object
type DatabaseGetcitiesbyid objects.BaseObject

// DatabaseGetcities type represents `database_getCities_response` API response object
type DatabaseGetcities struct {
	Count int                    `json:"count"` // Total number
	Items []objects.DatabaseCity `json:"items"`
}

// DatabaseGetcountriesbyid type represents `database_getCountriesById_response` API response object
type DatabaseGetcountriesbyid objects.BaseCountry

// DatabaseGetcountries type represents `database_getCountries_response` API response object
type DatabaseGetcountries struct {
	Count int                   `json:"count"` // Total number
	Items []objects.BaseCountry `json:"items"`
}

// DatabaseGetfaculties type represents `database_getFaculties_response` API response object
type DatabaseGetfaculties struct {
	Count int                       `json:"count"` // Total number
	Items []objects.DatabaseFaculty `json:"items"`
}

// DatabaseGetmetrostationsbyid type represents `database_getMetroStationsById_response` API response object
type DatabaseGetmetrostationsbyid objects.DatabaseStation

// DatabaseGetmetrostations type represents `database_getMetroStations_response` API response object
type DatabaseGetmetrostations struct {
	Count int                       `json:"count"` // Total number
	Items []objects.DatabaseStation `json:"items"`
}

// DatabaseGetregions type represents `database_getRegions_response` API response object
type DatabaseGetregions struct {
	Count int                      `json:"count"` // Total number
	Items []objects.DatabaseRegion `json:"items"`
}

// DatabaseGetschoolclasses type represents `database_getSchoolClasses_response` API response object
type DatabaseGetschoolclasses string

// DatabaseGetschools type represents `database_getSchools_response` API response object
type DatabaseGetschools struct {
	Count int                      `json:"count"` // Total number
	Items []objects.DatabaseSchool `json:"items"`
}

// DatabaseGetuniversities type represents `database_getUniversities_response` API response object
type DatabaseGetuniversities struct {
	Count int                          `json:"count"` // Total number
	Items []objects.DatabaseUniversity `json:"items"`
}
