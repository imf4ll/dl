package utils

import "fmt"

type writer struct {
    Total,
    Transferred int64
    Percentage float64
    Filename,
    Progress string
}

func (w *writer) Write(p []byte) (int, error) {
    var toHumanDownloaded string
    var toHumanTotal string
    var progress string

    current := len(p)
    w.Transferred += int64(current)
    w.Percentage = float64(w.Transferred) / float64(w.Total) * 100

    for i := 0; i < int(w.Percentage) / 10; i++ {
        progress += "###"
    
    }

    if w.Total < 1000000000 {
        toHumanDownloaded = fmt.Sprintf("%.1f MB", float64(w.Transferred) / 1000000)
        toHumanTotal = fmt.Sprintf("%.1f MB", float64(w.Total) / 1000000)

    } else {
        toHumanDownloaded = fmt.Sprintf("%.1f GB", float64(w.Transferred) / 1000000000)
        toHumanTotal = fmt.Sprintf("%.1f GB", float64(w.Total) / 1000000000)

    }

    fmt.Printf("\r%s%s, %s | %.1f%% (%s) <%s>%s", "\033[1;32m", w.Filename, toHumanTotal, w.Percentage, toHumanDownloaded, progress, "\033[m")

    return current, nil
}
