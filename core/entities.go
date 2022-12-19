package core

import (
	"fmt"
	"time"
)

type Person struct {
	FirstName    string
	LastName     string
	Birthday     string
	BirthdayDate Dob
}

type Dob struct {
	Day    int
	Month  int
	Year   int
	IsLeap bool
}

func (d Dob) Validate() bool {
	str := fmt.Sprintf("%d-%d-%d", d.Year, d.Month, d.Day)
	_, err := time.Parse("2006-01-02", str)
	return err == nil
}
