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
