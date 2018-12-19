package game

import (
	// "bufio"
	"fmt"
	// "os"

	"github.com/dbenoot/mars/astronaut"
	"github.com/dbenoot/mars/ship"
	"github.com/dbenoot/mars/util"
)

func NewGame() {

	a1 := astronaut.New("Jurgen", 2, 1, util.GetRand(0, 100), util.GetRand(0, 100), util.GetRand(0, 100), util.GetRand(0, 100), 100, true, true)
	a2 := astronaut.New("Kerbal", 2, 1, util.GetRand(0, 100), util.GetRand(0, 100), util.GetRand(0, 100), util.GetRand(0, 100), 100, true, true)
	a3 := astronaut.New("Buzz", 2, 2, util.GetRand(0, 100), util.GetRand(0, 100), util.GetRand(0, 100), util.GetRand(0, 100), 100, true, true)

	astronauts := []astronaut.Astronaut{a1, a2, a3}

	s := ship.New("Mars Explorer", 5000, 10, 250, 5000, 100)

	fmt.Println(s)
	fmt.Println(astronauts)
	fmt.Println("Game Started!")

	StartGame(s, astronauts)
}

func StartGame(s ship.Spaceship, a []astronaut.Astronaut) {
	day := 0
	days := util.GetRand(5, 10)
	var input string

	// Game loop

	for day < days {

		fmt.Printf("You have %v days left in transit", days-day)

		fmt.Println(s)
		fmt.Println(a)

		// get input

		fmt.Print("> ")

		fmt.Scan(&input)

		// process input

		switch input {

		// end the turn

		case "end":

			//process all ship actions
			s = s.Process()

			// loop over the astronauts and process all astronaut actions
			for i := len(a) - 1; i >= 0; i-- {
				a[i], s = a[i].Process(s)
				// Remove astronaut if health is <1
				if a[i].Health < 1 {
					a = append(a[:i], a[i+1:]...)
				}
			}

			// count up the day
			day++

			// check whether you have arrived and give stats

			if day == days {
				fmt.Println("You have arrived.")
				fmt.Println(s)
				fmt.Println(a)
			}

		default:
			fmt.Println("Command not found.")
		}

	}
}
