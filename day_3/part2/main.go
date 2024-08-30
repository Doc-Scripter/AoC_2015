package main

import (
	"bufio"
	"fmt"
	"os"
)

func Deliver(str string) int {
	countSanta := 1
	countRobo:=1
	var position [2]int
	var position2 [2]int
	house := make(map[[2]int]bool, 0)
	var x, y int
	var k,l int

	for j := 1; j < len(str)-1; j=j+2 {
		fmt.Println(string(str[j]))
		switch string(str[j]) {
		case "^":
			x++
		case "v":
			x--
		case "<":
			y--
		case ">":
			y++
		}
		position2 = [2]int{x, y}
		if !house[position2] {
			countRobo++
			house[position2] = true
		}
	}
	for i := 0; i < len(str)-1; i=i+2 {
		fmt.Println(string(str[i]))
		switch string(str[i]) {
		case "^":
			k++
		case "v":
			k--
		case "<":
			l--
		case ">":
			l++
		}
		position = [2]int{k, l}
		if !house[position] {
			countSanta++
			house[position] = true
		}
	}
	// fmt.Println(countRobo)
	count:=countSanta+countRobo
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
