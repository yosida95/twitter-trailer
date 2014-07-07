package trailer

type Tweet struct {
	Id uint64 `json:"id"`

	Favorited       bool   `json:"favorited"`
	Retweeted       bool   `json:"retweeted"`
	RetweetedStatus *Tweet `json:"retweeted_status"`

	RetweetCount   uint `json:"retweet_count"`
	FavoritesCount uint `json:"favourites_count"`

	Text string `json:"text"`

	User     User `json:"user"`
	Entities struct {
		Urls         []Url         `json:"urls"`
		Hashtags     []Hashtag     `json:"hashtags"`
		UserMentions []UserMention `json:"user_mentions"`
	} `json:"entities"`

	InReplyToUserId     uint64 `json:"in_reply_to_user_id"`
	InReplyToScreenName string `json:"in_reply_to_screen_name"`
	InReplyToStatusId   uint64 `json:"in_reply_to_status_id"`
}

type Indices [2]uint

type Entity struct {
	Indices Indices `json:"indices"`
}

type Url struct {
	ExpandedUrl string `json:"expanded_url"`
	Url         string `json:"url"`
	DisplayUrl  string `json:"display_url"`
	Entity
}

type Hashtag struct {
	Text string `json:"text"`
	Entity
}

type UserMention struct {
	Id         uint64 `json:"id"`
	ScreenName string `json:"screen_name"`
	Name       string `json:"name"`
	Entity
}

/*
{
  "coordinates": null,
  "truncated": false,
  "created_at": "Wed Jun 06 20:07:10 +0000 2012",
  "entities": {
    "user_mentions": [

    ]
  },
  "contributors": [
    14927800
  ],
  "geo": null,
  "possibly_sensitive": false,
  "place": null,
  "user": {
    "profile_sidebar_fill_color": "DDEEF6",
    "profile_sidebar_border_color": "C0DEED",
    "profile_background_tile": false,
    "profile_image_url": "http://a0.twimg.com/profile_images/2284174872/7df3h38zabcvjylnyfe3_normal.png",
    "created_at": "Wed May 23 06:01:13 +0000 2007",
    "location": "San Francisco, CA",
    "follow_request_sent": false,
    "profile_link_color": "0084B4",
    "is_translator": false,
    "default_profile": true,
    "contributors_enabled": true,
    "url": "http://dev.twitter.com",
    "profile_image_url_https": "https://si0.twimg.com/profile_images/2284174872/7df3h38zabcvjylnyfe3_normal.png",
    "utc_offset": -28800,
    "profile_use_background_image": true,
    "listed_count": 10774,
    "profile_text_color": "333333",
    "lang": "en",
    "followers_count": 1212963,
    "notifications": null,
    "profile_background_image_url_https": "https://si0.twimg.com/images/themes/theme1/bg.png",
    "profile_background_color": "C0DEED",
    "verified": true,
    "geo_enabled": true,
    "time_zone": "Pacific Time (US & Canada)",
    "default_profile_image": false,
    "profile_background_image_url": "http://a0.twimg.com/images/themes/theme1/bg.png",
    "statuses_count": 3333,
    "friends_count": 31,
    "following": true,
    "show_all_inline_media": false,
  },
  "source": "web",
}
*/
