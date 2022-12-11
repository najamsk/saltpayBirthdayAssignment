package core

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
