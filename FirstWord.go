package main

import (
	"fmt"
	"os"
)

func FirstWord(s string) string {

	if s == "" {
		return "\n"
	}
	var str string
	for _, ss := range s {
		if ss == ' ' && str != "" {
			return str + "\n"
		} else if ss != ' ' {
			str += string(ss)
		}

	}

	return s
}

var testCases = []struct {
	in   string
	want string
}{
	{"", "\n"},
	{"             a as", "a\n"},
	{"   f     d", "f\n"},
	{"   asd    ad", "asd\n"},
	{"   salut !!! ", "salut\n"},
	{" salut ! ! !", "salut\n"},
	{"salut ! !", "salut\n"},
}

func main() {
	for _, tc := range testCases {
		got := FirstWord(tc.in)
		if got != tc.want {
			fmt.Printf("FirstWord(%q) = %q instead of %q\n", tc.in, got, tc.want)
			os.Exit(1)
		}
	}
}
