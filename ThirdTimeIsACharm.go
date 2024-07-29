package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func ThirdTimeIsACharm(str string) string {
	if str == "" || len(str) < 3 {
		return "\n"
	}
	var res string
	for i := 2; i < len(str); i += 3 {
		res += string(str[i])
	}
	return res + "\n"
}

func main() {
	table := []string{"1234556789", "QKplq%QSw", "", "Kimetsu no Yaiba", "Z", "email123@live.fr", "8595485-52", "-552", "abc", "w58tw7474abc", "fifa world cup `2022`"}

	for _, s := range table {
		challenge.Function("ThirdTimeIsACharm", ThirdTimeIsACharm, solutions.ThirdTimeIsACharm, s)
	}
}
