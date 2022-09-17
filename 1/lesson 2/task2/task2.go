package main

import (
	"fmt"
	"time"
)

func main() {
	dobStr := "20.04.1998" // Replace this date with your birthday
	givenDate, err := time.Parse("02.01.2006", dobStr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s is %s", givenDate.Format("02-01-2006"), FindWeekday(givenDate))
}

func FindWeekday(date time.Time) (weekday string) {
	//
	// WRITE YOUR CODE HERE
	date.Weekday()
	weekday = date.Weekday().String()
	//
	return weekday
}
