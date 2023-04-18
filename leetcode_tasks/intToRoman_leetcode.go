package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Roman numerals are usually written largest to smallest from left to right.
// However, the numeral for four is not IIII. Instead, the number four is written
// as IV. Because the one is before the five we subtract it making four. The same
// principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

// I can be placed before V (5) and X (10) to make 4 and 9.
// X can be placed before L (50) and C (100) to make 40 and 90.
// C can be placed before D (500) and M (1000) to make 400 and 900.
// Given an integer, convert it to a roman numeral.

func intToRoman(num int) string {
	m := map[int]string{
		1:    "I",
		5:    "V",
		10:   "X",
		50:   "L",
		100:  "C",
		500:  "D",
		1000: "M",

		4:   "IV",
		9:   "IX",
		40:  "XL",
		90:  "XC",
		400: "CD",
		900: "CM",
	}

	s := ""

	var array []int
	qStr := strconv.Itoa(num)
	split := strings.Split(qStr, "")
	for _, el := range split {
		elInt, err := strconv.Atoi(el)
		if err != nil {
			panic(err)
		}
		array = append(array, elInt)
	}

	switch len(array) {
	case 1:
		array[0] *= 1

	case 2:
		array[0] *= 10

	case 3:
		array[0] *= 100
		array[1] *= 10

	case 4:
		array[0] *= 1000
		array[1] *= 100
		array[2] *= 10

	}

	for idx, el := range array {
		idxStr := strconv.Itoa(idx)

		if el == 0 {
			idxStr = ""
		}
		switch {
		case el >= 1 && el <= 9:
			if el < 4 {
				idxStr = strings.Repeat(m[1], el)
			} else if el > 5 && el < 9 {
				idxStr = m[5] + strings.Repeat(m[1], el-5)
			}
		case el >= 10 && el <= 99:
			if el < 40 {
				idxStr = strings.Repeat(m[10], el/10)
			} else if el > 50 && el < 90 {
				idxStr = m[50] + strings.Repeat(m[10], (el-50)/10)
			}
		case el >= 100 && el <= 999:
			if el < 400 {
				idxStr = strings.Repeat(m[100], el/100)
			} else if el > 500 && el < 900 {
				idxStr = m[500] + strings.Repeat(m[100], (el-500)/100)
			}
		case el >= 1000 && el < 4000:
			idxStr = strings.Repeat(m[1000], el/1000)

		}

		if m[el] != "" {
			idxStr = m[el]
		}
		s += idxStr
	}
	return s
}

func main() {
	q := []int{4, 25, 92, 485, 1234, 2873}

	for _, el := range q {
		w := intToRoman(el)
		fmt.Println(w)
	}

}
