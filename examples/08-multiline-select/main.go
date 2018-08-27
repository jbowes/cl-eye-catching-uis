package main

import (
	"fmt"
	"os"
)

var options = []string{
	"first option",
	"second option",
	"last option",
}

func getSpace() { fmt.Print("\n\n\n") }

func rehome() { fmt.Printf("\033[3A\r") }

func drawOptions(selected int) {
	rehome()
	fmt.Print("Select an option:")
	for i, option := range options {
		var ind string
		if i == selected {
			ind = ">"
		}

		fmt.Printf("\033[B\r  %1s %s", ind, option)
	}
}

func loop() {
	idx := 0
	for {
		drawOptions(idx)

		buf := []byte{0, 0, 0}
		n, _ := os.Stdin.Read(buf)
		if n == 0 { // This is very poor input parsing.
			continue
		}

		switch {
		case string(buf) == "\033[A": // escape sequence for up key
			if idx != 0 {
				idx--
			}
		case string(buf) == "\033[B": // escape sequence for down key
			if idx != 2 {
				idx++
			}
		case buf[0] == 'q':
			return
		default:
			fmt.Println(buf)
		}
	}
}

func main() {
	tios, _ := getTerminalIOs(os.Stdout)
	getSpace()
	makeRaw()
	fmt.Printf("\033[0m")   // Turn off character attributes, just in case
	fmt.Printf("\033[?25l") // Hide the cursor while we jump around

	loop()

	fmt.Printf("\033[?25h")         // Restore the cursor
	setTerminalIOs(os.Stdout, tios) // Restore previous terminal state
	fmt.Println()
}
