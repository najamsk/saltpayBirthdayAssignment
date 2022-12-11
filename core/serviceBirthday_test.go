package core

import (
	"reflect"
	"testing"
)

func TestBirthdayService_ListBirthdays_EmptyResponse_WithMockRepo(t *testing.T) {
	mr := NewMockRepo()
	s := NewBirthDay(mr)
	res := s.ListBirthdays(Dob{})

	if !reflect.DeepEqual(res, []Person{}) {
		t.Error("failed")
	}

}
