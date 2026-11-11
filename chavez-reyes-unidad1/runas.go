package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Ñandù"
	fmt.Println("Longitud:", len(s))
	fmt.Println("runas:", utf8.RuneCountInString(s))

	for i, r := range s {
		fmt.Printf("%d:%c ", i, r)
	}
	fmt.Println()

	fmt.Println(s[0])

	var u uint = 225
	u++
	fmt.Println(u)
}
