package addresses

import "fmt"

/////////////////////////////////////////////////////////////
// Addresses related API objects                           //
/////////////////////////////////////////////////////////////

type Fields string

func (f *Fields) MarshalJSON() ([]byte, error) {
	switch *f {
	case "id",
		"title",
		"address",
		"additional_address",
		"country_id",
		"city_id",
		"metro_station_id",
		"latitude",
		"longitude",
		"distance",
		"work_info_status",
		"timetable",
		"phone",
		"time_offset":
		return []byte(*f), nil
	}

	return []byte{}, fmt.Errorf("value is not in allowed range")
}

func (f *Fields) GetName() string {
	return string(*f)
}
