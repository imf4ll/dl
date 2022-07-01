package twitter

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "strings"
    "errors"
)

func getGuestToken() (string, error) {
    var guestToken TwitterGuestToken

    guestTokenReq, err := http.NewRequest("POST", "https://api.twitter.com/1.1/guest/activate.json", nil)
    if err != nil {
        return "", errors.New("Failed to request guest token")

    }

    guestTokenReq.Header.Add("Authorization", "Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA")
    
    guestTokenClient := &http.Client{}
    g, err := guestTokenClient.Do(guestTokenReq)
    if err != nil || g.StatusCode != 200 {
        return "", errors.New("Failed to request guest token")

    }

    token, err := ioutil.ReadAll(g.Body)
    if err != nil {
        return "", err

    }

    json.Unmarshal(token, &guestToken)

    return guestToken.Token, nil
}

func GetFormats(url string) ([]TwitterFormat, error) {
    video, err := GetVideo(url)
    if err != nil {
        return []TwitterFormat{}, err

    }

    return video.Formats, nil
}

func GetVideo(url string) (TwitterVideo, error) {
    var video TwitterVideo
    var videoFormats []TwitterFormat

    token, err := getGuestToken()
    if err != nil {
        return TwitterVideo{}, err

    }
        
    videoReq, err := http.NewRequest("GET", `https://twitter.com/i/api/graphql/LJ_TjoWGgNTXCl7gfx4Njw/TweetDetail?variables=%7B%22focalTweetId%22%3A%22` + strings.Split(url, "/")[5] + `%22%2C%22with_rux_injections%22%3Afalse%2C%22includePromotedContent%22%3Atrue%2C%22withCommunity%22%3Atrue%2C%22withQuickPromoteEligibilityTweetFields%22%3Atrue%2C%22withBirdwatchNotes%22%3Afalse%2C%22withSuperFollowsUserFields%22%3Atrue%2C%22withDownvotePerspective%22%3Afalse%2C%22withReactionsMetadata%22%3Afalse%2C%22withReactionsPerspective%22%3Afalse%2C%22withSuperFollowsTweetFields%22%3Atrue%2C%22withVoice%22%3Atrue%2C%22withV2Timeline%22%3Afalse%2C%22__fs_dont_mention_me_view_api_enabled%22%3Afalse%2C%22__fs_interactive_text_enabled%22%3Atrue%2C%22__fs_responsive_web_uc_gql_enabled%22%3Afalse%7D`, nil)
    if err != nil {
        return TwitterVideo{}, errors.New("Failed to request video information")

    }

    videoReq.Header = http.Header {
        "Host": []string{"twitter.com"},
        "Authorization": []string{"Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA"},
        "X-Guest-Token": []string{token},
    }
    
    videoClient := &http.Client{}
    v, err := videoClient.Do(videoReq)
    if err != nil || v.StatusCode != 200 {
        return TwitterVideo{}, errors.New("Failed to request video information")

    }

    b, _ := ioutil.ReadAll(v.Body)
    videoResponse := string(b)

    if !strings.Contains(string(videoResponse), `"variants":[`) {
        return TwitterVideo{}, errors.New("Invalid video information")

    }

    info := strings.Split(strings.Split(videoResponse, `"legacy":`)[2], `,"quick_promote`)[0]
    formats := strings.Split(strings.Split(videoResponse, `"variants":[`)[1], `]`)[0]
    user := strings.Split(strings.Split(videoResponse, `"user_results":{"result":`)[1], `}},"legacy"`)[0]
    json.Unmarshal(json.RawMessage(fmt.Sprintf(`{"information":%v,"formats":[%v],"user":%v}`, info, formats, user)), &video)

    for _, v := range video.Formats {
        if v.ContentType == "video/mp4" {
            videoFormats = append(videoFormats, TwitterFormat {
                Bitrate: v.Bitrate,
                ContentType: v.ContentType,
                URL: v.URL,
                Quality: strings.Split(strings.Split(v.URL, "vid/")[1], "/")[0],
            })
        }
    }

    video.Formats = videoFormats

    return video, nil
}
