package main

import (
	"fmt"
	"os"
)

func RepeatAlpha(s string) string {
	if s == "" {
		return ""
	}
	var res string
	for _, str := range s {
		if (str >= 'a' && str <= 'z')  {
			for i := 1; i <= int(str-'a'+1); i++ {
				res += string(str)
			}
		} else if (str >= 'A' && str <= 'Z') {
            for i := 1; i <= int(str-'A'+1); i++ {
				res += string(str)
			}
			
		} else {
            res += string(str)
        }

	}

	return res
}

func main() {
	testCases := []struct {
		in   string
		want string
	}{
		{"abc", "abbccc"},
		{"Choumi.", "CCChhhhhhhhooooooooooooooouuuuuuuuuuuuuuuuuuuuummmmmmmmmmmmmiiiiiiiii."},
		{"", ""},
		{"abacadaba 01!", "abbacccaddddabba 01!"},
	}
	for _, tc := range testCases {
		got := RepeatAlpha(tc.in)
		if got != tc.want {
			fmt.Printf("RepeatAlpha(%q) = %q instead of %q\n", tc.in, got, tc.want)
			os.Exit(1)
		}
	}
}
