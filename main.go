package main

import (
	"fmt"
	"github.com/dbenoot/mars/game"
)

func main() {

	var c string

	fmt.Println("Mars")
	fmt.Println("1. Start a new game")
	fmt.Print("2. Quit\n> ")

	fmt.Scan(&c)

	switch c {
	case "1":
		game.NewGame()
	case "2":
		break
	default:

	}

}
