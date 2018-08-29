package main

import "fmt"

func main() {
	fmt.Println("these change the whole line")
	fmt.Println("\033#3double size (top)")
	fmt.Println("\033#4double size (bottom)")
	fmt.Println("\033#6double width")
	fmt.Print("\033[0m")
}
