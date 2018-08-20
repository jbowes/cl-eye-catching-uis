package main

import "fmt"

const esc = "\033["

var resetCode = fmt.Sprintf("%s%dm ", esc, 0)

func main() {
	fmt.Println("these change the whole line")
	fmt.Println("\033#3double size (top)")
	fmt.Println("\033#4double size (bottom)")
	fmt.Println("\033#6double width")
	fmt.Printf(resetCode)
}
