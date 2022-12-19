package core

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBirthdayService_ListBirthdays_EmptyResponse_WithMockRepo(t *testing.T) {
	mr := NewMockRepo()
	s := NewBirthday(mr)
	iDate := Dob{Year: 2020, Month: 02, Day: 28, IsLeap: IsLeapYear(2020)}
	want := []Person{{FirstName: "Salma", LastName: "Riz", Birthday: "2020/02/28"}}

	got := s.ListBirthdays(iDate)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got:%v, want:%v \n", got, want)
	}

}

func TestListBirthDays(t *testing.T) {
	//arrange
	r := NewMockRepo()
	s := NewBirthday(r)

	type TestCase struct {
		description string
		input       Dob
		want        []Person
	}

	PeopleEmpty := []Person{}

	People28FebWithNormalYear := []Person{
		{FirstName: "Mark", LastName: "Curry", Birthday: "1988/02/29"},
		{FirstName: "Silly", LastName: "Goose", Birthday: "1996/02/29"},
		{FirstName: "Salma", LastName: "Riz", Birthday: "2020/02/28"},
	}
	People29Feb := []Person{
		{FirstName: "Mark", LastName: "Curry", Birthday: "1988/02/29"},
		{FirstName: "Silly", LastName: "Goose", Birthday: "1996/02/29"},
	}
	People28FebWithLeapYear := []Person{
		{FirstName: "Salma", LastName: "Riz", Birthday: "2020/02/28"},
	}
	// NonLeapYear29Feb := []data.Person{}

	tcs := []TestCase{
		{"29Feb With Leap Year", Dob{Year: 2024, Month: 02, Day: 29, IsLeap: IsLeapYear(2024)}, People29Feb},
		{"28Feb With Normal Year", Dob{Year: 2022, Month: 02, Day: 28, IsLeap: IsLeapYear(2022)}, People28FebWithNormalYear},
		{"28Feb With Leap Year", Dob{Year: 2020, Month: 02, Day: 28, IsLeap: IsLeapYear(2020)}, People28FebWithLeapYear},
		{"Skipping invalid birthday", Dob{Year: 1982, Month: 02, Day: 22, IsLeap: IsLeapYear(1982)}, PeopleEmpty},
	}

	//act
	for _, tc := range tcs {
		t.Run(tc.description, func(t *testing.T) {
			got := s.ListBirthdays(tc.input)
			fmt.Printf("Date:%v:, Got: %v \n", tc.input, got)
			//assert
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("failed with date:%v, want:%v, got:%v \n", tc.input, tc.want, got)
			}
		})
	}
}

func TestLeapYearTable(t *testing.T) {
	tcs := []struct {
		description string
		input       int
		want        bool
	}{
		{"2024 is leap year", 2024, true},
		{"2000 is leap year", 2000, true},
		{"2004 is leap year", 2004, true},
		{"1996 is leap year", 1996, true},
		{"2022 is not leap year", 2022, false},
		{"1995 is not leap year", 1995, false},
		{"1997 is not leap year", 1997, false},
		{"1998 is not leap year", 1998, false},
	}

	for _, tc := range tcs {
		t.Run(tc.description, func(t *testing.T) {
			got := IsLeapYear(tc.input)
			if got != tc.want {
				t.Errorf("failed with year:%v, want:%v, got:%v \n", tc.input, tc.want, got)
			}
		})

	}
}

// func TestBirthayStringToDobFormat(t *testing.T) {
// 	tcs := []struct {
// 		description string
// 		input       string
// 		want        error
// 	}{
// 		{"empty dob string", "", ErrParsingDate},
// 		{"jibberish string", "adfasdf", ErrParsingDate},
// 	}

// 	for _, tc := range tcs {
// 		t.Run(tc.description, func(t *testing.T) {
// 			_, got := birthdayDobFormat(tc.input)
// 			if got != tc.want {
// 				t.Errorf("input:%v, want:%v, got:%v \n", tc.input, tc.want, got)
// 			}
// 		})
// 	}

// 	tcsP := []struct {
// 		description string
// 		input       string
// 		want        Dob
// 	}{
// 		{"valid date as string", "2022/12/11", Dob{Year: 2022, Day: 11, Month: 12}},
// 	}

// 	for _, tc := range tcsP {
// 		t.Run(tc.description, func(t *testing.T) {
// 			got, _ := birthdayDobFormat(tc.input)
// 			if got != tc.want {
// 				t.Errorf("input:%v, want:%v, got:%v \n", tc.input, tc.want, got)
// 			}
// 		})
// 	}

// }
