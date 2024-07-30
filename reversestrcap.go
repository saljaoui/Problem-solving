package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	for _, text := range args {
		var res []string
		var resl string
		for i, s := range text {
			if i == 0 && s == ' ' {
				resl += " "
			}
			if i == len(text)-1 && s == ' ' {
				resl += " "
			}
			if s != ' ' {
				resl += string(s)
			} else if resl != "" {
				res = append(res, resl)
				resl = ""
			}
		}
		if resl != "" {
			res = append(res, resl)
		}
		var reslt string
		for i, ss := range res {
			n := len(ss)
			if n > 0 {
				reslt += Low(ss[:n-1])
				reslt += Up(ss[n-1:])
				if i < len(res)-1 {
					reslt += " "
				}
			}
		}
		fmt.Println(reslt)
	}
}

func Up(s string) string {
	var res string
	for _, ss := range s {
		if ss >= 'A' && ss <= 'Z' {
			res += string(ss)
		} else if ss >= 'a' && ss <= 'z' {
			res += string(ss - 32)
		} else {
			res += string(ss)
		}
	}
	return res
}

func Low(s string) string {
	var res string
	for _, ss := range s {
		if ss >= 'A' && ss <= 'Z' {
			res += string(ss + 32)
		} else if ss >= 'a' && ss <= 'z' {
			res += string(ss)
		} else {
			res += string(ss)
		}
	}
	return res
}
