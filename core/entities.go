package core

type Person struct {
	FirstName string
	LastName  string
	Birthday  string
}

type Dob struct {
	Day    int
	Month  int
	Year   int
	IsLeap bool
}
