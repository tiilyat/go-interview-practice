package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Read input from standard input
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()

		// Call the ReverseString function
		output := ReverseString(input)

		// Print the result
		fmt.Println(output)
	}
}

// ReverseString returns the reversed string of s.
func ReverseString(s string) string {
	if s == "" {
		return s
	}

	runeSlice := []rune(s)
	left := 0
	right := len(runeSlice) - 1

	for left < right {
		temp := runeSlice[left]
		runeSlice[left] = runeSlice[right]
		runeSlice[right] = temp

		left++
		right--
	}

	return string(runeSlice)
}
