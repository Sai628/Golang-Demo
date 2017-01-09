package main

import "fmt"

func f(from string)  {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}

func main()  {

    f("direct")
    // direct : 0
    // direct : 1
    // direct : 2

    go f("goroutine")

    go func(msg string) {
        fmt.Println(msg)
    }("going")

    // goroutine : 0
    // goroutine : 1
    // going
    // goroutine : 2

    var input string
    fmt.Scanln(&input)  // <press enter>
    fmt.Println("Done")  // Done
}
