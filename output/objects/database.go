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
// Code generator location: https://github.com/Burmuley/go-vkapi-gen                                       //
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

package objects

/////////////////////////////////////////////////////////////////////////////////////////////////////////////
// `database` group of objects
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

// DatabaseCity type represents `database_city` API object
type DatabaseCity struct {
	Area      string      `json:"area"`      // Area title
	Id        int         `json:"id"`        // Object ID
	Important BaseBoolInt `json:"important"` // Information whether the city is included in important cities list
	Region    string      `json:"region"`    // Region title
	Title     string      `json:"title"`     // Object title
}

// DatabaseFaculty type represents `database_faculty` API object
type DatabaseFaculty struct {
	Id    int    `json:"id"`    // Faculty ID
	Title string `json:"title"` // Faculty title
}

// DatabaseRegion type represents `database_region` API object
type DatabaseRegion struct {
	Id    int    `json:"id"`    // Region ID
	Title string `json:"title"` // Region title
}

// DatabaseSchool type represents `database_school` API object
type DatabaseSchool struct {
	Id    int    `json:"id"`    // School ID
	Title string `json:"title"` // School title
}

// DatabaseStation type represents `database_station` API object
type DatabaseStation struct {
	CityId int    `json:"city_id"` // City ID
	Color  string `json:"color"`   // Hex color code without #
	Id     int    `json:"id"`      // Station ID
	Name   string `json:"name"`    // Station name
}

// DatabaseUniversity type represents `database_university` API object
type DatabaseUniversity struct {
	Id    int    `json:"id"`    // University ID
	Title string `json:"title"` // University title
}
