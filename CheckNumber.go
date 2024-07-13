package main

import (
	"fmt"
)

func CheckNumber(arg string) bool {
	for _, arge := range arg {
		if arge < '9' && arge > '0' {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(CheckNumber("Hello"))
	fmt.Println(CheckNumber("Hello1"))
}
