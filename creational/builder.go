package creational

import (
	"fmt"
	"time"
)

type PersonProfile struct {
	Name     string
	Birthday time.Time
	Phone    string
	Email    string
}

type Person struct {
	Detail  PersonProfile
	Job     string
	Hobbies []string
}

func NewPerson(detail PersonProfile, job string, hobbies []string) Person {
	return Person{
		Detail:  detail,
		Job:     job,
		Hobbies: hobbies,
	}
}

// builder, split the construction and the object
// so the same object type can has different value
func CreatePerson() {
	richard := NewPerson(PersonProfile{
		Name:     "Richard",
		Birthday: time.Now(),
		Phone:    "11111111",
		Email:    "richard@gmail.com",
	}, "Backend", []string{"playing", "reading"})
	fmt.Println(richard)

	kara := NewPerson(PersonProfile{
		Name:     "Kara",
		Birthday: time.Now(),
		Phone:    "222222222",
		Email:    "kara@gmail.com",
	}, "QA", []string{"teasing", "eating"})
	fmt.Println(kara)
}
