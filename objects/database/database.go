package database

import "gitlab.com/Burmuley/go-vkapi/objects/base"

/////////////////////////////////////////////////////////////
// Database related API objects	                           //
/////////////////////////////////////////////////////////////

// City represents `database_city` API object
type City struct {
	base.Object
	Area      string       `json:"area"`
	Region    string       `json:"region"`
	Important base.BoolInt `json:"important"`
}

// Faculty represents `database_faculty` API object
type Faculty struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// Region represents `database_region` API object
type Region struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// School represents `database_school` API object
type School struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// Station represents `database_station` API object
type Station struct {
	ID     int    `json:"id"`
	CityID int    `json:"city_id"`
	Color  string `json:"color"`
	Name   string `json:"name"`
}

// University represents `database_university` API object
type University struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}
