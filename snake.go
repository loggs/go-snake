package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
)

func defaultBoard() [30][100]string {
   var board [30][100]string
    // Populate the board with 
    for i := 0; i < 30; i++ {
        for j := 0; j < 100; j++ {
            board[i][j] = "\033[48;5;239m "
        }
    }

    return board

}

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
        fmt.Printf("\033[2J")
        fmt.Printf("\033[?25h")
        os.Exit(1)
    }()

    var snake [3][2]int
    snake[1][0] = 1
    snake[2][0] = 2


    for {

        var board[30][100]string
        board = defaultBoard()
 
        // Replace board positions containing snake
        // with white tiles
        for _, pos := range snake {
            board[pos[1]][pos[0]] = "\u001b[48;5;15m "
        }
 
        // Build the board
        for _, row := range board {
            s := ""
            for _, str := range row {
                s = s + str
            }
            fmt.Printf(s + "\n")
        }

        fmt.Printf("\033[0;0H")

        if snake[2][0] < 99 {

            snake[0][0] = snake[0][0] + 1
            snake[1][0] = snake[1][0] + 1
            snake[2][0] = snake[2][0] + 1

        }

        // fmt.Printf("\033[1m\u001b[48;5;15m    \u001b[48;5;239m%v\n", strings.Repeat(" ", 96))
        // fmt.Printf("\033[1m\u001b[48;5;239m%v\n", strings.Repeat(" ", 100))
        // fmt.Printf("\033[1m\u001b[48;5;239m%v\n", strings.Repeat(" ", 100))
        // fmt.Printf("\033[1m\u001b[48;5;239m%v\n", strings.Repeat(" ", 100))
        // fmt.Printf("\033[1m\u001b[48;5;239m%v\n", strings.Repeat(" ", 100))
        // fmt.Printf("\033[1m\u001b[48;5;239m%v\n", strings.Repeat(" ", 100))
        // fmt.Printf("\033[1m\u001b[48;5;239m%v\n", strings.Repeat(" ", 100))
        // fmt.Printf("\033[1m\u001b[48;5;239m%v\n", strings.Repeat(" ", 100))
    }
}
