package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

var (
	errInvalidDateFormat = errors.New("invalid date format")
	errInvalidDate = errors.New("invalid date")
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go mm/dd/yyyy")
		return
	}

	argDate := os.Args[1]

	birthDate, err := convArgToDate(argDate)
	if err != nil {
		fmt.Println("Usage: go run main.go mm/dd/yyyy")
		return
	}

	timeAlive := time.Since(birthDate)
	if timeAlive.Seconds() < 0 {
		fmt.Println("Please enter a valid date that is not in the future.")
		return
	}

	yearsAlive := timeAlive.Hours() / 24 / 365
	daysAlive := timeAlive.Hours() / 24

	fmt.Printf("Years alive: %f\nDays alive: %f\nMinutes alive: %f\nSeconds alive: %f\n", yearsAlive, daysAlive, timeAlive.Minutes(), timeAlive.Seconds())

}

func convArgToDate(arg string) (time.Time, error) {
	// TODO: Need to fix it so it correctly handles months with less than 31 days

	birthMonthInt, err := strconv.Atoi(arg[:2])
	if err != nil {
		return time.Time{}, errInvalidDateFormat
	}
	if birthMonthInt <= 0 || birthMonthInt > 12 {
		return time.Time{}, errInvalidDate
	}
	var birthMonth time.Month = time.Month(birthMonthInt)

	birthDay, err := strconv.Atoi(arg[3:5])
	if err != nil {
		return time.Time{}, errInvalidDateFormat
	}
	if birthDay <= 0 || birthDay > 31 {
		return time.Time{}, errInvalidDate
	}

	birthYear, err := strconv.Atoi(arg[6:])
	if err != nil {
		return time.Time{}, errInvalidDateFormat
	}

	return time.Date(birthYear, birthMonth, birthDay, 0, 0, 0, 0, time.Local), nil
}
