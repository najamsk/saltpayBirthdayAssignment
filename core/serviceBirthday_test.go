package core

import (
	"reflect"
	"testing"
)

func TestBirthdayService_ListBirthdays_EmptyResponse_WithMockRepo(t *testing.T) {
	mr := NewMockRepo()
	s := NewBirthDay(mr)
	iDate := Dob{Year: 2020, Month: 02, Day: 28, IsLeap: IsLeapYear(2020)}
	want := []Person{{FirstName: "Salma", LastName: "Riz", Birthday: "2020/02/28"}}

	got := s.ListBirthdays(iDate)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got:%v, want:%v \n", got, want)
	}

}
