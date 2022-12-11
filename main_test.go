package main

import (
	"reflect"
	"salty/core"
	"salty/data"
	"testing"
)

func TestBirthdayServiceListBirthdays(t *testing.T) {
	r, err := data.NewRepo()
	if err != nil {
		t.Error("cant even create repo")

	}
	s := core.NewBirthDay(r)
	res := s.ListBirthdays(core.Dob{})

	if !reflect.DeepEqual(res, []core.Person{}) {
		t.Error("failed")
	}

}
