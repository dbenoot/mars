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
}

func New(name string, water int, food int, pilot int, engineering int, maintenance int, social int, Health int) Astronaut {
	a := Astronaut{name, water, food, pilot, engineering, maintenance, social, Health}
	return a
}

func (a Astronaut) Process(s ship.Spaceship) (Astronaut, ship.Spaceship) {

	s = s.Eat(a.food)

	s = s.Drink(a.water)

	return a, s

}
