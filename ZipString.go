package main

import (
	"fmt"

)


func ZipString(s string) string {
count := 1
var result string
	for i:= 0 ; i < len(s); i++ {
		if i < len(s)-1 && s[i] == s[i+1] {
			count+=1
			continue
		}
		if i < len(s)-1 && s[i] != s[i+1] {
			result +=  string(count+'0')
			result += string(s[i])
			count = 1
		} else if i == len(s)-1 {
			result +=  string(count+'0')
			result += string(s[i])
		}
	}
return result
}


func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}
