package main

import "fmt"

func printWithCode(code int, val string) {
	fmt.Printf("\033[%dm%s", code, val)
}

func reset() { printWithCode(0, "\n") }

func eight() {
	for i := 0; i < 8; i++ {
		printWithCode(i+30, "█")
	}
	reset()
}

func print256BGColor(code int) { fmt.Printf("\033[48;5;%dm ", code) }

func main() {
	fmt.Printf("original 8: ")
	eight()

	fmt.Printf("bold 8:     ")
	printWithCode(1, "")
	eight()

	fmt.Printf("dim 8:      ")
	printWithCode(2, "")
	eight()

	fmt.Println("Bold and dim only work for foreground colors")
	fmt.Println()

	fmt.Println("256 colors:")

	fmt.Print("original 8:       ")
	for i := 0; i < 8; i++ {
		print256BGColor(i)
	}
	reset()

	fmt.Print("high intensity 8: ")
	for i := 8; i < 16; i++ {
		print256BGColor(i)
	}
	reset()
	fmt.Println()

	fmt.Println("6 x 6 x 6 pallette:")
	for i := 1; i < 217; i++ {
		print256BGColor(i + 15)
		if i%36 == 0 {
			reset()
		}
	}
	reset()

	fmt.Print("greyscale: ")
	for i := 232; i < 256; i++ {
		print256BGColor(i)
	}
	reset()
	fmt.Println()

	fmt.Println("24 bit true color:")
	for i := 0; i < 8; i++ {
		for j := 0; j < 64; j++ {
			fmt.Printf("\033[48;2;%d;%d;%dm ", i*4+128, 0, j*2)
		}
		reset()
	}
}
