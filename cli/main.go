package main

import (
    "fmt"
    "flag"
    "os"
    "strings"

    "github.com/z3oxs/dl/cli/modules"
)

var (
    url,
    path string
    version bool
)

func init() {
    flag.StringVar(&url, "u", "", "URL from video")
    flag.StringVar(&path, "p", "video.mp4", "Path for downloaded video")
    flag.BoolVar(&version, "V", false, "Prints version")
    
    flag.Usage = func() {
        fmt.Printf(`Usage:
    dl [flags]

Description:
    A tool to download videos from your terminal

Flags:
    -u    URL from video
    -p    Path for downloaded video
    -V    Prints version
`)
    }

    flag.Parse()
}

func main() {
    if version {
        fmt.Println("v1.0.0")

        os.Exit(3)
    }

    if (url == "" && path == "") {
        panic("You need to specify at least URL to video")

    }

    if strings.Contains(url, "twitter") {
        modules.Twitter(url, path)

    } else if strings.Contains(url, "facebook") {
        modules.Facebook(url, path)

    } else if strings.Contains(url, "rumble") {
        modules.Rumble(url, path)

    } else {
        flag.Usage()

    }
}
