package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func TotalArea(dimension []int) int {
	l := dimension[0]
	w := dimension[1]
	h := dimension[2]
	smallArea := l * w
	if l*h < smallArea {
		smallArea = l * h
	}
	if h*w < smallArea {
		smallArea = h * w
	}

	surfaceArea := (2 * (l * w)) + (2 * (w * h)) + (2 * (h * l))

	return surfaceArea + smallArea
}

func main() {
	file, _ := os.Open(os.Args[1])

	defer file.Close()
	input := bufio.NewReader(file)

	scanner := bufio.NewScanner(input)

	var num []int

	var output int

	for scanner.Scan() {
		line := scanner.Text()
		dimension := strings.Split(line, "x")
		for i := range dimension {
			number, err := strconv.Atoi(dimension[i])
			if err != nil {
				return
			}
			num = append(num, number)

		}
		output += TotalArea(num)
		num = []int{}
	}
	fmt.Println(output)
}
