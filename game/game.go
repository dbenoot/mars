package game

import (
	"fmt"
	"github.com/dbenoot/mars/astronaut"
	"github.com/dbenoot/mars/ship"
	"github.com/dbenoot/mars/util"
)

func NewGame() {

	a1 := astronaut.New("Jurgen", 0, 0, util.GetRand(0, 100), util.GetRand(0, 100), util.GetRand(0, 100), util.GetRand(0, 100))
	a2 := astronaut.New("Kerbal", 0, 0, util.GetRand(0, 100), util.GetRand(0, 100), util.GetRand(0, 100), util.GetRand(0, 100))
	a3 := astronaut.New("Buzz", 0, 0, util.GetRand(0, 100), util.GetRand(0, 100), util.GetRand(0, 100), util.GetRand(0, 100))

	astronauts := []astronaut.Astronaut{a1, a2, a3}

	s := ship.New("Mars Explorer", 100, 100, 100, 100)

	StartGame(s, astronauts)
}

func StartGame(s ship.Spaceship, a []astronaut.Astronaut) {
	fmt.Println(s)
	fmt.Println(a)
	fmt.Println("Game Started!")

	days := 150

	for day := 0; day < days; day++ {

		s.Fuel = s.Fuel - 1

		fmt.Println(s)

	}
}
