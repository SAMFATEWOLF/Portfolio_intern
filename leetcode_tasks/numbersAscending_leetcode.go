package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// A sentence is a list of tokens separated by a single space with no leading or trailing spaces. E
// very token is either a positive number consisting of digits 0-9 with no leading zeros, or a
// word consisting of lowercase English letters.

// For example, "a puppy has 2 eyes 4 legs" is a sentence with seven tokens: "2" and "4"
// are numbers and the other tokens such as "puppy" are words.
// Given a string s representing a sentence, you need to check if all the numbers in s
// are strictly increasing from left to right (i.e., other than the last number, each
// 	number is strictly smaller than the number on its right in s).

// Return true if so, or false otherwise.

func areNumbersAscending(s string) bool {
	var arr []int
	split := strings.Split(s, " ")
	for _, el := range split {
		if unicode.IsDigit(rune(el[0])) {
			elInt, err := strconv.Atoi(el)
			if err != nil {
				panic(err)
			}
			arr = append(arr, elInt)
		}
	}
	counter := 0
	flag := 0
	for _, el := range arr {
		if flag < el {
			flag = el
			counter++
		}
	}
	if counter == len(arr) {
		return true
	}
	return false
}

func main() {

	s := "there is 2 eyes 2 legs 123 asd"
	fmt.Println(areNumbersAscending(s))
}
