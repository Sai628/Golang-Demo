package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {

    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        ucl := strings.ToUpper(scanner.Text())
        fmt.Println(ucl)
    }

    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error", err)
        os.Exit(1)
    }

    // To try out our line filter, first make a file with a few lowercase lines, 
    // then use the line filter to get uppercase lines.

    // $ echo "hello" > /tmp/lines
    // $ echo "filter" >> /tmp/lines

    // $ cat /tmp/lines | go run 058_line_filters.go
    // HELLO
    // FILTER
}
