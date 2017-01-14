package main

import "fmt"
import "time"

func main() {

    now := time.Now()    
    secs := now.Unix()
    nanos := now.UnixNano()
    fmt.Println(now)  // 2017-01-14 15:46:09.67084145 +0800 CST

    millis := nanos / 1000000
    fmt.Println(secs)  // 1484379969
    fmt.Println(millis)  // 1484379969670
    fmt.Println(nanos)  // 1484379969670841450

    fmt.Println(time.Unix(secs, 0))  // 2017-01-14 15:46:09 +0800 CST
    fmt.Println(time.Unix(0, nanos))  // 2017-01-14 15:46:09.67084145 +0800 CST
}
