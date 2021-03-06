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

	a1 := astronaut.New("Player", false, "Bridge", 2, 1, util.GetRand(0, 100), util.GetRand(0, 100), util.GetRand(0, 100), util.GetRand(0, 100), 100, 0, 0)
	a2 := astronaut.New("Kerbal", true, "Hub", 2, 1, util.GetRand(0, 100), util.GetRand(0, 100), util.GetRand(0, 100), util.GetRand(0, 100), 100, 0, 0)
	a3 := astronaut.New("Buzz", true, "Engineering", 2, 2, util.GetRand(0, 100), util.GetRand(0, 100), util.GetRand(0, 100), util.GetRand(0, 100), 100, 0, 0)

	astronauts := []astronaut.Astronaut{a1, a2, a3}

	s := ship.New("Mars Explorer", 5000, 750, 250, 21, 500, 100)

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

func removeSpecialism(lm map[string]ship.Location) map[string]ship.Location {
	for k, _ := range lm {

		var x = lm[k]
		x.Occupied = ""
		lm[k] = x

		lm[k].Specialist["Pilot"] = 0
		lm[k].Specialist["Engineering"] = 0
		lm[k].Specialist["Maintenance"] = 0
	}

	return lm
}

func addSpecialism(lm map[string]ship.Location, a astronaut.Astronaut) map[string]ship.Location {

	var x = lm[a.Location]
	x.Occupied = a.Name
	lm[a.Location] = x

	lm[a.Location].Specialist["Pilot"] = a.Pilot
	lm[a.Location].Specialist["Engineering"] = a.Engineering
	lm[a.Location].Specialist["Maintenance"] = a.Maintenance

	return lm
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

func printStat(s ship.Spaceship, a astronaut.Astronaut, lm map[string]ship.Location) {

	if _, ok := lm[a.Location].Subsystems["status"]; ok {
		fmt.Println("These are the ship's stats:")
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.AlignRight|tabwriter.Debug)

		if lm[a.Location].Subsystems["status"]+util.GetRand(0, 100) >= 100 {
			fmt.Fprintf(w, "\t%v\t%v\t%v\t\n", "Name", s.Name, " ")
		} else {
			fmt.Fprintf(w, "\t%v\t%v\t%v\t\n", "Name", "unable to access status", " ")
		}
		if lm[a.Location].Subsystems["status"]+util.GetRand(0, 100) >= 100 {
			fmt.Fprintf(w, "\t%v\t%v\t%v\t\n", "Water", s.Water, "L")
		} else {
			fmt.Fprintf(w, "\t%v\t%v\t%v\t\n", "Water", "unable to access status", "L")
		}
		if lm[a.Location].Subsystems["status"]+util.GetRand(0, 100) >= 100 {
			fmt.Fprintf(w, "\t%v\t%v\t%v\t\n", "Food", s.Food, "kg")
		} else {
			fmt.Fprintf(w, "\t%v\t%v\t%v\t\n", "Food", "unable to access status", "kg")
		}
		if lm[a.Location].Subsystems["status"]+util.GetRand(0, 100) >= 100 {
			fmt.Fprintf(w, "\t%v\t%v\t%v\t\n", "Fuel", s.Fuel, "m³")
		} else {
			fmt.Fprintf(w, "\t%v\t%v\t%v\t\n", "Fuel", "unable to access status", "m³")
		}
		if lm[a.Location].Subsystems["status"]+util.GetRand(0, 100) >= 100 {
			fmt.Fprintf(w, "\t%v\t%v\t%v\t\n", "Oxygen", s.Oxygen, "%")
		} else {
			fmt.Fprintf(w, "\t%v\t%v\t%v\t\n", "Oxygen", "unable to access status", "%")
		}
		if lm[a.Location].Subsystems["status"]+util.GetRand(0, 100) >= 100 {
			fmt.Fprintf(w, "\t%v\t%v\t%v\t\n", "CO2", s.CO2, "ppm")
		} else {
			fmt.Fprintf(w, "\t%v\t%v\t%v\t\n", "CO2", "unable to access status", "ppm")
		}
		if lm[a.Location].Subsystems["status"]+util.GetRand(0, 100) >= 100 {
			fmt.Fprintf(w, "\t%v\t%v\t%v\t\n", "Health", s.Health, "❤")
		} else {
			fmt.Fprintf(w, "\t%v\t%v\t%v\t\n", "Health", "unable to access status", "❤")
		}

		w.Flush()
	} else {
		fmt.Println("Status equipement not present in this location.")
	}

}

func save(day int, days int, s ship.Spaceship, a []astronaut.Astronaut, lm map[string]ship.Location) {

}

