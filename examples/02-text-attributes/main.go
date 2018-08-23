package main

import "fmt"

const esc = "\033["

var resetCode = fmt.Sprintf("%s%dm ", esc, 0)

func main() {
	fmt.Printf("\033[1mbold")
	fmt.Printf(resetCode)
	fmt.Printf("\033[2mdim")
	fmt.Printf(resetCode)
	fmt.Printf("\033[3mitalics")
	fmt.Printf(resetCode)
	fmt.Printf("\033[4munderline")
	fmt.Printf(resetCode)
	fmt.Printf("\033[5mblink")
	fmt.Printf(resetCode)
	fmt.Printf("\033[7mreverse")
	fmt.Printf(resetCode)
	fmt.Printf("\033[8minvisible")
	fmt.Printf(resetCode)
	fmt.Printf("\033[9mcrossed-out")
	fmt.Printf(resetCode)

	fmt.Println()
	fmt.Println("blink, invisible, and crossed-out are usually not supported")
}
