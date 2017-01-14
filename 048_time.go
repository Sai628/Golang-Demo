package main

import "fmt"
import "time"

func main() {
    
    p := fmt.Println

    now := time.Now()
    p(now)  // 2017-01-14 15:21:07.830807245 +0800 CST

    then := time.Date(2013, 11, 17, 20, 34, 58, 651387237, time.UTC)
    p(then)  // 2013-11-17 20:34:58.651387237 +0000 UTC

    p(then.Year())  // 2013
    p(then.Month())  // November
    p(then.Day())  // 17
    p(then.Hour())  // 20
    p(then.Minute())  // 34
    p(then.Second())  // 58
    p(then.Nanosecond())  // 651387237
    p(then.Location())  // UTC

    p(then.Weekday())  // Sunday

    p(then.Before(now))  // true
    p(then.After(now))  // false
    p(then.Equal(now))  // false

    diff := now.Sub(then)
    p(diff)  // 27682h46m9.179420008s

    p(diff.Hours())  // 27682.769216505556
    p(diff.Minutes())  // 1.6609661529903335e+06
    p(diff.Seconds())  // 9.965796917942001e+07
    p(diff.Nanoseconds())  // 99657969179420008

    p(then.Add(diff))  // 2017-01-14 07:21:07.830807245 +0000 UTC
    p(then.Add(-diff))  // 2010-09-21 09:48:49.471967229 +0000 UTC
}
