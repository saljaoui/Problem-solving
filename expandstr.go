package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	str := os.Args[1]
    var res string
    for i, s := range str {

        if s != ' ' {
            res += string(s)
        } else if len(str)-1 > i && s == ' ' && str[i+1] != ' ' && res != "" {
            res += "   "
        }
    }
    fmt.Println(res)
}
