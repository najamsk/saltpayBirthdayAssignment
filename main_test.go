package main

import "testing"

func TestBirthdayServiceListBirthdays(t *testing.T) {
	r, err := NewRepo()
	if err != nil {
		t.Error("cant even create repo")

	}
	s := NewBirthDay(r)
	res := s.ListBirthdays(targetDate)

}
