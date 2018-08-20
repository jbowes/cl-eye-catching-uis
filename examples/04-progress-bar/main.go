package main

import (
	"fmt"
	"strings"
	"time"
)

func drawBar(percent int) {
	cols := 10
	prog := strings.Repeat("=", cols*percent/100.0)
	fmt.Printf("\rdemo progress: %3[1]d%% |%-[3]*[2]s|",
		percent, prog, cols)
}

func main() {

	status := make(chan int)
	go func() {
		for i := 0; i <= 100; i++ {
			status <- i
			time.Sleep(100 * time.Millisecond)
		}
		close(status)
	}()

	for p := range status {
		drawBar(p)
	}

	// Start a new line for the next command prompt
	fmt.Println()
}
