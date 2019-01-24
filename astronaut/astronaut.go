package astronaut

import (
	"github.com/dbenoot/mars/ship"
	// "github.com/dbenoot/mars/util"
	"fmt"
)

type Astronaut struct {
	Name        string
	NPC         bool
	Location    string
	water       int
	food        int
	pilot       int
	engineering int
	maintenance int
	social      int
	Health      int
	fed         int
	drank       int
}

func New(Name string, npc bool, Location string, water int, food int, pilot int, engineering int, maintenance int, social int, Health int, fed int, drank int) Astronaut {
	a := Astronaut{Name, npc, Location, water, food, pilot, engineering, maintenance, social, Health, fed, drank}
	return a
}

func (a Astronaut) Process(s ship.Spaceship) (Astronaut, ship.Spaceship) {

	// eat
	s, fed := s.Eat(a.food)

	// drink
	s, drank := s.Drink(a.water)

	// calculate health

	if fed == true {
		a.Health = a.Health + 1
		a.fed = 0
	} else {
		a.fed = a.fed + 1
		a.Health = a.Health - a.fed
	}
	if drank == true {
		a.Health = a.Health + 1
		a.drank = 0
	} else {
		a.drank = a.drank + 10
		a.Health = a.Health - a.drank
	}
	if a.Health > 100 {
		a.Health = 100
	}

	return a, s

}
