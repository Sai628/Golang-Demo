package main

import (
    "fmt"
    "unsafe"
)

// Enum by using constants
const (
    Unknow = 0
    Female = 1
    Male = 2
)

const (
    A = "abc" 
    B = len(A)
    C = unsafe.Sizeof(A)
)

// iota
const (
    CONST_A = iota
    CONST_B
    CONST_C
    CONST_D = "ha"
    CONST_E
    CONST_F = 100
    CONST_G
    CONST_H = iota
    CONST_I
)

func main()  {
    fmt.Println(Unknow, Female, Male)  // 0 1 2
    fmt.Println(A, B, C)  // abc 3 16
    fmt.Println(CONST_A, CONST_B, CONST_C, CONST_D, CONST_E, CONST_F, CONST_G, CONST_H, CONST_I)  // 0 1 2 ha ha 100 100 7 8

    const LENGTH int = 10
    const WIDTH int = 5
    var area int
    const a, b, c = 1, false, "str"

    area = LENGTH * WIDTH
    fmt.Printf("area is: %d\n", area)  // area is: 50
    fmt.Println(a, b, c)  // 1 false str
}
