package main

import "os"
import "os/exec"
import "syscall"

func main() {

    binary, lookErr := exec.LookPath("ls")
    if lookErr != nil {
        panic(lookErr)
    }

    args := []string{"ls", "-a", "-l", "-h"}
    env := os.Environ()

    execErr := syscall.Exec(binary, args, env)
    if execErr != nil {
        panic(execErr)
    }

    // total 520
    // drwxr-xr-x  69 Sai  staff   2.3K Jan 16 18:36 .
    // drwxr-xr-x   4 Sai  staff   136B Jan 14 16:41 ..
    // ...
    // -rw-r--r--   1 Sai  staff   345B Jan 16 21:03 063_executing_processes.go
    // ...
}
