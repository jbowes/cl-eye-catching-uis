package main

import (
	"fmt"
	"os"
)

func main() {
	if isTTY(os.Stdout) {
		cols, rows, _ := terminalSize(os.Stdout)
		fmt.Printf("Detected %d columns, %d rows\n", cols, rows)
	} else {
		fmt.Println("This isn't a terminal")
	}
}
