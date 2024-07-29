package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func SaveAndMiss(arg string, num int) string {
	var res string
	var save int

	if num <= 0 {
		return arg
	}

	for i := 0; i < len(arg); i++ {
		if save < num {
			save++
			res += string(arg[i])
		} else if save == num {
			i += num - 1
			save = 0
		}
	}

	return res
}

func main() {
	structs := []struct {
		str string
		itt int
	}{
		{"123456789", 1},
		{"e 5£ @ 8* 7 =56 ;", 2},
		{"QKplq%QSw", 3},
		{"", 4},
		{"hello \\! n4ght cr3a8ure7 ", 5},
		{"Kimetsu no Yaiba", 6},
		{"8595485-52", 7},
		{"abcdefghijklmnopqrstuvwyz", 8},
		{"w58tw7474abc", 9},
		{"Po65 4o", 10},
	}

	for _, v := range structs {
		challenge.Function("SaveAndMiss", SaveAndMiss, solutions.SaveAndMiss, v.str, v.itt)
	}
}
