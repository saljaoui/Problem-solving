package main

import (
	"log"
	"reflect"
)

func RevConcatAlternate(slice1, slice2 []int) []int {
	slc1 := len(slice1)
	slc2 := len(slice2)
	
	res := []int{}
	if slc1 == slc2 {
		for i := slc1 - 1; i >= 0; i-- {
			res = append(res, slice1[i])
			res = append(res, slice2[i])

		}
	}
	if slc1 < slc2 {
		for i := slc2 - 1; i > slc1-1; i-- {
			res = append(res, slice2[i])
		}
		for i := slc1 - 1; i >= 0; i-- {
			res = append(res, slice1[i])
			res = append(res, slice2[i])
		}
	}
	if slc1 > slc2 {
		for i := slc1 - 1; i > slc2-1; i-- {
			res = append(res, slice1[i])
		}
		for i := slc2 - 1; i >= 0; i-- {
			res = append(res, slice1[i])
			res = append(res, slice2[i])

		}
	}


	return res
}

var testCases = []struct {
	args [][]int
	want []int
}{
	{
		[][]int{{1, 2, 3}, {4, 5, 6}},
		[]int{3, 6, 2, 5, 1, 4},
	},
	{
		[][]int{{4, 5, 6, 7, 8, 9}, {1, 2, 3}},
		[]int{9, 8, 7, 6, 3, 5, 2, 4, 1},
	},
	{
		[][]int{{1, 2, 3}, {4, 5, 6, 7, 8, 9}},
		[]int{9, 8, 7, 3, 6, 2, 5, 1, 4},
	},
	{
		[][]int{{1, 2, 3}, {4, 5}},
		[]int{3, 2, 5, 1, 4},
	},
	{
		[][]int{{}, {4, 5, 6}},
		[]int{6, 5, 4},
	},
	{
		[][]int{{1, 2, 3}, {}},
		[]int{3, 2, 1},
	},
	{
		[][]int{{}, {}},

		[]int{},
	},
	{
		[][]int{{1, 2, 4}, {10, 20, 30, 40, 50}},

		[]int{50, 40, 4, 30, 2, 20, 1, 10},
	},
}

func main() {
	for _, tc := range testCases {
		got := RevConcatAlternate(tc.args[0], tc.args[1])
		if !reflect.DeepEqual(got, tc.want) {
			log.Fatalf("%s(%#v, %#v) == %#v instead of %#v\n",
				"RevConcatAlternate",
				tc.args[0],
				tc.args[1],
				got,
				tc.want,
			)
		}
	}
}
