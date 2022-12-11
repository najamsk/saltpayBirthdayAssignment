package core

type Birthday struct {
}

type Repo interface {
	GetAll() []Person
}

func NewBirthDay(repo Repo) *Birthday {
	return &Birthday{}
}

func (s *Birthday) ListBirthdays(forDate Dob) []Person {
	return []Person{}
}
