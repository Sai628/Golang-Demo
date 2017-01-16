package main

import "fmt"
import "os"

func main() {

    argsWithProg := os.Args
    argsWithoutProg := os.Args[1:]
    arg := os.Args[3]

    fmt.Println(argsWithProg)
    fmt.Println(argsWithoutProg)
    fmt.Println(arg)

    // To experiment with command-line arguments it's best to build a binary with "go build" first:

    // $ go build 059_command-line_arguments.go
    // $ ./059_command-line_arguments a b c d
    
    // [./059_command-line_arguments a b c d]
    // [a b c d]
    // c
}
