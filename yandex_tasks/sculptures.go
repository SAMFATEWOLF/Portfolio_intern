package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var a, b int
	var firstRaw []int
	for i := 0; i < 3; i++ {
		fmt.Scan(&a)
		firstRaw = append(firstRaw, a)
	}
	var secondRaw []int
	for i := 0; i < firstRaw[0]; i++ {
		fmt.Scan(&b)
		secondRaw = append(secondRaw, b)
	}
	m := make(map[int]int)
	var time []int
	for _, el := range secondRaw {
		q := el - firstRaw[1]
		if q < firstRaw[2] {
			if q < 0 {
				q *= -1
			}
			m[q] = el
			time = append(time, q)
		}
	}
	sort.Ints(time)
	fmt.Println(m)
	fmt.Println(time)
	var w int
	var resArr1 []int
	for _, el := range time {
		w += el
		resArr1 = append(resArr1, el)
		if w > firstRaw[2] {
			break
		}
	}
	resArr := resArr1[:]
	for _, el := range resArr1 {
		e := 0
		e += el
		if e > firstRaw[2] {
			resArr = resArr1[:len(resArr1)-1]
		}
	}

	m2 := make(map[int]int)
	for idx, el := range m {
		for _, el1 := range resArr {
			if idx == el1 {
				m2[idx] = el
			}
		}
	}
	var ansArr []int
	for idx, el := range secondRaw {
		for _, el1 := range m2 {
			if el == el1 {
				ansArr = append(ansArr, idx+1)
			}
		}
	}
	answer := fmt.Sprintln(ansArr)
	answer = strings.Trim(answer, "[")
	answer = strings.ReplaceAll(answer, "]", "")

	fmt.Println(len(resArr))
	fmt.Println(answer)
}
