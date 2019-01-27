package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
    "strings"
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
        fmt.Printf("\033[1m\u001b[48;5;239m%v\n", strings.Repeat(" ", 100))
        fmt.Printf("\033[1m\u001b[48;5;239m%v\n", strings.Repeat(" ", 100))
        fmt.Printf("\033[1m\u001b[48;5;239m%v\n", strings.Repeat(" ", 100))
        fmt.Printf("\033[1m\u001b[48;5;239m%v\n", strings.Repeat(" ", 100))
        fmt.Printf("\033[1m\u001b[48;5;239m%v\n", strings.Repeat(" ", 100))
        fmt.Printf("\033[1m\u001b[48;5;239m%v\n", strings.Repeat(" ", 100))
        fmt.Printf("\033[1m\u001b[48;5;239m%v\n", strings.Repeat(" ", 100))
        fmt.Printf("\033[1m\u001b[48;5;239m%v\n", strings.Repeat(" ", 100))
        fmt.Printf("\033[0;0H")
    }
}
