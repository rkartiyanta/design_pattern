package creational

import "fmt"

type animal interface {
	GetName()
}

// first class
type Cat struct {
	Name string
}

func (c Cat) GetName() {
	fmt.Println(c.Name)
}

func NewCat() animal {
	return Cat{}
}

// second class
type Dog struct {
	Keeper string
	Name   string
}

func (c Dog) GetName() {
	fmt.Println(c.Name + ", " + c.Keeper)
}

func NewDog() animal {
	return Dog{}
}
