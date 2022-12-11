package core

type MockRepo struct {
	People []Person
}

func NewMockRepo() *MockRepo {
	return &MockRepo{
		People: []Person{
			// {FirstName: "John", LastName: "Doe", Birthday: "1982/10/08"},
			// {FirstName: "Bruce", LastName: "Wayne", Birthday: "1965/01/30"},
			// {FirstName: "Lady", LastName: "Gaga", Birthday: "1986/03/28"},
			{FirstName: "Mark", LastName: "Curry", Birthday: "1988/02/29"},
			// {FirstName: "Mark", LastName: "Cuban", Birthday: "2021/02/28"},
			{FirstName: "Silly", LastName: "Goose", Birthday: "1996/02/29"},
			{FirstName: "Salma", LastName: "Riz", Birthday: "2020/02/28"},
			{FirstName: "Khali", LastName: "Great", Birthday: ""},
			// {FirstName: "Eddie", LastName: "Tubbs", Birthday: "2021/12/11"},
		},
	}
}

func (r *MockRepo) GetAll() []Person {
	return r.People
}
