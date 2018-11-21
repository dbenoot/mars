package ship

type Spaceship struct {
	name   string
	Water  int
	Food   int
	Fuel   int
	Oxygen int
	Health int
}

func New(name string, water int, food int, fuel int, oxygen int, health int) Spaceship {
	s := Spaceship{name, water, food, fuel, oxygen, health}
	return s
}

func (s Spaceship) Process() Spaceship {
	s.Fuel = s.Fuel - 1
	return s
}

func (s Spaceship) Eat(f int) Spaceship {

	s.Food = s.Food - f

	return s
}

func (s Spaceship) Drink(f int) Spaceship {

	s.Water = s.Water - f

	return s
}
