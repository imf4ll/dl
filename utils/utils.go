package dlcore

import (
    "sync"
    "os"
    "os/exec"
    "errors"
    "fmt"
    "net/http"
    "io"
    "strings"

    "github.com/z3oxs/progressbar.go"
)

var (
    fail = errors.New("Failed to fetch video information.")
    invalid = errors.New("No one valid video was found.")
    wg sync.WaitGroup
    box = "█"
)

func Download(url, filename string) error {
    file, err := os.Create(filename)
    if err != nil {
        return errors.New("Failed to download video.")

    }

    defer file.Close()

    r, err := http.Get(url)
    if err != nil {
        return errors.New("Failed to download video.")

    }

    defer r.Body.Close()

    src := io.TeeReader(r.Body, &progressbar.Buffer {
        Total: r.ContentLength,
        Title: filename,
        ShowTitle: true,
        ShowPercent: true,
        ShowLength: true,
        ShowTime: true,
        Blocks: 30,
        Style: progressbar.Style {
            Start: "|",
            End: "|",
            Color: "red",
            Font: "bold",
        },
    })

    _, err = io.Copy(file, src)
    if err != nil {
        return errors.New("Failed to download video.")
    }

    fmt.Println()

    return nil
}

func Convert(filename string) error {
    name := ""
    
    if strings.Contains("webm", filename) {
        name = filename[:len(filename) - 4]
    
    } else {
        name = filename[:len(filename) - 3]

    }

    cmd := exec.Command("ffmpeg", "-i", filename, fmt.Sprintf("%v.mp3", name))
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Run(); err != nil {
        return errors.New("Failed to download video.")

    }

    os.Remove(filename)

    return nil
}
