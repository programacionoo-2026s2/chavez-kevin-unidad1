package main

import "fmt"

func main() {
	var fizz, buzz, fizzbuzz int

	for i := 1; i <= 30; i++ {
		switch {
		case i%15 == 0:
			fmt.Print("FizzBuzz ")
			fizzbuzz++
		case i%3 == 0:
			fmt.Print("Fizz ")
			fizz++
		case i%5 == 0:
			fmt.Print("Buzz ")
			buzz++
		default:
			fmt.Print(i, " ")
		}
	}

	fmt.Printf("\nFizz=%d Buzz=%d FizzBuzz=%d\n", fizz, buzz, fizzbuzz)
}
