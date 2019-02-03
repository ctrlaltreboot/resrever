package main

import (
	"fmt"
	"os"
	"strings"
)

type input []string

func (r input) reverse() {
	strInput := strings.Join(r[:], " ")
	sLen := len(strInput)
	revStr := make([]rune, sLen)
	for i, v := range strInput {
		revStr[sLen-(i+1)] = v
	}
	fmt.Printf("%s", string(revStr))
}

func main() {
	word := input(os.Args[1:])
	word.reverse()
}
