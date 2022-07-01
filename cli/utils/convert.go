package utils

import (
    "strings"
    "os/exec"
    "os"
    "errors"
    "fmt"
)

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
