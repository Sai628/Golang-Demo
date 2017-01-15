package main

import bs64 "encoding/base64"
import "fmt"

func main() {

    data := "abc123!?$*&()'-=@~"

    sEnc := bs64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println(sEnc)  // YWJjMTIzIT8kKiYoKSctPUB+

    sDec, _ := bs64.StdEncoding.DecodeString(sEnc)
    fmt.Println(string(sDec))  // abc123!?$*&()'-=@~

    uEnc := bs64.URLEncoding.EncodeToString([]byte(data))
    fmt.Println(uEnc)  // YWJjMTIzIT8kKiYoKSctPUB-

    uDec, _ := bs64.URLEncoding.DecodeString(uEnc)
    fmt.Println(string(uDec))  // abc123!?$*&()'-=@~
}
