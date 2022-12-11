package core

import "time"

type Birthday struct {
	repo Repo
}

type Repo interface {
	GetAll() []Person
}

func NewBirthDay(repo Repo) *Birthday {
	return &Birthday{repo: repo}
}

func (s *Birthday) ListBirthdays(forDate Dob) []Person {
	return []Person{}
}

func IsLeapYear(y int) bool {
	year := time.Date(y, time.December, 31, 0, 0, 0, 0, time.Local)
	days := year.YearDay()

	if days > 365 {
		return true
	} else {
		return false
	}
}
