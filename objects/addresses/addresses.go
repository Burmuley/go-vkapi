package addresses

import (
	"gitlab.com/Burmuley/go-vkapi/objects"
)

/////////////////////////////////////////////////////////////
// Addresses related API objects                           //
/////////////////////////////////////////////////////////////

type Fields string

func (f *Fields) MarshalJSON() ([]byte, error) {
	return objects.GetStringFromRange(string(*f), "id", "title", "address", "additional_address",
		"country_id", "city_id", "metro_station_id", "latitude", "longitude", "distance", "work_info_status",
		"timetable", "phone", "time_offset")
}

func (f *Fields) GetName() string {
	return string(*f)
}
