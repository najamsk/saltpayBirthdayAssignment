package core

import (
	"fmt"
	"time"
)

type constError string

func (err constError) Error() string {
	return string(err)
}

const (
	ErrParsingDate = constError("Cant parse the birthday string")
)

type Birthday struct {
	repo Repo
}

type Repo interface {
	GetAll() []Person
}

func NewBirthday(repo Repo) *Birthday {
	return &Birthday{repo: repo}
}

func (s *Birthday) ListBirthdays(forDate Dob) []Person {

	fmt.Println("service.ListBirthdays invoked")
	party := []Person{}
	// tDob := todayDobFormat()
	p := s.repo.GetAll()
	// fmt.Printf("got peole from repo:%v \n", len(p))

	for _, v := range p {
		chk := isYourBirthdayForDate(v, forDate)
		if chk {
			party = append(party, v)
		}
		// fmt.Printf("name: %s, birthday:%#v \n", v.FirstName, birthdayDobFormat(v.Birthday))
	}
	return party

}

func isYourBirthdayForDate(p Person, forDate Dob) bool {
	//if person year is leap and having feb as a month and date 29
	//	then check if today year is also leap
	//	else normal checking of month and date
	tDob := forDate

	// fmt.Printf("checking:%v, dob: %#v \n", p, pDob)

	if p.BirthdayDate.IsLeap && p.BirthdayDate.Month == 2 && p.BirthdayDate.Day == 29 {
		//match with today is leapYear
		if tDob.Day == 29 && tDob.Month == 2 && tDob.IsLeap {
			//if tody is 29 feb due to leap year
			return true
		} else if tDob.Day == 28 && tDob.Month == 2 && !tDob.IsLeap {
			//if tody is 28 feb due not leap year
			return true
		} else {
			return false
		}
	} else if p.BirthdayDate.Month == tDob.Month && p.BirthdayDate.Day == tDob.Day {
		return true
	}

	return false
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
