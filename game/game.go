package game

import (
	// "bufio"
	"encoding/json"
	"fmt"
	"github.com/dbenoot/mars/astronaut"
	"github.com/dbenoot/mars/ship"
	"github.com/dbenoot/mars/util"
	"io/ioutil"
	"os"
	"text/tabwriter"
)

func NewGame() {

	a1 := astronaut.New("Player", false, "Bridge", 2, 1, util.GetRand(0, 100), util.GetRand(0, 100), util.GetRand(0, 100), util.GetRand(0, 100), 100, true, true)
	a2 := astronaut.New("Kerbal", true, "Main Hall", 2, 1, util.GetRand(0, 100), util.GetRand(0, 100), util.GetRand(0, 100), util.GetRand(0, 100), 100, true, true)
	a3 := astronaut.New("Buzz", true, "Engineering", 2, 2, util.GetRand(0, 100), util.GetRand(0, 100), util.GetRand(0, 100), util.GetRand(0, 100), 100, true, true)

	astronauts := []astronaut.Astronaut{a1, a2, a3}

	s := ship.New("Mars Explorer", 5000, 10, 250, 5000, 100)

	locationMap := loadLoc("rooms.json")

	StartGame(s, astronauts, locationMap)
}

func processLocation(lm map[string]ship.Location, location string, astronauts []astronaut.Astronaut) {
	fmt.Println("You are here: ", location, ".")
	fmt.Println(lm[location].Description)
	fmt.Println("You can go to these places:")
	for index, loc := range lm[location].Transitions {
		fmt.Printf("\t%d - %s\n", index+1, loc)
	}
	for _, a := range astronauts {
		if location == a.Location && a.NPC == true {
			processInteraction(a)
		}
	}
}

func processInteraction(a astronaut.Astronaut) {
	fmt.Println(a.Name, " is here.")
}

func loadLoc(f string) map[string]ship.Location {
	var locationMap map[string]ship.Location

	a, err := ioutil.ReadFile(f)
	util.Check(err)

	err = json.Unmarshal(a, &locationMap)
	util.Check(err)

	return locationMap
}

func printSub(lm map[string]ship.Location, a astronaut.Astronaut) {
	fmt.Println("The following subsystems are present in this module:")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.AlignRight|tabwriter.Debug)
	for k, v := range lm[a.Location].Subsystems {
		fmt.Fprintf(w, "\t%v\t%v\t\n", k, v)
	}
	w.Flush()

}

func StartGame(s ship.Spaceship, a []astronaut.Astronaut, lm map[string]ship.Location) {
	day := 0
	days := util.GetRand(5, 10)
	var input string

	fmt.Println("Game Started!")

	// Game loop

	fmt.Printf("You are on the spaceship %v on a %v days transit to Mars.\n", s.Name, days)

	for day < days {
		fmt.Printf("Day %v of transfer, %v days to go.\n", day, days-day)
		processLocation(lm, a[0].Location, a)

		// get input

		fmt.Print("> ")
		fmt.Scan(&input)

		// process input

		switch input {

		// walk to other rooms

		case "1":
			a[0].Location = lm[a[0].Location].Transitions[0]

		case "2":
			a[0].Location = lm[a[0].Location].Transitions[1]

		case "3":
			a[0].Location = lm[a[0].Location].Transitions[2]

		case "sub":
			printSub(lm, a[0])

		case "exit", "quit":
			os.Exit(1)

		case "sleep":
			//this ends the turn (=day)

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
