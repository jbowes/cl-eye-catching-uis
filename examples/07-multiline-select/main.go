package main

import (
	"fmt"
)

var options = []string{
	"first option",
	"second option",
	"last option",
}

func getSpace() { fmt.Print("\n\n\n\n") }

func rehome() { fmt.Printf("\033[4A") }

func drawOptions(selected int) {
	rehome()
	fmt.Println("Select an option:")
	for i, option := range options {
		var ind string
		if i == selected {
			ind = ">"
		}

		fmt.Printf("  %1s %s\n", ind, option)
	}
}

func main() {
	getSpace()
	for {
		drawOptions(0)
	}
}
