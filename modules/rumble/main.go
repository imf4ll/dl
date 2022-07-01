package dlcore

import (
    "net/http"
    "encoding/json"
    "io/ioutil"
    "strings"
    "fmt"
)

func Rumble(url string) (RumbleVideo, error) {
    var video RumbleVideo

    v, err := http.Get(url)
    if err != nil {
        return RumbleVideo{}, fail

    }

    b, _ := ioutil.ReadAll(v.Body)
    videoID := strings.Split(strings.Split(string(b), `"video":"`)[1], `"`)[0]

    e, err := http.Get(`https://rumble.com/embedJS/u3/?request=video&ver=2&v=` + videoID +`&ext=%7B%22ad_count%22%3Anull%7D&ad_wt=9043`)
    if err != nil {
        return RumbleVideo{}, fail

    }

    b, _ = ioutil.ReadAll(e.Body)
    json.Unmarshal(b, &video)

    raw := strings.ReplaceAll(fmt.Sprintf(`{"formats":[%v]}`, strings.Split(strings.Split(string(b), `"ua":{"mp4":{`)[1], `}},"i"`)[0]), `},"webm":{`, `,`)
    formatsRaw := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(raw, `"240":`, ``), `"360":`, ``), `"480":`, ``), `"720":`, ``), `"1080":`, ``)
    json.Unmarshal(json.RawMessage(formatsRaw), &video)

    for i := range video.Formats {
        v := &video.Formats[i]

        if strings.Contains(v.URL, "mp4") {
            v.MimeType = "mp4"

        } else {
            v.MimeType = "webm"

        }
        
        v.Title = fmt.Sprintf("%v.%v", video.Title, v.MimeType)
    }

    return video, nil
}

func (video *RumbleVideoFormat) Download() error {
    err := Download(video.URL, video.Title)
    if err != nil {
        return err

    }

    return nil
}
