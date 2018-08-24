package main

import (
	"fmt"
	"time"
)

// http://fabiensanglard.net/fizzlefade/index.php
func fizzlefade() {
	var rndval uint32 = 1
	var x, y uint8

	for {
		y = uint8(rndval & 0x00FF)        // low 8 bits, enough for 24
		x = uint8((rndval & 0xFF00) >> 8) // high 8 bits, enough for 80

		lsb := rndval & 1 // get output bit
		rndval >>= 1      // shift register

		if lsb > 0 {
			rndval ^= 0x00012000
		}

		if x < 80 && y < 24 {
			fizzlechar(x+1, y+1)
			time.Sleep(time.Millisecond)
		}

		if rndval == 1 {
			break
		}
	}
}

func fizzlechar(x, y uint8) {
	// Write a space in column x, row y
	// This will overwrite any existing value there
	fmt.Printf("\033[%d;%dH ", y, x)
}

func main() {
	fmt.Printf("\033[0m")   // Turn off character attributes, just in case
	fmt.Printf("\033[?25l") // Hide the cursor while we jump around

	fizzlefade()
	time.Sleep(100 * time.Millisecond)

	fmt.Printf("\033[2J")   // Clear the entire screen
	fmt.Printf("\033[H")    // Move to the top left
	fmt.Printf("\033[?25h") // Restore the cursor
}
