package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("\033[?1049h") // enable alternate screen buffer
	fmt.Printf("\033[2J")     // clear the screen
	fmt.Printf("\033[12;20H")
	fmt.Printf("A whole new world!")
	time.Sleep(10 * time.Second)
	fmt.Printf("\033[?1049l") // return to original buffer
}
