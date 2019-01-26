package main

import (
    "fmt"
    "time"
)

func main() {
    var forever bool = true
    for forever {
        t := time.Now()
        fmt.Printf("\r %v", t)
    }
}
