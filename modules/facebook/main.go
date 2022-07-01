package facebook

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "strings"
    "errors"
)

func GetFormats(url string) ([]FacebookVideoFormat, error) {
    video, err := GetVideo(url)
    if err != nil {
        return []FacebookVideoFormat{}, err

    }

    return video.Formats, nil
}

func GetVideo(url string) (FacebookVideo, error) {
    var video FacebookVideo

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return FacebookVideo{}, errors.New("Failed to get video information")

    }
    
    req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")

    client := &http.Client{}
    r, err := client.Do(req)
    if err != nil || r.StatusCode != 200 {
        return FacebookVideo{}, errors.New("Failed to get video information")

    }

    b, _ := ioutil.ReadAll(r.Body)
    raw := string(b)

    if !strings.Contains(raw, `"representations":`) {
        return FacebookVideo{}, errors.New("Invalid video information")

    }
   
    formats := strings.Split(strings.Split(raw, `"representations":`)[1], `,"video_id"`)[0]
    feedback := strings.Split(strings.Split(raw, `"feedback":`)[1], `,"reaction_display_config"`)[0]
    qualities := strings.Split(strings.Split(raw, `init":null,`)[1], `,"spherical_`)[0]
    author := strings.Split(strings.Split(raw, `"owner_as_page":`)[1], `,"__isActor"`)[0]
    title := strings.Split(strings.Split(raw, `"color_ranges":[],"text":`)[1], `}`)[0]
    json.Unmarshal(json.RawMessage(fmt.Sprintf(`{"formats":%v,"feedback":%v},"qualities":{%v},"author":%v,"title":%v}`, formats, feedback, qualities, author, title)), &video)

    return video, nil
}
