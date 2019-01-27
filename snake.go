package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

type Game struct {
	closed chan struct{}
	snake  [3][2]int
}

func (g *Game) Start() {

	for {
		select {
		case <-g.closed:
			return
		default:
			var board [30][100]string
			board = defaultBoard()

			// Replace board positions containing snake
			// with white tiles
			for _, pos := range g.snake {
				board[pos[1]][pos[0]] = "\u001b[48;5;15m "
			}

			// Build the board
			for _, row := range board {
				fmt.Printf(join(row))
			}

			// Put the cursor back at 0,0
			fmt.Printf("\033[0;0H")

			// Move the snake forward
			g.Move()

			// Put delay in to slow render speed
			time.Sleep(50 * time.Millisecond)
		}
	}

}

func (g *Game) Move() {
	if g.snake[2][0] < 99 {

		g.snake[0][0] = g.snake[0][0] + 1
		g.snake[1][0] = g.snake[1][0] + 1
		g.snake[2][0] = g.snake[2][0] + 1

	}
}

func (g *Game) Stop() {
	close(g.closed)
}

func join(arr [100]string) string {
	// Join an array of strings with and empty string
	s := ""
	for _, str := range arr {
		s = s + str
	}

	return s + "\n"
}

func defaultBoard() [30][100]string {

	// Init the board, 30 * 100 pixels large
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
	// Create snake
	var snake [3][2]int
	snake[1][0] = 1
	snake[2][0] = 2

	game := &Game{
		closed: make(chan struct{}),
		snake:  snake,
	}

	// Hide cursor
	fmt.Printf("\033[?25l")
	// Clear screen
	fmt.Printf("\033[2J")
	// Move cursor to top left
	fmt.Printf("\033[0;0H")

	// On exit, make sure cursor is turned back on
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		// Retrun cursor
		fmt.Printf("\033[?25h")
		// Clear screen
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
		fmt.Printf("\033[H\033[2J")
		game.Stop()
	}()

	game.Start()

}
