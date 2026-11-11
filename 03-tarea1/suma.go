package main

import "fmt"

func suma() int {
	total := 0
	for i := 1; i <= 3; i++ {
		total = total + i
		_ = total
	}
	return total
}

func main() {
	fmt.Println("total:", suma())
}
