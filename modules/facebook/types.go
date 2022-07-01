package facebook

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

