package main

import "fmt"
import "time"

func main() {

    requests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        requests <- i
    }
    close(requests)

    limiter := time.Tick(time.Millisecond * 200)
    for req := range requests {
        <-limiter
        fmt.Println("request", req, time.Now())
    }


    burstyLimiter := make(chan time.Time, 3)
    for i := 0; i < 3; i++ {
        burstyLimiter <- time.Now()
    }

    go func() {
        for t := range time.Tick(time.Millisecond * 200) {
            burstyLimiter <- t
        }
    }()

    burstyRequests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        burstyRequests <- i
    }
    close(burstyRequests)

    for req := range burstyRequests {
        <-burstyLimiter
        fmt.Println("request", req, time.Now())
    }

    // (the first batch of requests handled once every ~200ms)
    // request 1 2017-01-10 14:47:30.627100604 +0800 CST
    // request 2 2017-01-10 14:47:30.827544069 +0800 CST
    // request 3 2017-01-10 14:47:31.024911414 +0800 CST
    // request 4 2017-01-10 14:47:31.223696144 +0800 CST
    // request 5 2017-01-10 14:47:31.423485974 +0800 CST


    // (the second batch of requests, serve the first 3 immediately because of the burstable rate limiting, then serve the remaining 2 with ~200ms delays each.):
    // request 1 2017-01-10 14:47:31.423596988 +0800 CST
    // request 2 2017-01-10 14:47:31.423613199 +0800 CST
    // request 3 2017-01-10 14:47:31.423623054 +0800 CST
    // request 4 2017-01-10 14:47:31.623977218 +0800 CST
    // request 5 2017-01-10 14:47:31.823706382 +0800 CST
}