func endTurn(day int, days int, s ship.Spaceship, a []astronaut.Astronaut, lm map[string]ship.Location) (int, ship.Spaceship, []astronaut.Astronaut, map[string]ship.Location) {

	//process all ship actions
	lm = removeSpecialism(lm)
	s = s.Process(lm, len(a))

	// loop over the astronauts and process all astronaut stats
	for i := len(a) - 1; i >= 0; i-- {
		lm = addSpecialism(lm, a[i])
		a[i], s = a[i].Process(s)
		// Stop game if the commander dies
		if a[i].Health < 1 && a[i].NPC == false {
			fmt.Println("Alas, you died! Game over!") //alternatively, make another astronaut the captain?
			os.Exit(1)
		}
		// Remove astronaut if health is <1
		if a[i].Health < 1 {
			fmt.Printf("Alas, astronaut %v has died!", a[i].Name)
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

	return day, s, a, lm
}

func StartGame(s ship.Spaceship, a []astronaut.Astronaut, lm map[string]ship.Location) {
	day := 0
	days := util.GetRand(220, 250)
	var input string
	var command []string

	// Sync astronaut info with ship info

	for i := len(a) - 1; i >= 0; i-- {
		lm = addSpecialism(lm, a[i])
	}

	// Welcome the player

	fmt.Printf("Game Started!\nYou are on the spaceship %v on a %v days transit to Mars.\n", s.Name, days)

	// Game loop

	for day < days {
		fmt.Printf("Day %v of transfer, %v days to go.\n", day, days-day)
		processLocation(lm, a[0].Location, a)

		// get input

		fmt.Print("> ")
		fmt.Scan(&input)

		// process input

		command = append(command, input)

		switch input {

		// walk to other rooms

		case "1":
			if lm[a[0].Location].Occupied == a[0].Name {
				var x = lm[a[0].Location]
				x.Occupied = ""
				lm[a[0].Location] = x
			}

			if lm[lm[a[0].Location].Transitions[0]].Occupied == "" {
				var y = lm[lm[a[0].Location].Transitions[0]]
				y.Occupied = a[0].Name
				lm[lm[a[0].Location].Transitions[0]] = y
			}

			a[0].Location = lm[a[0].Location].Transitions[0]

		case "2":
			if lm[a[0].Location].Occupied == a[0].Name {
				var x = lm[a[0].Location]
				x.Occupied = ""
				lm[a[0].Location] = x
			}

			if lm[lm[a[0].Location].Transitions[1]].Occupied == "" {
				var y = lm[lm[a[0].Location].Transitions[1]]
				y.Occupied = a[0].Name
				lm[lm[a[0].Location].Transitions[1]] = y
			}

			a[0].Location = lm[a[0].Location].Transitions[1]

		case "3":
			if lm[a[0].Location].Occupied == a[0].Name {
				var x = lm[a[0].Location]
				x.Occupied = ""
				lm[a[0].Location] = x
			}

			if lm[lm[a[0].Location].Transitions[2]].Occupied == "" {
				var y = lm[lm[a[0].Location].Transitions[2]]
				y.Occupied = a[0].Name
				lm[lm[a[0].Location].Transitions[2]] = y
			}

			a[0].Location = lm[a[0].Location].Transitions[2]

		// overview of subsystems

		case "sub", "subsystem", "subsystems":
			printSub(lm, a[0])

		case "loc", "location":
			for k, _ := range a {
				fmt.Printf("Astronauts: %v - %v\n", a[k].Name, a[k].Location)
			}
			for _, value := range lm {
				fmt.Printf("Locations: %v - %v\n", value.Name, value.Occupied)
			}

		case "stat", "status":
			printStat(s, a[0], lm)

		case "talk":
			for k, _ := range a {
				if a[k].Location == a[0].Location && a[k].NPC == true {
					fmt.Printf("%v says hi.\n", a[k].Name)
				}
			}

		case "assign":

			// in the specific location, iterate over astronauts and check which NPC is present

			for k, _ := range a {
				if a[k].Location == a[0].Location && a[k].NPC == true {
					// For the astronaut that is present, show all location where they can go
					for loc, val := range lm {
						if len(val.Occupied) < 1 {
							fmt.Printf("\t- %s\n", loc)
						}
					}

					//ask the player to which location the astronaut should be moved

					var in string

					fmt.Printf("Assign %v to > ", a[k].Name)
					fmt.Scan(&in)

					// Check that the location exists and if it does, move the astronaut. Update both the location in a and the occupied in lm (this should be 1 table! TODO)

					if val, ok := lm[in]; ok {
						// clear old values
						// lm[a[k].Location].Occupied = ""
						// above does not work, hacky workaround necessary
						var x = lm[a[k].Location]
						x. Occupied = ""
						lm[a[k].Location] = x
						//also clear for the astronaut
						a[k].Location = ""

						// complete new location for the astronaut
						a[k].Location = val.Name
						// hacky workaround for lm, as the below does not work directly
						// lm[a[k].Location].Occupied = a[k].Name
						var y = lm[a[k].Location]
						y.Occupied = a[k].Name
						lm[a[k].Location] = y

						fmt.Printf("Astronaut %v moved to location %v.\n", a[k].Name, a[k].Location)
					} else {
						fmt.Printf("Location %v does not exist.\n", in)
					}

				}
			}

		// save : doesn't work yet

		case "save":
			save(day, days, s, a, lm)

		// quit

		case "exit", "quit":
			os.Exit(1)

		// end the turn

		case "end":

			day, s, a, lm = endTurn(day, days, s, a, lm)

		default:
			fmt.Println("Command not found.")
		}

	}
}
