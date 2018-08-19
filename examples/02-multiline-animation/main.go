package main

import (
	"fmt"
	"time"
)

var options = []string{
	"first option",
	"second option",
	"last option",
}

// http://fabiensanglard.net/fizzlefade/index.php
func fizzlefade() {
	var rndval uint16 = 1
	var x, y uint8

	for {
		y = uint8(rndval & 0x00FF)        // low 5 bits, enough for 24
		x = uint8((rndval & 0x7F00) >> 8) // high 7 bits, enough for 80

		lsb := rndval & 1 // get output bit
		rndval >>= 1      // shift register

		if lsb > 0 {
			rndval ^= 0xD030

			if x < 80 && y < 24 {
				fizzlechar(x, y)
			}

			time.Sleep(0 * time.Millisecond)
		}
		if rndval == 1 {
			break
		}
	}
}

func fizzlechar(x, y uint8) {
	fmt.Printf("\033[%d;%dHX", y, x)
}

func main() {
	fizzlefade()
}
