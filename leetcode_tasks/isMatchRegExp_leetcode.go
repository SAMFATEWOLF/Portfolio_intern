package main

import (
	"fmt"
	"regexp"
)

// Given an input string s and a pattern p, implement regular 
//expression matching with support for '.' and '*' where:

// '.' Matches any single character.​​​​
// '*' Matches zero or more of the preceding element.
// The matching should cover the entire input string (not partial).

func isMatch(s string, p string) bool {
	res, err := regexp.Compile(fmt.Sprintf("^%s$", p))
	if err != nil {
		panic(err)
	}
	resBool := res.Match([]byte(s))
	return resBool

}

func main() {
	s := "aa"
	p := ".*"
	fmt.Println(isMatch(s, p))
}
