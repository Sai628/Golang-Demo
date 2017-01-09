package main

import "fmt"

func main() {
    jobs := make(chan int, 5)
    done := make(chan bool)

    go func() {
        for {
            j, more := <-jobs
            if more {
                fmt.Println("received job:", j)
            } else {
                fmt.Println("received all jobs")
                done <- true
                return
            }
        }
    }()

    for j := 0; j < 3; j++ {
        jobs <- j
        fmt.Println("sent job:", j)
    }
    close(jobs)
    fmt.Println("sent all jobs")

    <-done

    // output like the following(they are base the go runtime):

    // sent job: 0
    // received job: 0
    // sent job: 1
    // sent job: 2
    // received job: 1
    // sent all jobs
    // received job: 2
    // received all jobs
}
