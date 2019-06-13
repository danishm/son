package person

import (
	"time"
)

// Person represents a basic person and their attributes in the world
type Person struct {
	Name string    `json:"name"`
	DOB  time.Time `json:"dob"`
}

// Age returns the age of a person
func (p *Person) Age() Age {
	return AgeFromTime(p.DOB, time.Now())
}
