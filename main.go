package main

import "fmt"

func increment(n *int) {
	*n++
}

func main() {
	var x int
	fmt.Scan(&x)
	increment(&x)
	fmt.Println(x)
}
