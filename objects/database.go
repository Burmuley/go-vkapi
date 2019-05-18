package objects

/////////////////////////////////////////////////////////////
// Database related API objects	                           //
/////////////////////////////////////////////////////////////

// City represents `database_city` API object
type DatabaseCity struct {
	BaseObject
	Area      string      `json:"area"`
	Region    string      `json:"region"`
	Important BaseBoolInt `json:"important"`
}

// Faculty represents `database_faculty` API object
type DatabaseFaculty struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// Region represents `database_region` API object
type DatabaseRegion struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// School represents `database_school` API object
type DatabaseSchool struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// Station represents `database_station` API object
type DatabaseStation struct {
	ID     int    `json:"id"`
	CityID int    `json:"city_id"`
	Color  string `json:"color"`
	Name   string `json:"name"`
}

// University represents `database_university` API object
type DatabaseUniversity struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}
