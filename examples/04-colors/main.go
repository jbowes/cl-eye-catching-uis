package main

import "fmt"

const esc = "\033["

type attribute int

const reset attribute = 0

// The possible colors of text inside the application.
//
// These constants are called through the use of the Styler function.
const (
	FGBlack attribute = iota + 30
	FGRed
	FGGreen
	FGYellow
	FGBlue
	FGMagenta
	FGCyan
	FGWhite
)

// The possible background colors of text inside the application.
//
// These constants are called through the use of the Styler function.
const (
	BGBlack attribute = iota + 40
	BGRed
	BGGreen
	BGYellow
	BGBlue
	BGMagenta
	BGCyan
	BGWhite
)

var resetCode = fmt.Sprintf("%s%dm ", esc, reset)

func eight(attrs ...int) {
	for i := 0; i < 8; i++ {
		var extra string
		if len(attrs) > 0 {
			extra = fmt.Sprintf("%d;", attrs[0])
		}
		fmt.Printf("\033[%s%dm█", extra, i+30)
	}
	fmt.Printf(resetCode)
	fmt.Println()
}

func main() {
	fmt.Printf("original 8: ")
	eight()

	fmt.Printf("bold 8:     ")
	fmt.Printf("\033[1m")
	eight(1)

	fmt.Printf("dim 8:      ")
	fmt.Printf("\033[2m")
	eight(2)

	fmt.Println("Bold and dim only work for foreground colors")
	fmt.Println()

	fmt.Println("256 colors:")

	fmt.Print("original 8:       ")
	for i := 0; i < 8; i++ {
		fmt.Printf("\033[48;5;%dm ", i)
	}
	fmt.Printf(resetCode)
	fmt.Println()

	fmt.Print("high intensity 8: ")
	for i := 8; i < 16; i++ {
		fmt.Printf("\033[48;5;%dm ", i)
	}
	fmt.Printf(resetCode)
	fmt.Println()
	fmt.Println()

	fmt.Println("6 x 6 x 6 pallette:")
	for i := 1; i < 217; i++ {
		fmt.Printf("\033[48;5;%dm ", i+15)
		if i%36 == 0 {
			fmt.Printf(resetCode)
			fmt.Println()
		}
	}
	fmt.Printf(resetCode)
	fmt.Println()

	fmt.Print("greyscale: ")
	for i := 232; i < 256; i++ {
		fmt.Printf("\033[48;5;%dm ", i)
	}
	fmt.Printf(resetCode)
	fmt.Println()
	fmt.Println()

	fmt.Println("24 bit true color:")
	for i := 0; i < 8; i++ {
		for j := 0; j < 64; j++ {
			fmt.Printf("\033[48;2;%d;0;%dm ", i*4+128, j*2)
		}
		fmt.Println(resetCode)
	}
}
