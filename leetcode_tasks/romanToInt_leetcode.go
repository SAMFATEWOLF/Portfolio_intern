package main

import (
	"fmt"
	"strings"
)

// // Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII.
// Instead, the number four is written as IV. Because the one is before the five we subtract it making four.
// The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

// // I can be placed before V (5) and X (10) to make 4 and 9.
// // X can be placed before L (50) and C (100) to make 40 and 90.
// // C can be placed before D (500) and M (1000) to make 400 and 900.
// // Given a roman numeral, convert it to an integer.
func romanToInt(s string) int {
	m := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,

		"F": 4,
		"N": 9,
		"R": 40,
		"T": 90,
		"Y": 400,
		"U": 900,
	}
	var q int

	if strings.Contains(s, "IV") {
		s = strings.ReplaceAll(s, "IV", "F")
	}
	if strings.Contains(s, "IX") {
		s = strings.ReplaceAll(s, "IX", "N")
	}
	if strings.Contains(s, "XL") {
		s = strings.ReplaceAll(s, "XL", "R")
	}
	if strings.Contains(s, "XC") {
		s = strings.ReplaceAll(s, "XC", "T")
	}
	if strings.Contains(s, "CD") {
		s = strings.ReplaceAll(s, "CD", "Y")
	}
	if strings.Contains(s, "CM") {
		s = strings.ReplaceAll(s, "CM", "U")
	}

	for _, el := range s {
		elStr := string(el)
		q += m[elStr]
	}
	return q
}
func main() {
	s := []string{"X", "XXVI", "MCMXCVII"}
	for _, el := range s {
		w := romanToInt(el)
		fmt.Println(w)
	}
}
