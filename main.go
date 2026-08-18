package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)

	words := strings.Fields(line)
	distinctWords := make(map[string]int)

	for _, w := range words {
		distinctWords[w]++
	}

	fmt.Println(len(distinctWords))
}
