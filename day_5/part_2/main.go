package main

import (
	"bufio"
	"fmt"
	"os"
)

func CountPairs(str string) int {
	var count int

	for i := 0; i < len(str)-1; i++ {
		first := str[i:i+2]
		for j := i + 2; j < len(str)-1; j++ {
			if first == str[j:j+2] {
				fmt.Println(first,str[j:j+2])
				count++
			}
		}
	}

	return count
}

func checkPairs(str string) bool {
	// var count int
	count := CountPairs(str)
	// fmt.Println(count)
	return count >= 1
}

func checkRepeat(str string) bool {
	count := 0
	for i := 0; i < len(str)-1; i++ {
		if i+2 < len(str) && str[i] == str[i+2] {
			count++
		}
	}
	// fmt.Println(count)
	return count >= 1
}

// checkLetters validates the string based on specific criteria.
func isValid(str string) bool {
	if !checkPairs(str) || !checkRepeat(str) {
		return false
	}

	return true
}

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	Scanner := bufio.NewScanner(file)
	var count int

	// Scan each line in the file
	for Scanner.Scan() {
		line := Scanner.Text()
		if isValid(line) {
			// fmt.Println(line)
			count++
		}
	}

	fmt.Println(count)
}
