package main

import "fmt"
import "os"

func main() {

    defer fmt.Println("!")

    os.Exit(3)

    // (If you run 065_exit.go using go run, the exit will be picked up by go and printed.)
    // $ go run 065_exit.go
    // exit status 3

    // (By building and executing a binary you can see the status in the terminal.)
    // $ go build 065_exit.go
    // $ ./065_exit.go
    // $ echo $?
    // 3
}
