package main

import (
	"fmt"
)

func HashCode(dec string) string {
	var str string
	for _, s := range dec{
		str += string(int(s)+len(dec)%127)
	}
	return str
}

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}
