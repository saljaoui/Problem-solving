package main

import (
	"fmt"
)

func DigitLen(n, base int) int {
	if base < 2 || base > 32 {
		return -1
	}
	var r int
	if n < 0 {
		n = n * -1
	}
	for n > 0 {
		n = n / base
		r++
	}
	return r
}

func main() {
	fmt.Println(DigitLen(100, 10))
	fmt.Println(DigitLen(100, 2))
	fmt.Println(DigitLen(-100, 16))
	fmt.Println(DigitLen(100, -1))
}
