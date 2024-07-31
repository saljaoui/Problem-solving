package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func FishAndChips(n int) string {
	if n < 0 {
		return "error: number is negative"
	}
	if n%2 == 0 && n%3 == 0 {
		return "fish and chips"
	}
	if n%2 == 0 {
		return "fish"
	}
	if n%3 == 0 {
		return "chips"
	}
		return "error: non divisible"
}

func main() {
	args := []int{20, 15, -125, 10, 5, 4, -9, 9, 24, 0, -85, 6, 11}

	for i := 0; i < len(args); i++ {
		challenge.Function("FishAndChips", FishAndChips, solutions.FishAndChips, args[i])
	}
}
