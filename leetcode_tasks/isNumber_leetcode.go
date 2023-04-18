package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// A valid number can be split up into these components (in order):

// A decimal number or an integer.
// (Optional) An 'e' or 'E', followed by an integer.
// A decimal number can be split up into these components (in order):

// (Optional) A sign character (either '+' or '-').
// One of the following formats:
// One or more digits, followed by a dot '.'.
// One or more digits, followed by a dot '.', followed by one or more digits.
// A dot '.', followed by one or more digits.
// An integer can be split up into these components (in order):

// (Optional) A sign character (either '+' or '-').
// One or more digits.

// Given a string s, return true if s is a valid number.

func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	if unicode.IsLetter(rune(s[0])) {
		return false
	}
	if len(s) == 1 {
		switch s {
		case "+", "-", ".":
			return false
		}
	}
	switch s[0] {
	case '+', '-':
		if unicode.IsLetter(rune(s[1])) {
			return false
		}
	}
	_, err := strconv.ParseFloat(s, 64)
	return err == nil || strings.Contains(err.Error(), "value out of range")
}

func main() {
	s := "-8115e957"
	fmt.Println(isNumber(s))

}
