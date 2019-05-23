package objects

// DomainResolvedType represents `utils_domain_resolved_type` API object
type UtilsDomainResolvedType string

func (d *UtilsDomainResolvedType) MarshalJSON() ([]byte, error) {
	return GetStringFromRange(string(*d), "user", "group", "application", "page")
}

func (d *UtilsDomainResolvedType) String() string {
	return string(*d)
}

// DomainResolved represents `utils_domain_resolved` API object
type UtilsDomainResolved struct {
	ObjectID int                     `json:"object_id"`
	Type     UtilsDomainResolvedType `json:"type"`
}

// LastShortenedLink represents `utils_last_shortened_link` API object
type UtilsLastShortenedLink struct {
	AccessKey string `json:"access_key"`
	Key       string `json:"key"`
	ShortURL  string `json:"short_url"`
	Timestamp int    `json:"timestamp"`
	URL       string `json:"url"`
	Views     int    `json:"views"`
}

// LinkCheckedStatus represents `utils_link_checked_status` API object
type UtilsLinkCheckedStatus string

func (l *UtilsLinkCheckedStatus) MarshalJSON() ([]byte, error) {
	return GetStringFromRange(string(*l), "banned", "not_banned", "processing")
}

func (l *UtilsLinkCheckedStatus) String() string {
	return string(*l)
}

// LinkChecked represents `utils_link_checked` API object
type UtilsLinkChecked struct {
	Link   string                 `json:"link"`
	Status UtilsLinkCheckedStatus `json:"status"`
}

// LinkStats represents `utils_link_stats` API object
type UtilsLinkStats struct {
	Key   string       `json:"key"` //Link key (characters after vk.cc/)
	Stats []UtilsStats `json:"stats"`
}

// LinkStatsExtended represents `utils_link_stats_extended` API object
type UtilsLinkStatsExtended struct {
	Key   string               `json:"key"` //Link key (characters after vk.cc/)
	Stats []UtilsStatsExtended `json:"stats"`
}

// ShortLink represents `utils_short_link` API object
type UtilsShortLink struct {
	AccessKey string `json:"access_key"`
	Key       string `json:"key"`
	ShortURL  string `json:"short_url"`
	URL       string `json:"url"`
}

// Stats represents `utils_stats` API object
type UtilsStats struct {
	Timestamp int `json:"timestamp"`
	Views     int `json:"views"`
}

// StatsExtended represents `utils_stats_extended` API object
type UtilsStatsExtended struct {
	Cities    []StatsCity    `json:"cities"`
	Countries []StatsCountry `json:"countries"`
	SexAge    []StatsSexAge  `json:"sex_age"`
	Timestamp int            `json:"timestamp"`
	Views     int            `json:"views"`
}

// StatsSexAge represents `utils_stats_sex_age` API object
type UtilsStatsSexAge struct {
	AgeRange string `json:"age_range"`
	Female   int    `json:"female"`
	Male     int    `json:"males"`
}

//StatsCountry represents `utils_stats_country` API object
type UtilsStatsCountry struct {
	CountryID int `json:"country_id"`
	Views     int `json:"views"`
}

//StatsCity represents `utils_stats_city` API object
type UtilsStatsCity struct {
	CountryID int `json:"city_id"`
	Views     int `json:"views"`
}
