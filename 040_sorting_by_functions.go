package main

import "fmt"
import "sort"

type ByLength []string

func (s ByLength) Len() int {
    return len(s)
}

func (s ByLength) Less(i, j int) bool {
    return len(s[i]) < len(s[j])
}

func (s ByLength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func main() {
    fruites := []string{"peach", "banana", "kiwi"}
    sort.Sort(ByLength(fruites))
    fmt.Println(fruites)  // [kiwi peach banana]
}
