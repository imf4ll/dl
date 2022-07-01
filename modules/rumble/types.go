package rumble

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
