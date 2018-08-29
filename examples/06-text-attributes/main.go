package main

import "fmt"

func reset() { fmt.Print(" \033[0m") }

func main() {
	fmt.Printf("\033[1mbold")
	reset()

	fmt.Printf("\033[2mdim")
	reset()

	fmt.Printf("\033[3mitalics")
	reset()

	fmt.Printf("\033[4munderline")
	reset()

	fmt.Printf("\033[5mblink")
	reset()

	fmt.Printf("\033[7mreverse")
	reset()

	fmt.Printf("\033[8minvisible") // don't use this for hiding passwords
	reset()

	fmt.Printf("\033[9mcrossed-out")
	reset()

	fmt.Println()
	fmt.Println("blink, invisible, and crossed-out are usually not supported")
}
