package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Given a signed 32-bit integer x, return x with its digits reversed.
// If reversing x causes the value to go outside the signed 32-bit integer
// range [-231, 231 - 1], then return 0.

// Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

func reverse(x int) int {

	xStr := strconv.Itoa(x)
	xsp := strings.Split(xStr, "")
	for i := len(xsp)/2 - 1; i >= 0; i-- {
		opp := len(xsp) - 1 - i
		xsp[i], xsp[opp] = xsp[opp], xsp[i]
	}
	xStr = strings.Join(xsp, "")
	for strings.HasSuffix(xStr, "-") {
		xStr = strings.Trim(xStr, "-")
		xStr = "-" + xStr
	}
	res, err := strconv.Atoi(xStr)
	if err != nil {
		panic(err)
	}
	if res > 2147483647 {
		return 0
	} else if res <= -2147483648 {
		return 0
	}
	return res
}

func main() {

	fmt.Println(reverse(-123))

}
