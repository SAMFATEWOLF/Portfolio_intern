package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

// Implement the myAtoi(string s) function, which converts a string to a 32-bit signed integer (similar to C/C++'s atoi function).

// // The algorithm for myAtoi(string s) is as follows:

// // Read in and ignore any leading whitespace.
// // Check if the next character (if not already at the end of the string) is '-' or '+'.
// // Read this character in if it is either. This determines if the final result is negative
// // or positive respectively. Assume the result is positive if neither is present.
// // Read in next the characters until the next non-digit character or the end of the input
// // is reached. The rest of the string is ignored.
// // Convert these digits into an integer (i.e. "123" -> 123, "0032" -> 32). If no digits
// // were read, then the integer is 0. Change the sign as necessary (from step 2).
// // If the integer is out of the 32-bit signed integer range [-231, 231 - 1], then clamp the
// // integer so that it remains in the range. Specifically, integers less than -231 should be
// // clamped to -231, and integers greater than 231 - 1 should be clamped to 231 - 1.
// // Return the integer as the final result.

func myAtoi(s string) int {

	res := 0

	for strings.HasPrefix(s, " ") {
		s = strings.Trim(s, " ")
	}
	symbols := []string{"+", "-", ".", ",", ""}
	for _, el := range symbols {
		if s == el {
			return 0
		}
	}

	sign := 0
	switch s[0] {
	case '+':
		sign = 1
		s = s[1:]
	case '-':
		sign = -1
		s = s[1:]
	case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
		sign = 1
	default:
		return 0
	}
	if !unicode.IsDigit(rune(s[0])) {
		return 0
	}

	for idx, el := range s {
		if !unicode.IsDigit(el) {
			s = s[:idx]
			break
		}
	}
	if len(s) >= 11 {
		switch s[0] {
		case '0':
			for strings.HasPrefix(s, "0") {
				s = strings.Trim(s, "0")
			}
		default:
			if sign == -1 {
				return math.MinInt32
			} else if sign == 1 {
				return math.MaxInt32
			}
		}
	}

	res, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}

	res *= sign
	if res > math.MaxInt32 {
		return math.MaxInt32
	}

	if res < math.MinInt32 {
		return math.MinInt32
	}
	return res

}

func main() {
	s := "-000000000000000322333333"
	fmt.Println(myAtoi(s))

}
