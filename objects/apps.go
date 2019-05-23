package objects

/////////////////////////////////////////////////////////////
// Applications related API objects	                       //
/////////////////////////////////////////////////////////////

// App represents `apps_app` API object
type AppsApp struct {
	AuthorGroup     int                    `json:"author_group"`
	AuthorID        int                    `json:"author_id"`
	AuthorUrl       string                 `json:"author_url"`
	Banner1120      string                 `json:"banner_1120"`
	Banner560       string                 `json:"banner_560"`
	CatalogPosition int                    `json:"catalog_position"`
	Description     string                 `json:"description"`
	Genre           string                 `json:"genre"`
	GenreID         int                    `json:"genre_id"`
	Icon139         string                 `json:"icon_139"`
	Icon150         string                 `json:"icon_150"`
	Icon278         string                 `json:"icon_278"`
	Icon75          string                 `json:"icon_75"`
	ID              int                    `json:"id"`
	International   int                    `json:"international"`
	IsInCatalog     int                    `json:"is_in_catalog"`
	LeaderboardType AppsAppLeaderboardType `json:"leaderboard_type"`
	MembersCount    int                    `json:"members_count"`
	PlatformID      int                    `json:"platform_id"`
	PublishedDate   int                    `json:"published_date"`
	ScreenName      string                 `json:"screen_name"`
	Screenshots     []PhotosPhoto          `json:"screenshots"`
	Section         string                 `json:"section"`
	Title           string                 `json:"title"`
	Type            AppsAppType            `json:"type"`
}

// Leaderboard represents `apps_leaderboard` API object
type AppsLeaderboard struct {
	Level  int `json:"level"`
	Points int `json:"points"`
	Score  int `json:"score"`
	UserID int `json:"user_id"`
}

// AppLeaderboardType represents `apps_app_leaderboard_type` API object
type AppsAppLeaderboardType int

func (a *AppsAppLeaderboardType) MarshalJSON() ([]byte, error) {
	return GetIntFromRange(int(*a), 0, 2, true)
}

func (a *AppsAppLeaderboardType) String() string {
	switch *a {
	case 0:
		return "not supported"
	case 1:
		return "levels"
	case 2:
		return "points"
	default:
		return "unknown"
	}
}

// AppType represents `apps_app_type` API object
type AppsAppType string

func (a *AppsAppType) MarshalJSON() ([]byte, error) {
	return GetStringFromRange(string(*a), "app", "game", "site", "standalone", "vk_app",
		"community_app", "html5_game")
}

func (a *AppsAppType) String() string {
	return string(*a)
}

// Scope represents `apps_scope` API object
type AppsScope struct {
	Name  string `json:"name"`
	Title string `json:"title"`
}
