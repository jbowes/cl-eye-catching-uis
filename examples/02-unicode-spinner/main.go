package main

import (
	"fmt"
	"time"
)

var (
	braille = []rune{'⠋', '⠙', '⠹', '⠸', '⠼', '⠴', '⠦', '⠧', '⠇', '⠏'}
	clock   = []rune{'🕒', '🕓', '🕔', '🕕', '🕖', '🕗', '🕘', '🕙', '🕚', '🕛', '🕐', '🕑'}
)

func drawSpinners(i int) {
	fmt.Printf("\rdrawing spinners: %c  %c",
		braille[i%len(braille)],
		clock[i%len(clock)])
}

func main() {

	// This is the same setup as for the progress bar
	status := make(chan int)
	go func() {
		for i := 0; i <= 100; i++ {
			status <- i
			time.Sleep(100 * time.Millisecond)
		}
		close(status)
	}()

	for p := range status {
		// If we don't know our progress, we can just spin or
		// pulse something. We do need something to hold state though.
		drawSpinners(p)
	}

	// Start a new line for the next command prompt
	fmt.Println()
}
