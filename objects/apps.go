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
// `apps` group of objects
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

// AppsApp type represents `apps_app` API object
type AppsApp struct {
	AppsAppMin      AppsAppMin             `json:"AppsAppMin"`
	AuthorGroup     int                    `json:"author_group"`     // Official community's ID
	AuthorUrl       string                 `json:"author_url"`       // Application author's URL
	Banner1120      string                 `json:"banner_1120"`      // URL of the app banner with 1120 px in width
	Banner560       string                 `json:"banner_560"`       // URL of the app banner with 560 px in width
	CatalogPosition int                    `json:"catalog_position"` // Catalog position
	Description     string                 `json:"description"`      // Application description
	Friends         []int                  `json:"friends"`
	Genre           string                 `json:"genre"`         // Genre name
	GenreId         int                    `json:"genre_id"`      // Genre ID
	International   int                    `json:"international"` // Information whether the application is multilanguage
	IsInCatalog     int                    `json:"is_in_catalog"` // Information whether application is in mobile catalog
	LeaderboardType AppsAppLeaderboardType `json:"leaderboard_type"`
	MembersCount    int                    `json:"members_count"`  // Members number
	PlatformId      int                    `json:"platform_id"`    // Application ID in store
	PublishedDate   int                    `json:"published_date"` // Date when the application has been published in Unixtime
	ScreenName      string                 `json:"screen_name"`    // Screen name
	Screenshots     []PhotosPhoto          `json:"screenshots"`
	Section         string                 `json:"section"` // Application section name
}

// AppsAppLeaderboardType type represents `apps_app_leaderboard_type` API object
type AppsAppLeaderboardType int // Leaderboard type

// AppsAppMin type represents `apps_app_min` API object
type AppsAppMin struct {
	AuthorId int         `json:"author_id"` // Application author's ID
	Icon139  string      `json:"icon_139"`  // URL of the app icon with 139 px in width
	Icon150  string      `json:"icon_150"`  // URL of the app icon with 150 px in width
	Icon278  string      `json:"icon_278"`  // URL of the app icon with 278 px in width
	Icon75   string      `json:"icon_75"`   // URL of the app icon with 75 px in width
	Id       int         `json:"id"`        // Application ID
	Title    string      `json:"title"`     // Application title
	Type     AppsAppType `json:"type"`
}

// AppsAppType type represents `apps_app_type` API object
type AppsAppType string // Application type

// AppsLeaderboard type represents `apps_leaderboard` API object
type AppsLeaderboard struct {
	Level  int `json:"level"`   // Level
	Points int `json:"points"`  // Points number
	Score  int `json:"score"`   // Score number
	UserId int `json:"user_id"` // User ID
}

// AppsScope type represents `apps_scope` API object
type AppsScope struct {
	Name  string `json:"name"`  // Scope name
	Title string `json:"title"` // Scope title
}
