package main

import (
	"fmt"
	"os"
)

func main() {
	s := os.Args[1]
	var str string
	var Str []string

	for _, f := range s {
		if f == '(' || f == '[' || f == '{' || f == ')' || f == ']' || f == '}' {
			str += string(f)
			Str = append(Str, string(f))
		}
	}

	for i := 0; i < len(Str)-1; i++ {
		if (Str[i] == "(" && Str[i+1] == ")") ||
			(Str[i] == "[" && Str[i+1] == "]") ||
			(Str[i] == "{" && Str[i+1] == "}") {
			Str = append(Str[:i], Str[i+2:]...)
			i = -1
		}
	}

	if len(Str) == 0 {
		fmt.Println("OK")
		return
	}
	fmt.Println("Error")
}
