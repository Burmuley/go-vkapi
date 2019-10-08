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
// `stats` group of objects
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

// StatsActivity type represents `stats_activity` API object
type StatsActivity struct {
	Comments     int `json:"comments"`     // Comments number
	Copies       int `json:"copies"`       // Reposts number
	Hidden       int `json:"hidden"`       // Hidden from news count
	Likes        int `json:"likes"`        // Likes number
	Subscribed   int `json:"subscribed"`   // New subscribers count
	Unsubscribed int `json:"unsubscribed"` // Unsubscribed count
}

// StatsWallpostStat type represents `stats_wallpost_stat` API object
type StatsWallpostStat struct {
	Hide             int `json:"hide"`              // Hidings number
	JoinGroup        int `json:"join_group"`        // People have joined the group
	Links            int `json:"links"`             // Link clickthrough
	ReachSubscribers int `json:"reach_subscribers"` // Subscribers reach
	ReachTotal       int `json:"reach_total"`       // Total reach
	Report           int `json:"report"`            // Reports number
	ToGroup          int `json:"to_group"`          // Clickthrough to community
	Unsubscribe      int `json:"unsubscribe"`       // Unsubscribed members
}

// StatsCountry type represents `stats_country` API object
type StatsCountry struct {
	Code  string `json:"code"`  // Country code
	Count int    `json:"count"` // Visitors number
	Name  string `json:"name"`  // Country name
	Value int    `json:"value"` // Country ID
}

// StatsPeriod type represents `stats_period` API object
type StatsPeriod struct {
	Activity   StatsActivity `json:"activity"`
	PeriodFrom int           `json:"period_from"` // Unix timestamp
	PeriodTo   int           `json:"period_to"`   // Unix timestamp
	Reach      StatsReach    `json:"reach"`
	Visitors   StatsViews    `json:"visitors"`
}

// StatsCity type represents `stats_city` API object
type StatsCity struct {
	Count int    `json:"count"` // Visitors number
	Name  string `json:"name"`  // City name
	Value int    `json:"value"` // City ID
}

// StatsSexAge type represents `stats_sex_age` API object
type StatsSexAge struct {
	Count int    `json:"count"` // Visitors number
	Value string `json:"value"` // Sex/age value
}

// StatsViews type represents `stats_views` API object
type StatsViews struct {
	Age         []StatsSexAge  `json:"age"`
	Cities      []StatsCity    `json:"cities"`
	Countries   []StatsCountry `json:"countries"`
	MobileViews int            `json:"mobile_views"` // Number of views from mobile devices
	Sex         []StatsSexAge  `json:"sex"`
	SexAge      []StatsSexAge  `json:"sex_age"`
	Views       int            `json:"views"`    // Views number
	Visitors    int            `json:"visitors"` // Visitors number
}

// StatsReach type represents `stats_reach` API object
type StatsReach struct {
	Age              []StatsSexAge  `json:"age"`
	Cities           []StatsCity    `json:"cities"`
	Countries        []StatsCountry `json:"countries"`
	MobileReach      int            `json:"mobile_reach"`      // Reach count from mobile devices
	Reach            int            `json:"reach"`             // Reach count
	ReachSubscribers int            `json:"reach_subscribers"` // Subscribers reach count
	Sex              []StatsSexAge  `json:"sex"`
	SexAge           []StatsSexAge  `json:"sex_age"`
}
