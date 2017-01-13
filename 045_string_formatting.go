package main

import "fmt"
import "os"

type point struct {
    x, y int
}

func main() {
    
    p := point{1, 2}

    fmt.Printf("%v\n", p)  // {1 2}
    fmt.Printf("%+v\n", p)  // {x:1 y:2}
    fmt.Printf("%#v\n", p)  // main.point{x:1, y:2}
    fmt.Printf("%T\n", p)  // main.point

    fmt.Printf("%t\n", true)  // true

    fmt.Printf("%d\n", 123)  // 123
    fmt.Printf("%b\n", 14)  // 1110 (print a binary representation)
    fmt.Printf("%c\n", 33)  // !
    fmt.Printf("%x\n", 456)  // 1c8 (print hex encoding)

    fmt.Printf("%f\n", 78.9)  // 78.900000
    fmt.Printf("%e\n", 123400000.0)  // 1.234000e+08
    fmt.Printf("%E\n", 123400000.0)  // 1.234000E+08

    fmt.Printf("%s\n", "\"string\"")  // "string"
    fmt.Printf("%q\n", "\"string\"")  // "\"string\""

    fmt.Printf("%x\n", "hex this")  // 6865782074686973 (renders the string in base-16, with two output characters per byte of input)
    fmt.Printf("%p\n", &p)  // 0xc420072030

    fmt.Printf("|%6d|%6d|\n", 12, 345)         // |    12|   345|
    fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)   // |  1.20|  3.45|
    fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45) // |1.20  |3.45  |
    
    fmt.Printf("|%6s|%6s|\n", "foo", "b")      // |   foo|     b|
    fmt.Printf("|%-6s|%-6s|\n", "foo", "b")    // |foo   |b     |
    
    s := fmt.Sprintf("a %s", "string")
    fmt.Println(s)  // a string

    fmt.Fprintf(os.Stderr, "an %s\n", "error")  // an error
}
