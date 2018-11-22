package astronaut

import (
	"github.com/dbenoot/mars/ship"
	// "github.com/dbenoot/mars/util"
)

type Astronaut struct {
	name        string
	water       int
	food        int
	pilot       int
	engineering int
	maintenance int
	social      int
	Health      int
	fed         bool
	drank       bool
}

func New(name string, water int, food int, pilot int, engineering int, maintenance int, social int, Health int, fed bool, drank bool) Astronaut {
	a := Astronaut{name, water, food, pilot, engineering, maintenance, social, Health, fed, drank}
	return a
}

func (a Astronaut) Process(s ship.Spaceship) (Astronaut, ship.Spaceship) {

	// eat
	s, a.fed = s.Eat(a.food)

	// drink
	s, a.drank = s.Drink(a.water)

	// calculate health
	// for now we do fed +5 health and drank +5; not fed -1 and not drank -10

	if a.fed == true {
		a.Health = a.Health + 5
	} else {
		a.Health = a.Health - 7
	}
	if a.drank == true {
		a.Health = a.Health + 5
	} else {
		a.Health = a.Health - 7
	}
	if a.Health > 100 {
		a.Health = 100
	}

	return a, s

}
