package astronaut

type Astronaut struct {
	name        string
	water       int
	food        int
	pilot       int
	engineering int
	maintenance int
	social      int
}

func New(name string, water int, food int, pilot int, engineering int, maintenance int, social int) Astronaut {
	a := Astronaut{name, water, food, pilot, engineering, maintenance, social}
	return a
}
