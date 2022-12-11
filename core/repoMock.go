package core

type MockRepo struct {
	People []Person
}

func NewMockRepo() *MockRepo {
	return &MockRepo{
		People: []Person{},
	}
}

func (r *MockRepo) GetAll() []Person {
	return r.People
}
