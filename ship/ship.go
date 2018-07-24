package ship

type Spaceship struct {
	name   string
	water  int
	food   int
	Fuel   int
	health int
}

func New(name string, water int, food int, fuel int, health int) Spaceship {
	s := Spaceship{name, water, food, fuel, health}
	return s
}
