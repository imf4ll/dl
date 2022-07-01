package modules

import (
    "fmt"

    "github.com/z3oxs/dl/modules/twitter"
    "github.com/z3oxs/dl/cli/utils"
)

func Twitter(url, path string) {
    var choose int

    video, err := twitter.GetVideo(url)
    if err != nil {
        panic(err)

    } else if len(video.Formats) == 0 {
        panic("Video returned 0 formats")

    }

    fmt.Printf(`
%sAuthor:%s %s (%s)
%sCaption:%s %s

`, "\033[1;32m", "\033[m", video.User.Info.Name, video.User.Info.Nickname, "\033[1;32m", "\033[m", video.Information.Caption)

    for k, format := range video.Formats {
        fmt.Printf("%s[%d] %s%s%s\n", "\033[1;34m", k, "\033[0;34m", format.Quality, "\033[m")

    }

    fmt.Print("> ")
    fmt.Scanf("%d", &choose)

    utils.Download(video.Formats[choose].URL, path)
}
