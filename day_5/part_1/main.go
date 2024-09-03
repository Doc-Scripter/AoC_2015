package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// IsVowel checks if a rune is a vowel.
func IsVowel(r rune) bool {
	vowels := []rune{'a', 'e', 'i', 'o', 'u'}
	for _, v := range vowels {
		if v == r {
			return true
		}
	}
	return false
}

// checkLetters validates the string based on specific criteria.
func checkLetters(str string) bool {
	// Check for invalid substrings
	if strings.Contains(str, "ab") || strings.Contains(str, "cd") || strings.Contains(str, "pq") || strings.Contains(str, "xy") {
		return false
	}

	var count int
	double := false

	// Iterate through the string and count vowels and check for double letters
	for i := 0; i < len(str); i++ {
		if IsVowel(rune(str[i])) {
			count++
		}
		if i != len(str)-1 && str[i] == str[i+1] {
			double = true
		}
	}

	// Check for exactly 3 vowels and at least one pair of adjacent letters
	if count <3|| !double {
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
		if checkLetters(line) {
			count++
		}
	}

	fmt.Println(count)
}
