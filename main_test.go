package main

import (
	"reflect"
	"salty/core"
	"salty/data"
	"testing"
)

func TestBirthdayService_WithRealRepoListBirthdays(t *testing.T) {
	r, err := data.NewRepo()
	if err != nil {
		t.Error("cant even create repo")

	}
	s := core.NewBirthday(r)
	res := s.ListBirthdays(core.Dob{})

	if !reflect.DeepEqual(res, []core.Person{}) {
		t.Error("failed")
	}

}
