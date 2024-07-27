package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println()
		return
	}

	str1 := os.Args[1]
	str2 := os.Args[2]
	
	strall := str1 + str2

	seen := make(map[rune]bool)
	result := ""

	for _, char := range strall {
		if !seen[char] {
			seen[char] = true
			result += string(char)
		}
	}
	fmt.Println(result)
}
