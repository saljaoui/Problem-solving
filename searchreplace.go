package main

import "os"

func main() {
    if len(os.Args) != 4 {
        return
    }
	str := os.Args[1:]
    var res string

    for _, s :=  range str[0] {
        if string(s) == str[1] {
            res += str[2]
        } else {
            res += string(s)
        }
    }
    println(res)
}
