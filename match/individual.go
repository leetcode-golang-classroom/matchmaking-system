package match

import (
	"fmt"
	"log"
)

type Individual struct {
	id                int
	gender            Gender
	age               int
	intro             *Intro
	habbits           Habbits
	coord             *Coord
	matchedIndividual *MatchedIndividual
}

func NewIndividual(ID int, gender Gender, age int, intro *Intro, habbits Habbits, coord *Coord) *Individual {
	return &Individual{
		id:                ID,
		gender:            gender,
		age:               age,
		intro:             intro,
		habbits:           habbits,
		coord:             coord,
		matchedIndividual: nil,
	}
}

func (i *Individual) String() string {
	return fmt.Sprintf("id:%d, gender: %v, age: %d, intro: %s\nhabbits: %s\ncoord:%s\n", i.id, i.gender, i.age, i.intro, i.habbits, i.coord)
}

func (i *Individual) SetMatchedIndividual(id *Individual) {
	if i.matchedIndividual != nil {
		log.Fatal("this guy has been matched")
	}
	i.matchedIndividual = NewMatchedIndividual(id)
}
func (i *Individual) ClearMatchedIndividual() {
	if i.matchedIndividual != nil {
		i.matchedIndividual = nil
	}
}
func (i *Individual) GetMatched() *MatchedIndividual {
	return i.matchedIndividual
}
func (i *Individual) GetIntro() string {
	return i.intro.String()
}
