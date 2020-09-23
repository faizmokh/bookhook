package hooks

type Tweet struct {
	ForUserID         string             `json:"for_user_id"`
	TweetCreateEvents []TweetCreateEvent `json:"tweet_create_events"`
}

type TweetCreateEvent struct {
	CreatedAt             string                `json:"created_at"`
	ID                    float64               `json:"id"`
	IDStr                 string                `json:"id_str"`
	Text                  string                `json:"text"`
	Truncated             bool                  `json:"truncated"`
	QuotedStatusID        float64               `json:"quoted_status_id"`
	QuotedStatusIDStr     string                `json:"quoted_status_id_str"`
	QuotedStatus          QuotedStatus          `json:"quoted_status"`
	QuotedStatusPermalink QuotedStatusPermalink `json:"quoted_status_permalink"`
	Entities              Entities              `json:"entities"`
	Favorited             bool                  `json:"favorited"`
	Retweeted             bool                  `json:"retweeted"`
	TimestampMS           string                `json:"timestamp_ms"`
}

type Entities struct {
	Hashtags     []Hashtag     `json:"hashtags"`
	Urls         []URL         `json:"urls"`
	UserMentions []interface{} `json:"user_mentions"`
	Symbols      []interface{} `json:"symbols"`
}

type Hashtag struct {
	Text    string  `json:"text"`
	Indices []int64 `json:"indices"`
}

type URL struct {
	URL         string  `json:"url"`
	ExpandedURL string  `json:"expanded_url"`
	DisplayURL  string  `json:"display_url"`
	Indices     []int64 `json:"indices"`
}

type QuotedStatus struct {
	CreatedAt     string   `json:"created_at"`
	ID            float64  `json:"id"`
	IDStr         string   `json:"id_str"`
	Text          string   `json:"text"`
	Truncated     bool     `json:"truncated"`
	IsQuoteStatus bool     `json:"is_quote_status"`
	QuoteCount    int64    `json:"quote_count"`
	ReplyCount    int64    `json:"reply_count"`
	RetweetCount  int64    `json:"retweet_count"`
	FavoriteCount int64    `json:"favorite_count"`
	Entities      Entities `json:"entities"`
}

type QuotedStatusPermalink struct {
	URL      string `json:"url"`
	Expanded string `json:"expanded"`
	Display  string `json:"display"`
}

type SendMessage struct {
	ChatID                string `json:"chat_id"`
	Text                  string `json:"text"`
	DisableWebPagePreview bool   `json:"disable_web_page_preview"`
	DisableNotification   bool   `json:"disable_notification"`
	ParseMode             string `json:"parse_mode"`
}
