package dlcore

type Writer struct {
    Total int64
    Transferred int64
    Title string
}

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

type InstagramVideo struct {
    Info struct {
        ID int `json:"pk"`
        Caption string `json:"text"`
        UserID int `json:"user_id"`
    } `json:"info"`
    Author struct {
        ID int `json:"pk"`
        Username string `json:"username"`
        Name string `json:"full_name"`
        IsPrivate bool `json:"is_private"`
        IsVerified bool `json:"is_verified"`
        ProfilePhoto string `json:"profile_pic_url"`
    } `json:"author"`
    Formats []InstagramVideoFormat `json:"formats"`
}

type InstagramVideoFormat struct {
    ID string `json:"id"`
    Type string `json:"type"`
    Width int `json:"width"`
    Height int `json:"height"`
    URL string `json:"url"`
}

type YouTubeVideos struct {
    Title,
    Duration string
    Formats []YouTubeVideo
}

type YouTubeVideo struct {
    Quality,
    MimeType,
    URL,
    Title,
    Audio string
    DownloadAudio bool
}

type PlaylistInfo struct {
    Author,
    Title string
    Size int
    Videos []YouTubePlaylistVideos
}

type YouTubePlaylistVideos struct {
    Title string
    Format YouTubeVideo
}

type FacebookVideo struct {
    Feedback FacebookVideoFeedback `json:"feedback"`
    Title string `json:"title"`
    Author struct {
        ID string `json:"id"`
        Name string `json:"name"`
        ProfilePhoto string `json:"profile_pic_uri"`
    } `json:"author"`
    Formats []FacebookVideoFormat `json:"formats"`
    Qualities FacebookVideoQualities `json:"qualities"`
}

type FacebookVideoFeedback struct {
    Comment struct {
        Count int `json:"total_count"`
        IsEmpty bool `json:"is_empty"`
    } `json:"comment_count"`
    Reaction struct {
        Count int `json:"count"`
        IsEmpty bool `json:"is_empty"`
    } `json:"reaction_count"`
    I18NCommentCount string `json:"i18n_comment_count"`
    I18NReactionCount string `json:"i18n_reaction_count"`
}

type FacebookVideoFormat struct {
    ID string `json:"representation_id"`
    Codec string `json:"codecs"`
    MimeType string `json:"mime_type"`
    URL string `json:"base_url"`
    Bitrate int `json:"bandwidth"`
    Height int `json:"height"`
    Width int `json:"width"`
}

type FacebookVideoQualities struct {
    SD string `json:"playable_url"`
    HD string `json:"playable_url_quality_hd"`
}

type RumbleVideo struct {
    FPS int `json:"fps"`
    Width int `json:"w"`
    Height int `json:"h"`
    Format struct {
        Standard RumbleVideoFormat `json:"mp4"`
    } `json:"u"`
    Formats []RumbleVideoFormat `json:"formats"`
    Thumbnail string `json:"i"`
    Link string `json:"l"`
    Title string `json:"title"`
    Author struct {
        Name string `json:"name"`
        URL string `json:"url"`
    } `json:"author"`
    Duration int `json:"duration"`
    CreatedAt string `json:"pubDate"`
}

type RumbleVideoFormat struct {
    URL string `json:"url"`
    Meta struct {
        Bitrate int `json:"bitrate"`
        Size int `json:"size"`
        Width int `json:"w"`
        Height int `json:"h"`
    } `json:"meta"`
    Title,
    MimeType string
}
