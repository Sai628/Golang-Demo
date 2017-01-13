package main

import "fmt"
import "sort"

func main() {
    
    strs := []string{"c", "a", "b"}
    sort.Strings(strs)
    fmt.Println("Strings:", strs)  // Strings: [a b c]

    ints := []int{7, 2, 4}
    sort.Ints(ints)
    fmt.Println("Ints:   ",ints)  // Ints:    [2 4 7]

    s := sort.StringsAreSorted(strs)
    fmt.Println("Sorted:", s)  // Sorted: true
}
