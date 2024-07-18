package main

import (
	"fmt"
)

func CamelToSnakeCase(s string) string {
	var str []string
	for i, ss := range s {
		if i+1 < len(s) && ss >= 'A' && ss <= 'Z' && i != 0 && !IsUp(s[i+1]) && !IsUp(s[i-1]) {
			str = append(str, "_")
		}
		str = append(str, string(ss))
	}
	var re string
	for _, r := range str {
		re += r
	}

	return re
}

func IsUp(s byte) bool {
	ss := rune(s)
	if ss >= 'A' && ss <= 'Z' {
		return true
	}
	return false
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}
