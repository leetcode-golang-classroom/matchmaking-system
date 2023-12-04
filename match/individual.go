package match

import (
	"fmt"
)

type Individual struct {
	id      int
	gender  Gender
	age     int
	intro   *Intro
	habbits Habbits
	coord   *Coord
}

func NewIndividual(ID int, gender Gender, age int, intro *Intro, habbits Habbits, coord *Coord) *Individual {
	return &Individual{
		id:      ID,
		gender:  gender,
		age:     age,
		intro:   intro,
		habbits: habbits,
		coord:   coord,
	}
}

func (i *Individual) String() string {
	return fmt.Sprintf("id:%d, gender: %v, age: %d, intro: %s\nhabbits: %s\ncoord:%s\n", i.id, i.gender, i.age, i.intro, i.habbits, i.coord)
}

func (i *Individual) GetIntro() string {
	return i.intro.String()
}
