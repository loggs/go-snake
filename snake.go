package main

import (
    "fmt"
    "time"
)

func main() {
    var forever bool = true
    fmt.Printf("\033[2J")
    fmt.Printf("\033[0;0H")
    for forever {
        t := time.Now()
        fmt.Printf("\033[1m Hello: %v\n", t)
        fmt.Printf("\033[1m Hello: %v\n", t)
        fmt.Printf("\033[1m Hello: %v\n", t)
        fmt.Printf("\033[0;0H")
    }
}
