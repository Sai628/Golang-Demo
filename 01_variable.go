package main

var x, y int
var (
    a int
    b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

func main() {
    g, h := 123, "hello"
    println(x, y, a, b, c, d, e, f, g, h) // 0 0 0 false 1 2 123 hello 123 hello
}
