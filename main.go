package main

import (
	"fmt"
	"log"
	"salty/core"
	"salty/data"
	"time"
)

func main() {
	fmt.Println("hello")
	repo, err := data.NewRepo()
	if err != nil {
		log.Fatalf("Cant create repo because of error: %v \n", err)
	}
	srv := core.NewBirthday(repo)

	// resB := srv.ListTodayBirthdays()
	today := time.Now()

	fDate := core.Dob{Year: today.Year(), Month: int(today.Month()), Day: today.Day()}
	// fDate := core.Dob{Year: today.Year(), Month: 13, Day: today.Day()}
	if !fDate.Validate() {
		log.Fatalf("invalid target date supplied: %v", fDate)

	}
	resB := srv.ListBirthdays(fDate)
	for _, v := range resB {
		fmt.Printf("birthday : %#v \n", v)
	}
}
