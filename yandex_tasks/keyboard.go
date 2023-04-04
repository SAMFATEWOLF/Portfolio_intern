package main

import (
	"fmt"
)

func main() {

	var a, b, c, d, e int
	fmt.Scan(&a)

	var arr2, arr3 []int
	for i := 0; i < a; i++ {
		fmt.Scan(&b)
		arr2 = append(arr2, b)
	}
	for i := 0; i < a; i++ {
		fmt.Scan(&c)
		arr3 = append(arr3, c)
	}
	fmt.Scan(&d)
	var arr5 []int
	for i := 0; i < d; i++ {
		fmt.Scan(&e)
		arr5 = append(arr5, e)
	}

	counter := 0
	flag := 0
	m := make(map[int]int)
	for idx, el := range arr2 {
		idx = arr3[idx]
		m[el] = idx
	}
	for _, el := range arr5 {
		if flag != m[el] {
			flag = m[el]
			counter++
		}
	}
	res := counter - 1
	if res < 0 {
		res = 0
	}

	fmt.Println(res)
}
