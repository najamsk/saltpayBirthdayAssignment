package data

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"salty/core"
)

type Repo struct {
	People []core.Person
}

func NewRepo() (*Repo, error) {
	p, err := setupData()
	if err != nil {
		return nil, err
	}

	return &Repo{People: p}, nil
}

func (r *Repo) GetAll() []core.Person {
	return r.People
}
func setupData() ([]core.Person, error) {
	filepath := parseFileFlag()

	people := []core.Person{}
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		// log.Printf("Erorr while reading from %s: %v \n", filepath, err)
		return people, err
	}

	type BirthDaysData [][]string
	var lines BirthDaysData
	// unmarshall it
	err = json.Unmarshal(data, &lines)
	if err != nil {
		// log.Printf("Erorr while parsing data form %s: %v \n", filepath, err)
		return people, err
	}

	// transform into our data.Person object
	for _, v := range lines {
		// fmt.Printf("record: %#v \n", v)
		p := core.Person{FirstName: v[1], LastName: v[0], Birthday: v[2]}
		people = append(people, p)
	}
	return people, nil
}

func parseFileFlag() string {
	iFile := flag.String("file", "data.json", "provide file to read")
	flag.Parse()
	fmt.Printf("will read file: %s \n", *iFile)
	return *iFile
}
