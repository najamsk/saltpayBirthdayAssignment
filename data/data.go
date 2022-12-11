package data

import "salty/core"

type Repo struct {
}

func NewRepo() (*Repo, error) {
	return &Repo{}, nil
}

func (r *Repo) GetAll() []core.Person {
	return []core.Person{}
}
