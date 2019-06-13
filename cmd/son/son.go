package main

import (
	"fmt"
	"time"

	"github.com/danishm/son/pkg/person"
)

func main() {
	newYork, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println("Unable to load timezone")
	}
	son := person.Person{
		DOB: time.Date(2016, time.March, 30, 1, 15, 0, 0, newYork),
	}
	fmt.Printf("%+v\n", son.Age())
}
