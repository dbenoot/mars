package game

import (
	"bufio"
	"fmt"
	"os"

	"github.com/dbenoot/mars/astronaut"
	"github.com/dbenoot/mars/ship"
	"github.com/dbenoot/mars/util"
)

func NewGame() {

	a1 := astronaut.New("Jurgen", 2, 1, util.GetRand(0, 100), util.GetRand(0, 100), util.GetRand(0, 100), util.GetRand(0, 100), 100)
	a2 := astronaut.New("Kerbal", 2, 1, util.GetRand(0, 100), util.GetRand(0, 100), util.GetRand(0, 100), util.GetRand(0, 100), 100)
	a3 := astronaut.New("Buzz", 2, 2, util.GetRand(0, 100), util.GetRand(0, 100), util.GetRand(0, 100), util.GetRand(0, 100), 100)

	astronauts := []astronaut.Astronaut{a1, a2, a3}

	s := ship.New("Mars Explorer", 5000, 5000, 250, 5000, 100)

	fmt.Println(s)
	fmt.Println(astronauts)
	fmt.Println("Game Started!")

	StartGame(s, astronauts)
}

func StartGame(s ship.Spaceship, a []astronaut.Astronaut) {

	days := util.GetRand(200, 250)

	// Game loop

	for day := 0; day < days; day++ {

		fmt.Printf("You have %v days left in transit", days-day)

		// s.Fuel = s.Fuel - 1
		s = s.Process()

		for i := len(a) - 1; i >= 0; i-- {
			a[i], s = a[i].Process(s)

			// Remove astronaut if health is <1
			if a[i].Health < 1 {
				a = append(a[:i], a[i+1:]...)
			}

		}

		fmt.Println(s)
		fmt.Println(a)

		buf := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		_, _ = buf.ReadBytes('\n')
	}
}
