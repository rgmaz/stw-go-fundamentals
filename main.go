package main

import "fmt"

const (
	fizz     = "Fizz"
	buzz     = "Buzz"
	fizzBuzz = fizz + buzz
)

func main() {
	var n int
	fmt.Scan(&n)

	switch {
	case n%15 == 0:
		fmt.Println(fizzBuzz)
	case n%5 == 0:
		fmt.Println(buzz)
	case n%3 == 0:
		fmt.Println(fizz)
	default:
		fmt.Println(n)
	}
}
