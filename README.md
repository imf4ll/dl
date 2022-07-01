<div align="center">
    <img src="./assets/logo.png" width="300" />
    <h3>A tool and core to download videos from your terminal</h3>
</div>

## ❗️ Install (CLI tool):
Arch based:
```bash
git clone https://aur.archlinux.org/dl.git
cd dl
makepkg -si
```

Using AUR Helper:
```bash
yay -S dl
```

&nbsp;

Others:
```bash
git clone https://github.com/z3oxs/dl
cd dl/cli/
make install
```

&nbsp;
## 🚀 Usage:
Downloading a video from Twitter:
```bash
dl -u "https://twitter.com/USER/status/SOMESTATUS"
```

As you can see, you need to only parse the video URL with "-u" flag.

&nbsp;

Download a video and parsing a custom filename and path:
```bash
dl -u "https://twitter.com/USER/status/SOMESTATUS" -p "some/path/video.mp4"
```

&nbsp;
##  📘 Docs (Core)
First, install the package:
```bash
go get github.com/z3oxs/dl
```
&nbsp;

Making a twitter video information request:
```go
package main

import (
    "fmt"

    "github.com/z3oxs/dl/modules/twitter"
)

func main() {
    // Will get all video information
    video, err := twitter.GetVideo("VIDEOURL")
    if err != nil {
        panic(err)

    }

    // Will print the video URL
    fmt.Println(video.Formats[0].URL)
```

### Available modules
- Twitter
- Facebook
- Rumble

&nbsp;
**If you want to contribute, make a issue requesting a new module or a pull request implementing a new module**
