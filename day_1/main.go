package main

import (
	"fmt"
	"os"
	"strings"
)

func Floor(move []string) int {
	var count int

	for i := range move {
		if move[i] == "(" {
			count++
		} else {
			count--
		}
		if count==-1{
			return i
		}
	}
	return count
}

func main() {
	input, _ := os.ReadFile(os.Args[1])
	str := strings.Split(string(input), "")

	fmt.Println(Floor(str)+1)
}
