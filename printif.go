package main

import (
	"fmt"
)


func PrintIf(str string) string {
	// fmt.Println(len(str))
	if len(str) >= 1 &&  len(str) <= 3 {
		return "Invalid Output\n"
	}
	return "G\n"

}


func main() {
	fmt.Print(PrintIf("abcdefz"))
	fmt.Print(PrintIf("abc"))
	fmt.Print(PrintIf(""))
	fmt.Print(PrintIf("14"))
}
