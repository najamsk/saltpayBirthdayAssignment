package data

import (
	"encoding/json"
	"flag"
	"os"
	"salty/core"
	"strconv"
	"strings"
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

func birthdayDobFormat(d string) (core.Dob, error) {
	dd := strings.Split(d, "/")
	yy, err := strconv.Atoi(dd[0])
	if err != nil {
		return core.Dob{}, core.ErrParsingDate
	}
	days, err := strconv.Atoi(dd[2])
	if err != nil {
		return core.Dob{}, core.ErrParsingDate
	}
	mm, err := strconv.Atoi(dd[1])
	if err != nil {
		return core.Dob{}, core.ErrParsingDate
	}

	dob := core.Dob{Year: yy, Month: mm, Day: days, IsLeap: core.IsLeapYear(yy)}
	return dob, nil
}
func (r *Repo) GetAll() []core.Person {
	res := []core.Person{}
	for _, v := range r.People {

		bd, err := birthdayDobFormat(v.Birthday)

		if err != nil || !bd.Validate() {
			continue
		}
		v.BirthdayDate = bd
		res = append(res, v)

	}
	return res
}
func setupData() ([]core.Person, error) {
	filepath := parseFileFlag()

	people := []core.Person{}
	// data, err := ioutil.ReadFile(filepath)
	data, err := os.ReadFile(filepath)
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
		//XXX: parse birthday to Dob object here?
		p := core.Person{FirstName: v[1], LastName: v[0], Birthday: v[2]}
		people = append(people, p)
	}
	return people, nil
}

func parseFileFlag() string {
	iFile := flag.String("file", "data.json", "provide file to read")
	flag.Parse()
	// fmt.Printf("will read file: %s \n", *iFile)
	return *iFile
}
