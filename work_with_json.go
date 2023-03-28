// На стандартный ввод подаются данные о студентах университетской группы в формате JSON:

// {
//     "ID":134,
//     "Number":"ИЛМ-1274",
//     "Year":2,
//     "Students":[
//         {
//             "LastName":"Вещий",
//             "FirstName":"Лифон",
//             "MiddleName":"Вениаминович",
//             "Birthday":"4апреля1970года",
//             "Address":"632432,г.Тобольск,ул.Киевская,дом6,квартира23",
//             "Phone":"+7(948)709-47-24",
//             "Rating":[1,2,3]
//         },
//         {
//             // ...
//         }
//     ]
// }
// В сведениях о каждом студенте содержится информация о полученных им оценках (Rating). 
// Требуется прочитать данные, и рассчитать среднее количество оценок, полученное студентами группы. 
// Ответ на задачу требуется записать на стандартный вывод в формате JSON в следующей форме:

// {
//     "Average": 14.1
// }

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

type Group struct {
	ID       int
	Number   string
	Year     int
	Students []Student
}
type Student struct {
	LastName   string
	FirstName  string
	MiddleName string
	Birthday   string
	Adress     string
	Phone      string
	Rating     []int
}
type Rate struct {
	Average float32
}

func main() {

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil && err != io.EOF {
		panic(err)
	}

	var Test Group
	if err := json.Unmarshal(data, &Test); err != nil {
		panic(err)
	}
	counter := 0
	gradeAmmount := float32(0)
	for _, elStudent := range Test.Students {
		if elStudent.LastName != "" {
			counter++
		}
		gradeAmmount += float32(len(elStudent.Rating))
	}

	var R2 = Rate{}

	q := fmt.Sprintf("%.1f", gradeAmmount/float32(counter))
	fl, err := strconv.ParseFloat(q, 32)
	if err != nil {
		panic(err)
	}
	R2.Average = float32(fl)

	res, err := json.MarshalIndent(R2, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(res))

}
