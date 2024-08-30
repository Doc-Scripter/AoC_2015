package main

import (
	"bufio"
	"fmt"
	"os"
)

func Deliver(str string) int {
	count := 1
	var inner [2]int
	house := make(map[[2]int]bool, 0)
	var x, y int
	house[[2]int{x, y}] = true

	for i := 0; i < len(str)-1; i++ {
		switch string(str[i]) {
		case "^":
			x++
		case "v":
			x--
		case "<":
			y--
		case ">":
			y++
		}
		inner = [2]int{x, y}
		if !house[inner] {
			count++
			house[inner] = true
		}
	}

	return count
}

func main() {
	file, _ := os.Open(os.Args[1])

	defer file.Close()

	input := bufio.NewReader(file)

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(Deliver(line))

	}
}
