package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	name, _ := r.ReadString('\n')
	name = strings.TrimRight(name, "\r\n")
	ageStr, _ := r.ReadString('\n')
	ageStr = strings.TrimRight(ageStr, "\r\n")
	age, _ := strconv.Atoi(ageStr)

	fmt.Printf("Hi, %s! You are %d years old.", name, age)
}
