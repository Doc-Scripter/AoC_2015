package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func toggleLights(instructions map[[2]int]string, currentGrid map[[2]int]string) int {
	var count int
	for i:=0;i<1000;i++{
		for j:=0;j<1000;j++{
			if currentGrid[[2]int{i,j}] == "off"&&instructions[[2]int{i,j}] == "on"{
				currentGrid[[2]int{i,j}] = "tn"
				count++
			}
		}
	}
	fmt.Println(currentGrid)
	return count
}

func originalGrid() map[[2]int]string {
	origin := make(map[[2]int]string)

	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			origin[[2]int{i, j}] = "off"
		}
	}

	return origin
}

func getInstructions(s []string) map[[2]int]string {
	commands := make(map[[2]int]string)

	if s[0] == "toggle" {
		command := "toggle"
		startx, _ := strconv.Atoi(s[1])
		starty, _ := strconv.Atoi(s[2])
		endx, _ := strconv.Atoi(s[4])
		endy, _ := strconv.Atoi(s[5])

		for i := startx; i <= endx; i++ {
			for j := starty; i <= endy; i++ {
				commands[[2]int{i, j}] = command
			}
		}

	}
	if s[0] == "turn" && s[1] == "off" {
		command := "tff"
		startx, _ := strconv.Atoi(s[2])
		starty, _ := strconv.Atoi(s[3])
		endx, _ := strconv.Atoi(s[5])
		endy, _ := strconv.Atoi(s[6])

		for i := startx; i <= endx; i++ {
			for j := starty; i <= endy; i++ {
				commands[[2]int{i, j}] = command
			}
		}

	}
	if s[0] == "turn" && s[1] == "on" {
		command := "on"
		startx, _ := strconv.Atoi(s[2])
		starty, _ := strconv.Atoi(s[3])
		endx, _ := strconv.Atoi(s[5])
		
		endy, _ := strconv.Atoi(s[6])

		for i := startx; i <= endx; i++ {
			for j := starty; j <= endy; j++ {
				commands[[2]int{i, j}] = command
			}
		}
		return commands
	}

	return commands
}

func countLights(s []string) int {
	grid:=originalGrid()
	instructions := getInstructions(s)
	// fmt.Println(instructions)
	count:=toggleLights(instructions,grid)
	return count
}

func main() {
	file, _ := os.Open("file.txt")

	defer file.Close()

	input := bufio.NewReader(file)

	scan := bufio.NewScanner(input)

	var out int

	for scan.Scan() {
		line := scan.Text()
		input := strings.Split(line, ",")
		inputStr := strings.Join(input, " ")
		input = strings.Fields(inputStr)
		out = countLights(input)
		// count += count
	}
	fmt.Println(out)
}
