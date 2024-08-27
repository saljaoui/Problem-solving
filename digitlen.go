package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func DigitLen(n, base int) int {
	if base < 2 || base > 36 {
		return -1
	}
	if n < 0 {
		n = -n
	}
	cont := 0
	for n != 0 {
		cont++
		n = n / base
	}
	return cont

}

func main() {
	table := [][]int{
		{10, 2},
		{10, 3},
		{10, 4},
		{10, 5},
		{10, 6},
		{100, -1},
		{100, 10},
		{-100, 10},
		{-45, -1},
		{-100, 2},
		{456, 2},
		{-233, 15},
		{-100, 8},
		{-432, 16},
		{100, 36},
		{-100, 36},
		{100, 37},
		{-100, 37},
		{100, 1},
		{-100, 1},
	}

	for _, arg := range table {
		challenge.Function("DigitLen", DigitLen, solutions.DigitLen, arg[0], arg[1])
	}
}

