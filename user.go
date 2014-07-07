package trailer

type User struct {
	Id          uint64 `json:"id"`
	ScreenName  string `json:"screen_name"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Url         string `json:"url"`

	StatusesCount  uint64 `json:"statuses_count"`
	FriendsCount   uint64 `json:"friends_count"`
	FollowersCount uint64 `json:"followers_count"`
	FavoritesCount uint64 `json:"favourites_count"`
	ListedCount    uint64 `json:"listed_count"`

	Verified            bool `json:"verified"`
	IsTranslator        bool `json:"is_translator"`
	IsTranslatorEnabled bool `json:"is_translation_enabled"`

	ProfileImageUrl      string `json:"profile_image_url"`
	ProfileImageUrlHttps string `json:"profile_image_url_https"`
	DefaultProfileImage  bool   `json:"default_profile_image"`

	ProfileBackgroundImageUrl      string `json:"profile_background_image_url"`
	ProfileBackgroundImageUrlHttps string `json:"profile_background_image_url_https"`
	ProfileBackgroundTile          bool   `json:"profile_background_tile"`
	ProfileBackgroundColor         string `json:"profile_background_color"`
	ProfileUseBackgroundImage      bool   `json:"profile_use_background_image"`

	ProfileLinkColor string `json:"profile_link_color"`
	ProfileTextColor string `json:"profile_text_color"`

	ProfileSidebarBorderColor string `json:"profile_sidebar_border_color"`
	ProfileSidebarFillColor   string `json:"profile_sidebar_fill_color"`

	DefaultProfile bool `json:"default_profile"`
}

/*
var (
	User = map[string]interface{}{
		"created_at":           "Sat Jun 06 11:34:55 +0000 2009",
		"contributors_enabled": bool(false),
		"following":            nil,
		"lang":                 "ja",
		"utc_offset":           float64(32400),
		"time_zone":            "Tokyo",
		"geo_enabled":          bool(false),
		"follow_request_sent":  nil,
		"notifications":        nil,
		"location":             "",
		"protected":            bool(false),
	}
)
*/
