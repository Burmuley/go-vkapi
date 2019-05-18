package objects

/////////////////////////////////////////////////////////////
// Addresses related API objects                           //
/////////////////////////////////////////////////////////////

type AddressesFields string

func (f *AddressesFields) MarshalJSON() ([]byte, error) {
	return GetStringFromRange(string(*f), "id", "title", "address", "additional_address",
		"country_id", "city_id", "metro_station_id", "latitude", "longitude", "distance", "work_info_status",
		"timetable", "phone", "time_offset")
}

func (f *AddressesFields) GetName() string {
	return string(*f)
}
