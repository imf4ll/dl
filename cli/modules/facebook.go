package modules

import (
    "fmt"

    "github.com/z3oxs/dl/modules/facebook"
    "github.com/z3oxs/dl/cli/utils"
)

func Facebook(url, path string) {
    var choose int

    video, err := facebook.GetVideo(url)
    if err != nil {
        panic(err)

    } else if len(video.Formats) == 0 {
        panic("Video returned 0 formats")

    }

    fmt.Printf(`
%sAuthor:%s %s
%sTitle:%s %s

`, "\033[1;32m", "\033[m", video.Author.Name, "\033[1;32m", video.Title, "\033[m")

    for k, format := range video.Formats {
        fmt.Printf("%s[%d] %s%dx%d%s\n", "\033[1;34m", k, "\033[0;34m", format.Width, format.Height, "\033[m")

    }

    fmt.Print("> ")
    fmt.Scanf("%d", &choose)

    utils.Download(video.Formats[choose].URL, path)
}
