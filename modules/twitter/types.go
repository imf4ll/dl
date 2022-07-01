package twitter

type TwitterGuestToken struct {
    Token string `json:"guest_token"`
}

type TwitterVideo struct {
    Information TwitterInfo `json:"information"`
    Formats []TwitterFormat `json:"formats"`
    User TwitterUser `json:"user"`
}

type TwitterInfo struct {
    Favorite int `json:"favourite_count"`
    Caption string `json:"full_text"`
    Language string `json:"lang"`
    Quote int `json:"quote_count"`
    Reply int `json:"reply_count"`
    Retweeted int `json:"retweet_count"`
}

type TwitterFormat struct {
    Bitrate int `json:"bitrate"`  
    ContentType string `json:"content_type"`  
    URL string `json:"url"`  
    Quality string
}

type TwitterUser struct {
    Type string `json:"__typename"`
    ID string `json:"id"`
    Info struct {
        CreatedAt string `json:"created_at"`
        Description string `json:"description"`
        Entities struct {
            URL struct {
                URLs []struct {
                    DisplayURL string `json:"display_url"`
                    ExpandedURL string `json:"expanded_url"`
                    URL string `json:"url"`
                } `json:"urls"`
            } `json:"url"`
        } `json:"entities"`
        Favourites int `json:"favourites_count"`
        Followers int `json:"followers_count"`
        Friends int `json:"friends_count"`
        Listed int `json:"listed_count"`
        Location string `json:"location"`
        Media int `json:"media_count"`
        Name string `json:"name"`
        NormalFollowers int `json:"normal_followers_count"`
        PinnedTweets []string `json:"pinned_tweet_ids_str"`
        BannerURL string `json:"profile_banner_url"`
        ProfilePhoto string `json:"profile_image_url_https"`
        Nickname string `json:"screen_name"`
        Status string `json:"statuses_count"`
        Protected bool `json:"protected"`
    } `json:"legacy"`
}

