package main

import (
	"fmt"
	"unicode/utf8"

	"github.com/mattn/go-runewidth"
)

func main() {
	fmt.Println("Missing characters:")
	missing := "The power symbol is new in Unicode 9: ⏻"
	fmt.Println(missing)

	fmt.Println()
	fmt.Println("Miscounting length:")
	miscounted := "😀"
	// 4. That can really mess up positioning and wrapping!
	fmt.Println(miscounted, "has len", len(miscounted))
	fmt.Println(miscounted, "has runes", utf8.RuneCount([]byte(miscounted)))

	fmt.Println()
	fmt.Println("Wide characters:")
	wide := "十"
	fmt.Println(wide)
	fmt.Println("01")

	fmt.Println("runes:", utf8.RuneCount([]byte(wide)))
	fmt.Println("width:", runewidth.StringWidth(wide))
}
