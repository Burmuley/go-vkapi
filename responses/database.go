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

// DatabaseGetchair type represents `database_getChairs_response` API response object
type DatabaseGetchair struct {
	Count int                   `json:"count"` // Total number
	Items []*objects.BaseObject `json:"items"`
}

// DatabaseGetcitiesbyid type represents `database_getCitiesById_response` API response object
type DatabaseGetcitiesbyid *objects.BaseObject

// DatabaseGetciti type represents `database_getCities_response` API response object
type DatabaseGetciti struct {
	Count int                     `json:"count"` // Total number
	Items []*objects.DatabaseCity `json:"items"`
}

// DatabaseGetcountriesbyid type represents `database_getCountriesById_response` API response object
type DatabaseGetcountriesbyid *objects.BaseCountry

// DatabaseGetcountri type represents `database_getCountries_response` API response object
type DatabaseGetcountri struct {
	Count int                    `json:"count"` // Total number
	Items []*objects.BaseCountry `json:"items"`
}

// DatabaseGetfaculti type represents `database_getFaculties_response` API response object
type DatabaseGetfaculti struct {
	Count int                        `json:"count"` // Total number
	Items []*objects.DatabaseFaculty `json:"items"`
}

// DatabaseGetmetrostationsbyid type represents `database_getMetroStationsById_response` API response object
type DatabaseGetmetrostationsbyid *objects.DatabaseStation

// DatabaseGetmetrostati type represents `database_getMetroStations_response` API response object
type DatabaseGetmetrostati struct {
	Count int                        `json:"count"` // Total number
	Items []*objects.DatabaseStation `json:"items"`
}

// DatabaseGetregi type represents `database_getRegions_response` API response object
type DatabaseGetregi struct {
	Count int                       `json:"count"` // Total number
	Items []*objects.DatabaseRegion `json:"items"`
}

// DatabaseGetschoolcla type represents `database_getSchoolClasses_response` API response object
type DatabaseGetschoolcla string

// DatabaseGetschool type represents `database_getSchools_response` API response object
type DatabaseGetschool struct {
	Count int                       `json:"count"` // Total number
	Items []*objects.DatabaseSchool `json:"items"`
}

// DatabaseGetuniversiti type represents `database_getUniversities_response` API response object
type DatabaseGetuniversiti struct {
	Count int                           `json:"count"` // Total number
	Items []*objects.DatabaseUniversity `json:"items"`
}
