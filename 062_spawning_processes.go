package main

import "fmt"
import "io/ioutil"
import "os/exec"

func main() {

    dateCmd := exec.Command("date")
    dateOut, err := dateCmd.Output()
    if err != nil {
        panic(err)
    }

    fmt.Println("> date")
    fmt.Println(string(dateOut))
    // > date
    // Mon Jan 16 18:31:56 CST 2017


    grepCmd := exec.Command("grep", "hello")
    grepIn, _ := grepCmd.StdinPipe()
    grepOut, _ := grepCmd.StdoutPipe()
    grepCmd.Start()
    grepIn.Write([]byte("hello grep\ngoodbye grep"))
    grepIn.Close()
    grepBytes, _ := ioutil.ReadAll(grepOut)
    grepCmd.Wait()

    fmt.Println("> grep hello")
    fmt.Println(string(grepBytes))
    // > grep hello
    // hello grep


    lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
    lsOut, err := lsCmd.Output()
    if err != nil {
        panic(err)
    }

    fmt.Println("> ls -a -l -h")
    fmt.Println(string(lsOut))
    // > ls -a -l -h
    // total 512
    // drwxr-xr-x  68 Sai  staff   2.3K Jan 16 18:21 .
    // drwxr-xr-x   4 Sai  staff   136B Jan 14 16:41 ..
    // ...
    // -rw-r--r--   1 Sai  staff   801B Jan 16 18:31 062_spawning_processes.go
    // ...
}
