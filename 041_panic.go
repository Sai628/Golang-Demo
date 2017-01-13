package main

import "os"

func main() {

    panic("a problem")    

    _, err := os.Create("/tmp/file")
    if err != nil {
        panic(err)
    }

    // panic: a problem

    // goroutine 1 [running]:
    // panic(0x5cae0, 0xc42000a290)
    //          /usr/local/go/src/runtime/panic.go:500 +0x1a1
    // main.main()
    //         /.../041_panic.go:7 +0x6d
    // exit status 2
}
