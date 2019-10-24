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

package objects

/////////////////////////////////////////////////////////////////////////////////////////////////////////////
// `utils` group of objects
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

// UtilsDomainResolved type represents `utils_domain_resolved` API object
type UtilsDomainResolved struct {
	ObjectId int                      `json:"object_id"` // Object ID
	Type     *UtilsDomainResolvedType `json:"type"`
}

// UtilsDomainResolvedType type represents `utils_domain_resolved_type` API object
type UtilsDomainResolvedType string // Object type

// UtilsLastShortenedLink type represents `utils_last_shortened_link` API object
type UtilsLastShortenedLink struct {
	AccessKey string `json:"access_key"` // Access key for private stats
	Key       string `json:"key"`        // Link key (characters after vk.cc/)
	ShortUrl  string `json:"short_url"`  // Short link URL
	Timestamp int    `json:"timestamp"`  // Creation time in Unixtime
	Url       string `json:"url"`        // Full URL
	Views     int    `json:"views"`      // Total views number
}

// UtilsLinkChecked type represents `utils_link_checked` API object
type UtilsLinkChecked struct {
	Link   string                  `json:"link"` // Link URL
	Status *UtilsLinkCheckedStatus `json:"status"`
}

// UtilsLinkCheckedStatus type represents `utils_link_checked_status` API object
type UtilsLinkCheckedStatus string // Link status

// UtilsLinkStats type represents `utils_link_stats` API object
type UtilsLinkStats struct {
	Key   string        `json:"key"` // Link key (characters after vk.cc/)
	Stats []*UtilsStats `json:"stats"`
}

// UtilsLinkStatsExtended type represents `utils_link_stats_extended` API object
type UtilsLinkStatsExtended struct {
	Key   string                `json:"key"` // Link key (characters after vk.cc/)
	Stats []*UtilsStatsExtended `json:"stats"`
}

// UtilsShortLink type represents `utils_short_link` API object
type UtilsShortLink struct {
	AccessKey string `json:"access_key"` // Access key for private stats
	Key       string `json:"key"`        // Link key (characters after vk.cc/)
	ShortUrl  string `json:"short_url"`  // Short link URL
	Url       string `json:"url"`        // Full URL
}

// UtilsStats type represents `utils_stats` API object
type UtilsStats struct {
	Timestamp int `json:"timestamp"` // Start time
	Views     int `json:"views"`     // Total views number
}

// UtilsStatsCity type represents `utils_stats_city` API object
type UtilsStatsCity struct {
	CityId int `json:"city_id"` // City ID
	Views  int `json:"views"`   // Views number
}

// UtilsStatsCountry type represents `utils_stats_country` API object
type UtilsStatsCountry struct {
	CountryId int `json:"country_id"` // Country ID
	Views     int `json:"views"`      // Views number
}

// UtilsStatsExtended type represents `utils_stats_extended` API object
type UtilsStatsExtended struct {
	Cities    []*UtilsStatsCity    `json:"cities"`
	Countries []*UtilsStatsCountry `json:"countries"`
	SexAge    []*UtilsStatsSexAge  `json:"sex_age"`
	Timestamp int                  `json:"timestamp"` // Start time
	Views     int                  `json:"views"`     // Total views number
}

// UtilsStatsSexAge type represents `utils_stats_sex_age` API object
type UtilsStatsSexAge struct {
	AgeRange string `json:"age_range"` // Age denotation
	Female   int    `json:"female"`    //  Views by female users
	Male     int    `json:"male"`      //  Views by male users
}
