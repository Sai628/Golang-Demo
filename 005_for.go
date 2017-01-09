package main

import "fmt"

func main() {
    
    i := 1
    for i <= 3 {
        fmt.Println(i)  // 1, 2, 3
        i = i + 1
    }

    for j := 7; j <= 9; j++ {
        fmt.Println(j)  // 7, 8, 9
    }

    for {
        fmt.Println("loop")  // loop
        break
    }
}
