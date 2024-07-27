package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	str1 := os.Args[1]
	str2 := os.Args[2]
	var str2arry []string
	for _, s2 := range str2 {
		str2arry = append(str2arry, string(s2))
	}
	var count int
	var oldi int = 0
	for _, s1 := range str1 {
		for i, s2 := range str2arry {
			if s2 == string(s1) {
				if i >= oldi {
					oldi = i
					str2arry[i] = ""
					count++
					break
				}
			}
		}
	}
	if count >= len(str1) {
		fmt.Println(str1)
	}
}
