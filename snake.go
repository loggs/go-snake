package main

import (
    "fmt"
    "time"
    "os"
    "os/signal"
    "syscall"
)


func main() {
    // Hide cursor
    fmt.Printf("\033[?25l")
    // Clear screen
    fmt.Printf("\033[2J")
    // Move cursor to top left
    fmt.Printf("\033[0;0H")

    // On exit, make sure cursor is turned back on
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    go func(){
        <-c
        fmt.Printf("\033[?25h")
        os.Exit(1)
    }()
   
    // Exit game code
    for {
        t := time.Now()
        fmt.Printf("\033[1m Hello: %v\n", t)
        fmt.Printf("\033[1m Hello: %v\n", t)
        fmt.Printf("\033[1m Hello: %v\n", t)
        fmt.Printf("\033[0;0H")
    }
}
