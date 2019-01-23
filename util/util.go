package util

import (
	"math/rand"
	"time"
)

func GetRand(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn((max-min)+1) + min
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
