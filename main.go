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

	words := strings.FieldsSeq(line)
	distinctWords := make(map[string]int)

	for w := range words {
		distinctWords[w]++
	}

	fmt.Println(len(distinctWords))
}
