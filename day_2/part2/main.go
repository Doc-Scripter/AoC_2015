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
	perimeter := 2 * (l + w)

	if l*h <= smallArea && l*h <= w*h {
		perimeter = 2 * (l + h)
	} else if h*w <= l*h && h*w <= smallArea {
		perimeter = 2 * (h + w)
	}

	cubic := h * w * l

	return cubic + perimeter
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
