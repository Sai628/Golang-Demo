package main

import "flag"
import "fmt"

func main() {

    wordPtr := flag.String("word", "foo", "a string")
    numbPtr := flag.Int("numb", 42, "an int")
    boolPtr := flag.Bool("fork", false, "a bool")

    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var")

    // Must be called after all flags are defined and before flags are accessed by the program.
    flag.Parse()

    fmt.Println("word:", *wordPtr)
    fmt.Println("numb:", *numbPtr)
    fmt.Println("boolPtr:", *boolPtr)
    fmt.Println("svar:", svar)
    fmt.Println("tail:", flag.Args())

    // (To experiment with the command-line flags program it's best to first compile it and then run the resulting binary directly.)
    // $ go build 060_command-line-flags.go


    // (Try out the built program by first giving it values for all flags.)
    // word: opt
    // numb: 7
    // boolPtr: true
    // svar: flag
    // tail: []

    // (If you omit flags they automatically take their default values.)
    // word: opt
    // numb: 42
    // boolPtr: false
    // svar: bar
    // tail: []

    // (Trailing positional arguments can be provided after any flags.)
    // word: opt
    // numb: 42
    // boolPtr: false
    // svar: bar
    // tail: [a1 a2 a3]
}
