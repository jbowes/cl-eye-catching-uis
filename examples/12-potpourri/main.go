package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/campoy/tools/imgcat"
)

func loop(draw func(x, y int)) {
	for {

		buf := [6]byte{}
		os.Stdin.Read(buf[:])

		// This is very poor input parsing
		switch {
		case string(buf[:3]) == "\033[M": // Mouse info!
			// Mouse coordinates are encoded in a byte with their value +32
			// so they are printable characters.
			// They start at 1,1. Cursor positioning starts at 0,0
			draw(int(buf[4]-33), int(buf[5]-33))
		case buf[0] == 'q':
			return
		}
	}
}

func main() {
	tios, _ := getTerminalIOs(os.Stdout)
	makeRaw()
	fmt.Printf("\033[0m")   // Turn off character attributes, just in case
	fmt.Printf("\033[?25l") // Hide the cursor while we jump around

	fmt.Printf("\033[?1049h") // enable alternate screen buffer
	fmt.Printf("\033[2J")     // clear the screen

	fmt.Printf("\033[?1003h") // track mouse coordinates

	draw := drawMouse
	if os.Getenv("TERM_PROGRAM") == "iTerm.app" {
		draw = drawForITerm()
	}

	loop(draw)

	fmt.Printf("\033[?1003l")       // disable mouse tracking
	fmt.Printf("\033[?1049l")       // return to original buffer
	fmt.Printf("\033[?25h")         // Restore the cursor
	setTerminalIOs(os.Stdout, tios) // Restore previous terminal state
}

func drawMouse(x, y int) {
	fmt.Printf("\033[D ")           // clear the last position
	fmt.Printf("\033[%d;%dH", y, x) // move cursor
	fmt.Print("🐭")
}

func drawForITerm() func(x, y int) {
	var buf bytes.Buffer
	enc, err := imgcat.NewEncoder(&buf,
		imgcat.Inline(true),
		imgcat.Width(imgcat.Cells(8)),
		imgcat.Height(imgcat.Auto()),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	f, err := os.Open("12-potpourri/gophercon.jpg")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	enc.Encode(f)

	return func(x, y int) {
		fmt.Printf("\033[2J")           // clear the screen
		fmt.Printf("\033[%d;%dH", y, x) // move cursor
		fmt.Print(buf.String())
	}
}
