package main

import (
	"fmt"
)

func CountAlpha(s string) int {
	var count int
	for _, S := range s {
		if S > 'a' && S < 'z' {
			count++
		}
		if S > 'A' && S < 'Z' {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(CountAlpha("Hello world"))
	fmt.Println(CountAlpha("H e l l o"))
	fmt.Println(CountAlpha("H1e2l3l4o"))
}
