package utils

import (
	"gitlab.com/Burmuley/go-vkapi/objects"
)

// DomainResolvedType represents `utils_domain_resolved_type` API object
type DomainResolvedType string

func (d *DomainResolvedType) MarshalJSON() ([]byte, error) {
	return objects.GetStringFromRange(string(*d), "user", "group", "application", "page")
}

func (d *DomainResolvedType) GetName() string {
	return string(*d)
}

// DomainResolved represents `utils_domain_resolved` API object
type DomainResolved struct {
	ObjectID int                `json:"object_id"`
	Type     DomainResolvedType `json:"type"`
}

// LastShortenedLink represents `utils_last_shortened_link` API object
type LastShortenedLink struct {
	AccessKey string `json:"access_key"`
	Key       string `json:"key"`
	ShortURL  string `json:"short_url"`
	Timestamp int    `json:"timestamp"`
	URL       string `json:"url"`
	Views     int    `json:"views"`
}

// LinkCheckedStatus represents `utils_link_checked_status` API object
type LinkCheckedStatus string

func (l *LinkCheckedStatus) MarshalJSON() ([]byte, error) {
	return objects.GetStringFromRange(string(*l), "banned", "not_banned", "processing")
}

func (l *LinkCheckedStatus) GetName() string {
	return string(*l)
}

// LinkChecked represents `utils_link_checked` API object
type LinkChecked struct {
	Link   string            `json:"link"`
	Status LinkCheckedStatus `json:"status"`
}

// LinkStats represents `utils_link_stats` API object
type LinkStats struct {
	Key   string  `json:"key"` //Link key (characters after vk.cc/)
	Stats []Stats `json:"stats"`
}

// LinkStatsExtended represents `utils_link_stats_extended` API object
type LinkStatsExtended struct {
	Key   string          `json:"key"` //Link key (characters after vk.cc/)
	Stats []StatsExtended `json:"stats"`
}

// ShortLink represents `utils_short_link` API object
type ShortLink struct {
	AccessKey string `json:"access_key"`
	Key       string `json:"key"`
	ShortURL  string `json:"short_url"`
	URL       string `json:"url"`
}

// Stats represents `utils_stats` API object
type Stats struct {
	Timestamp int `json:"timestamp"`
	Views     int `json:"views"`
}

// StatsExtended represents `utils_stats_extended` API object
type StatsExtended struct {
	Cities    []StatsCity    `json:"cities"`
	Countries []StatsCountry `json:"countries"`
	SexAge    []StatsSexAge  `json:"sex_age"`
	Timestamp int            `json:"timestamp"`
	Views     int            `json:"views"`
}

// StatsSexAge represents `utils_stats_sex_age` API object
type StatsSexAge struct {
	AgeRange string `json:"age_range"`
	Female   int    `json:"female"`
	Male     int    `json:"males"`
}

//StatsCountry represents `utils_stats_country` API object
type StatsCountry struct {
	CountryID int `json:"country_id"`
	Views     int `json:"views"`
}

//StatsCity represents `utils_stats_city` API object
type StatsCity struct {
	CountryID int `json:"city_id"`
	Views     int `json:"views"`
}
