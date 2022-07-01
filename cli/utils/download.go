package utils

import (
    "os"
    "errors"
    "net/http"
    "io"
    "fmt"
)

func Download(url, filename string) error {
    fmt.Println()

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

    src := io.TeeReader(r.Body, &writer {
        Total: r.ContentLength,
        Filename: filename,
    })

    _, err = io.Copy(file, src)
    if err != nil {
        return errors.New("Failed to download video.")
    }

    fmt.Println()

    return nil
}
