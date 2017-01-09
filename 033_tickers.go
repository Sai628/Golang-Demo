package main

import "fmt"
import "time"

func main() {

    ticker := time.NewTicker(time.Millisecond * 500)
    go func() {
        for t := range ticker.C {
            fmt.Println("Tick at", t)
        }
    }()

    time.Sleep(time.Millisecond * 1600)
    // Tick at 2017-01-09 16:05:21.412931505 +0800 CST
    // Tick at 2017-01-09 16:05:21.913499876 +0800 CST
    // Tick at 2017-01-09 16:05:22.413425163 +0800 CST

    ticker.Stop()
    fmt.Println("Ticker stopped")  // Ticker stopped
}
