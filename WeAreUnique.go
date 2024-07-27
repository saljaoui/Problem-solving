package main

import (


	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}
	s1 := make(map[rune]bool)
	s2 := make(map[rune]bool)
	for _, s := range str1 {
		s1[s] = true
	}
	for _, s := range str2 {
		s2[s] = true
	}
	var count int
	for i := range s1 {
		for j := range s2 {
			if i == j {
				count ++
			}
		}
	}
	return (len(s1) + len(s2)) - count*2
}

func main() {
	table := [][]string{
		{"abc", "def"},
		{"hello", "yoall"},
		{"everyone", ""},
		{"hello world", "fam"},
		{"abc", "abc"},
		{"", ""},
		{"pomme", "pomme"},
		{"+265", "265"},
		{"123231", "123231"},
		{"w^p@@j", "w^p@@j"},
		{"26235e5", "4478q92"},
		{"		", "		 "},
		{"AB$%d.52", "eepqdl.52"},
		{"", "eveRyone"},
		{"_55w1se", "55w1se"},
	}
	for _, arg := range table {
		challenge.Function("WeAreUnique", WeAreUnique, solutions.WeAreUnique, arg[0], arg[1])
	}
}
