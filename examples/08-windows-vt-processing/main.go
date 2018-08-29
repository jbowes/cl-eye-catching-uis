package main

import "fmt"

func reset() { fmt.Print(" \033[0m") }

func main() {
	initOutput()

	fmt.Printf("\033[1mbold")
	reset()

	fmt.Printf("\033[4munderline")
	reset()

	fmt.Printf("\033[7mreverse")
	reset()

	fmt.Println()
}
