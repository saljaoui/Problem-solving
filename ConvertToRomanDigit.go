package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args[1:]) != 1 {
		fmt.Println("not correct")
		return
	}
	text := os.Args[1]
	if !chekcIfNumber(text) {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}
	nb, _ := strconv.Atoi(text)
	res := ConvertToRomanDigit(nb)
	fmt.Println(res)
}

func ConvertToRomanDigit(nb int) string {
	var str string
	var text string

	for i := 1; i <= nb; i++ {
		if nb >= 1 && nb < 5 { // 1
			if nb < 4 {
				text += "I"
				text += "+"
				str += "I"
			}
			if i == 4 {
				str += "IV"
				text += "(V-I)"
				text += "+"
			}
		} else if nb >= 5 && nb < 10 { // 2
			if i == 9 {
				str += "IX"
				text += "(X-I)"
				text += "+"
			}
			if nb < 9 {
				text += "V"
				text += "+"
				str += "V"
				i = 0
				nb -= 5
			}
		} else if nb >= 10 && nb < 50 { // 3
			if nb < 40 {
				text += "X"
				text += "+"
				str += "X"
				i = 0
				nb -= 10

			}
			if i == 40 {
				str += "XL"
				text += "(L-X)"
				text += "+"
				nb = nb - 40
				i = 0
			}
		} else if nb >= 50 && nb < 100 {
			if nb < 90 {
				str += "L"
				text += "L"
				text += "+"

				i = 0
				nb -= 50

			}
			if i == 90 {
				str += "XC"
				text += "(C-X)"
				text += "+"
				i = 0
				nb = nb - 90
			}
		} else if nb >= 100 && nb < 500 {
			if nb < 400 {
				str += "C"
				text += "C"
				text += "+"
				i = 0
				nb -= 100

			}
			if i == 400 {
				str += "CD"
				text += "(D-C)"
				text += "+"
				i = 0
				nb = nb - 400
			}
		} else if nb >= 500 && nb < 1000 {
			if nb < 900 {
				str += "D"
				text += "D"
				text += "+"
				i = 0
				nb -= 500

			}
			if i == 900 {
				str += "CM"
				text += "(M-C)"
				text += "+"
				i = 0
				nb = nb - 900
			}
		} else if nb >= 1000 && nb < 4000 {
			str += "M"
			text += "M"
			text += "+"
			nb -= 1000
		}
	}
	newtext := text[:len(text)-1]
	fmt.Println(newtext)
	return str
}

func chekcIfNumber(nb string) bool {
	for _, s := range nb {
		if s < '0' || s > '9' {
			return false
		}
	}
	return true
}
