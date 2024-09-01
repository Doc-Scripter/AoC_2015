package main

import (
	"bufio"
	"fmt"
	"os"
)

func Deliver(str string) int {
	countSanta := 1
	countRobo := 0
	var position [2]int
	var position2 [2]int

	house := make(map[[2]int]bool, 0)
	var x, y int
	var k, l int
	position = [2]int{x, y}
	house[position] = true
	
	
	for j := 0; j < len(str); j++ {
		if j%2 == 0 {
			
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
			
			position = [2]int{x, y}
			
			if !house[position] {
				countSanta++
				house[position] = true
			}
		} else if j%2 != 0 {
			switch string(str[j]) {
			case "^":
				k++
			case "v":
				k--
			case "<":
				l--
			case ">":
				l++
			}
			position2 = [2]int{k, l}
			if !house[position2] {
				
				countRobo++
				house[position2] = true
				
			}
		}
	}

	count := countSanta + countRobo

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
