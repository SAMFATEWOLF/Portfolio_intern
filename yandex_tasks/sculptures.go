package main

import (
	"fmt"
	"sort"
	"strings"
)


// К Новому-преновому году работники Тындекса построили N ледяных скульптур, 
// i-я скульптура состоит из ai килограмм льда.
// Но они не посоветовались с Кузей! А ведь Кузя знает, что идеальная скульптура 
// состоит из ровно X килограмм льда, не больше и не меньше.
// Новый-преновый год уже совсем скоро, до него осталось всего T минут. За одну 
// минуту Кузя может выбрать одну скульптуру и добавить или удалить ровно 1 килограмм льда из неё.
// Вас, как отличника художественной школы, Кузя просит найти максимальное 
// количество идеальных скульптур в момент наступления праздника.
// Формат ввода
// В первой строке вводятся три целых числа N, X, T(1≤N≤2⋅105;0≤X≤109;0≤T≤3⋅1014) — 
// количество скульптур, идеальное количество льда в скульптуре и оставшееся количество минут до наступления праздника.
// Во второй строке вводятся через пробел N целых чисел ai(1≤ai≤109) — количество 
// килограмм льда в i-й скульптуре.
// Формат вывода
// В первой строке выведите целое число K(0≤K≤N) — максимально возможное количество 
// идеальных скульптур в момент наступления праздника.
// Во второй строке выведите через пробел K различных целых чисел bi(1≤bi≤N) — номера 
// скульптур, которые будут идеальными в момент наступления Нового-пренового года.
// Скульптуры нумеруются с 1 в порядке ввода.
// Если оптимальных ответов несколько, то выведите любой из оптимальных.


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
