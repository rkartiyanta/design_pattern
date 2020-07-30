package structural

type person struct {
	name string
	age  int
}

type Person interface {
	GetName() string
	GetAge() int
}

func NewRichard() Person {
	return &person{
		name: "Richard",
		age:  20,
	}
}

func (c *person) GetName() string {
	return c.name
}

func (c *person) GetAge() int {
	return c.age
}
