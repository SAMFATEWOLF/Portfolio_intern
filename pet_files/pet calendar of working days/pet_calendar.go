package main

import (
	"fmt"
	"time"
)

// Программа, высчитывающая количество рабочих дней между выбранными датами с учетом выходных и праздничных дней на 2023 год.

func isWeekend(date time.Time) bool {
	return date.Weekday() == time.Saturday || date.Weekday() == time.Sunday
}
func isHolliday(date time.Time) bool {
	hollidayArr := []string{
		"01-01-2023",
		"02-01-2023",
		"03-01-2023",
		"04-01-2023",
		"05-01-2023",
		"06-01-2023",
		"07-01-2023",
		"08-01-2023",
		"23-02-2023",
		"08-03-2023",
		"01-05-2023",
		"09-05-2023",
		"12-06-2023",
		"06-11-2023",
	}
	formatDate := date.Format("02-01-2006")
	for _, el := range hollidayArr {
		if el == formatDate {
			return true
		}
	}
	return false
}

func main() {
	var startDate string
	fmt.Scanln(&startDate)
	var endDate string
	fmt.Scanln(&endDate)

	startParse, err := time.Parse("02-01-2006", startDate)
	if err != nil {
		panic(err)
	}
	endParse, err := time.Parse("02-01-2006", endDate)
	if err != nil {
		panic(err)
	}

	daysAmmount := endParse.Sub(startParse).Hours() / 24
	for i := startParse; i != endParse; i = i.AddDate(0, 0, 1) {
		if isHolliday(i) || isWeekend(i) {
			daysAmmount--
		}
	}
	resStr := fmt.Sprint(daysAmmount)

	switch resStr[len(resStr)-1:] {
	case "1":
		fmt.Printf("Между выбранными датами %v рабочий день", daysAmmount)
	case "2", "3", "4":
		fmt.Printf("Между выбранными датами %v рабочих дня", daysAmmount)
	default:
		fmt.Printf("Между выбранными датами %v рабочих дней", daysAmmount)
	}
}
