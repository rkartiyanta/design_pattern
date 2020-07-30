package structural

import "fmt"

type Soldier interface {
	Attack()
}

type Knight struct {
	Power int
}

func (c Knight) Attack() {
	fmt.Println(c.Power)
}

func NewKnight() Soldier {
	return Knight{Power: 10}
}

type Sword struct {
	Power int
}

func (c Sword) Attack() {
	fmt.Println(c.Power)
}

func NewSword() Soldier {
	return Sword{Power: 2}
}

type KnightSword struct {
	Knight
	Sword
}

func (c KnightSword) Attack() {
	fmt.Println(c.Knight.Power + c.Sword.Power)
}
